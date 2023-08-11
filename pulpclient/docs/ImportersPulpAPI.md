# \ImportersPulpAPI

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ImportersCorePulpCreate**](ImportersPulpAPI.md#ImportersCorePulpCreate) | **Post** /pulp/api/v3/importers/core/pulp/ | Create a pulp importer
[**ImportersCorePulpDelete**](ImportersPulpAPI.md#ImportersCorePulpDelete) | **Delete** {pulp_importer_href} | Delete a pulp importer
[**ImportersCorePulpList**](ImportersPulpAPI.md#ImportersCorePulpList) | **Get** /pulp/api/v3/importers/core/pulp/ | List pulp importers
[**ImportersCorePulpPartialUpdate**](ImportersPulpAPI.md#ImportersCorePulpPartialUpdate) | **Patch** {pulp_importer_href} | Update a pulp importer
[**ImportersCorePulpRead**](ImportersPulpAPI.md#ImportersCorePulpRead) | **Get** {pulp_importer_href} | Inspect a pulp importer
[**ImportersCorePulpUpdate**](ImportersPulpAPI.md#ImportersCorePulpUpdate) | **Put** {pulp_importer_href} | Update a pulp importer



## ImportersCorePulpCreate

> PulpImporterResponse ImportersCorePulpCreate(ctx).PulpImporter(pulpImporter).Execute()

Create a pulp importer



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
    pulpImporter := *openapiclient.NewPulpImporter("Name_example") // PulpImporter | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ImportersPulpAPI.ImportersCorePulpCreate(context.Background()).PulpImporter(pulpImporter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImportersPulpAPI.ImportersCorePulpCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ImportersCorePulpCreate`: PulpImporterResponse
    fmt.Fprintf(os.Stdout, "Response from `ImportersPulpAPI.ImportersCorePulpCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiImportersCorePulpCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pulpImporter** | [**PulpImporter**](PulpImporter.md) |  | 

### Return type

[**PulpImporterResponse**](PulpImporterResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImportersCorePulpDelete

> ImportersCorePulpDelete(ctx, pulpImporterHref).Execute()

Delete a pulp importer



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
    pulpImporterHref := "pulpImporterHref_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ImportersPulpAPI.ImportersCorePulpDelete(context.Background(), pulpImporterHref).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImportersPulpAPI.ImportersCorePulpDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pulpImporterHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiImportersCorePulpDeleteRequest struct via the builder pattern


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


## ImportersCorePulpList

> PaginatedPulpImporterResponseList ImportersCorePulpList(ctx).Limit(limit).Name(name).NameContains(nameContains).NameIcontains(nameIcontains).NameIn(nameIn).NameStartswith(nameStartswith).Offset(offset).Ordering(ordering).PulpHrefIn(pulpHrefIn).PulpIdIn(pulpIdIn).Fields(fields).ExcludeFields(excludeFields).Execute()

List pulp importers



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
    ordering := []string{"Ordering_example"} // []string | Ordering  * `pulp_id` - Pulp id * `-pulp_id` - Pulp id (descending) * `pulp_created` - Pulp created * `-pulp_created` - Pulp created (descending) * `pulp_last_updated` - Pulp last updated * `-pulp_last_updated` - Pulp last updated (descending) * `pulp_type` - Pulp type * `-pulp_type` - Pulp type (descending) * `name` - Name * `-name` - Name (descending) * `pk` - Pk * `-pk` - Pk (descending) (optional)
    pulpHrefIn := []string{"Inner_example"} // []string | Multiple values may be separated by commas. (optional)
    pulpIdIn := []string{"Inner_example"} // []string | Multiple values may be separated by commas. (optional)
    fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
    excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ImportersPulpAPI.ImportersCorePulpList(context.Background()).Limit(limit).Name(name).NameContains(nameContains).NameIcontains(nameIcontains).NameIn(nameIn).NameStartswith(nameStartswith).Offset(offset).Ordering(ordering).PulpHrefIn(pulpHrefIn).PulpIdIn(pulpIdIn).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImportersPulpAPI.ImportersCorePulpList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ImportersCorePulpList`: PaginatedPulpImporterResponseList
    fmt.Fprintf(os.Stdout, "Response from `ImportersPulpAPI.ImportersCorePulpList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiImportersCorePulpListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Number of results to return per page. | 
 **name** | **string** | Filter results where name matches value | 
 **nameContains** | **string** | Filter results where name contains value | 
 **nameIcontains** | **string** | Filter results where name contains value | 
 **nameIn** | **[]string** | Filter results where name is in a comma-separated list of values | 
 **nameStartswith** | **string** | Filter results where name starts with value | 
 **offset** | **int32** | The initial index from which to return the results. | 
 **ordering** | **[]string** | Ordering  * &#x60;pulp_id&#x60; - Pulp id * &#x60;-pulp_id&#x60; - Pulp id (descending) * &#x60;pulp_created&#x60; - Pulp created * &#x60;-pulp_created&#x60; - Pulp created (descending) * &#x60;pulp_last_updated&#x60; - Pulp last updated * &#x60;-pulp_last_updated&#x60; - Pulp last updated (descending) * &#x60;pulp_type&#x60; - Pulp type * &#x60;-pulp_type&#x60; - Pulp type (descending) * &#x60;name&#x60; - Name * &#x60;-name&#x60; - Name (descending) * &#x60;pk&#x60; - Pk * &#x60;-pk&#x60; - Pk (descending) | 
 **pulpHrefIn** | **[]string** | Multiple values may be separated by commas. | 
 **pulpIdIn** | **[]string** | Multiple values may be separated by commas. | 
 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

### Return type

[**PaginatedPulpImporterResponseList**](PaginatedPulpImporterResponseList.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImportersCorePulpPartialUpdate

> PulpImporterResponse ImportersCorePulpPartialUpdate(ctx, pulpImporterHref).PatchedPulpImporter(patchedPulpImporter).Execute()

Update a pulp importer



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
    pulpImporterHref := "pulpImporterHref_example" // string | 
    patchedPulpImporter := *openapiclient.NewPatchedPulpImporter() // PatchedPulpImporter | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ImportersPulpAPI.ImportersCorePulpPartialUpdate(context.Background(), pulpImporterHref).PatchedPulpImporter(patchedPulpImporter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImportersPulpAPI.ImportersCorePulpPartialUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ImportersCorePulpPartialUpdate`: PulpImporterResponse
    fmt.Fprintf(os.Stdout, "Response from `ImportersPulpAPI.ImportersCorePulpPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pulpImporterHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiImportersCorePulpPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedPulpImporter** | [**PatchedPulpImporter**](PatchedPulpImporter.md) |  | 

### Return type

[**PulpImporterResponse**](PulpImporterResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImportersCorePulpRead

> PulpImporterResponse ImportersCorePulpRead(ctx, pulpImporterHref).Fields(fields).ExcludeFields(excludeFields).Execute()

Inspect a pulp importer



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
    pulpImporterHref := "pulpImporterHref_example" // string | 
    fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
    excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ImportersPulpAPI.ImportersCorePulpRead(context.Background(), pulpImporterHref).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImportersPulpAPI.ImportersCorePulpRead``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ImportersCorePulpRead`: PulpImporterResponse
    fmt.Fprintf(os.Stdout, "Response from `ImportersPulpAPI.ImportersCorePulpRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pulpImporterHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiImportersCorePulpReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

### Return type

[**PulpImporterResponse**](PulpImporterResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImportersCorePulpUpdate

> PulpImporterResponse ImportersCorePulpUpdate(ctx, pulpImporterHref).PulpImporter(pulpImporter).Execute()

Update a pulp importer



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
    pulpImporterHref := "pulpImporterHref_example" // string | 
    pulpImporter := *openapiclient.NewPulpImporter("Name_example") // PulpImporter | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ImportersPulpAPI.ImportersCorePulpUpdate(context.Background(), pulpImporterHref).PulpImporter(pulpImporter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImportersPulpAPI.ImportersCorePulpUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ImportersCorePulpUpdate`: PulpImporterResponse
    fmt.Fprintf(os.Stdout, "Response from `ImportersPulpAPI.ImportersCorePulpUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pulpImporterHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiImportersCorePulpUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pulpImporter** | [**PulpImporter**](PulpImporter.md) |  | 

### Return type

[**PulpImporterResponse**](PulpImporterResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

