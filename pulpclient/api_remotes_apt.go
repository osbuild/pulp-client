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


// RemotesAptAPIService RemotesAptAPI service
type RemotesAptAPIService service

type RemotesAptAPIRemotesDebAptCreateRequest struct {
	ctx context.Context
	ApiService *RemotesAptAPIService
	debAptRemote *DebAptRemote
}

func (r RemotesAptAPIRemotesDebAptCreateRequest) DebAptRemote(debAptRemote DebAptRemote) RemotesAptAPIRemotesDebAptCreateRequest {
	r.debAptRemote = &debAptRemote
	return r
}

func (r RemotesAptAPIRemotesDebAptCreateRequest) Execute() (*DebAptRemoteResponse, *http.Response, error) {
	return r.ApiService.RemotesDebAptCreateExecute(r)
}

/*
RemotesDebAptCreate Create an apt remote

An AptRemote represents an external APT repository content source.

It contains the location of the upstream APT repository, as well as the user options that are
applied when using the remote to synchronize the upstream repository to Pulp.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return RemotesAptAPIRemotesDebAptCreateRequest
*/
func (a *RemotesAptAPIService) RemotesDebAptCreate(ctx context.Context) RemotesAptAPIRemotesDebAptCreateRequest {
	return RemotesAptAPIRemotesDebAptCreateRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return DebAptRemoteResponse
func (a *RemotesAptAPIService) RemotesDebAptCreateExecute(r RemotesAptAPIRemotesDebAptCreateRequest) (*DebAptRemoteResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *DebAptRemoteResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RemotesAptAPIService.RemotesDebAptCreate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pulp/api/v3/remotes/deb/apt/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.debAptRemote == nil {
		return localVarReturnValue, nil, reportError("debAptRemote is required and must be specified")
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
	localVarPostBody = r.debAptRemote
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

type RemotesAptAPIRemotesDebAptDeleteRequest struct {
	ctx context.Context
	ApiService *RemotesAptAPIService
	debAptRemoteHref string
}

func (r RemotesAptAPIRemotesDebAptDeleteRequest) Execute() (*AsyncOperationResponse, *http.Response, error) {
	return r.ApiService.RemotesDebAptDeleteExecute(r)
}

/*
RemotesDebAptDelete Delete an apt remote

Trigger an asynchronous delete task

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param debAptRemoteHref
 @return RemotesAptAPIRemotesDebAptDeleteRequest
*/
func (a *RemotesAptAPIService) RemotesDebAptDelete(ctx context.Context, debAptRemoteHref string) RemotesAptAPIRemotesDebAptDeleteRequest {
	return RemotesAptAPIRemotesDebAptDeleteRequest{
		ApiService: a,
		ctx: ctx,
		debAptRemoteHref: debAptRemoteHref,
	}
}

// Execute executes the request
//  @return AsyncOperationResponse
func (a *RemotesAptAPIService) RemotesDebAptDeleteExecute(r RemotesAptAPIRemotesDebAptDeleteRequest) (*AsyncOperationResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AsyncOperationResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RemotesAptAPIService.RemotesDebAptDelete")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "{deb_apt_remote_href}"
	localVarPath = strings.Replace(localVarPath, "{"+"deb_apt_remote_href"+"}", parameterValueToString(r.debAptRemoteHref, "debAptRemoteHref"), -1)  // NOTE: paths aren't escaped because Pulp uses hrefs as path parameters

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

type RemotesAptAPIRemotesDebAptListRequest struct {
	ctx context.Context
	ApiService *RemotesAptAPIService
	limit *int32
	name *string
	nameContains *string
	nameIcontains *string
	nameIn *[]string
	nameStartswith *string
	offset *int32
	ordering *[]string
	pulpHrefIn *[]string
	pulpIdIn *[]string
	pulpLabelSelect *string
	pulpLastUpdated *time.Time
	pulpLastUpdatedGt *time.Time
	pulpLastUpdatedGte *time.Time
	pulpLastUpdatedLt *time.Time
	pulpLastUpdatedLte *time.Time
	pulpLastUpdatedRange *[]time.Time
	fields *[]string
	excludeFields *[]string
}

// Number of results to return per page.
func (r RemotesAptAPIRemotesDebAptListRequest) Limit(limit int32) RemotesAptAPIRemotesDebAptListRequest {
	r.limit = &limit
	return r
}

// Filter results where name matches value
func (r RemotesAptAPIRemotesDebAptListRequest) Name(name string) RemotesAptAPIRemotesDebAptListRequest {
	r.name = &name
	return r
}

// Filter results where name contains value
func (r RemotesAptAPIRemotesDebAptListRequest) NameContains(nameContains string) RemotesAptAPIRemotesDebAptListRequest {
	r.nameContains = &nameContains
	return r
}

// Filter results where name contains value
func (r RemotesAptAPIRemotesDebAptListRequest) NameIcontains(nameIcontains string) RemotesAptAPIRemotesDebAptListRequest {
	r.nameIcontains = &nameIcontains
	return r
}

// Filter results where name is in a comma-separated list of values
func (r RemotesAptAPIRemotesDebAptListRequest) NameIn(nameIn []string) RemotesAptAPIRemotesDebAptListRequest {
	r.nameIn = &nameIn
	return r
}

// Filter results where name starts with value
func (r RemotesAptAPIRemotesDebAptListRequest) NameStartswith(nameStartswith string) RemotesAptAPIRemotesDebAptListRequest {
	r.nameStartswith = &nameStartswith
	return r
}

// The initial index from which to return the results.
func (r RemotesAptAPIRemotesDebAptListRequest) Offset(offset int32) RemotesAptAPIRemotesDebAptListRequest {
	r.offset = &offset
	return r
}

// Ordering  * &#x60;pulp_id&#x60; - Pulp id * &#x60;-pulp_id&#x60; - Pulp id (descending) * &#x60;pulp_created&#x60; - Pulp created * &#x60;-pulp_created&#x60; - Pulp created (descending) * &#x60;pulp_last_updated&#x60; - Pulp last updated * &#x60;-pulp_last_updated&#x60; - Pulp last updated (descending) * &#x60;pulp_type&#x60; - Pulp type * &#x60;-pulp_type&#x60; - Pulp type (descending) * &#x60;name&#x60; - Name * &#x60;-name&#x60; - Name (descending) * &#x60;pulp_labels&#x60; - Pulp labels * &#x60;-pulp_labels&#x60; - Pulp labels (descending) * &#x60;url&#x60; - Url * &#x60;-url&#x60; - Url (descending) * &#x60;ca_cert&#x60; - Ca cert * &#x60;-ca_cert&#x60; - Ca cert (descending) * &#x60;client_cert&#x60; - Client cert * &#x60;-client_cert&#x60; - Client cert (descending) * &#x60;client_key&#x60; - Client key * &#x60;-client_key&#x60; - Client key (descending) * &#x60;tls_validation&#x60; - Tls validation * &#x60;-tls_validation&#x60; - Tls validation (descending) * &#x60;username&#x60; - Username * &#x60;-username&#x60; - Username (descending) * &#x60;password&#x60; - Password * &#x60;-password&#x60; - Password (descending) * &#x60;proxy_url&#x60; - Proxy url * &#x60;-proxy_url&#x60; - Proxy url (descending) * &#x60;proxy_username&#x60; - Proxy username * &#x60;-proxy_username&#x60; - Proxy username (descending) * &#x60;proxy_password&#x60; - Proxy password * &#x60;-proxy_password&#x60; - Proxy password (descending) * &#x60;download_concurrency&#x60; - Download concurrency * &#x60;-download_concurrency&#x60; - Download concurrency (descending) * &#x60;max_retries&#x60; - Max retries * &#x60;-max_retries&#x60; - Max retries (descending) * &#x60;policy&#x60; - Policy * &#x60;-policy&#x60; - Policy (descending) * &#x60;total_timeout&#x60; - Total timeout * &#x60;-total_timeout&#x60; - Total timeout (descending) * &#x60;connect_timeout&#x60; - Connect timeout * &#x60;-connect_timeout&#x60; - Connect timeout (descending) * &#x60;sock_connect_timeout&#x60; - Sock connect timeout * &#x60;-sock_connect_timeout&#x60; - Sock connect timeout (descending) * &#x60;sock_read_timeout&#x60; - Sock read timeout * &#x60;-sock_read_timeout&#x60; - Sock read timeout (descending) * &#x60;headers&#x60; - Headers * &#x60;-headers&#x60; - Headers (descending) * &#x60;rate_limit&#x60; - Rate limit * &#x60;-rate_limit&#x60; - Rate limit (descending) * &#x60;pk&#x60; - Pk * &#x60;-pk&#x60; - Pk (descending)
func (r RemotesAptAPIRemotesDebAptListRequest) Ordering(ordering []string) RemotesAptAPIRemotesDebAptListRequest {
	r.ordering = &ordering
	return r
}

// Multiple values may be separated by commas.
func (r RemotesAptAPIRemotesDebAptListRequest) PulpHrefIn(pulpHrefIn []string) RemotesAptAPIRemotesDebAptListRequest {
	r.pulpHrefIn = &pulpHrefIn
	return r
}

// Multiple values may be separated by commas.
func (r RemotesAptAPIRemotesDebAptListRequest) PulpIdIn(pulpIdIn []string) RemotesAptAPIRemotesDebAptListRequest {
	r.pulpIdIn = &pulpIdIn
	return r
}

// Filter labels by search string
func (r RemotesAptAPIRemotesDebAptListRequest) PulpLabelSelect(pulpLabelSelect string) RemotesAptAPIRemotesDebAptListRequest {
	r.pulpLabelSelect = &pulpLabelSelect
	return r
}

// Filter results where pulp_last_updated matches value
func (r RemotesAptAPIRemotesDebAptListRequest) PulpLastUpdated(pulpLastUpdated time.Time) RemotesAptAPIRemotesDebAptListRequest {
	r.pulpLastUpdated = &pulpLastUpdated
	return r
}

// Filter results where pulp_last_updated is greater than value
func (r RemotesAptAPIRemotesDebAptListRequest) PulpLastUpdatedGt(pulpLastUpdatedGt time.Time) RemotesAptAPIRemotesDebAptListRequest {
	r.pulpLastUpdatedGt = &pulpLastUpdatedGt
	return r
}

// Filter results where pulp_last_updated is greater than or equal to value
func (r RemotesAptAPIRemotesDebAptListRequest) PulpLastUpdatedGte(pulpLastUpdatedGte time.Time) RemotesAptAPIRemotesDebAptListRequest {
	r.pulpLastUpdatedGte = &pulpLastUpdatedGte
	return r
}

// Filter results where pulp_last_updated is less than value
func (r RemotesAptAPIRemotesDebAptListRequest) PulpLastUpdatedLt(pulpLastUpdatedLt time.Time) RemotesAptAPIRemotesDebAptListRequest {
	r.pulpLastUpdatedLt = &pulpLastUpdatedLt
	return r
}

// Filter results where pulp_last_updated is less than or equal to value
func (r RemotesAptAPIRemotesDebAptListRequest) PulpLastUpdatedLte(pulpLastUpdatedLte time.Time) RemotesAptAPIRemotesDebAptListRequest {
	r.pulpLastUpdatedLte = &pulpLastUpdatedLte
	return r
}

// Filter results where pulp_last_updated is between two comma separated values
func (r RemotesAptAPIRemotesDebAptListRequest) PulpLastUpdatedRange(pulpLastUpdatedRange []time.Time) RemotesAptAPIRemotesDebAptListRequest {
	r.pulpLastUpdatedRange = &pulpLastUpdatedRange
	return r
}

// A list of fields to include in the response.
func (r RemotesAptAPIRemotesDebAptListRequest) Fields(fields []string) RemotesAptAPIRemotesDebAptListRequest {
	r.fields = &fields
	return r
}

// A list of fields to exclude from the response.
func (r RemotesAptAPIRemotesDebAptListRequest) ExcludeFields(excludeFields []string) RemotesAptAPIRemotesDebAptListRequest {
	r.excludeFields = &excludeFields
	return r
}

func (r RemotesAptAPIRemotesDebAptListRequest) Execute() (*PaginateddebAptRemoteResponseList, *http.Response, error) {
	return r.ApiService.RemotesDebAptListExecute(r)
}

/*
RemotesDebAptList List apt remotes

An AptRemote represents an external APT repository content source.

It contains the location of the upstream APT repository, as well as the user options that are
applied when using the remote to synchronize the upstream repository to Pulp.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return RemotesAptAPIRemotesDebAptListRequest
*/
func (a *RemotesAptAPIService) RemotesDebAptList(ctx context.Context) RemotesAptAPIRemotesDebAptListRequest {
	return RemotesAptAPIRemotesDebAptListRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return PaginateddebAptRemoteResponseList
func (a *RemotesAptAPIService) RemotesDebAptListExecute(r RemotesAptAPIRemotesDebAptListRequest) (*PaginateddebAptRemoteResponseList, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PaginateddebAptRemoteResponseList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RemotesAptAPIService.RemotesDebAptList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pulp/api/v3/remotes/deb/apt/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "")
	}
	if r.name != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "name", r.name, "")
	}
	if r.nameContains != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "name__contains", r.nameContains, "")
	}
	if r.nameIcontains != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "name__icontains", r.nameIcontains, "")
	}
	if r.nameIn != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "name__in", r.nameIn, "csv")
	}
	if r.nameStartswith != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "name__startswith", r.nameStartswith, "")
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
	if r.pulpLabelSelect != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pulp_label_select", r.pulpLabelSelect, "")
	}
	if r.pulpLastUpdated != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pulp_last_updated", r.pulpLastUpdated, "")
	}
	if r.pulpLastUpdatedGt != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pulp_last_updated__gt", r.pulpLastUpdatedGt, "")
	}
	if r.pulpLastUpdatedGte != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pulp_last_updated__gte", r.pulpLastUpdatedGte, "")
	}
	if r.pulpLastUpdatedLt != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pulp_last_updated__lt", r.pulpLastUpdatedLt, "")
	}
	if r.pulpLastUpdatedLte != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pulp_last_updated__lte", r.pulpLastUpdatedLte, "")
	}
	if r.pulpLastUpdatedRange != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pulp_last_updated__range", r.pulpLastUpdatedRange, "csv")
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

