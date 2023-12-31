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
	"os"
)


// PulpAnsibleArtifactsCollectionsV3APIService PulpAnsibleArtifactsCollectionsV3API service
type PulpAnsibleArtifactsCollectionsV3APIService service

type PulpAnsibleArtifactsCollectionsV3APIPulpAnsibleGalaxyApiV3ArtifactsCollectionsCreateRequest struct {
	ctx context.Context
	ApiService *PulpAnsibleArtifactsCollectionsV3APIService
	path string
	file *os.File
	sha256 *string
	expectedNamespace *string
	expectedName *string
	expectedVersion *string
}

// The Collection tarball.
func (r PulpAnsibleArtifactsCollectionsV3APIPulpAnsibleGalaxyApiV3ArtifactsCollectionsCreateRequest) File(file *os.File) PulpAnsibleArtifactsCollectionsV3APIPulpAnsibleGalaxyApiV3ArtifactsCollectionsCreateRequest {
	r.file = file
	return r
}

// An optional sha256 checksum of the uploaded file.
func (r PulpAnsibleArtifactsCollectionsV3APIPulpAnsibleGalaxyApiV3ArtifactsCollectionsCreateRequest) Sha256(sha256 string) PulpAnsibleArtifactsCollectionsV3APIPulpAnsibleGalaxyApiV3ArtifactsCollectionsCreateRequest {
	r.sha256 = &sha256
	return r
}

// The expected &#39;namespace&#39; of the Collection to be verified against the metadata during import.
func (r PulpAnsibleArtifactsCollectionsV3APIPulpAnsibleGalaxyApiV3ArtifactsCollectionsCreateRequest) ExpectedNamespace(expectedNamespace string) PulpAnsibleArtifactsCollectionsV3APIPulpAnsibleGalaxyApiV3ArtifactsCollectionsCreateRequest {
	r.expectedNamespace = &expectedNamespace
	return r
}

// The expected &#39;name&#39; of the Collection to be verified against the metadata during import.
func (r PulpAnsibleArtifactsCollectionsV3APIPulpAnsibleGalaxyApiV3ArtifactsCollectionsCreateRequest) ExpectedName(expectedName string) PulpAnsibleArtifactsCollectionsV3APIPulpAnsibleGalaxyApiV3ArtifactsCollectionsCreateRequest {
	r.expectedName = &expectedName
	return r
}

// The expected version of the Collection to be verified against the metadata during import.
func (r PulpAnsibleArtifactsCollectionsV3APIPulpAnsibleGalaxyApiV3ArtifactsCollectionsCreateRequest) ExpectedVersion(expectedVersion string) PulpAnsibleArtifactsCollectionsV3APIPulpAnsibleGalaxyApiV3ArtifactsCollectionsCreateRequest {
	r.expectedVersion = &expectedVersion
	return r
}

func (r PulpAnsibleArtifactsCollectionsV3APIPulpAnsibleGalaxyApiV3ArtifactsCollectionsCreateRequest) Execute() (*AsyncOperationResponse, *http.Response, error) {
	return r.ApiService.PulpAnsibleGalaxyApiV3ArtifactsCollectionsCreateExecute(r)
}

/*
PulpAnsibleGalaxyApiV3ArtifactsCollectionsCreate Upload a collection

Create an artifact and trigger an asynchronous task to create Collection content from it.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param path
 @return PulpAnsibleArtifactsCollectionsV3APIPulpAnsibleGalaxyApiV3ArtifactsCollectionsCreateRequest

Deprecated
*/
func (a *PulpAnsibleArtifactsCollectionsV3APIService) PulpAnsibleGalaxyApiV3ArtifactsCollectionsCreate(ctx context.Context, path string) PulpAnsibleArtifactsCollectionsV3APIPulpAnsibleGalaxyApiV3ArtifactsCollectionsCreateRequest {
	return PulpAnsibleArtifactsCollectionsV3APIPulpAnsibleGalaxyApiV3ArtifactsCollectionsCreateRequest{
		ApiService: a,
		ctx: ctx,
		path: path,
	}
}

