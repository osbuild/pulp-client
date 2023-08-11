# \AccessPoliciesAPI

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AccessPoliciesList**](AccessPoliciesAPI.md#AccessPoliciesList) | **Get** /pulp/api/v3/access_policies/ | List access policys
[**AccessPoliciesPartialUpdate**](AccessPoliciesAPI.md#AccessPoliciesPartialUpdate) | **Patch** {access_policy_href} | Update an access policy
[**AccessPoliciesRead**](AccessPoliciesAPI.md#AccessPoliciesRead) | **Get** {access_policy_href} | Inspect an access policy
[**AccessPoliciesReset**](AccessPoliciesAPI.md#AccessPoliciesReset) | **Post** {access_policy_href}reset/ | 
[**AccessPoliciesUpdate**](AccessPoliciesAPI.md#AccessPoliciesUpdate) | **Put** {access_policy_href} | Update an access policy



## AccessPoliciesList

> PaginatedAccessPolicyResponseList AccessPoliciesList(ctx).Customized(customized).Limit(limit).Offset(offset).Ordering(ordering).PulpHrefIn(pulpHrefIn).PulpIdIn(pulpIdIn).ViewsetName(viewsetName).ViewsetNameContains(viewsetNameContains).ViewsetNameIcontains(viewsetNameIcontains).ViewsetNameIn(viewsetNameIn).ViewsetNameStartswith(viewsetNameStartswith).Fields(fields).ExcludeFields(excludeFields).Execute()

List access policys



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
    customized := true // bool | Filter results where customized matches value (optional)
    limit := int32(56) // int32 | Number of results to return per page. (optional)
    offset := int32(56) // int32 | The initial index from which to return the results. (optional)
    ordering := []string{"Ordering_example"} // []string | Ordering  * `pulp_id` - Pulp id * `-pulp_id` - Pulp id (descending) * `pulp_created` - Pulp created * `-pulp_created` - Pulp created (descending) * `pulp_last_updated` - Pulp last updated * `-pulp_last_updated` - Pulp last updated (descending) * `creation_hooks` - Creation hooks * `-creation_hooks` - Creation hooks (descending) * `statements` - Statements * `-statements` - Statements (descending) * `viewset_name` - Viewset name * `-viewset_name` - Viewset name (descending) * `customized` - Customized * `-customized` - Customized (descending) * `queryset_scoping` - Queryset scoping * `-queryset_scoping` - Queryset scoping (descending) * `pk` - Pk * `-pk` - Pk (descending) (optional)
    pulpHrefIn := []string{"Inner_example"} // []string | Multiple values may be separated by commas. (optional)
    pulpIdIn := []string{"Inner_example"} // []string | Multiple values may be separated by commas. (optional)
    viewsetName := "viewsetName_example" // string | Filter results where viewset_name matches value (optional)
    viewsetNameContains := "viewsetNameContains_example" // string | Filter results where viewset_name contains value (optional)
    viewsetNameIcontains := "viewsetNameIcontains_example" // string | Filter results where viewset_name contains value (optional)
    viewsetNameIn := []string{"Inner_example"} // []string | Filter results where viewset_name is in a comma-separated list of values (optional)
    viewsetNameStartswith := "viewsetNameStartswith_example" // string | Filter results where viewset_name starts with value (optional)
    fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
    excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccessPoliciesAPI.AccessPoliciesList(context.Background()).Customized(customized).Limit(limit).Offset(offset).Ordering(ordering).PulpHrefIn(pulpHrefIn).PulpIdIn(pulpIdIn).ViewsetName(viewsetName).ViewsetNameContains(viewsetNameContains).ViewsetNameIcontains(viewsetNameIcontains).ViewsetNameIn(viewsetNameIn).ViewsetNameStartswith(viewsetNameStartswith).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccessPoliciesAPI.AccessPoliciesList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AccessPoliciesList`: PaginatedAccessPolicyResponseList
    fmt.Fprintf(os.Stdout, "Response from `AccessPoliciesAPI.AccessPoliciesList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAccessPoliciesListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **customized** | **bool** | Filter results where customized matches value | 
 **limit** | **int32** | Number of results to return per page. | 
 **offset** | **int32** | The initial index from which to return the results. | 
 **ordering** | **[]string** | Ordering  * &#x60;pulp_id&#x60; - Pulp id * &#x60;-pulp_id&#x60; - Pulp id (descending) * &#x60;pulp_created&#x60; - Pulp created * &#x60;-pulp_created&#x60; - Pulp created (descending) * &#x60;pulp_last_updated&#x60; - Pulp last updated * &#x60;-pulp_last_updated&#x60; - Pulp last updated (descending) * &#x60;creation_hooks&#x60; - Creation hooks * &#x60;-creation_hooks&#x60; - Creation hooks (descending) * &#x60;statements&#x60; - Statements * &#x60;-statements&#x60; - Statements (descending) * &#x60;viewset_name&#x60; - Viewset name * &#x60;-viewset_name&#x60; - Viewset name (descending) * &#x60;customized&#x60; - Customized * &#x60;-customized&#x60; - Customized (descending) * &#x60;queryset_scoping&#x60; - Queryset scoping * &#x60;-queryset_scoping&#x60; - Queryset scoping (descending) * &#x60;pk&#x60; - Pk * &#x60;-pk&#x60; - Pk (descending) | 
 **pulpHrefIn** | **[]string** | Multiple values may be separated by commas. | 
 **pulpIdIn** | **[]string** | Multiple values may be separated by commas. | 
 **viewsetName** | **string** | Filter results where viewset_name matches value | 
 **viewsetNameContains** | **string** | Filter results where viewset_name contains value | 
 **viewsetNameIcontains** | **string** | Filter results where viewset_name contains value | 
 **viewsetNameIn** | **[]string** | Filter results where viewset_name is in a comma-separated list of values | 
 **viewsetNameStartswith** | **string** | Filter results where viewset_name starts with value | 
 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

### Return type

[**PaginatedAccessPolicyResponseList**](PaginatedAccessPolicyResponseList.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AccessPoliciesPartialUpdate

> AccessPolicyResponse AccessPoliciesPartialUpdate(ctx, accessPolicyHref).PatchedAccessPolicy(patchedAccessPolicy).Execute()

Update an access policy



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
    accessPolicyHref := "accessPolicyHref_example" // string | 
    patchedAccessPolicy := *openapiclient.NewPatchedAccessPolicy() // PatchedAccessPolicy | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccessPoliciesAPI.AccessPoliciesPartialUpdate(context.Background(), accessPolicyHref).PatchedAccessPolicy(patchedAccessPolicy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccessPoliciesAPI.AccessPoliciesPartialUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AccessPoliciesPartialUpdate`: AccessPolicyResponse
    fmt.Fprintf(os.Stdout, "Response from `AccessPoliciesAPI.AccessPoliciesPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accessPolicyHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAccessPoliciesPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedAccessPolicy** | [**PatchedAccessPolicy**](PatchedAccessPolicy.md) |  | 

### Return type

[**AccessPolicyResponse**](AccessPolicyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AccessPoliciesRead

> AccessPolicyResponse AccessPoliciesRead(ctx, accessPolicyHref).Fields(fields).ExcludeFields(excludeFields).Execute()

Inspect an access policy



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
    accessPolicyHref := "accessPolicyHref_example" // string | 
    fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
    excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccessPoliciesAPI.AccessPoliciesRead(context.Background(), accessPolicyHref).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccessPoliciesAPI.AccessPoliciesRead``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AccessPoliciesRead`: AccessPolicyResponse
    fmt.Fprintf(os.Stdout, "Response from `AccessPoliciesAPI.AccessPoliciesRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accessPolicyHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAccessPoliciesReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

### Return type

[**AccessPolicyResponse**](AccessPolicyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AccessPoliciesReset

> AccessPolicyResponse AccessPoliciesReset(ctx, accessPolicyHref).Execute()





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
    accessPolicyHref := "accessPolicyHref_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccessPoliciesAPI.AccessPoliciesReset(context.Background(), accessPolicyHref).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccessPoliciesAPI.AccessPoliciesReset``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AccessPoliciesReset`: AccessPolicyResponse
    fmt.Fprintf(os.Stdout, "Response from `AccessPoliciesAPI.AccessPoliciesReset`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accessPolicyHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAccessPoliciesResetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AccessPolicyResponse**](AccessPolicyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AccessPoliciesUpdate

> AccessPolicyResponse AccessPoliciesUpdate(ctx, accessPolicyHref).AccessPolicy(accessPolicy).Execute()

Update an access policy



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
    accessPolicyHref := "accessPolicyHref_example" // string | 
    accessPolicy := *openapiclient.NewAccessPolicy([]map[string]interface{}{map[string]interface{}(123)}) // AccessPolicy | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccessPoliciesAPI.AccessPoliciesUpdate(context.Background(), accessPolicyHref).AccessPolicy(accessPolicy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccessPoliciesAPI.AccessPoliciesUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AccessPoliciesUpdate`: AccessPolicyResponse
    fmt.Fprintf(os.Stdout, "Response from `AccessPoliciesAPI.AccessPoliciesUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accessPolicyHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAccessPoliciesUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accessPolicy** | [**AccessPolicy**](AccessPolicy.md) |  | 

### Return type

[**AccessPolicyResponse**](AccessPolicyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

