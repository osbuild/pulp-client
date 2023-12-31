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
	"time"
	"reflect"
)


// PublicationsAptAPIService PublicationsAptAPI service
type PublicationsAptAPIService service

type PublicationsAptAPIPublicationsDebAptCreateRequest struct {
	ctx context.Context
	ApiService *PublicationsAptAPIService
	debAptPublication *DebAptPublication
}

func (r PublicationsAptAPIPublicationsDebAptCreateRequest) DebAptPublication(debAptPublication DebAptPublication) PublicationsAptAPIPublicationsDebAptCreateRequest {
	r.debAptPublication = &debAptPublication
	return r
}

func (r PublicationsAptAPIPublicationsDebAptCreateRequest) Execute() (*AsyncOperationResponse, *http.Response, error) {
	return r.ApiService.PublicationsDebAptCreateExecute(r)
}

/*
PublicationsDebAptCreate Create an apt publication

Trigger an asynchronous task to publish content

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return PublicationsAptAPIPublicationsDebAptCreateRequest
*/
func (a *PublicationsAptAPIService) PublicationsDebAptCreate(ctx context.Context) PublicationsAptAPIPublicationsDebAptCreateRequest {
	return PublicationsAptAPIPublicationsDebAptCreateRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return AsyncOperationResponse
func (a *PublicationsAptAPIService) PublicationsDebAptCreateExecute(r PublicationsAptAPIPublicationsDebAptCreateRequest) (*AsyncOperationResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AsyncOperationResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PublicationsAptAPIService.PublicationsDebAptCreate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pulp/api/v3/publications/deb/apt/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.debAptPublication == nil {
		return localVarReturnValue, nil, reportError("debAptPublication is required and must be specified")
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
	localVarPostBody = r.debAptPublication
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

type PublicationsAptAPIPublicationsDebAptDeleteRequest struct {
	ctx context.Context
	ApiService *PublicationsAptAPIService
	debAptPublicationHref string
}

func (r PublicationsAptAPIPublicationsDebAptDeleteRequest) Execute() (*http.Response, error) {
	return r.ApiService.PublicationsDebAptDeleteExecute(r)
}

/*
PublicationsDebAptDelete Delete an apt publication

An AptPublication is the ready to serve Pulp-internal representation of an AptRepositoryVersion.

When creating an APT publication, users must use simple or structured mode (or both). If the
publication should include '.deb' packages that were manually uploaded to the relevant
AptRepository, users must use 'simple=true'. Conversely, 'structured=true' is only useful for
publishing content obtained via synchronization. Once a Pulp publication has been created, it
can be served by creating a Pulp distribution (in a near atomic action).

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param debAptPublicationHref
 @return PublicationsAptAPIPublicationsDebAptDeleteRequest
*/
func (a *PublicationsAptAPIService) PublicationsDebAptDelete(ctx context.Context, debAptPublicationHref string) PublicationsAptAPIPublicationsDebAptDeleteRequest {
	return PublicationsAptAPIPublicationsDebAptDeleteRequest{
		ApiService: a,
		ctx: ctx,
		debAptPublicationHref: debAptPublicationHref,
	}
}

// Execute executes the request
func (a *PublicationsAptAPIService) PublicationsDebAptDeleteExecute(r PublicationsAptAPIPublicationsDebAptDeleteRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PublicationsAptAPIService.PublicationsDebAptDelete")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "{deb_apt_publication_href}"
	localVarPath = strings.Replace(localVarPath, "{"+"deb_apt_publication_href"+"}", parameterValueToString(r.debAptPublicationHref, "debAptPublicationHref"), -1)  // NOTE: paths aren't escaped because Pulp uses hrefs as path parameters

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
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type PublicationsAptAPIPublicationsDebAptListRequest struct {
	ctx context.Context
	ApiService *PublicationsAptAPIService
	content *string
	contentIn *string
	limit *int32
	offset *int32
	ordering *[]string
	pulpCreated *time.Time
	pulpCreatedGt *time.Time
	pulpCreatedGte *time.Time
	pulpCreatedLt *time.Time
	pulpCreatedLte *time.Time
	pulpCreatedRange *[]time.Time
	pulpHrefIn *[]string
	pulpIdIn *[]string
	repository *string
	repositoryVersion *string
	fields *[]string
	excludeFields *[]string
}

// Content Unit referenced by HREF
func (r PublicationsAptAPIPublicationsDebAptListRequest) Content(content string) PublicationsAptAPIPublicationsDebAptListRequest {
	r.content = &content
	return r
}

// Content Unit referenced by HREF
func (r PublicationsAptAPIPublicationsDebAptListRequest) ContentIn(contentIn string) PublicationsAptAPIPublicationsDebAptListRequest {
	r.contentIn = &contentIn
	return r
}

// Number of results to return per page.
func (r PublicationsAptAPIPublicationsDebAptListRequest) Limit(limit int32) PublicationsAptAPIPublicationsDebAptListRequest {
	r.limit = &limit
	return r
}

// The initial index from which to return the results.
func (r PublicationsAptAPIPublicationsDebAptListRequest) Offset(offset int32) PublicationsAptAPIPublicationsDebAptListRequest {
	r.offset = &offset
	return r
}

// Ordering  * &#x60;pulp_id&#x60; - Pulp id * &#x60;-pulp_id&#x60; - Pulp id (descending) * &#x60;pulp_created&#x60; - Pulp created * &#x60;-pulp_created&#x60; - Pulp created (descending) * &#x60;pulp_last_updated&#x60; - Pulp last updated * &#x60;-pulp_last_updated&#x60; - Pulp last updated (descending) * &#x60;pulp_type&#x60; - Pulp type * &#x60;-pulp_type&#x60; - Pulp type (descending) * &#x60;complete&#x60; - Complete * &#x60;-complete&#x60; - Complete (descending) * &#x60;pass_through&#x60; - Pass through * &#x60;-pass_through&#x60; - Pass through (descending) * &#x60;pk&#x60; - Pk * &#x60;-pk&#x60; - Pk (descending)
func (r PublicationsAptAPIPublicationsDebAptListRequest) Ordering(ordering []string) PublicationsAptAPIPublicationsDebAptListRequest {
	r.ordering = &ordering
	return r
}

// Filter results where pulp_created matches value
func (r PublicationsAptAPIPublicationsDebAptListRequest) PulpCreated(pulpCreated time.Time) PublicationsAptAPIPublicationsDebAptListRequest {
	r.pulpCreated = &pulpCreated
	return r
}

// Filter results where pulp_created is greater than value
func (r PublicationsAptAPIPublicationsDebAptListRequest) PulpCreatedGt(pulpCreatedGt time.Time) PublicationsAptAPIPublicationsDebAptListRequest {
	r.pulpCreatedGt = &pulpCreatedGt
	return r
}

// Filter results where pulp_created is greater than or equal to value
func (r PublicationsAptAPIPublicationsDebAptListRequest) PulpCreatedGte(pulpCreatedGte time.Time) PublicationsAptAPIPublicationsDebAptListRequest {
	r.pulpCreatedGte = &pulpCreatedGte
	return r
}

// Filter results where pulp_created is less than value
func (r PublicationsAptAPIPublicationsDebAptListRequest) PulpCreatedLt(pulpCreatedLt time.Time) PublicationsAptAPIPublicationsDebAptListRequest {
	r.pulpCreatedLt = &pulpCreatedLt
	return r
}

// Filter results where pulp_created is less than or equal to value
func (r PublicationsAptAPIPublicationsDebAptListRequest) PulpCreatedLte(pulpCreatedLte time.Time) PublicationsAptAPIPublicationsDebAptListRequest {
	r.pulpCreatedLte = &pulpCreatedLte
	return r
}

// Filter results where pulp_created is between two comma separated values
func (r PublicationsAptAPIPublicationsDebAptListRequest) PulpCreatedRange(pulpCreatedRange []time.Time) PublicationsAptAPIPublicationsDebAptListRequest {
	r.pulpCreatedRange = &pulpCreatedRange
	return r
}

// Multiple values may be separated by commas.
func (r PublicationsAptAPIPublicationsDebAptListRequest) PulpHrefIn(pulpHrefIn []string) PublicationsAptAPIPublicationsDebAptListRequest {
	r.pulpHrefIn = &pulpHrefIn
	return r
}

// Multiple values may be separated by commas.
func (r PublicationsAptAPIPublicationsDebAptListRequest) PulpIdIn(pulpIdIn []string) PublicationsAptAPIPublicationsDebAptListRequest {
	r.pulpIdIn = &pulpIdIn
	return r
}

// Repository referenced by HREF
func (r PublicationsAptAPIPublicationsDebAptListRequest) Repository(repository string) PublicationsAptAPIPublicationsDebAptListRequest {
	r.repository = &repository
	return r
}

// Repository Version referenced by HREF
func (r PublicationsAptAPIPublicationsDebAptListRequest) RepositoryVersion(repositoryVersion string) PublicationsAptAPIPublicationsDebAptListRequest {
	r.repositoryVersion = &repositoryVersion
	return r
}

// A list of fields to include in the response.
func (r PublicationsAptAPIPublicationsDebAptListRequest) Fields(fields []string) PublicationsAptAPIPublicationsDebAptListRequest {
	r.fields = &fields
	return r
}

// A list of fields to exclude from the response.
func (r PublicationsAptAPIPublicationsDebAptListRequest) ExcludeFields(excludeFields []string) PublicationsAptAPIPublicationsDebAptListRequest {
	r.excludeFields = &excludeFields
	return r
}

func (r PublicationsAptAPIPublicationsDebAptListRequest) Execute() (*PaginateddebAptPublicationResponseList, *http.Response, error) {
	return r.ApiService.PublicationsDebAptListExecute(r)
}

/*
PublicationsDebAptList List apt publications

An AptPublication is the ready to serve Pulp-internal representation of an AptRepositoryVersion.

When creating an APT publication, users must use simple or structured mode (or both). If the
publication should include '.deb' packages that were manually uploaded to the relevant
AptRepository, users must use 'simple=true'. Conversely, 'structured=true' is only useful for
publishing content obtained via synchronization. Once a Pulp publication has been created, it
can be served by creating a Pulp distribution (in a near atomic action).

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return PublicationsAptAPIPublicationsDebAptListRequest
*/
func (a *PublicationsAptAPIService) PublicationsDebAptList(ctx context.Context) PublicationsAptAPIPublicationsDebAptListRequest {
	return PublicationsAptAPIPublicationsDebAptListRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return PaginateddebAptPublicationResponseList
func (a *PublicationsAptAPIService) PublicationsDebAptListExecute(r PublicationsAptAPIPublicationsDebAptListRequest) (*PaginateddebAptPublicationResponseList, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PaginateddebAptPublicationResponseList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PublicationsAptAPIService.PublicationsDebAptList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pulp/api/v3/publications/deb/apt/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.content != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "content", r.content, "")
	}
	if r.contentIn != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "content__in", r.contentIn, "")
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
	if r.pulpCreated != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pulp_created", r.pulpCreated, "")
	}
	if r.pulpCreatedGt != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pulp_created__gt", r.pulpCreatedGt, "")
	}
	if r.pulpCreatedGte != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pulp_created__gte", r.pulpCreatedGte, "")
	}
	if r.pulpCreatedLt != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pulp_created__lt", r.pulpCreatedLt, "")
	}
	if r.pulpCreatedLte != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pulp_created__lte", r.pulpCreatedLte, "")
	}
	if r.pulpCreatedRange != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pulp_created__range", r.pulpCreatedRange, "csv")
	}
	if r.pulpHrefIn != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pulp_href__in", r.pulpHrefIn, "csv")
	}
	if r.pulpIdIn != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pulp_id__in", r.pulpIdIn, "csv")
	}
	if r.repository != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "repository", r.repository, "")
	}
	if r.repositoryVersion != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "repository_version", r.repositoryVersion, "")
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