// Execute executes the request
//  @return AsyncOperationResponse
// Deprecated
func (a *PulpAnsibleArtifactsCollectionsV3APIService) PulpAnsibleGalaxyApiV3ArtifactsCollectionsCreateExecute(r PulpAnsibleArtifactsCollectionsV3APIPulpAnsibleGalaxyApiV3ArtifactsCollectionsCreateRequest) (*AsyncOperationResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AsyncOperationResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PulpAnsibleArtifactsCollectionsV3APIService.PulpAnsibleGalaxyApiV3ArtifactsCollectionsCreate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pulp_ansible/galaxy/{path}/api/v3/artifacts/collections/"
	localVarPath = strings.Replace(localVarPath, "{"+"path"+"}", parameterValueToString(r.path, "path"), -1)  // NOTE: paths aren't escaped because Pulp uses hrefs as path parameters

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.file == nil {
		return localVarReturnValue, nil, reportError("file is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"multipart/form-data", "application/x-www-form-urlencoded"}

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
	var fileLocalVarFormFileName string
	var fileLocalVarFileName     string
	var fileLocalVarFileBytes    []byte

	fileLocalVarFormFileName = "file"


	fileLocalVarFile := r.file

	if fileLocalVarFile != nil {
		fbs, _ := io.ReadAll(fileLocalVarFile)

		fileLocalVarFileBytes = fbs
		fileLocalVarFileName = fileLocalVarFile.Name()
		fileLocalVarFile.Close()
		formFiles = append(formFiles, formFile{fileBytes: fileLocalVarFileBytes, fileName: fileLocalVarFileName, formFileName: fileLocalVarFormFileName})
	}
	if r.sha256 != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "sha256", r.sha256, "")
	}
	if r.expectedNamespace != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "expected_namespace", r.expectedNamespace, "")
	}
	if r.expectedName != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "expected_name", r.expectedName, "")
	}
	if r.expectedVersion != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "expected_version", r.expectedVersion, "")
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

type PulpAnsibleArtifactsCollectionsV3APIPulpAnsibleGalaxyApiV3PluginAnsibleContentCollectionsArtifactsCreateRequest struct {
	ctx context.Context
	ApiService *PulpAnsibleArtifactsCollectionsV3APIService
	distroBasePath string
	path string
	file *os.File
	sha256 *string
	expectedNamespace *string
	expectedName *string
	expectedVersion *string
}

// The Collection tarball.
func (r PulpAnsibleArtifactsCollectionsV3APIPulpAnsibleGalaxyApiV3PluginAnsibleContentCollectionsArtifactsCreateRequest) File(file *os.File) PulpAnsibleArtifactsCollectionsV3APIPulpAnsibleGalaxyApiV3PluginAnsibleContentCollectionsArtifactsCreateRequest {
	r.file = file
	return r
}

// An optional sha256 checksum of the uploaded file.
func (r PulpAnsibleArtifactsCollectionsV3APIPulpAnsibleGalaxyApiV3PluginAnsibleContentCollectionsArtifactsCreateRequest) Sha256(sha256 string) PulpAnsibleArtifactsCollectionsV3APIPulpAnsibleGalaxyApiV3PluginAnsibleContentCollectionsArtifactsCreateRequest {
	r.sha256 = &sha256
	return r
}

// The expected &#39;namespace&#39; of the Collection to be verified against the metadata during import.
func (r PulpAnsibleArtifactsCollectionsV3APIPulpAnsibleGalaxyApiV3PluginAnsibleContentCollectionsArtifactsCreateRequest) ExpectedNamespace(expectedNamespace string) PulpAnsibleArtifactsCollectionsV3APIPulpAnsibleGalaxyApiV3PluginAnsibleContentCollectionsArtifactsCreateRequest {
	r.expectedNamespace = &expectedNamespace
	return r
}

