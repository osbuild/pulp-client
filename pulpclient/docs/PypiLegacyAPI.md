# \PypiLegacyAPI

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PypiLegacyCreate**](PypiLegacyAPI.md#PypiLegacyCreate) | **Post** /pypi/{path}/legacy/ | Upload a package



## PypiLegacyCreate

> PackageUploadTaskResponse PypiLegacyCreate(ctx, path).Content(content).Sha256Digest(sha256Digest).Action(action).Execute()

Upload a package



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
    path := "path_example" // string | 
    content := os.NewFile(1234, "some_file") // *os.File | A Python package release file to upload to the index.
    sha256Digest := "sha256Digest_example" // string | SHA256 of package to validate upload integrity.
    action := "action_example" // string | Defaults to `file_upload`, don't change it or request will fail! (optional) (default to "file_upload")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PypiLegacyAPI.PypiLegacyCreate(context.Background(), path).Content(content).Sha256Digest(sha256Digest).Action(action).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PypiLegacyAPI.PypiLegacyCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PypiLegacyCreate`: PackageUploadTaskResponse
    fmt.Fprintf(os.Stdout, "Response from `PypiLegacyAPI.PypiLegacyCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**path** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPypiLegacyCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **content** | ***os.File** | A Python package release file to upload to the index. | 
 **sha256Digest** | **string** | SHA256 of package to validate upload integrity. | 
 **action** | **string** | Defaults to &#x60;file_upload&#x60;, don&#39;t change it or request will fail! | [default to &quot;file_upload&quot;]

### Return type

[**PackageUploadTaskResponse**](PackageUploadTaskResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: multipart/form-data, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