type PublicationsAptAPIPublicationsDebAptReadRequest struct {
	ctx context.Context
	ApiService *PublicationsAptAPIService
	debAptPublicationHref string
	fields *[]string
	excludeFields *[]string
}

// A list of fields to include in the response.
func (r PublicationsAptAPIPublicationsDebAptReadRequest) Fields(fields []string) PublicationsAptAPIPublicationsDebAptReadRequest {
	r.fields = &fields
	return r
}

// A list of fields to exclude from the response.
func (r PublicationsAptAPIPublicationsDebAptReadRequest) ExcludeFields(excludeFields []string) PublicationsAptAPIPublicationsDebAptReadRequest {
	r.excludeFields = &excludeFields
	return r
}

func (r PublicationsAptAPIPublicationsDebAptReadRequest) Execute() (*DebAptPublicationResponse, *http.Response, error) {
	return r.ApiService.PublicationsDebAptReadExecute(r)
}

/*
PublicationsDebAptRead Inspect an apt publication

An AptPublication is the ready to serve Pulp-internal representation of an AptRepositoryVersion.

When creating an APT publication, users must use simple or structured mode (or both). If the
publication should include '.deb' packages that were manually uploaded to the relevant
AptRepository, users must use 'simple=true'. Conversely, 'structured=true' is only useful for
publishing content obtained via synchronization. Once a Pulp publication has been created, it
can be served by creating a Pulp distribution (in a near atomic action).

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param debAptPublicationHref
 @return PublicationsAptAPIPublicationsDebAptReadRequest
*/
func (a *PublicationsAptAPIService) PublicationsDebAptRead(ctx context.Context, debAptPublicationHref string) PublicationsAptAPIPublicationsDebAptReadRequest {
	return PublicationsAptAPIPublicationsDebAptReadRequest{
		ApiService: a,
		ctx: ctx,
		debAptPublicationHref: debAptPublicationHref,
	}
}

// Execute executes the request
//  @return DebAptPublicationResponse
func (a *PublicationsAptAPIService) PublicationsDebAptReadExecute(r PublicationsAptAPIPublicationsDebAptReadRequest) (*DebAptPublicationResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *DebAptPublicationResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PublicationsAptAPIService.PublicationsDebAptRead")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "{deb_apt_publication_href}"
	localVarPath = strings.Replace(localVarPath, "{"+"deb_apt_publication_href"+"}", parameterValueToString(r.debAptPublicationHref, "debAptPublicationHref"), -1)  // NOTE: paths aren't escaped because Pulp uses hrefs as path parameters

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