// The expected &#39;name&#39; of the Collection to be verified against the metadata during import.
func (r PulpAnsibleArtifactsCollectionsV3APIPulpAnsibleGalaxyApiV3PluginAnsibleContentCollectionsArtifactsCreateRequest) ExpectedName(expectedName string) PulpAnsibleArtifactsCollectionsV3APIPulpAnsibleGalaxyApiV3PluginAnsibleContentCollectionsArtifactsCreateRequest {
	r.expectedName = &expectedName
	return r
}

// The expected version of the Collection to be verified against the metadata during import.
func (r PulpAnsibleArtifactsCollectionsV3APIPulpAnsibleGalaxyApiV3PluginAnsibleContentCollectionsArtifactsCreateRequest) ExpectedVersion(expectedVersion string) PulpAnsibleArtifactsCollectionsV3APIPulpAnsibleGalaxyApiV3PluginAnsibleContentCollectionsArtifactsCreateRequest {
	r.expectedVersion = &expectedVersion
	return r
}

func (r PulpAnsibleArtifactsCollectionsV3APIPulpAnsibleGalaxyApiV3PluginAnsibleContentCollectionsArtifactsCreateRequest) Execute() (*AsyncOperationResponse, *http.Response, error) {
	return r.ApiService.PulpAnsibleGalaxyApiV3PluginAnsibleContentCollectionsArtifactsCreateExecute(r)
}

/*
PulpAnsibleGalaxyApiV3PluginAnsibleContentCollectionsArtifactsCreate Upload a collection

Create an artifact and trigger an asynchronous task to create Collection content from it.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param distroBasePath
 @param path
 @return PulpAnsibleArtifactsCollectionsV3APIPulpAnsibleGalaxyApiV3PluginAnsibleContentCollectionsArtifactsCreateRequest
*/
func (a *PulpAnsibleArtifactsCollectionsV3APIService) PulpAnsibleGalaxyApiV3PluginAnsibleContentCollectionsArtifactsCreate(ctx context.Context, distroBasePath string, path string) PulpAnsibleArtifactsCollectionsV3APIPulpAnsibleGalaxyApiV3PluginAnsibleContentCollectionsArtifactsCreateRequest {
	return PulpAnsibleArtifactsCollectionsV3APIPulpAnsibleGalaxyApiV3PluginAnsibleContentCollectionsArtifactsCreateRequest{
		ApiService: a,
		ctx: ctx,
		distroBasePath: distroBasePath,
		path: path,
	}
}

// Execute executes the request
//  @return AsyncOperationResponse
func (a *PulpAnsibleArtifactsCollectionsV3APIService) PulpAnsibleGalaxyApiV3PluginAnsibleContentCollectionsArtifactsCreateExecute(r PulpAnsibleArtifactsCollectionsV3APIPulpAnsibleGalaxyApiV3PluginAnsibleContentCollectionsArtifactsCreateRequest) (*AsyncOperationResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AsyncOperationResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PulpAnsibleArtifactsCollectionsV3APIService.PulpAnsibleGalaxyApiV3PluginAnsibleContentCollectionsArtifactsCreate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pulp_ansible/galaxy/{path}/api/v3/plugin/ansible/content/{distro_base_path}/collections/artifacts/"
	localVarPath = strings.Replace(localVarPath, "{"+"distro_base_path"+"}", parameterValueToString(r.distroBasePath, "distroBasePath"), -1)  // NOTE: paths aren't escaped because Pulp uses hrefs as path parameters
	localVarPath = strings.Replace(localVarPath, "{"+"path"+"}", parameterValueToString(r.path, "path"), -1)  // NOTE: paths aren't escaped because Pulp uses hrefs as path parameters

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.file == nil {
		return localVarReturnValue, nil, reportError("file is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"multipart/form-data", "application/x-www-form-urlencoded"}

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
	var fileLocalVarFormFileName string
	var fileLocalVarFileName     string
	var fileLocalVarFileBytes    []byte

	fileLocalVarFormFileName = "file"


	fileLocalVarFile := r.file

	if fileLocalVarFile != nil {
		fbs, _ := io.ReadAll(fileLocalVarFile)

		fileLocalVarFileBytes = fbs
		fileLocalVarFileName = fileLocalVarFile.Name()
		fileLocalVarFile.Close()
		formFiles = append(formFiles, formFile{fileBytes: fileLocalVarFileBytes, fileName: fileLocalVarFileName, formFileName: fileLocalVarFormFileName})
	}
	if r.sha256 != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "sha256", r.sha256, "")
	}
	if r.expectedNamespace != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "expected_namespace", r.expectedNamespace, "")
	}
	if r.expectedName != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "expected_name", r.expectedName, "")
	}
	if r.expectedVersion != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "expected_version", r.expectedVersion, "")
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

