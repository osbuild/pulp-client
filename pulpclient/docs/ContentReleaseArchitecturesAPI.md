# \ContentReleaseArchitecturesAPI

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ContentDebReleaseArchitecturesCreate**](ContentReleaseArchitecturesAPI.md#ContentDebReleaseArchitecturesCreate) | **Post** /pulp/api/v3/content/deb/release_architectures/ | Create a release architecture
[**ContentDebReleaseArchitecturesList**](ContentReleaseArchitecturesAPI.md#ContentDebReleaseArchitecturesList) | **Get** /pulp/api/v3/content/deb/release_architectures/ | List release architectures
[**ContentDebReleaseArchitecturesRead**](ContentReleaseArchitecturesAPI.md#ContentDebReleaseArchitecturesRead) | **Get** {deb_release_architecture_href} | Inspect a release architecture



## ContentDebReleaseArchitecturesCreate

> DebReleaseArchitectureResponse ContentDebReleaseArchitecturesCreate(ctx).DebReleaseArchitecture(debReleaseArchitecture).Execute()

Create a release architecture



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
    debReleaseArchitecture := *openapiclient.NewDebReleaseArchitecture("Architecture_example", "Distribution_example", "Codename_example", "Suite_example") // DebReleaseArchitecture | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContentReleaseArchitecturesAPI.ContentDebReleaseArchitecturesCreate(context.Background()).DebReleaseArchitecture(debReleaseArchitecture).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContentReleaseArchitecturesAPI.ContentDebReleaseArchitecturesCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContentDebReleaseArchitecturesCreate`: DebReleaseArchitectureResponse
    fmt.Fprintf(os.Stdout, "Response from `ContentReleaseArchitecturesAPI.ContentDebReleaseArchitecturesCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiContentDebReleaseArchitecturesCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **debReleaseArchitecture** | [**DebReleaseArchitecture**](DebReleaseArchitecture.md) |  | 

### Return type

[**DebReleaseArchitectureResponse**](DebReleaseArchitectureResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContentDebReleaseArchitecturesList

> PaginateddebReleaseArchitectureResponseList ContentDebReleaseArchitecturesList(ctx).Architecture(architecture).Codename(codename).Distribution(distribution).Limit(limit).Offset(offset).Ordering(ordering).PulpHrefIn(pulpHrefIn).PulpIdIn(pulpIdIn).RepositoryVersion(repositoryVersion).RepositoryVersionAdded(repositoryVersionAdded).RepositoryVersionRemoved(repositoryVersionRemoved).Suite(suite).Fields(fields).ExcludeFields(excludeFields).Execute()

List release architectures



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
    codename := "codename_example" // string | Filter results where codename matches value (optional)
    distribution := "distribution_example" // string | Filter results where distribution matches value (optional)
    limit := int32(56) // int32 | Number of results to return per page. (optional)
    offset := int32(56) // int32 | The initial index from which to return the results. (optional)
    ordering := []string{"Ordering_example"} // []string | Ordering  * `pulp_id` - Pulp id * `-pulp_id` - Pulp id (descending) * `pulp_created` - Pulp created * `-pulp_created` - Pulp created (descending) * `pulp_last_updated` - Pulp last updated * `-pulp_last_updated` - Pulp last updated (descending) * `pulp_type` - Pulp type * `-pulp_type` - Pulp type (descending) * `upstream_id` - Upstream id * `-upstream_id` - Upstream id (descending) * `timestamp_of_interest` - Timestamp of interest * `-timestamp_of_interest` - Timestamp of interest (descending) * `architecture` - Architecture * `-architecture` - Architecture (descending) * `distribution` - Distribution * `-distribution` - Distribution (descending) * `codename` - Codename * `-codename` - Codename (descending) * `suite` - Suite * `-suite` - Suite (descending) * `pk` - Pk * `-pk` - Pk (descending) (optional)
    pulpHrefIn := []string{"Inner_example"} // []string | Multiple values may be separated by commas. (optional)
    pulpIdIn := []string{"Inner_example"} // []string | Multiple values may be separated by commas. (optional)
    repositoryVersion := "repositoryVersion_example" // string | Repository Version referenced by HREF (optional)
    repositoryVersionAdded := "repositoryVersionAdded_example" // string | Repository Version referenced by HREF (optional)
    repositoryVersionRemoved := "repositoryVersionRemoved_example" // string | Repository Version referenced by HREF (optional)
    suite := "suite_example" // string | Filter results where suite matches value (optional)
    fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
    excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContentReleaseArchitecturesAPI.ContentDebReleaseArchitecturesList(context.Background()).Architecture(architecture).Codename(codename).Distribution(distribution).Limit(limit).Offset(offset).Ordering(ordering).PulpHrefIn(pulpHrefIn).PulpIdIn(pulpIdIn).RepositoryVersion(repositoryVersion).RepositoryVersionAdded(repositoryVersionAdded).RepositoryVersionRemoved(repositoryVersionRemoved).Suite(suite).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContentReleaseArchitecturesAPI.ContentDebReleaseArchitecturesList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContentDebReleaseArchitecturesList`: PaginateddebReleaseArchitectureResponseList
    fmt.Fprintf(os.Stdout, "Response from `ContentReleaseArchitecturesAPI.ContentDebReleaseArchitecturesList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiContentDebReleaseArchitecturesListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **architecture** | **string** | Filter results where architecture matches value | 
 **codename** | **string** | Filter results where codename matches value | 
 **distribution** | **string** | Filter results where distribution matches value | 
 **limit** | **int32** | Number of results to return per page. | 
 **offset** | **int32** | The initial index from which to return the results. | 
 **ordering** | **[]string** | Ordering  * &#x60;pulp_id&#x60; - Pulp id * &#x60;-pulp_id&#x60; - Pulp id (descending) * &#x60;pulp_created&#x60; - Pulp created * &#x60;-pulp_created&#x60; - Pulp created (descending) * &#x60;pulp_last_updated&#x60; - Pulp last updated * &#x60;-pulp_last_updated&#x60; - Pulp last updated (descending) * &#x60;pulp_type&#x60; - Pulp type * &#x60;-pulp_type&#x60; - Pulp type (descending) * &#x60;upstream_id&#x60; - Upstream id * &#x60;-upstream_id&#x60; - Upstream id (descending) * &#x60;timestamp_of_interest&#x60; - Timestamp of interest * &#x60;-timestamp_of_interest&#x60; - Timestamp of interest (descending) * &#x60;architecture&#x60; - Architecture * &#x60;-architecture&#x60; - Architecture (descending) * &#x60;distribution&#x60; - Distribution * &#x60;-distribution&#x60; - Distribution (descending) * &#x60;codename&#x60; - Codename * &#x60;-codename&#x60; - Codename (descending) * &#x60;suite&#x60; - Suite * &#x60;-suite&#x60; - Suite (descending) * &#x60;pk&#x60; - Pk * &#x60;-pk&#x60; - Pk (descending) | 
 **pulpHrefIn** | **[]string** | Multiple values may be separated by commas. | 
 **pulpIdIn** | **[]string** | Multiple values may be separated by commas. | 
 **repositoryVersion** | **string** | Repository Version referenced by HREF | 
 **repositoryVersionAdded** | **string** | Repository Version referenced by HREF | 
 **repositoryVersionRemoved** | **string** | Repository Version referenced by HREF | 
 **suite** | **string** | Filter results where suite matches value | 
 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

### Return type

[**PaginateddebReleaseArchitectureResponseList**](PaginateddebReleaseArchitectureResponseList.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContentDebReleaseArchitecturesRead

> DebReleaseArchitectureResponse ContentDebReleaseArchitecturesRead(ctx, debReleaseArchitectureHref).Fields(fields).ExcludeFields(excludeFields).Execute()

Inspect a release architecture



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
    debReleaseArchitectureHref := "debReleaseArchitectureHref_example" // string | 
    fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
    excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContentReleaseArchitecturesAPI.ContentDebReleaseArchitecturesRead(context.Background(), debReleaseArchitectureHref).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContentReleaseArchitecturesAPI.ContentDebReleaseArchitecturesRead``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContentDebReleaseArchitecturesRead`: DebReleaseArchitectureResponse
    fmt.Fprintf(os.Stdout, "Response from `ContentReleaseArchitecturesAPI.ContentDebReleaseArchitecturesRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**debReleaseArchitectureHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiContentDebReleaseArchitecturesReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

### Return type

[**DebReleaseArchitectureResponse**](DebReleaseArchitectureResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

