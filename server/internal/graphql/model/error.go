package model

import (
	"github.com/khulnasoft/meshplay/meshkit/errors"
)

// Please reference the following before contributing an error code:
// https://docs.meshplay.khulnasoft.com/project/contributing/contributing-error
// https://github.com/khulnasoft/meshkit/blob/master/errors/errors.go
const (
	ErrNilClientCode                            = "1300"
	ErrCreateDataCode                           = "1301"
	ErrQueryCode                                = "1302"
	ErrMeshsyncSubscriptionCode                 = "1303"
	ErrMeshplayClientCode                        = "1304"
	ErrSubscribeChannelCode                     = "1305"
	ErrPublishBrokerCode                        = "1306"
	ErrEmptyHandlerCode                         = "1307"
	ErrApplyHelmChartCode                       = "1308"
	ErrMeshplayControllersStatusSubscriptionCode = "1309"
	ErrMeshSyncEventsSubscriptionCode           = "1310"
	ErrMeshplayClientNilCode                     = "1311"
	ErrUpdateDataCode                           = "1312"
	ErrDeleteDataCode                           = "1313"
)

var (
	ErrEmptyHandler     = errors.New(ErrEmptyHandlerCode, errors.Alert, []string{"Database handler not initialized"}, []string{"Meshplay Database handler is not accessible to perform operations"}, []string{"Meshplay Database is crashed or not reachable"}, []string{"Restart Meshplay Server", "Please check if Meshplay server is accessible to the Database"})
	ErrMeshplayClientNil = errors.New(ErrMeshplayClientNilCode, errors.Alert, []string{"Meshplay kubernetes client not initialized"}, []string{"Kubernetes config is not initialized with Meshplay 2"}, []string{}, []string{"Upload your kubernetes config via the settings dashboard. If uploaded, wait for a minute for it to get initialized"})
)

func ErrMeshSyncEventsSubscription(err error) error {
	return errors.New(ErrMeshSyncEventsSubscriptionCode, errors.Alert, []string{"Could not create MeshSync events subcription", err.Error()}, []string{"Meshplay controller handlers may not be available"}, []string{""}, []string{"Make sure the controllers are deployed and their handlers are configured"})
}

func ErrMeshplayControllersStatusSubscription(err error) error {
	return errors.New(ErrMeshplayControllersStatusSubscriptionCode, errors.Alert, []string{"Could not create meshplay controllers status subcription", err.Error()}, []string{"Meshplay controller handlers may not be available"}, []string{""}, []string{"Make sure the controllers are deployed and their handlers are configured"})
}

func ErrCreateData(err error) error {
	return errors.New(ErrCreateDataCode, errors.Alert, []string{"Error while writing meshsync data", err.Error()}, []string{"Unable to write MeshSync data to the Meshplay Database"}, []string{"Meshplay Database is crashed or not reachable"}, []string{"Restart Meshplay Server", "Please check if Meshplay server is accessible to the Database"})
}

func ErrUpdateData(err error) error {
	return errors.New(ErrUpdateDataCode, errors.Alert, []string{"Error while updating meshsync data", err.Error()}, []string{"Unable to update MeshSync data to the Meshplay Database"}, []string{"Meshplay Database is crashed or not reachable"}, []string{"Restart Meshplay Server", "Please check if Meshplay server is accessible to the Database"})
}

func ErrDeleteData(err error) error {
	return errors.New(ErrDeleteDataCode, errors.Alert, []string{"Error while deleting meshsync data", err.Error()}, []string{"Unable to read MeshSync data to the Meshplay Database"}, []string{"Meshplay Database is crashed or not reachable"}, []string{"Restart Meshplay Server", "Please check if Meshplay server is accessible to the Database"})
}

func ErrQuery(err error) error {
	return errors.New(ErrQueryCode, errors.Alert, []string{"Error while querying data", err.Error()}, []string{"Invalid Query performed in Meshplay Database"}, []string{}, []string{})
}

func ErrMeshsyncSubscription(err error) error {
	return errors.New(ErrMeshsyncSubscriptionCode, errors.Alert, []string{"MeshSync Subscription failed", err.Error()}, []string{"GraphQL subscription for MeshSync stopped"}, []string{"Could be a network issue"}, []string{"Check if meshplay server is reachable from the browser"})
}

func ErrSubscribeChannel(err error) error {
	return errors.New(ErrSubscribeChannelCode, errors.Alert, []string{"Unable to subscribe to channel", err.Error()}, []string{"Unable to create a broker subscription"}, []string{"Could be a network issue", "Meshplay Broker could have crashed"}, []string{"Check if Meshplay Broker is reachable from Meshplay Server", "Check if Meshplay Broker is up and running inside the configured cluster"})
}

func ErrPublishBroker(err error) error {
	return errors.New(ErrPublishBrokerCode, errors.Alert, []string{"Unable to publish to broker", err.Error()}, []string{"Unable to create a broker publisher"}, []string{"Could be a network issue", "Meshplay Broker could have crashed"}, []string{"Check if Meshplay Broker is reachable from Meshplay Server", "Check if Meshplay Broker is up and running inside the configured cluster"})
}

func ErrMeshplayClient(err error) error {
	return errors.New(ErrMeshplayClientCode, errors.Alert, []string{"Meshplay kubernetes client not initialized", err.Error()}, []string{"Kubernetes config is not initialized with Meshplay 1"}, []string{}, []string{"Upload your kubernetes config via the settings dashboard. If uploaded, wait for a minute for it to get initialized"})
}

// ErrApplyHelmChart is the error which occurs while applying helm chart
func ErrApplyHelmChart(err error) error {
	return errors.New(ErrApplyHelmChartCode, errors.Alert, []string{"Error occurred while applying Helm Chart"}, []string{err.Error()}, []string{"Kubernetes cluster might not be connected", "Leftover resources from partial install"}, []string{"Try reinstalling", "Try reconnecting your kubernetes cluster", "Clean up artifacts from preinstalled helm release manually"})
}
