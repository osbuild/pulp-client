/*
Pulp 3 API

Fetch, Upload, Organize, and Distribute Software Packages

API version: v3
Contact: pulp-list@redhat.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package pulpclient

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
	"reflect"
)


// ContentInstallerFileIndicesAPIService ContentInstallerFileIndicesAPI service
type ContentInstallerFileIndicesAPIService service

type ContentInstallerFileIndicesAPIContentDebInstallerFileIndicesCreateRequest struct {
	ctx context.Context
	ApiService *ContentInstallerFileIndicesAPIService
	debInstallerFileIndex *DebInstallerFileIndex
}

func (r ContentInstallerFileIndicesAPIContentDebInstallerFileIndicesCreateRequest) DebInstallerFileIndex(debInstallerFileIndex DebInstallerFileIndex) ContentInstallerFileIndicesAPIContentDebInstallerFileIndicesCreateRequest {
	r.debInstallerFileIndex = &debInstallerFileIndex
	return r
}

func (r ContentInstallerFileIndicesAPIContentDebInstallerFileIndicesCreateRequest) Execute() (*DebInstallerFileIndexResponse, *http.Response, error) {
	return r.ApiService.ContentDebInstallerFileIndicesCreateExecute(r)
}

/*
ContentDebInstallerFileIndicesCreate Create an installer file index

An InstallerFileIndex represents the indices for a set of installer files.

Associated artifacts: Exactly one 'SHA256SUMS' and/or 'MD5SUMS' file.

Each InstallerFileIndes is associated with a single component-architecture combination within
a single Release. Note that installer files are currently used exclusively for verbatim
publications. The APT publisher (both simple and structured mode) does not make use of installer
content.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ContentInstallerFileIndicesAPIContentDebInstallerFileIndicesCreateRequest
*/
func (a *ContentInstallerFileIndicesAPIService) ContentDebInstallerFileIndicesCreate(ctx context.Context) ContentInstallerFileIndicesAPIContentDebInstallerFileIndicesCreateRequest {
	return ContentInstallerFileIndicesAPIContentDebInstallerFileIndicesCreateRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return DebInstallerFileIndexResponse
func (a *ContentInstallerFileIndicesAPIService) ContentDebInstallerFileIndicesCreateExecute(r ContentInstallerFileIndicesAPIContentDebInstallerFileIndicesCreateRequest) (*DebInstallerFileIndexResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *DebInstallerFileIndexResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ContentInstallerFileIndicesAPIService.ContentDebInstallerFileIndicesCreate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pulp/api/v3/content/deb/installer_file_indices/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.debInstallerFileIndex == nil {
		return localVarReturnValue, nil, reportError("debInstallerFileIndex is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json", "application/x-www-form-urlencoded", "multipart/form-data"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.debInstallerFileIndex
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ContentInstallerFileIndicesAPIContentDebInstallerFileIndicesListRequest struct {
	ctx context.Context
	ApiService *ContentInstallerFileIndicesAPIService
	architecture *string
	component *string
	limit *int32
	offset *int32
	ordering *[]string
	pulpHrefIn *[]string
	pulpIdIn *[]string
	relativePath *string
	repositoryVersion *string
	repositoryVersionAdded *string
	repositoryVersionRemoved *string
	sha256 *string
	fields *[]string
	excludeFields *[]string
}

// Filter results where architecture matches value
func (r ContentInstallerFileIndicesAPIContentDebInstallerFileIndicesListRequest) Architecture(architecture string) ContentInstallerFileIndicesAPIContentDebInstallerFileIndicesListRequest {
	r.architecture = &architecture
	return r
}

// Filter results where component matches value
func (r ContentInstallerFileIndicesAPIContentDebInstallerFileIndicesListRequest) Component(component string) ContentInstallerFileIndicesAPIContentDebInstallerFileIndicesListRequest {
	r.component = &component
	return r
}

// Number of results to return per page.
func (r ContentInstallerFileIndicesAPIContentDebInstallerFileIndicesListRequest) Limit(limit int32) ContentInstallerFileIndicesAPIContentDebInstallerFileIndicesListRequest {
	r.limit = &limit
	return r
}

// The initial index from which to return the results.
func (r ContentInstallerFileIndicesAPIContentDebInstallerFileIndicesListRequest) Offset(offset int32) ContentInstallerFileIndicesAPIContentDebInstallerFileIndicesListRequest {
	r.offset = &offset
	return r
}

// Ordering  * &#x60;pulp_id&#x60; - Pulp id * &#x60;-pulp_id&#x60; - Pulp id (descending) * &#x60;pulp_created&#x60; - Pulp created * &#x60;-pulp_created&#x60; - Pulp created (descending) * &#x60;pulp_last_updated&#x60; - Pulp last updated * &#x60;-pulp_last_updated&#x60; - Pulp last updated (descending) * &#x60;pulp_type&#x60; - Pulp type * &#x60;-pulp_type&#x60; - Pulp type (descending) * &#x60;upstream_id&#x60; - Upstream id * &#x60;-upstream_id&#x60; - Upstream id (descending) * &#x60;timestamp_of_interest&#x60; - Timestamp of interest * &#x60;-timestamp_of_interest&#x60; - Timestamp of interest (descending) * &#x60;component&#x60; - Component * &#x60;-component&#x60; - Component (descending) * &#x60;architecture&#x60; - Architecture * &#x60;-architecture&#x60; - Architecture (descending) * &#x60;relative_path&#x60; - Relative path * &#x60;-relative_path&#x60; - Relative path (descending) * &#x60;sha256&#x60; - Sha256 * &#x60;-sha256&#x60; - Sha256 (descending) * &#x60;pk&#x60; - Pk * &#x60;-pk&#x60; - Pk (descending)
func (r ContentInstallerFileIndicesAPIContentDebInstallerFileIndicesListRequest) Ordering(ordering []string) ContentInstallerFileIndicesAPIContentDebInstallerFileIndicesListRequest {
	r.ordering = &ordering
	return r
}

// Multiple values may be separated by commas.
func (r ContentInstallerFileIndicesAPIContentDebInstallerFileIndicesListRequest) PulpHrefIn(pulpHrefIn []string) ContentInstallerFileIndicesAPIContentDebInstallerFileIndicesListRequest {
	r.pulpHrefIn = &pulpHrefIn
	return r
}

// Multiple values may be separated by commas.
func (r ContentInstallerFileIndicesAPIContentDebInstallerFileIndicesListRequest) PulpIdIn(pulpIdIn []string) ContentInstallerFileIndicesAPIContentDebInstallerFileIndicesListRequest {
	r.pulpIdIn = &pulpIdIn
	return r
}

// Filter results where relative_path matches value
func (r ContentInstallerFileIndicesAPIContentDebInstallerFileIndicesListRequest) RelativePath(relativePath string) ContentInstallerFileIndicesAPIContentDebInstallerFileIndicesListRequest {
	r.relativePath = &relativePath
	return r
}

// Repository Version referenced by HREF
func (r ContentInstallerFileIndicesAPIContentDebInstallerFileIndicesListRequest) RepositoryVersion(repositoryVersion string) ContentInstallerFileIndicesAPIContentDebInstallerFileIndicesListRequest {
	r.repositoryVersion = &repositoryVersion
	return r
}

// Repository Version referenced by HREF
func (r ContentInstallerFileIndicesAPIContentDebInstallerFileIndicesListRequest) RepositoryVersionAdded(repositoryVersionAdded string) ContentInstallerFileIndicesAPIContentDebInstallerFileIndicesListRequest {
	r.repositoryVersionAdded = &repositoryVersionAdded
	return r
}

// Repository Version referenced by HREF
func (r ContentInstallerFileIndicesAPIContentDebInstallerFileIndicesListRequest) RepositoryVersionRemoved(repositoryVersionRemoved string) ContentInstallerFileIndicesAPIContentDebInstallerFileIndicesListRequest {
	r.repositoryVersionRemoved = &repositoryVersionRemoved
	return r
}

// Filter results where sha256 matches value
func (r ContentInstallerFileIndicesAPIContentDebInstallerFileIndicesListRequest) Sha256(sha256 string) ContentInstallerFileIndicesAPIContentDebInstallerFileIndicesListRequest {
	r.sha256 = &sha256
	return r
}

// A list of fields to include in the response.
func (r ContentInstallerFileIndicesAPIContentDebInstallerFileIndicesListRequest) Fields(fields []string) ContentInstallerFileIndicesAPIContentDebInstallerFileIndicesListRequest {
	r.fields = &fields
	return r
}

// A list of fields to exclude from the response.
func (r ContentInstallerFileIndicesAPIContentDebInstallerFileIndicesListRequest) ExcludeFields(excludeFields []string) ContentInstallerFileIndicesAPIContentDebInstallerFileIndicesListRequest {
	r.excludeFields = &excludeFields
	return r
}

func (r ContentInstallerFileIndicesAPIContentDebInstallerFileIndicesListRequest) Execute() (*PaginateddebInstallerFileIndexResponseList, *http.Response, error) {
	return r.ApiService.ContentDebInstallerFileIndicesListExecute(r)
}

/*
ContentDebInstallerFileIndicesList List InstallerFileIndices

An InstallerFileIndex represents the indices for a set of installer files.

Associated artifacts: Exactly one 'SHA256SUMS' and/or 'MD5SUMS' file.

Each InstallerFileIndes is associated with a single component-architecture combination within
a single Release. Note that installer files are currently used exclusively for verbatim
publications. The APT publisher (both simple and structured mode) does not make use of installer
content.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ContentInstallerFileIndicesAPIContentDebInstallerFileIndicesListRequest
*/
func (a *ContentInstallerFileIndicesAPIService) ContentDebInstallerFileIndicesList(ctx context.Context) ContentInstallerFileIndicesAPIContentDebInstallerFileIndicesListRequest {
	return ContentInstallerFileIndicesAPIContentDebInstallerFileIndicesListRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return PaginateddebInstallerFileIndexResponseList
func (a *ContentInstallerFileIndicesAPIService) ContentDebInstallerFileIndicesListExecute(r ContentInstallerFileIndicesAPIContentDebInstallerFileIndicesListRequest) (*PaginateddebInstallerFileIndexResponseList, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PaginateddebInstallerFileIndexResponseList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ContentInstallerFileIndicesAPIService.ContentDebInstallerFileIndicesList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pulp/api/v3/content/deb/installer_file_indices/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.architecture != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "architecture", r.architecture, "")
	}
	if r.component != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "component", r.component, "")
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "")
	}
	if r.offset != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "offset", r.offset, "")
	}
	if r.ordering != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "ordering", r.ordering, "csv")
	}
	if r.pulpHrefIn != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pulp_href__in", r.pulpHrefIn, "csv")
	}
	if r.pulpIdIn != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pulp_id__in", r.pulpIdIn, "csv")
	}
	if r.relativePath != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "relative_path", r.relativePath, "")
	}
	if r.repositoryVersion != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "repository_version", r.repositoryVersion, "")
	}
	if r.repositoryVersionAdded != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "repository_version_added", r.repositoryVersionAdded, "")
	}
	if r.repositoryVersionRemoved != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "repository_version_removed", r.repositoryVersionRemoved, "")
	}
	if r.sha256 != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sha256", r.sha256, "")
	}
	if r.fields != nil {
		t := *r.fields
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "fields", s.Index(i), "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "fields", t, "multi")
		}
	}
	if r.excludeFields != nil {
		t := *r.excludeFields
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "exclude_fields", s.Index(i), "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "exclude_fields", t, "multi")
		}
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ContentInstallerFileIndicesAPIContentDebInstallerFileIndicesReadRequest struct {
	ctx context.Context
	ApiService *ContentInstallerFileIndicesAPIService
	debInstallerFileIndexHref string
	fields *[]string
	excludeFields *[]string
}

