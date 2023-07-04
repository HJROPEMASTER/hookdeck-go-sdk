// This file was auto-generated by Fern from our API Definition.

package api

import (
	json "encoding/json"
	fmt "fmt"
)

// Lifecyle status of the event
type CreateEventBulkRetryRequestQueryStatus struct {
	typeName        string
	EventStatus     EventStatus
	EventStatusList []EventStatus
}

func NewCreateEventBulkRetryRequestQueryStatusFromEventStatus(value EventStatus) *CreateEventBulkRetryRequestQueryStatus {
	return &CreateEventBulkRetryRequestQueryStatus{typeName: "eventStatus", EventStatus: value}
}

func NewCreateEventBulkRetryRequestQueryStatusFromEventStatusList(value []EventStatus) *CreateEventBulkRetryRequestQueryStatus {
	return &CreateEventBulkRetryRequestQueryStatus{typeName: "eventStatusList", EventStatusList: value}
}

func (c *CreateEventBulkRetryRequestQueryStatus) UnmarshalJSON(data []byte) error {
	var valueEventStatus EventStatus
	if err := json.Unmarshal(data, &valueEventStatus); err == nil {
		c.typeName = "eventStatus"
		c.EventStatus = valueEventStatus
		return nil
	}
	var valueEventStatusList []EventStatus
	if err := json.Unmarshal(data, &valueEventStatusList); err == nil {
		c.typeName = "eventStatusList"
		c.EventStatusList = valueEventStatusList
		return nil
	}
	return fmt.Errorf("%s cannot be deserialized as a %T", data, c)
}

func (c CreateEventBulkRetryRequestQueryStatus) MarshalJSON() ([]byte, error) {
	switch c.typeName {
	default:
		return nil, fmt.Errorf("invalid type %s in %T", c.typeName, c)
	case "eventStatus":
		return json.Marshal(c.EventStatus)
	case "eventStatusList":
		return json.Marshal(c.EventStatusList)
	}
}

type CreateEventBulkRetryRequestQueryStatusVisitor interface {
	VisitEventStatus(EventStatus) error
	VisitEventStatusList([]EventStatus) error
}

func (c *CreateEventBulkRetryRequestQueryStatus) Accept(v CreateEventBulkRetryRequestQueryStatusVisitor) error {
	switch c.typeName {
	default:
		return fmt.Errorf("invalid type %s in %T", c.typeName, c)
	case "eventStatus":
		return v.VisitEventStatus(c.EventStatus)
	case "eventStatusList":
		return v.VisitEventStatusList(c.EventStatusList)
	}
}