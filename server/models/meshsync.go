package models

import (
	"github.com/khulnasoft/meshplay/meshsync/pkg/model"
)

type MeshSyncResourcesAPIResponse struct {
	Page       int             `json:"page"`
	PageSize   int             `json:"page_size"`
	TotalCount int64             `json:"total_count"`
	Resources   []model.KubernetesResource `json:"resources"`
}