// This file was auto-generated by Fern from our API Definition.

package api

type ToggleWebhookNotificationsRequest struct {
	// Enable or disable webhook notifications on the workspace
	Enabled *bool `json:"enabled,omitempty"`
	// List of topics to send notifications for
	Topics *[]TopicsValue `json:"topics,omitempty"`
	// The Hookdeck Source to send the webhook to
	SourceId *string `json:"source_id,omitempty"`
}