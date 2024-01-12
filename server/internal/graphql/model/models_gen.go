// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
	"time"
)

type AdapterStatusInput struct {
	TargetStatus Status `json:"targetStatus"`
	TargetPort   string `json:"targetPort"`
	Adapter      string `json:"adapter"`
}

type AddonList struct {
	Name  string `json:"name"`
	Owner string `json:"owner"`
}

type AddonStatusInput struct {
	Selector     *MeshType `json:"selector,omitempty"`
	K8scontextID string    `json:"k8scontextID"`
	TargetStatus Status    `json:"targetStatus"`
}

type ApplicationPage struct {
	Page         int                  `json:"page"`
	PageSize     int                  `json:"page_size"`
	TotalCount   int                  `json:"total_count"`
	Applications []*ApplicationResult `json:"applications,omitempty"`
}

type ApplicationResult struct {
	ID              string      `json:"id"`
	Name            string      `json:"name"`
	ApplicationFile string      `json:"application_file"`
	Type            *NullString `json:"type"`
	UserID          string      `json:"user_id"`
	Location        *Location   `json:"location"`
	Visibility      string      `json:"visibility"`
	CreatedAt       *string     `json:"created_at,omitempty"`
	UpdatedAt       *string     `json:"updated_at,omitempty"`
}

type CatalogFilter struct {
	ID             string                 `json:"id"`
	Name           string                 `json:"name"`
	FilterFile     string                 `json:"filter_file"`
	UserID         string                 `json:"user_id"`
	Location       *Location              `json:"location"`
	FilterResource string                 `json:"filter_resource"`
	Visibility     string                 `json:"visibility"`
	CatalogData    map[string]interface{} `json:"catalog_data,omitempty"`
	CreatedAt      *string                `json:"created_at,omitempty"`
	UpdatedAt      *string                `json:"updated_at,omitempty"`
}

type CatalogPattern struct {
	ID          string                 `json:"id"`
	Name        string                 `json:"name"`
	UserID      string                 `json:"user_id"`
	PatternFile string                 `json:"pattern_file"`
	Location    *Location              `json:"location"`
	Visibility  string                 `json:"visibility"`
	CatalogData map[string]interface{} `json:"catalog_data,omitempty"`
	CreatedAt   *string                `json:"created_at,omitempty"`
	UpdatedAt   *string                `json:"updated_at,omitempty"`
}

type CatalogSelector struct {
	Page     string `json:"page"`
	Pagesize string `json:"pagesize"`
	Search   string `json:"search"`
	Order    string `json:"order"`
}

type ClusterResources struct {
	Resources []*Resource `json:"resources"`
}

type ConfigurationPage struct {
	Applications *ApplicationPage   `json:"applications,omitempty"`
	Patterns     *PatternPageResult `json:"patterns,omitempty"`
	Filters      *FilterPage        `json:"filters,omitempty"`
}

type Container struct {
	ControlPlaneMemberName string           `json:"controlPlaneMemberName"`
	ContainerName          string           `json:"containerName"`
	Image                  string           `json:"image"`
	Status                 *ContainerStatus `json:"status,omitempty"`
	Ports                  []*ContainerPort `json:"ports,omitempty"`
	Resources              interface{}      `json:"resources,omitempty"`
}

type ContainerPort struct {
	Name          *string `json:"name,omitempty"`
	ContainerPort int     `json:"containerPort"`
	Protocol      string  `json:"protocol"`
}

type ContainerStatus struct {
	ContainerStatusName string      `json:"containerStatusName"`
	Image               string      `json:"image"`
	State               interface{} `json:"state,omitempty"`
	LastState           interface{} `json:"lastState,omitempty"`
	Ready               bool        `json:"ready"`
	RestartCount        interface{} `json:"restartCount,omitempty"`
	Started             bool        `json:"started"`
	ImageID             interface{} `json:"imageID,omitempty"`
	ContainerID         interface{} `json:"containerID,omitempty"`
}

type ControlPlane struct {
	Name    string                `json:"name"`
	Members []*ControlPlaneMember `json:"members"`
}

type ControlPlaneMember struct {
	Name       string       `json:"name"`
	Component  string       `json:"component"`
	Version    string       `json:"version"`
	Namespace  string       `json:"namespace"`
	DataPlanes []*Container `json:"data_planes,omitempty"`
}

type DataPlane struct {
	Name    string       `json:"name"`
	Proxies []*Container `json:"proxies"`
}

