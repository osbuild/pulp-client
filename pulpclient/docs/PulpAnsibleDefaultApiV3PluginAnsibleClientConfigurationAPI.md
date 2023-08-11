# \PulpAnsibleDefaultApiV3PluginAnsibleClientConfigurationAPI

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PulpAnsibleGalaxyDefaultApiV3PluginAnsibleClientConfigurationRead**](PulpAnsibleDefaultApiV3PluginAnsibleClientConfigurationAPI.md#PulpAnsibleGalaxyDefaultApiV3PluginAnsibleClientConfigurationRead) | **Get** /pulp_ansible/galaxy/default/api/v3/plugin/ansible/client-configuration/ | 



## PulpAnsibleGalaxyDefaultApiV3PluginAnsibleClientConfigurationRead

> ClientConfigurationResponse PulpAnsibleGalaxyDefaultApiV3PluginAnsibleClientConfigurationRead(ctx).Execute()





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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PulpAnsibleDefaultApiV3PluginAnsibleClientConfigurationAPI.PulpAnsibleGalaxyDefaultApiV3PluginAnsibleClientConfigurationRead(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PulpAnsibleDefaultApiV3PluginAnsibleClientConfigurationAPI.PulpAnsibleGalaxyDefaultApiV3PluginAnsibleClientConfigurationRead``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PulpAnsibleGalaxyDefaultApiV3PluginAnsibleClientConfigurationRead`: ClientConfigurationResponse
    fmt.Fprintf(os.Stdout, "Response from `PulpAnsibleDefaultApiV3PluginAnsibleClientConfigurationAPI.PulpAnsibleGalaxyDefaultApiV3PluginAnsibleClientConfigurationRead`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPulpAnsibleGalaxyDefaultApiV3PluginAnsibleClientConfigurationReadRequest struct via the builder pattern


### Return type

[**ClientConfigurationResponse**](ClientConfigurationResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

