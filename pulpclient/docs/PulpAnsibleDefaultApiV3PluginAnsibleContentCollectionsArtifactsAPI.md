# \PulpAnsibleDefaultApiV3PluginAnsibleContentCollectionsArtifactsAPI

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsArtifactsDownload**](PulpAnsibleDefaultApiV3PluginAnsibleContentCollectionsArtifactsAPI.md#PulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsArtifactsDownload) | **Get** /pulp_ansible/galaxy/default/api/v3/plugin/ansible/content/{distro_base_path}/collections/artifacts/{filename} | 



## PulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsArtifactsDownload

> PulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsArtifactsDownload(ctx, distroBasePath, filename).Fields(fields).ExcludeFields(excludeFields).Execute()





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
    filename := "filename_example" // string | 
    fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
    excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PulpAnsibleDefaultApiV3PluginAnsibleContentCollectionsArtifactsAPI.PulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsArtifactsDownload(context.Background(), distroBasePath, filename).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PulpAnsibleDefaultApiV3PluginAnsibleContentCollectionsArtifactsAPI.PulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsArtifactsDownload``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**distroBasePath** | **string** |  | 
**filename** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsArtifactsDownloadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