type Error struct {
	Code        string `json:"code"`
	Description string `json:"description"`
}

type Event struct {
	ID          string                 `json:"id"`
	UserID      string                 `json:"userID"`
	ActedUpon   string                 `json:"actedUpon"`
	OperationID string                 `json:"operationID"`
	SystemID    string                 `json:"systemID"`
	Severity    Severity               `json:"severity"`
	Action      string                 `json:"action"`
	Status      string                 `json:"status"`
	Category    string                 `json:"category"`
	Description string                 `json:"description"`
	Metadata    map[string]interface{} `json:"metadata,omitempty"`
	CreatedAt   time.Time              `json:"createdAt"`
	UpdatedAt   time.Time              `json:"updatedAt"`
	DeletedAt   *time.Time             `json:"deletedAt,omitempty"`
}

type FilterPage struct {
	Page       int             `json:"page"`
	PageSize   int             `json:"page_size"`
	TotalCount int             `json:"total_count"`
	Filters    []*FilterResult `json:"filters,omitempty"`
}

type FilterResult struct {
	ID             string                 `json:"id"`
	Name           string                 `json:"name"`
	FilterFile     string                 `json:"filter_file"`
	FilterResource string                 `json:"filter_resource"`
	UserID         string                 `json:"user_id"`
	Location       *Location              `json:"location"`
	Visibility     string                 `json:"visibility"`
	CatalogData    map[string]interface{} `json:"catalog_data,omitempty"`
	CreatedAt      *string                `json:"created_at,omitempty"`
	UpdatedAt      *string                `json:"updated_at,omitempty"`
}

type K8sContext struct {
	ID                 string `json:"id"`
	Name               string `json:"name"`
	Server             string `json:"server"`
	Owner              string `json:"owner"`
	CreatedBy          string `json:"created_by"`
	MeshplayInstanceID  string `json:"meshery_instance_id"`
	KubernetesServerID string `json:"kubernetes_server_id"`
	DeploymentType     string `json:"deployment_type"`
	Version            string `json:"version"`
	UpdatedAt          string `json:"updated_at"`
	CreatedAt          string `json:"created_at"`
	ConnectionID       string `json:"connection_id"`
}

type K8sContextsPage struct {
	TotalCount int           `json:"total_count"`
	Contexts   []*K8sContext `json:"contexts"`
}

type KctlDescribeDetails struct {
	Describe *string `json:"describe,omitempty"`
	Ctxid    *string `json:"ctxid,omitempty"`
}

type Location struct {
	Branch *string `json:"branch,omitempty"`
	Host   *string `json:"host,omitempty"`
	Path   *string `json:"path,omitempty"`
	Type   *string `json:"type,omitempty"`
}

type MeshModelComponent struct {
	Name  string `json:"name"`
	Count int    `json:"count"`
}

type MeshModelRelationship struct {
	Name  string `json:"name"`
	Count int    `json:"count"`
}

type MeshModelSummary struct {
	Components    []*MeshModelComponent    `json:"components,omitempty"`
	Relationships []*MeshModelRelationship `json:"relationships,omitempty"`
}

type MeshModelSummarySelector struct {
	Type string `json:"type"`
}

type MeshSyncEvent struct {
	Type      string      `json:"type"`
	Object    interface{} `json:"object"`
	ContextID string      `json:"contextId"`
}

type MeshplayControllersStatusListItem struct {
	ContextID  string                  `json:"contextId"`
	Controller MeshplayController       `json:"controller"`
	Status     MeshplayControllerStatus `json:"status"`
	Version    *string                 `json:"version,omitempty"`
}

type MeshplayResult struct {
	MeshplayID          *string                `json:"meshery_id,omitempty"`
	Name               *string                `json:"name,omitempty"`
	Mesh               *string                `json:"mesh,omitempty"`
	PerformanceProfile *string                `json:"performance_profile,omitempty"`
	TestID             *string                `json:"test_id,omitempty"`
	RunnerResults      map[string]interface{} `json:"runner_results,omitempty"`
	ServerMetrics      *string                `json:"server_metrics,omitempty"`
	ServerBoardConfig  *string                `json:"server_board_config,omitempty"`
	TestStartTime      *string                `json:"test_start_time,omitempty"`
	UserID             *string                `json:"user_id,omitempty"`
	UpdatedAt          *string                `json:"updated_at,omitempty"`
	CreatedAt          *string                `json:"created_at,omitempty"`
}

