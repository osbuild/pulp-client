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


// PulpAnsibleApiV3CollectionsVersionsAPIService PulpAnsibleApiV3CollectionsVersionsAPI service
type PulpAnsibleApiV3CollectionsVersionsAPIService service

type PulpAnsibleApiV3CollectionsVersionsAPIPulpAnsibleGalaxyApiV3CollectionsVersionsDeleteRequest struct {
	ctx context.Context
	ApiService *PulpAnsibleApiV3CollectionsVersionsAPIService
	name string
	namespace string
	path string
	version string
}

func (r PulpAnsibleApiV3CollectionsVersionsAPIPulpAnsibleGalaxyApiV3CollectionsVersionsDeleteRequest) Execute() (*AsyncOperationResponse, *http.Response, error) {
	return r.ApiService.PulpAnsibleGalaxyApiV3CollectionsVersionsDeleteExecute(r)
}

/*
PulpAnsibleGalaxyApiV3CollectionsVersionsDelete Method for PulpAnsibleGalaxyApiV3CollectionsVersionsDelete

Legacy v3 endpoint.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param name
 @param namespace
 @param path
 @param version
 @return PulpAnsibleApiV3CollectionsVersionsAPIPulpAnsibleGalaxyApiV3CollectionsVersionsDeleteRequest

Deprecated
*/
func (a *PulpAnsibleApiV3CollectionsVersionsAPIService) PulpAnsibleGalaxyApiV3CollectionsVersionsDelete(ctx context.Context, name string, namespace string, path string, version string) PulpAnsibleApiV3CollectionsVersionsAPIPulpAnsibleGalaxyApiV3CollectionsVersionsDeleteRequest {
	return PulpAnsibleApiV3CollectionsVersionsAPIPulpAnsibleGalaxyApiV3CollectionsVersionsDeleteRequest{
		ApiService: a,
		ctx: ctx,
		name: name,
		namespace: namespace,
		path: path,
		version: version,
	}
}

