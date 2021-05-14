// Copyright 2019 James Cote
// All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated by gen-esign; DO NOT EDIT.

// Package diagnostics implements the DocuSign SDK
// category Diagnostics.
//
// The Diagnostics category provides miscellaneous end points.
//
// They include:
// * Requesting and managing the API call-logging feature. (Perfect for debugging your app!)
// * Getting information on the API's resources and versions.
//
// Service Api documentation may be found at:
// https://developers.docusign.com/esign-rest-api/reference/Diagnostics
// Usage example:
//
//   import (
//       "github.com/pwaterz/esign"
//       "github.com/pwaterz/esign/v2.1/diagnostics"
//       "github.com/pwaterz/esign/v2.1/model"
//   )
//   ...
//   diagnosticsService := diagnostics.New(esignCredential)
package diagnostics // import "github.com/pwaterz/esign/v2.1/diagnostics"

import (
	"context"
	"net/url"
	"strings"

	"github.com/pwaterz/esign"
	"github.com/pwaterz/esign/v2.1/model"
)

// Service implements DocuSign Diagnostics Category API operations
type Service struct {
	credential esign.Credential
}

// New initializes a diagnostics service using cred to authorize ops.
func New(cred esign.Credential) *Service {
	return &Service{credential: cred}
}

// RequestLogsDelete deletes the request log files.
//
// https://developers.docusign.com/esign-rest-api/reference/diagnostics/requestlogs/delete
//
// SDK Method Diagnostics::deleteRequestLogs
func (s *Service) RequestLogsDelete() *RequestLogsDeleteOp {
	return &RequestLogsDeleteOp{
		Credential: s.credential,
		Method:     "DELETE",
		Path:       "/v2.1/diagnostics/request_logs",
		QueryOpts:  make(url.Values),
		Version:    esign.VersionV21,
	}
}

// RequestLogsDeleteOp implements DocuSign API SDK Diagnostics::deleteRequestLogs
type RequestLogsDeleteOp esign.Op

// Do executes the op.  A nil context will return error.
func (op *RequestLogsDeleteOp) Do(ctx context.Context) error {
	return ((*esign.Op)(op)).Do(ctx, nil)
}

// RequestLogsGet gets a request logging log file.
//
// https://developers.docusign.com/esign-rest-api/reference/diagnostics/requestlogs/get
//
// SDK Method Diagnostics::getRequestLog
func (s *Service) RequestLogsGet(requestLogID string) *RequestLogsGetOp {
	return &RequestLogsGetOp{
		Credential: s.credential,
		Method:     "GET",
		Path:       strings.Join([]string{"", "v2.1", "diagnostics", "request_logs", requestLogID}, "/"),
		Accept:     "text/plain",
		QueryOpts:  make(url.Values),
		Version:    esign.VersionV21,
	}
}

// RequestLogsGetOp implements DocuSign API SDK Diagnostics::getRequestLog
type RequestLogsGetOp esign.Op

// Do executes the op.  A nil context will return error.
func (op *RequestLogsGetOp) Do(ctx context.Context) (*esign.Download, error) {
	var res *esign.Download
	return res, ((*esign.Op)(op)).Do(ctx, &res)
}

// RequestLogsGetSettings gets the API request logging settings.
//
// https://developers.docusign.com/esign-rest-api/reference/diagnostics/requestlogs/getsettings
//
// SDK Method Diagnostics::getRequestLogSettings
func (s *Service) RequestLogsGetSettings() *RequestLogsGetSettingsOp {
	return &RequestLogsGetSettingsOp{
		Credential: s.credential,
		Method:     "GET",
		Path:       "/v2.1/diagnostics/settings",
		QueryOpts:  make(url.Values),
		Version:    esign.VersionV21,
	}
}

// RequestLogsGetSettingsOp implements DocuSign API SDK Diagnostics::getRequestLogSettings
type RequestLogsGetSettingsOp esign.Op

// Do executes the op.  A nil context will return error.
func (op *RequestLogsGetSettingsOp) Do(ctx context.Context) (*model.DiagnosticsSettingsInformation, error) {
	var res *model.DiagnosticsSettingsInformation
	return res, ((*esign.Op)(op)).Do(ctx, &res)
}

