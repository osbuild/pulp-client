# \ContentInstallerFileIndicesAPI

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ContentDebInstallerFileIndicesCreate**](ContentInstallerFileIndicesAPI.md#ContentDebInstallerFileIndicesCreate) | **Post** /pulp/api/v3/content/deb/installer_file_indices/ | Create an installer file index
[**ContentDebInstallerFileIndicesList**](ContentInstallerFileIndicesAPI.md#ContentDebInstallerFileIndicesList) | **Get** /pulp/api/v3/content/deb/installer_file_indices/ | List InstallerFileIndices
[**ContentDebInstallerFileIndicesRead**](ContentInstallerFileIndicesAPI.md#ContentDebInstallerFileIndicesRead) | **Get** {deb_installer_file_index_href} | Inspect an installer file index



## ContentDebInstallerFileIndicesCreate

> DebInstallerFileIndexResponse ContentDebInstallerFileIndicesCreate(ctx).DebInstallerFileIndex(debInstallerFileIndex).Execute()

Create an installer file index



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
    debInstallerFileIndex := *openapiclient.NewDebInstallerFileIndex(map[string]interface{}(123), "Component_example", "Architecture_example") // DebInstallerFileIndex | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContentInstallerFileIndicesAPI.ContentDebInstallerFileIndicesCreate(context.Background()).DebInstallerFileIndex(debInstallerFileIndex).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContentInstallerFileIndicesAPI.ContentDebInstallerFileIndicesCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContentDebInstallerFileIndicesCreate`: DebInstallerFileIndexResponse
    fmt.Fprintf(os.Stdout, "Response from `ContentInstallerFileIndicesAPI.ContentDebInstallerFileIndicesCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiContentDebInstallerFileIndicesCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **debInstallerFileIndex** | [**DebInstallerFileIndex**](DebInstallerFileIndex.md) |  | 

### Return type

[**DebInstallerFileIndexResponse**](DebInstallerFileIndexResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContentDebInstallerFileIndicesList

> PaginateddebInstallerFileIndexResponseList ContentDebInstallerFileIndicesList(ctx).Architecture(architecture).Component(component).Limit(limit).Offset(offset).Ordering(ordering).PulpHrefIn(pulpHrefIn).PulpIdIn(pulpIdIn).RelativePath(relativePath).RepositoryVersion(repositoryVersion).RepositoryVersionAdded(repositoryVersionAdded).RepositoryVersionRemoved(repositoryVersionRemoved).Sha256(sha256).Fields(fields).ExcludeFields(excludeFields).Execute()

List InstallerFileIndices



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
    architecture := "architecture_example" // string | Filter results where architecture matches value (optional)
    component := "component_example" // string | Filter results where component matches value (optional)
    limit := int32(56) // int32 | Number of results to return per page. (optional)
    offset := int32(56) // int32 | The initial index from which to return the results. (optional)
    ordering := []string{"Ordering_example"} // []string | Ordering  * `pulp_id` - Pulp id * `-pulp_id` - Pulp id (descending) * `pulp_created` - Pulp created * `-pulp_created` - Pulp created (descending) * `pulp_last_updated` - Pulp last updated * `-pulp_last_updated` - Pulp last updated (descending) * `pulp_type` - Pulp type * `-pulp_type` - Pulp type (descending) * `upstream_id` - Upstream id * `-upstream_id` - Upstream id (descending) * `timestamp_of_interest` - Timestamp of interest * `-timestamp_of_interest` - Timestamp of interest (descending) * `component` - Component * `-component` - Component (descending) * `architecture` - Architecture * `-architecture` - Architecture (descending) * `relative_path` - Relative path * `-relative_path` - Relative path (descending) * `sha256` - Sha256 * `-sha256` - Sha256 (descending) * `pk` - Pk * `-pk` - Pk (descending) (optional)
    pulpHrefIn := []string{"Inner_example"} // []string | Multiple values may be separated by commas. (optional)
    pulpIdIn := []string{"Inner_example"} // []string | Multiple values may be separated by commas. (optional)
    relativePath := "relativePath_example" // string | Filter results where relative_path matches value (optional)
    repositoryVersion := "repositoryVersion_example" // string | Repository Version referenced by HREF (optional)
    repositoryVersionAdded := "repositoryVersionAdded_example" // string | Repository Version referenced by HREF (optional)
    repositoryVersionRemoved := "repositoryVersionRemoved_example" // string | Repository Version referenced by HREF (optional)
    sha256 := "sha256_example" // string | Filter results where sha256 matches value (optional)
    fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
    excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContentInstallerFileIndicesAPI.ContentDebInstallerFileIndicesList(context.Background()).Architecture(architecture).Component(component).Limit(limit).Offset(offset).Ordering(ordering).PulpHrefIn(pulpHrefIn).PulpIdIn(pulpIdIn).RelativePath(relativePath).RepositoryVersion(repositoryVersion).RepositoryVersionAdded(repositoryVersionAdded).RepositoryVersionRemoved(repositoryVersionRemoved).Sha256(sha256).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContentInstallerFileIndicesAPI.ContentDebInstallerFileIndicesList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContentDebInstallerFileIndicesList`: PaginateddebInstallerFileIndexResponseList
    fmt.Fprintf(os.Stdout, "Response from `ContentInstallerFileIndicesAPI.ContentDebInstallerFileIndicesList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiContentDebInstallerFileIndicesListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **architecture** | **string** | Filter results where architecture matches value | 
 **component** | **string** | Filter results where component matches value | 
 **limit** | **int32** | Number of results to return per page. | 
 **offset** | **int32** | The initial index from which to return the results. | 
 **ordering** | **[]string** | Ordering  * &#x60;pulp_id&#x60; - Pulp id * &#x60;-pulp_id&#x60; - Pulp id (descending) * &#x60;pulp_created&#x60; - Pulp created * &#x60;-pulp_created&#x60; - Pulp created (descending) * &#x60;pulp_last_updated&#x60; - Pulp last updated * &#x60;-pulp_last_updated&#x60; - Pulp last updated (descending) * &#x60;pulp_type&#x60; - Pulp type * &#x60;-pulp_type&#x60; - Pulp type (descending) * &#x60;upstream_id&#x60; - Upstream id * &#x60;-upstream_id&#x60; - Upstream id (descending) * &#x60;timestamp_of_interest&#x60; - Timestamp of interest * &#x60;-timestamp_of_interest&#x60; - Timestamp of interest (descending) * &#x60;component&#x60; - Component * &#x60;-component&#x60; - Component (descending) * &#x60;architecture&#x60; - Architecture * &#x60;-architecture&#x60; - Architecture (descending) * &#x60;relative_path&#x60; - Relative path * &#x60;-relative_path&#x60; - Relative path (descending) * &#x60;sha256&#x60; - Sha256 * &#x60;-sha256&#x60; - Sha256 (descending) * &#x60;pk&#x60; - Pk * &#x60;-pk&#x60; - Pk (descending) | 
 **pulpHrefIn** | **[]string** | Multiple values may be separated by commas. | 
 **pulpIdIn** | **[]string** | Multiple values may be separated by commas. | 
 **relativePath** | **string** | Filter results where relative_path matches value | 
 **repositoryVersion** | **string** | Repository Version referenced by HREF | 
 **repositoryVersionAdded** | **string** | Repository Version referenced by HREF | 
 **repositoryVersionRemoved** | **string** | Repository Version referenced by HREF | 
 **sha256** | **string** | Filter results where sha256 matches value | 
 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

### Return type

[**PaginateddebInstallerFileIndexResponseList**](PaginateddebInstallerFileIndexResponseList.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContentDebInstallerFileIndicesRead

> DebInstallerFileIndexResponse ContentDebInstallerFileIndicesRead(ctx, debInstallerFileIndexHref).Fields(fields).ExcludeFields(excludeFields).Execute()

Inspect an installer file index



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
    debInstallerFileIndexHref := "debInstallerFileIndexHref_example" // string | 
    fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
    excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContentInstallerFileIndicesAPI.ContentDebInstallerFileIndicesRead(context.Background(), debInstallerFileIndexHref).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContentInstallerFileIndicesAPI.ContentDebInstallerFileIndicesRead``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContentDebInstallerFileIndicesRead`: DebInstallerFileIndexResponse
    fmt.Fprintf(os.Stdout, "Response from `ContentInstallerFileIndicesAPI.ContentDebInstallerFileIndicesRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**debInstallerFileIndexHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiContentDebInstallerFileIndicesReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

### Return type

[**DebInstallerFileIndexResponse**](DebInstallerFileIndexResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

