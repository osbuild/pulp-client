# \ContentPackageReleaseComponentsAPI

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ContentDebPackageReleaseComponentsCreate**](ContentPackageReleaseComponentsAPI.md#ContentDebPackageReleaseComponentsCreate) | **Post** /pulp/api/v3/content/deb/package_release_components/ | Create a package release component
[**ContentDebPackageReleaseComponentsList**](ContentPackageReleaseComponentsAPI.md#ContentDebPackageReleaseComponentsList) | **Get** /pulp/api/v3/content/deb/package_release_components/ | List package release components
[**ContentDebPackageReleaseComponentsRead**](ContentPackageReleaseComponentsAPI.md#ContentDebPackageReleaseComponentsRead) | **Get** {deb_package_release_component_href} | Inspect a package release component



## ContentDebPackageReleaseComponentsCreate

> DebPackageReleaseComponentResponse ContentDebPackageReleaseComponentsCreate(ctx).DebPackageReleaseComponent(debPackageReleaseComponent).Execute()

Create a package release component



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
    debPackageReleaseComponent := *openapiclient.NewDebPackageReleaseComponent("Package_example", "ReleaseComponent_example") // DebPackageReleaseComponent | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContentPackageReleaseComponentsAPI.ContentDebPackageReleaseComponentsCreate(context.Background()).DebPackageReleaseComponent(debPackageReleaseComponent).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContentPackageReleaseComponentsAPI.ContentDebPackageReleaseComponentsCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContentDebPackageReleaseComponentsCreate`: DebPackageReleaseComponentResponse
    fmt.Fprintf(os.Stdout, "Response from `ContentPackageReleaseComponentsAPI.ContentDebPackageReleaseComponentsCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiContentDebPackageReleaseComponentsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **debPackageReleaseComponent** | [**DebPackageReleaseComponent**](DebPackageReleaseComponent.md) |  | 

### Return type

[**DebPackageReleaseComponentResponse**](DebPackageReleaseComponentResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContentDebPackageReleaseComponentsList

> PaginateddebPackageReleaseComponentResponseList ContentDebPackageReleaseComponentsList(ctx).Limit(limit).Offset(offset).Ordering(ordering).Package_(package_).PulpHrefIn(pulpHrefIn).PulpIdIn(pulpIdIn).ReleaseComponent(releaseComponent).RepositoryVersion(repositoryVersion).RepositoryVersionAdded(repositoryVersionAdded).RepositoryVersionRemoved(repositoryVersionRemoved).Fields(fields).ExcludeFields(excludeFields).Execute()

List package release components



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
    limit := int32(56) // int32 | Number of results to return per page. (optional)
    offset := int32(56) // int32 | The initial index from which to return the results. (optional)
    ordering := []string{"Ordering_example"} // []string | Ordering  * `pulp_id` - Pulp id * `-pulp_id` - Pulp id (descending) * `pulp_created` - Pulp created * `-pulp_created` - Pulp created (descending) * `pulp_last_updated` - Pulp last updated * `-pulp_last_updated` - Pulp last updated (descending) * `pulp_type` - Pulp type * `-pulp_type` - Pulp type (descending) * `upstream_id` - Upstream id * `-upstream_id` - Upstream id (descending) * `timestamp_of_interest` - Timestamp of interest * `-timestamp_of_interest` - Timestamp of interest (descending) * `pk` - Pk * `-pk` - Pk (descending) (optional)
    package_ := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Filter results where package matches value (optional)
    pulpHrefIn := []string{"Inner_example"} // []string | Multiple values may be separated by commas. (optional)
    pulpIdIn := []string{"Inner_example"} // []string | Multiple values may be separated by commas. (optional)
    releaseComponent := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Filter results where release_component matches value (optional)
    repositoryVersion := "repositoryVersion_example" // string | Repository Version referenced by HREF (optional)
    repositoryVersionAdded := "repositoryVersionAdded_example" // string | Repository Version referenced by HREF (optional)
    repositoryVersionRemoved := "repositoryVersionRemoved_example" // string | Repository Version referenced by HREF (optional)
    fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
    excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContentPackageReleaseComponentsAPI.ContentDebPackageReleaseComponentsList(context.Background()).Limit(limit).Offset(offset).Ordering(ordering).Package_(package_).PulpHrefIn(pulpHrefIn).PulpIdIn(pulpIdIn).ReleaseComponent(releaseComponent).RepositoryVersion(repositoryVersion).RepositoryVersionAdded(repositoryVersionAdded).RepositoryVersionRemoved(repositoryVersionRemoved).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContentPackageReleaseComponentsAPI.ContentDebPackageReleaseComponentsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContentDebPackageReleaseComponentsList`: PaginateddebPackageReleaseComponentResponseList
    fmt.Fprintf(os.Stdout, "Response from `ContentPackageReleaseComponentsAPI.ContentDebPackageReleaseComponentsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiContentDebPackageReleaseComponentsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Number of results to return per page. | 
 **offset** | **int32** | The initial index from which to return the results. | 
 **ordering** | **[]string** | Ordering  * &#x60;pulp_id&#x60; - Pulp id * &#x60;-pulp_id&#x60; - Pulp id (descending) * &#x60;pulp_created&#x60; - Pulp created * &#x60;-pulp_created&#x60; - Pulp created (descending) * &#x60;pulp_last_updated&#x60; - Pulp last updated * &#x60;-pulp_last_updated&#x60; - Pulp last updated (descending) * &#x60;pulp_type&#x60; - Pulp type * &#x60;-pulp_type&#x60; - Pulp type (descending) * &#x60;upstream_id&#x60; - Upstream id * &#x60;-upstream_id&#x60; - Upstream id (descending) * &#x60;timestamp_of_interest&#x60; - Timestamp of interest * &#x60;-timestamp_of_interest&#x60; - Timestamp of interest (descending) * &#x60;pk&#x60; - Pk * &#x60;-pk&#x60; - Pk (descending) | 
 **package_** | **string** | Filter results where package matches value | 
 **pulpHrefIn** | **[]string** | Multiple values may be separated by commas. | 
 **pulpIdIn** | **[]string** | Multiple values may be separated by commas. | 
 **releaseComponent** | **string** | Filter results where release_component matches value | 
 **repositoryVersion** | **string** | Repository Version referenced by HREF | 
 **repositoryVersionAdded** | **string** | Repository Version referenced by HREF | 
 **repositoryVersionRemoved** | **string** | Repository Version referenced by HREF | 
 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

### Return type

[**PaginateddebPackageReleaseComponentResponseList**](PaginateddebPackageReleaseComponentResponseList.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContentDebPackageReleaseComponentsRead

> DebPackageReleaseComponentResponse ContentDebPackageReleaseComponentsRead(ctx, debPackageReleaseComponentHref).Fields(fields).ExcludeFields(excludeFields).Execute()

Inspect a package release component



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
    debPackageReleaseComponentHref := "debPackageReleaseComponentHref_example" // string | 
    fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
    excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContentPackageReleaseComponentsAPI.ContentDebPackageReleaseComponentsRead(context.Background(), debPackageReleaseComponentHref).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContentPackageReleaseComponentsAPI.ContentDebPackageReleaseComponentsRead``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContentDebPackageReleaseComponentsRead`: DebPackageReleaseComponentResponse
    fmt.Fprintf(os.Stdout, "Response from `ContentPackageReleaseComponentsAPI.ContentDebPackageReleaseComponentsRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**debPackageReleaseComponentHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiContentDebPackageReleaseComponentsReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

### Return type

[**DebPackageReleaseComponentResponse**](DebPackageReleaseComponentResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

