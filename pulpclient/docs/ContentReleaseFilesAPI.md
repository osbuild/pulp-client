# \ContentReleaseFilesAPI

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ContentDebReleaseFilesCreate**](ContentReleaseFilesAPI.md#ContentDebReleaseFilesCreate) | **Post** /pulp/api/v3/content/deb/release_files/ | Create a release file
[**ContentDebReleaseFilesList**](ContentReleaseFilesAPI.md#ContentDebReleaseFilesList) | **Get** /pulp/api/v3/content/deb/release_files/ | List release files
[**ContentDebReleaseFilesRead**](ContentReleaseFilesAPI.md#ContentDebReleaseFilesRead) | **Get** {deb_release_file_href} | Inspect a release file



## ContentDebReleaseFilesCreate

> DebReleaseFileResponse ContentDebReleaseFilesCreate(ctx).DebReleaseFile(debReleaseFile).Execute()

Create a release file



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/osbuild/pulp-client/pulpclient"
)

func main() {
    debReleaseFile := *openapiclient.NewDebReleaseFile(map[string]interface{}(123), "Distribution_example") // DebReleaseFile | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContentReleaseFilesAPI.ContentDebReleaseFilesCreate(context.Background()).DebReleaseFile(debReleaseFile).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContentReleaseFilesAPI.ContentDebReleaseFilesCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContentDebReleaseFilesCreate`: DebReleaseFileResponse
    fmt.Fprintf(os.Stdout, "Response from `ContentReleaseFilesAPI.ContentDebReleaseFilesCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiContentDebReleaseFilesCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **debReleaseFile** | [**DebReleaseFile**](DebReleaseFile.md) |  | 

### Return type

[**DebReleaseFileResponse**](DebReleaseFileResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContentDebReleaseFilesList

> PaginateddebReleaseFileResponseList ContentDebReleaseFilesList(ctx).Codename(codename).Limit(limit).Offset(offset).Ordering(ordering).PulpHrefIn(pulpHrefIn).PulpIdIn(pulpIdIn).RelativePath(relativePath).RepositoryVersion(repositoryVersion).RepositoryVersionAdded(repositoryVersionAdded).RepositoryVersionRemoved(repositoryVersionRemoved).Sha256(sha256).Suite(suite).Fields(fields).ExcludeFields(excludeFields).Execute()

List release files



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/osbuild/pulp-client/pulpclient"
)

func main() {
    codename := "codename_example" // string | Filter results where codename matches value (optional)
    limit := int32(56) // int32 | Number of results to return per page. (optional)
    offset := int32(56) // int32 | The initial index from which to return the results. (optional)
    ordering := []string{"Ordering_example"} // []string | Ordering  * `pulp_id` - Pulp id * `-pulp_id` - Pulp id (descending) * `pulp_created` - Pulp created * `-pulp_created` - Pulp created (descending) * `pulp_last_updated` - Pulp last updated * `-pulp_last_updated` - Pulp last updated (descending) * `pulp_type` - Pulp type * `-pulp_type` - Pulp type (descending) * `upstream_id` - Upstream id * `-upstream_id` - Upstream id (descending) * `timestamp_of_interest` - Timestamp of interest * `-timestamp_of_interest` - Timestamp of interest (descending) * `codename` - Codename * `-codename` - Codename (descending) * `suite` - Suite * `-suite` - Suite (descending) * `distribution` - Distribution * `-distribution` - Distribution (descending) * `components` - Components * `-components` - Components (descending) * `architectures` - Architectures * `-architectures` - Architectures (descending) * `relative_path` - Relative path * `-relative_path` - Relative path (descending) * `sha256` - Sha256 * `-sha256` - Sha256 (descending) * `artifact_set_sha256` - Artifact set sha256 * `-artifact_set_sha256` - Artifact set sha256 (descending) * `pk` - Pk * `-pk` - Pk (descending) (optional)
    pulpHrefIn := []string{"Inner_example"} // []string | Multiple values may be separated by commas. (optional)
    pulpIdIn := []string{"Inner_example"} // []string | Multiple values may be separated by commas. (optional)
    relativePath := "relativePath_example" // string | Filter results where relative_path matches value (optional)
    repositoryVersion := "repositoryVersion_example" // string | Repository Version referenced by HREF (optional)
    repositoryVersionAdded := "repositoryVersionAdded_example" // string | Repository Version referenced by HREF (optional)
    repositoryVersionRemoved := "repositoryVersionRemoved_example" // string | Repository Version referenced by HREF (optional)
    sha256 := "sha256_example" // string | Filter results where sha256 matches value (optional)
    suite := "suite_example" // string | Filter results where suite matches value (optional)
    fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
    excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContentReleaseFilesAPI.ContentDebReleaseFilesList(context.Background()).Codename(codename).Limit(limit).Offset(offset).Ordering(ordering).PulpHrefIn(pulpHrefIn).PulpIdIn(pulpIdIn).RelativePath(relativePath).RepositoryVersion(repositoryVersion).RepositoryVersionAdded(repositoryVersionAdded).RepositoryVersionRemoved(repositoryVersionRemoved).Sha256(sha256).Suite(suite).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContentReleaseFilesAPI.ContentDebReleaseFilesList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContentDebReleaseFilesList`: PaginateddebReleaseFileResponseList
    fmt.Fprintf(os.Stdout, "Response from `ContentReleaseFilesAPI.ContentDebReleaseFilesList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiContentDebReleaseFilesListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **codename** | **string** | Filter results where codename matches value | 
 **limit** | **int32** | Number of results to return per page. | 
 **offset** | **int32** | The initial index from which to return the results. | 
 **ordering** | **[]string** | Ordering  * &#x60;pulp_id&#x60; - Pulp id * &#x60;-pulp_id&#x60; - Pulp id (descending) * &#x60;pulp_created&#x60; - Pulp created * &#x60;-pulp_created&#x60; - Pulp created (descending) * &#x60;pulp_last_updated&#x60; - Pulp last updated * &#x60;-pulp_last_updated&#x60; - Pulp last updated (descending) * &#x60;pulp_type&#x60; - Pulp type * &#x60;-pulp_type&#x60; - Pulp type (descending) * &#x60;upstream_id&#x60; - Upstream id * &#x60;-upstream_id&#x60; - Upstream id (descending) * &#x60;timestamp_of_interest&#x60; - Timestamp of interest * &#x60;-timestamp_of_interest&#x60; - Timestamp of interest (descending) * &#x60;codename&#x60; - Codename * &#x60;-codename&#x60; - Codename (descending) * &#x60;suite&#x60; - Suite * &#x60;-suite&#x60; - Suite (descending) * &#x60;distribution&#x60; - Distribution * &#x60;-distribution&#x60; - Distribution (descending) * &#x60;components&#x60; - Components * &#x60;-components&#x60; - Components (descending) * &#x60;architectures&#x60; - Architectures * &#x60;-architectures&#x60; - Architectures (descending) * &#x60;relative_path&#x60; - Relative path * &#x60;-relative_path&#x60; - Relative path (descending) * &#x60;sha256&#x60; - Sha256 * &#x60;-sha256&#x60; - Sha256 (descending) * &#x60;artifact_set_sha256&#x60; - Artifact set sha256 * &#x60;-artifact_set_sha256&#x60; - Artifact set sha256 (descending) * &#x60;pk&#x60; - Pk * &#x60;-pk&#x60; - Pk (descending) | 
 **pulpHrefIn** | **[]string** | Multiple values may be separated by commas. | 
 **pulpIdIn** | **[]string** | Multiple values may be separated by commas. | 
 **relativePath** | **string** | Filter results where relative_path matches value | 
 **repositoryVersion** | **string** | Repository Version referenced by HREF | 
 **repositoryVersionAdded** | **string** | Repository Version referenced by HREF | 
 **repositoryVersionRemoved** | **string** | Repository Version referenced by HREF | 
 **sha256** | **string** | Filter results where sha256 matches value | 
 **suite** | **string** | Filter results where suite matches value | 
 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

### Return type

[**PaginateddebReleaseFileResponseList**](PaginateddebReleaseFileResponseList.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContentDebReleaseFilesRead

> DebReleaseFileResponse ContentDebReleaseFilesRead(ctx, debReleaseFileHref).Fields(fields).ExcludeFields(excludeFields).Execute()

Inspect a release file



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/osbuild/pulp-client/pulpclient"
)

func main() {
    debReleaseFileHref := "debReleaseFileHref_example" // string | 
    fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
    excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContentReleaseFilesAPI.ContentDebReleaseFilesRead(context.Background(), debReleaseFileHref).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContentReleaseFilesAPI.ContentDebReleaseFilesRead``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContentDebReleaseFilesRead`: DebReleaseFileResponse
    fmt.Fprintf(os.Stdout, "Response from `ContentReleaseFilesAPI.ContentDebReleaseFilesRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**debReleaseFileHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiContentDebReleaseFilesReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

### Return type

[**DebReleaseFileResponse**](DebReleaseFileResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

