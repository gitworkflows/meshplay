FROM golang:1.21 as meshplay-server
ARG TOKEN
ARG GIT_VERSION
ARG GIT_COMMITSHA
ARG RELEASE_CHANNEL

RUN adduser --disabled-login appuser
WORKDIR /github.com/khulnasoft/meshplay
ADD . .
RUN go clean -modcache; cd server; cd cmd; GOPROXY=https://proxy.golang.org GOSUMDB=off go build -ldflags="-w -s -X main.globalTokenForAnonymousResults=$TOKEN -X main.version=$GIT_VERSION -X main.commitsha=$GIT_COMMITSHA -X main.releasechannel=$RELEASE_CHANNEL" -tags draft -a -o /meshplay .

FROM node:lts as ui
ADD ui ui
RUN cd ui; npm install --omit-dev; npm run build && npm run export; mv out /; mv public /

FROM node:lts as provider-ui
ADD provider-ui provider-ui
RUN cd provider-ui; npm install --omit=dev; npm run build && npm run export; mv out /

FROM ubuntu as wrk2
RUN apt-get -y update && apt-get -y install git && apt-get clean && rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/* /usr/share/man/?? /usr/share/man/??_*
RUN apt-get -y update && apt-get -y  install build-essential libssl-dev git zlib1g-dev
RUN git config --global user.email "meshplay@khulnasoft.com"
RUN git config --global user.name "meshplay"
RUN git clone --depth=1 https://github.com/khulnasoft/meshplay/wrk2 && cd wrk2 && make

FROM alpine:3.15.3 as seed_content
RUN apk add --no-cache curl
WORKDIR /
RUN lines=$(curl -s https://api.github.com/repos/khulnasoft/wasm-filters/releases/latest | grep "browser_download_url.*wasm" | cut -d : -f 2,3 | sed 's/"//g') \
    && mkdir -p seed_content/filters/binaries \ 
    && cd seed_content/filters/binaries  \
    for line in $lines \
    do \
    curl -LO $line \
    done 

# bundling patterns
RUN curl -L -s https://github.com/service-mesh-patterns/service-mesh-patterns/tarball/master -o service-mesh-patterns.tgz \
    && mkdir service-mesh-patterns \
    && mkdir -p /seed_content/patterns \
    && tar xzf service-mesh-patterns.tgz --directory=service-mesh-patterns \
    && mv service-mesh-patterns/*/samples/* /seed_content/patterns/

# bundling applications
RUN mkdir -p /seed_content/applications && cd /seed_content/applications \
    && curl -LO https://raw.githubusercontent.com/istio/istio/master/samples/bookinfo/platform/kube/bookinfo.yaml \
    && curl -LO https://raw.githubusercontent.com/istio/istio/master/samples/httpbin/httpbin.yaml \
    && curl -L https://raw.githubusercontent.com/khulnasoft/image-hub/master/deployment.yaml -o imagehub.yaml \
    && mkdir /emojivoto && (cd /emojivoto && curl --remote-name-all -L https://raw.githubusercontent.com/BuoyantIO/emojivoto/main/kustomize/deployment/emoji.yml \
    https://raw.githubusercontent.com/BuoyantIO/emojivoto/main/kustomize/deployment/vote-bot.yml \
    https://raw.githubusercontent.com/BuoyantIO/emojivoto/main/kustomize/deployment/voting.yml \
    https://raw.githubusercontent.com/BuoyantIO/emojivoto/main/kustomize/deployment/web.yml) \
    && awk 'FNR==1 && NR>1 { printf("\n%s\n\n","---") } 1' /emojivoto/*.yml > /seed_content/applications/emojivoto.yml

#FROM ubuntu as nighthawk
#RUN apt-get -y update && apt-get -y install git && apt-get clean && rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/* /usr/share/man/?? /usr/share/man/??_*
#RUN apt-get -y update && apt-get -y  install build-essential libssl-dev git zlib1g-dev
#RUN git config --global user.email "meshplay@khulnasoft.com"
#RUN git config --global user.name "meshplay"
#RUN git clone https://github.com/khulnasoft/meshplay/getnighthawk
#RUN cd getnighthawk/apinighthawk/bin && chmod +x ./nighthawk_client


FROM alpine:3.15.3
ENV GLIBC_REPO=https://github.com/sgerrand/alpine-pkg-glibc
ENV GLIBC_VERSION=2.34-r0
# To solve: ./meshplay: /usr/glibc-compat/lib/libc.so.6: version `GLIBC_2.34' not found (required by ./meshplay)
RUN set -ex && \
    apk --update add libstdc++ curl ca-certificates && \
    for pkg in glibc-${GLIBC_VERSION} glibc-bin-${GLIBC_VERSION}; \
    do curl -sSL ${GLIBC_REPO}/releases/download/${GLIBC_VERSION}/${pkg}.apk -o /tmp/${pkg}.apk; done && \
    apk add --allow-untrusted /tmp/*.apk && \
    rm -v /tmp/*.apk && \
    /usr/glibc-compat/sbin/ldconfig /lib /usr/glibc-compat/lib
RUN update-ca-certificates

# FROM frolvlad/alpine-glibc:alpine-3.13_glibc-2.32
#RUN apt-get update; apt-get install -y ca-certificates; update-ca-certificates && rm -rf /var/lib/apt/lists/*
# RUN apk update && apk add ca-certificates; update-ca-certificates && rm -rf /var/cache/apk/*
# RUN apk upgrade --no-cache && \
#     apk add --no-cache libstdc++

COPY ./server/meshmodel /app/server/meshmodel
COPY --from=meshplay-server /meshplay /app/server/cmd/
COPY --from=meshplay-server /etc/passwd /etc/passwd
COPY --from=meshplay-server /github.com/khulnasoft/meshplay/server/helpers/swagger.yaml /app/server/helpers/swagger.yaml
COPY --from=ui /out /app/ui/out
COPY --from=ui /public /app/ui/public
COPY --from=provider-ui /out /app/provider-ui/out
COPY --from=wrk2 /wrk2 /app/server/cmd/wrk2
COPY --from=wrk2 /wrk2/wrk /usr/local/bin
COPY --from=seed_content /seed_content /home/appuser/.meshplay/seed_content
COPY --from=khulnasoft/khulnasoft/meshplay/wrk2latest /usr/local/bin/nighthawk_service /app/server/cmd/
COPY --from=khulnasoft/khulnasoft/meshplay/wrk2latest /usr/local/bin/nighthawk_output_transform /app/server/cmd/

RUN addgroup docker
RUN mkdir -p /var/run; chown -R appuser:docker /var/run/
RUN mkdir -p /home/appuser/.meshplay/config; chown -R appuser /home/appuser/; chown -R appuser /app/ui/public
USER appuser
WORKDIR /app/server/cmd
CMD ./meshplay
