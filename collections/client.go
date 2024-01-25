// This file was auto-generated by Fern from our API Definition.

package collections

import (
	context "context"
	fmt "fmt"
	hookdeckgosdk "github.com/hookdeck/hookdeck-go-sdk"
	core "github.com/hookdeck/hookdeck-go-sdk/core"
	http "net/http"
	url "net/url"
)

type Client struct {
	baseURL string
	caller  *core.Caller
	header  http.Header
}

func NewClient(opts ...core.ClientOption) *Client {
	options := core.NewClientOptions()
	for _, opt := range opts {
		opt(options)
	}
	return &Client{
		baseURL: options.BaseURL,
		caller:  core.NewCaller(options.HTTPClient),
		header:  options.ToHeader(),
	}
}

// Name of the collection to create shards for
func (c *Client) CreateShardKey(ctx context.Context, collectionName string, request *hookdeckgosdk.CreateShardKeyRequest) (*hookdeckgosdk.CreateShardKeyResponse, error) {
	baseURL := "http://localhost:6333"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	endpointURL := fmt.Sprintf(baseURL+"/"+"collections/%v/shards", collectionName)

	queryParams := make(url.Values)
	if request.Timeout != nil {
		queryParams.Add("timeout", fmt.Sprintf("%v", *request.Timeout))
	}
	if len(queryParams) > 0 {
		endpointURL += "?" + queryParams.Encode()
	}

	var response *hookdeckgosdk.CreateShardKeyResponse
	if err := c.caller.Call(
		ctx,
		&core.CallParams{
			URL:      endpointURL,
			Method:   http.MethodPut,
			Headers:  c.header,
			Request:  request,
			Response: &response,
		},
	); err != nil {
		return nil, err
	}
	return response, nil
}

// Name of the collection to create shards for
func (c *Client) DeleteShardKey(ctx context.Context, collectionName string, request *hookdeckgosdk.DeleteShardKeyRequest) (*hookdeckgosdk.DeleteShardKeyResponse, error) {
	baseURL := "http://localhost:6333"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	endpointURL := fmt.Sprintf(baseURL+"/"+"collections/%v/shards/delete", collectionName)

	queryParams := make(url.Values)
	if request.Timeout != nil {
		queryParams.Add("timeout", fmt.Sprintf("%v", *request.Timeout))
	}
	if len(queryParams) > 0 {
		endpointURL += "?" + queryParams.Encode()
	}

	var response *hookdeckgosdk.DeleteShardKeyResponse
	if err := c.caller.Call(
		ctx,
		&core.CallParams{
			URL:      endpointURL,
			Method:   http.MethodPost,
			Headers:  c.header,
			Request:  request,
			Response: &response,
		},
	); err != nil {
		return nil, err
	}
	return response, nil
}

// Get list name of all existing collections
func (c *Client) GetCollections(ctx context.Context) (*hookdeckgosdk.GetCollectionsResponse, error) {
	baseURL := "http://localhost:6333"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	endpointURL := baseURL + "/" + "collections"

	var response *hookdeckgosdk.GetCollectionsResponse
	if err := c.caller.Call(
		ctx,
		&core.CallParams{
			URL:      endpointURL,
			Method:   http.MethodGet,
			Headers:  c.header,
			Response: &response,
		},
	); err != nil {
		return nil, err
	}
	return response, nil
}

// Get detailed information about specified existing collection
//
// Name of the collection to retrieve
func (c *Client) GetCollection(ctx context.Context, collectionName string) (*hookdeckgosdk.GetCollectionResponse, error) {
	baseURL := "http://localhost:6333"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	endpointURL := fmt.Sprintf(baseURL+"/"+"collections/%v", collectionName)

	var response *hookdeckgosdk.GetCollectionResponse
	if err := c.caller.Call(
		ctx,
		&core.CallParams{
			URL:      endpointURL,
			Method:   http.MethodGet,
			Headers:  c.header,
			Response: &response,
		},
	); err != nil {
		return nil, err
	}
	return response, nil
}

