// This file was auto-generated by Fern from our API Definition.

package connections

import (
	bytes "bytes"
	context "context"
	json "encoding/json"
	errors "errors"
	fmt "fmt"
	hookdeckgosdk "github.com/hookdeck/hookdeck-go-sdk"
	core "github.com/hookdeck/hookdeck-go-sdk/core"
	io "io"
	http "net/http"
	url "net/url"
	time "time"
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

func (c *Client) GetConnections(ctx context.Context, request *hookdeckgosdk.GetConnectionsRequest) (*hookdeckgosdk.ConnectionPaginatedResult, error) {
	baseURL := "https://api.hookdeck.com/2023-07-01"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	endpointURL := baseURL + "/" + "connections"

	queryParams := make(url.Values)
	if request.Id != nil {
		queryParams.Add("id", fmt.Sprintf("%v", *request.Id))
	}
	if request.Name != nil {
		queryParams.Add("name", fmt.Sprintf("%v", *request.Name))
	}
	if request.DestinationId != nil {
		queryParams.Add("destination_id", fmt.Sprintf("%v", *request.DestinationId))
	}
	if request.SourceId != nil {
		queryParams.Add("source_id", fmt.Sprintf("%v", *request.SourceId))
	}
	if request.Archived != nil {
		queryParams.Add("archived", fmt.Sprintf("%v", *request.Archived))
	}
	if request.ArchivedAt != nil {
		queryParams.Add("archived_at", fmt.Sprintf("%v", request.ArchivedAt.Format(time.RFC3339)))
	}
	if request.FullName != nil {
		queryParams.Add("full_name", fmt.Sprintf("%v", *request.FullName))
	}
	if request.PausedAt != nil {
		queryParams.Add("paused_at", fmt.Sprintf("%v", request.PausedAt.Format(time.RFC3339)))
	}
	if request.OrderBy != nil {
		queryParams.Add("order_by", fmt.Sprintf("%v", *request.OrderBy))
	}
	if request.Dir != nil {
		queryParams.Add("dir", fmt.Sprintf("%v", *request.Dir))
	}
	if request.Limit != nil {
		queryParams.Add("limit", fmt.Sprintf("%v", *request.Limit))
	}
	if request.Next != nil {
		queryParams.Add("next", fmt.Sprintf("%v", *request.Next))
	}
	if request.Prev != nil {
		queryParams.Add("prev", fmt.Sprintf("%v", *request.Prev))
	}
	if len(queryParams) > 0 {
		endpointURL += "?" + queryParams.Encode()
	}

	errorDecoder := func(statusCode int, body io.Reader) error {
		raw, err := io.ReadAll(body)
		if err != nil {
			return err
		}
		apiError := core.NewAPIError(statusCode, errors.New(string(raw)))
		decoder := json.NewDecoder(bytes.NewReader(raw))
		switch statusCode {
		case 400:
			value := new(hookdeckgosdk.BadRequestError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		case 422:
			value := new(hookdeckgosdk.UnprocessableEntityError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		}
		return apiError
	}

	var response *hookdeckgosdk.ConnectionPaginatedResult
	if err := core.DoRequest(
		ctx,
		c.httpClient,
		endpointURL,
		http.MethodGet,
		request,
		&response,
		false,
		c.header,
		errorDecoder,
	); err != nil {
		return response, err
	}
	return response, nil
}

func (c *Client) CreateConnection(ctx context.Context, request *hookdeckgosdk.CreateConnectionRequest) (*hookdeckgosdk.Connection, error) {
	baseURL := "https://api.hookdeck.com/2023-07-01"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	endpointURL := baseURL + "/" + "connections"

	errorDecoder := func(statusCode int, body io.Reader) error {
		raw, err := io.ReadAll(body)
		if err != nil {
			return err
		}
		apiError := core.NewAPIError(statusCode, errors.New(string(raw)))
		decoder := json.NewDecoder(bytes.NewReader(raw))
		switch statusCode {
		case 400:
			value := new(hookdeckgosdk.BadRequestError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		case 422:
			value := new(hookdeckgosdk.UnprocessableEntityError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		}
		return apiError
	}

	var response *hookdeckgosdk.Connection
	if err := core.DoRequest(
		ctx,
		c.httpClient,
		endpointURL,
		http.MethodPost,
		request,
		&response,
		false,
		c.header,
		errorDecoder,
	); err != nil {
		return response, err
	}
	return response, nil
}

func (c *Client) UpsertConnection(ctx context.Context, request *hookdeckgosdk.UpsertConnectionRequest) (*hookdeckgosdk.Connection, error) {
	baseURL := "https://api.hookdeck.com/2023-07-01"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	endpointURL := baseURL + "/" + "connections"

	errorDecoder := func(statusCode int, body io.Reader) error {
		raw, err := io.ReadAll(body)
		if err != nil {
			return err
		}
		apiError := core.NewAPIError(statusCode, errors.New(string(raw)))
		decoder := json.NewDecoder(bytes.NewReader(raw))
		switch statusCode {
		case 400:
			value := new(hookdeckgosdk.BadRequestError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		case 422:
			value := new(hookdeckgosdk.UnprocessableEntityError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		}
		return apiError
	}

	var response *hookdeckgosdk.Connection
	if err := core.DoRequest(
		ctx,
		c.httpClient,
		endpointURL,
		http.MethodPut,
		request,
		&response,
		false,
		c.header,
		errorDecoder,
	); err != nil {
		return response, err
	}
	return response, nil
}

func (c *Client) GetConnection(ctx context.Context, id string) (*hookdeckgosdk.Connection, error) {
	baseURL := "https://api.hookdeck.com/2023-07-01"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	endpointURL := fmt.Sprintf(baseURL+"/"+"connections/%v", id)

	errorDecoder := func(statusCode int, body io.Reader) error {
		raw, err := io.ReadAll(body)
		if err != nil {
			return err
		}
		apiError := core.NewAPIError(statusCode, errors.New(string(raw)))
		decoder := json.NewDecoder(bytes.NewReader(raw))
		switch statusCode {
		case 404:
			value := new(hookdeckgosdk.NotFoundError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		}
		return apiError
	}

	var response *hookdeckgosdk.Connection
	if err := core.DoRequest(
		ctx,
		c.httpClient,
		endpointURL,
		http.MethodGet,
		nil,
		&response,
		false,
		c.header,
		errorDecoder,
	); err != nil {
		return response, err
	}
	return response, nil
}

func (c *Client) UpdateConnection(ctx context.Context, id string, request *hookdeckgosdk.UpdateConnectionRequest) (*hookdeckgosdk.Connection, error) {
	baseURL := "https://api.hookdeck.com/2023-07-01"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	endpointURL := fmt.Sprintf(baseURL+"/"+"connections/%v", id)

	errorDecoder := func(statusCode int, body io.Reader) error {
		raw, err := io.ReadAll(body)
		if err != nil {
			return err
		}
		apiError := core.NewAPIError(statusCode, errors.New(string(raw)))
		decoder := json.NewDecoder(bytes.NewReader(raw))
		switch statusCode {
		case 400:
			value := new(hookdeckgosdk.BadRequestError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		case 404:
			value := new(hookdeckgosdk.NotFoundError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		case 422:
			value := new(hookdeckgosdk.UnprocessableEntityError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		}
		return apiError
	}

	var response *hookdeckgosdk.Connection
	if err := core.DoRequest(
		ctx,
		c.httpClient,
		endpointURL,
		http.MethodPut,
		request,
		&response,
		false,
		c.header,
		errorDecoder,
	); err != nil {
		return response, err
	}
	return response, nil
}

func (c *Client) DeleteConnection(ctx context.Context, id string) (*hookdeckgosdk.DeleteConnectionResponse, error) {
	baseURL := "https://api.hookdeck.com/2023-07-01"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	endpointURL := fmt.Sprintf(baseURL+"/"+"connections/%v", id)

	errorDecoder := func(statusCode int, body io.Reader) error {
		raw, err := io.ReadAll(body)
		if err != nil {
			return err
		}
		apiError := core.NewAPIError(statusCode, errors.New(string(raw)))
		decoder := json.NewDecoder(bytes.NewReader(raw))
		switch statusCode {
		case 404:
			value := new(hookdeckgosdk.NotFoundError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		}
		return apiError
	}

	var response *hookdeckgosdk.DeleteConnectionResponse
	if err := core.DoRequest(
		ctx,
		c.httpClient,
		endpointURL,
		http.MethodDelete,
		nil,
		&response,
		false,
		c.header,
		errorDecoder,
	); err != nil {
		return response, err
	}
	return response, nil
}

func (c *Client) ArchiveConnection(ctx context.Context, id string) (*hookdeckgosdk.Connection, error) {
	baseURL := "https://api.hookdeck.com/2023-07-01"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	endpointURL := fmt.Sprintf(baseURL+"/"+"connections/%v/archive", id)

	errorDecoder := func(statusCode int, body io.Reader) error {
		raw, err := io.ReadAll(body)
		if err != nil {
			return err
		}
		apiError := core.NewAPIError(statusCode, errors.New(string(raw)))
		decoder := json.NewDecoder(bytes.NewReader(raw))
		switch statusCode {
		case 404:
			value := new(hookdeckgosdk.NotFoundError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		}
		return apiError
	}

	var response *hookdeckgosdk.Connection
	if err := core.DoRequest(
		ctx,
		c.httpClient,
		endpointURL,
		http.MethodPut,
		nil,
		&response,
		false,
		c.header,
		errorDecoder,
	); err != nil {
		return response, err
	}
	return response, nil
}

func (c *Client) UnarchiveConnection(ctx context.Context, id string) (*hookdeckgosdk.Connection, error) {
	baseURL := "https://api.hookdeck.com/2023-07-01"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	endpointURL := fmt.Sprintf(baseURL+"/"+"connections/%v/unarchive", id)

	errorDecoder := func(statusCode int, body io.Reader) error {
		raw, err := io.ReadAll(body)
		if err != nil {
			return err
		}
		apiError := core.NewAPIError(statusCode, errors.New(string(raw)))
		decoder := json.NewDecoder(bytes.NewReader(raw))
		switch statusCode {
		case 404:
			value := new(hookdeckgosdk.NotFoundError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		}
		return apiError
	}

	var response *hookdeckgosdk.Connection
	if err := core.DoRequest(
		ctx,
		c.httpClient,
		endpointURL,
		http.MethodPut,
		nil,
		&response,
		false,
		c.header,
		errorDecoder,
	); err != nil {
		return response, err
	}
	return response, nil
}

func (c *Client) PauseConnection(ctx context.Context, id string) (*hookdeckgosdk.Connection, error) {
	baseURL := "https://api.hookdeck.com/2023-07-01"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	endpointURL := fmt.Sprintf(baseURL+"/"+"connections/%v/pause", id)

	errorDecoder := func(statusCode int, body io.Reader) error {
		raw, err := io.ReadAll(body)
		if err != nil {
			return err
		}
		apiError := core.NewAPIError(statusCode, errors.New(string(raw)))
		decoder := json.NewDecoder(bytes.NewReader(raw))
		switch statusCode {
		case 404:
			value := new(hookdeckgosdk.NotFoundError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		}
		return apiError
	}

	var response *hookdeckgosdk.Connection
	if err := core.DoRequest(
		ctx,
		c.httpClient,
		endpointURL,
		http.MethodPut,
		nil,
		&response,
		false,
		c.header,
		errorDecoder,
	); err != nil {
		return response, err
	}
	return response, nil
}

func (c *Client) UnpauseConnection(ctx context.Context, id string) (*hookdeckgosdk.Connection, error) {
	baseURL := "https://api.hookdeck.com/2023-07-01"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	endpointURL := fmt.Sprintf(baseURL+"/"+"connections/%v/unpause", id)

	errorDecoder := func(statusCode int, body io.Reader) error {
		raw, err := io.ReadAll(body)
		if err != nil {
			return err
		}
		apiError := core.NewAPIError(statusCode, errors.New(string(raw)))
		decoder := json.NewDecoder(bytes.NewReader(raw))
		switch statusCode {
		case 404:
			value := new(hookdeckgosdk.NotFoundError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		}
		return apiError
	}

	var response *hookdeckgosdk.Connection
	if err := core.DoRequest(
		ctx,
		c.httpClient,
		endpointURL,
		http.MethodPut,
		nil,
		&response,
		false,
		c.header,
		errorDecoder,
	); err != nil {
		return response, err
	}
	return response, nil
}
