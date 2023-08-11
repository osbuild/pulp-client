# \VersionsAPI

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV1RolesVersionsList**](VersionsAPI.md#ApiV1RolesVersionsList) | **Get** {ansible_role_href}versions/ | 
[**ApiV2CollectionVersionsList**](VersionsAPI.md#ApiV2CollectionVersionsList) | **Get** {ansible_collection_version_href}versions/ | 



## ApiV1RolesVersionsList

> PaginatedGalaxyRoleVersionResponseList ApiV1RolesVersionsList(ctx, ansibleRoleHref).Limit(limit).Offset(offset).Fields(fields).ExcludeFields(excludeFields).Execute()





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
    ansibleRoleHref := "ansibleRoleHref_example" // string | 
    limit := int32(56) // int32 | Number of results to return per page. (optional)
    offset := int32(56) // int32 | The initial index from which to return the results. (optional)
    fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
    excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VersionsAPI.ApiV1RolesVersionsList(context.Background(), ansibleRoleHref).Limit(limit).Offset(offset).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VersionsAPI.ApiV1RolesVersionsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV1RolesVersionsList`: PaginatedGalaxyRoleVersionResponseList
    fmt.Fprintf(os.Stdout, "Response from `VersionsAPI.ApiV1RolesVersionsList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ansibleRoleHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV1RolesVersionsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Number of results to return per page. | 
 **offset** | **int32** | The initial index from which to return the results. | 
 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

### Return type

[**PaginatedGalaxyRoleVersionResponseList**](PaginatedGalaxyRoleVersionResponseList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2CollectionVersionsList

> PaginatedGalaxyCollectionVersionResponseList ApiV2CollectionVersionsList(ctx, ansibleCollectionVersionHref).Page(page).Fields(fields).ExcludeFields(excludeFields).Execute()





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
    ansibleCollectionVersionHref := "ansibleCollectionVersionHref_example" // string | 
    page := int32(56) // int32 | A page number within the paginated result set. (optional)
    fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
    excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VersionsAPI.ApiV2CollectionVersionsList(context.Background(), ansibleCollectionVersionHref).Page(page).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VersionsAPI.ApiV2CollectionVersionsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV2CollectionVersionsList`: PaginatedGalaxyCollectionVersionResponseList
    fmt.Fprintf(os.Stdout, "Response from `VersionsAPI.ApiV2CollectionVersionsList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ansibleCollectionVersionHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2CollectionVersionsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | A page number within the paginated result set. | 
 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

### Return type

[**PaginatedGalaxyCollectionVersionResponseList**](PaginatedGalaxyCollectionVersionResponseList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

