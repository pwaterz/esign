// Copyright 2017 James Cote and Liberty Fund, Inc.
// All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated by go-swagger; DO NOT EDIT.

// Package folders implements the DocuSign SDK
// category Folders.
// 
// Use the Folders category to manage envelopes in your folders. 
// 
// You can list the folder contents and move envelopes between folders.
// Api documentation may be found at:
// https://docs.docusign.com/esign/restapi/Folders
package folders

import (
    "fmt"
    "net/url"
    "time"
    
    "golang.org/x/net/context"
    
    "github.com/jfcote87/esign"
    "github.com/jfcote87/esign/model"
)

// Service generates DocuSign Folders Category API calls
type Service struct {
    credential esign.Credential 
}

// New initializes a folders service using cred to authorize calls.
func New(cred esign.Credential) *Service {
    return &Service{credential: cred}
}

// ListItems gets a list of the envelopes in the specified folder.
// SDK Method Folders::listItems
// https://docs.docusign.com/esign/restapi/Folders/Folders/listItems
func (s *Service) ListItems(folderID string) *ListItemsCall {
    return &ListItemsCall{
        &esign.Call{
            Credential: s.credential,
    		Method:  "GET",
            Path: "folders/{folderId}",
            PathParameters: map[string]string{ 
                "{folderId}": folderID,
            },
            QueryOpts: make(url.Values),
        },
    }
}

// ListItemsCall implements DocuSign API SDK Folders::listItems
type ListItemsCall struct {
    *esign.Call
}

// Do executes the call.  A nil context will return error.
func (op *ListItemsCall) Do(ctx context.Context)  (*model.FolderItemsResponse, error) {
    var res *model.FolderItemsResponse
    return res, op.Call.Do(ctx, &res)
}

// FromDate only return items on or after this date. If no value is provided, the default search is the previous 30 days.
func (op *ListItemsCall) FromDate(val time.Time) *ListItemsCall {
    op.QueryOpts.Set("from_date", val.Format(time.RFC3339))
    return op
}

// IncludeItems set the call query parameter include_items
func (op *ListItemsCall) IncludeItems(val string) *ListItemsCall {
    op.QueryOpts.Set("include_items", val)
    return op
}

// OwnerEmail is the email of the folder owner.
func (op *ListItemsCall) OwnerEmail(val string) *ListItemsCall {
    op.QueryOpts.Set("owner_email", val)
    return op
}

// OwnerName is the name of the folder owner.
func (op *ListItemsCall) OwnerName(val string) *ListItemsCall {
    op.QueryOpts.Set("owner_name", val)
    return op
}

// SearchText is the search text used to search the items of the envelope. The search looks at recipient names and emails, envelope custom fields, sender name, and subject.
func (op *ListItemsCall) SearchText(val string) *ListItemsCall {
    op.QueryOpts.Set("search_text", val)
    return op
}

// StartPosition is the position of the folder items to return. This is used for repeated calls, when the number of envelopes returned is too much for one return (calls return 100 envelopes at a time). The default value is 0.
func (op *ListItemsCall) StartPosition(val int) *ListItemsCall {
    op.QueryOpts.Set("start_position", fmt.Sprintf("%d", val ))
    return op
}

// Status is the current status of the envelope. If no value is provided, the default search is all/any status.
func (op *ListItemsCall) Status(val string) *ListItemsCall {
    op.QueryOpts.Set("status", val)
    return op
}

// ToDate only return items up to this date. If no value is provided, the default search is to the current date.
func (op *ListItemsCall) ToDate(val time.Time) *ListItemsCall {
    op.QueryOpts.Set("to_date", val.Format(time.RFC3339))
    return op
}

// List gets a list of the folders for the account.
// SDK Method Folders::list
// https://docs.docusign.com/esign/restapi/Folders/Folders/list
func (s *Service) List() *ListCall {
    return &ListCall{
        &esign.Call{
            Credential: s.credential,
    		Method:  "GET",
            Path: "folders",
            QueryOpts: make(url.Values),
        },
    }
}

// ListCall implements DocuSign API SDK Folders::list
type ListCall struct {
    *esign.Call
}

// Do executes the call.  A nil context will return error.
func (op *ListCall) Do(ctx context.Context)  (*model.FoldersResponse, error) {
    var res *model.FoldersResponse
    return res, op.Call.Do(ctx, &res)
}

// Include reserved for DocuSign.
func (op *ListCall) Include(val string) *ListCall {
    op.QueryOpts.Set("include", val)
    return op
}

