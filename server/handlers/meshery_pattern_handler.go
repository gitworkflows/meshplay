package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/gofrs/uuid"
	"github.com/gorilla/mux"
	"github.com/khulnasoft/meshplay/server/meshes"
	"github.com/khulnasoft/meshplay/server/models"
	"github.com/khulnasoft/meshkit/errors"
	"github.com/khulnasoft/meshkit/models/events"
	"io"
	"net/http"
	"strings"

	pCore "github.com/khulnasoft/meshplay/server/models/pattern/core"
	"github.com/khulnasoft/meshplay/server/models/pattern/stages"
	"github.com/sirupsen/logrus"
)

// MeshplayPatternRequestBody refers to the type of request body that
// SaveMeshplayPattern would receive
type MeshplayPatternRequestBody struct {
	Name          string                 `json:"name,omitempty"`
	URL           string                 `json:"url,omitempty"`
	Path          string                 `json:"path,omitempty"`
	Save          bool                   `json:"save,omitempty"`
	PatternData   *models.MeshplayPattern `json:"pattern_data,omitempty"`
	CytoscapeJSON string                 `json:"cytoscape_json,omitempty"`
}

// PatternFileRequestHandler will handle requests of both type GET and POST
// on the route /api/pattern
func (h *Handler) PatternFileRequestHandler(
	rw http.ResponseWriter,
	r *http.Request,
	prefObj *models.Preference,
	user *models.User,
	provider models.Provider,
) {
	if r.Method == http.MethodGet {
		h.GetMeshplayPatternsHandler(rw, r, prefObj, user, provider)
		return
	}

	if r.Method == http.MethodPost {
		h.handlePatternPOST(rw, r, prefObj, user, provider)
		return
	}
}

// swagger:route POST /api/pattern PatternsAPI idPostPatternFile
// Handle POST requests for patterns
//
// Edit/update a meshplay pattern
// responses:
// 	200: meshplayPatternResponseWrapper

