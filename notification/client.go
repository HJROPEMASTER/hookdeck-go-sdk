// This file was auto-generated by Fern from our API Definition.

package notification

import (
	context "context"
	hookdeckgosdk "github.com/hookdeck/hookdeck-go-sdk"
	core "github.com/hookdeck/hookdeck-go-sdk/core"
	http "net/http"
)

type Client struct {
	baseURL    string
	httpClient core.HTTPClient
	header     http.Header
}

func NewClient(opts ...core.ClientOption) *Client {
	options := core.NewClientOptions()
	for _, opt := range opts {
		opt(options)
	}
	return &Client{
		baseURL:    options.BaseURL,
		httpClient: options.HTTPClient,
		header:     options.ToHeader(),
	}
}

func (c *Client) Update(ctx context.Context, request *hookdeckgosdk.NotificationUpdateRequest) (*hookdeckgosdk.ToggleWebhookNotifications, error) {
	baseURL := "https://api.hookdeck.com/2024-03-01"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	endpointURL := baseURL + "/" + "notifications/webhooks"

	var response *hookdeckgosdk.ToggleWebhookNotifications
	if err := core.DoRequest(
		ctx,
		c.httpClient,
		endpointURL,
		http.MethodPut,
		request,
		&response,
		false,
		c.header,
		nil,
	); err != nil {
		return response, err
	}
	return response, nil
}
