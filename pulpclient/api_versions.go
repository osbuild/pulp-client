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


// VersionsAPIService VersionsAPI service
type VersionsAPIService service

type VersionsAPIApiV1RolesVersionsListRequest struct {
	ctx context.Context
	ApiService *VersionsAPIService
	ansibleRoleHref string
	limit *int32
	offset *int32
	fields *[]string
	excludeFields *[]string
}

// Number of results to return per page.
func (r VersionsAPIApiV1RolesVersionsListRequest) Limit(limit int32) VersionsAPIApiV1RolesVersionsListRequest {
	r.limit = &limit
	return r
}

// The initial index from which to return the results.
func (r VersionsAPIApiV1RolesVersionsListRequest) Offset(offset int32) VersionsAPIApiV1RolesVersionsListRequest {
	r.offset = &offset
	return r
}

// A list of fields to include in the response.
func (r VersionsAPIApiV1RolesVersionsListRequest) Fields(fields []string) VersionsAPIApiV1RolesVersionsListRequest {
	r.fields = &fields
	return r
}

// A list of fields to exclude from the response.
func (r VersionsAPIApiV1RolesVersionsListRequest) ExcludeFields(excludeFields []string) VersionsAPIApiV1RolesVersionsListRequest {
	r.excludeFields = &excludeFields
	return r
}

func (r VersionsAPIApiV1RolesVersionsListRequest) Execute() (*PaginatedGalaxyRoleVersionResponseList, *http.Response, error) {
	return r.ApiService.ApiV1RolesVersionsListExecute(r)
}

/*
ApiV1RolesVersionsList Method for ApiV1RolesVersionsList

APIView for Role Versions.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param ansibleRoleHref
 @return VersionsAPIApiV1RolesVersionsListRequest
*/
func (a *VersionsAPIService) ApiV1RolesVersionsList(ctx context.Context, ansibleRoleHref string) VersionsAPIApiV1RolesVersionsListRequest {
	return VersionsAPIApiV1RolesVersionsListRequest{
		ApiService: a,
		ctx: ctx,
		ansibleRoleHref: ansibleRoleHref,
	}
}

// Execute executes the request
//  @return PaginatedGalaxyRoleVersionResponseList
func (a *VersionsAPIService) ApiV1RolesVersionsListExecute(r VersionsAPIApiV1RolesVersionsListRequest) (*PaginatedGalaxyRoleVersionResponseList, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PaginatedGalaxyRoleVersionResponseList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "VersionsAPIService.ApiV1RolesVersionsList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "{ansible_role_href}versions/"
	localVarPath = strings.Replace(localVarPath, "{"+"ansible_role_href"+"}", parameterValueToString(r.ansibleRoleHref, "ansibleRoleHref"), -1)  // NOTE: paths aren't escaped because Pulp uses hrefs as path parameters

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "")
	}
	if r.offset != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "offset", r.offset, "")
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

type VersionsAPIApiV2CollectionVersionsListRequest struct {
	ctx context.Context
	ApiService *VersionsAPIService
	ansibleCollectionVersionHref string
	page *int32
	fields *[]string
	excludeFields *[]string
}

// A page number within the paginated result set.
func (r VersionsAPIApiV2CollectionVersionsListRequest) Page(page int32) VersionsAPIApiV2CollectionVersionsListRequest {
	r.page = &page
	return r
}

// A list of fields to include in the response.
func (r VersionsAPIApiV2CollectionVersionsListRequest) Fields(fields []string) VersionsAPIApiV2CollectionVersionsListRequest {
	r.fields = &fields
	return r
}

// A list of fields to exclude from the response.
func (r VersionsAPIApiV2CollectionVersionsListRequest) ExcludeFields(excludeFields []string) VersionsAPIApiV2CollectionVersionsListRequest {
	r.excludeFields = &excludeFields
	return r
}

func (r VersionsAPIApiV2CollectionVersionsListRequest) Execute() (*PaginatedGalaxyCollectionVersionResponseList, *http.Response, error) {
	return r.ApiService.ApiV2CollectionVersionsListExecute(r)
}

/*
ApiV2CollectionVersionsList Method for ApiV2CollectionVersionsList

APIView for Collections by namespace/name.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param ansibleCollectionVersionHref
 @return VersionsAPIApiV2CollectionVersionsListRequest
*/
func (a *VersionsAPIService) ApiV2CollectionVersionsList(ctx context.Context, ansibleCollectionVersionHref string) VersionsAPIApiV2CollectionVersionsListRequest {
	return VersionsAPIApiV2CollectionVersionsListRequest{
		ApiService: a,
		ctx: ctx,
		ansibleCollectionVersionHref: ansibleCollectionVersionHref,
	}
}

// Execute executes the request
//  @return PaginatedGalaxyCollectionVersionResponseList
func (a *VersionsAPIService) ApiV2CollectionVersionsListExecute(r VersionsAPIApiV2CollectionVersionsListRequest) (*PaginatedGalaxyCollectionVersionResponseList, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PaginatedGalaxyCollectionVersionResponseList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "VersionsAPIService.ApiV2CollectionVersionsList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "{ansible_collection_version_href}versions/"
	localVarPath = strings.Replace(localVarPath, "{"+"ansible_collection_version_href"+"}", parameterValueToString(r.ansibleCollectionVersionHref, "ansibleCollectionVersionHref"), -1)  // NOTE: paths aren't escaped because Pulp uses hrefs as path parameters

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.page != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page, "")
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
