# \PulpAnsibleApiV3PluginAnsibleContentCollectionsIndexVersionsDocsBlobAPI

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PulpAnsibleGalaxyApiV3PluginAnsibleContentCollectionsIndexVersionsDocsBlobRead**](PulpAnsibleApiV3PluginAnsibleContentCollectionsIndexVersionsDocsBlobAPI.md#PulpAnsibleGalaxyApiV3PluginAnsibleContentCollectionsIndexVersionsDocsBlobRead) | **Get** /pulp_ansible/galaxy/{path}/api/v3/plugin/ansible/content/{distro_base_path}/collections/index/{namespace}/{name}/versions/{version}/docs-blob/ | 



## PulpAnsibleGalaxyApiV3PluginAnsibleContentCollectionsIndexVersionsDocsBlobRead

> CollectionVersionDocsResponse PulpAnsibleGalaxyApiV3PluginAnsibleContentCollectionsIndexVersionsDocsBlobRead(ctx, distroBasePath, name, namespace, path, version).Fields(fields).ExcludeFields(excludeFields).Execute()





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
    version := "version_example" // string | 
    fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
    excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PulpAnsibleApiV3PluginAnsibleContentCollectionsIndexVersionsDocsBlobAPI.PulpAnsibleGalaxyApiV3PluginAnsibleContentCollectionsIndexVersionsDocsBlobRead(context.Background(), distroBasePath, name, namespace, path, version).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PulpAnsibleApiV3PluginAnsibleContentCollectionsIndexVersionsDocsBlobAPI.PulpAnsibleGalaxyApiV3PluginAnsibleContentCollectionsIndexVersionsDocsBlobRead``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PulpAnsibleGalaxyApiV3PluginAnsibleContentCollectionsIndexVersionsDocsBlobRead`: CollectionVersionDocsResponse
    fmt.Fprintf(os.Stdout, "Response from `PulpAnsibleApiV3PluginAnsibleContentCollectionsIndexVersionsDocsBlobAPI.PulpAnsibleGalaxyApiV3PluginAnsibleContentCollectionsIndexVersionsDocsBlobRead`: %v\n", resp)
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
**version** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPulpAnsibleGalaxyApiV3PluginAnsibleContentCollectionsIndexVersionsDocsBlobReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

### Return type

[**CollectionVersionDocsResponse**](CollectionVersionDocsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