// RequestLogsList gets the API request logging log files.
//
// https://developers.docusign.com/esign-rest-api/reference/diagnostics/requestlogs/list
//
// SDK Method Diagnostics::listRequestLogs
func (s *Service) RequestLogsList() *RequestLogsListOp {
	return &RequestLogsListOp{
		Credential: s.credential,
		Method:     "GET",
		Path:       "/v2.1/diagnostics/request_logs",
		QueryOpts:  make(url.Values),
		Version:    esign.VersionV21,
	}
}

// RequestLogsListOp implements DocuSign API SDK Diagnostics::listRequestLogs
type RequestLogsListOp esign.Op

// Do executes the op.  A nil context will return error.
func (op *RequestLogsListOp) Do(ctx context.Context) (*model.APIRequestLogsResult, error) {
	var res *model.APIRequestLogsResult
	return res, ((*esign.Op)(op)).Do(ctx, &res)
}

// Encoding set the call query parameter encoding
func (op *RequestLogsListOp) Encoding(val string) *RequestLogsListOp {
	if op != nil {
		op.QueryOpts.Set("encoding", val)
	}
	return op
}

// Zip returns a zip file containing log files by setting
// the Accept header to application/zip
//
// **not included in swagger definition
func (op *RequestLogsListOp) Zip(ctx context.Context) (*esign.Download, error) {
	var res *esign.Download
	if op == nil {
		return nil, esign.ErrNilOp
	}
	newOp := esign.Op(*op)
	newOp.Accept = "application/zip"
	return res, (&newOp).Do(ctx, &res)
}

// RequestLogsUpdateSettings enables or disables API request logging for troubleshooting.
//
// https://developers.docusign.com/esign-rest-api/reference/diagnostics/requestlogs/updatesettings
//
// SDK Method Diagnostics::updateRequestLogSettings
func (s *Service) RequestLogsUpdateSettings(requestLogs *model.DiagnosticsSettingsInformation) *RequestLogsUpdateSettingsOp {
	return &RequestLogsUpdateSettingsOp{
		Credential: s.credential,
		Method:     "PUT",
		Path:       "/v2.1/diagnostics/settings",
		Payload:    requestLogs,
		QueryOpts:  make(url.Values),
		Version:    esign.VersionV21,
	}
}

// RequestLogsUpdateSettingsOp implements DocuSign API SDK Diagnostics::updateRequestLogSettings
type RequestLogsUpdateSettingsOp esign.Op

// Do executes the op.  A nil context will return error.
func (op *RequestLogsUpdateSettingsOp) Do(ctx context.Context) (*model.DiagnosticsSettingsInformation, error) {
	var res *model.DiagnosticsSettingsInformation
	return res, ((*esign.Op)(op)).Do(ctx, &res)
}

// ResourcesGet lists resources for REST version specified
//
// https://developers.docusign.com/esign-rest-api/reference/diagnostics/resources/get
//
// SDK Method Diagnostics::getResources
func (s *Service) ResourcesGet() *ResourcesGetOp {
	return &ResourcesGetOp{
		Credential: s.credential,
		Method:     "GET",
		Path:       "/v2.1",
		QueryOpts:  make(url.Values),
		Version:    esign.VersionV21,
	}
}

// ResourcesGetOp implements DocuSign API SDK Diagnostics::getResources
type ResourcesGetOp esign.Op

// Do executes the op.  A nil context will return error.
func (op *ResourcesGetOp) Do(ctx context.Context) (*model.ResourceInformation, error) {
	var res *model.ResourceInformation
	return res, ((*esign.Op)(op)).Do(ctx, &res)
}

// ServicesGet retrieves the available REST API versions.
//
// https://developers.docusign.com/esign-rest-api/reference/diagnostics/services/get
//
// SDK Method Diagnostics::getService
func (s *Service) ServicesGet() *ServicesGetOp {
	return &ServicesGetOp{
		Credential: s.credential,
		Method:     "GET",
		Path:       "/service_information",
		QueryOpts:  make(url.Values),
		Version:    esign.VersionV21,
	}
}

// ServicesGetOp implements DocuSign API SDK Diagnostics::getService
type ServicesGetOp esign.Op

// Do executes the op.  A nil context will return error.
func (op *ServicesGetOp) Do(ctx context.Context) (*model.ServiceInformation, error) {
	var res *model.ServiceInformation
	return res, ((*esign.Op)(op)).Do(ctx, &res)
}
