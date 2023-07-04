// This file was auto-generated by Fern from our API Definition.

package api

type EventAttempt struct {
	// Attempt ID
	Id string `json:"id,omitempty"`
	// Team ID
	TeamId string `json:"team_id,omitempty"`
	// Event ID
	EventId string `json:"event_id,omitempty"`
	// Attempt's HTTP response code
	ResponseStatus *int `json:"response_status,omitempty"`
	// Sequential number of attempts (up to and including this one) made for the associated event
	AttemptNumber *int               `json:"attempt_number,omitempty"`
	Trigger       *AttemptTrigger    `json:"trigger,omitempty"`
	ErrorCode     *AttemptErrorCodes `json:"error_code,omitempty"`
	// Response body from the destination
	Body *EventAttemptBody `json:"body,omitempty"`
	// URL of the destination where delivery was attempted
	RequestedUrl *string `json:"requested_url,omitempty"`
	// ID of associated bulk retry
	BulkRetryId *string       `json:"bulk_retry_id,omitempty"`
	Status      AttemptStatus `json:"status,omitempty"`
	// Date the attempt was successful
	SuccessfulAt *string `json:"successful_at,omitempty"`
	// Date the attempt was delivered
	DeliveredAt *string `json:"delivered_at,omitempty"`
	// Date the destination responded to this attempt
	RespondedAt *string `json:"responded_at,omitempty"`
	// Time elapsed between attempt initiation and final delivery (in ms)
	DeliveryLatency *int `json:"delivery_latency,omitempty"`
	// Time elapsed between attempt initiation and a response from the destination (in ms)
	ResponseLatency *int `json:"response_latency,omitempty"`
	// Date the attempt was last updated
	UpdatedAt string `json:"updated_at,omitempty"`
	// Date the attempt was created
	CreatedAt string        `json:"created_at,omitempty"`
	State     *AttemptState `json:"state,omitempty"`
	// Date the attempt was archived
	ArchivedAt    *string `json:"archived_at,omitempty"`
	DestinationId *string `json:"destination_id,omitempty"`
}