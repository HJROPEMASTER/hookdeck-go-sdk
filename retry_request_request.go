// This file was auto-generated by Fern from our API Definition.

package api

type RetryRequestRequest struct {
	// Subset of webhook_ids to re-run the event logic on. Useful to retry only specific ignored_events
	WebhookIds []string `json:"webhook_ids,omitempty"`
}