type PulpAnsibleArtifactsCollectionsV3APIPulpAnsibleGalaxyDefaultApiV3ArtifactsCollectionsCreateRequest struct {
	ctx context.Context
	ApiService *PulpAnsibleArtifactsCollectionsV3APIService
	file *os.File
	sha256 *string
	expectedNamespace *string
	expectedName *string
	expectedVersion *string
}

// The Collection tarball.
func (r PulpAnsibleArtifactsCollectionsV3APIPulpAnsibleGalaxyDefaultApiV3ArtifactsCollectionsCreateRequest) File(file *os.File) PulpAnsibleArtifactsCollectionsV3APIPulpAnsibleGalaxyDefaultApiV3ArtifactsCollectionsCreateRequest {
	r.file = file
	return r
}

// An optional sha256 checksum of the uploaded file.
func (r PulpAnsibleArtifactsCollectionsV3APIPulpAnsibleGalaxyDefaultApiV3ArtifactsCollectionsCreateRequest) Sha256(sha256 string) PulpAnsibleArtifactsCollectionsV3APIPulpAnsibleGalaxyDefaultApiV3ArtifactsCollectionsCreateRequest {
	r.sha256 = &sha256
	return r
}

// The expected &#39;namespace&#39; of the Collection to be verified against the metadata during import.
func (r PulpAnsibleArtifactsCollectionsV3APIPulpAnsibleGalaxyDefaultApiV3ArtifactsCollectionsCreateRequest) ExpectedNamespace(expectedNamespace string) PulpAnsibleArtifactsCollectionsV3APIPulpAnsibleGalaxyDefaultApiV3ArtifactsCollectionsCreateRequest {
	r.expectedNamespace = &expectedNamespace
	return r
}

// The expected &#39;name&#39; of the Collection to be verified against the metadata during import.
func (r PulpAnsibleArtifactsCollectionsV3APIPulpAnsibleGalaxyDefaultApiV3ArtifactsCollectionsCreateRequest) ExpectedName(expectedName string) PulpAnsibleArtifactsCollectionsV3APIPulpAnsibleGalaxyDefaultApiV3ArtifactsCollectionsCreateRequest {
	r.expectedName = &expectedName
	return r
}

// The expected version of the Collection to be verified against the metadata during import.
func (r PulpAnsibleArtifactsCollectionsV3APIPulpAnsibleGalaxyDefaultApiV3ArtifactsCollectionsCreateRequest) ExpectedVersion(expectedVersion string) PulpAnsibleArtifactsCollectionsV3APIPulpAnsibleGalaxyDefaultApiV3ArtifactsCollectionsCreateRequest {
	r.expectedVersion = &expectedVersion
	return r
}

