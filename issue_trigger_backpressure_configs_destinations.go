// This file was auto-generated by Fern from our API Definition.

package api

import (
	json "encoding/json"
	fmt "fmt"
)

// A pattern to match on the destination name or array of destination IDs. Use `*` as wildcard.
type IssueTriggerBackpressureConfigsDestinations struct {
	typeName   string
	String     string
	StringList []string
}

func NewIssueTriggerBackpressureConfigsDestinationsFromString(value string) *IssueTriggerBackpressureConfigsDestinations {
	return &IssueTriggerBackpressureConfigsDestinations{typeName: "string", String: value}
}

func NewIssueTriggerBackpressureConfigsDestinationsFromStringList(value []string) *IssueTriggerBackpressureConfigsDestinations {
	return &IssueTriggerBackpressureConfigsDestinations{typeName: "stringList", StringList: value}
}

func (i *IssueTriggerBackpressureConfigsDestinations) UnmarshalJSON(data []byte) error {
	var valueString string
	if err := json.Unmarshal(data, &valueString); err == nil {
		i.typeName = "string"
		i.String = valueString
		return nil
	}
	var valueStringList []string
	if err := json.Unmarshal(data, &valueStringList); err == nil {
		i.typeName = "stringList"
		i.StringList = valueStringList
		return nil
	}
	return fmt.Errorf("%s cannot be deserialized as a %T", data, i)
}

func (i IssueTriggerBackpressureConfigsDestinations) MarshalJSON() ([]byte, error) {
	switch i.typeName {
	default:
		return nil, fmt.Errorf("invalid type %s in %T", i.typeName, i)
	case "string":
		return json.Marshal(i.String)
	case "stringList":
		return json.Marshal(i.StringList)
	}
}

type IssueTriggerBackpressureConfigsDestinationsVisitor interface {
	VisitString(string) error
	VisitStringList([]string) error
}

func (i *IssueTriggerBackpressureConfigsDestinations) Accept(v IssueTriggerBackpressureConfigsDestinationsVisitor) error {
	switch i.typeName {
	default:
		return fmt.Errorf("invalid type %s in %T", i.typeName, i)
	case "string":
		return v.VisitString(i.String)
	case "stringList":
		return v.VisitStringList(i.StringList)
	}
}