type RemotesAptAPIRemotesDebAptPartialUpdateRequest struct {
	ctx context.Context
	ApiService *RemotesAptAPIService
	debAptRemoteHref string
	patcheddebAptRemote *PatcheddebAptRemote
}

func (r RemotesAptAPIRemotesDebAptPartialUpdateRequest) PatcheddebAptRemote(patcheddebAptRemote PatcheddebAptRemote) RemotesAptAPIRemotesDebAptPartialUpdateRequest {
	r.patcheddebAptRemote = &patcheddebAptRemote
	return r
}

func (r RemotesAptAPIRemotesDebAptPartialUpdateRequest) Execute() (*AsyncOperationResponse, *http.Response, error) {
	return r.ApiService.RemotesDebAptPartialUpdateExecute(r)
}

/*
RemotesDebAptPartialUpdate Update an apt remote

Trigger an asynchronous partial update task

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param debAptRemoteHref
 @return RemotesAptAPIRemotesDebAptPartialUpdateRequest
*/
func (a *RemotesAptAPIService) RemotesDebAptPartialUpdate(ctx context.Context, debAptRemoteHref string) RemotesAptAPIRemotesDebAptPartialUpdateRequest {
	return RemotesAptAPIRemotesDebAptPartialUpdateRequest{
		ApiService: a,
		ctx: ctx,
		debAptRemoteHref: debAptRemoteHref,
	}
}

