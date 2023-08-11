# \ExportersPulpExportsAPI

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ExportersCorePulpExportsCreate**](ExportersPulpExportsAPI.md#ExportersCorePulpExportsCreate) | **Post** {pulp_exporter_href}exports/ | Create a pulp export
[**ExportersCorePulpExportsDelete**](ExportersPulpExportsAPI.md#ExportersCorePulpExportsDelete) | **Delete** {pulp_pulp_export_href} | Delete a pulp export
[**ExportersCorePulpExportsList**](ExportersPulpExportsAPI.md#ExportersCorePulpExportsList) | **Get** {pulp_exporter_href}exports/ | List pulp exports
[**ExportersCorePulpExportsRead**](ExportersPulpExportsAPI.md#ExportersCorePulpExportsRead) | **Get** {pulp_pulp_export_href} | Inspect a pulp export



## ExportersCorePulpExportsCreate

> AsyncOperationResponse ExportersCorePulpExportsCreate(ctx, pulpExporterHref).PulpExport(pulpExport).Execute()

Create a pulp export



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
    pulpExporterHref := "pulpExporterHref_example" // string | 
    pulpExport := *openapiclient.NewPulpExport() // PulpExport | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ExportersPulpExportsAPI.ExportersCorePulpExportsCreate(context.Background(), pulpExporterHref).PulpExport(pulpExport).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExportersPulpExportsAPI.ExportersCorePulpExportsCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExportersCorePulpExportsCreate`: AsyncOperationResponse
    fmt.Fprintf(os.Stdout, "Response from `ExportersPulpExportsAPI.ExportersCorePulpExportsCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pulpExporterHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiExportersCorePulpExportsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pulpExport** | [**PulpExport**](PulpExport.md) |  | 

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


## ExportersCorePulpExportsDelete

> ExportersCorePulpExportsDelete(ctx, pulpPulpExportHref).Execute()

Delete a pulp export



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
    pulpPulpExportHref := "pulpPulpExportHref_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ExportersPulpExportsAPI.ExportersCorePulpExportsDelete(context.Background(), pulpPulpExportHref).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExportersPulpExportsAPI.ExportersCorePulpExportsDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pulpPulpExportHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiExportersCorePulpExportsDeleteRequest struct via the builder pattern


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


## ExportersCorePulpExportsList

> PaginatedPulpExportResponseList ExportersCorePulpExportsList(ctx, pulpExporterHref).Limit(limit).Offset(offset).Fields(fields).ExcludeFields(excludeFields).Execute()

List pulp exports



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
    pulpExporterHref := "pulpExporterHref_example" // string | 
    limit := int32(56) // int32 | Number of results to return per page. (optional)
    offset := int32(56) // int32 | The initial index from which to return the results. (optional)
    fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
    excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ExportersPulpExportsAPI.ExportersCorePulpExportsList(context.Background(), pulpExporterHref).Limit(limit).Offset(offset).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExportersPulpExportsAPI.ExportersCorePulpExportsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExportersCorePulpExportsList`: PaginatedPulpExportResponseList
    fmt.Fprintf(os.Stdout, "Response from `ExportersPulpExportsAPI.ExportersCorePulpExportsList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pulpExporterHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiExportersCorePulpExportsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Number of results to return per page. | 
 **offset** | **int32** | The initial index from which to return the results. | 
 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

### Return type

[**PaginatedPulpExportResponseList**](PaginatedPulpExportResponseList.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportersCorePulpExportsRead

> PulpExportResponse ExportersCorePulpExportsRead(ctx, pulpPulpExportHref).Fields(fields).ExcludeFields(excludeFields).Execute()

Inspect a pulp export



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
    pulpPulpExportHref := "pulpPulpExportHref_example" // string | 
    fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
    excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ExportersPulpExportsAPI.ExportersCorePulpExportsRead(context.Background(), pulpPulpExportHref).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExportersPulpExportsAPI.ExportersCorePulpExportsRead``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExportersCorePulpExportsRead`: PulpExportResponse
    fmt.Fprintf(os.Stdout, "Response from `ExportersPulpExportsAPI.ExportersCorePulpExportsRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pulpPulpExportHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiExportersCorePulpExportsReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

### Return type

[**PulpExportResponse**](PulpExportResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

