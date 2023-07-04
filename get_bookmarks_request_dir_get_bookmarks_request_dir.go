// This file was auto-generated by Fern from our API Definition.

package api

import (
	json "encoding/json"
	fmt "fmt"
	strconv "strconv"
)

type GetBookmarksRequestDirGetBookmarksRequestDir uint8

const (
	GetBookmarksRequestDirGetBookmarksRequestDirAsc GetBookmarksRequestDirGetBookmarksRequestDir = iota + 1
	GetBookmarksRequestDirGetBookmarksRequestDirDesc
)

func (g GetBookmarksRequestDirGetBookmarksRequestDir) String() string {
	switch g {
	default:
		return strconv.Itoa(int(g))
	case GetBookmarksRequestDirGetBookmarksRequestDirAsc:
		return "asc"
	case GetBookmarksRequestDirGetBookmarksRequestDirDesc:
		return "desc"
	}
}

func (g GetBookmarksRequestDirGetBookmarksRequestDir) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%q", g.String())), nil
}

func (g *GetBookmarksRequestDirGetBookmarksRequestDir) UnmarshalJSON(data []byte) error {
	var raw string
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	switch raw {
	case "asc":
		value := GetBookmarksRequestDirGetBookmarksRequestDirAsc
		*g = value
	case "desc":
		value := GetBookmarksRequestDirGetBookmarksRequestDirDesc
		*g = value
	}
	return nil
}