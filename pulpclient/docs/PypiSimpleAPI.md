# \PypiSimpleAPI

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PypiSimpleCreate**](PypiSimpleAPI.md#PypiSimpleCreate) | **Post** /pypi/{path}/simple/ | Upload a package
[**PypiSimplePackageRead**](PypiSimpleAPI.md#PypiSimplePackageRead) | **Get** /pypi/{path}/simple/{package}/ | Get package simple page
[**PypiSimpleRead**](PypiSimpleAPI.md#PypiSimpleRead) | **Get** /pypi/{path}/simple/ | Get index simple page



## PypiSimpleCreate

> PackageUploadTaskResponse PypiSimpleCreate(ctx, path).Content(content).Sha256Digest(sha256Digest).Action(action).Execute()

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
    resp, r, err := apiClient.PypiSimpleAPI.PypiSimpleCreate(context.Background(), path).Content(content).Sha256Digest(sha256Digest).Action(action).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PypiSimpleAPI.PypiSimpleCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PypiSimpleCreate`: PackageUploadTaskResponse
    fmt.Fprintf(os.Stdout, "Response from `PypiSimpleAPI.PypiSimpleCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**path** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPypiSimpleCreateRequest struct via the builder pattern


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


## PypiSimplePackageRead

> PypiSimplePackageRead(ctx, package_, path).Fields(fields).ExcludeFields(excludeFields).Execute()

Get package simple page



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
    package_ := "package__example" // string | 
    path := "path_example" // string | 
    fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
    excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PypiSimpleAPI.PypiSimplePackageRead(context.Background(), package_, path).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PypiSimpleAPI.PypiSimplePackageRead``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**package_** | **string** |  | 
**path** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPypiSimplePackageReadRequest struct via the builder pattern


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


## PypiSimpleRead

> PypiSimpleRead(ctx, path).Fields(fields).ExcludeFields(excludeFields).Execute()

Get index simple page



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
    fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
    excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PypiSimpleAPI.PypiSimpleRead(context.Background(), path).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PypiSimpleAPI.PypiSimpleRead``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**path** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPypiSimpleReadRequest struct via the builder pattern


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

