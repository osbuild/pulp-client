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


// ContentSignaturesAPIService ContentSignaturesAPI service
type ContentSignaturesAPIService service

type ContentSignaturesAPIContentContainerSignaturesListRequest struct {
	ctx context.Context
	ApiService *ContentSignaturesAPIService
	digest *string
	digestIn *[]string
	keyId *string
	keyIdIn *[]string
	limit *int32
	manifest *[]string
	name *string
	nameContains *string
	nameIcontains *string
	nameIn *[]string
	nameStartswith *string
	offset *int32
	ordering *[]string
	pulpHrefIn *[]string
	pulpIdIn *[]string
	repositoryVersion *string
	repositoryVersionAdded *string
	repositoryVersionRemoved *string
	fields *[]string
	excludeFields *[]string
}

// Filter results where digest matches value
func (r ContentSignaturesAPIContentContainerSignaturesListRequest) Digest(digest string) ContentSignaturesAPIContentContainerSignaturesListRequest {
	r.digest = &digest
	return r
}

// Filter results where digest is in a comma-separated list of values
func (r ContentSignaturesAPIContentContainerSignaturesListRequest) DigestIn(digestIn []string) ContentSignaturesAPIContentContainerSignaturesListRequest {
	r.digestIn = &digestIn
	return r
}

// Filter results where key_id matches value
func (r ContentSignaturesAPIContentContainerSignaturesListRequest) KeyId(keyId string) ContentSignaturesAPIContentContainerSignaturesListRequest {
	r.keyId = &keyId
	return r
}

// Filter results where key_id is in a comma-separated list of values
func (r ContentSignaturesAPIContentContainerSignaturesListRequest) KeyIdIn(keyIdIn []string) ContentSignaturesAPIContentContainerSignaturesListRequest {
	r.keyIdIn = &keyIdIn
	return r
}

// Number of results to return per page.
func (r ContentSignaturesAPIContentContainerSignaturesListRequest) Limit(limit int32) ContentSignaturesAPIContentContainerSignaturesListRequest {
	r.limit = &limit
	return r
}

// Multiple values may be separated by commas.
func (r ContentSignaturesAPIContentContainerSignaturesListRequest) Manifest(manifest []string) ContentSignaturesAPIContentContainerSignaturesListRequest {
	r.manifest = &manifest
	return r
}

// Filter results where name matches value
func (r ContentSignaturesAPIContentContainerSignaturesListRequest) Name(name string) ContentSignaturesAPIContentContainerSignaturesListRequest {
	r.name = &name
	return r
}

// Filter results where name contains value
func (r ContentSignaturesAPIContentContainerSignaturesListRequest) NameContains(nameContains string) ContentSignaturesAPIContentContainerSignaturesListRequest {
	r.nameContains = &nameContains
	return r
}

// Filter results where name contains value
func (r ContentSignaturesAPIContentContainerSignaturesListRequest) NameIcontains(nameIcontains string) ContentSignaturesAPIContentContainerSignaturesListRequest {
	r.nameIcontains = &nameIcontains
	return r
}

// Filter results where name is in a comma-separated list of values
func (r ContentSignaturesAPIContentContainerSignaturesListRequest) NameIn(nameIn []string) ContentSignaturesAPIContentContainerSignaturesListRequest {
	r.nameIn = &nameIn
	return r
}

// Filter results where name starts with value
func (r ContentSignaturesAPIContentContainerSignaturesListRequest) NameStartswith(nameStartswith string) ContentSignaturesAPIContentContainerSignaturesListRequest {
	r.nameStartswith = &nameStartswith
	return r
}

// The initial index from which to return the results.
func (r ContentSignaturesAPIContentContainerSignaturesListRequest) Offset(offset int32) ContentSignaturesAPIContentContainerSignaturesListRequest {
	r.offset = &offset
	return r
}

