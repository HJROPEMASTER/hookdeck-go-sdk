// This file was auto-generated by Fern from our API Definition.

package request

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

func (c *Client) List(ctx context.Context, request *hookdeckgosdk.RequestListRequest) (*hookdeckgosdk.RequestPaginatedResult, error) {
	baseURL := "https://api.hookdeck.com/2024-03-01"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	endpointURL := baseURL + "/" + "requests"

	queryParams := make(url.Values)
	if request.Id != nil {
		queryParams.Add("id", fmt.Sprintf("%v", *request.Id))
	}
	if request.Status != nil {
		queryParams.Add("status", fmt.Sprintf("%v", *request.Status))
	}
	if request.RejectionCause != nil {
		queryParams.Add("rejection_cause", fmt.Sprintf("%v", *request.RejectionCause))
	}
	if request.SourceId != nil {
		queryParams.Add("source_id", fmt.Sprintf("%v", *request.SourceId))
	}
	if request.Verified != nil {
		queryParams.Add("verified", fmt.Sprintf("%v", *request.Verified))
	}
	if request.SearchTerm != nil {
		queryParams.Add("search_term", fmt.Sprintf("%v", *request.SearchTerm))
	}
	if request.Headers != nil {
		queryParams.Add("headers", fmt.Sprintf("%v", *request.Headers))
	}
	if request.Body != nil {
		queryParams.Add("body", fmt.Sprintf("%v", *request.Body))
	}
	if request.ParsedQuery != nil {
		queryParams.Add("parsed_query", fmt.Sprintf("%v", *request.ParsedQuery))
	}
	if request.Path != nil {
		queryParams.Add("path", fmt.Sprintf("%v", *request.Path))
	}
	if request.IgnoredCount != nil {
		queryParams.Add("ignored_count", fmt.Sprintf("%v", *request.IgnoredCount))
	}
	if request.EventsCount != nil {
		queryParams.Add("events_count", fmt.Sprintf("%v", *request.EventsCount))
	}
	if request.IngestedAt != nil {
		queryParams.Add("ingested_at", fmt.Sprintf("%v", request.IngestedAt.Format(time.RFC3339)))
	}
	if request.BulkRetryId != nil {
		queryParams.Add("bulk_retry_id", fmt.Sprintf("%v", *request.BulkRetryId))
	}
	if request.Include != nil {
		queryParams.Add("include", fmt.Sprintf("%v", *request.Include))
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

	var response *hookdeckgosdk.RequestPaginatedResult
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

func (c *Client) Retrieve(ctx context.Context, id string) (*hookdeckgosdk.Request, error) {
	baseURL := "https://api.hookdeck.com/2024-03-01"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	endpointURL := fmt.Sprintf(baseURL+"/"+"requests/%v", id)

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

	var response *hookdeckgosdk.Request
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

func (c *Client) RetrieveBody(ctx context.Context, id string) (*hookdeckgosdk.RawBody, error) {
	baseURL := "https://api.hookdeck.com/2024-03-01"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	endpointURL := fmt.Sprintf(baseURL+"/"+"requests/%v/raw_body", id)

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

	var response *hookdeckgosdk.RawBody
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

func (c *Client) Retry(ctx context.Context, id string, request *hookdeckgosdk.RequestRetryRequest) (*hookdeckgosdk.RetryRequest, error) {
	baseURL := "https://api.hookdeck.com/2024-03-01"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	endpointURL := fmt.Sprintf(baseURL+"/"+"requests/%v/retry", id)

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

	var response *hookdeckgosdk.RetryRequest
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

func (c *Client) ListEvent(ctx context.Context, id string, request *hookdeckgosdk.RequestListEventRequest) (*hookdeckgosdk.EventPaginatedResult, error) {
	baseURL := "https://api.hookdeck.com/2024-03-01"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	endpointURL := fmt.Sprintf(baseURL+"/"+"requests/%v/events", id)

	queryParams := make(url.Values)
	if request.Id != nil {
		queryParams.Add("id", fmt.Sprintf("%v", *request.Id))
	}
	if request.Status != nil {
		queryParams.Add("status", fmt.Sprintf("%v", *request.Status))
	}
	if request.WebhookId != nil {
		queryParams.Add("webhook_id", fmt.Sprintf("%v", *request.WebhookId))
	}
	if request.DestinationId != nil {
		queryParams.Add("destination_id", fmt.Sprintf("%v", *request.DestinationId))
	}
	if request.SourceId != nil {
		queryParams.Add("source_id", fmt.Sprintf("%v", *request.SourceId))
	}
	if request.Attempts != nil {
		queryParams.Add("attempts", fmt.Sprintf("%v", *request.Attempts))
	}
	if request.ResponseStatus != nil {
		queryParams.Add("response_status", fmt.Sprintf("%v", *request.ResponseStatus))
	}
	if request.SuccessfulAt != nil {
		queryParams.Add("successful_at", fmt.Sprintf("%v", request.SuccessfulAt.Format(time.RFC3339)))
	}
	if request.CreatedAt != nil {
		queryParams.Add("created_at", fmt.Sprintf("%v", request.CreatedAt.Format(time.RFC3339)))
	}
	if request.ErrorCode != nil {
		queryParams.Add("error_code", fmt.Sprintf("%v", *request.ErrorCode))
	}
	if request.CliId != nil {
		queryParams.Add("cli_id", fmt.Sprintf("%v", *request.CliId))
	}
	if request.LastAttemptAt != nil {
		queryParams.Add("last_attempt_at", fmt.Sprintf("%v", request.LastAttemptAt.Format(time.RFC3339)))
	}
	if request.SearchTerm != nil {
		queryParams.Add("search_term", fmt.Sprintf("%v", *request.SearchTerm))
	}
	if request.Headers != nil {
		queryParams.Add("headers", fmt.Sprintf("%v", *request.Headers))
	}
	if request.Body != nil {
		queryParams.Add("body", fmt.Sprintf("%v", *request.Body))
	}
	if request.ParsedQuery != nil {
		queryParams.Add("parsed_query", fmt.Sprintf("%v", *request.ParsedQuery))
	}
	if request.Path != nil {
		queryParams.Add("path", fmt.Sprintf("%v", *request.Path))
	}
	if request.CliUserId != nil {
		queryParams.Add("cli_user_id", fmt.Sprintf("%v", *request.CliUserId))
	}
	if request.IssueId != nil {
		queryParams.Add("issue_id", fmt.Sprintf("%v", *request.IssueId))
	}
	if request.EventDataId != nil {
		queryParams.Add("event_data_id", fmt.Sprintf("%v", *request.EventDataId))
	}
	if request.BulkRetryId != nil {
		queryParams.Add("bulk_retry_id", fmt.Sprintf("%v", *request.BulkRetryId))
	}
	if request.Include != nil {
		queryParams.Add("include", fmt.Sprintf("%v", *request.Include))
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

	var response *hookdeckgosdk.EventPaginatedResult
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

func (c *Client) ListIgnoredEvent(ctx context.Context, id string, request *hookdeckgosdk.RequestListIgnoredEventRequest) (*hookdeckgosdk.IgnoredEventPaginatedResult, error) {
	baseURL := "https://api.hookdeck.com/2024-03-01"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	endpointURL := fmt.Sprintf(baseURL+"/"+"requests/%v/ignored_events", id)

	queryParams := make(url.Values)
	if request.Id != nil {
		queryParams.Add("id", fmt.Sprintf("%v", *request.Id))
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

	var response *hookdeckgosdk.IgnoredEventPaginatedResult
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