type NameSpace struct {
	Namespace string `json:"namespace"`
}

type NullString struct {
	String string `json:"String"`
	Valid  bool   `json:"Valid"`
}

type OAMCapability struct {
	OamDefinition interface{} `json:"oam_definition,omitempty"`
	ID            *string     `json:"id,omitempty"`
	OamRefSchema  *string     `json:"oam_ref_schema,omitempty"`
	Host          *string     `json:"host,omitempty"`
	Restricted    *bool       `json:"restricted,omitempty"`
	Metadata      interface{} `json:"metadata,omitempty"`
}

type OperatorControllerStatus struct {
	Name      string `json:"name"`
	Version   string `json:"version"`
	Status    Status `json:"status"`
	Error     *Error `json:"error,omitempty"`
	ContextID string `json:"contextID"`
}

type OperatorControllerStatusPerK8sContext struct {
	ContextID                string                    `json:"contextID"`
	OperatorControllerStatus *OperatorControllerStatus `json:"OperatorControllerStatus"`
}

type OperatorStatus struct {
	Status      Status                      `json:"status"`
	Version     string                      `json:"version"`
	Controllers []*OperatorControllerStatus `json:"controllers"`
	Error       *Error                      `json:"error,omitempty"`
	ContextID   string                      `json:"contextID"`
}

type OperatorStatusInput struct {
	TargetStatus Status `json:"targetStatus"`
	ContextID    string `json:"contextID"`
}

type OperatorStatusPerK8sContext struct {
	ContextID      string          `json:"contextID"`
	OperatorStatus *OperatorStatus `json:"operatorStatus"`
}

type PageFilter struct {
	Page         string   `json:"page"`
	PageSize     string   `json:"pageSize"`
	Order        *string  `json:"order,omitempty"`
	Search       *string  `json:"search,omitempty"`
	From         *string  `json:"from,omitempty"`
	To           *string  `json:"to,omitempty"`
	UpdatedAfter *string  `json:"updated_after,omitempty"`
	Visibility   []string `json:"visibility,omitempty"`
}

type PatternPageResult struct {
	Page       int              `json:"page"`
	PageSize   int              `json:"page_size"`
	TotalCount int              `json:"total_count"`
	Patterns   []*PatternResult `json:"patterns,omitempty"`
}

type PatternResult struct {
	ID          string                 `json:"id"`
	Name        string                 `json:"name"`
	UserID      string                 `json:"user_id"`
	Location    *Location              `json:"location"`
	PatternFile string                 `json:"pattern_file"`
	Visibility  string                 `json:"visibility"`
	CatalogData map[string]interface{} `json:"catalog_data,omitempty"`
	CanSupport  bool                   `json:"canSupport"`
	Errmsg      *string                `json:"errmsg,omitempty"`
	CreatedAt   *string                `json:"created_at,omitempty"`
	UpdatedAt   *string                `json:"updated_at,omitempty"`
}

type PerfPageProfiles struct {
	Page       int            `json:"page"`
	PageSize   int            `json:"page_size"`
	TotalCount int            `json:"total_count"`
	Profiles   []*PerfProfile `json:"profiles,omitempty"`
}

type PerfPageResult struct {
	Page       int              `json:"page"`
	PageSize   int              `json:"page_size"`
	TotalCount int              `json:"total_count"`
	Results    []*MeshplayResult `json:"results,omitempty"`
}

type PerfProfile struct {
	ConcurrentRequest int                    `json:"concurrent_request"`
	CreatedAt         *string                `json:"created_at,omitempty"`
	Duration          string                 `json:"duration"`
	Endpoints         []*string              `json:"endpoints,omitempty"`
	ID                string                 `json:"id"`
	LastRun           *string                `json:"last_run,omitempty"`
	LoadGenerators    []*string              `json:"load_generators,omitempty"`
	Name              *string                `json:"name,omitempty"`
	QPS               *int                   `json:"qps,omitempty"`
	TotalResults      *int                   `json:"total_results,omitempty"`
	UpdatedAt         *string                `json:"updated_at,omitempty"`
	UserID            string                 `json:"user_id"`
	RequestHeaders    *string                `json:"request_headers,omitempty"`
	RequestCookies    *string                `json:"request_cookies,omitempty"`
	RequestBody       *string                `json:"request_body,omitempty"`
	ContentType       *string                `json:"content_type,omitempty"`
	ServiceMesh       *string                `json:"service_mesh,omitempty"`
	Metadata          map[string]interface{} `json:"metadata,omitempty"`
}