func (h *Handler) handlePatternPOST(
	rw http.ResponseWriter,
	r *http.Request,
	_ *models.Preference,
	user *models.User,
	provider models.Provider,
) {
	defer func() {
		_ = r.Body.Close()
	}()

	var err error
	userID := uuid.FromStringOrNil(user.ID)
	eventBuilder := events.NewEvent().FromUser(userID).FromSystem(*h.SystemID).WithCategory("pattern").WithAction("create").ActedUpon(userID).WithSeverity(events.Informational)

	res := meshes.EventsResponse{
		Component:     "core",
		ComponentName: "Design",
		OperationId:   uuid.Nil.String(), // to be removed
		EventType:     meshes.EventType_INFO,
	}
	var parsedBody *MeshplayPatternRequestBody
	if err := json.NewDecoder(r.Body).Decode(&parsedBody); err != nil {
		h.log.Error(ErrRequestBody(err))
		http.Error(rw, ErrRequestBody(err).Error(), http.StatusBadRequest)
		event := eventBuilder.WithSeverity(events.Error).WithMetadata(map[string]interface{}{
			"error": ErrRequestBody(err),
		}).WithDescription("Unable to parse uploaded pattern.").Build()

		_ = provider.PersistEvent(event)
		go h.config.EventBroadcaster.Publish(userID, event)
		return
	}

	actedUpon := &userID
	if parsedBody.PatternData != nil && parsedBody.PatternData.ID != nil {
		actedUpon = parsedBody.PatternData.ID
	}

	eventBuilder.ActedUpon(*actedUpon)

	token, err := provider.GetProviderToken(r)
	if err != nil {
		h.log.Error(ErrRetrieveUserToken(err))
		http.Error(rw, ErrRetrieveUserToken(err).Error(), http.StatusInternalServerError)
		event := eventBuilder.WithSeverity(events.Critical).WithMetadata(map[string]interface{}{
			"error": ErrRetrieveUserToken(err),
		}).WithDescription("No auth token provided in the request.").Build()

		_ = provider.PersistEvent(event)
		go h.config.EventBroadcaster.Publish(userID, event)

		return
	}

	format := r.URL.Query().Get("output")

	if parsedBody.CytoscapeJSON != "" {
		pf, err := pCore.NewPatternFileFromCytoscapeJSJSON(parsedBody.Name, []byte(parsedBody.CytoscapeJSON))
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(rw, "%s", err)
			event := eventBuilder.WithSeverity(events.Error).WithMetadata(map[string]interface{}{
				"error": ErrSavePattern(err),
			}).WithDescription("Pattern save failed, cytoJSON could be malformed.").Build()

			_ = provider.PersistEvent(event)
			go h.config.EventBroadcaster.Publish(userID, event)
			return
		}

		pfByt, err := pf.ToYAML()
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(rw, "%s", err)
			event := eventBuilder.WithSeverity(events.Error).WithMetadata(map[string]interface{}{
				"error": ErrSavePattern(err),
			}).WithDescription(ErrSavePattern(err).Error()).Build()

			_ = provider.PersistEvent(event)
			go h.config.EventBroadcaster.Publish(userID, event)
			return
		}

		patternName, err := models.GetPatternName(string(pfByt))
		if err != nil {
			h.log.Error(ErrSavePattern(err))
			http.Error(rw, ErrSavePattern(err).Error(), http.StatusBadRequest)
			event := eventBuilder.WithSeverity(events.Error).WithMetadata(map[string]interface{}{
				"error": ErrSavePattern(err),
			}).WithDescription("unable to get \"name\" from the pattern.").Build()

			_ = provider.PersistEvent(event)
			go h.config.EventBroadcaster.Publish(userID, event)
			return
		}

		meshplayPattern := &models.MeshplayPattern{
			Name:        patternName,
			PatternFile: string(pfByt),
			Location: map[string]interface{}{
				"host": "",
				"path": "",
				"type": "local",
			},
			CatalogData: parsedBody.PatternData.CatalogData,
		}
		if parsedBody.PatternData != nil {
			meshplayPattern.ID = parsedBody.PatternData.ID
		}

		if parsedBody.Save {
			resp, err := provider.SaveMeshplayPattern(token, meshplayPattern)
			if err != nil {
				h.log.Error(ErrSavePattern(err))
				http.Error(rw, ErrSavePattern(err).Error(), http.StatusInternalServerError)
				event := eventBuilder.WithSeverity(events.Error).WithMetadata(map[string]interface{}{
					"error": ErrSavePattern(err),
				}).WithDescription(ErrSavePattern(err).Error()).Build()

				_ = provider.PersistEvent(event)
				go h.config.EventBroadcaster.Publish(userID, event)
				return
			}

			h.formatPatternOutput(rw, resp, format, &res, eventBuilder)
			event := eventBuilder.Build()
			_ = provider.PersistEvent(event)
			// Do not send pattern save event if pattern is in cyto format as user is on meshmap and every node move will result in save request flooding user's screen.
			// go h.config.EventBroadcaster.Publish(userID, event)
			go h.config.PatternChannel.Publish(uuid.FromStringOrNil(user.ID), struct{}{})
			return
		}

		byt, err := json.Marshal([]models.MeshplayPattern{*meshplayPattern})
		if err != nil {
			h.log.Error(ErrEncodePattern(err))
			http.Error(rw, ErrEncodePattern(err).Error(), http.StatusInternalServerError)
			addMeshkitErr(&res, ErrEncodePattern(err))
			go h.EventsBuffer.Publish(&res)
			return
		}

		h.formatPatternOutput(rw, byt, format, &res, eventBuilder)

		return
	}
	// If Content is not empty then assume it's a local upload
	if parsedBody.PatternData != nil {
		// Check if the pattern is valid
		err := pCore.IsValidPattern(parsedBody.PatternData.PatternFile)
		if err != nil {
			h.log.Error(ErrInvalidPattern(err))
			http.Error(rw, ErrInvalidPattern(err).Error(), http.StatusBadRequest)
			addMeshkitErr(&res, ErrInvalidPattern(err))
			go h.EventsBuffer.Publish(&res)
			return
		}
		// Assign a name if no name is provided
		if parsedBody.PatternData.Name == "" {
			patternName, err := models.GetPatternName(parsedBody.PatternData.PatternFile)
			if err != nil {
				h.log.Error(ErrSavePattern(err))
				http.Error(rw, ErrSavePattern(err).Error(), http.StatusBadRequest)
				event := eventBuilder.WithSeverity(events.Error).WithMetadata(map[string]interface{}{
					"error": ErrSavePattern(err),
				}).WithDescription("unable to get \"name\" from the pattern.").Build()

				_ = provider.PersistEvent(event)
				go h.config.EventBroadcaster.Publish(userID, event)
				return
			}
			parsedBody.PatternData.Name = patternName
		}

		// Assign a location if no location is specified
		if parsedBody.PatternData.Location == nil {
			parsedBody.PatternData.Location = map[string]interface{}{
				"host":   "",
				"path":   "",
				"type":   "local",
				"branch": "",
			}
		}

		meshplayPattern := parsedBody.PatternData

		if parsedBody.Save {
			resp, err := provider.SaveMeshplayPattern(token, meshplayPattern)
			if err != nil {
				h.log.Error(ErrSavePattern(err))
				http.Error(rw, ErrSavePattern(err).Error(), http.StatusInternalServerError)

				event := eventBuilder.WithSeverity(events.Error).WithMetadata(map[string]interface{}{
					"error": ErrSavePattern(err),
				}).WithDescription(ErrSavePattern(err).Error()).Build()

				_ = provider.PersistEvent(event)
				go h.config.EventBroadcaster.Publish(userID, event)
				return
			}

			h.formatPatternOutput(rw, resp, format, &res, eventBuilder)
			event := eventBuilder.Build()
			_ = provider.PersistEvent(event)
			go h.config.EventBroadcaster.Publish(userID, event)
			go h.config.PatternChannel.Publish(uuid.FromStringOrNil(user.ID), struct{}{})
			return
		}

		byt, err := json.Marshal([]models.MeshplayPattern{*meshplayPattern})
		if err != nil {
			h.log.Error(ErrEncodePattern(err))
			http.Error(rw, ErrEncodePattern(err).Error(), http.StatusInternalServerError)
			addMeshkitErr(&res, ErrSavePattern(err))
			go h.EventsBuffer.Publish(&res)
			return
		}

		h.formatPatternOutput(rw, byt, format, &res, eventBuilder)
		event := eventBuilder.Build()
		_ = provider.PersistEvent(event)
		go h.config.EventBroadcaster.Publish(userID, event)
		return
	}

	if parsedBody.URL != "" {
		resp, err := provider.RemotePatternFile(r, parsedBody.URL, parsedBody.Path, parsedBody.Save)

		if err != nil {
			h.log.Error(ErrImportPattern(err))
			http.Error(rw, ErrImportPattern(err).Error(), http.StatusInternalServerError)
			event := eventBuilder.WithSeverity(events.Error).WithMetadata(map[string]interface{}{
				"error": ErrImportPattern(err),
			}).WithDescription(ErrImportPattern(err).Error()).Build()

			_ = provider.PersistEvent(event)
			go h.config.EventBroadcaster.Publish(userID, event)
			return
		}

		h.formatPatternOutput(rw, resp, format, &res, eventBuilder)
		event := eventBuilder.Build()
		_ = provider.PersistEvent(event)
		go h.config.EventBroadcaster.Publish(userID, event)
		return
	}
	//Depracated: The below logic was used when applications were stored as k8s_manifests.
	// if parsedBody.K8sManifest != "" {
	// 	pattern, err := pCore.NewPatternFileFromK8sManifest(parsedBody.K8sManifest, false)
	// 	if err != nil {
	// 		http.Error(rw, fmt.Sprintf("failed to convert to pattern: %s", err), http.StatusBadRequest)
	// 		fmt.Println("err: ", err)
	// 		addMeshkitErr(&res, err) //this error is already a meshkit error so no further wrapping required
	// 		go h.EventsBuffer.Publish(&res)
	// 		return
	// 	}

	// 	patternYAML, err := pattern.ToYAML()
	// 	if err != nil {
	// 		http.Error(rw, fmt.Sprintf("failed to generate pattern: %s", err), http.StatusInternalServerError)
	// 		addMeshkitErr(&res, ErrSavePattern(err))
	// 		go h.EventsBuffer.Publish(&res)
	// 		return
	// 	}

	// 	name, err := models.GetPatternName(string(patternYAML))
	// 	if err != nil {
	// 		http.Error(rw, fmt.Sprintf("failed to get pattern name: %s", err), http.StatusInternalServerError)
	// 		addMeshkitErr(&res, ErrSavePattern(err))
	// 		go h.EventsBuffer.Publish(&res)
	// 		return
	// 	}

	// 	patternModel := &models.MeshplayPattern{
	// 		Name:        name,
	// 		PatternFile: string(patternYAML),
	// 		Location: sql.Map{
	// 			"host": "",
	// 			"path": "",
	// 			"type": "local",
	// 		},
	// 	}

	// 	if parsedBody.Save {
	// 		resp, err := provider.SaveMeshplayPattern(token, patternModel)
	// 		if err != nil {
	// 			http.Error(rw, fmt.Sprintf("failed to save the pattern: %s", err), http.StatusInternalServerError)
	// 			addMeshkitErr(&res, ErrSavePattern(err))
	// 			go h.EventsBuffer.Publish(&res)
	// 			return
	// 		}

	// 		go h.config.PatternChannel.Publish(userID, struct{}{})
	// 		h.formatPatternOutput(rw, resp, format, &res)
	// 		return
	// 	}

	// 	byt, err := json.Marshal([]models.MeshplayPattern{*patternModel})
	// 	if err != nil {
	// 		http.Error(rw, fmt.Sprintf("failed to encode pattern: %s", err), http.StatusInternalServerError)
	// 		addMeshkitErr(&res, ErrSavePattern(err))
	// 		go h.EventsBuffer.Publish(&res)
	// 		return
	// 	}

	// 	h.formatPatternOutput(rw, byt, format, &res)
	// 	return
	// }
}

