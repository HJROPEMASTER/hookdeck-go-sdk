// This file was auto-generated by Fern from our API Definition.

package api

import (
	json "encoding/json"
	fmt "fmt"
)

// Filter by requests IDs
type CreateRequestBulkRetryRequestQueryId struct {
	typeName string
	// Request ID <span style="white-space: nowrap">`<= 255 characters`</span>
	String     string
	StringList []string
}

func NewCreateRequestBulkRetryRequestQueryIdFromString(value string) *CreateRequestBulkRetryRequestQueryId {
	return &CreateRequestBulkRetryRequestQueryId{typeName: "string", String: value}
}

func NewCreateRequestBulkRetryRequestQueryIdFromStringList(value []string) *CreateRequestBulkRetryRequestQueryId {
	return &CreateRequestBulkRetryRequestQueryId{typeName: "stringList", StringList: value}
}

func (c *CreateRequestBulkRetryRequestQueryId) UnmarshalJSON(data []byte) error {
	var valueString string
	if err := json.Unmarshal(data, &valueString); err == nil {
		c.typeName = "string"
		c.String = valueString
		return nil
	}
	var valueStringList []string
	if err := json.Unmarshal(data, &valueStringList); err == nil {
		c.typeName = "stringList"
		c.StringList = valueStringList
		return nil
	}
	return fmt.Errorf("%s cannot be deserialized as a %T", data, c)
}

func (c CreateRequestBulkRetryRequestQueryId) MarshalJSON() ([]byte, error) {
	switch c.typeName {
	default:
		return nil, fmt.Errorf("invalid type %s in %T", c.typeName, c)
	case "string":
		return json.Marshal(c.String)
	case "stringList":
		return json.Marshal(c.StringList)
	}
}

type CreateRequestBulkRetryRequestQueryIdVisitor interface {
	VisitString(string) error
	VisitStringList([]string) error
}

func (c *CreateRequestBulkRetryRequestQueryId) Accept(v CreateRequestBulkRetryRequestQueryIdVisitor) error {
	switch c.typeName {
	default:
		return fmt.Errorf("invalid type %s in %T", c.typeName, c)
	case "string":
		return v.VisitString(c.String)
	case "stringList":
		return v.VisitStringList(c.StringList)
	}
}