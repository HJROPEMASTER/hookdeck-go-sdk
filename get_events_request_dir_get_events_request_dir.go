// This file was auto-generated by Fern from our API Definition.

package api

import (
	json "encoding/json"
	fmt "fmt"
	strconv "strconv"
)

type GetEventsRequestDirGetEventsRequestDir uint8

const (
	GetEventsRequestDirGetEventsRequestDirAsc GetEventsRequestDirGetEventsRequestDir = iota + 1
	GetEventsRequestDirGetEventsRequestDirDesc
)

func (g GetEventsRequestDirGetEventsRequestDir) String() string {
	switch g {
	default:
		return strconv.Itoa(int(g))
	case GetEventsRequestDirGetEventsRequestDirAsc:
		return "asc"
	case GetEventsRequestDirGetEventsRequestDirDesc:
		return "desc"
	}
}

func (g GetEventsRequestDirGetEventsRequestDir) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%q", g.String())), nil
}

func (g *GetEventsRequestDirGetEventsRequestDir) UnmarshalJSON(data []byte) error {
	var raw string
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	switch raw {
	case "asc":
		value := GetEventsRequestDirGetEventsRequestDirAsc
		*g = value
	case "desc":
		value := GetEventsRequestDirGetEventsRequestDirDesc
		*g = value
	}
	return nil
}