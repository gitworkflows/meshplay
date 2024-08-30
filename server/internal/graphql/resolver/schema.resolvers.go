package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.45

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/gofrs/uuid"
	"github.com/khulnasoft/meshkit/broker"
	"github.com/khulnasoft/meshkit/utils"
	"github.com/khulnasoft/meshplay/server/handlers"
	"github.com/khulnasoft/meshplay/server/internal/graphql/generated"
	"github.com/khulnasoft/meshplay/server/internal/graphql/model"
	"github.com/khulnasoft/meshplay/server/machines/kubernetes"
	"github.com/khulnasoft/meshplay/server/models"
)

// ChangeOperatorStatus is the resolver for the changeOperatorStatus field.
func (r *mutationResolver) ChangeOperatorStatus(ctx context.Context, input *model.OperatorStatusInput) (model.Status, error) {
	provider := ctx.Value(models.ProviderCtxKey).(models.Provider)
	return r.changeOperatorStatus(ctx, provider, input.TargetStatus, input.ContextID)
}

// ChangeAdapterStatus is the resolver for the changeAdapterStatus field.
func (r *mutationResolver) ChangeAdapterStatus(ctx context.Context, input *model.AdapterStatusInput) (model.Status, error) {
	provider := ctx.Value(models.ProviderCtxKey).(models.Provider)
	return r.changeAdapterStatus(ctx, provider, input.TargetStatus, input.Adapter, input.TargetPort)
}

// GetAvailableAddons is the resolver for the getAvailableAddons field.
func (r *queryResolver) GetAvailableAddons(ctx context.Context, filter *model.ServiceMeshFilter) ([]*model.AddonList, error) {
	provider := ctx.Value(models.ProviderCtxKey).(models.Provider)
	if filter != nil {
		return r.getAvailableAddons(ctx, provider, filter)
	}

	return nil, ErrInvalidRequest
}

// GetControlPlanes is the resolver for the getControlPlanes field.
func (r *queryResolver) GetControlPlanes(ctx context.Context, filter *model.ServiceMeshFilter) ([]*model.ControlPlane, error) {
	provider := ctx.Value(models.ProviderCtxKey).(models.Provider)
	if filter != nil {
		return r.getControlPlanes(ctx, provider, filter)
	}

	return nil, ErrInvalidRequest
}

// GetDataPlanes is the resolver for the getDataPlanes field.
func (r *queryResolver) GetDataPlanes(ctx context.Context, filter *model.ServiceMeshFilter) ([]*model.DataPlane, error) {
	provider := ctx.Value(models.ProviderCtxKey).(models.Provider)
	if filter != nil {
		return r.getDataPlanes(ctx, provider, filter)
	}

	return nil, ErrInvalidRequest
}

// GetOperatorStatus is the resolver for the getOperatorStatus field.
func (r *queryResolver) GetOperatorStatus(ctx context.Context, connectionID string) (*model.MeshplayControllersStatusListItem, error) {
	provider := ctx.Value(models.ProviderCtxKey).(models.Provider)
	return r.getOperatorStatus(ctx, provider, connectionID)
}

// ResyncCluster is the resolver for the resyncCluster field.
func (r *queryResolver) ResyncCluster(ctx context.Context, selector *model.ReSyncActions, k8scontextID string) (model.Status, error) {
	provider := ctx.Value(models.ProviderCtxKey).(models.Provider)
	return r.resyncCluster(ctx, provider, selector, k8scontextID)
}

// GetMeshsyncStatus is the resolver for the getMeshsyncStatus field.
func (r *queryResolver) GetMeshsyncStatus(ctx context.Context, connectionID string) (*model.OperatorControllerStatus, error) {
	provider := ctx.Value(models.ProviderCtxKey).(models.Provider)
	return r.getMeshsyncStatus(ctx, provider, connectionID)
}

// GetNatsStatus is the resolver for the getNatsStatus field.
func (r *queryResolver) GetNatsStatus(ctx context.Context, connectionID string) (*model.OperatorControllerStatus, error) {
	provider := ctx.Value(models.ProviderCtxKey).(models.Provider)
	return r.getNatsStatus(ctx, provider, connectionID)
}

// GetAvailableNamespaces is the resolver for the getAvailableNamespaces field.
func (r *queryResolver) GetAvailableNamespaces(ctx context.Context, k8sClusterIDs []string) ([]*model.NameSpace, error) {
	provider := ctx.Value(models.ProviderCtxKey).(models.Provider)
	return r.getAvailableNamespaces(ctx, provider, k8sClusterIDs)
}