// Execute executes the request
//  @return AsyncOperationResponse
func (a *RemotesAptAPIService) RemotesDebAptPartialUpdateExecute(r RemotesAptAPIRemotesDebAptPartialUpdateRequest) (*AsyncOperationResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AsyncOperationResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RemotesAptAPIService.RemotesDebAptPartialUpdate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "{deb_apt_remote_href}"
	localVarPath = strings.Replace(localVarPath, "{"+"deb_apt_remote_href"+"}", parameterValueToString(r.debAptRemoteHref, "debAptRemoteHref"), -1)  // NOTE: paths aren't escaped because Pulp uses hrefs as path parameters

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.patcheddebAptRemote == nil {
		return localVarReturnValue, nil, reportError("patcheddebAptRemote is required and must be specified")
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
	localVarPostBody = r.patcheddebAptRemote
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

type RemotesAptAPIRemotesDebAptReadRequest struct {
	ctx context.Context
	ApiService *RemotesAptAPIService
	debAptRemoteHref string
	fields *[]string
	excludeFields *[]string
}

// A list of fields to include in the response.
func (r RemotesAptAPIRemotesDebAptReadRequest) Fields(fields []string) RemotesAptAPIRemotesDebAptReadRequest {
	r.fields = &fields
	return r
}

// A list of fields to exclude from the response.
func (r RemotesAptAPIRemotesDebAptReadRequest) ExcludeFields(excludeFields []string) RemotesAptAPIRemotesDebAptReadRequest {
	r.excludeFields = &excludeFields
	return r
}

func (r RemotesAptAPIRemotesDebAptReadRequest) Execute() (*DebAptRemoteResponse, *http.Response, error) {
	return r.ApiService.RemotesDebAptReadExecute(r)
}

/*
RemotesDebAptRead Inspect an apt remote

An AptRemote represents an external APT repository content source.

It contains the location of the upstream APT repository, as well as the user options that are
applied when using the remote to synchronize the upstream repository to Pulp.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param debAptRemoteHref
 @return RemotesAptAPIRemotesDebAptReadRequest
*/
func (a *RemotesAptAPIService) RemotesDebAptRead(ctx context.Context, debAptRemoteHref string) RemotesAptAPIRemotesDebAptReadRequest {
	return RemotesAptAPIRemotesDebAptReadRequest{
		ApiService: a,
		ctx: ctx,
		debAptRemoteHref: debAptRemoteHref,
	}
}

// Execute executes the request
//  @return DebAptRemoteResponse
func (a *RemotesAptAPIService) RemotesDebAptReadExecute(r RemotesAptAPIRemotesDebAptReadRequest) (*DebAptRemoteResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *DebAptRemoteResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RemotesAptAPIService.RemotesDebAptRead")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "{deb_apt_remote_href}"
	localVarPath = strings.Replace(localVarPath, "{"+"deb_apt_remote_href"+"}", parameterValueToString(r.debAptRemoteHref, "debAptRemoteHref"), -1)  // NOTE: paths aren't escaped because Pulp uses hrefs as path parameters

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

type RemotesAptAPIRemotesDebAptUpdateRequest struct {
	ctx context.Context
	ApiService *RemotesAptAPIService
	debAptRemoteHref string
	debAptRemote *DebAptRemote
}

func (r RemotesAptAPIRemotesDebAptUpdateRequest) DebAptRemote(debAptRemote DebAptRemote) RemotesAptAPIRemotesDebAptUpdateRequest {
	r.debAptRemote = &debAptRemote
	return r
}

func (r RemotesAptAPIRemotesDebAptUpdateRequest) Execute() (*AsyncOperationResponse, *http.Response, error) {
	return r.ApiService.RemotesDebAptUpdateExecute(r)
}

/*
RemotesDebAptUpdate Update an apt remote

Trigger an asynchronous update task

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param debAptRemoteHref
 @return RemotesAptAPIRemotesDebAptUpdateRequest
*/
func (a *RemotesAptAPIService) RemotesDebAptUpdate(ctx context.Context, debAptRemoteHref string) RemotesAptAPIRemotesDebAptUpdateRequest {
	return RemotesAptAPIRemotesDebAptUpdateRequest{
		ApiService: a,
		ctx: ctx,
		debAptRemoteHref: debAptRemoteHref,
	}
}

// Execute executes the request
//  @return AsyncOperationResponse
func (a *RemotesAptAPIService) RemotesDebAptUpdateExecute(r RemotesAptAPIRemotesDebAptUpdateRequest) (*AsyncOperationResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AsyncOperationResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RemotesAptAPIService.RemotesDebAptUpdate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "{deb_apt_remote_href}"
	localVarPath = strings.Replace(localVarPath, "{"+"deb_apt_remote_href"+"}", parameterValueToString(r.debAptRemoteHref, "debAptRemoteHref"), -1)  // NOTE: paths aren't escaped because Pulp uses hrefs as path parameters

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.debAptRemote == nil {
		return localVarReturnValue, nil, reportError("debAptRemote is required and must be specified")
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
	localVarPostBody = r.debAptRemote
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
