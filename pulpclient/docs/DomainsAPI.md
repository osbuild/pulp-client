# \DomainsAPI

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DomainsCreate**](DomainsAPI.md#DomainsCreate) | **Post** /pulp/api/v3/domains/ | Create a domain
[**DomainsDelete**](DomainsAPI.md#DomainsDelete) | **Delete** {domain_href} | Delete a domain
[**DomainsList**](DomainsAPI.md#DomainsList) | **Get** /pulp/api/v3/domains/ | List domains
[**DomainsPartialUpdate**](DomainsAPI.md#DomainsPartialUpdate) | **Patch** {domain_href} | Update a domain
[**DomainsRead**](DomainsAPI.md#DomainsRead) | **Get** {domain_href} | Inspect a domain
[**DomainsUpdate**](DomainsAPI.md#DomainsUpdate) | **Put** {domain_href} | Update a domain



## DomainsCreate

> DomainResponse DomainsCreate(ctx).Domain(domain).Execute()

Create a domain



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
    domain := *openapiclient.NewDomain("Name_example", openapiclient.StorageClassEnum("pulpcore.app.models.storage.FileSystem"), map[string]interface{}(123)) // Domain | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DomainsAPI.DomainsCreate(context.Background()).Domain(domain).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DomainsAPI.DomainsCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DomainsCreate`: DomainResponse
    fmt.Fprintf(os.Stdout, "Response from `DomainsAPI.DomainsCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDomainsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **domain** | [**Domain**](Domain.md) |  | 

### Return type

[**DomainResponse**](DomainResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DomainsDelete

> AsyncOperationResponse DomainsDelete(ctx, domainHref).Execute()

Delete a domain



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
    domainHref := "domainHref_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DomainsAPI.DomainsDelete(context.Background(), domainHref).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DomainsAPI.DomainsDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DomainsDelete`: AsyncOperationResponse
    fmt.Fprintf(os.Stdout, "Response from `DomainsAPI.DomainsDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domainHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDomainsDeleteRequest struct via the builder pattern


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


## DomainsList

> PaginatedDomainResponseList DomainsList(ctx).Limit(limit).Name(name).NameContains(nameContains).NameIcontains(nameIcontains).NameIn(nameIn).NameStartswith(nameStartswith).Offset(offset).Ordering(ordering).PulpHrefIn(pulpHrefIn).PulpIdIn(pulpIdIn).Fields(fields).ExcludeFields(excludeFields).Execute()

List domains



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
    ordering := []string{"Ordering_example"} // []string | Ordering  * `pulp_id` - Pulp id * `-pulp_id` - Pulp id (descending) * `pulp_created` - Pulp created * `-pulp_created` - Pulp created (descending) * `pulp_last_updated` - Pulp last updated * `-pulp_last_updated` - Pulp last updated (descending) * `name` - Name * `-name` - Name (descending) * `description` - Description * `-description` - Description (descending) * `storage_class` - Storage class * `-storage_class` - Storage class (descending) * `storage_settings` - Storage settings * `-storage_settings` - Storage settings (descending) * `redirect_to_object_storage` - Redirect to object storage * `-redirect_to_object_storage` - Redirect to object storage (descending) * `hide_guarded_distributions` - Hide guarded distributions * `-hide_guarded_distributions` - Hide guarded distributions (descending) * `pk` - Pk * `-pk` - Pk (descending) (optional)
    pulpHrefIn := []string{"Inner_example"} // []string | Multiple values may be separated by commas. (optional)
    pulpIdIn := []string{"Inner_example"} // []string | Multiple values may be separated by commas. (optional)
    fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
    excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DomainsAPI.DomainsList(context.Background()).Limit(limit).Name(name).NameContains(nameContains).NameIcontains(nameIcontains).NameIn(nameIn).NameStartswith(nameStartswith).Offset(offset).Ordering(ordering).PulpHrefIn(pulpHrefIn).PulpIdIn(pulpIdIn).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DomainsAPI.DomainsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DomainsList`: PaginatedDomainResponseList
    fmt.Fprintf(os.Stdout, "Response from `DomainsAPI.DomainsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDomainsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Number of results to return per page. | 
 **name** | **string** | Filter results where name matches value | 
 **nameContains** | **string** | Filter results where name contains value | 
 **nameIcontains** | **string** | Filter results where name contains value | 
 **nameIn** | **[]string** | Filter results where name is in a comma-separated list of values | 
 **nameStartswith** | **string** | Filter results where name starts with value | 
 **offset** | **int32** | The initial index from which to return the results. | 
 **ordering** | **[]string** | Ordering  * &#x60;pulp_id&#x60; - Pulp id * &#x60;-pulp_id&#x60; - Pulp id (descending) * &#x60;pulp_created&#x60; - Pulp created * &#x60;-pulp_created&#x60; - Pulp created (descending) * &#x60;pulp_last_updated&#x60; - Pulp last updated * &#x60;-pulp_last_updated&#x60; - Pulp last updated (descending) * &#x60;name&#x60; - Name * &#x60;-name&#x60; - Name (descending) * &#x60;description&#x60; - Description * &#x60;-description&#x60; - Description (descending) * &#x60;storage_class&#x60; - Storage class * &#x60;-storage_class&#x60; - Storage class (descending) * &#x60;storage_settings&#x60; - Storage settings * &#x60;-storage_settings&#x60; - Storage settings (descending) * &#x60;redirect_to_object_storage&#x60; - Redirect to object storage * &#x60;-redirect_to_object_storage&#x60; - Redirect to object storage (descending) * &#x60;hide_guarded_distributions&#x60; - Hide guarded distributions * &#x60;-hide_guarded_distributions&#x60; - Hide guarded distributions (descending) * &#x60;pk&#x60; - Pk * &#x60;-pk&#x60; - Pk (descending) | 
 **pulpHrefIn** | **[]string** | Multiple values may be separated by commas. | 
 **pulpIdIn** | **[]string** | Multiple values may be separated by commas. | 
 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

### Return type

[**PaginatedDomainResponseList**](PaginatedDomainResponseList.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DomainsPartialUpdate

> AsyncOperationResponse DomainsPartialUpdate(ctx, domainHref).PatchedDomain(patchedDomain).Execute()

Update a domain



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
    domainHref := "domainHref_example" // string | 
    patchedDomain := *openapiclient.NewPatchedDomain() // PatchedDomain | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DomainsAPI.DomainsPartialUpdate(context.Background(), domainHref).PatchedDomain(patchedDomain).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DomainsAPI.DomainsPartialUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DomainsPartialUpdate`: AsyncOperationResponse
    fmt.Fprintf(os.Stdout, "Response from `DomainsAPI.DomainsPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domainHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDomainsPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedDomain** | [**PatchedDomain**](PatchedDomain.md) |  | 

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


## DomainsRead

> DomainResponse DomainsRead(ctx, domainHref).Fields(fields).ExcludeFields(excludeFields).Execute()

Inspect a domain



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
    domainHref := "domainHref_example" // string | 
    fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
    excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DomainsAPI.DomainsRead(context.Background(), domainHref).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DomainsAPI.DomainsRead``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DomainsRead`: DomainResponse
    fmt.Fprintf(os.Stdout, "Response from `DomainsAPI.DomainsRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domainHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDomainsReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

### Return type

[**DomainResponse**](DomainResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DomainsUpdate

> AsyncOperationResponse DomainsUpdate(ctx, domainHref).Domain(domain).Execute()

Update a domain



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
    domainHref := "domainHref_example" // string | 
    domain := *openapiclient.NewDomain("Name_example", openapiclient.StorageClassEnum("pulpcore.app.models.storage.FileSystem"), map[string]interface{}(123)) // Domain | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DomainsAPI.DomainsUpdate(context.Background(), domainHref).Domain(domain).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DomainsAPI.DomainsUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DomainsUpdate`: AsyncOperationResponse
    fmt.Fprintf(os.Stdout, "Response from `DomainsAPI.DomainsUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domainHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDomainsUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **domain** | [**Domain**](Domain.md) |  | 

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

