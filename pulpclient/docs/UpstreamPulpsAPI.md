# \UpstreamPulpsAPI

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UpstreamPulpsCreate**](UpstreamPulpsAPI.md#UpstreamPulpsCreate) | **Post** /pulp/api/v3/upstream-pulps/ | Create an upstream pulp
[**UpstreamPulpsDelete**](UpstreamPulpsAPI.md#UpstreamPulpsDelete) | **Delete** {upstream_pulp_href} | Delete an upstream pulp
[**UpstreamPulpsList**](UpstreamPulpsAPI.md#UpstreamPulpsList) | **Get** /pulp/api/v3/upstream-pulps/ | List upstream pulps
[**UpstreamPulpsPartialUpdate**](UpstreamPulpsAPI.md#UpstreamPulpsPartialUpdate) | **Patch** {upstream_pulp_href} | Update an upstream pulp
[**UpstreamPulpsRead**](UpstreamPulpsAPI.md#UpstreamPulpsRead) | **Get** {upstream_pulp_href} | Inspect an upstream pulp
[**UpstreamPulpsReplicate**](UpstreamPulpsAPI.md#UpstreamPulpsReplicate) | **Post** {upstream_pulp_href}replicate/ | Replicate
[**UpstreamPulpsUpdate**](UpstreamPulpsAPI.md#UpstreamPulpsUpdate) | **Put** {upstream_pulp_href} | Update an upstream pulp



## UpstreamPulpsCreate

> UpstreamPulpResponse UpstreamPulpsCreate(ctx).UpstreamPulp(upstreamPulp).Execute()

Create an upstream pulp



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
    upstreamPulp := *openapiclient.NewUpstreamPulp("Name_example", "BaseUrl_example", "ApiRoot_example") // UpstreamPulp | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UpstreamPulpsAPI.UpstreamPulpsCreate(context.Background()).UpstreamPulp(upstreamPulp).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UpstreamPulpsAPI.UpstreamPulpsCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpstreamPulpsCreate`: UpstreamPulpResponse
    fmt.Fprintf(os.Stdout, "Response from `UpstreamPulpsAPI.UpstreamPulpsCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpstreamPulpsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **upstreamPulp** | [**UpstreamPulp**](UpstreamPulp.md) |  | 

### Return type

[**UpstreamPulpResponse**](UpstreamPulpResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpstreamPulpsDelete

> UpstreamPulpsDelete(ctx, upstreamPulpHref).Execute()

Delete an upstream pulp



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
    upstreamPulpHref := "upstreamPulpHref_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.UpstreamPulpsAPI.UpstreamPulpsDelete(context.Background(), upstreamPulpHref).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UpstreamPulpsAPI.UpstreamPulpsDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**upstreamPulpHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpstreamPulpsDeleteRequest struct via the builder pattern


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


## UpstreamPulpsList

> PaginatedUpstreamPulpResponseList UpstreamPulpsList(ctx).Limit(limit).Offset(offset).Fields(fields).ExcludeFields(excludeFields).Execute()

List upstream pulps



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
    offset := int32(56) // int32 | The initial index from which to return the results. (optional)
    fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
    excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UpstreamPulpsAPI.UpstreamPulpsList(context.Background()).Limit(limit).Offset(offset).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UpstreamPulpsAPI.UpstreamPulpsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpstreamPulpsList`: PaginatedUpstreamPulpResponseList
    fmt.Fprintf(os.Stdout, "Response from `UpstreamPulpsAPI.UpstreamPulpsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpstreamPulpsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Number of results to return per page. | 
 **offset** | **int32** | The initial index from which to return the results. | 
 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

### Return type

[**PaginatedUpstreamPulpResponseList**](PaginatedUpstreamPulpResponseList.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpstreamPulpsPartialUpdate

> UpstreamPulpResponse UpstreamPulpsPartialUpdate(ctx, upstreamPulpHref).PatchedUpstreamPulp(patchedUpstreamPulp).Execute()

Update an upstream pulp



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
    upstreamPulpHref := "upstreamPulpHref_example" // string | 
    patchedUpstreamPulp := *openapiclient.NewPatchedUpstreamPulp() // PatchedUpstreamPulp | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UpstreamPulpsAPI.UpstreamPulpsPartialUpdate(context.Background(), upstreamPulpHref).PatchedUpstreamPulp(patchedUpstreamPulp).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UpstreamPulpsAPI.UpstreamPulpsPartialUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpstreamPulpsPartialUpdate`: UpstreamPulpResponse
    fmt.Fprintf(os.Stdout, "Response from `UpstreamPulpsAPI.UpstreamPulpsPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**upstreamPulpHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpstreamPulpsPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedUpstreamPulp** | [**PatchedUpstreamPulp**](PatchedUpstreamPulp.md) |  | 

### Return type

[**UpstreamPulpResponse**](UpstreamPulpResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpstreamPulpsRead

> UpstreamPulpResponse UpstreamPulpsRead(ctx, upstreamPulpHref).Fields(fields).ExcludeFields(excludeFields).Execute()

Inspect an upstream pulp



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
    upstreamPulpHref := "upstreamPulpHref_example" // string | 
    fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
    excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UpstreamPulpsAPI.UpstreamPulpsRead(context.Background(), upstreamPulpHref).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UpstreamPulpsAPI.UpstreamPulpsRead``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpstreamPulpsRead`: UpstreamPulpResponse
    fmt.Fprintf(os.Stdout, "Response from `UpstreamPulpsAPI.UpstreamPulpsRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**upstreamPulpHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpstreamPulpsReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

### Return type

[**UpstreamPulpResponse**](UpstreamPulpResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpstreamPulpsReplicate

> TaskGroupOperationResponse UpstreamPulpsReplicate(ctx, upstreamPulpHref).Execute()

Replicate



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
    upstreamPulpHref := "upstreamPulpHref_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UpstreamPulpsAPI.UpstreamPulpsReplicate(context.Background(), upstreamPulpHref).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UpstreamPulpsAPI.UpstreamPulpsReplicate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpstreamPulpsReplicate`: TaskGroupOperationResponse
    fmt.Fprintf(os.Stdout, "Response from `UpstreamPulpsAPI.UpstreamPulpsReplicate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**upstreamPulpHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpstreamPulpsReplicateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TaskGroupOperationResponse**](TaskGroupOperationResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpstreamPulpsUpdate

> UpstreamPulpResponse UpstreamPulpsUpdate(ctx, upstreamPulpHref).UpstreamPulp(upstreamPulp).Execute()

Update an upstream pulp



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
    upstreamPulpHref := "upstreamPulpHref_example" // string | 
    upstreamPulp := *openapiclient.NewUpstreamPulp("Name_example", "BaseUrl_example", "ApiRoot_example") // UpstreamPulp | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UpstreamPulpsAPI.UpstreamPulpsUpdate(context.Background(), upstreamPulpHref).UpstreamPulp(upstreamPulp).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UpstreamPulpsAPI.UpstreamPulpsUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpstreamPulpsUpdate`: UpstreamPulpResponse
    fmt.Fprintf(os.Stdout, "Response from `UpstreamPulpsAPI.UpstreamPulpsUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**upstreamPulpHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpstreamPulpsUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **upstreamPulp** | [**UpstreamPulp**](UpstreamPulp.md) |  | 

### Return type

[**UpstreamPulpResponse**](UpstreamPulpResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

