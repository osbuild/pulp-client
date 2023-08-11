# \ContentguardsRbacAPI

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ContentguardsCoreRbacAddRole**](ContentguardsRbacAPI.md#ContentguardsCoreRbacAddRole) | **Post** {r_b_a_c_content_guard_href}add_role/ | 
[**ContentguardsCoreRbacCreate**](ContentguardsRbacAPI.md#ContentguardsCoreRbacCreate) | **Post** /pulp/api/v3/contentguards/core/rbac/ | Create a rbac content guard
[**ContentguardsCoreRbacDelete**](ContentguardsRbacAPI.md#ContentguardsCoreRbacDelete) | **Delete** {r_b_a_c_content_guard_href} | Delete a rbac content guard
[**ContentguardsCoreRbacList**](ContentguardsRbacAPI.md#ContentguardsCoreRbacList) | **Get** /pulp/api/v3/contentguards/core/rbac/ | List rbac content guards
[**ContentguardsCoreRbacListRoles**](ContentguardsRbacAPI.md#ContentguardsCoreRbacListRoles) | **Get** {r_b_a_c_content_guard_href}list_roles/ | 
[**ContentguardsCoreRbacMyPermissions**](ContentguardsRbacAPI.md#ContentguardsCoreRbacMyPermissions) | **Get** {r_b_a_c_content_guard_href}my_permissions/ | 
[**ContentguardsCoreRbacPartialUpdate**](ContentguardsRbacAPI.md#ContentguardsCoreRbacPartialUpdate) | **Patch** {r_b_a_c_content_guard_href} | Update a rbac content guard
[**ContentguardsCoreRbacRead**](ContentguardsRbacAPI.md#ContentguardsCoreRbacRead) | **Get** {r_b_a_c_content_guard_href} | Inspect a rbac content guard
[**ContentguardsCoreRbacRemoveRole**](ContentguardsRbacAPI.md#ContentguardsCoreRbacRemoveRole) | **Post** {r_b_a_c_content_guard_href}remove_role/ | 
[**ContentguardsCoreRbacUpdate**](ContentguardsRbacAPI.md#ContentguardsCoreRbacUpdate) | **Put** {r_b_a_c_content_guard_href} | Update a rbac content guard



## ContentguardsCoreRbacAddRole

> NestedRoleResponse ContentguardsCoreRbacAddRole(ctx, rBACContentGuardHref).NestedRole(nestedRole).Execute()





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
    rBACContentGuardHref := "rBACContentGuardHref_example" // string | 
    nestedRole := *openapiclient.NewNestedRole("Role_example") // NestedRole | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContentguardsRbacAPI.ContentguardsCoreRbacAddRole(context.Background(), rBACContentGuardHref).NestedRole(nestedRole).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContentguardsRbacAPI.ContentguardsCoreRbacAddRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContentguardsCoreRbacAddRole`: NestedRoleResponse
    fmt.Fprintf(os.Stdout, "Response from `ContentguardsRbacAPI.ContentguardsCoreRbacAddRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**rBACContentGuardHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiContentguardsCoreRbacAddRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **nestedRole** | [**NestedRole**](NestedRole.md) |  | 

### Return type

[**NestedRoleResponse**](NestedRoleResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContentguardsCoreRbacCreate

> RBACContentGuardResponse ContentguardsCoreRbacCreate(ctx).RBACContentGuard(rBACContentGuard).Execute()

Create a rbac content guard



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
    rBACContentGuard := *openapiclient.NewRBACContentGuard("Name_example") // RBACContentGuard | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContentguardsRbacAPI.ContentguardsCoreRbacCreate(context.Background()).RBACContentGuard(rBACContentGuard).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContentguardsRbacAPI.ContentguardsCoreRbacCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContentguardsCoreRbacCreate`: RBACContentGuardResponse
    fmt.Fprintf(os.Stdout, "Response from `ContentguardsRbacAPI.ContentguardsCoreRbacCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiContentguardsCoreRbacCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **rBACContentGuard** | [**RBACContentGuard**](RBACContentGuard.md) |  | 

### Return type

[**RBACContentGuardResponse**](RBACContentGuardResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContentguardsCoreRbacDelete

> ContentguardsCoreRbacDelete(ctx, rBACContentGuardHref).Execute()

Delete a rbac content guard



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
    rBACContentGuardHref := "rBACContentGuardHref_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ContentguardsRbacAPI.ContentguardsCoreRbacDelete(context.Background(), rBACContentGuardHref).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContentguardsRbacAPI.ContentguardsCoreRbacDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**rBACContentGuardHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiContentguardsCoreRbacDeleteRequest struct via the builder pattern


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


## ContentguardsCoreRbacList

> PaginatedRBACContentGuardResponseList ContentguardsCoreRbacList(ctx).Limit(limit).Name(name).NameContains(nameContains).NameIcontains(nameIcontains).NameIn(nameIn).NameStartswith(nameStartswith).Offset(offset).Ordering(ordering).PulpHrefIn(pulpHrefIn).PulpIdIn(pulpIdIn).Fields(fields).ExcludeFields(excludeFields).Execute()

List rbac content guards



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
    resp, r, err := apiClient.ContentguardsRbacAPI.ContentguardsCoreRbacList(context.Background()).Limit(limit).Name(name).NameContains(nameContains).NameIcontains(nameIcontains).NameIn(nameIn).NameStartswith(nameStartswith).Offset(offset).Ordering(ordering).PulpHrefIn(pulpHrefIn).PulpIdIn(pulpIdIn).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContentguardsRbacAPI.ContentguardsCoreRbacList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContentguardsCoreRbacList`: PaginatedRBACContentGuardResponseList
    fmt.Fprintf(os.Stdout, "Response from `ContentguardsRbacAPI.ContentguardsCoreRbacList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiContentguardsCoreRbacListRequest struct via the builder pattern


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

[**PaginatedRBACContentGuardResponseList**](PaginatedRBACContentGuardResponseList.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContentguardsCoreRbacListRoles

> ObjectRolesResponse ContentguardsCoreRbacListRoles(ctx, rBACContentGuardHref).Fields(fields).ExcludeFields(excludeFields).Execute()





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
    rBACContentGuardHref := "rBACContentGuardHref_example" // string | 
    fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
    excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContentguardsRbacAPI.ContentguardsCoreRbacListRoles(context.Background(), rBACContentGuardHref).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContentguardsRbacAPI.ContentguardsCoreRbacListRoles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContentguardsCoreRbacListRoles`: ObjectRolesResponse
    fmt.Fprintf(os.Stdout, "Response from `ContentguardsRbacAPI.ContentguardsCoreRbacListRoles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**rBACContentGuardHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiContentguardsCoreRbacListRolesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

### Return type

[**ObjectRolesResponse**](ObjectRolesResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContentguardsCoreRbacMyPermissions

> MyPermissionsResponse ContentguardsCoreRbacMyPermissions(ctx, rBACContentGuardHref).Fields(fields).ExcludeFields(excludeFields).Execute()





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
    rBACContentGuardHref := "rBACContentGuardHref_example" // string | 
    fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
    excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContentguardsRbacAPI.ContentguardsCoreRbacMyPermissions(context.Background(), rBACContentGuardHref).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContentguardsRbacAPI.ContentguardsCoreRbacMyPermissions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContentguardsCoreRbacMyPermissions`: MyPermissionsResponse
    fmt.Fprintf(os.Stdout, "Response from `ContentguardsRbacAPI.ContentguardsCoreRbacMyPermissions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**rBACContentGuardHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiContentguardsCoreRbacMyPermissionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

### Return type

[**MyPermissionsResponse**](MyPermissionsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContentguardsCoreRbacPartialUpdate

> RBACContentGuardResponse ContentguardsCoreRbacPartialUpdate(ctx, rBACContentGuardHref).PatchedRBACContentGuard(patchedRBACContentGuard).Execute()

Update a rbac content guard



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
    rBACContentGuardHref := "rBACContentGuardHref_example" // string | 
    patchedRBACContentGuard := *openapiclient.NewPatchedRBACContentGuard() // PatchedRBACContentGuard | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContentguardsRbacAPI.ContentguardsCoreRbacPartialUpdate(context.Background(), rBACContentGuardHref).PatchedRBACContentGuard(patchedRBACContentGuard).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContentguardsRbacAPI.ContentguardsCoreRbacPartialUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContentguardsCoreRbacPartialUpdate`: RBACContentGuardResponse
    fmt.Fprintf(os.Stdout, "Response from `ContentguardsRbacAPI.ContentguardsCoreRbacPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**rBACContentGuardHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiContentguardsCoreRbacPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedRBACContentGuard** | [**PatchedRBACContentGuard**](PatchedRBACContentGuard.md) |  | 

### Return type

[**RBACContentGuardResponse**](RBACContentGuardResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContentguardsCoreRbacRead

> RBACContentGuardResponse ContentguardsCoreRbacRead(ctx, rBACContentGuardHref).Fields(fields).ExcludeFields(excludeFields).Execute()

Inspect a rbac content guard



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
    rBACContentGuardHref := "rBACContentGuardHref_example" // string | 
    fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
    excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContentguardsRbacAPI.ContentguardsCoreRbacRead(context.Background(), rBACContentGuardHref).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContentguardsRbacAPI.ContentguardsCoreRbacRead``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContentguardsCoreRbacRead`: RBACContentGuardResponse
    fmt.Fprintf(os.Stdout, "Response from `ContentguardsRbacAPI.ContentguardsCoreRbacRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**rBACContentGuardHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiContentguardsCoreRbacReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

### Return type

[**RBACContentGuardResponse**](RBACContentGuardResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContentguardsCoreRbacRemoveRole

> NestedRoleResponse ContentguardsCoreRbacRemoveRole(ctx, rBACContentGuardHref).NestedRole(nestedRole).Execute()





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
    rBACContentGuardHref := "rBACContentGuardHref_example" // string | 
    nestedRole := *openapiclient.NewNestedRole("Role_example") // NestedRole | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContentguardsRbacAPI.ContentguardsCoreRbacRemoveRole(context.Background(), rBACContentGuardHref).NestedRole(nestedRole).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContentguardsRbacAPI.ContentguardsCoreRbacRemoveRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContentguardsCoreRbacRemoveRole`: NestedRoleResponse
    fmt.Fprintf(os.Stdout, "Response from `ContentguardsRbacAPI.ContentguardsCoreRbacRemoveRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**rBACContentGuardHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiContentguardsCoreRbacRemoveRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **nestedRole** | [**NestedRole**](NestedRole.md) |  | 

### Return type

[**NestedRoleResponse**](NestedRoleResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContentguardsCoreRbacUpdate

> RBACContentGuardResponse ContentguardsCoreRbacUpdate(ctx, rBACContentGuardHref).RBACContentGuard(rBACContentGuard).Execute()

Update a rbac content guard



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
    rBACContentGuardHref := "rBACContentGuardHref_example" // string | 
    rBACContentGuard := *openapiclient.NewRBACContentGuard("Name_example") // RBACContentGuard | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContentguardsRbacAPI.ContentguardsCoreRbacUpdate(context.Background(), rBACContentGuardHref).RBACContentGuard(rBACContentGuard).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContentguardsRbacAPI.ContentguardsCoreRbacUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContentguardsCoreRbacUpdate`: RBACContentGuardResponse
    fmt.Fprintf(os.Stdout, "Response from `ContentguardsRbacAPI.ContentguardsCoreRbacUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**rBACContentGuardHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiContentguardsCoreRbacUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **rBACContentGuard** | [**RBACContentGuard**](RBACContentGuard.md) |  | 

### Return type

[**RBACContentGuardResponse**](RBACContentGuardResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