// Create new collection with given parameters
//
// Name of the new collection
func (c *Client) CreateCollection(ctx context.Context, collectionName string, request *hookdeckgosdk.CreateCollection) (*hookdeckgosdk.CreateCollectionResponse, error) {
	baseURL := "http://localhost:6333"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	endpointURL := fmt.Sprintf(baseURL+"/"+"collections/%v", collectionName)

	queryParams := make(url.Values)
	if request.Timeout != nil {
		queryParams.Add("timeout", fmt.Sprintf("%v", *request.Timeout))
	}
	if len(queryParams) > 0 {
		endpointURL += "?" + queryParams.Encode()
	}

	var response *hookdeckgosdk.CreateCollectionResponse
	if err := c.caller.Call(
		ctx,
		&core.CallParams{
			URL:      endpointURL,
			Method:   http.MethodPut,
			Headers:  c.header,
			Request:  request,
			Response: &response,
		},
	); err != nil {
		return nil, err
	}
	return response, nil
}

// Drop collection and all associated data
//
// Name of the collection to delete
func (c *Client) DeleteCollection(ctx context.Context, collectionName string, request *hookdeckgosdk.DeleteCollectionRequest) (*hookdeckgosdk.DeleteCollectionResponse, error) {
	baseURL := "http://localhost:6333"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	endpointURL := fmt.Sprintf(baseURL+"/"+"collections/%v", collectionName)

	queryParams := make(url.Values)
	if request.Timeout != nil {
		queryParams.Add("timeout", fmt.Sprintf("%v", *request.Timeout))
	}
	if len(queryParams) > 0 {
		endpointURL += "?" + queryParams.Encode()
	}

	var response *hookdeckgosdk.DeleteCollectionResponse
	if err := c.caller.Call(
		ctx,
		&core.CallParams{
			URL:      endpointURL,
			Method:   http.MethodDelete,
			Headers:  c.header,
			Response: &response,
		},
	); err != nil {
		return nil, err
	}
	return response, nil
}

// Update parameters of the existing collection
//
// Name of the collection to update
func (c *Client) UpdateCollection(ctx context.Context, collectionName string, request *hookdeckgosdk.UpdateCollection) (*hookdeckgosdk.UpdateCollectionResponse, error) {
	baseURL := "http://localhost:6333"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	endpointURL := fmt.Sprintf(baseURL+"/"+"collections/%v", collectionName)

	queryParams := make(url.Values)
	if request.Timeout != nil {
		queryParams.Add("timeout", fmt.Sprintf("%v", *request.Timeout))
	}
	if len(queryParams) > 0 {
		endpointURL += "?" + queryParams.Encode()
	}

	var response *hookdeckgosdk.UpdateCollectionResponse
	if err := c.caller.Call(
		ctx,
		&core.CallParams{
			URL:      endpointURL,
			Method:   http.MethodPatch,
			Headers:  c.header,
			Request:  request,
			Response: &response,
		},
	); err != nil {
		return nil, err
	}
	return response, nil
}

func (c *Client) UpdateAliases(ctx context.Context, request *hookdeckgosdk.ChangeAliasesOperation) (*hookdeckgosdk.UpdateAliasesResponse, error) {
	baseURL := "http://localhost:6333"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	endpointURL := baseURL + "/" + "collections/aliases"

	queryParams := make(url.Values)
	if request.Timeout != nil {
		queryParams.Add("timeout", fmt.Sprintf("%v", *request.Timeout))
	}
	if len(queryParams) > 0 {
		endpointURL += "?" + queryParams.Encode()
	}

	var response *hookdeckgosdk.UpdateAliasesResponse
	if err := c.caller.Call(
		ctx,
		&core.CallParams{
			URL:      endpointURL,
			Method:   http.MethodPost,
			Headers:  c.header,
			Request:  request,
			Response: &response,
		},
	); err != nil {
		return nil, err
	}
	return response, nil
}

// Create index for field in collection
//
// Name of the collection
func (c *Client) CreateFieldIndex(ctx context.Context, collectionName string, request *hookdeckgosdk.CreateFieldIndex) (*hookdeckgosdk.CreateFieldIndexResponse, error) {
	baseURL := "http://localhost:6333"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	endpointURL := fmt.Sprintf(baseURL+"/"+"collections/%v/index", collectionName)

	queryParams := make(url.Values)
	if request.Wait != nil {
		queryParams.Add("wait", fmt.Sprintf("%v", *request.Wait))
	}
	if request.Ordering != nil {
		queryParams.Add("ordering", fmt.Sprintf("%v", *request.Ordering))
	}
	if len(queryParams) > 0 {
		endpointURL += "?" + queryParams.Encode()
	}

	var response *hookdeckgosdk.CreateFieldIndexResponse
	if err := c.caller.Call(
		ctx,
		&core.CallParams{
			URL:      endpointURL,
			Method:   http.MethodPut,
			Headers:  c.header,
			Request:  request,
			Response: &response,
		},
	); err != nil {
		return nil, err
	}
	return response, nil
}

