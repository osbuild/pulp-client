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


// ContentPackageIndicesAPIService ContentPackageIndicesAPI service
type ContentPackageIndicesAPIService service

type ContentPackageIndicesAPIContentDebPackageIndicesCreateRequest struct {
	ctx context.Context
	ApiService *ContentPackageIndicesAPIService
	debPackageIndex *DebPackageIndex
}

func (r ContentPackageIndicesAPIContentDebPackageIndicesCreateRequest) DebPackageIndex(debPackageIndex DebPackageIndex) ContentPackageIndicesAPIContentDebPackageIndicesCreateRequest {
	r.debPackageIndex = &debPackageIndex
	return r
}

func (r ContentPackageIndicesAPIContentDebPackageIndicesCreateRequest) Execute() (*DebPackageIndexResponse, *http.Response, error) {
	return r.ApiService.ContentDebPackageIndicesCreateExecute(r)
}

/*
ContentDebPackageIndicesCreate Create a package index

A PackageIndex represents the package indices of a single component-architecture combination.

Associated artifacts: Exactly one 'Packages' file. May optionally include one or more of
'Packages.gz', 'Packages.xz', 'Release'. If included, the 'Release' file is a legacy
per-component-and-architecture Release file.

Note: The verbatim publisher will republish all associated artifacts, while the APT publisher
(both simple and structured mode) will generate any 'Packages' files it needs when creating the
publication. It does not make use of PackageIndex content.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ContentPackageIndicesAPIContentDebPackageIndicesCreateRequest
*/
func (a *ContentPackageIndicesAPIService) ContentDebPackageIndicesCreate(ctx context.Context) ContentPackageIndicesAPIContentDebPackageIndicesCreateRequest {
	return ContentPackageIndicesAPIContentDebPackageIndicesCreateRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return DebPackageIndexResponse
func (a *ContentPackageIndicesAPIService) ContentDebPackageIndicesCreateExecute(r ContentPackageIndicesAPIContentDebPackageIndicesCreateRequest) (*DebPackageIndexResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *DebPackageIndexResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ContentPackageIndicesAPIService.ContentDebPackageIndicesCreate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pulp/api/v3/content/deb/package_indices/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.debPackageIndex == nil {
		return localVarReturnValue, nil, reportError("debPackageIndex is required and must be specified")
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
	localVarPostBody = r.debPackageIndex
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

type ContentPackageIndicesAPIContentDebPackageIndicesListRequest struct {
	ctx context.Context
	ApiService *ContentPackageIndicesAPIService
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
func (r ContentPackageIndicesAPIContentDebPackageIndicesListRequest) Architecture(architecture string) ContentPackageIndicesAPIContentDebPackageIndicesListRequest {
	r.architecture = &architecture
	return r
}

// Filter results where component matches value
func (r ContentPackageIndicesAPIContentDebPackageIndicesListRequest) Component(component string) ContentPackageIndicesAPIContentDebPackageIndicesListRequest {
	r.component = &component
	return r
}

// Number of results to return per page.
func (r ContentPackageIndicesAPIContentDebPackageIndicesListRequest) Limit(limit int32) ContentPackageIndicesAPIContentDebPackageIndicesListRequest {
	r.limit = &limit
	return r
}

// The initial index from which to return the results.
func (r ContentPackageIndicesAPIContentDebPackageIndicesListRequest) Offset(offset int32) ContentPackageIndicesAPIContentDebPackageIndicesListRequest {
	r.offset = &offset
	return r
}

// Ordering  * &#x60;pulp_id&#x60; - Pulp id * &#x60;-pulp_id&#x60; - Pulp id (descending) * &#x60;pulp_created&#x60; - Pulp created * &#x60;-pulp_created&#x60; - Pulp created (descending) * &#x60;pulp_last_updated&#x60; - Pulp last updated * &#x60;-pulp_last_updated&#x60; - Pulp last updated (descending) * &#x60;pulp_type&#x60; - Pulp type * &#x60;-pulp_type&#x60; - Pulp type (descending) * &#x60;upstream_id&#x60; - Upstream id * &#x60;-upstream_id&#x60; - Upstream id (descending) * &#x60;timestamp_of_interest&#x60; - Timestamp of interest * &#x60;-timestamp_of_interest&#x60; - Timestamp of interest (descending) * &#x60;component&#x60; - Component * &#x60;-component&#x60; - Component (descending) * &#x60;architecture&#x60; - Architecture * &#x60;-architecture&#x60; - Architecture (descending) * &#x60;relative_path&#x60; - Relative path * &#x60;-relative_path&#x60; - Relative path (descending) * &#x60;sha256&#x60; - Sha256 * &#x60;-sha256&#x60; - Sha256 (descending) * &#x60;artifact_set_sha256&#x60; - Artifact set sha256 * &#x60;-artifact_set_sha256&#x60; - Artifact set sha256 (descending) * &#x60;pk&#x60; - Pk * &#x60;-pk&#x60; - Pk (descending)
func (r ContentPackageIndicesAPIContentDebPackageIndicesListRequest) Ordering(ordering []string) ContentPackageIndicesAPIContentDebPackageIndicesListRequest {
	r.ordering = &ordering
	return r
}

// Multiple values may be separated by commas.
func (r ContentPackageIndicesAPIContentDebPackageIndicesListRequest) PulpHrefIn(pulpHrefIn []string) ContentPackageIndicesAPIContentDebPackageIndicesListRequest {
	r.pulpHrefIn = &pulpHrefIn
	return r
}

// Multiple values may be separated by commas.
func (r ContentPackageIndicesAPIContentDebPackageIndicesListRequest) PulpIdIn(pulpIdIn []string) ContentPackageIndicesAPIContentDebPackageIndicesListRequest {
	r.pulpIdIn = &pulpIdIn
	return r
}

// Filter results where relative_path matches value
func (r ContentPackageIndicesAPIContentDebPackageIndicesListRequest) RelativePath(relativePath string) ContentPackageIndicesAPIContentDebPackageIndicesListRequest {
	r.relativePath = &relativePath
	return r
}

// Repository Version referenced by HREF
func (r ContentPackageIndicesAPIContentDebPackageIndicesListRequest) RepositoryVersion(repositoryVersion string) ContentPackageIndicesAPIContentDebPackageIndicesListRequest {
	r.repositoryVersion = &repositoryVersion
	return r
}

// Repository Version referenced by HREF
func (r ContentPackageIndicesAPIContentDebPackageIndicesListRequest) RepositoryVersionAdded(repositoryVersionAdded string) ContentPackageIndicesAPIContentDebPackageIndicesListRequest {
	r.repositoryVersionAdded = &repositoryVersionAdded
	return r
}

// Repository Version referenced by HREF
func (r ContentPackageIndicesAPIContentDebPackageIndicesListRequest) RepositoryVersionRemoved(repositoryVersionRemoved string) ContentPackageIndicesAPIContentDebPackageIndicesListRequest {
	r.repositoryVersionRemoved = &repositoryVersionRemoved
	return r
}

// Filter results where sha256 matches value
func (r ContentPackageIndicesAPIContentDebPackageIndicesListRequest) Sha256(sha256 string) ContentPackageIndicesAPIContentDebPackageIndicesListRequest {
	r.sha256 = &sha256
	return r
}

// A list of fields to include in the response.
func (r ContentPackageIndicesAPIContentDebPackageIndicesListRequest) Fields(fields []string) ContentPackageIndicesAPIContentDebPackageIndicesListRequest {
	r.fields = &fields
	return r
}

// A list of fields to exclude from the response.
func (r ContentPackageIndicesAPIContentDebPackageIndicesListRequest) ExcludeFields(excludeFields []string) ContentPackageIndicesAPIContentDebPackageIndicesListRequest {
	r.excludeFields = &excludeFields
	return r
}

func (r ContentPackageIndicesAPIContentDebPackageIndicesListRequest) Execute() (*PaginateddebPackageIndexResponseList, *http.Response, error) {
	return r.ApiService.ContentDebPackageIndicesListExecute(r)
}

/*
ContentDebPackageIndicesList List PackageIndices

A PackageIndex represents the package indices of a single component-architecture combination.

Associated artifacts: Exactly one 'Packages' file. May optionally include one or more of
'Packages.gz', 'Packages.xz', 'Release'. If included, the 'Release' file is a legacy
per-component-and-architecture Release file.

Note: The verbatim publisher will republish all associated artifacts, while the APT publisher
(both simple and structured mode) will generate any 'Packages' files it needs when creating the
publication. It does not make use of PackageIndex content.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ContentPackageIndicesAPIContentDebPackageIndicesListRequest
*/
func (a *ContentPackageIndicesAPIService) ContentDebPackageIndicesList(ctx context.Context) ContentPackageIndicesAPIContentDebPackageIndicesListRequest {
	return ContentPackageIndicesAPIContentDebPackageIndicesListRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return PaginateddebPackageIndexResponseList
func (a *ContentPackageIndicesAPIService) ContentDebPackageIndicesListExecute(r ContentPackageIndicesAPIContentDebPackageIndicesListRequest) (*PaginateddebPackageIndexResponseList, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PaginateddebPackageIndexResponseList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ContentPackageIndicesAPIService.ContentDebPackageIndicesList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pulp/api/v3/content/deb/package_indices/"

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

type ContentPackageIndicesAPIContentDebPackageIndicesReadRequest struct {
	ctx context.Context
	ApiService *ContentPackageIndicesAPIService
	debPackageIndexHref string
	fields *[]string
	excludeFields *[]string
}

// A list of fields to include in the response.
func (r ContentPackageIndicesAPIContentDebPackageIndicesReadRequest) Fields(fields []string) ContentPackageIndicesAPIContentDebPackageIndicesReadRequest {
	r.fields = &fields
	return r
}

// A list of fields to exclude from the response.
func (r ContentPackageIndicesAPIContentDebPackageIndicesReadRequest) ExcludeFields(excludeFields []string) ContentPackageIndicesAPIContentDebPackageIndicesReadRequest {
	r.excludeFields = &excludeFields
	return r
}

func (r ContentPackageIndicesAPIContentDebPackageIndicesReadRequest) Execute() (*DebPackageIndexResponse, *http.Response, error) {
	return r.ApiService.ContentDebPackageIndicesReadExecute(r)
}

/*
ContentDebPackageIndicesRead Inspect a package index

A PackageIndex represents the package indices of a single component-architecture combination.

Associated artifacts: Exactly one 'Packages' file. May optionally include one or more of
'Packages.gz', 'Packages.xz', 'Release'. If included, the 'Release' file is a legacy
per-component-and-architecture Release file.

Note: The verbatim publisher will republish all associated artifacts, while the APT publisher
(both simple and structured mode) will generate any 'Packages' files it needs when creating the
publication. It does not make use of PackageIndex content.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param debPackageIndexHref
 @return ContentPackageIndicesAPIContentDebPackageIndicesReadRequest
*/
func (a *ContentPackageIndicesAPIService) ContentDebPackageIndicesRead(ctx context.Context, debPackageIndexHref string) ContentPackageIndicesAPIContentDebPackageIndicesReadRequest {
	return ContentPackageIndicesAPIContentDebPackageIndicesReadRequest{
		ApiService: a,
		ctx: ctx,
		debPackageIndexHref: debPackageIndexHref,
	}
}

// Execute executes the request
//  @return DebPackageIndexResponse
func (a *ContentPackageIndicesAPIService) ContentDebPackageIndicesReadExecute(r ContentPackageIndicesAPIContentDebPackageIndicesReadRequest) (*DebPackageIndexResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *DebPackageIndexResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ContentPackageIndicesAPIService.ContentDebPackageIndicesRead")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "{deb_package_index_href}"
	localVarPath = strings.Replace(localVarPath, "{"+"deb_package_index_href"+"}", parameterValueToString(r.debPackageIndexHref, "debPackageIndexHref"), -1)  // NOTE: paths aren't escaped because Pulp uses hrefs as path parameters

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
