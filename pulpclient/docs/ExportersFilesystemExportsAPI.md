# \ExportersFilesystemExportsAPI

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ExportersCoreFilesystemExportsCreate**](ExportersFilesystemExportsAPI.md#ExportersCoreFilesystemExportsCreate) | **Post** {filesystem_exporter_href}exports/ | Create a filesystem export
[**ExportersCoreFilesystemExportsDelete**](ExportersFilesystemExportsAPI.md#ExportersCoreFilesystemExportsDelete) | **Delete** {filesystem_filesystem_export_href} | Delete a filesystem export
[**ExportersCoreFilesystemExportsList**](ExportersFilesystemExportsAPI.md#ExportersCoreFilesystemExportsList) | **Get** {filesystem_exporter_href}exports/ | List filesystem exports
[**ExportersCoreFilesystemExportsRead**](ExportersFilesystemExportsAPI.md#ExportersCoreFilesystemExportsRead) | **Get** {filesystem_filesystem_export_href} | Inspect a filesystem export



## ExportersCoreFilesystemExportsCreate

> AsyncOperationResponse ExportersCoreFilesystemExportsCreate(ctx, filesystemExporterHref).FilesystemExport(filesystemExport).Execute()

Create a filesystem export



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
    filesystemExporterHref := "filesystemExporterHref_example" // string | 
    filesystemExport := *openapiclient.NewFilesystemExport() // FilesystemExport | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ExportersFilesystemExportsAPI.ExportersCoreFilesystemExportsCreate(context.Background(), filesystemExporterHref).FilesystemExport(filesystemExport).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExportersFilesystemExportsAPI.ExportersCoreFilesystemExportsCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExportersCoreFilesystemExportsCreate`: AsyncOperationResponse
    fmt.Fprintf(os.Stdout, "Response from `ExportersFilesystemExportsAPI.ExportersCoreFilesystemExportsCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**filesystemExporterHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiExportersCoreFilesystemExportsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filesystemExport** | [**FilesystemExport**](FilesystemExport.md) |  | 

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


## ExportersCoreFilesystemExportsDelete

> ExportersCoreFilesystemExportsDelete(ctx, filesystemFilesystemExportHref).Execute()

Delete a filesystem export



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
    filesystemFilesystemExportHref := "filesystemFilesystemExportHref_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ExportersFilesystemExportsAPI.ExportersCoreFilesystemExportsDelete(context.Background(), filesystemFilesystemExportHref).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExportersFilesystemExportsAPI.ExportersCoreFilesystemExportsDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**filesystemFilesystemExportHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiExportersCoreFilesystemExportsDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## ExportersCoreFilesystemExportsList

> PaginatedFilesystemExportResponseList ExportersCoreFilesystemExportsList(ctx, filesystemExporterHref).Limit(limit).Offset(offset).Fields(fields).ExcludeFields(excludeFields).Execute()

List filesystem exports



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
    filesystemExporterHref := "filesystemExporterHref_example" // string | 
    limit := int32(56) // int32 | Number of results to return per page. (optional)
    offset := int32(56) // int32 | The initial index from which to return the results. (optional)
    fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
    excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ExportersFilesystemExportsAPI.ExportersCoreFilesystemExportsList(context.Background(), filesystemExporterHref).Limit(limit).Offset(offset).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExportersFilesystemExportsAPI.ExportersCoreFilesystemExportsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExportersCoreFilesystemExportsList`: PaginatedFilesystemExportResponseList
    fmt.Fprintf(os.Stdout, "Response from `ExportersFilesystemExportsAPI.ExportersCoreFilesystemExportsList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**filesystemExporterHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiExportersCoreFilesystemExportsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Number of results to return per page. | 
 **offset** | **int32** | The initial index from which to return the results. | 
 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

### Return type

[**PaginatedFilesystemExportResponseList**](PaginatedFilesystemExportResponseList.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportersCoreFilesystemExportsRead

> FilesystemExportResponse ExportersCoreFilesystemExportsRead(ctx, filesystemFilesystemExportHref).Fields(fields).ExcludeFields(excludeFields).Execute()

Inspect a filesystem export



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
    filesystemFilesystemExportHref := "filesystemFilesystemExportHref_example" // string | 
    fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
    excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ExportersFilesystemExportsAPI.ExportersCoreFilesystemExportsRead(context.Background(), filesystemFilesystemExportHref).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExportersFilesystemExportsAPI.ExportersCoreFilesystemExportsRead``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExportersCoreFilesystemExportsRead`: FilesystemExportResponse
    fmt.Fprintf(os.Stdout, "Response from `ExportersFilesystemExportsAPI.ExportersCoreFilesystemExportsRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**filesystemFilesystemExportHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiExportersCoreFilesystemExportsReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

### Return type

[**FilesystemExportResponse**](FilesystemExportResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