// GetPerfResult is the resolver for the getPerfResult field.
func (r *queryResolver) GetPerfResult(ctx context.Context, id string) (*model.MeshplayResult, error) {
	provider := ctx.Value(models.ProviderCtxKey).(models.Provider)
	return r.getPerfResult(ctx, provider, id)
	// panic(fmt.Errorf("not implemented"))
}

// FetchResults is the resolver for the fetchResults field.
func (r *queryResolver) FetchResults(ctx context.Context, selector model.PageFilter, profileID string) (*model.PerfPageResult, error) {
	provider := ctx.Value(models.ProviderCtxKey).(models.Provider)
	return r.fetchResults(ctx, provider, selector, profileID)
}

// GetPerformanceProfiles is the resolver for the getPerformanceProfiles field.
func (r *queryResolver) GetPerformanceProfiles(ctx context.Context, selector model.PageFilter) (*model.PerfPageProfiles, error) {
	provider := ctx.Value(models.ProviderCtxKey).(models.Provider)
	return r.getPerformanceProfiles(ctx, provider, selector)
}

// FetchAllResults is the resolver for the fetchAllResults field.
func (r *queryResolver) FetchAllResults(ctx context.Context, selector model.PageFilter) (*model.PerfPageResult, error) {
	provider := ctx.Value(models.ProviderCtxKey).(models.Provider)
	return r.fetchAllResults(ctx, provider, selector)
}

// FetchPatterns is the resolver for the fetchPatterns field.
func (r *queryResolver) FetchPatterns(ctx context.Context, selector model.PageFilter) (*model.PatternPageResult, error) {
	provider := ctx.Value(models.ProviderCtxKey).(models.Provider)
	return r.fetchPatterns(ctx, provider, selector)
}

// GetKubectlDescribe is the resolver for the getKubectlDescribe field.
func (r *queryResolver) GetKubectlDescribe(ctx context.Context, name string, kind string, namespace string) (*model.KctlDescribeDetails, error) {
	return r.getKubectlDescribe(ctx, name, kind, namespace)
}

// FetchPatternCatalogContent is the resolver for the fetchPatternCatalogContent field.
func (r *queryResolver) FetchPatternCatalogContent(ctx context.Context, selector *model.CatalogSelector) ([]*model.CatalogPattern, error) {
	provider := ctx.Value(models.ProviderCtxKey).(models.Provider)
	return r.fetchCatalogPattern(ctx, provider, selector)
}

// FetchFilterCatalogContent is the resolver for the fetchFilterCatalogContent field.
func (r *queryResolver) FetchFilterCatalogContent(ctx context.Context, selector *model.CatalogSelector) ([]*model.CatalogFilter, error) {
	provider := ctx.Value(models.ProviderCtxKey).(models.Provider)
	return r.fetchCatalogFilter(ctx, provider, selector)
}

// GetMeshModelSummary is the resolver for the getMeshModelSummary field.
func (r *queryResolver) GetMeshModelSummary(ctx context.Context, selector model.MeshModelSummarySelector) (*model.MeshModelSummary, error) {
	provider := ctx.Value(models.ProviderCtxKey).(models.Provider)
	return r.getMeshModelSummary(ctx, provider, selector)
}

// FetchTelemetryComponents is the resolver for the fetchTelemetryComponents field.
func (r *queryResolver) FetchTelemetryComponents(ctx context.Context, contexts []string) ([]*model.TelemetryComp, error) {
	provider := ctx.Value(models.ProviderCtxKey).(models.Provider)
	return r.getTelemetryComps(ctx, provider, contexts)
}

// SubscribePerfProfiles is the resolver for the subscribePerfProfiles field.
func (r *subscriptionResolver) SubscribePerfProfiles(ctx context.Context, selector model.PageFilter) (<-chan *model.PerfPageProfiles, error) {
	provider := ctx.Value(models.ProviderCtxKey).(models.Provider)
	return r.subscribePerfProfiles(ctx, provider, selector)
}

// SubscribePerfResults is the resolver for the subscribePerfResults field.
func (r *subscriptionResolver) SubscribePerfResults(ctx context.Context, selector model.PageFilter, profileID string) (<-chan *model.PerfPageResult, error) {
	provider := ctx.Value(models.ProviderCtxKey).(models.Provider)
	return r.subscribePerfResults(ctx, provider, selector, profileID)
}

