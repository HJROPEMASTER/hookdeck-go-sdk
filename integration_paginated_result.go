// This file was auto-generated by Fern from our API Definition.

package api

type IntegrationPaginatedResult struct {
	Pagination *SeekPagination `json:"pagination,omitempty"`
	Count      *int            `json:"count,omitempty"`
	Models     *[]*Integration `json:"models,omitempty"`
}