// swagger:route GET /api/pattern PatternsAPI idGetPatternFiles
// Handle GET request for patterns
//
// Returns the list of all the patterns saved by the current user
// This will return all the patterns with their details
//
// ```?order={field}``` orders on the passed field
//
// ```?search=<design name>``` A string matching is done on the specified design name
//
// ```?page={page-number}``` Default page number is 1
//
// ```?pagesize={pagesize}``` Default pagesize is 10
// responses:
//
//	200: meshplayPatternsResponseWrapper
func (h *Handler) GetMeshplayPatternsHandler(
	rw http.ResponseWriter,
	r *http.Request,
	_ *models.Preference,
	_ *models.User,
	provider models.Provider,
) {
	q := r.URL.Query()
	tokenString := r.Context().Value(models.TokenCtxKey).(string)
	updateAfter := q.Get("updated_after")
	resp, err := provider.GetMeshplayPatterns(tokenString, q.Get("page"), q.Get("pagesize"), q.Get("search"), q.Get("order"), updateAfter)
	if err != nil {
		h.log.Error(ErrFetchPattern(err))
		http.Error(rw, ErrFetchPattern(err).Error(), http.StatusInternalServerError)
		return
	}

	// token, err := provider.GetProviderToken(r)
	if err != nil {
		http.Error(rw, "failed to get user token", http.StatusInternalServerError)
		return
	}
	// mc := NewContentModifier(token, provider, prefObj, user.UserID)
	// //acts like a middleware, modifying the bytes lazily just before sending them back
	// err = mc.AddMetadataForPatterns(r.Context(), &resp)
	if err != nil {
		fmt.Println("Could not add metadata about pattern's current support ", err.Error())
	}
	rw.Header().Set("Content-Type", "application/json")
	fmt.Fprint(rw, string(resp))
}