// SubscribeMeshplayControllersStatus is the resolver for the subscribeMeshplayControllersStatus field.
func (r *subscriptionResolver) SubscribeMeshplayControllersStatus(ctx context.Context, connectionIDs []string) (<-chan []*model.MeshplayControllersStatusListItem, error) {
	resChan := make(chan []*model.MeshplayControllersStatusListItem)
	handler, ok := ctx.Value(models.HandlerKey).(*handlers.Handler)

	if !ok {
		er := model.ErrMeshplayControllersStatusSubscription(fmt.Errorf("controller handlers are not configured for any of the contexts"))
		r.Log.Error(er)
		return nil, er
	}

	statusMapPerConnection := make(map[string]map[models.MeshplayController]models.MeshplayControllerStatusAndVersion)
	// initialize the map

	for _, connectionID := range connectionIDs {
		inst, ok := handler.ConnectionToStateMachineInstanceTracker.Get(uuid.FromStringOrNil(connectionID))
		if ok && inst != nil {
			machinectx, err := utils.Cast[*kubernetes.MachineCtx](inst.Context)
			if err != nil {
				r.Log.Error(model.ErrMeshplayControllersStatusSubscription(err))
				continue
			}
			ctrlHandlers := machinectx.MeshplayCtrlsHelper.GetControllerHandlersForEachContext()
			for controller, handler := range ctrlHandlers {
				if _, ok := statusMapPerConnection[connectionID]; !ok {
					statusMapPerConnection[connectionID] = make(map[models.MeshplayController]models.MeshplayControllerStatusAndVersion)
				}
				version, err := handler.GetVersion()
				if err != nil {
					er := model.ErrMeshplayControllersStatusSubscription(err)
					r.Log.Error(er)
				}

				statusMapPerConnection[connectionID][controller] = models.MeshplayControllerStatusAndVersion{
					Status:  handler.GetStatus(),
					Version: version,
				}
			}

		}
	}
	go func() {
		ctrlsStatusList := make([]*model.MeshplayControllersStatusListItem, 0)
		// first send the initial status of the controllers
		for connectionID, controllerMap := range statusMapPerConnection {
			for controller, statusAndVersion := range controllerMap {
				ctrlsStatusList = append(ctrlsStatusList, &model.MeshplayControllersStatusListItem{
					ConnectionID: connectionID,
					Controller:   model.GetInternalController(controller),
					Status:       model.GetInternalControllerStatus(statusAndVersion.Status),
					Version:      statusAndVersion.Version,
				})
			}
		}
		resChan <- ctrlsStatusList
		ctrlsStatusList = make([]*model.MeshplayControllersStatusListItem, 0)
		// do this every 5 seconds
		for {
			for _, connectionID := range connectionIDs {
				inst, ok := handler.ConnectionToStateMachineInstanceTracker.Get(uuid.FromStringOrNil(connectionID))
				if ok && inst != nil {
					machinectx, err := utils.Cast[*kubernetes.MachineCtx](inst.Context)
					if err != nil {
						r.Log.Error(model.ErrMeshplayControllersStatusSubscription(err))
						continue
					}
					ctrlHandlers := machinectx.MeshplayCtrlsHelper.GetControllerHandlersForEachContext()
					for controller, handler := range ctrlHandlers {
						newStatus := handler.GetStatus()

						version, err := handler.GetVersion()
						if err != nil {
							er := model.ErrMeshplayControllersStatusSubscription(err)
							r.Log.Error(er)
						}
						// if the status has changed, send that to the subscription
						if newStatus != statusMapPerConnection[connectionID][controller].Status {
							ctrlsStatusList = append(ctrlsStatusList, &model.MeshplayControllersStatusListItem{
								ConnectionID: connectionID,
								Controller:   model.GetInternalController(controller),
								Status:       model.GetInternalControllerStatus(newStatus),
								Version:      version,
							})
							resChan <- ctrlsStatusList
						}
						// update the status list with newStatus
						if _, ok := statusMapPerConnection[connectionID][controller]; ok {
							statusMapPerConnection[connectionID][controller] = models.MeshplayControllerStatusAndVersion{
								Status:  newStatus,
								Version: version,
							}
						}
						ctrlsStatusList = make([]*model.MeshplayControllersStatusListItem, 0)
					}
				}
			}
			// establish a watch connection to get updates, ideally in meshplay-operator
			time.Sleep(time.Second * 5)
		}
	}()
	return resChan, nil
}

