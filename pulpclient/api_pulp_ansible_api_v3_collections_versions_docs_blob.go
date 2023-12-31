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


// PulpAnsibleApiV3CollectionsVersionsDocsBlobAPIService PulpAnsibleApiV3CollectionsVersionsDocsBlobAPI service
type PulpAnsibleApiV3CollectionsVersionsDocsBlobAPIService service

type PulpAnsibleApiV3CollectionsVersionsDocsBlobAPIPulpAnsibleGalaxyApiV3CollectionsVersionsDocsBlobReadRequest struct {
	ctx context.Context
	ApiService *PulpAnsibleApiV3CollectionsVersionsDocsBlobAPIService
	name string
	namespace string
	path string
	version string
	fields *[]string
	excludeFields *[]string
}

// A list of fields to include in the response.
func (r PulpAnsibleApiV3CollectionsVersionsDocsBlobAPIPulpAnsibleGalaxyApiV3CollectionsVersionsDocsBlobReadRequest) Fields(fields []string) PulpAnsibleApiV3CollectionsVersionsDocsBlobAPIPulpAnsibleGalaxyApiV3CollectionsVersionsDocsBlobReadRequest {
	r.fields = &fields
	return r
}

// A list of fields to exclude from the response.
func (r PulpAnsibleApiV3CollectionsVersionsDocsBlobAPIPulpAnsibleGalaxyApiV3CollectionsVersionsDocsBlobReadRequest) ExcludeFields(excludeFields []string) PulpAnsibleApiV3CollectionsVersionsDocsBlobAPIPulpAnsibleGalaxyApiV3CollectionsVersionsDocsBlobReadRequest {
	r.excludeFields = &excludeFields
	return r
}

func (r PulpAnsibleApiV3CollectionsVersionsDocsBlobAPIPulpAnsibleGalaxyApiV3CollectionsVersionsDocsBlobReadRequest) Execute() (*CollectionVersionDocsResponse, *http.Response, error) {
	return r.ApiService.PulpAnsibleGalaxyApiV3CollectionsVersionsDocsBlobReadExecute(r)
}

/*
PulpAnsibleGalaxyApiV3CollectionsVersionsDocsBlobRead Method for PulpAnsibleGalaxyApiV3CollectionsVersionsDocsBlobRead

Legacy v3 endpoint.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param name
 @param namespace
 @param path
 @param version
 @return PulpAnsibleApiV3CollectionsVersionsDocsBlobAPIPulpAnsibleGalaxyApiV3CollectionsVersionsDocsBlobReadRequest

Deprecated
*/
func (a *PulpAnsibleApiV3CollectionsVersionsDocsBlobAPIService) PulpAnsibleGalaxyApiV3CollectionsVersionsDocsBlobRead(ctx context.Context, name string, namespace string, path string, version string) PulpAnsibleApiV3CollectionsVersionsDocsBlobAPIPulpAnsibleGalaxyApiV3CollectionsVersionsDocsBlobReadRequest {
	return PulpAnsibleApiV3CollectionsVersionsDocsBlobAPIPulpAnsibleGalaxyApiV3CollectionsVersionsDocsBlobReadRequest{
		ApiService: a,
		ctx: ctx,
		name: name,
		namespace: namespace,
		path: path,
		version: version,
	}
}

// Execute executes the request
//  @return CollectionVersionDocsResponse
// Deprecated
func (a *PulpAnsibleApiV3CollectionsVersionsDocsBlobAPIService) PulpAnsibleGalaxyApiV3CollectionsVersionsDocsBlobReadExecute(r PulpAnsibleApiV3CollectionsVersionsDocsBlobAPIPulpAnsibleGalaxyApiV3CollectionsVersionsDocsBlobReadRequest) (*CollectionVersionDocsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CollectionVersionDocsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PulpAnsibleApiV3CollectionsVersionsDocsBlobAPIService.PulpAnsibleGalaxyApiV3CollectionsVersionsDocsBlobRead")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pulp_ansible/galaxy/{path}/api/v3/collections/{namespace}/{name}/versions/{version}/docs-blob/"
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
