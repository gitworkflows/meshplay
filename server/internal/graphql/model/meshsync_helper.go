package model

import (
	"github.com/khulnasoft/meshplay/meshkit/broker"
	"github.com/khulnasoft/meshplay/meshkit/database"
	meshsyncmodel "github.com/khulnasoft/meshplay/meshsync/pkg/model"

	"gorm.io/gorm"
)

func recordMeshSyncData(eventtype broker.EventType, handler *database.Handler, object *meshsyncmodel.KubernetesResource) error {
	if handler == nil {
		return ErrEmptyHandler
	}

	handler.Lock()
	defer handler.Unlock()

	switch eventtype {
	case broker.Add, broker.Update:
		result := handler.Create(object)
		if result.Error != nil {
			result = handler.Session(&gorm.Session{FullSaveAssociations: true}).Updates(object)
			if result.Error != nil {
				return ErrCreateData(result.Error)
			}
		}
	case broker.Delete:
		result := handler.Delete(object)
		if result.Error != nil {
			return ErrDeleteData(result.Error)
		}
	case broker.ErrorEvent:
		return nil
	}

	return nil
}
