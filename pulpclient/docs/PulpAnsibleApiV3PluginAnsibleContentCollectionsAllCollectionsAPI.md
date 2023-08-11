# \PulpAnsibleApiV3PluginAnsibleContentCollectionsAllCollectionsAPI

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PulpAnsibleGalaxyApiV3PluginAnsibleContentCollectionsAllCollectionsList**](PulpAnsibleApiV3PluginAnsibleContentCollectionsAllCollectionsAPI.md#PulpAnsibleGalaxyApiV3PluginAnsibleContentCollectionsAllCollectionsList) | **Get** /pulp_ansible/galaxy/{path}/api/v3/plugin/ansible/content/{distro_base_path}/collections/all-collections/ | 



## PulpAnsibleGalaxyApiV3PluginAnsibleContentCollectionsAllCollectionsList

> []CollectionResponse PulpAnsibleGalaxyApiV3PluginAnsibleContentCollectionsAllCollectionsList(ctx, distroBasePath, path).Deprecated(deprecated).Name(name).Namespace(namespace).Ordering(ordering).PulpHrefIn(pulpHrefIn).PulpIdIn(pulpIdIn).Fields(fields).ExcludeFields(excludeFields).Execute()





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
    name := "name_example" // string |  (optional)
    namespace := "namespace_example" // string |  (optional)
    ordering := []string{"Ordering_example"} // []string | Ordering  * `pulp_id` - Pulp id * `-pulp_id` - Pulp id (descending) * `pulp_created` - Pulp created * `-pulp_created` - Pulp created (descending) * `pulp_last_updated` - Pulp last updated * `-pulp_last_updated` - Pulp last updated (descending) * `namespace` - Namespace * `-namespace` - Namespace (descending) * `name` - Name * `-name` - Name (descending) * `pk` - Pk * `-pk` - Pk (descending) (optional)
    pulpHrefIn := []string{"Inner_example"} // []string | Multiple values may be separated by commas. (optional)
    pulpIdIn := []string{"Inner_example"} // []string | Multiple values may be separated by commas. (optional)
    fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
    excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PulpAnsibleApiV3PluginAnsibleContentCollectionsAllCollectionsAPI.PulpAnsibleGalaxyApiV3PluginAnsibleContentCollectionsAllCollectionsList(context.Background(), distroBasePath, path).Deprecated(deprecated).Name(name).Namespace(namespace).Ordering(ordering).PulpHrefIn(pulpHrefIn).PulpIdIn(pulpIdIn).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PulpAnsibleApiV3PluginAnsibleContentCollectionsAllCollectionsAPI.PulpAnsibleGalaxyApiV3PluginAnsibleContentCollectionsAllCollectionsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PulpAnsibleGalaxyApiV3PluginAnsibleContentCollectionsAllCollectionsList`: []CollectionResponse
    fmt.Fprintf(os.Stdout, "Response from `PulpAnsibleApiV3PluginAnsibleContentCollectionsAllCollectionsAPI.PulpAnsibleGalaxyApiV3PluginAnsibleContentCollectionsAllCollectionsList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**distroBasePath** | **string** |  | 
**path** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPulpAnsibleGalaxyApiV3PluginAnsibleContentCollectionsAllCollectionsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **deprecated** | **bool** |  | 
 **name** | **string** |  | 
 **namespace** | **string** |  | 
 **ordering** | **[]string** | Ordering  * &#x60;pulp_id&#x60; - Pulp id * &#x60;-pulp_id&#x60; - Pulp id (descending) * &#x60;pulp_created&#x60; - Pulp created * &#x60;-pulp_created&#x60; - Pulp created (descending) * &#x60;pulp_last_updated&#x60; - Pulp last updated * &#x60;-pulp_last_updated&#x60; - Pulp last updated (descending) * &#x60;namespace&#x60; - Namespace * &#x60;-namespace&#x60; - Namespace (descending) * &#x60;name&#x60; - Name * &#x60;-name&#x60; - Name (descending) * &#x60;pk&#x60; - Pk * &#x60;-pk&#x60; - Pk (descending) | 
 **pulpHrefIn** | **[]string** | Multiple values may be separated by commas. | 
 **pulpIdIn** | **[]string** | Multiple values may be separated by commas. | 
 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

### Return type

[**[]CollectionResponse**](CollectionResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

