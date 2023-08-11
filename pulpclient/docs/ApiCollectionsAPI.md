# \ApiCollectionsAPI

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV2CollectionsGet**](ApiCollectionsAPI.md#ApiV2CollectionsGet) | **Get** {ansible_collection_href}api/v2/collections/ | 
[**ApiV2CollectionsPost**](ApiCollectionsAPI.md#ApiV2CollectionsPost) | **Post** {ansible_collection_href}api/v2/collections/ | 



## ApiV2CollectionsGet

> PaginatedGalaxyCollectionResponseList ApiV2CollectionsGet(ctx, ansibleCollectionHref).Page(page).Fields(fields).ExcludeFields(excludeFields).Execute()





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
    ansibleCollectionHref := "ansibleCollectionHref_example" // string | 
    page := int32(56) // int32 | A page number within the paginated result set. (optional)
    fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
    excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApiCollectionsAPI.ApiV2CollectionsGet(context.Background(), ansibleCollectionHref).Page(page).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApiCollectionsAPI.ApiV2CollectionsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV2CollectionsGet`: PaginatedGalaxyCollectionResponseList
    fmt.Fprintf(os.Stdout, "Response from `ApiCollectionsAPI.ApiV2CollectionsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ansibleCollectionHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2CollectionsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | A page number within the paginated result set. | 
 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

### Return type

[**PaginatedGalaxyCollectionResponseList**](PaginatedGalaxyCollectionResponseList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2CollectionsPost

> GalaxyCollectionResponse ApiV2CollectionsPost(ctx, ansibleCollectionHref).GalaxyCollection(galaxyCollection).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "github.com/osbuild/pulp-client/pulpclient"
)

func main() {
    ansibleCollectionHref := "ansibleCollectionHref_example" // string | 
    galaxyCollection := *openapiclient.NewGalaxyCollection("Id_example", "Name_example", time.Now(), time.Now()) // GalaxyCollection | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApiCollectionsAPI.ApiV2CollectionsPost(context.Background(), ansibleCollectionHref).GalaxyCollection(galaxyCollection).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApiCollectionsAPI.ApiV2CollectionsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV2CollectionsPost`: GalaxyCollectionResponse
    fmt.Fprintf(os.Stdout, "Response from `ApiCollectionsAPI.ApiV2CollectionsPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ansibleCollectionHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2CollectionsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **galaxyCollection** | [**GalaxyCollection**](GalaxyCollection.md) |  | 

### Return type

[**GalaxyCollectionResponse**](GalaxyCollectionResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

