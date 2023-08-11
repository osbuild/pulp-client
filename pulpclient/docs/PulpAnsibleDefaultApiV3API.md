# \PulpAnsibleDefaultApiV3API

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PulpAnsibleGalaxyDefaultApiV3Read**](PulpAnsibleDefaultApiV3API.md#PulpAnsibleGalaxyDefaultApiV3Read) | **Get** /pulp_ansible/galaxy/default/api/v3/ | 



## PulpAnsibleGalaxyDefaultApiV3Read

> RepoMetadataResponse PulpAnsibleGalaxyDefaultApiV3Read(ctx).Execute()





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
    resp, r, err := apiClient.PulpAnsibleDefaultApiV3API.PulpAnsibleGalaxyDefaultApiV3Read(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PulpAnsibleDefaultApiV3API.PulpAnsibleGalaxyDefaultApiV3Read``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PulpAnsibleGalaxyDefaultApiV3Read`: RepoMetadataResponse
    fmt.Fprintf(os.Stdout, "Response from `PulpAnsibleDefaultApiV3API.PulpAnsibleGalaxyDefaultApiV3Read`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPulpAnsibleGalaxyDefaultApiV3ReadRequest struct via the builder pattern


### Return type

[**RepoMetadataResponse**](RepoMetadataResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