func (r PulpAnsibleArtifactsCollectionsV3APIPulpAnsibleGalaxyDefaultApiV3ArtifactsCollectionsCreateRequest) Execute() (*AsyncOperationResponse, *http.Response, error) {
	return r.ApiService.PulpAnsibleGalaxyDefaultApiV3ArtifactsCollectionsCreateExecute(r)
}

/*
PulpAnsibleGalaxyDefaultApiV3ArtifactsCollectionsCreate Upload a collection

Create an artifact and trigger an asynchronous task to create Collection content from it.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return PulpAnsibleArtifactsCollectionsV3APIPulpAnsibleGalaxyDefaultApiV3ArtifactsCollectionsCreateRequest

Deprecated
*/
func (a *PulpAnsibleArtifactsCollectionsV3APIService) PulpAnsibleGalaxyDefaultApiV3ArtifactsCollectionsCreate(ctx context.Context) PulpAnsibleArtifactsCollectionsV3APIPulpAnsibleGalaxyDefaultApiV3ArtifactsCollectionsCreateRequest {
	return PulpAnsibleArtifactsCollectionsV3APIPulpAnsibleGalaxyDefaultApiV3ArtifactsCollectionsCreateRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return AsyncOperationResponse
// Deprecated
func (a *PulpAnsibleArtifactsCollectionsV3APIService) PulpAnsibleGalaxyDefaultApiV3ArtifactsCollectionsCreateExecute(r PulpAnsibleArtifactsCollectionsV3APIPulpAnsibleGalaxyDefaultApiV3ArtifactsCollectionsCreateRequest) (*AsyncOperationResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AsyncOperationResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PulpAnsibleArtifactsCollectionsV3APIService.PulpAnsibleGalaxyDefaultApiV3ArtifactsCollectionsCreate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pulp_ansible/galaxy/default/api/v3/artifacts/collections/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.file == nil {
		return localVarReturnValue, nil, reportError("file is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"multipart/form-data", "application/x-www-form-urlencoded"}

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
	var fileLocalVarFormFileName string
	var fileLocalVarFileName     string
	var fileLocalVarFileBytes    []byte

	fileLocalVarFormFileName = "file"


	fileLocalVarFile := r.file

	if fileLocalVarFile != nil {
		fbs, _ := io.ReadAll(fileLocalVarFile)

		fileLocalVarFileBytes = fbs
		fileLocalVarFileName = fileLocalVarFile.Name()
		fileLocalVarFile.Close()
		formFiles = append(formFiles, formFile{fileBytes: fileLocalVarFileBytes, fileName: fileLocalVarFileName, formFileName: fileLocalVarFormFileName})
	}
	if r.sha256 != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "sha256", r.sha256, "")
	}
	if r.expectedNamespace != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "expected_namespace", r.expectedNamespace, "")
	}
	if r.expectedName != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "expected_name", r.expectedName, "")
	}
	if r.expectedVersion != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "expected_version", r.expectedVersion, "")
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

type PulpAnsibleArtifactsCollectionsV3APIPulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsArtifactsCreateRequest struct {
	ctx context.Context
	ApiService *PulpAnsibleArtifactsCollectionsV3APIService
	distroBasePath string
	file *os.File
	sha256 *string
	expectedNamespace *string
	expectedName *string
	expectedVersion *string
}

// The Collection tarball.
func (r PulpAnsibleArtifactsCollectionsV3APIPulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsArtifactsCreateRequest) File(file *os.File) PulpAnsibleArtifactsCollectionsV3APIPulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsArtifactsCreateRequest {
	r.file = file
	return r
}

// An optional sha256 checksum of the uploaded file.
func (r PulpAnsibleArtifactsCollectionsV3APIPulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsArtifactsCreateRequest) Sha256(sha256 string) PulpAnsibleArtifactsCollectionsV3APIPulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsArtifactsCreateRequest {
	r.sha256 = &sha256
	return r
}