// Execute executes the request
//  @return AsyncOperationResponse
// Deprecated
func (a *PulpAnsibleApiV3CollectionsVersionsAPIService) PulpAnsibleGalaxyApiV3CollectionsVersionsDeleteExecute(r PulpAnsibleApiV3CollectionsVersionsAPIPulpAnsibleGalaxyApiV3CollectionsVersionsDeleteRequest) (*AsyncOperationResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AsyncOperationResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PulpAnsibleApiV3CollectionsVersionsAPIService.PulpAnsibleGalaxyApiV3CollectionsVersionsDelete")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pulp_ansible/galaxy/{path}/api/v3/collections/{namespace}/{name}/versions/{version}/"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", parameterValueToString(r.name, "name"), -1)  // NOTE: paths aren't escaped because Pulp uses hrefs as path parameters
	localVarPath = strings.Replace(localVarPath, "{"+"namespace"+"}", parameterValueToString(r.namespace, "namespace"), -1)  // NOTE: paths aren't escaped because Pulp uses hrefs as path parameters
	localVarPath = strings.Replace(localVarPath, "{"+"path"+"}", parameterValueToString(r.path, "path"), -1)  // NOTE: paths aren't escaped because Pulp uses hrefs as path parameters
	localVarPath = strings.Replace(localVarPath, "{"+"version"+"}", parameterValueToString(r.version, "version"), -1)  // NOTE: paths aren't escaped because Pulp uses hrefs as path parameters

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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

type PulpAnsibleApiV3CollectionsVersionsAPIPulpAnsibleGalaxyApiV3CollectionsVersionsListRequest struct {
	ctx context.Context
	ApiService *PulpAnsibleApiV3CollectionsVersionsAPIService
	name string
	namespace string
	path string
	isHighest *bool
	limit *int32
	name2 *string
	namespace2 *string
	offset *int32
	ordering *[]string
	pulpHrefIn *[]string
	pulpIdIn *[]string
	q *string
	repositoryVersion *string
	repositoryVersionAdded *string
	repositoryVersionRemoved *string
	tags *string
	version *string
	fields *[]string
	excludeFields *[]string
}

func (r PulpAnsibleApiV3CollectionsVersionsAPIPulpAnsibleGalaxyApiV3CollectionsVersionsListRequest) IsHighest(isHighest bool) PulpAnsibleApiV3CollectionsVersionsAPIPulpAnsibleGalaxyApiV3CollectionsVersionsListRequest {
	r.isHighest = &isHighest
	return r
}

// Number of results to return per page.
func (r PulpAnsibleApiV3CollectionsVersionsAPIPulpAnsibleGalaxyApiV3CollectionsVersionsListRequest) Limit(limit int32) PulpAnsibleApiV3CollectionsVersionsAPIPulpAnsibleGalaxyApiV3CollectionsVersionsListRequest {
	r.limit = &limit
	return r
}

func (r PulpAnsibleApiV3CollectionsVersionsAPIPulpAnsibleGalaxyApiV3CollectionsVersionsListRequest) Name2(name2 string) PulpAnsibleApiV3CollectionsVersionsAPIPulpAnsibleGalaxyApiV3CollectionsVersionsListRequest {
	r.name2 = &name2
	return r
}

func (r PulpAnsibleApiV3CollectionsVersionsAPIPulpAnsibleGalaxyApiV3CollectionsVersionsListRequest) Namespace2(namespace2 string) PulpAnsibleApiV3CollectionsVersionsAPIPulpAnsibleGalaxyApiV3CollectionsVersionsListRequest {
	r.namespace2 = &namespace2
	return r
}

// The initial index from which to return the results.
func (r PulpAnsibleApiV3CollectionsVersionsAPIPulpAnsibleGalaxyApiV3CollectionsVersionsListRequest) Offset(offset int32) PulpAnsibleApiV3CollectionsVersionsAPIPulpAnsibleGalaxyApiV3CollectionsVersionsListRequest {
	r.offset = &offset
	return r
}

// Ordering  * &#x60;pulp_id&#x60; - Pulp id * &#x60;-pulp_id&#x60; - Pulp id (descending) * &#x60;pulp_created&#x60; - Pulp created * &#x60;-pulp_created&#x60; - Pulp created (descending) * &#x60;pulp_last_updated&#x60; - Pulp last updated * &#x60;-pulp_last_updated&#x60; - Pulp last updated (descending) * &#x60;pulp_type&#x60; - Pulp type * &#x60;-pulp_type&#x60; - Pulp type (descending) * &#x60;upstream_id&#x60; - Upstream id * &#x60;-upstream_id&#x60; - Upstream id (descending) * &#x60;timestamp_of_interest&#x60; - Timestamp of interest * &#x60;-timestamp_of_interest&#x60; - Timestamp of interest (descending) * &#x60;authors&#x60; - Authors * &#x60;-authors&#x60; - Authors (descending) * &#x60;contents&#x60; - Contents * &#x60;-contents&#x60; - Contents (descending) * &#x60;dependencies&#x60; - Dependencies * &#x60;-dependencies&#x60; - Dependencies (descending) * &#x60;description&#x60; - Description * &#x60;-description&#x60; - Description (descending) * &#x60;docs_blob&#x60; - Docs blob * &#x60;-docs_blob&#x60; - Docs blob (descending) * &#x60;manifest&#x60; - Manifest * &#x60;-manifest&#x60; - Manifest (descending) * &#x60;files&#x60; - Files * &#x60;-files&#x60; - Files (descending) * &#x60;documentation&#x60; - Documentation * &#x60;-documentation&#x60; - Documentation (descending) * &#x60;homepage&#x60; - Homepage * &#x60;-homepage&#x60; - Homepage (descending) * &#x60;issues&#x60; - Issues * &#x60;-issues&#x60; - Issues (descending) * &#x60;license&#x60; - License * &#x60;-license&#x60; - License (descending) * &#x60;name&#x60; - Name * &#x60;-name&#x60; - Name (descending) * &#x60;namespace&#x60; - Namespace * &#x60;-namespace&#x60; - Namespace (descending) * &#x60;repository&#x60; - Repository * &#x60;-repository&#x60; - Repository (descending) * &#x60;requires_ansible&#x60; - Requires ansible * &#x60;-requires_ansible&#x60; - Requires ansible (descending) * &#x60;version&#x60; - Version * &#x60;-version&#x60; - Version (descending) * &#x60;version_major&#x60; - Version major * &#x60;-version_major&#x60; - Version major (descending) * &#x60;version_minor&#x60; - Version minor * &#x60;-version_minor&#x60; - Version minor (descending) * &#x60;version_patch&#x60; - Version patch * &#x60;-version_patch&#x60; - Version patch (descending) * &#x60;version_prerelease&#x60; - Version prerelease * &#x60;-version_prerelease&#x60; - Version prerelease (descending) * &#x60;is_highest&#x60; - Is highest * &#x60;-is_highest&#x60; - Is highest (descending) * &#x60;search_vector&#x60; - Search vector * &#x60;-search_vector&#x60; - Search vector (descending) * &#x60;pk&#x60; - Pk * &#x60;-pk&#x60; - Pk (descending)
func (r PulpAnsibleApiV3CollectionsVersionsAPIPulpAnsibleGalaxyApiV3CollectionsVersionsListRequest) Ordering(ordering []string) PulpAnsibleApiV3CollectionsVersionsAPIPulpAnsibleGalaxyApiV3CollectionsVersionsListRequest {
	r.ordering = &ordering
	return r
}

// Multiple values may be separated by commas.
func (r PulpAnsibleApiV3CollectionsVersionsAPIPulpAnsibleGalaxyApiV3CollectionsVersionsListRequest) PulpHrefIn(pulpHrefIn []string) PulpAnsibleApiV3CollectionsVersionsAPIPulpAnsibleGalaxyApiV3CollectionsVersionsListRequest {
	r.pulpHrefIn = &pulpHrefIn
	return r
}

// Multiple values may be separated by commas.
func (r PulpAnsibleApiV3CollectionsVersionsAPIPulpAnsibleGalaxyApiV3CollectionsVersionsListRequest) PulpIdIn(pulpIdIn []string) PulpAnsibleApiV3CollectionsVersionsAPIPulpAnsibleGalaxyApiV3CollectionsVersionsListRequest {
	r.pulpIdIn = &pulpIdIn
	return r
}

func (r PulpAnsibleApiV3CollectionsVersionsAPIPulpAnsibleGalaxyApiV3CollectionsVersionsListRequest) Q(q string) PulpAnsibleApiV3CollectionsVersionsAPIPulpAnsibleGalaxyApiV3CollectionsVersionsListRequest {
	r.q = &q
	return r
}

// Repository Version referenced by HREF
func (r PulpAnsibleApiV3CollectionsVersionsAPIPulpAnsibleGalaxyApiV3CollectionsVersionsListRequest) RepositoryVersion(repositoryVersion string) PulpAnsibleApiV3CollectionsVersionsAPIPulpAnsibleGalaxyApiV3CollectionsVersionsListRequest {
	r.repositoryVersion = &repositoryVersion
	return r
}

// Repository Version referenced by HREF
func (r PulpAnsibleApiV3CollectionsVersionsAPIPulpAnsibleGalaxyApiV3CollectionsVersionsListRequest) RepositoryVersionAdded(repositoryVersionAdded string) PulpAnsibleApiV3CollectionsVersionsAPIPulpAnsibleGalaxyApiV3CollectionsVersionsListRequest {
	r.repositoryVersionAdded = &repositoryVersionAdded
	return r
}

// Repository Version referenced by HREF
func (r PulpAnsibleApiV3CollectionsVersionsAPIPulpAnsibleGalaxyApiV3CollectionsVersionsListRequest) RepositoryVersionRemoved(repositoryVersionRemoved string) PulpAnsibleApiV3CollectionsVersionsAPIPulpAnsibleGalaxyApiV3CollectionsVersionsListRequest {
	r.repositoryVersionRemoved = &repositoryVersionRemoved
	return r
}

// Filter by comma separate list of tags that must all be matched
func (r PulpAnsibleApiV3CollectionsVersionsAPIPulpAnsibleGalaxyApiV3CollectionsVersionsListRequest) Tags(tags string) PulpAnsibleApiV3CollectionsVersionsAPIPulpAnsibleGalaxyApiV3CollectionsVersionsListRequest {
	r.tags = &tags
	return r
}

// Filter results where version matches value
func (r PulpAnsibleApiV3CollectionsVersionsAPIPulpAnsibleGalaxyApiV3CollectionsVersionsListRequest) Version(version string) PulpAnsibleApiV3CollectionsVersionsAPIPulpAnsibleGalaxyApiV3CollectionsVersionsListRequest {
	r.version = &version
	return r
}

// A list of fields to include in the response.
func (r PulpAnsibleApiV3CollectionsVersionsAPIPulpAnsibleGalaxyApiV3CollectionsVersionsListRequest) Fields(fields []string) PulpAnsibleApiV3CollectionsVersionsAPIPulpAnsibleGalaxyApiV3CollectionsVersionsListRequest {
	r.fields = &fields
	return r
}

// A list of fields to exclude from the response.
func (r PulpAnsibleApiV3CollectionsVersionsAPIPulpAnsibleGalaxyApiV3CollectionsVersionsListRequest) ExcludeFields(excludeFields []string) PulpAnsibleApiV3CollectionsVersionsAPIPulpAnsibleGalaxyApiV3CollectionsVersionsListRequest {
	r.excludeFields = &excludeFields
	return r
}

func (r PulpAnsibleApiV3CollectionsVersionsAPIPulpAnsibleGalaxyApiV3CollectionsVersionsListRequest) Execute() (*PaginatedCollectionVersionListResponseList, *http.Response, error) {
	return r.ApiService.PulpAnsibleGalaxyApiV3CollectionsVersionsListExecute(r)
}

/*
PulpAnsibleGalaxyApiV3CollectionsVersionsList Method for PulpAnsibleGalaxyApiV3CollectionsVersionsList

Legacy v3 endpoint.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param name
 @param namespace
 @param path
 @return PulpAnsibleApiV3CollectionsVersionsAPIPulpAnsibleGalaxyApiV3CollectionsVersionsListRequest

Deprecated
*/
func (a *PulpAnsibleApiV3CollectionsVersionsAPIService) PulpAnsibleGalaxyApiV3CollectionsVersionsList(ctx context.Context, name string, namespace string, path string) PulpAnsibleApiV3CollectionsVersionsAPIPulpAnsibleGalaxyApiV3CollectionsVersionsListRequest {
	return PulpAnsibleApiV3CollectionsVersionsAPIPulpAnsibleGalaxyApiV3CollectionsVersionsListRequest{
		ApiService: a,
		ctx: ctx,
		name: name,
		namespace: namespace,
		path: path,
	}
}

// Execute executes the request
//  @return PaginatedCollectionVersionListResponseList
// Deprecated
func (a *PulpAnsibleApiV3CollectionsVersionsAPIService) PulpAnsibleGalaxyApiV3CollectionsVersionsListExecute(r PulpAnsibleApiV3CollectionsVersionsAPIPulpAnsibleGalaxyApiV3CollectionsVersionsListRequest) (*PaginatedCollectionVersionListResponseList, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PaginatedCollectionVersionListResponseList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PulpAnsibleApiV3CollectionsVersionsAPIService.PulpAnsibleGalaxyApiV3CollectionsVersionsList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pulp_ansible/galaxy/{path}/api/v3/collections/{namespace}/{name}/versions/"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", parameterValueToString(r.name, "name"), -1)  // NOTE: paths aren't escaped because Pulp uses hrefs as path parameters
	localVarPath = strings.Replace(localVarPath, "{"+"namespace"+"}", parameterValueToString(r.namespace, "namespace"), -1)  // NOTE: paths aren't escaped because Pulp uses hrefs as path parameters
	localVarPath = strings.Replace(localVarPath, "{"+"path"+"}", parameterValueToString(r.path, "path"), -1)  // NOTE: paths aren't escaped because Pulp uses hrefs as path parameters

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.isHighest != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "is_highest", r.isHighest, "")
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "")
	}
	if r.name2 != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "name", r.name2, "")
	}
	if r.namespace2 != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "namespace", r.namespace2, "")
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
	if r.q != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "q", r.q, "")
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
	if r.tags != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "tags", r.tags, "")
	}
	if r.version != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "version", r.version, "")
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