// swagger:route GET /api/pattern/catalog PatternsAPI idGetCatalogMeshplayPatternsHandler
// Handle GET request for catalog patterns
//
// # Patterns can be further filtered through query parameter
//
// ```?order={field}``` orders on the passed field
//
// ```?page={page-number}``` Default page number is 0
//
// ```?pagesize={pagesize}``` Default pagesize is 10.
//
// ```?search={patternname}``` If search is non empty then a greedy search is performed
// responses:
//
//	200: meshplayPatternsResponseWrapper
func (h *Handler) GetCatalogMeshplayPatternsHandler(
	rw http.ResponseWriter,
	r *http.Request,
	_ *models.Preference,
	_ *models.User,
	provider models.Provider,
) {
	q := r.URL.Query()
	tokenString := r.Context().Value(models.TokenCtxKey).(string)

	resp, err := provider.GetCatalogMeshplayPatterns(tokenString, q.Get("page"), q.Get("pagesize"), q.Get("search"), q.Get("order"))
	if err != nil {
		h.log.Error(ErrFetchPattern(err))
		http.Error(rw, ErrFetchPattern(err).Error(), http.StatusInternalServerError)
		return
	}

	rw.Header().Set("Content-Type", "application/json")
	fmt.Fprint(rw, string(resp))
}

