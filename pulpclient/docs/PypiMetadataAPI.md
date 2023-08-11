# \PypiMetadataAPI

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PypiPypiRead**](PypiMetadataAPI.md#PypiPypiRead) | **Get** /pypi/{path}/pypi/{meta}/ | Get package metadata



## PypiPypiRead

> PackageMetadataResponse PypiPypiRead(ctx, meta, path).Fields(fields).ExcludeFields(excludeFields).Execute()

Get package metadata



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
    meta := "meta_example" // string | 
    path := "path_example" // string | 
    fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
    excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PypiMetadataAPI.PypiPypiRead(context.Background(), meta, path).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PypiMetadataAPI.PypiPypiRead``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PypiPypiRead`: PackageMetadataResponse
    fmt.Fprintf(os.Stdout, "Response from `PypiMetadataAPI.PypiPypiRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**meta** | **string** |  | 
**path** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPypiPypiReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

### Return type

[**PackageMetadataResponse**](PackageMetadataResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