// The expected &#39;namespace&#39; of the Collection to be verified against the metadata during import.
func (r PulpAnsibleArtifactsCollectionsV3APIPulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsArtifactsCreateRequest) ExpectedNamespace(expectedNamespace string) PulpAnsibleArtifactsCollectionsV3APIPulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsArtifactsCreateRequest {
	r.expectedNamespace = &expectedNamespace
	return r
}

// The expected &#39;name&#39; of the Collection to be verified against the metadata during import.
func (r PulpAnsibleArtifactsCollectionsV3APIPulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsArtifactsCreateRequest) ExpectedName(expectedName string) PulpAnsibleArtifactsCollectionsV3APIPulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsArtifactsCreateRequest {
	r.expectedName = &expectedName
	return r
}

// The expected version of the Collection to be verified against the metadata during import.
func (r PulpAnsibleArtifactsCollectionsV3APIPulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsArtifactsCreateRequest) ExpectedVersion(expectedVersion string) PulpAnsibleArtifactsCollectionsV3APIPulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsArtifactsCreateRequest {
	r.expectedVersion = &expectedVersion
	return r
}

func (r PulpAnsibleArtifactsCollectionsV3APIPulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsArtifactsCreateRequest) Execute() (*AsyncOperationResponse, *http.Response, error) {
	return r.ApiService.PulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsArtifactsCreateExecute(r)
}

/*
PulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsArtifactsCreate Upload a collection

Create an artifact and trigger an asynchronous task to create Collection content from it.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param distroBasePath
 @return PulpAnsibleArtifactsCollectionsV3APIPulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsArtifactsCreateRequest
*/
func (a *PulpAnsibleArtifactsCollectionsV3APIService) PulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsArtifactsCreate(ctx context.Context, distroBasePath string) PulpAnsibleArtifactsCollectionsV3APIPulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsArtifactsCreateRequest {
	return PulpAnsibleArtifactsCollectionsV3APIPulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsArtifactsCreateRequest{
		ApiService: a,
		ctx: ctx,
		distroBasePath: distroBasePath,
	}
}

// Execute executes the request
//  @return AsyncOperationResponse
func (a *PulpAnsibleArtifactsCollectionsV3APIService) PulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsArtifactsCreateExecute(r PulpAnsibleArtifactsCollectionsV3APIPulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsArtifactsCreateRequest) (*AsyncOperationResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AsyncOperationResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PulpAnsibleArtifactsCollectionsV3APIService.PulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsArtifactsCreate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pulp_ansible/galaxy/default/api/v3/plugin/ansible/content/{distro_base_path}/collections/artifacts/"
	localVarPath = strings.Replace(localVarPath, "{"+"distro_base_path"+"}", parameterValueToString(r.distroBasePath, "distroBasePath"), -1)  // NOTE: paths aren't escaped because Pulp uses hrefs as path parameters

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.file == nil {
		return localVarReturnValue, nil, reportError("file is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"multipart/form-data", "application/x-www-form-urlencoded"}

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
	var fileLocalVarFormFileName string
	var fileLocalVarFileName     string
	var fileLocalVarFileBytes    []byte

	fileLocalVarFormFileName = "file"


	fileLocalVarFile := r.file

	if fileLocalVarFile != nil {
		fbs, _ := io.ReadAll(fileLocalVarFile)

		fileLocalVarFileBytes = fbs
		fileLocalVarFileName = fileLocalVarFile.Name()
		fileLocalVarFile.Close()
		formFiles = append(formFiles, formFile{fileBytes: fileLocalVarFileBytes, fileName: fileLocalVarFileName, formFileName: fileLocalVarFormFileName})
	}
	if r.sha256 != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "sha256", r.sha256, "")
	}
	if r.expectedNamespace != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "expected_namespace", r.expectedNamespace, "")
	}
	if r.expectedName != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "expected_name", r.expectedName, "")
	}
	if r.expectedVersion != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "expected_version", r.expectedVersion, "")
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
