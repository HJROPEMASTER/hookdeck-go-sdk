// This file was auto-generated by Fern from our API Definition.

package api

import (
	json "encoding/json"
	fmt "fmt"
)

type RequestDataParsedQuery struct {
	typeName                                     string
	StringOptional                               *string
	RequestDataParsedQueryRequestDataParsedQuery *RequestDataParsedQueryRequestDataParsedQuery
}

func NewRequestDataParsedQueryFromStringOptional(value *string) *RequestDataParsedQuery {
	return &RequestDataParsedQuery{typeName: "stringOptional", StringOptional: value}
}

func NewRequestDataParsedQueryFromRequestDataParsedQueryRequestDataParsedQuery(value *RequestDataParsedQueryRequestDataParsedQuery) *RequestDataParsedQuery {
	return &RequestDataParsedQuery{typeName: "requestDataParsedQueryRequestDataParsedQuery", RequestDataParsedQueryRequestDataParsedQuery: value}
}

func (r *RequestDataParsedQuery) UnmarshalJSON(data []byte) error {
	var valueStringOptional *string
	if err := json.Unmarshal(data, &valueStringOptional); err == nil {
		r.typeName = "stringOptional"
		r.StringOptional = valueStringOptional
		return nil
	}
	valueRequestDataParsedQueryRequestDataParsedQuery := new(RequestDataParsedQueryRequestDataParsedQuery)
	if err := json.Unmarshal(data, &valueRequestDataParsedQueryRequestDataParsedQuery); err == nil {
		r.typeName = "requestDataParsedQueryRequestDataParsedQuery"
		r.RequestDataParsedQueryRequestDataParsedQuery = valueRequestDataParsedQueryRequestDataParsedQuery
		return nil
	}
	return fmt.Errorf("%s cannot be deserialized as a %T", data, r)
}

func (r RequestDataParsedQuery) MarshalJSON() ([]byte, error) {
	switch r.typeName {
	default:
		return nil, fmt.Errorf("invalid type %s in %T", r.typeName, r)
	case "stringOptional":
		return json.Marshal(r.StringOptional)
	case "requestDataParsedQueryRequestDataParsedQuery":
		return json.Marshal(r.RequestDataParsedQueryRequestDataParsedQuery)
	}
}

type RequestDataParsedQueryVisitor interface {
	VisitStringOptional(*string) error
	VisitRequestDataParsedQueryRequestDataParsedQuery(*RequestDataParsedQueryRequestDataParsedQuery) error
}

func (r *RequestDataParsedQuery) Accept(v RequestDataParsedQueryVisitor) error {
	switch r.typeName {
	default:
		return fmt.Errorf("invalid type %s in %T", r.typeName, r)
	case "stringOptional":
		return v.VisitStringOptional(r.StringOptional)
	case "requestDataParsedQueryRequestDataParsedQuery":
		return v.VisitRequestDataParsedQueryRequestDataParsedQuery(r.RequestDataParsedQueryRequestDataParsedQuery)
	}
}