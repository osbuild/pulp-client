# \PulpAnsibleDefaultApiV3CollectionsAPI

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PulpAnsibleGalaxyDefaultApiV3CollectionsDelete**](PulpAnsibleDefaultApiV3CollectionsAPI.md#PulpAnsibleGalaxyDefaultApiV3CollectionsDelete) | **Delete** /pulp_ansible/galaxy/default/api/v3/collections/{namespace}/{name}/ | 
[**PulpAnsibleGalaxyDefaultApiV3CollectionsList**](PulpAnsibleDefaultApiV3CollectionsAPI.md#PulpAnsibleGalaxyDefaultApiV3CollectionsList) | **Get** /pulp_ansible/galaxy/default/api/v3/collections/ | 
[**PulpAnsibleGalaxyDefaultApiV3CollectionsRead**](PulpAnsibleDefaultApiV3CollectionsAPI.md#PulpAnsibleGalaxyDefaultApiV3CollectionsRead) | **Get** /pulp_ansible/galaxy/default/api/v3/collections/{namespace}/{name}/ | 
[**PulpAnsibleGalaxyDefaultApiV3CollectionsUpdate**](PulpAnsibleDefaultApiV3CollectionsAPI.md#PulpAnsibleGalaxyDefaultApiV3CollectionsUpdate) | **Patch** /pulp_ansible/galaxy/default/api/v3/collections/{namespace}/{name}/ | 



## PulpAnsibleGalaxyDefaultApiV3CollectionsDelete

> AsyncOperationResponse PulpAnsibleGalaxyDefaultApiV3CollectionsDelete(ctx, name, namespace).Execute()





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
    name := "name_example" // string | 
    namespace := "namespace_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PulpAnsibleDefaultApiV3CollectionsAPI.PulpAnsibleGalaxyDefaultApiV3CollectionsDelete(context.Background(), name, namespace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PulpAnsibleDefaultApiV3CollectionsAPI.PulpAnsibleGalaxyDefaultApiV3CollectionsDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PulpAnsibleGalaxyDefaultApiV3CollectionsDelete`: AsyncOperationResponse
    fmt.Fprintf(os.Stdout, "Response from `PulpAnsibleDefaultApiV3CollectionsAPI.PulpAnsibleGalaxyDefaultApiV3CollectionsDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 
**namespace** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPulpAnsibleGalaxyDefaultApiV3CollectionsDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**AsyncOperationResponse**](AsyncOperationResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PulpAnsibleGalaxyDefaultApiV3CollectionsList

> PaginatedCollectionResponseList PulpAnsibleGalaxyDefaultApiV3CollectionsList(ctx).Deprecated(deprecated).Limit(limit).Name(name).Namespace(namespace).Offset(offset).Ordering(ordering).PulpHrefIn(pulpHrefIn).PulpIdIn(pulpIdIn).Fields(fields).ExcludeFields(excludeFields).Execute()





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
    deprecated := true // bool |  (optional)
    limit := int32(56) // int32 | Number of results to return per page. (optional)
    name := "name_example" // string |  (optional)
    namespace := "namespace_example" // string |  (optional)
    offset := int32(56) // int32 | The initial index from which to return the results. (optional)
    ordering := []string{"Ordering_example"} // []string | Ordering  * `pulp_id` - Pulp id * `-pulp_id` - Pulp id (descending) * `pulp_created` - Pulp created * `-pulp_created` - Pulp created (descending) * `pulp_last_updated` - Pulp last updated * `-pulp_last_updated` - Pulp last updated (descending) * `namespace` - Namespace * `-namespace` - Namespace (descending) * `name` - Name * `-name` - Name (descending) * `pk` - Pk * `-pk` - Pk (descending) (optional)
    pulpHrefIn := []string{"Inner_example"} // []string | Multiple values may be separated by commas. (optional)
    pulpIdIn := []string{"Inner_example"} // []string | Multiple values may be separated by commas. (optional)
    fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
    excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PulpAnsibleDefaultApiV3CollectionsAPI.PulpAnsibleGalaxyDefaultApiV3CollectionsList(context.Background()).Deprecated(deprecated).Limit(limit).Name(name).Namespace(namespace).Offset(offset).Ordering(ordering).PulpHrefIn(pulpHrefIn).PulpIdIn(pulpIdIn).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PulpAnsibleDefaultApiV3CollectionsAPI.PulpAnsibleGalaxyDefaultApiV3CollectionsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PulpAnsibleGalaxyDefaultApiV3CollectionsList`: PaginatedCollectionResponseList
    fmt.Fprintf(os.Stdout, "Response from `PulpAnsibleDefaultApiV3CollectionsAPI.PulpAnsibleGalaxyDefaultApiV3CollectionsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPulpAnsibleGalaxyDefaultApiV3CollectionsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deprecated** | **bool** |  | 
 **limit** | **int32** | Number of results to return per page. | 
 **name** | **string** |  | 
 **namespace** | **string** |  | 
 **offset** | **int32** | The initial index from which to return the results. | 
 **ordering** | **[]string** | Ordering  * &#x60;pulp_id&#x60; - Pulp id * &#x60;-pulp_id&#x60; - Pulp id (descending) * &#x60;pulp_created&#x60; - Pulp created * &#x60;-pulp_created&#x60; - Pulp created (descending) * &#x60;pulp_last_updated&#x60; - Pulp last updated * &#x60;-pulp_last_updated&#x60; - Pulp last updated (descending) * &#x60;namespace&#x60; - Namespace * &#x60;-namespace&#x60; - Namespace (descending) * &#x60;name&#x60; - Name * &#x60;-name&#x60; - Name (descending) * &#x60;pk&#x60; - Pk * &#x60;-pk&#x60; - Pk (descending) | 
 **pulpHrefIn** | **[]string** | Multiple values may be separated by commas. | 
 **pulpIdIn** | **[]string** | Multiple values may be separated by commas. | 
 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

### Return type

[**PaginatedCollectionResponseList**](PaginatedCollectionResponseList.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PulpAnsibleGalaxyDefaultApiV3CollectionsRead

> CollectionResponse PulpAnsibleGalaxyDefaultApiV3CollectionsRead(ctx, name, namespace).Fields(fields).ExcludeFields(excludeFields).Execute()





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
    name := "name_example" // string | 
    namespace := "namespace_example" // string | 
    fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
    excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PulpAnsibleDefaultApiV3CollectionsAPI.PulpAnsibleGalaxyDefaultApiV3CollectionsRead(context.Background(), name, namespace).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PulpAnsibleDefaultApiV3CollectionsAPI.PulpAnsibleGalaxyDefaultApiV3CollectionsRead``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PulpAnsibleGalaxyDefaultApiV3CollectionsRead`: CollectionResponse
    fmt.Fprintf(os.Stdout, "Response from `PulpAnsibleDefaultApiV3CollectionsAPI.PulpAnsibleGalaxyDefaultApiV3CollectionsRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 
**namespace** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPulpAnsibleGalaxyDefaultApiV3CollectionsReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

### Return type

[**CollectionResponse**](CollectionResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PulpAnsibleGalaxyDefaultApiV3CollectionsUpdate

> AsyncOperationResponse PulpAnsibleGalaxyDefaultApiV3CollectionsUpdate(ctx, name, namespace).PatchedCollection(patchedCollection).Execute()





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
    name := "name_example" // string | 
    namespace := "namespace_example" // string | 
    patchedCollection := *openapiclient.NewPatchedCollection() // PatchedCollection | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PulpAnsibleDefaultApiV3CollectionsAPI.PulpAnsibleGalaxyDefaultApiV3CollectionsUpdate(context.Background(), name, namespace).PatchedCollection(patchedCollection).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PulpAnsibleDefaultApiV3CollectionsAPI.PulpAnsibleGalaxyDefaultApiV3CollectionsUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PulpAnsibleGalaxyDefaultApiV3CollectionsUpdate`: AsyncOperationResponse
    fmt.Fprintf(os.Stdout, "Response from `PulpAnsibleDefaultApiV3CollectionsAPI.PulpAnsibleGalaxyDefaultApiV3CollectionsUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 
**namespace** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPulpAnsibleGalaxyDefaultApiV3CollectionsUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **patchedCollection** | [**PatchedCollection**](PatchedCollection.md) |  | 

### Return type

[**AsyncOperationResponse**](AsyncOperationResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

