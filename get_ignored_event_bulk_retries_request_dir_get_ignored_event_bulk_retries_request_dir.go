// This file was auto-generated by Fern from our API Definition.

package api

import (
	json "encoding/json"
	fmt "fmt"
	strconv "strconv"
)

type GetIgnoredEventBulkRetriesRequestDirGetIgnoredEventBulkRetriesRequestDir uint8

const (
	GetIgnoredEventBulkRetriesRequestDirGetIgnoredEventBulkRetriesRequestDirAsc GetIgnoredEventBulkRetriesRequestDirGetIgnoredEventBulkRetriesRequestDir = iota + 1
	GetIgnoredEventBulkRetriesRequestDirGetIgnoredEventBulkRetriesRequestDirDesc
)

func (g GetIgnoredEventBulkRetriesRequestDirGetIgnoredEventBulkRetriesRequestDir) String() string {
	switch g {
	default:
		return strconv.Itoa(int(g))
	case GetIgnoredEventBulkRetriesRequestDirGetIgnoredEventBulkRetriesRequestDirAsc:
		return "asc"
	case GetIgnoredEventBulkRetriesRequestDirGetIgnoredEventBulkRetriesRequestDirDesc:
		return "desc"
	}
}

func (g GetIgnoredEventBulkRetriesRequestDirGetIgnoredEventBulkRetriesRequestDir) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%q", g.String())), nil
}

func (g *GetIgnoredEventBulkRetriesRequestDirGetIgnoredEventBulkRetriesRequestDir) UnmarshalJSON(data []byte) error {
	var raw string
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	switch raw {
	case "asc":
		value := GetIgnoredEventBulkRetriesRequestDirGetIgnoredEventBulkRetriesRequestDirAsc
		*g = value
	case "desc":
		value := GetIgnoredEventBulkRetriesRequestDirGetIgnoredEventBulkRetriesRequestDirDesc
		*g = value
	}
	return nil
}