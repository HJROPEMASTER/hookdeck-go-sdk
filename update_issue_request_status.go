// This file was auto-generated by Fern from our API Definition.

package api

import (
	json "encoding/json"
	fmt "fmt"
	strconv "strconv"
)

// New status
type UpdateIssueRequestStatus uint8

const (
	UpdateIssueRequestStatusOpened UpdateIssueRequestStatus = iota + 1
	UpdateIssueRequestStatusIgnored
	UpdateIssueRequestStatusAcknowledged
	UpdateIssueRequestStatusResolved
)

func (u UpdateIssueRequestStatus) String() string {
	switch u {
	default:
		return strconv.Itoa(int(u))
	case UpdateIssueRequestStatusOpened:
		return "OPENED"
	case UpdateIssueRequestStatusIgnored:
		return "IGNORED"
	case UpdateIssueRequestStatusAcknowledged:
		return "ACKNOWLEDGED"
	case UpdateIssueRequestStatusResolved:
		return "RESOLVED"
	}
}

func (u UpdateIssueRequestStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%q", u.String())), nil
}

func (u *UpdateIssueRequestStatus) UnmarshalJSON(data []byte) error {
	var raw string
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	switch raw {
	case "OPENED":
		value := UpdateIssueRequestStatusOpened
		*u = value
	case "IGNORED":
		value := UpdateIssueRequestStatusIgnored
		*u = value
	case "ACKNOWLEDGED":
		value := UpdateIssueRequestStatusAcknowledged
		*u = value
	case "RESOLVED":
		value := UpdateIssueRequestStatusResolved
		*u = value
	}
	return nil
}