type PulpAnsibleApiV3CollectionsVersionsAPIPulpAnsibleGalaxyApiV3CollectionsVersionsReadRequest struct {
	ctx context.Context
	ApiService *PulpAnsibleApiV3CollectionsVersionsAPIService
	name string
	namespace string
	path string
	version string
	fields *[]string
	excludeFields *[]string
}

// A list of fields to include in the response.
func (r PulpAnsibleApiV3CollectionsVersionsAPIPulpAnsibleGalaxyApiV3CollectionsVersionsReadRequest) Fields(fields []string) PulpAnsibleApiV3CollectionsVersionsAPIPulpAnsibleGalaxyApiV3CollectionsVersionsReadRequest {
	r.fields = &fields
	return r
}

// A list of fields to exclude from the response.
func (r PulpAnsibleApiV3CollectionsVersionsAPIPulpAnsibleGalaxyApiV3CollectionsVersionsReadRequest) ExcludeFields(excludeFields []string) PulpAnsibleApiV3CollectionsVersionsAPIPulpAnsibleGalaxyApiV3CollectionsVersionsReadRequest {
	r.excludeFields = &excludeFields
	return r
}

func (r PulpAnsibleApiV3CollectionsVersionsAPIPulpAnsibleGalaxyApiV3CollectionsVersionsReadRequest) Execute() (*CollectionVersionResponse, *http.Response, error) {
	return r.ApiService.PulpAnsibleGalaxyApiV3CollectionsVersionsReadExecute(r)
}

