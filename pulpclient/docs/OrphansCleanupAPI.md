# \OrphansCleanupAPI

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OrphansCleanupCleanup**](OrphansCleanupAPI.md#OrphansCleanupCleanup) | **Post** /pulp/api/v3/orphans/cleanup/ | 



## OrphansCleanupCleanup

> AsyncOperationResponse OrphansCleanupCleanup(ctx).OrphansCleanup(orphansCleanup).Execute()





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
    orphansCleanup := *openapiclient.NewOrphansCleanup() // OrphansCleanup | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrphansCleanupAPI.OrphansCleanupCleanup(context.Background()).OrphansCleanup(orphansCleanup).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrphansCleanupAPI.OrphansCleanupCleanup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrphansCleanupCleanup`: AsyncOperationResponse
    fmt.Fprintf(os.Stdout, "Response from `OrphansCleanupAPI.OrphansCleanupCleanup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOrphansCleanupCleanupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **orphansCleanup** | [**OrphansCleanup**](OrphansCleanup.md) |  | 

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