type ReSyncActions struct {
	ClearDb   string `json:"clearDB"`
	ReSync    string `json:"ReSync"`
	HardReset string `json:"hardReset"`
}

type Resource struct {
	Kind  string `json:"kind"`
	Count int    `json:"count"`
}

type ServiceMeshFilter struct {
	Type          *MeshType `json:"type,omitempty"`
	K8sClusterIDs []string  `json:"k8sClusterIDs,omitempty"`
}

type TelemetryComp struct {
	Name   string `json:"name"`
	Spec   string `json:"spec"`
	Status string `json:"status"`
}

type MeshSyncEventType string

const (
	MeshSyncEventTypeAdded    MeshSyncEventType = "ADDED"
	MeshSyncEventTypeModified MeshSyncEventType = "MODIFIED"
	MeshSyncEventTypeDeleted  MeshSyncEventType = "DELETED"
)

var AllMeshSyncEventType = []MeshSyncEventType{
	MeshSyncEventTypeAdded,
	MeshSyncEventTypeModified,
	MeshSyncEventTypeDeleted,
}

func (e MeshSyncEventType) IsValid() bool {
	switch e {
	case MeshSyncEventTypeAdded, MeshSyncEventTypeModified, MeshSyncEventTypeDeleted:
		return true
	}
	return false
}

func (e MeshSyncEventType) String() string {
	return string(e)
}

func (e *MeshSyncEventType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = MeshSyncEventType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid MeshSyncEventType", str)
	}
	return nil
}

func (e MeshSyncEventType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type MeshType string

const (
	MeshTypeAllMesh            MeshType = "ALL_MESH"
	MeshTypeInvalidMesh        MeshType = "INVALID_MESH"
	MeshTypeAppMesh            MeshType = "APP_MESH"
	MeshTypeCitrixServiceMesh  MeshType = "CITRIX_SERVICE_MESH"
	MeshTypeConsul             MeshType = "CONSUL"
	MeshTypeIstio              MeshType = "ISTIO"
	MeshTypeKuma               MeshType = "KUMA"
	MeshTypeLinkerd            MeshType = "LINKERD"
	MeshTypeTraefikMesh        MeshType = "TRAEFIK_MESH"
	MeshTypeOctarine           MeshType = "OCTARINE"
	MeshTypeNetworkServiceMesh MeshType = "NETWORK_SERVICE_MESH"
	MeshTypeTanzu              MeshType = "TANZU"
	MeshTypeOpenServiceMesh    MeshType = "OPEN_SERVICE_MESH"
	MeshTypeNginxServiceMesh   MeshType = "NGINX_SERVICE_MESH"
	MeshTypeCiliumServiceMesh  MeshType = "CILIUM_SERVICE_MESH"
)

var AllMeshType = []MeshType{
	MeshTypeAllMesh,
	MeshTypeInvalidMesh,
	MeshTypeAppMesh,
	MeshTypeCitrixServiceMesh,
	MeshTypeConsul,
	MeshTypeIstio,
	MeshTypeKuma,
	MeshTypeLinkerd,
	MeshTypeTraefikMesh,
	MeshTypeOctarine,
	MeshTypeNetworkServiceMesh,
	MeshTypeTanzu,
	MeshTypeOpenServiceMesh,
	MeshTypeNginxServiceMesh,
	MeshTypeCiliumServiceMesh,
}

func (e MeshType) IsValid() bool {
	switch e {
	case MeshTypeAllMesh, MeshTypeInvalidMesh, MeshTypeAppMesh, MeshTypeCitrixServiceMesh, MeshTypeConsul, MeshTypeIstio, MeshTypeKuma, MeshTypeLinkerd, MeshTypeTraefikMesh, MeshTypeOctarine, MeshTypeNetworkServiceMesh, MeshTypeTanzu, MeshTypeOpenServiceMesh, MeshTypeNginxServiceMesh, MeshTypeCiliumServiceMesh:
		return true
	}
	return false
}

func (e MeshType) String() string {
	return string(e)
}

func (e *MeshType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = MeshType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid MeshType", str)
	}
	return nil
}

