# \ContentguardsX509API

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ContentguardsCertguardX509Create**](ContentguardsX509API.md#ContentguardsCertguardX509Create) | **Post** /pulp/api/v3/contentguards/certguard/x509/ | Create a x509 cert guard
[**ContentguardsCertguardX509Delete**](ContentguardsX509API.md#ContentguardsCertguardX509Delete) | **Delete** {certguard_x509_cert_guard_href} | Delete a x509 cert guard
[**ContentguardsCertguardX509List**](ContentguardsX509API.md#ContentguardsCertguardX509List) | **Get** /pulp/api/v3/contentguards/certguard/x509/ | List x509 cert guards
[**ContentguardsCertguardX509PartialUpdate**](ContentguardsX509API.md#ContentguardsCertguardX509PartialUpdate) | **Patch** {certguard_x509_cert_guard_href} | Update a x509 cert guard
[**ContentguardsCertguardX509Read**](ContentguardsX509API.md#ContentguardsCertguardX509Read) | **Get** {certguard_x509_cert_guard_href} | Inspect a x509 cert guard
[**ContentguardsCertguardX509Update**](ContentguardsX509API.md#ContentguardsCertguardX509Update) | **Put** {certguard_x509_cert_guard_href} | Update a x509 cert guard



## ContentguardsCertguardX509Create

> CertguardX509CertGuardResponse ContentguardsCertguardX509Create(ctx).CertguardX509CertGuard(certguardX509CertGuard).Execute()

Create a x509 cert guard



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
    certguardX509CertGuard := *openapiclient.NewCertguardX509CertGuard("Name_example", "CaCertificate_example") // CertguardX509CertGuard | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContentguardsX509API.ContentguardsCertguardX509Create(context.Background()).CertguardX509CertGuard(certguardX509CertGuard).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContentguardsX509API.ContentguardsCertguardX509Create``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContentguardsCertguardX509Create`: CertguardX509CertGuardResponse
    fmt.Fprintf(os.Stdout, "Response from `ContentguardsX509API.ContentguardsCertguardX509Create`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiContentguardsCertguardX509CreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **certguardX509CertGuard** | [**CertguardX509CertGuard**](CertguardX509CertGuard.md) |  | 

### Return type

[**CertguardX509CertGuardResponse**](CertguardX509CertGuardResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContentguardsCertguardX509Delete

> ContentguardsCertguardX509Delete(ctx, certguardX509CertGuardHref).Execute()

Delete a x509 cert guard



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
    certguardX509CertGuardHref := "certguardX509CertGuardHref_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ContentguardsX509API.ContentguardsCertguardX509Delete(context.Background(), certguardX509CertGuardHref).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContentguardsX509API.ContentguardsCertguardX509Delete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**certguardX509CertGuardHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiContentguardsCertguardX509DeleteRequest struct via the builder pattern


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


## ContentguardsCertguardX509List

> PaginatedcertguardX509CertGuardResponseList ContentguardsCertguardX509List(ctx).Limit(limit).Name(name).NameContains(nameContains).NameIcontains(nameIcontains).NameIn(nameIn).NameStartswith(nameStartswith).Offset(offset).Ordering(ordering).PulpHrefIn(pulpHrefIn).PulpIdIn(pulpIdIn).Fields(fields).ExcludeFields(excludeFields).Execute()

List x509 cert guards



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
    ordering := []string{"Ordering_example"} // []string | Ordering  * `pulp_id` - Pulp id * `-pulp_id` - Pulp id (descending) * `pulp_created` - Pulp created * `-pulp_created` - Pulp created (descending) * `pulp_last_updated` - Pulp last updated * `-pulp_last_updated` - Pulp last updated (descending) * `pulp_type` - Pulp type * `-pulp_type` - Pulp type (descending) * `name` - Name * `-name` - Name (descending) * `description` - Description * `-description` - Description (descending) * `pk` - Pk * `-pk` - Pk (descending) (optional)
    pulpHrefIn := []string{"Inner_example"} // []string | Multiple values may be separated by commas. (optional)
    pulpIdIn := []string{"Inner_example"} // []string | Multiple values may be separated by commas. (optional)
    fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
    excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContentguardsX509API.ContentguardsCertguardX509List(context.Background()).Limit(limit).Name(name).NameContains(nameContains).NameIcontains(nameIcontains).NameIn(nameIn).NameStartswith(nameStartswith).Offset(offset).Ordering(ordering).PulpHrefIn(pulpHrefIn).PulpIdIn(pulpIdIn).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContentguardsX509API.ContentguardsCertguardX509List``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContentguardsCertguardX509List`: PaginatedcertguardX509CertGuardResponseList
    fmt.Fprintf(os.Stdout, "Response from `ContentguardsX509API.ContentguardsCertguardX509List`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiContentguardsCertguardX509ListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Number of results to return per page. | 
 **name** | **string** | Filter results where name matches value | 
 **nameContains** | **string** | Filter results where name contains value | 
 **nameIcontains** | **string** | Filter results where name contains value | 
 **nameIn** | **[]string** | Filter results where name is in a comma-separated list of values | 
 **nameStartswith** | **string** | Filter results where name starts with value | 
 **offset** | **int32** | The initial index from which to return the results. | 
 **ordering** | **[]string** | Ordering  * &#x60;pulp_id&#x60; - Pulp id * &#x60;-pulp_id&#x60; - Pulp id (descending) * &#x60;pulp_created&#x60; - Pulp created * &#x60;-pulp_created&#x60; - Pulp created (descending) * &#x60;pulp_last_updated&#x60; - Pulp last updated * &#x60;-pulp_last_updated&#x60; - Pulp last updated (descending) * &#x60;pulp_type&#x60; - Pulp type * &#x60;-pulp_type&#x60; - Pulp type (descending) * &#x60;name&#x60; - Name * &#x60;-name&#x60; - Name (descending) * &#x60;description&#x60; - Description * &#x60;-description&#x60; - Description (descending) * &#x60;pk&#x60; - Pk * &#x60;-pk&#x60; - Pk (descending) | 
 **pulpHrefIn** | **[]string** | Multiple values may be separated by commas. | 
 **pulpIdIn** | **[]string** | Multiple values may be separated by commas. | 
 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

### Return type

[**PaginatedcertguardX509CertGuardResponseList**](PaginatedcertguardX509CertGuardResponseList.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContentguardsCertguardX509PartialUpdate

> CertguardX509CertGuardResponse ContentguardsCertguardX509PartialUpdate(ctx, certguardX509CertGuardHref).PatchedcertguardX509CertGuard(patchedcertguardX509CertGuard).Execute()

Update a x509 cert guard



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
    certguardX509CertGuardHref := "certguardX509CertGuardHref_example" // string | 
    patchedcertguardX509CertGuard := *openapiclient.NewPatchedcertguardX509CertGuard() // PatchedcertguardX509CertGuard | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContentguardsX509API.ContentguardsCertguardX509PartialUpdate(context.Background(), certguardX509CertGuardHref).PatchedcertguardX509CertGuard(patchedcertguardX509CertGuard).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContentguardsX509API.ContentguardsCertguardX509PartialUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContentguardsCertguardX509PartialUpdate`: CertguardX509CertGuardResponse
    fmt.Fprintf(os.Stdout, "Response from `ContentguardsX509API.ContentguardsCertguardX509PartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**certguardX509CertGuardHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiContentguardsCertguardX509PartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedcertguardX509CertGuard** | [**PatchedcertguardX509CertGuard**](PatchedcertguardX509CertGuard.md) |  | 

### Return type

[**CertguardX509CertGuardResponse**](CertguardX509CertGuardResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContentguardsCertguardX509Read

> CertguardX509CertGuardResponse ContentguardsCertguardX509Read(ctx, certguardX509CertGuardHref).Fields(fields).ExcludeFields(excludeFields).Execute()

Inspect a x509 cert guard



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
    certguardX509CertGuardHref := "certguardX509CertGuardHref_example" // string | 
    fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
    excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContentguardsX509API.ContentguardsCertguardX509Read(context.Background(), certguardX509CertGuardHref).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContentguardsX509API.ContentguardsCertguardX509Read``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContentguardsCertguardX509Read`: CertguardX509CertGuardResponse
    fmt.Fprintf(os.Stdout, "Response from `ContentguardsX509API.ContentguardsCertguardX509Read`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**certguardX509CertGuardHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiContentguardsCertguardX509ReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

### Return type

[**CertguardX509CertGuardResponse**](CertguardX509CertGuardResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContentguardsCertguardX509Update

> CertguardX509CertGuardResponse ContentguardsCertguardX509Update(ctx, certguardX509CertGuardHref).CertguardX509CertGuard(certguardX509CertGuard).Execute()

Update a x509 cert guard



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
    certguardX509CertGuardHref := "certguardX509CertGuardHref_example" // string | 
    certguardX509CertGuard := *openapiclient.NewCertguardX509CertGuard("Name_example", "CaCertificate_example") // CertguardX509CertGuard | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContentguardsX509API.ContentguardsCertguardX509Update(context.Background(), certguardX509CertGuardHref).CertguardX509CertGuard(certguardX509CertGuard).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContentguardsX509API.ContentguardsCertguardX509Update``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContentguardsCertguardX509Update`: CertguardX509CertGuardResponse
    fmt.Fprintf(os.Stdout, "Response from `ContentguardsX509API.ContentguardsCertguardX509Update`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**certguardX509CertGuardHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiContentguardsCertguardX509UpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **certguardX509CertGuard** | [**CertguardX509CertGuard**](CertguardX509CertGuard.md) |  | 

### Return type

[**CertguardX509CertGuardResponse**](CertguardX509CertGuardResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