// A list of fields to include in the response.
func (r ContentInstallerFileIndicesAPIContentDebInstallerFileIndicesReadRequest) Fields(fields []string) ContentInstallerFileIndicesAPIContentDebInstallerFileIndicesReadRequest {
	r.fields = &fields
	return r
}

// A list of fields to exclude from the response.
func (r ContentInstallerFileIndicesAPIContentDebInstallerFileIndicesReadRequest) ExcludeFields(excludeFields []string) ContentInstallerFileIndicesAPIContentDebInstallerFileIndicesReadRequest {
	r.excludeFields = &excludeFields
	return r
}

func (r ContentInstallerFileIndicesAPIContentDebInstallerFileIndicesReadRequest) Execute() (*DebInstallerFileIndexResponse, *http.Response, error) {
	return r.ApiService.ContentDebInstallerFileIndicesReadExecute(r)
}

/*
ContentDebInstallerFileIndicesRead Inspect an installer file index

An InstallerFileIndex represents the indices for a set of installer files.

Associated artifacts: Exactly one 'SHA256SUMS' and/or 'MD5SUMS' file.

Each InstallerFileIndes is associated with a single component-architecture combination within
a single Release. Note that installer files are currently used exclusively for verbatim
publications. The APT publisher (both simple and structured mode) does not make use of installer
content.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param debInstallerFileIndexHref
 @return ContentInstallerFileIndicesAPIContentDebInstallerFileIndicesReadRequest
*/
func (a *ContentInstallerFileIndicesAPIService) ContentDebInstallerFileIndicesRead(ctx context.Context, debInstallerFileIndexHref string) ContentInstallerFileIndicesAPIContentDebInstallerFileIndicesReadRequest {
	return ContentInstallerFileIndicesAPIContentDebInstallerFileIndicesReadRequest{
		ApiService: a,
		ctx: ctx,
		debInstallerFileIndexHref: debInstallerFileIndexHref,
	}
}

// Execute executes the request
//  @return DebInstallerFileIndexResponse
func (a *ContentInstallerFileIndicesAPIService) ContentDebInstallerFileIndicesReadExecute(r ContentInstallerFileIndicesAPIContentDebInstallerFileIndicesReadRequest) (*DebInstallerFileIndexResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *DebInstallerFileIndexResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ContentInstallerFileIndicesAPIService.ContentDebInstallerFileIndicesRead")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "{deb_installer_file_index_href}"
	localVarPath = strings.Replace(localVarPath, "{"+"deb_installer_file_index_href"+"}", parameterValueToString(r.debInstallerFileIndexHref, "debInstallerFileIndexHref"), -1)  // NOTE: paths aren't escaped because Pulp uses hrefs as path parameters

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.fields != nil {
		t := *r.fields
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "fields", s.Index(i), "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "fields", t, "multi")
		}
	}
	if r.excludeFields != nil {
		t := *r.excludeFields
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "exclude_fields", s.Index(i), "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "exclude_fields", t, "multi")
		}
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
