package models

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/gofrs/uuid"
	"gorm.io/gorm"

	"github.com/khulnasoft/meshkit/database"
)

// MeshplayPatternPersister is the persister for persisting
// performance profiles on the database
type MeshplayPatternPersister struct {
	DB *database.Handler
}

// MeshplayPatternPage represents a page of performance profiles
type MeshplayPatternPage struct {
	Page       uint64            `json:"page"`
	PageSize   uint64            `json:"page_size"`
	TotalCount int               `json:"total_count"`
	Patterns   []*MeshplayPattern `json:"patterns"`
}

// GetMeshplayPatterns returns all of the 'private' patterns. Though private has no meaning here since there is only
// one local user. We make this distinction to be consistent with the remote provider
func (mpp *MeshplayPatternPersister) GetMeshplayPatterns(search, order string, page, pageSize uint64, updatedAfter string, visibility []string) ([]byte, error) {
	order = SanitizeOrderInput(order, []string{"created_at", "updated_at", "name"})

	if order == "" {
		order = "updated_at desc"
	}

	count := int64(0)
	patterns := []*MeshplayPattern{}
	var query *gorm.DB
	if len(visibility) == 0 {
		query = mpp.DB.Where("visibility in (?)", visibility)
	}
	query = query.Where("updated_at > ?", updatedAfter).Order(order)

	if search != "" {
		like := "%" + strings.ToLower(search) + "%"
		query = query.Where("(lower(meshery_patterns.name) like ?)", like)
	}

	query.Table("meshery_patterns").Count(&count)

	Paginate(uint(page), uint(pageSize))(query).Find(&patterns)

	mesheryPatternPage := &MeshplayPatternPage{
		Page:       page,
		PageSize:   pageSize,
		TotalCount: int(count),
		Patterns:   patterns,
	}

	return marshalMeshplayPatternPage(mesheryPatternPage), nil
}

// GetMeshplayCatalogPatterns returns all of the published patterns
func (mpp *MeshplayPatternPersister) GetMeshplayCatalogPatterns(page, pageSize, search, order string) ([]byte, error) {
	var err error
	order = SanitizeOrderInput(order, []string{"created_at", "updated_at", "name"})

	if order == "" {
		order = "updated_at desc"
	}

	var pg int
	if page != "" {
		pg, err = strconv.Atoi(page)

		if err != nil || pg < 0 {
			pg = 0
		}
	} else {
		pg = 0
	}

	// 0 page size is for all records
	var pgSize int
	if pageSize != "" {
		pgSize, err = strconv.Atoi(pageSize)

		if err != nil || pgSize < 0 {
			pgSize = 0
		}
	} else {
		pgSize = 0
	}

	patterns := []MeshplayPattern{}

	query := mpp.DB.Where("visibility = ?", Published).Order(order)

	if search != "" {
		like := "%" + strings.ToLower(search) + "%"
		query = query.Where("(lower(meshery_patterns.name) like ?)", like)
	}

	var count int64
	err = query.Model(&MeshplayPattern{}).Count(&count).Error

	if err != nil {
		return nil, err
	}

	if pgSize != 0 {
		Paginate(uint(pg), uint(pgSize))(query).Find(&patterns)
	} else {
		query.Find(&patterns)
	}

	response := PatternsAPIResponse{
		Page:       uint(pg),
		PageSize:   uint(pgSize),
		TotalCount: uint(count),
		Patterns:   patterns,
	}

	marshalledResponse, _ := json.Marshal(response)
	return marshalledResponse, nil
}

// CloneMeshplayPattern clones meshery pattern to private
func (mpp *MeshplayPatternPersister) CloneMeshplayPattern(patternID string, clonePatternRequest *MeshplayClonePatternRequestBody) ([]byte, error) {
	var mesheryPattern MeshplayPattern
	patternUUID, _ := uuid.FromString(patternID)
	err := mpp.DB.First(&mesheryPattern, patternUUID).Error
	if err != nil || *mesheryPattern.ID == uuid.Nil {
		return nil, fmt.Errorf("unable to get design: %w", err)
	}

	id, err := uuid.NewV4()
	if err != nil {
		return nil, err
	}

	mesheryPattern.Visibility = Private
	mesheryPattern.ID = &id
	mesheryPattern.Name = clonePatternRequest.Name

	return mpp.SaveMeshplayPattern(&mesheryPattern)
}

// DeleteMeshplayPattern takes in a profile id and delete it if it already exists
func (mpp *MeshplayPatternPersister) DeleteMeshplayPattern(id uuid.UUID) ([]byte, error) {
	pattern := MeshplayPattern{ID: &id}
	mpp.DB.Delete(&pattern)

	return marshalMeshplayPattern(&pattern), nil
}

// DeleteMeshplayPatterns takes in a meshery-patterns and delete those if exist
func (mpp *MeshplayPatternPersister) DeleteMeshplayPatterns(patterns MeshplayPatternDeleteRequestBody) ([]byte, error) {
	var deletedMaptterns []MeshplayPattern
	for _, pObj := range patterns.Patterns {
		id := uuid.FromStringOrNil(pObj.ID)
		pattern := MeshplayPattern{ID: &id}
		mpp.DB.Delete(&pattern)
		deletedMaptterns = append(deletedMaptterns, pattern)
	}

	return marshalMeshplayPatterns(deletedMaptterns), nil
}

func (mpp *MeshplayPatternPersister) SaveMeshplayPattern(pattern *MeshplayPattern) ([]byte, error) {
	if pattern.Visibility == "" {
		pattern.Visibility = Private
	}
	if pattern.ID == nil {
		id, err := uuid.NewV4()
		if err != nil {
			return nil, ErrGenerateUUID(err)
		}

		pattern.ID = &id
	}

	return marshalMeshplayPatterns([]MeshplayPattern{*pattern}), mpp.DB.Save(pattern).Error
}

// SaveMeshplayPatterns batch inserts the given patterns
func (mpp *MeshplayPatternPersister) SaveMeshplayPatterns(patterns []MeshplayPattern) ([]byte, error) {
	finalPatterns := []MeshplayPattern{}
	nilUserID := ""
	for _, pattern := range patterns {
		if pattern.Visibility == "" {
			pattern.Visibility = Private
		}
		pattern.UserID = &nilUserID
		if pattern.ID == nil {
			id, err := uuid.NewV4()
			if err != nil {
				return nil, ErrGenerateUUID(err)
			}

			pattern.ID = &id
		}

		finalPatterns = append(finalPatterns, pattern)
	}

	return marshalMeshplayPatterns(finalPatterns), mpp.DB.Create(finalPatterns).Error
}

func (mpp *MeshplayPatternPersister) GetMeshplayPattern(id uuid.UUID) ([]byte, error) {
	var mesheryPattern MeshplayPattern

	err := mpp.DB.First(&mesheryPattern, id).Error
	return marshalMeshplayPattern(&mesheryPattern), err
}

func (mpp *MeshplayPatternPersister) GetMeshplayPatternSource(id uuid.UUID) ([]byte, error) {
	var mesheryPattern MeshplayPattern
	err := mpp.DB.First(&mesheryPattern, id).Error
	return mesheryPattern.SourceContent, err
}

func marshalMeshplayPatternPage(mpp *MeshplayPatternPage) []byte {
	res, _ := json.Marshal(mpp)

	return res
}

func marshalMeshplayPattern(mp *MeshplayPattern) []byte {
	res, _ := json.Marshal(mp)

	return res
}

func marshalMeshplayPatterns(mps []MeshplayPattern) []byte {
	res, _ := json.Marshal(mps)

	return res
}
