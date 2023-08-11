# \ImportersPulpImportCheckAPI

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PulpImportCheckPost**](ImportersPulpImportCheckAPI.md#PulpImportCheckPost) | **Post** /pulp/api/v3/importers/core/pulp/import-check/ | Validate the parameters to be used for a PulpImport call



## PulpImportCheckPost

> PulpImportCheckResponse PulpImportCheckPost(ctx).PulpImportCheck(pulpImportCheck).Execute()

Validate the parameters to be used for a PulpImport call



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
    pulpImportCheck := *openapiclient.NewPulpImportCheck() // PulpImportCheck | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ImportersPulpImportCheckAPI.PulpImportCheckPost(context.Background()).PulpImportCheck(pulpImportCheck).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImportersPulpImportCheckAPI.PulpImportCheckPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PulpImportCheckPost`: PulpImportCheckResponse
    fmt.Fprintf(os.Stdout, "Response from `ImportersPulpImportCheckAPI.PulpImportCheckPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPulpImportCheckPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pulpImportCheck** | [**PulpImportCheck**](PulpImportCheck.md) |  | 

### Return type

[**PulpImportCheckResponse**](PulpImportCheckResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

