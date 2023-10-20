package config

import (
	"time"
)

var (
	Server = map[string]string{
		"name":      "meshplay-meshsync",
		"port":      "11000",
		"version":   "latest",
		"startedat": time.Now().String(),
	}

	Pipelines = map[string]PipelineConfigs{
		GlobalResourceKey: []PipelineConfig{
			// Core Resources
			{
				Name:      "namespaces.v1.",
				PublishTo: "meshplay.meshsync.core",
			},
			{
				Name:      "configmaps.v1.",
				PublishTo: "meshplay.meshsync.core",
			},
			{
				Name:      "nodes.v1.",
				PublishTo: "meshplay.meshsync.core",
			},
			{
				Name:      "secrets.v1.",
				PublishTo: "meshplay.meshsync.core",
			},
			{
				Name:      "persistentvolumes.v1.",
				PublishTo: "meshplay.meshsync.core",
			},
			{
				Name:      "persistentvolumeclaims.v1.",
				PublishTo: "meshplay.meshsync.core",
			},
		},
		LocalResourceKey: []PipelineConfig{
			// Core Resources
			{
				Name:      "replicasets.v1.apps",
				PublishTo: "meshplay.meshsync.core",
			},
			{
				Name:      "pods.v1.",
				PublishTo: "meshplay.meshsync.core",
			},
			{
				Name:      "services.v1.",
				PublishTo: "meshplay.meshsync.core",
			},
			{
				Name:      "deployments.v1.apps",
				PublishTo: "meshplay.meshsync.core",
			},
			{
				Name:      "statefulsets.v1.apps",
				PublishTo: "meshplay.meshsync.core",
			},
			{
				Name:      "daemonsets.v1.apps",
				PublishTo: "meshplay.meshsync.core",
			},
			//Added Ingress support
			{
				Name:      "ingresses.v1.networking.k8s.io",
				PublishTo: "meshplay.meshsync.core",
			},
			//Added endpoint support
			{
				Name:      "endpoints.v1.",
				PublishTo: "meshplay.meshsync.core",
			},
			//Added endpointslice support
			{
				Name:      "endpointslices.v1.discovery.k8s.io",
				PublishTo: "meshplay.meshsync.core",
			},
			//Added cronJob support
			{
				Name:      "cronjobs.v1.batch",
				PublishTo: "meshplay.meshsync.core",
			},
			//Added ReplicationController support
			{
				Name:      "replicationcontrollers.v1.",
				PublishTo: "meshplay.meshsync.core",
			},
			//Added storageClass support
			{
				Name:      "storageclasses.v1.storage.k8s.io",
				PublishTo: "meshplay.meshsync.core",
			},
			//Added ClusterRole support
			{
				Name:      "clusterroles.v1.rbac.authorization.k8s.io",
				PublishTo: "meshplay.meshsync.core",
			},
			//Added VolumeAttachment support
			{
				Name:      "volumeattachments.v1.storage.k8s.io",
				PublishTo: "meshplay.meshsync.core",
			},
			//Added apiservice support
			{
				Name:      "apiservices.v1.apiregistration.k8s.io",
				PublishTo: "meshplay.meshsync.core",
			},
			// Istio Resources
			// {
			// 	Name:      "virtualservices.v1beta1.networking.istio.io",
			// 	PublishTo: "meshplay.meshsync.istio",
			// },
			// {
			// 	Name:      "gateways.v1beta1.networking.istio.io",
			// 	PublishTo: "meshplay.meshsync.istio",
			// },
			// {
			// 	Name:      "destinationrules.v1beta1.networking.istio.io",
			// 	PublishTo: "meshplay.meshsync.istio",
			// },
		},
	}

	Listeners = map[string]ListenerConfig{
		LogStream: {
			Name:           LogStream,
			ConnectionName: "meshsync-logstream",
			PublishTo:      "meshplay.meshsync.logs",
		},
		ExecShell: {
			Name:           ExecShell,
			ConnectionName: "meshsync-exec",
			PublishTo:      "meshplay.meshsync.exec",
		},
		RequestStream: {
			Name:           RequestStream,
			ConnectionName: "meshsync-request-stream",
			SubscribeTo:    "meshplay.meshsync.request",
		},
	}
)
