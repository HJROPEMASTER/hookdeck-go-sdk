// This file was auto-generated by Fern from our API Definition.

package issuetriggers

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

func (c *Client) GetIssueTriggers(ctx context.Context, request *hookdeckgosdk.GetIssueTriggersRequest) (*hookdeckgosdk.IssueTriggerPaginatedResult, error) {
	baseURL := "https://api.hookdeck.com/2023-07-01"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	endpointURL := baseURL + "/" + "issue-triggers"

	queryParams := make(url.Values)
	if request.Name != nil {
		queryParams.Add("name", fmt.Sprintf("%v", *request.Name))
	}
	if request.Type != nil {
		queryParams.Add("type", fmt.Sprintf("%v", *request.Type))
	}
	if request.DisabledAt != nil {
		queryParams.Add("disabled_at", fmt.Sprintf("%v", request.DisabledAt.Format(time.RFC3339)))
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

	var response *hookdeckgosdk.IssueTriggerPaginatedResult
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

func (c *Client) CreateIssueTrigger(ctx context.Context, request *hookdeckgosdk.CreateIssueTriggerRequest) (*hookdeckgosdk.IssueTrigger, error) {
	baseURL := "https://api.hookdeck.com/2023-07-01"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	endpointURL := baseURL + "/" + "issue-triggers"

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

	var response *hookdeckgosdk.IssueTrigger
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

func (c *Client) UpsertIssueTrigger(ctx context.Context, request *hookdeckgosdk.UpsertIssueTriggerRequest) (*hookdeckgosdk.IssueTrigger, error) {
	baseURL := "https://api.hookdeck.com/2023-07-01"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	endpointURL := baseURL + "/" + "issue-triggers"

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

	var response *hookdeckgosdk.IssueTrigger
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

func (c *Client) GetIssueTrigger(ctx context.Context, id string) (*hookdeckgosdk.IssueTrigger, error) {
	baseURL := "https://api.hookdeck.com/2023-07-01"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	endpointURL := fmt.Sprintf(baseURL+"/"+"issue-triggers/%v", id)

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

	var response *hookdeckgosdk.IssueTrigger
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

func (c *Client) UpdateIssueTrigger(ctx context.Context, id string, request *hookdeckgosdk.UpdateIssueTriggerRequest) (*hookdeckgosdk.IssueTrigger, error) {
	baseURL := "https://api.hookdeck.com/2023-07-01"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	endpointURL := fmt.Sprintf(baseURL+"/"+"issue-triggers/%v", id)

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

	var response *hookdeckgosdk.IssueTrigger
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

func (c *Client) DeleteIssueTrigger(ctx context.Context, id string) (*hookdeckgosdk.DeletedIssueTriggerResponse, error) {
	baseURL := "https://api.hookdeck.com/2023-07-01"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	endpointURL := fmt.Sprintf(baseURL+"/"+"issue-triggers/%v", id)

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

	var response *hookdeckgosdk.DeletedIssueTriggerResponse
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

func (c *Client) DisableIssueTrigger(ctx context.Context, id string) (*hookdeckgosdk.IssueTrigger, error) {
	baseURL := "https://api.hookdeck.com/2023-07-01"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	endpointURL := fmt.Sprintf(baseURL+"/"+"issue-triggers/%v/disable", id)

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

	var response *hookdeckgosdk.IssueTrigger
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

func (c *Client) EnableIssueTrigger(ctx context.Context, id string) (*hookdeckgosdk.IssueTrigger, error) {
	baseURL := "https://api.hookdeck.com/2023-07-01"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	endpointURL := fmt.Sprintf(baseURL+"/"+"issue-triggers/%v/enable", id)

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

	var response *hookdeckgosdk.IssueTrigger
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