// swagger:route DELETE /api/pattern/{id} PatternsAPI idDeleteMeshplayPattern
// Handle Delete for a Meshplay Pattern
//
// Deletes a meshplay pattern with ID: id
// responses:
//
//	200: noContentWrapper
//
// DeleteMeshplayPatternHandler deletes a pattern with the given id
func (h *Handler) DeleteMeshplayPatternHandler(
	rw http.ResponseWriter,
	r *http.Request,
	_ *models.Preference,
	user *models.User,
	provider models.Provider,
) {
	patternID := mux.Vars(r)["id"]
	userID := uuid.FromStringOrNil(user.ID)
	eventBuilder := events.NewEvent().FromUser(userID).FromSystem(*h.SystemID).WithCategory("pattern").WithAction("delete").ActedUpon(uuid.FromStringOrNil(patternID))

	meshplayPattern := models.MeshplayPattern{}

	resp, err := provider.DeleteMeshplayPattern(r, patternID)
	if err != nil {
		errPatternDelete := ErrDeletePattern(err)

		h.log.Error(errPatternDelete)
		http.Error(rw, errPatternDelete.Error(), http.StatusInternalServerError)
		event := eventBuilder.WithSeverity(events.Error).WithMetadata(map[string]interface{}{
			"error": errPatternDelete,
		}).WithDescription("Error deleting pattern.").Build()
		http.Error(rw, errPatternDelete.Error(), http.StatusInternalServerError)
		_ = provider.PersistEvent(event)
		go h.config.EventBroadcaster.Publish(userID, event)
		return
	}

	_ = json.Unmarshal(resp, &meshplayPattern)
	event := eventBuilder.WithSeverity(events.Informational).WithDescription(fmt.Sprintf("Pattern %s deleted.", meshplayPattern.Name)).Build()
	_ = provider.PersistEvent(event)
	go h.config.EventBroadcaster.Publish(userID, event)
	go h.config.PatternChannel.Publish(uuid.FromStringOrNil(user.ID), struct{}{})

	rw.Header().Set("Content-Type", "application/json")
	fmt.Fprint(rw, string(resp))
}

