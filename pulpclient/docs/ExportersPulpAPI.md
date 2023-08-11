# \ExportersPulpAPI

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ExportersCorePulpCreate**](ExportersPulpAPI.md#ExportersCorePulpCreate) | **Post** /pulp/api/v3/exporters/core/pulp/ | Create a pulp exporter
[**ExportersCorePulpDelete**](ExportersPulpAPI.md#ExportersCorePulpDelete) | **Delete** {pulp_exporter_href} | Delete a pulp exporter
[**ExportersCorePulpList**](ExportersPulpAPI.md#ExportersCorePulpList) | **Get** /pulp/api/v3/exporters/core/pulp/ | List pulp exporters
[**ExportersCorePulpPartialUpdate**](ExportersPulpAPI.md#ExportersCorePulpPartialUpdate) | **Patch** {pulp_exporter_href} | Update a pulp exporter
[**ExportersCorePulpRead**](ExportersPulpAPI.md#ExportersCorePulpRead) | **Get** {pulp_exporter_href} | Inspect a pulp exporter
[**ExportersCorePulpUpdate**](ExportersPulpAPI.md#ExportersCorePulpUpdate) | **Put** {pulp_exporter_href} | Update a pulp exporter



## ExportersCorePulpCreate

> PulpExporterResponse ExportersCorePulpCreate(ctx).PulpExporter(pulpExporter).Execute()

Create a pulp exporter



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
    pulpExporter := *openapiclient.NewPulpExporter("Name_example", "Path_example", []string{"Repositories_example"}) // PulpExporter | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ExportersPulpAPI.ExportersCorePulpCreate(context.Background()).PulpExporter(pulpExporter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExportersPulpAPI.ExportersCorePulpCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExportersCorePulpCreate`: PulpExporterResponse
    fmt.Fprintf(os.Stdout, "Response from `ExportersPulpAPI.ExportersCorePulpCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExportersCorePulpCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pulpExporter** | [**PulpExporter**](PulpExporter.md) |  | 

### Return type

[**PulpExporterResponse**](PulpExporterResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportersCorePulpDelete

> AsyncOperationResponse ExportersCorePulpDelete(ctx, pulpExporterHref).Execute()

Delete a pulp exporter



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ExportersPulpAPI.ExportersCorePulpDelete(context.Background(), pulpExporterHref).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExportersPulpAPI.ExportersCorePulpDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExportersCorePulpDelete`: AsyncOperationResponse
    fmt.Fprintf(os.Stdout, "Response from `ExportersPulpAPI.ExportersCorePulpDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pulpExporterHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiExportersCorePulpDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AsyncOperationResponse**](AsyncOperationResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportersCorePulpList

> PaginatedPulpExporterResponseList ExportersCorePulpList(ctx).Limit(limit).Name(name).NameContains(nameContains).NameIcontains(nameIcontains).NameIn(nameIn).NameStartswith(nameStartswith).Offset(offset).Ordering(ordering).PulpHrefIn(pulpHrefIn).PulpIdIn(pulpIdIn).Fields(fields).ExcludeFields(excludeFields).Execute()

List pulp exporters



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
    limit := int32(56) // int32 | Number of results to return per page. (optional)
    name := "name_example" // string | Filter results where name matches value (optional)
    nameContains := "nameContains_example" // string | Filter results where name contains value (optional)
    nameIcontains := "nameIcontains_example" // string | Filter results where name contains value (optional)
    nameIn := []string{"Inner_example"} // []string | Filter results where name is in a comma-separated list of values (optional)
    nameStartswith := "nameStartswith_example" // string | Filter results where name starts with value (optional)
    offset := int32(56) // int32 | The initial index from which to return the results. (optional)
    ordering := []string{"Ordering_example"} // []string | Ordering  * `pulp_id` - Pulp id * `-pulp_id` - Pulp id (descending) * `pulp_created` - Pulp created * `-pulp_created` - Pulp created (descending) * `pulp_last_updated` - Pulp last updated * `-pulp_last_updated` - Pulp last updated (descending) * `pulp_type` - Pulp type * `-pulp_type` - Pulp type (descending) * `name` - Name * `-name` - Name (descending) * `path` - Path * `-path` - Path (descending) * `pk` - Pk * `-pk` - Pk (descending) (optional)
    pulpHrefIn := []string{"Inner_example"} // []string | Multiple values may be separated by commas. (optional)
    pulpIdIn := []string{"Inner_example"} // []string | Multiple values may be separated by commas. (optional)
    fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
    excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ExportersPulpAPI.ExportersCorePulpList(context.Background()).Limit(limit).Name(name).NameContains(nameContains).NameIcontains(nameIcontains).NameIn(nameIn).NameStartswith(nameStartswith).Offset(offset).Ordering(ordering).PulpHrefIn(pulpHrefIn).PulpIdIn(pulpIdIn).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExportersPulpAPI.ExportersCorePulpList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExportersCorePulpList`: PaginatedPulpExporterResponseList
    fmt.Fprintf(os.Stdout, "Response from `ExportersPulpAPI.ExportersCorePulpList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExportersCorePulpListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Number of results to return per page. | 
 **name** | **string** | Filter results where name matches value | 
 **nameContains** | **string** | Filter results where name contains value | 
 **nameIcontains** | **string** | Filter results where name contains value | 
 **nameIn** | **[]string** | Filter results where name is in a comma-separated list of values | 
 **nameStartswith** | **string** | Filter results where name starts with value | 
 **offset** | **int32** | The initial index from which to return the results. | 
 **ordering** | **[]string** | Ordering  * &#x60;pulp_id&#x60; - Pulp id * &#x60;-pulp_id&#x60; - Pulp id (descending) * &#x60;pulp_created&#x60; - Pulp created * &#x60;-pulp_created&#x60; - Pulp created (descending) * &#x60;pulp_last_updated&#x60; - Pulp last updated * &#x60;-pulp_last_updated&#x60; - Pulp last updated (descending) * &#x60;pulp_type&#x60; - Pulp type * &#x60;-pulp_type&#x60; - Pulp type (descending) * &#x60;name&#x60; - Name * &#x60;-name&#x60; - Name (descending) * &#x60;path&#x60; - Path * &#x60;-path&#x60; - Path (descending) * &#x60;pk&#x60; - Pk * &#x60;-pk&#x60; - Pk (descending) | 
 **pulpHrefIn** | **[]string** | Multiple values may be separated by commas. | 
 **pulpIdIn** | **[]string** | Multiple values may be separated by commas. | 
 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

### Return type

[**PaginatedPulpExporterResponseList**](PaginatedPulpExporterResponseList.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportersCorePulpPartialUpdate

> AsyncOperationResponse ExportersCorePulpPartialUpdate(ctx, pulpExporterHref).PatchedPulpExporter(patchedPulpExporter).Execute()

Update a pulp exporter



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
    patchedPulpExporter := *openapiclient.NewPatchedPulpExporter() // PatchedPulpExporter | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ExportersPulpAPI.ExportersCorePulpPartialUpdate(context.Background(), pulpExporterHref).PatchedPulpExporter(patchedPulpExporter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExportersPulpAPI.ExportersCorePulpPartialUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExportersCorePulpPartialUpdate`: AsyncOperationResponse
    fmt.Fprintf(os.Stdout, "Response from `ExportersPulpAPI.ExportersCorePulpPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pulpExporterHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiExportersCorePulpPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedPulpExporter** | [**PatchedPulpExporter**](PatchedPulpExporter.md) |  | 

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


## ExportersCorePulpRead

> PulpExporterResponse ExportersCorePulpRead(ctx, pulpExporterHref).Fields(fields).ExcludeFields(excludeFields).Execute()

Inspect a pulp exporter



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
    fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
    excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ExportersPulpAPI.ExportersCorePulpRead(context.Background(), pulpExporterHref).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExportersPulpAPI.ExportersCorePulpRead``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExportersCorePulpRead`: PulpExporterResponse
    fmt.Fprintf(os.Stdout, "Response from `ExportersPulpAPI.ExportersCorePulpRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pulpExporterHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiExportersCorePulpReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

### Return type

[**PulpExporterResponse**](PulpExporterResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportersCorePulpUpdate

> AsyncOperationResponse ExportersCorePulpUpdate(ctx, pulpExporterHref).PulpExporter(pulpExporter).Execute()

Update a pulp exporter



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
    pulpExporter := *openapiclient.NewPulpExporter("Name_example", "Path_example", []string{"Repositories_example"}) // PulpExporter | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ExportersPulpAPI.ExportersCorePulpUpdate(context.Background(), pulpExporterHref).PulpExporter(pulpExporter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExportersPulpAPI.ExportersCorePulpUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExportersCorePulpUpdate`: AsyncOperationResponse
    fmt.Fprintf(os.Stdout, "Response from `ExportersPulpAPI.ExportersCorePulpUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pulpExporterHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiExportersCorePulpUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pulpExporter** | [**PulpExporter**](PulpExporter.md) |  | 

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

