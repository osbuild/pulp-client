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


// PulpAnsibleApiV3APIService PulpAnsibleApiV3API service
type PulpAnsibleApiV3APIService service

type PulpAnsibleApiV3APIPulpAnsibleGalaxyApiV3ReadRequest struct {
	ctx context.Context
	ApiService *PulpAnsibleApiV3APIService
	path string
	fields *[]string
	excludeFields *[]string
}

// A list of fields to include in the response.
func (r PulpAnsibleApiV3APIPulpAnsibleGalaxyApiV3ReadRequest) Fields(fields []string) PulpAnsibleApiV3APIPulpAnsibleGalaxyApiV3ReadRequest {
	r.fields = &fields
	return r
}

// A list of fields to exclude from the response.
func (r PulpAnsibleApiV3APIPulpAnsibleGalaxyApiV3ReadRequest) ExcludeFields(excludeFields []string) PulpAnsibleApiV3APIPulpAnsibleGalaxyApiV3ReadRequest {
	r.excludeFields = &excludeFields
	return r
}

func (r PulpAnsibleApiV3APIPulpAnsibleGalaxyApiV3ReadRequest) Execute() (*RepoMetadataResponse, *http.Response, error) {
	return r.ApiService.PulpAnsibleGalaxyApiV3ReadExecute(r)
}

/*
PulpAnsibleGalaxyApiV3Read Method for PulpAnsibleGalaxyApiV3Read

Legacy v3 endpoint.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param path
 @return PulpAnsibleApiV3APIPulpAnsibleGalaxyApiV3ReadRequest

Deprecated
*/
func (a *PulpAnsibleApiV3APIService) PulpAnsibleGalaxyApiV3Read(ctx context.Context, path string) PulpAnsibleApiV3APIPulpAnsibleGalaxyApiV3ReadRequest {
	return PulpAnsibleApiV3APIPulpAnsibleGalaxyApiV3ReadRequest{
		ApiService: a,
		ctx: ctx,
		path: path,
	}
}

// Execute executes the request
//  @return RepoMetadataResponse
// Deprecated
func (a *PulpAnsibleApiV3APIService) PulpAnsibleGalaxyApiV3ReadExecute(r PulpAnsibleApiV3APIPulpAnsibleGalaxyApiV3ReadRequest) (*RepoMetadataResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *RepoMetadataResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PulpAnsibleApiV3APIService.PulpAnsibleGalaxyApiV3Read")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pulp_ansible/galaxy/{path}/api/v3/"
	localVarPath = strings.Replace(localVarPath, "{"+"path"+"}", parameterValueToString(r.path, "path"), -1)  // NOTE: paths aren't escaped because Pulp uses hrefs as path parameters

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