// swagger:route GET /api/pattern/{id} PatternsAPI idGetMeshplayPattern
// Handle GET request for Meshplay Pattern with the given id
//
// Get the pattern with the given id
// responses:
//  200:

// GetMeshplayPatternHandler returns the pattern file with the given id

func (h *Handler) DownloadMeshplayPatternHandler(
	rw http.ResponseWriter,
	r *http.Request,
	_ *models.Preference,
	_ *models.User,
	provider models.Provider,
) {
	patternID := mux.Vars(r)["id"]
	resp, err := provider.GetMeshplayPattern(r, patternID)
	if err != nil {
		h.log.Error(ErrGetPattern(err))
		http.Error(rw, ErrGetPattern(err).Error(), http.StatusNotFound)
		return
	}

	pattern := &models.MeshplayPattern{}

	err = json.Unmarshal(resp, &pattern)
	if err != nil {
		obj := "download pattern"
		h.log.Error(models.ErrUnmarshal(err, obj))
		http.Error(rw, models.ErrUnmarshal(err, obj).Error(), http.StatusInternalServerError)
		return
	}

	rw.Header().Set("Content-Type", "application/x-yaml")
	if _, err := io.Copy(rw, strings.NewReader(pattern.PatternFile)); err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
}

// swagger:route POST /api/pattern/clone/{id} PatternsAPI idCloneMeshplayPattern
// Handle Clone for a Meshplay Pattern
//
// Creates a local copy of a published pattern with id: id
// responses:
//
//	200 : noContentWrapper
//
// CloneMeshplayPatternHandler clones a pattern with the given id
func (h *Handler) CloneMeshplayPatternHandler(
	rw http.ResponseWriter,
	r *http.Request,
	_ *models.Preference,
	user *models.User,
	provider models.Provider,
) {
	patternID := mux.Vars(r)["id"]
	var parsedBody *models.MeshplayClonePatternRequestBody
	if err := json.NewDecoder(r.Body).Decode(&parsedBody); err != nil || patternID == "" {
		h.log.Error(ErrRequestBody(err))
		http.Error(rw, ErrRequestBody(err).Error(), http.StatusBadRequest)
		return
	}

	resp, err := provider.CloneMeshplayPattern(r, patternID, parsedBody)
	if err != nil {
		h.log.Error(ErrClonePattern(err))
		http.Error(rw, ErrClonePattern(err).Error(), http.StatusInternalServerError)
		return
	}
	go h.config.PatternChannel.Publish(uuid.FromStringOrNil(user.ID), struct{}{})
	rw.Header().Set("Content-Type", "application/json")
	fmt.Fprint(rw, string(resp))
}

// swagger:route POST /api/pattern/catalog/publish PatternsAPI idPublishCatalogPatternHandler
// Handle Publish for a Meshplay Pattern
//
// Publishes pattern to Meshplay Catalog by setting visibility to published and setting catalog data
// responses:
//
//	202: noContentWrapper
//
// PublishCatalogPatternHandler sets visibility of pattern with given id as published
func (h *Handler) PublishCatalogPatternHandler(
	rw http.ResponseWriter,
	r *http.Request,
	_ *models.Preference,
	user *models.User,
	provider models.Provider,
) {
	defer func() {
		_ = r.Body.Close()
	}()

	var parsedBody *models.MeshplayCatalogPatternRequestBody
	if err := json.NewDecoder(r.Body).Decode(&parsedBody); err != nil {
		h.log.Error(ErrRequestBody(err))
		http.Error(rw, ErrRequestBody(err).Error(), http.StatusBadRequest)
		return
	}
	resp, err := provider.PublishCatalogPattern(r, parsedBody)
	if err != nil {
		h.log.Error(ErrPublishCatalogPattern(err))
		http.Error(rw, ErrPublishCatalogPattern(err).Error(), http.StatusInternalServerError)
		return
	}
	go h.config.PatternChannel.Publish(uuid.FromStringOrNil(user.ID), struct{}{})
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusAccepted)
	fmt.Fprint(rw, string(resp))
}

