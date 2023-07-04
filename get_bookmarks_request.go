// This file was auto-generated by Fern from our API Definition.

package api

type GetBookmarksRequest struct {
	Id          *string                                       `json:"-"`
	Name        *string                                       `json:"-"`
	WebhookId   *string                                       `json:"-"`
	EventDataId *string                                       `json:"-"`
	Label       *string                                       `json:"-"`
	LastUsedAt  *string                                       `json:"-"`
	OrderBy     *string                                       `json:"-"`
	Dir         *GetBookmarksRequestDirGetBookmarksRequestDir `json:"-"`
	Limit       *int                                          `json:"-"`
	Next        *string                                       `json:"-"`
	Prev        *string                                       `json:"-"`
}