#-----------------------------------------------------------------------------
# Meshplay Playground Production Environment
#-----------------------------------------------------------------------------

DEBUG=true
PROD_PLAYGROUND_PORT=80
REMOTE_PROVIDER=https://meshery.layer5.io

#-----------------------------------------------------------------------------
# Meshplay Playground Production Deployment
#-----------------------------------------------------------------------------
.PHONY: play-prod-run play-prod-stop 

## Deploys Meshplay Playground in Prod
play-prod-run:
	(docker rm -f meshery) || true
	docker run -d --name meshery \
	-e PROVIDER_BASE_URLS=$(REMOTE_PROVIDER) \
	-e DEBUG=$(DEBUG) \
	-e PLAYGROUND=true \
	-e MESHPLAY_SERVER_CALLBACK_URL=https://playground.khulnasoft.com/api/user/token \ 
	-p $(PROD_PLAYGROUND_PORT):8080 \
	layer5/meshery:playground-latest ./meshery
	
# docker run -d --name meshery \
# 	-e PROVIDER_BASE_URLS=https://meshery.layer5.io \
# 	-e PLAYGROUND=true \
# 	-p 80:8080 \
# 	layer5/meshery:playground-latest ./meshery

play-prod-stop:
	docker rm -f meshery