// swagger:route DELETE /api/pattern/catalog/unpublish PatternsAPI idUnPublishCatalogPatternHandler
// Handle Publish for a Meshplay Pattern
//
// Unpublishes pattern from Meshplay Catalog by setting visibility to private and removing catalog data from website
// responses:
//
//	200: noContentWrapper
//
// UnPublishCatalogPatternHandler sets visibility of pattern with given id as private
func (h *Handler) UnPublishCatalogPatternHandler(
	rw http.ResponseWriter,
	r *http.Request,
	_ *models.Preference,
	user *models.User,
	provider models.Provider,
) {
	defer func() {
		_ = r.Body.Close()
	}()

	var parsedBody *models.MeshplayCatalogPatternRequestBody
	if err := json.NewDecoder(r.Body).Decode(&parsedBody); err != nil {
		h.log.Error(ErrRequestBody(err))
		http.Error(rw, ErrRequestBody(err).Error(), http.StatusBadRequest)
		return
	}
	resp, err := provider.UnPublishCatalogPattern(r, parsedBody)
	if err != nil {
		h.log.Error(ErrPublishCatalogPattern(err))
		http.Error(rw, ErrPublishCatalogPattern(err).Error(), http.StatusInternalServerError)
		return
	}
	go h.config.PatternChannel.Publish(uuid.FromStringOrNil(user.ID), struct{}{})
	rw.Header().Set("Content-Type", "application/json")
	fmt.Fprint(rw, string(resp))
}

// swagger:route DELETE /api/patterns PatternsAPI idDeleteMeshplayPattern
// Handle Delete for multiple Meshplay Patterns
//
// DeleteMultiMeshplayPatternsHandler deletes patterns with the given ids
func (h *Handler) DeleteMultiMeshplayPatternsHandler(
	rw http.ResponseWriter,
	r *http.Request,
	_ *models.Preference,
	user *models.User,
	provider models.Provider,
) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		logrus.Error(rw, "err deleting pattern, converting bytes: ", err)
	}
	var patterns models.MeshplayPatternDeleteRequestBody
	err = json.Unmarshal([]byte(body), &patterns)
	if err != nil {
		logrus.Error("error marshaling patterns json: ", err)
	}

	logrus.Debugf("patterns to be deleted: %+v", patterns)

	resp, err := provider.DeleteMeshplayPatterns(r, patterns)

	if err != nil {
		http.Error(rw, fmt.Sprintf("failed to delete the pattern: %s", err), http.StatusInternalServerError)
		return
	}
	go h.config.PatternChannel.Publish(uuid.FromStringOrNil(user.ID), struct{}{})
	rw.Header().Set("Content-Type", "application/json")
	fmt.Fprint(rw, string(resp))
}

// swagger:route GET /api/pattern/{id} PatternsAPI idGetMeshplayPattern
// Handle GET for a Meshplay Pattern
//
// Fetches the pattern with the given id
// responses:
// 	200: meshplayPatternResponseWrapper

// GetMeshplayPatternHandler fetched the pattern with the given id
func (h *Handler) GetMeshplayPatternHandler(
	rw http.ResponseWriter,
	r *http.Request,
	_ *models.Preference,
	_ *models.User,
	provider models.Provider,
) {
	patternID := mux.Vars(r)["id"]

	resp, err := provider.GetMeshplayPattern(r, patternID)
	if err != nil {
		h.log.Error(ErrGetPattern(err))
		http.Error(rw, ErrGetPattern(err).Error(), http.StatusNotFound)
		return
	}

	rw.Header().Set("Content-Type", "application/json")
	fmt.Fprint(rw, string(resp))
}