// Ordering  * &#x60;pulp_id&#x60; - Pulp id * &#x60;-pulp_id&#x60; - Pulp id (descending) * &#x60;pulp_created&#x60; - Pulp created * &#x60;-pulp_created&#x60; - Pulp created (descending) * &#x60;pulp_last_updated&#x60; - Pulp last updated * &#x60;-pulp_last_updated&#x60; - Pulp last updated (descending) * &#x60;pulp_type&#x60; - Pulp type * &#x60;-pulp_type&#x60; - Pulp type (descending) * &#x60;upstream_id&#x60; - Upstream id * &#x60;-upstream_id&#x60; - Upstream id (descending) * &#x60;timestamp_of_interest&#x60; - Timestamp of interest * &#x60;-timestamp_of_interest&#x60; - Timestamp of interest (descending) * &#x60;name&#x60; - Name * &#x60;-name&#x60; - Name (descending) * &#x60;digest&#x60; - Digest * &#x60;-digest&#x60; - Digest (descending) * &#x60;type&#x60; - Type * &#x60;-type&#x60; - Type (descending) * &#x60;key_id&#x60; - Key id * &#x60;-key_id&#x60; - Key id (descending) * &#x60;timestamp&#x60; - Timestamp * &#x60;-timestamp&#x60; - Timestamp (descending) * &#x60;creator&#x60; - Creator * &#x60;-creator&#x60; - Creator (descending) * &#x60;data&#x60; - Data * &#x60;-data&#x60; - Data (descending) * &#x60;pk&#x60; - Pk * &#x60;-pk&#x60; - Pk (descending)
func (r ContentSignaturesAPIContentContainerSignaturesListRequest) Ordering(ordering []string) ContentSignaturesAPIContentContainerSignaturesListRequest {
	r.ordering = &ordering
	return r
}

// Multiple values may be separated by commas.
func (r ContentSignaturesAPIContentContainerSignaturesListRequest) PulpHrefIn(pulpHrefIn []string) ContentSignaturesAPIContentContainerSignaturesListRequest {
	r.pulpHrefIn = &pulpHrefIn
	return r
}

// Multiple values may be separated by commas.
func (r ContentSignaturesAPIContentContainerSignaturesListRequest) PulpIdIn(pulpIdIn []string) ContentSignaturesAPIContentContainerSignaturesListRequest {
	r.pulpIdIn = &pulpIdIn
	return r
}

// Repository Version referenced by HREF
func (r ContentSignaturesAPIContentContainerSignaturesListRequest) RepositoryVersion(repositoryVersion string) ContentSignaturesAPIContentContainerSignaturesListRequest {
	r.repositoryVersion = &repositoryVersion
	return r
}

// Repository Version referenced by HREF
func (r ContentSignaturesAPIContentContainerSignaturesListRequest) RepositoryVersionAdded(repositoryVersionAdded string) ContentSignaturesAPIContentContainerSignaturesListRequest {
	r.repositoryVersionAdded = &repositoryVersionAdded
	return r
}

// Repository Version referenced by HREF
func (r ContentSignaturesAPIContentContainerSignaturesListRequest) RepositoryVersionRemoved(repositoryVersionRemoved string) ContentSignaturesAPIContentContainerSignaturesListRequest {
	r.repositoryVersionRemoved = &repositoryVersionRemoved
	return r
}

// A list of fields to include in the response.
func (r ContentSignaturesAPIContentContainerSignaturesListRequest) Fields(fields []string) ContentSignaturesAPIContentContainerSignaturesListRequest {
	r.fields = &fields
	return r
}

// A list of fields to exclude from the response.
func (r ContentSignaturesAPIContentContainerSignaturesListRequest) ExcludeFields(excludeFields []string) ContentSignaturesAPIContentContainerSignaturesListRequest {
	r.excludeFields = &excludeFields
	return r
}

func (r ContentSignaturesAPIContentContainerSignaturesListRequest) Execute() (*PaginatedcontainerManifestSignatureResponseList, *http.Response, error) {
	return r.ApiService.ContentContainerSignaturesListExecute(r)
}

