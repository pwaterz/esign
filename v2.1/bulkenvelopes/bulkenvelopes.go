// Copyright 2019 James Cote
// All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated by gen-esign; DO NOT EDIT.

// Package bulkenvelopes implements the DocuSign SDK
// category BulkEnvelopes.
//
// Use the BulkEnvelopes category to manage the sending of envelopes to multiple recipients.
//
// Service Api documentation may be found at:
// https://developers.docusign.com/esign-rest-api/reference/BulkEnvelopes
// Usage example:
//
//   import (
//       "github.com/pwaterz/esign"
//       "github.com/pwaterz/esign/v2.1/bulkenvelopes"
//       "github.com/pwaterz/esign/v2.1/model"
//   )
//   ...
//   bulkenvelopesService := bulkenvelopes.New(esignCredential)
package bulkenvelopes // import "github.com/pwaterz/esign/v2.1/bulkenvelopes"

import (
	"context"
	"fmt"
	"net/url"
	"strings"

	"github.com/pwaterz/esign"
	"github.com/pwaterz/esign/v2.1/model"
)

// Service implements DocuSign BulkEnvelopes Category API operations
type Service struct {
	credential esign.Credential
}

// New initializes a bulkenvelopes service using cred to authorize ops.
func New(cred esign.Credential) *Service {
	return &Service{credential: cred}
}

// Get gets the status of a specified bulk send operation.
//
// https://developers.docusign.com/esign-rest-api/reference/bulkenvelopes/bulkenvelopes/get
//
// SDK Method BulkEnvelopes::get
func (s *Service) Get(batchID string) *GetOp {
	return &GetOp{
		Credential: s.credential,
		Method:     "GET",
		Path:       strings.Join([]string{"bulk_envelopes", batchID}, "/"),
		QueryOpts:  make(url.Values),
		Version:    esign.VersionV21,
	}
}

// GetOp implements DocuSign API SDK BulkEnvelopes::get
type GetOp esign.Op

// Do executes the op.  A nil context will return error.
func (op *GetOp) Do(ctx context.Context) (*model.BulkEnvelopeStatus, error) {
	var res *model.BulkEnvelopeStatus
	return res, ((*esign.Op)(op)).Do(ctx, &res)
}

// Count is the maximum number of results to return.
func (op *GetOp) Count(val int) *GetOp {
	if op != nil {
		op.QueryOpts.Set("count", fmt.Sprintf("%d", val))
	}
	return op
}

// Include specifies which entries to include in the response. You can include multiple entries by using commas in the query string. For example: `?include="failed,queued"`.
//
// Valid query values:
// * `all` - Returns all entries. This overrides all other query settings, and is the default if no query string is provided.
// * `failed` - Return entries with a failed status.
// * `processing` - Return entries with a processing status.
// * `queued` - Return entries with a queued status.
// * `sent` - Return entries with a sent status.
func (op *GetOp) Include(val ...string) *GetOp {
	if op != nil {
		op.QueryOpts.Set("include", strings.Join(val, ","))
	}
	return op
}

// StartPosition specifies the position in the list of envelopes from which to start returning bulk recipient batch status information.
func (op *GetOp) StartPosition(val int) *GetOp {
	if op != nil {
		op.QueryOpts.Set("start_position", fmt.Sprintf("%d", val))
	}
	return op
}

// List gets status information about bulk recipient batches.
//
// https://developers.docusign.com/esign-rest-api/reference/bulkenvelopes/bulkenvelopes/list
//
// SDK Method BulkEnvelopes::list
func (s *Service) List() *ListOp {
	return &ListOp{
		Credential: s.credential,
		Method:     "GET",
		Path:       "bulk_envelopes",
		QueryOpts:  make(url.Values),
		Version:    esign.VersionV21,
	}
}

// ListOp implements DocuSign API SDK BulkEnvelopes::list
type ListOp esign.Op

// Do executes the op.  A nil context will return error.
func (op *ListOp) Do(ctx context.Context) (*model.BulkEnvelopesResponse, error) {
	var res *model.BulkEnvelopesResponse
	return res, ((*esign.Op)(op)).Do(ctx, &res)
}

// Count is the number of results to return. Can be 1 to 20.
func (op *ListOp) Count(val int) *ListOp {
	if op != nil {
		op.QueryOpts.Set("count", fmt.Sprintf("%d", val))
	}
	return op
}