func (h *Handler) formatPatternOutput(rw http.ResponseWriter, content []byte, format string, res *meshes.EventsResponse, eventBuilder *events.EventBuilder) {
	contentMeshplayPatternSlice := make([]models.MeshplayPattern, 0)

	if err := json.Unmarshal(content, &contentMeshplayPatternSlice); err != nil {
		http.Error(rw, ErrDecodePattern(err).Error(), http.StatusInternalServerError)
		addMeshkitErr(res, ErrDecodePattern(err))
		go h.EventsBuffer.Publish(res)
		return
	}

	result := []models.MeshplayPattern{}
	names := []string{}
	for _, content := range contentMeshplayPatternSlice {
		if content.ID != nil {
			eventBuilder.ActedUpon(*content.ID)
		}
		if format == "cytoscape" {
			patternFile, err := pCore.NewPatternFile([]byte(content.PatternFile))
			if err != nil {
				http.Error(rw, ErrParsePattern(err).Error(), http.StatusBadRequest)

				eventBuilder.WithSeverity(events.Error).WithMetadata(map[string]interface{}{
					"error": ErrParsePattern(err),
				}).WithDescription("Unable to parse pattern file, pattern could be malformed.").Build()
				return
			}

			//TODO: The below line has to go away once the client fully supports referencing variables  and pattern imports inside design
			newpatternfile := evalImportAndReferenceStage(&patternFile)

			cyjs, _ := newpatternfile.ToCytoscapeJS()

			bytes, err := json.Marshal(&cyjs)
			if err != nil {
				http.Error(rw, ErrConvertPattern(err).Error(), http.StatusInternalServerError)
				addMeshkitErr(res, ErrConvertPattern(err))
				go h.EventsBuffer.Publish(res)
				return
			}

			// Replace the patternfile with cytoscape type data
			content.PatternFile = string(bytes)
		}

		result = append(result, content)
		names = append(names, content.Name)
	}

	data, err := json.Marshal(&result)
	if err != nil {
		obj := "pattern file"
		http.Error(rw, models.ErrMarshal(err, obj).Error(), http.StatusInternalServerError)
		addMeshkitErr(res, models.ErrMarshal(err, obj))

		go h.EventsBuffer.Publish(res)
		return
	}
	eventBuilder.WithDescription(fmt.Sprintf("Design %s saved", strings.Join(names, ",")))
	rw.Header().Set("Content-Type", "application/json")
	fmt.Fprint(rw, string(data))
	res.Details = "\"" + strings.Join(names, ",") + "\" design saved"
	res.Summary = "Changes to the \"" + strings.Join(names, ",") + "\" design have been saved."
	// go h.EventsBuffer.Publish(res)
}

// Since the client currently does not support pattern imports and externalized variables, the first(import) stage of pattern engine
// is evaluated here to simplify the pattern file such that it is valid when a deploy takes place
func evalImportAndReferenceStage(p *pCore.Pattern) (newp pCore.Pattern) {
	sap := &serviceActionProvider{}
	sip := &serviceInfoProvider{}
	chain := stages.CreateChain()
	chain.
		Add(stages.Import(sip, sap)).
		Add(stages.Filler(false)).
		Add(func(data *stages.Data, err error, next stages.ChainStageNextFunction) {
			data.Lock.Lock()
			newp = *data.Pattern
			data.Lock.Unlock()
		}).
		Process(&stages.Data{
			Pattern: p,
		})
	return newp
}

// Only pass Meshkit err here or there will be a panic
func addMeshkitErr(res *meshes.EventsResponse, err error) {
	if err != nil {
		res.EventType = meshes.EventType_ERROR
		res.ProbableCause = errors.GetCause(err)
		res.SuggestedRemediation = errors.GetRemedy(err)
		res.Details = err.Error()
		res.Summary = errors.GetSDescription(err)
		res.ErrorCode = errors.GetCode(err)
	}
}
