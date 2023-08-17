// This file was auto-generated by Fern from our API Definition.

package api

import (
	time "time"
)

type GetRequestEventsRequest struct {
	Id             *string                         `json:"-"`
	Status         *EventStatus                    `json:"-"`
	WebhookId      *string                         `json:"-"`
	DestinationId  *string                         `json:"-"`
	SourceId       *string                         `json:"-"`
	Attempts       *int                            `json:"-"`
	ResponseStatus *int                            `json:"-"`
	SuccessfulAt   *time.Time                      `json:"-"`
	CreatedAt      *time.Time                      `json:"-"`
	ErrorCode      *AttemptErrorCodes              `json:"-"`
	CliId          *string                         `json:"-"`
	LastAttemptAt  *time.Time                      `json:"-"`
	SearchTerm     *string                         `json:"-"`
	Headers        *string                         `json:"-"`
	Body           *string                         `json:"-"`
	ParsedQuery    *string                         `json:"-"`
	Path           *string                         `json:"-"`
	CliUserId      *string                         `json:"-"`
	IssueId        *string                         `json:"-"`
	EventDataId    *string                         `json:"-"`
	BulkRetryId    *string                         `json:"-"`
	Include        *string                         `json:"-"`
	OrderBy        *GetRequestEventsRequestOrderBy `json:"-"`
	Dir            *GetRequestEventsRequestDir     `json:"-"`
	Limit          *int                            `json:"-"`
	Next           *string                         `json:"-"`
	Prev           *string                         `json:"-"`
}

type GetRequestIgnoredEventsRequest struct {
	Id      *string                            `json:"-"`
	OrderBy *string                            `json:"-"`
	Dir     *GetRequestIgnoredEventsRequestDir `json:"-"`
	Limit   *int                               `json:"-"`
	Next    *string                            `json:"-"`
	Prev    *string                            `json:"-"`
}

type GetRequestsRequest struct {
	Id             *string                    `json:"-"`
	Status         *GetRequestsRequestStatus  `json:"-"`
	RejectionCause *RequestRejectionCause     `json:"-"`
	SourceId       *string                    `json:"-"`
	Verified       *bool                      `json:"-"`
	SearchTerm     *string                    `json:"-"`
	Headers        *string                    `json:"-"`
	Body           *string                    `json:"-"`
	ParsedQuery    *string                    `json:"-"`
	Path           *string                    `json:"-"`
	IgnoredCount   *int                       `json:"-"`
	EventsCount    *int                       `json:"-"`
	IngestedAt     *time.Time                 `json:"-"`
	BulkRetryId    *string                    `json:"-"`
	Include        *string                    `json:"-"`
	OrderBy        *GetRequestsRequestOrderBy `json:"-"`
	Dir            *GetRequestsRequestDir     `json:"-"`
	Limit          *int                       `json:"-"`
	Next           *string                    `json:"-"`
	Prev           *string                    `json:"-"`
}

type RetryRequestRequest struct {
	// Subset of webhook_ids to re-run the event logic on. Useful to retry only specific ignored_events
	WebhookIds []string `json:"webhook_ids,omitempty"`
}