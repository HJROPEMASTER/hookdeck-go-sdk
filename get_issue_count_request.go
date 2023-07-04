// This file was auto-generated by Fern from our API Definition.

package api

type GetIssueCountRequest struct {
	Id             *string                                                 `json:"-"`
	IssueTriggerId *string                                                 `json:"-"`
	Type           *GetIssueCountRequestTypeGetIssueCountRequestType       `json:"-"`
	Status         *GetIssueCountRequestStatusGetIssueCountRequestStatus   `json:"-"`
	MergedWith     *string                                                 `json:"-"`
	CreatedAt      *string                                                 `json:"-"`
	FirstSeenAt    *string                                                 `json:"-"`
	LastSeenAt     *string                                                 `json:"-"`
	DismissedAt    *string                                                 `json:"-"`
	OrderBy        *GetIssueCountRequestOrderByGetIssueCountRequestOrderBy `json:"-"`
	Dir            *GetIssueCountRequestDirGetIssueCountRequestDir         `json:"-"`
	Limit          *int                                                    `json:"-"`
	Next           *string                                                 `json:"-"`
	Prev           *string                                                 `json:"-"`
}