func (e MeshType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type MeshplayController string

const (
	MeshplayControllerBroker   MeshplayController = "BROKER"
	MeshplayControllerOperator MeshplayController = "OPERATOR"
	MeshplayControllerMeshsync MeshplayController = "MESHSYNC"
)

var AllMeshplayController = []MeshplayController{
	MeshplayControllerBroker,
	MeshplayControllerOperator,
	MeshplayControllerMeshsync,
}

func (e MeshplayController) IsValid() bool {
	switch e {
	case MeshplayControllerBroker, MeshplayControllerOperator, MeshplayControllerMeshsync:
		return true
	}
	return false
}

func (e MeshplayController) String() string {
	return string(e)
}

func (e *MeshplayController) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = MeshplayController(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid MeshplayController", str)
	}
	return nil
}

func (e MeshplayController) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type MeshplayControllerStatus string

const (
	MeshplayControllerStatusDeployed    MeshplayControllerStatus = "DEPLOYED"
	MeshplayControllerStatusNotdeployed MeshplayControllerStatus = "NOTDEPLOYED"
	MeshplayControllerStatusDeploying   MeshplayControllerStatus = "DEPLOYING"
	MeshplayControllerStatusUnkown      MeshplayControllerStatus = "UNKOWN"
	MeshplayControllerStatusUndeployed  MeshplayControllerStatus = "UNDEPLOYED"
	MeshplayControllerStatusEnabled     MeshplayControllerStatus = "ENABLED"
	MeshplayControllerStatusRunning     MeshplayControllerStatus = "RUNNING"
	MeshplayControllerStatusConnected   MeshplayControllerStatus = "CONNECTED"
)

var AllMeshplayControllerStatus = []MeshplayControllerStatus{
	MeshplayControllerStatusDeployed,
	MeshplayControllerStatusNotdeployed,
	MeshplayControllerStatusDeploying,
	MeshplayControllerStatusUnkown,
	MeshplayControllerStatusUndeployed,
	MeshplayControllerStatusEnabled,
	MeshplayControllerStatusRunning,
	MeshplayControllerStatusConnected,
}

func (e MeshplayControllerStatus) IsValid() bool {
	switch e {
	case MeshplayControllerStatusDeployed, MeshplayControllerStatusNotdeployed, MeshplayControllerStatusDeploying, MeshplayControllerStatusUnkown, MeshplayControllerStatusUndeployed, MeshplayControllerStatusEnabled, MeshplayControllerStatusRunning, MeshplayControllerStatusConnected:
		return true
	}
	return false
}

func (e MeshplayControllerStatus) String() string {
	return string(e)
}

func (e *MeshplayControllerStatus) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = MeshplayControllerStatus(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid MeshplayControllerStatus", str)
	}
	return nil
}

func (e MeshplayControllerStatus) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type Severity string

const (
	SeverityAlert         Severity = "alert"
	SeverityCritical      Severity = "critical"
	SeverityDebug         Severity = "debug"
	SeverityEmergency     Severity = "emergency"
	SeverityError         Severity = "error"
	SeverityWarning       Severity = "warning"
	SeverityInformational Severity = "informational"
)

var AllSeverity = []Severity{
	SeverityAlert,
	SeverityCritical,
	SeverityDebug,
	SeverityEmergency,
	SeverityError,
	SeverityWarning,
	SeverityInformational,
}

func (e Severity) IsValid() bool {
	switch e {
	case SeverityAlert, SeverityCritical, SeverityDebug, SeverityEmergency, SeverityError, SeverityWarning, SeverityInformational:
		return true
	}
	return false
}

func (e Severity) String() string {
	return string(e)
}

func (e *Severity) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Severity(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Severity", str)
	}
	return nil
}

func (e Severity) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type Status string

const (
	StatusEnabled    Status = "ENABLED"
	StatusConnected  Status = "CONNECTED"
	StatusDisabled   Status = "DISABLED"
	StatusProcessing Status = "PROCESSING"
	StatusUnknown    Status = "UNKNOWN"
)

var AllStatus = []Status{
	StatusEnabled,
	StatusConnected,
	StatusDisabled,
	StatusProcessing,
	StatusUnknown,
}

func (e Status) IsValid() bool {
	switch e {
	case StatusEnabled, StatusConnected, StatusDisabled, StatusProcessing, StatusUnknown:
		return true
	}
	return false
}

func (e Status) String() string {
	return string(e)
}

func (e *Status) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Status(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Status", str)
	}
	return nil
}

func (e Status) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
