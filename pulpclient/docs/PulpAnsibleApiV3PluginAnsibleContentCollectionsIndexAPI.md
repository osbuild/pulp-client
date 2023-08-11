# \PulpAnsibleApiV3PluginAnsibleContentCollectionsIndexAPI

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PulpAnsibleGalaxyApiV3PluginAnsibleContentCollectionsIndexDelete**](PulpAnsibleApiV3PluginAnsibleContentCollectionsIndexAPI.md#PulpAnsibleGalaxyApiV3PluginAnsibleContentCollectionsIndexDelete) | **Delete** /pulp_ansible/galaxy/{path}/api/v3/plugin/ansible/content/{distro_base_path}/collections/index/{namespace}/{name}/ | 
[**PulpAnsibleGalaxyApiV3PluginAnsibleContentCollectionsIndexList**](PulpAnsibleApiV3PluginAnsibleContentCollectionsIndexAPI.md#PulpAnsibleGalaxyApiV3PluginAnsibleContentCollectionsIndexList) | **Get** /pulp_ansible/galaxy/{path}/api/v3/plugin/ansible/content/{distro_base_path}/collections/index/ | 
[**PulpAnsibleGalaxyApiV3PluginAnsibleContentCollectionsIndexRead**](PulpAnsibleApiV3PluginAnsibleContentCollectionsIndexAPI.md#PulpAnsibleGalaxyApiV3PluginAnsibleContentCollectionsIndexRead) | **Get** /pulp_ansible/galaxy/{path}/api/v3/plugin/ansible/content/{distro_base_path}/collections/index/{namespace}/{name}/ | 
[**PulpAnsibleGalaxyApiV3PluginAnsibleContentCollectionsIndexUpdate**](PulpAnsibleApiV3PluginAnsibleContentCollectionsIndexAPI.md#PulpAnsibleGalaxyApiV3PluginAnsibleContentCollectionsIndexUpdate) | **Patch** /pulp_ansible/galaxy/{path}/api/v3/plugin/ansible/content/{distro_base_path}/collections/index/{namespace}/{name}/ | 



## PulpAnsibleGalaxyApiV3PluginAnsibleContentCollectionsIndexDelete

> AsyncOperationResponse PulpAnsibleGalaxyApiV3PluginAnsibleContentCollectionsIndexDelete(ctx, distroBasePath, name, namespace, path).Execute()





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
    distroBasePath := "distroBasePath_example" // string | 
    name := "name_example" // string | 
    namespace := "namespace_example" // string | 
    path := "path_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PulpAnsibleApiV3PluginAnsibleContentCollectionsIndexAPI.PulpAnsibleGalaxyApiV3PluginAnsibleContentCollectionsIndexDelete(context.Background(), distroBasePath, name, namespace, path).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PulpAnsibleApiV3PluginAnsibleContentCollectionsIndexAPI.PulpAnsibleGalaxyApiV3PluginAnsibleContentCollectionsIndexDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PulpAnsibleGalaxyApiV3PluginAnsibleContentCollectionsIndexDelete`: AsyncOperationResponse
    fmt.Fprintf(os.Stdout, "Response from `PulpAnsibleApiV3PluginAnsibleContentCollectionsIndexAPI.PulpAnsibleGalaxyApiV3PluginAnsibleContentCollectionsIndexDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**distroBasePath** | **string** |  | 
**name** | **string** |  | 
**namespace** | **string** |  | 
**path** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPulpAnsibleGalaxyApiV3PluginAnsibleContentCollectionsIndexDeleteRequest struct via the builder pattern


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


## PulpAnsibleGalaxyApiV3PluginAnsibleContentCollectionsIndexList

> PaginatedCollectionResponseList PulpAnsibleGalaxyApiV3PluginAnsibleContentCollectionsIndexList(ctx, distroBasePath, path).Deprecated(deprecated).Limit(limit).Name(name).Namespace(namespace).Offset(offset).Ordering(ordering).PulpHrefIn(pulpHrefIn).PulpIdIn(pulpIdIn).Fields(fields).ExcludeFields(excludeFields).Execute()





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
    distroBasePath := "distroBasePath_example" // string | 
    path := "path_example" // string | 
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
    resp, r, err := apiClient.PulpAnsibleApiV3PluginAnsibleContentCollectionsIndexAPI.PulpAnsibleGalaxyApiV3PluginAnsibleContentCollectionsIndexList(context.Background(), distroBasePath, path).Deprecated(deprecated).Limit(limit).Name(name).Namespace(namespace).Offset(offset).Ordering(ordering).PulpHrefIn(pulpHrefIn).PulpIdIn(pulpIdIn).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PulpAnsibleApiV3PluginAnsibleContentCollectionsIndexAPI.PulpAnsibleGalaxyApiV3PluginAnsibleContentCollectionsIndexList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PulpAnsibleGalaxyApiV3PluginAnsibleContentCollectionsIndexList`: PaginatedCollectionResponseList
    fmt.Fprintf(os.Stdout, "Response from `PulpAnsibleApiV3PluginAnsibleContentCollectionsIndexAPI.PulpAnsibleGalaxyApiV3PluginAnsibleContentCollectionsIndexList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**distroBasePath** | **string** |  | 
**path** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPulpAnsibleGalaxyApiV3PluginAnsibleContentCollectionsIndexListRequest struct via the builder pattern


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


## PulpAnsibleGalaxyApiV3PluginAnsibleContentCollectionsIndexRead

> CollectionResponse PulpAnsibleGalaxyApiV3PluginAnsibleContentCollectionsIndexRead(ctx, distroBasePath, name, namespace, path).Fields(fields).ExcludeFields(excludeFields).Execute()





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
    distroBasePath := "distroBasePath_example" // string | 
    name := "name_example" // string | 
    namespace := "namespace_example" // string | 
    path := "path_example" // string | 
    fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
    excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PulpAnsibleApiV3PluginAnsibleContentCollectionsIndexAPI.PulpAnsibleGalaxyApiV3PluginAnsibleContentCollectionsIndexRead(context.Background(), distroBasePath, name, namespace, path).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PulpAnsibleApiV3PluginAnsibleContentCollectionsIndexAPI.PulpAnsibleGalaxyApiV3PluginAnsibleContentCollectionsIndexRead``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PulpAnsibleGalaxyApiV3PluginAnsibleContentCollectionsIndexRead`: CollectionResponse
    fmt.Fprintf(os.Stdout, "Response from `PulpAnsibleApiV3PluginAnsibleContentCollectionsIndexAPI.PulpAnsibleGalaxyApiV3PluginAnsibleContentCollectionsIndexRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**distroBasePath** | **string** |  | 
**name** | **string** |  | 
**namespace** | **string** |  | 
**path** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPulpAnsibleGalaxyApiV3PluginAnsibleContentCollectionsIndexReadRequest struct via the builder pattern


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


## PulpAnsibleGalaxyApiV3PluginAnsibleContentCollectionsIndexUpdate

> AsyncOperationResponse PulpAnsibleGalaxyApiV3PluginAnsibleContentCollectionsIndexUpdate(ctx, distroBasePath, name, namespace, path).PatchedCollection(patchedCollection).Execute()





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
    distroBasePath := "distroBasePath_example" // string | 
    name := "name_example" // string | 
    namespace := "namespace_example" // string | 
    path := "path_example" // string | 
    patchedCollection := *openapiclient.NewPatchedCollection() // PatchedCollection | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PulpAnsibleApiV3PluginAnsibleContentCollectionsIndexAPI.PulpAnsibleGalaxyApiV3PluginAnsibleContentCollectionsIndexUpdate(context.Background(), distroBasePath, name, namespace, path).PatchedCollection(patchedCollection).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PulpAnsibleApiV3PluginAnsibleContentCollectionsIndexAPI.PulpAnsibleGalaxyApiV3PluginAnsibleContentCollectionsIndexUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PulpAnsibleGalaxyApiV3PluginAnsibleContentCollectionsIndexUpdate`: AsyncOperationResponse
    fmt.Fprintf(os.Stdout, "Response from `PulpAnsibleApiV3PluginAnsibleContentCollectionsIndexAPI.PulpAnsibleGalaxyApiV3PluginAnsibleContentCollectionsIndexUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**distroBasePath** | **string** |  | 
**name** | **string** |  | 
**namespace** | **string** |  | 
**path** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPulpAnsibleGalaxyApiV3PluginAnsibleContentCollectionsIndexUpdateRequest struct via the builder pattern


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