// SubscribeMeshSyncEvents is the resolver for the subscribeMeshSyncEvents field.
func (r *subscriptionResolver) SubscribeMeshSyncEvents(ctx context.Context, connectionIDs []string, eventTypes []model.MeshSyncEventType) (<-chan *model.MeshSyncEvent, error) {
	resChan := make(chan *model.MeshSyncEvent)
	isSubscriptionFlushed := false
	brokerEventTypes := model.GetMeshplayBrokerEventTypesFromArray(eventTypes)

	handler, ok := ctx.Value(models.HandlerKey).(*handlers.Handler)
	if !ok {
		er := model.ErrMeshSyncEventsSubscription(fmt.Errorf("meshsync data handlers are not configured for any of the contexts"))
		r.Log.Error(er)
		return nil, er
	}

	for _, connectionID := range connectionIDs {
		inst, ok := handler.ConnectionToStateMachineInstanceTracker.Get(uuid.FromStringOrNil(connectionID))
		if ok && inst != nil {
			machinectx, err := utils.Cast[*kubernetes.MachineCtx](inst.Context)
			if err != nil {
				r.Log.Error(model.ErrMeshplayControllersStatusSubscription(err))
				continue
			}

			dataHandler := machinectx.MeshplayCtrlsHelper.GetMeshSyncDataHandlersForEachContext()
			if dataHandler == nil {
				r.Log.Info("skipping meshsync events subscription for connection Id: %s", connectionID)
				r.Log.Info("connection to broker and datahandler is not yet initialised")
				continue
			}
			brokerEventsChan := make(chan *broker.Message)
			err = dataHandler.ListenToMeshSyncEvents(brokerEventsChan)
			if err != nil {
				r.Log.Warn(err)
				r.Log.Info("skipping meshsync events subscription for connection Id: %s", connectionID)
				continue
			}
			go func(connectionID string, brokerEventsChan chan *broker.Message) {
				publishHandlerWithProcessing := processAndRateLimitTheResponseOnGqlChannel(handler.MeshsyncChannel, resChan, r, 5*time.Second)
				for event := range brokerEventsChan {
					if event.EventType == broker.ErrorEvent || isSubscriptionFlushed { // better close the parent channel, but it is throwing panic
						// TODO: Handle errors accordingly
						continue
					}

					// skip event that UI doesn't want to listen to
					if !model.CheckIfBrokerEventExistsInArray(event.EventType, brokerEventTypes) {
						continue
					}

					// handle the events
					res := &model.MeshSyncEvent{
						ConnectionID: connectionID,
						Type:         string(event.EventType),
						Object:       event.Object,
					}
					publishHandlerWithProcessing(res)
				}
			}(connectionID, brokerEventsChan)
		}
	}

	// handle subscription dispose
	go func() {
	loop:
		for {
			select {
			case <-ctx.Done(): // executes when the subscription is dumped
				isSubscriptionFlushed = true
				break loop
			default:
			}
		}
	}()

	return resChan, nil
}

// SubscribeConfiguration is the resolver for the subscribeConfiguration field.
func (r *subscriptionResolver) SubscribeConfiguration(ctx context.Context, patternSelector model.PageFilter, filterSelector model.PageFilter) (<-chan *model.ConfigurationPage, error) {
	provider := ctx.Value(models.ProviderCtxKey).(models.Provider)
	user := ctx.Value(models.UserCtxKey).(*models.User)
	return r.subscribeConfiguration(ctx, provider, *user, patternSelector, filterSelector)
}

// SubscribeClusterResources is the resolver for the subscribeClusterResources field.
func (r *subscriptionResolver) SubscribeClusterResources(ctx context.Context, k8scontextIDs []string, namespace string) (<-chan *model.ClusterResources, error) {
	provider := ctx.Value(models.ProviderCtxKey).(models.Provider)
	return r.subscribeClusterResources(ctx, provider, k8scontextIDs, namespace)
}

// SubscribeK8sContext is the resolver for the subscribeK8sContext field.
func (r *subscriptionResolver) SubscribeK8sContext(ctx context.Context, selector model.PageFilter) (<-chan *model.K8sContextsPage, error) {
	provider := ctx.Value(models.ProviderCtxKey).(models.Provider)
	return r.subscribeK8sContexts(ctx, provider, selector)
}

// SubscribeMeshModelSummary is the resolver for the subscribeMeshModelSummary field.
func (r *subscriptionResolver) SubscribeMeshModelSummary(ctx context.Context, selector model.MeshModelSummarySelector) (<-chan *model.MeshModelSummary, error) {
	provider := ctx.Value(models.ProviderCtxKey).(models.Provider)
	return r.subscribeMeshModelSummary(ctx, provider, selector)
}