// Delete field index for collection
//
// Name of the collection
// Name of the field where to delete the index
func (c *Client) DeleteFieldIndex(ctx context.Context, collectionName string, fieldName string, request *hookdeckgosdk.DeleteFieldIndexRequest) (*hookdeckgosdk.DeleteFieldIndexResponse, error) {
	baseURL := "http://localhost:6333"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	endpointURL := fmt.Sprintf(baseURL+"/"+"collections/%v/index/%v", collectionName, fieldName)

	queryParams := make(url.Values)
	if request.Wait != nil {
		queryParams.Add("wait", fmt.Sprintf("%v", *request.Wait))
	}
	if request.Ordering != nil {
		queryParams.Add("ordering", fmt.Sprintf("%v", *request.Ordering))
	}
	if len(queryParams) > 0 {
		endpointURL += "?" + queryParams.Encode()
	}

	var response *hookdeckgosdk.DeleteFieldIndexResponse
	if err := c.caller.Call(
		ctx,
		&core.CallParams{
			URL:      endpointURL,
			Method:   http.MethodDelete,
			Headers:  c.header,
			Response: &response,
		},
	); err != nil {
		return nil, err
	}
	return response, nil
}

// Get cluster information for a collection
//
// Name of the collection to retrieve the cluster info for
func (c *Client) CollectionClusterInfo(ctx context.Context, collectionName string) (*hookdeckgosdk.CollectionClusterInfoResponse, error) {
	baseURL := "http://localhost:6333"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	endpointURL := fmt.Sprintf(baseURL+"/"+"collections/%v/cluster", collectionName)

	var response *hookdeckgosdk.CollectionClusterInfoResponse
	if err := c.caller.Call(
		ctx,
		&core.CallParams{
			URL:      endpointURL,
			Method:   http.MethodGet,
			Headers:  c.header,
			Response: &response,
		},
	); err != nil {
		return nil, err
	}
	return response, nil
}

// Name of the collection on which to to apply the cluster update operation
func (c *Client) UpdateCollectionCluster(ctx context.Context, collectionName string, request *hookdeckgosdk.UpdateCollectionClusterRequest) (*hookdeckgosdk.UpdateCollectionClusterResponse, error) {
	baseURL := "http://localhost:6333"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	endpointURL := fmt.Sprintf(baseURL+"/"+"collections/%v/cluster", collectionName)

	queryParams := make(url.Values)
	if request.Timeout != nil {
		queryParams.Add("timeout", fmt.Sprintf("%v", *request.Timeout))
	}
	if len(queryParams) > 0 {
		endpointURL += "?" + queryParams.Encode()
	}

	var response *hookdeckgosdk.UpdateCollectionClusterResponse
	if err := c.caller.Call(
		ctx,
		&core.CallParams{
			URL:      endpointURL,
			Method:   http.MethodPost,
			Headers:  c.header,
			Request:  request,
			Response: &response,
		},
	); err != nil {
		return nil, err
	}
	return response, nil
}

// Get list of all aliases for a collection
//
// Name of the collection
func (c *Client) GetCollectionAliases(ctx context.Context, collectionName string) (*hookdeckgosdk.GetCollectionAliasesResponse, error) {
	baseURL := "http://localhost:6333"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	endpointURL := fmt.Sprintf(baseURL+"/"+"collections/%v/aliases", collectionName)

	var response *hookdeckgosdk.GetCollectionAliasesResponse
	if err := c.caller.Call(
		ctx,
		&core.CallParams{
			URL:      endpointURL,
			Method:   http.MethodGet,
			Headers:  c.header,
			Response: &response,
		},
	); err != nil {
		return nil, err
	}
	return response, nil
}

// Get list of all existing collections aliases
func (c *Client) GetCollectionsAliases(ctx context.Context) (*hookdeckgosdk.GetCollectionsAliasesResponse, error) {
	baseURL := "http://localhost:6333"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	endpointURL := baseURL + "/" + "aliases"

	var response *hookdeckgosdk.GetCollectionsAliasesResponse
	if err := c.caller.Call(
		ctx,
		&core.CallParams{
			URL:      endpointURL,
			Method:   http.MethodGet,
			Headers:  c.header,
			Response: &response,
		},
	); err != nil {
		return nil, err
	}
	return response, nil
}
