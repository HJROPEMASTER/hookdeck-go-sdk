// This file was auto-generated by Fern from our API Definition.

package api

import (
	bytes "bytes"
	context "context"
	json "encoding/json"
	errors "errors"
	fmt "fmt"
	core "github.com/fern-hookdeck/hookdeck-go/core"
	io "io"
	http "net/http"
	url "net/url"
)

type RulesetsClient interface {
	GetRulesets(ctx context.Context, request *GetRulesetsRequest) (*RulesetPaginatedResult, error)
	CreateRuleset(ctx context.Context, request *CreateRulesetRequest) (*Ruleset, error)
	UpsertRuleset(ctx context.Context, request *UpsertRulesetRequest) (*Ruleset, error)
	GetRuleset(ctx context.Context, id string) (*Ruleset, error)
	UpdateRuleset(ctx context.Context, id string, request *UpdateRulesetRequest) (*Ruleset, error)
	ArchiveRuleset(ctx context.Context, id string) (*Ruleset, error)
	UnarchiveRuleset(ctx context.Context, id string) (*Ruleset, error)
}

func NewRulesetsClient(opts ...core.ClientOption) RulesetsClient {
	options := core.NewClientOptions()
	for _, opt := range opts {
		opt(options)
	}
	return &rulesetsClient{
		baseURL:    options.BaseURL,
		httpClient: options.HTTPClient,
		header:     options.ToHeader(),
	}
}

type rulesetsClient struct {
	baseURL    string
	httpClient core.HTTPClient
	header     http.Header
}

func (r *rulesetsClient) GetRulesets(ctx context.Context, request *GetRulesetsRequest) (*RulesetPaginatedResult, error) {
	baseURL := "https://api.hookdeck.com/2023-01-01"
	if r.baseURL != "" {
		baseURL = r.baseURL
	}
	endpointURL := baseURL + "/" + "rulesets"

	queryParams := make(url.Values)
	var idDefaultValue *string
	if request.Id != idDefaultValue {
		queryParams.Add("id", fmt.Sprintf("%v", *request.Id))
	}
	var nameDefaultValue *string
	if request.Name != nameDefaultValue {
		queryParams.Add("name", fmt.Sprintf("%v", *request.Name))
	}
	var archivedDefaultValue *bool
	if request.Archived != archivedDefaultValue {
		queryParams.Add("archived", fmt.Sprintf("%v", *request.Archived))
	}
	var archivedAtDefaultValue *string
	if request.ArchivedAt != archivedAtDefaultValue {
		queryParams.Add("archived_at", fmt.Sprintf("%v", *request.ArchivedAt))
	}
	var orderByDefaultValue *string
	if request.OrderBy != orderByDefaultValue {
		queryParams.Add("order_by", fmt.Sprintf("%v", *request.OrderBy))
	}
	var dirDefaultValue *GetRulesetsRequestDirGetRulesetsRequestDir
	if request.Dir != dirDefaultValue {
		queryParams.Add("dir", fmt.Sprintf("%v", *request.Dir))
	}
	var limitDefaultValue *int
	if request.Limit != limitDefaultValue {
		queryParams.Add("limit", fmt.Sprintf("%v", *request.Limit))
	}
	var nextDefaultValue *string
	if request.Next != nextDefaultValue {
		queryParams.Add("next", fmt.Sprintf("%v", *request.Next))
	}
	var prevDefaultValue *string
	if request.Prev != prevDefaultValue {
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
			value := new(BadRequestError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return err
			}
			return value
		case 422:
			value := new(UnprocessableEntityError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return err
			}
			return value
		}
		return apiError
	}

	response := new(RulesetPaginatedResult)
	if err := core.DoRequest(
		ctx,
		r.httpClient,
		endpointURL,
		http.MethodGet,
		request,
		&response,
		r.header,
		errorDecoder,
	); err != nil {
		return response, err
	}
	return response, nil
}

func (r *rulesetsClient) CreateRuleset(ctx context.Context, request *CreateRulesetRequest) (*Ruleset, error) {
	baseURL := "https://api.hookdeck.com/2023-01-01"
	if r.baseURL != "" {
		baseURL = r.baseURL
	}
	endpointURL := baseURL + "/" + "rulesets"

	errorDecoder := func(statusCode int, body io.Reader) error {
		raw, err := io.ReadAll(body)
		if err != nil {
			return err
		}
		apiError := core.NewAPIError(statusCode, errors.New(string(raw)))
		decoder := json.NewDecoder(bytes.NewReader(raw))
		switch statusCode {
		case 400:
			value := new(BadRequestError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return err
			}
			return value
		case 422:
			value := new(UnprocessableEntityError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return err
			}
			return value
		}
		return apiError
	}

	var response Ruleset
	if err := core.DoRequest(
		ctx,
		r.httpClient,
		endpointURL,
		http.MethodPost,
		request,
		&response,
		r.header,
		errorDecoder,
	); err != nil {
		return response, err
	}
	return response, nil
}