// Include specifies which entries to include in the response. You can include multiple entries by using commas in the query string. For example: `?include="failed,queued"`.
//
// Valid query values:
// * `all` - Returns all entries. This overrides all other query settings, and is the default if no query string is provided.
// * `failed` - Return entries with a failed status.
// * `processing` - Return entries with a processing status.
// * `queued` - Return entries with a queued status.
// * `sent` - Return entries with a sent status.
func (op *ListOp) Include(val ...string) *ListOp {
	if op != nil {
		op.QueryOpts.Set("include", strings.Join(val, ","))
	}
	return op
}

// StartPosition specifies the position of the bulk envelope items from which to start returning results. You can use this parameter for repeated calls, when the number of bulk envelopes returned is too large to be included in a single response. The default value is 0.
func (op *ListOp) StartPosition(val int) *ListOp {
	if op != nil {
		op.QueryOpts.Set("start_position", fmt.Sprintf("%d", val))
	}
	return op
}

// RecipientsDelete deletes the bulk recipient file from an envelope.
//
// https://developers.docusign.com/esign-rest-api/reference/bulkenvelopes/envelopebulkrecipients/delete
//
// SDK Method BulkEnvelopes::deleteRecipients
func (s *Service) RecipientsDelete(envelopeID string, recipientID string) *RecipientsDeleteOp {
	return &RecipientsDeleteOp{
		Credential: s.credential,
		Method:     "DELETE",
		Path:       strings.Join([]string{"envelopes", envelopeID, "recipients", recipientID, "bulk_recipients"}, "/"),
		QueryOpts:  make(url.Values),
		Version:    esign.VersionV21,
	}
}

// RecipientsDeleteOp implements DocuSign API SDK BulkEnvelopes::deleteRecipients
type RecipientsDeleteOp esign.Op

// Do executes the op.  A nil context will return error.
func (op *RecipientsDeleteOp) Do(ctx context.Context) (*model.BulkRecipientsUpdateResponse, error) {
	var res *model.BulkRecipientsUpdateResponse
	return res, ((*esign.Op)(op)).Do(ctx, &res)
}

// RecipientsList gets the bulk recipient file from an envelope.
//
// https://developers.docusign.com/esign-rest-api/reference/bulkenvelopes/envelopebulkrecipients/list
//
// SDK Method BulkEnvelopes::getRecipients
func (s *Service) RecipientsList(envelopeID string, recipientID string) *RecipientsListOp {
	return &RecipientsListOp{
		Credential: s.credential,
		Method:     "GET",
		Path:       strings.Join([]string{"envelopes", envelopeID, "recipients", recipientID, "bulk_recipients"}, "/"),
		QueryOpts:  make(url.Values),
		Version:    esign.VersionV21,
	}
}

// RecipientsListOp implements DocuSign API SDK BulkEnvelopes::getRecipients
type RecipientsListOp esign.Op

// Do executes the op.  A nil context will return error.
func (op *RecipientsListOp) Do(ctx context.Context) (*model.BulkRecipientsResponse, error) {
	var res *model.BulkRecipientsResponse
	return res, ((*esign.Op)(op)).Do(ctx, &res)
}

// IncludeTabs when set to **true**, the response includes the tabs associated with the recipients.
func (op *RecipientsListOp) IncludeTabs() *RecipientsListOp {
	if op != nil {
		op.QueryOpts.Set("include_tabs", "true")
	}
	return op
}

// StartPosition is the starting position in the results set.
func (op *RecipientsListOp) StartPosition(val int) *RecipientsListOp {
	if op != nil {
		op.QueryOpts.Set("start_position", fmt.Sprintf("%d", val))
	}
	return op
}

// RecipientsUpdate adds or replaces envelope bulk recipients.
//
// https://developers.docusign.com/esign-rest-api/reference/bulkenvelopes/envelopebulkrecipients/update
//
// SDK Method BulkEnvelopes::updateRecipients
func (s *Service) RecipientsUpdate(envelopeID string, recipientID string, bulkRecipientsRequest *model.BulkRecipientsRequest) *RecipientsUpdateOp {
	return &RecipientsUpdateOp{
		Credential: s.credential,
		Method:     "PUT",
		Path:       strings.Join([]string{"envelopes", envelopeID, "recipients", recipientID, "bulk_recipients"}, "/"),
		Payload:    bulkRecipientsRequest,
		QueryOpts:  make(url.Values),
		Version:    esign.VersionV21,
	}
}

// RecipientsUpdateOp implements DocuSign API SDK BulkEnvelopes::updateRecipients
type RecipientsUpdateOp esign.Op

// Do executes the op.  A nil context will return error.
func (op *RecipientsUpdateOp) Do(ctx context.Context) (*model.BulkRecipientsSummaryResponse, error) {
	var res *model.BulkRecipientsSummaryResponse
	return res, ((*esign.Op)(op)).Do(ctx, &res)
}
