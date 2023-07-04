// This file was auto-generated by Fern from our API Definition.

package api

// Filter properties for the events to be included in the bulk retry, use query parameters of [Requests](#requests)
type CreateRequestBulkRetryRequestQuery struct {
	// Filter by requests IDs
	Id     *CreateRequestBulkRetryRequestQueryId     `json:"id,omitempty"`
	Status *CreateRequestBulkRetryRequestQueryStatus `json:"status,omitempty"`
	// Filter by rejection cause
	RejectionCause *CreateRequestBulkRetryRequestQueryRejectionCause `json:"rejection_cause,omitempty"`
	IgnoredCount   *CreateRequestBulkRetryRequestQueryIgnoredCount   `json:"ignored_count,omitempty"`
	EventsCount    *CreateRequestBulkRetryRequestQueryEventsCount    `json:"events_count,omitempty"`
	// Filter by source IDs
	SourceId *CreateRequestBulkRetryRequestQuerySourceId `json:"source_id,omitempty"`
	// Filter by verification status
	Verified *bool `json:"verified,omitempty"`
	// URL Encoded string of the JSON to match to the data headers
	Headers *CreateRequestBulkRetryRequestQueryHeaders `json:"headers,omitempty"`
	// URL Encoded string of the JSON to match to the data body
	Body *CreateRequestBulkRetryRequestQueryBody `json:"body,omitempty"`
	// URL Encoded string of the JSON to match to the parsed query (JSON representation of the query)
	ParsedQuery *CreateRequestBulkRetryRequestQueryParsedQuery `json:"parsed_query,omitempty"`
	// URL Encoded string of the string to match partially to the path
	Path        *string                                        `json:"path,omitempty"`
	CreatedAt   *CreateRequestBulkRetryRequestQueryCreatedAt   `json:"created_at,omitempty"`
	IngestedAt  *CreateRequestBulkRetryRequestQueryIngestedAt  `json:"ingested_at,omitempty"`
	BulkRetryId *CreateRequestBulkRetryRequestQueryBulkRetryId `json:"bulk_retry_id,omitempty"`
}