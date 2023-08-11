# \PulpAnsibleDefaultApiV3PluginAnsibleContentCollectionsIndexAPI

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsIndexDelete**](PulpAnsibleDefaultApiV3PluginAnsibleContentCollectionsIndexAPI.md#PulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsIndexDelete) | **Delete** /pulp_ansible/galaxy/default/api/v3/plugin/ansible/content/{distro_base_path}/collections/index/{namespace}/{name}/ | 
[**PulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsIndexList**](PulpAnsibleDefaultApiV3PluginAnsibleContentCollectionsIndexAPI.md#PulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsIndexList) | **Get** /pulp_ansible/galaxy/default/api/v3/plugin/ansible/content/{distro_base_path}/collections/index/ | 
[**PulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsIndexRead**](PulpAnsibleDefaultApiV3PluginAnsibleContentCollectionsIndexAPI.md#PulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsIndexRead) | **Get** /pulp_ansible/galaxy/default/api/v3/plugin/ansible/content/{distro_base_path}/collections/index/{namespace}/{name}/ | 
[**PulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsIndexUpdate**](PulpAnsibleDefaultApiV3PluginAnsibleContentCollectionsIndexAPI.md#PulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsIndexUpdate) | **Patch** /pulp_ansible/galaxy/default/api/v3/plugin/ansible/content/{distro_base_path}/collections/index/{namespace}/{name}/ | 



## PulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsIndexDelete

> AsyncOperationResponse PulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsIndexDelete(ctx, distroBasePath, name, namespace).Execute()





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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PulpAnsibleDefaultApiV3PluginAnsibleContentCollectionsIndexAPI.PulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsIndexDelete(context.Background(), distroBasePath, name, namespace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PulpAnsibleDefaultApiV3PluginAnsibleContentCollectionsIndexAPI.PulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsIndexDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsIndexDelete`: AsyncOperationResponse
    fmt.Fprintf(os.Stdout, "Response from `PulpAnsibleDefaultApiV3PluginAnsibleContentCollectionsIndexAPI.PulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsIndexDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**distroBasePath** | **string** |  | 
**name** | **string** |  | 
**namespace** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsIndexDeleteRequest struct via the builder pattern


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


## PulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsIndexList

> PaginatedCollectionResponseList PulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsIndexList(ctx, distroBasePath).Deprecated(deprecated).Limit(limit).Name(name).Namespace(namespace).Offset(offset).Ordering(ordering).PulpHrefIn(pulpHrefIn).PulpIdIn(pulpIdIn).Fields(fields).ExcludeFields(excludeFields).Execute()





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
    resp, r, err := apiClient.PulpAnsibleDefaultApiV3PluginAnsibleContentCollectionsIndexAPI.PulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsIndexList(context.Background(), distroBasePath).Deprecated(deprecated).Limit(limit).Name(name).Namespace(namespace).Offset(offset).Ordering(ordering).PulpHrefIn(pulpHrefIn).PulpIdIn(pulpIdIn).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PulpAnsibleDefaultApiV3PluginAnsibleContentCollectionsIndexAPI.PulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsIndexList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsIndexList`: PaginatedCollectionResponseList
    fmt.Fprintf(os.Stdout, "Response from `PulpAnsibleDefaultApiV3PluginAnsibleContentCollectionsIndexAPI.PulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsIndexList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**distroBasePath** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsIndexListRequest struct via the builder pattern


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


## PulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsIndexRead

> CollectionResponse PulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsIndexRead(ctx, distroBasePath, name, namespace).Fields(fields).ExcludeFields(excludeFields).Execute()





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
    fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
    excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PulpAnsibleDefaultApiV3PluginAnsibleContentCollectionsIndexAPI.PulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsIndexRead(context.Background(), distroBasePath, name, namespace).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PulpAnsibleDefaultApiV3PluginAnsibleContentCollectionsIndexAPI.PulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsIndexRead``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsIndexRead`: CollectionResponse
    fmt.Fprintf(os.Stdout, "Response from `PulpAnsibleDefaultApiV3PluginAnsibleContentCollectionsIndexAPI.PulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsIndexRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**distroBasePath** | **string** |  | 
**name** | **string** |  | 
**namespace** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsIndexReadRequest struct via the builder pattern


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


## PulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsIndexUpdate

> AsyncOperationResponse PulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsIndexUpdate(ctx, distroBasePath, name, namespace).PatchedCollection(patchedCollection).Execute()





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
    patchedCollection := *openapiclient.NewPatchedCollection() // PatchedCollection | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PulpAnsibleDefaultApiV3PluginAnsibleContentCollectionsIndexAPI.PulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsIndexUpdate(context.Background(), distroBasePath, name, namespace).PatchedCollection(patchedCollection).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PulpAnsibleDefaultApiV3PluginAnsibleContentCollectionsIndexAPI.PulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsIndexUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsIndexUpdate`: AsyncOperationResponse
    fmt.Fprintf(os.Stdout, "Response from `PulpAnsibleDefaultApiV3PluginAnsibleContentCollectionsIndexAPI.PulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsIndexUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**distroBasePath** | **string** |  | 
**name** | **string** |  | 
**namespace** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsIndexUpdateRequest struct via the builder pattern


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