/*
PulpAnsibleGalaxyApiV3CollectionsVersionsRead Method for PulpAnsibleGalaxyApiV3CollectionsVersionsRead

Legacy v3 endpoint.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param name
 @param namespace
 @param path
 @param version
 @return PulpAnsibleApiV3CollectionsVersionsAPIPulpAnsibleGalaxyApiV3CollectionsVersionsReadRequest

Deprecated
*/
func (a *PulpAnsibleApiV3CollectionsVersionsAPIService) PulpAnsibleGalaxyApiV3CollectionsVersionsRead(ctx context.Context, name string, namespace string, path string, version string) PulpAnsibleApiV3CollectionsVersionsAPIPulpAnsibleGalaxyApiV3CollectionsVersionsReadRequest {
	return PulpAnsibleApiV3CollectionsVersionsAPIPulpAnsibleGalaxyApiV3CollectionsVersionsReadRequest{
		ApiService: a,
		ctx: ctx,
		name: name,
		namespace: namespace,
		path: path,
		version: version,
	}
}

// Execute executes the request
//  @return CollectionVersionResponse
// Deprecated
func (a *PulpAnsibleApiV3CollectionsVersionsAPIService) PulpAnsibleGalaxyApiV3CollectionsVersionsReadExecute(r PulpAnsibleApiV3CollectionsVersionsAPIPulpAnsibleGalaxyApiV3CollectionsVersionsReadRequest) (*CollectionVersionResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CollectionVersionResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PulpAnsibleApiV3CollectionsVersionsAPIService.PulpAnsibleGalaxyApiV3CollectionsVersionsRead")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pulp_ansible/galaxy/{path}/api/v3/collections/{namespace}/{name}/versions/{version}/"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", parameterValueToString(r.name, "name"), -1)  // NOTE: paths aren't escaped because Pulp uses hrefs as path parameters
	localVarPath = strings.Replace(localVarPath, "{"+"namespace"+"}", parameterValueToString(r.namespace, "namespace"), -1)  // NOTE: paths aren't escaped because Pulp uses hrefs as path parameters
	localVarPath = strings.Replace(localVarPath, "{"+"path"+"}", parameterValueToString(r.path, "path"), -1)  // NOTE: paths aren't escaped because Pulp uses hrefs as path parameters
	localVarPath = strings.Replace(localVarPath, "{"+"version"+"}", parameterValueToString(r.version, "version"), -1)  // NOTE: paths aren't escaped because Pulp uses hrefs as path parameters

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