// IncludeItems set the call query parameter include_items
func (op *ListCall) IncludeItems(val string) *ListCall {
    op.QueryOpts.Set("include_items", val)
    return op
}

// StartPosition reserved for DocuSign.
func (op *ListCall) StartPosition(val int) *ListCall {
    op.QueryOpts.Set("start_position", fmt.Sprintf("%d", val ))
    return op
}

// Template specifies the items that are returned. Valid values are: 
// 
// * include - The folder list will return normal folders plus template folders. 
// * only - Only the list of template folders are returned.
func (op *ListCall) Template(val string) *ListCall {
    op.QueryOpts.Set("template", val)
    return op
}

// UserFilter reserved for DocuSign.
func (op *ListCall) UserFilter(val string) *ListCall {
    op.QueryOpts.Set("user_filter", val)
    return op
}

// MoveEnvelopes moves an envelope from its current folder to the specified folder.
// SDK Method Folders::moveEnvelopes
// https://docs.docusign.com/esign/restapi/Folders/Folders/moveEnvelopes
func (s *Service) MoveEnvelopes(folderID string, foldersRequest *model.FoldersRequest) *MoveEnvelopesCall {
    return &MoveEnvelopesCall{
        &esign.Call{
            Credential: s.credential,
    		Method:  "PUT",
            Path: "folders/{folderId}",
            PathParameters: map[string]string{ 
                "{folderId}": folderID,
            },
            Payload: foldersRequest,
            QueryOpts: make(url.Values),
        },
    }
}

// MoveEnvelopesCall implements DocuSign API SDK Folders::moveEnvelopes
type MoveEnvelopesCall struct {
    *esign.Call
}

// Do executes the call.  A nil context will return error.
func (op *MoveEnvelopesCall) Do(ctx context.Context)  (*model.FoldersResponse, error) {
    var res *model.FoldersResponse
    return res, op.Call.Do(ctx, &res)
}

// Search gets a list of envelopes in folders matching the specified criteria.
// SDK Method Folders::search
// https://docs.docusign.com/esign/restapi/Folders/Folders/search
func (s *Service) Search(searchFolderID string) *SearchCall {
    return &SearchCall{
        &esign.Call{
            Credential: s.credential,
    		Method:  "GET",
            Path: "search_folders/{searchFolderId}",
            PathParameters: map[string]string{ 
                "{searchFolderId}": searchFolderID,
            },
            QueryOpts: make(url.Values),
        },
    }
}

// SearchCall implements DocuSign API SDK Folders::search
type SearchCall struct {
    *esign.Call
}

// Do executes the call.  A nil context will return error.
func (op *SearchCall) Do(ctx context.Context)  error {
    
    return op.Call.Do(ctx, nil)
}

// All specifies that all envelopes that match the criteria are returned.
func (op *SearchCall) All() *SearchCall {
    op.QueryOpts.Set("all", "true")
    return op
}

// Count specifies the number of records returned in the cache. The number must be greater than 0 and less than or equal to 100.
func (op *SearchCall) Count(val int) *SearchCall {
    op.QueryOpts.Set("count", fmt.Sprintf("%d", val ))
    return op
}

// FromDate specifies the start of the date range to return. If no value is provided, the default search is the previous 30 days.
func (op *SearchCall) FromDate(val time.Time) *SearchCall {
    op.QueryOpts.Set("from_date", val.Format(time.RFC3339))
    return op
}

// IncludeRecipients when set to **true**, the recipient information is returned in the response.
func (op *SearchCall) IncludeRecipients() *SearchCall {
    op.QueryOpts.Set("include_recipients", "true")
    return op
}

// Order specifies the order in which the list is returned. Valid values are: `asc` for ascending order, and `desc` for descending order.
func (op *SearchCall) Order(val string) *SearchCall {
    op.QueryOpts.Set("order", val)
    return op
}

// OrderBy specifies the property used to sort the list. Valid values are: `action_required`, `created`, `completed`, `sent`, `signer_list`, `status`, or `subject`.
func (op *SearchCall) OrderBy(val string) *SearchCall {
    op.QueryOpts.Set("order_by", val)
    return op
}

// StartPosition specifies the the starting location in the result set of the items that are returned.
func (op *SearchCall) StartPosition(val int) *SearchCall {
    op.QueryOpts.Set("start_position", fmt.Sprintf("%d", val ))
    return op
}

// ToDate specifies the end of the date range to return.
func (op *SearchCall) ToDate(val time.Time) *SearchCall {
    op.QueryOpts.Set("to_date", val.Format(time.RFC3339))
    return op
}