// SubscribeEvents is the resolver for the subscribeEvents field.
func (r *subscriptionResolver) SubscribeEvents(ctx context.Context) (<-chan *model.Event, error) {
	provider := ctx.Value(models.ProviderCtxKey).(models.Provider)
	user := ctx.Value(models.UserCtxKey).(*models.User)
	return r.eventsResolver(ctx, provider, *user)
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Subscription returns generated.SubscriptionResolver implementation.
func (r *Resolver) Subscription() generated.SubscriptionResolver { return &subscriptionResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type subscriptionResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//   - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//     it when you're done.
//   - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *queryResolver) GetClusterResources(ctx context.Context, k8scontextIDs []string, namespace string) (*model.ClusterResources, error) {
	provider := ctx.Value(models.ProviderCtxKey).(models.Provider)
	return r.getClusterResources(ctx, provider, k8scontextIDs, namespace)
}
func (r *subscriptionResolver) ListenToAddonState(ctx context.Context, filter *model.ServiceMeshFilter) (<-chan []*model.AddonList, error) {
	provider := ctx.Value(models.ProviderCtxKey).(models.Provider)
	if filter != nil {
		return r.listenToAddonState(ctx, provider, filter)
	}

	return nil, ErrInvalidRequest
}
func (r *subscriptionResolver) ListenToControlPlaneState(ctx context.Context, filter *model.ServiceMeshFilter) (<-chan []*model.ControlPlane, error) {
	provider := ctx.Value(models.ProviderCtxKey).(models.Provider)
	if filter != nil {
		return r.listenToControlPlaneState(ctx, provider, filter)
	}

	return nil, ErrInvalidRequest
}
func (r *subscriptionResolver) ListenToDataPlaneState(ctx context.Context, filter *model.ServiceMeshFilter) (<-chan []*model.DataPlane, error) {
	provider := ctx.Value(models.ProviderCtxKey).(models.Provider)
	if filter != nil {
		return r.listenToDataPlaneState(ctx, provider, filter)
	}

	return nil, ErrInvalidRequest
}
func processAndRateLimitTheResponseOnGqlChannel(meshsyncChan chan struct{}, publishChannel chan *model.MeshSyncEvent, r *subscriptionResolver, d time.Duration) func(meshsyncEvent *model.MeshSyncEvent) {
	shouldWait := false
	type syncedProcessMap struct {
		mu         sync.Mutex
		processMap map[string]*model.MeshSyncEvent
	}

	processMap := syncedProcessMap{processMap: make(map[string]*model.MeshSyncEvent)}

	isLast := false

	return func(meshsyncEvent *model.MeshSyncEvent) {
		// create a key to uniquely identify meshsync objects with its type, purpose, ctx and resource uniqueId
		var key string
		key += meshsyncEvent.Type
		key += (meshsyncEvent.Object).(map[string]interface{})["kind"].(string)
		key += meshsyncEvent.ConnectionID
		metadata := (meshsyncEvent.Object).(map[string]interface{})["metadata"]
		// the metadata.uid could alone be used as key, but has a danger that it may not be avaiable
		// this uid is coming from the ObjectMeta of the resource, but it is possible that modified and Add or Delete may come under one key, which isn't desired
		if metadata != nil && metadata != "" && metadata.(map[string]interface{})["uid"] != nil {
			key += metadata.(map[string]interface{})["uid"].(string)
		}

		processMap.mu.Lock()
		// Deduplicates the same event by storing it as a map rather
		processMap.processMap[key] = meshsyncEvent
		processMap.mu.Unlock()

		if !shouldWait {
			shouldWait = true
			isLast = false

			go func() {
				<-time.After(d)
				shouldWait = false

				processMap.mu.Lock()
				for k, v := range processMap.processMap {
					publishChannel <- v
					go r.Config.DashboardK8sResourcesChan.PublishDashboardK8sResources()
					// instead send the MeshSync event itself and update the plugin to send out this particular event instead of again querying the database.
					go func() {
						meshsyncChan <- struct{}{}
					}()
					// delete the key once processed to collect new entries
					delete(processMap.processMap, k)
				}
				processMap.mu.Unlock()
			}()
		} else {
			// don't let the last items stay in the queue until next event, but execute once the timer ends
			isLast = true
			go func() {
				<-time.After(d)
				if isLast {
					processMap.mu.Lock()
					for k, v := range processMap.processMap {
						publishChannel <- v
						go r.Config.DashboardK8sResourcesChan.PublishDashboardK8sResources()

						// delete the key once processed to collect new entries
						delete(processMap.processMap, k)
					}
					processMap.mu.Unlock()
				}
			}()
		}
	}
}