func (r *rulesetsClient) UpsertRuleset(ctx context.Context, request *UpsertRulesetRequest) (*Ruleset, error) {
	baseURL := "https://api.hookdeck.com/2023-01-01"
	if r.baseURL != "" {
		baseURL = r.baseURL
	}
	endpointURL := baseURL + "/" + "rulesets"

	errorDecoder := func(statusCode int, body io.Reader) error {
		raw, err := io.ReadAll(body)
		if err != nil {
			return err
		}
		apiError := core.NewAPIError(statusCode, errors.New(string(raw)))
		decoder := json.NewDecoder(bytes.NewReader(raw))
		switch statusCode {
		case 400:
			value := new(BadRequestError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return err
			}
			return value
		case 422:
			value := new(UnprocessableEntityError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return err
			}
			return value
		}
		return apiError
	}

	var response Ruleset
	if err := core.DoRequest(
		ctx,
		r.httpClient,
		endpointURL,
		http.MethodPut,
		request,
		&response,
		r.header,
		errorDecoder,
	); err != nil {
		return response, err
	}
	return response, nil
}

func (r *rulesetsClient) GetRuleset(ctx context.Context, id string) (*Ruleset, error) {
	baseURL := "https://api.hookdeck.com/2023-01-01"
	if r.baseURL != "" {
		baseURL = r.baseURL
	}
	endpointURL := fmt.Sprintf(baseURL+"/"+"rulesets/%v", id)

	errorDecoder := func(statusCode int, body io.Reader) error {
		raw, err := io.ReadAll(body)
		if err != nil {
			return err
		}
		apiError := core.NewAPIError(statusCode, errors.New(string(raw)))
		decoder := json.NewDecoder(bytes.NewReader(raw))
		switch statusCode {
		case 404:
			value := new(NotFoundError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return err
			}
			return value
		}
		return apiError
	}

	var response Ruleset
	if err := core.DoRequest(
		ctx,
		r.httpClient,
		endpointURL,
		http.MethodGet,
		nil,
		&response,
		r.header,
		errorDecoder,
	); err != nil {
		return response, err
	}
	return response, nil
}

func (r *rulesetsClient) UpdateRuleset(ctx context.Context, id string, request *UpdateRulesetRequest) (*Ruleset, error) {
	baseURL := "https://api.hookdeck.com/2023-01-01"
	if r.baseURL != "" {
		baseURL = r.baseURL
	}
	endpointURL := fmt.Sprintf(baseURL+"/"+"rulesets/%v", id)

	errorDecoder := func(statusCode int, body io.Reader) error {
		raw, err := io.ReadAll(body)
		if err != nil {
			return err
		}
		apiError := core.NewAPIError(statusCode, errors.New(string(raw)))
		decoder := json.NewDecoder(bytes.NewReader(raw))
		switch statusCode {
		case 400:
			value := new(BadRequestError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return err
			}
			return value
		case 404:
			value := new(NotFoundError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return err
			}
			return value
		case 422:
			value := new(UnprocessableEntityError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return err
			}
			return value
		}
		return apiError
	}

	var response Ruleset
	if err := core.DoRequest(
		ctx,
		r.httpClient,
		endpointURL,
		http.MethodPut,
		request,
		&response,
		r.header,
		errorDecoder,
	); err != nil {
		return response, err
	}
	return response, nil
}

func (r *rulesetsClient) ArchiveRuleset(ctx context.Context, id string) (*Ruleset, error) {
	baseURL := "https://api.hookdeck.com/2023-01-01"
	if r.baseURL != "" {
		baseURL = r.baseURL
	}
	endpointURL := fmt.Sprintf(baseURL+"/"+"rulesets/%v/archive", id)

	errorDecoder := func(statusCode int, body io.Reader) error {
		raw, err := io.ReadAll(body)
		if err != nil {
			return err
		}
		apiError := core.NewAPIError(statusCode, errors.New(string(raw)))
		decoder := json.NewDecoder(bytes.NewReader(raw))
		switch statusCode {
		case 404:
			value := new(NotFoundError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return err
			}
			return value
		}
		return apiError
	}

	var response Ruleset
	if err := core.DoRequest(
		ctx,
		r.httpClient,
		endpointURL,
		http.MethodPut,
		nil,
		&response,
		r.header,
		errorDecoder,
	); err != nil {
		return response, err
	}
	return response, nil
}

func (r *rulesetsClient) UnarchiveRuleset(ctx context.Context, id string) (*Ruleset, error) {
	baseURL := "https://api.hookdeck.com/2023-01-01"
	if r.baseURL != "" {
		baseURL = r.baseURL
	}
	endpointURL := fmt.Sprintf(baseURL+"/"+"rulesets/%v/unarchive", id)

	errorDecoder := func(statusCode int, body io.Reader) error {
		raw, err := io.ReadAll(body)
		if err != nil {
			return err
		}
		apiError := core.NewAPIError(statusCode, errors.New(string(raw)))
		decoder := json.NewDecoder(bytes.NewReader(raw))
		switch statusCode {
		case 404:
			value := new(NotFoundError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return err
			}
			return value
		}
		return apiError
	}

	var response Ruleset
	if err := core.DoRequest(
		ctx,
		r.httpClient,
		endpointURL,
		http.MethodPut,
		nil,
		&response,
		r.header,
		errorDecoder,
	); err != nil {
		return response, err
	}
	return response, nil
}