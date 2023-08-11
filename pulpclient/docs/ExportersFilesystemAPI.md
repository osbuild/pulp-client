# \ExportersFilesystemAPI

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ExportersCoreFilesystemCreate**](ExportersFilesystemAPI.md#ExportersCoreFilesystemCreate) | **Post** /pulp/api/v3/exporters/core/filesystem/ | Create a filesystem exporter
[**ExportersCoreFilesystemDelete**](ExportersFilesystemAPI.md#ExportersCoreFilesystemDelete) | **Delete** {filesystem_exporter_href} | Delete a filesystem exporter
[**ExportersCoreFilesystemList**](ExportersFilesystemAPI.md#ExportersCoreFilesystemList) | **Get** /pulp/api/v3/exporters/core/filesystem/ | List filesystem exporters
[**ExportersCoreFilesystemPartialUpdate**](ExportersFilesystemAPI.md#ExportersCoreFilesystemPartialUpdate) | **Patch** {filesystem_exporter_href} | Update a filesystem exporter
[**ExportersCoreFilesystemRead**](ExportersFilesystemAPI.md#ExportersCoreFilesystemRead) | **Get** {filesystem_exporter_href} | Inspect a filesystem exporter
[**ExportersCoreFilesystemUpdate**](ExportersFilesystemAPI.md#ExportersCoreFilesystemUpdate) | **Put** {filesystem_exporter_href} | Update a filesystem exporter



## ExportersCoreFilesystemCreate

> FilesystemExporterResponse ExportersCoreFilesystemCreate(ctx).FilesystemExporter(filesystemExporter).Execute()

Create a filesystem exporter



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
    filesystemExporter := *openapiclient.NewFilesystemExporter("Name_example", "Path_example") // FilesystemExporter | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ExportersFilesystemAPI.ExportersCoreFilesystemCreate(context.Background()).FilesystemExporter(filesystemExporter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExportersFilesystemAPI.ExportersCoreFilesystemCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExportersCoreFilesystemCreate`: FilesystemExporterResponse
    fmt.Fprintf(os.Stdout, "Response from `ExportersFilesystemAPI.ExportersCoreFilesystemCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExportersCoreFilesystemCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filesystemExporter** | [**FilesystemExporter**](FilesystemExporter.md) |  | 

### Return type

[**FilesystemExporterResponse**](FilesystemExporterResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportersCoreFilesystemDelete

> AsyncOperationResponse ExportersCoreFilesystemDelete(ctx, filesystemExporterHref).Execute()

Delete a filesystem exporter



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ExportersFilesystemAPI.ExportersCoreFilesystemDelete(context.Background(), filesystemExporterHref).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExportersFilesystemAPI.ExportersCoreFilesystemDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExportersCoreFilesystemDelete`: AsyncOperationResponse
    fmt.Fprintf(os.Stdout, "Response from `ExportersFilesystemAPI.ExportersCoreFilesystemDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**filesystemExporterHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiExportersCoreFilesystemDeleteRequest struct via the builder pattern


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


## ExportersCoreFilesystemList

> PaginatedFilesystemExporterResponseList ExportersCoreFilesystemList(ctx).Limit(limit).Name(name).NameContains(nameContains).NameIcontains(nameIcontains).NameIn(nameIn).NameStartswith(nameStartswith).Offset(offset).Ordering(ordering).PulpHrefIn(pulpHrefIn).PulpIdIn(pulpIdIn).Fields(fields).ExcludeFields(excludeFields).Execute()

List filesystem exporters



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
    ordering := []string{"Ordering_example"} // []string | Ordering  * `pulp_id` - Pulp id * `-pulp_id` - Pulp id (descending) * `pulp_created` - Pulp created * `-pulp_created` - Pulp created (descending) * `pulp_last_updated` - Pulp last updated * `-pulp_last_updated` - Pulp last updated (descending) * `pulp_type` - Pulp type * `-pulp_type` - Pulp type (descending) * `name` - Name * `-name` - Name (descending) * `path` - Path * `-path` - Path (descending) * `method` - Method * `-method` - Method (descending) * `pk` - Pk * `-pk` - Pk (descending) (optional)
    pulpHrefIn := []string{"Inner_example"} // []string | Multiple values may be separated by commas. (optional)
    pulpIdIn := []string{"Inner_example"} // []string | Multiple values may be separated by commas. (optional)
    fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
    excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ExportersFilesystemAPI.ExportersCoreFilesystemList(context.Background()).Limit(limit).Name(name).NameContains(nameContains).NameIcontains(nameIcontains).NameIn(nameIn).NameStartswith(nameStartswith).Offset(offset).Ordering(ordering).PulpHrefIn(pulpHrefIn).PulpIdIn(pulpIdIn).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExportersFilesystemAPI.ExportersCoreFilesystemList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExportersCoreFilesystemList`: PaginatedFilesystemExporterResponseList
    fmt.Fprintf(os.Stdout, "Response from `ExportersFilesystemAPI.ExportersCoreFilesystemList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExportersCoreFilesystemListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Number of results to return per page. | 
 **name** | **string** | Filter results where name matches value | 
 **nameContains** | **string** | Filter results where name contains value | 
 **nameIcontains** | **string** | Filter results where name contains value | 
 **nameIn** | **[]string** | Filter results where name is in a comma-separated list of values | 
 **nameStartswith** | **string** | Filter results where name starts with value | 
 **offset** | **int32** | The initial index from which to return the results. | 
 **ordering** | **[]string** | Ordering  * &#x60;pulp_id&#x60; - Pulp id * &#x60;-pulp_id&#x60; - Pulp id (descending) * &#x60;pulp_created&#x60; - Pulp created * &#x60;-pulp_created&#x60; - Pulp created (descending) * &#x60;pulp_last_updated&#x60; - Pulp last updated * &#x60;-pulp_last_updated&#x60; - Pulp last updated (descending) * &#x60;pulp_type&#x60; - Pulp type * &#x60;-pulp_type&#x60; - Pulp type (descending) * &#x60;name&#x60; - Name * &#x60;-name&#x60; - Name (descending) * &#x60;path&#x60; - Path * &#x60;-path&#x60; - Path (descending) * &#x60;method&#x60; - Method * &#x60;-method&#x60; - Method (descending) * &#x60;pk&#x60; - Pk * &#x60;-pk&#x60; - Pk (descending) | 
 **pulpHrefIn** | **[]string** | Multiple values may be separated by commas. | 
 **pulpIdIn** | **[]string** | Multiple values may be separated by commas. | 
 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

### Return type

[**PaginatedFilesystemExporterResponseList**](PaginatedFilesystemExporterResponseList.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportersCoreFilesystemPartialUpdate

> AsyncOperationResponse ExportersCoreFilesystemPartialUpdate(ctx, filesystemExporterHref).PatchedFilesystemExporter(patchedFilesystemExporter).Execute()

Update a filesystem exporter



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
    patchedFilesystemExporter := *openapiclient.NewPatchedFilesystemExporter() // PatchedFilesystemExporter | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ExportersFilesystemAPI.ExportersCoreFilesystemPartialUpdate(context.Background(), filesystemExporterHref).PatchedFilesystemExporter(patchedFilesystemExporter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExportersFilesystemAPI.ExportersCoreFilesystemPartialUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExportersCoreFilesystemPartialUpdate`: AsyncOperationResponse
    fmt.Fprintf(os.Stdout, "Response from `ExportersFilesystemAPI.ExportersCoreFilesystemPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**filesystemExporterHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiExportersCoreFilesystemPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedFilesystemExporter** | [**PatchedFilesystemExporter**](PatchedFilesystemExporter.md) |  | 

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


## ExportersCoreFilesystemRead

> FilesystemExporterResponse ExportersCoreFilesystemRead(ctx, filesystemExporterHref).Fields(fields).ExcludeFields(excludeFields).Execute()

Inspect a filesystem exporter



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
    fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
    excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ExportersFilesystemAPI.ExportersCoreFilesystemRead(context.Background(), filesystemExporterHref).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExportersFilesystemAPI.ExportersCoreFilesystemRead``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExportersCoreFilesystemRead`: FilesystemExporterResponse
    fmt.Fprintf(os.Stdout, "Response from `ExportersFilesystemAPI.ExportersCoreFilesystemRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**filesystemExporterHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiExportersCoreFilesystemReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

### Return type

[**FilesystemExporterResponse**](FilesystemExporterResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportersCoreFilesystemUpdate

> AsyncOperationResponse ExportersCoreFilesystemUpdate(ctx, filesystemExporterHref).FilesystemExporter(filesystemExporter).Execute()

Update a filesystem exporter



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
    filesystemExporter := *openapiclient.NewFilesystemExporter("Name_example", "Path_example") // FilesystemExporter | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ExportersFilesystemAPI.ExportersCoreFilesystemUpdate(context.Background(), filesystemExporterHref).FilesystemExporter(filesystemExporter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExportersFilesystemAPI.ExportersCoreFilesystemUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExportersCoreFilesystemUpdate`: AsyncOperationResponse
    fmt.Fprintf(os.Stdout, "Response from `ExportersFilesystemAPI.ExportersCoreFilesystemUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**filesystemExporterHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiExportersCoreFilesystemUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filesystemExporter** | [**FilesystemExporter**](FilesystemExporter.md) |  | 

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

