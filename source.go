// This file was auto-generated by Fern from our API Definition.

package api

import (
	fmt "fmt"
	core "github.com/hookdeck/hookdeck-go-sdk/core"
	time "time"
)

type SourceCreateRequest struct {
	// A unique name for the source
	Name string `json:"name"`
	// Description for the source
	Description        *core.Optional[string]                  `json:"description,omitempty"`
	AllowedHttpMethods *core.Optional[SourceAllowedHttpMethod] `json:"allowed_http_methods,omitempty"`
	CustomResponse     *core.Optional[SourceCustomResponse]    `json:"custom_response,omitempty"`
	Verification       *core.Optional[VerificationConfig]      `json:"verification,omitempty"`
}

type SourceListRequest struct {
	Id         []*string                 `json:"-"`
	Name       *string                   `json:"-"`
	Disabled   *bool                     `json:"-"`
	DisabledAt *time.Time                `json:"-"`
	OrderBy    *SourceListRequestOrderBy `json:"-"`
	Dir        *SourceListRequestDir     `json:"-"`
	Limit      *int                      `json:"-"`
	Next       *string                   `json:"-"`
	Prev       *string                   `json:"-"`
}

type SourceRetrieveRequest struct {
	Include *string `json:"-"`
}

type SourceDeleteResponse struct {
	// ID of the source
	Id string `json:"id"`
}

type SourceListRequestDir string

const (
	SourceListRequestDirAsc  SourceListRequestDir = "asc"
	SourceListRequestDirDesc SourceListRequestDir = "desc"
)

func NewSourceListRequestDirFromString(s string) (SourceListRequestDir, error) {
	switch s {
	case "asc":
		return SourceListRequestDirAsc, nil
	case "desc":
		return SourceListRequestDirDesc, nil
	}
	var t SourceListRequestDir
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (s SourceListRequestDir) Ptr() *SourceListRequestDir {
	return &s
}

type SourceListRequestOrderBy string

const (
	SourceListRequestOrderByCreatedAt SourceListRequestOrderBy = "created_at"
)

func NewSourceListRequestOrderByFromString(s string) (SourceListRequestOrderBy, error) {
	switch s {
	case "created_at":
		return SourceListRequestOrderByCreatedAt, nil
	}
	var t SourceListRequestOrderBy
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (s SourceListRequestOrderBy) Ptr() *SourceListRequestOrderBy {
	return &s
}

type SourceUpdateRequest struct {
	// A unique name for the source
	Name *core.Optional[string] `json:"name,omitempty"`
	// Description for the source
	Description        *core.Optional[string]                  `json:"description,omitempty"`
	AllowedHttpMethods *core.Optional[SourceAllowedHttpMethod] `json:"allowed_http_methods,omitempty"`
	CustomResponse     *core.Optional[SourceCustomResponse]    `json:"custom_response,omitempty"`
	Verification       *core.Optional[VerificationConfig]      `json:"verification,omitempty"`
}

type SourceUpsertRequest struct {
	// A unique name for the source
	Name string `json:"name"`
	// Description for the source
	Description        *core.Optional[string]                  `json:"description,omitempty"`
	AllowedHttpMethods *core.Optional[SourceAllowedHttpMethod] `json:"allowed_http_methods,omitempty"`
	CustomResponse     *core.Optional[SourceCustomResponse]    `json:"custom_response,omitempty"`
	Verification       *core.Optional[VerificationConfig]      `json:"verification,omitempty"`
}