/*
ContentContainerSignaturesList List manifest signatures

ViewSet for image signatures.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ContentSignaturesAPIContentContainerSignaturesListRequest
*/
func (a *ContentSignaturesAPIService) ContentContainerSignaturesList(ctx context.Context) ContentSignaturesAPIContentContainerSignaturesListRequest {
	return ContentSignaturesAPIContentContainerSignaturesListRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return PaginatedcontainerManifestSignatureResponseList
func (a *ContentSignaturesAPIService) ContentContainerSignaturesListExecute(r ContentSignaturesAPIContentContainerSignaturesListRequest) (*PaginatedcontainerManifestSignatureResponseList, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PaginatedcontainerManifestSignatureResponseList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ContentSignaturesAPIService.ContentContainerSignaturesList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pulp/api/v3/content/container/signatures/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.digest != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "digest", r.digest, "")
	}
	if r.digestIn != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "digest__in", r.digestIn, "csv")
	}
	if r.keyId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "key_id", r.keyId, "")
	}
	if r.keyIdIn != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "key_id__in", r.keyIdIn, "csv")
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "")
	}
	if r.manifest != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "manifest", r.manifest, "csv")
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
	if r.repositoryVersion != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "repository_version", r.repositoryVersion, "")
	}
	if r.repositoryVersionAdded != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "repository_version_added", r.repositoryVersionAdded, "")
	}
	if r.repositoryVersionRemoved != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "repository_version_removed", r.repositoryVersionRemoved, "")
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

type ContentSignaturesAPIContentContainerSignaturesReadRequest struct {
	ctx context.Context
	ApiService *ContentSignaturesAPIService
	containerManifestSignatureHref string
	fields *[]string
	excludeFields *[]string
}

// A list of fields to include in the response.
func (r ContentSignaturesAPIContentContainerSignaturesReadRequest) Fields(fields []string) ContentSignaturesAPIContentContainerSignaturesReadRequest {
	r.fields = &fields
	return r
}

// A list of fields to exclude from the response.
func (r ContentSignaturesAPIContentContainerSignaturesReadRequest) ExcludeFields(excludeFields []string) ContentSignaturesAPIContentContainerSignaturesReadRequest {
	r.excludeFields = &excludeFields
	return r
}

func (r ContentSignaturesAPIContentContainerSignaturesReadRequest) Execute() (*ContainerManifestSignatureResponse, *http.Response, error) {
	return r.ApiService.ContentContainerSignaturesReadExecute(r)
}

/*
ContentContainerSignaturesRead Inspect a manifest signature

ViewSet for image signatures.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param containerManifestSignatureHref
 @return ContentSignaturesAPIContentContainerSignaturesReadRequest
*/
func (a *ContentSignaturesAPIService) ContentContainerSignaturesRead(ctx context.Context, containerManifestSignatureHref string) ContentSignaturesAPIContentContainerSignaturesReadRequest {
	return ContentSignaturesAPIContentContainerSignaturesReadRequest{
		ApiService: a,
		ctx: ctx,
		containerManifestSignatureHref: containerManifestSignatureHref,
	}
}

// Execute executes the request
//  @return ContainerManifestSignatureResponse
func (a *ContentSignaturesAPIService) ContentContainerSignaturesReadExecute(r ContentSignaturesAPIContentContainerSignaturesReadRequest) (*ContainerManifestSignatureResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ContainerManifestSignatureResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ContentSignaturesAPIService.ContentContainerSignaturesRead")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "{container_manifest_signature_href}"
	localVarPath = strings.Replace(localVarPath, "{"+"container_manifest_signature_href"+"}", parameterValueToString(r.containerManifestSignatureHref, "containerManifestSignatureHref"), -1)  // NOTE: paths aren't escaped because Pulp uses hrefs as path parameters

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