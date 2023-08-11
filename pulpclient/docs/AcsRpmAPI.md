# \AcsRpmAPI

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AcsRpmRpmAddRole**](AcsRpmAPI.md#AcsRpmRpmAddRole) | **Post** {rpm_rpm_alternate_content_source_href}add_role/ | 
[**AcsRpmRpmCreate**](AcsRpmAPI.md#AcsRpmRpmCreate) | **Post** /pulp/api/v3/acs/rpm/rpm/ | Create a rpm alternate content source
[**AcsRpmRpmDelete**](AcsRpmAPI.md#AcsRpmRpmDelete) | **Delete** {rpm_rpm_alternate_content_source_href} | Delete a rpm alternate content source
[**AcsRpmRpmList**](AcsRpmAPI.md#AcsRpmRpmList) | **Get** /pulp/api/v3/acs/rpm/rpm/ | List rpm alternate content sources
[**AcsRpmRpmListRoles**](AcsRpmAPI.md#AcsRpmRpmListRoles) | **Get** {rpm_rpm_alternate_content_source_href}list_roles/ | 
[**AcsRpmRpmMyPermissions**](AcsRpmAPI.md#AcsRpmRpmMyPermissions) | **Get** {rpm_rpm_alternate_content_source_href}my_permissions/ | 
[**AcsRpmRpmPartialUpdate**](AcsRpmAPI.md#AcsRpmRpmPartialUpdate) | **Patch** {rpm_rpm_alternate_content_source_href} | Update a rpm alternate content source
[**AcsRpmRpmRead**](AcsRpmAPI.md#AcsRpmRpmRead) | **Get** {rpm_rpm_alternate_content_source_href} | Inspect a rpm alternate content source
[**AcsRpmRpmRefresh**](AcsRpmAPI.md#AcsRpmRpmRefresh) | **Post** {rpm_rpm_alternate_content_source_href}refresh/ | 
[**AcsRpmRpmRemoveRole**](AcsRpmAPI.md#AcsRpmRpmRemoveRole) | **Post** {rpm_rpm_alternate_content_source_href}remove_role/ | 
[**AcsRpmRpmUpdate**](AcsRpmAPI.md#AcsRpmRpmUpdate) | **Put** {rpm_rpm_alternate_content_source_href} | Update a rpm alternate content source



## AcsRpmRpmAddRole

> NestedRoleResponse AcsRpmRpmAddRole(ctx, rpmRpmAlternateContentSourceHref).NestedRole(nestedRole).Execute()





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
    rpmRpmAlternateContentSourceHref := "rpmRpmAlternateContentSourceHref_example" // string | 
    nestedRole := *openapiclient.NewNestedRole("Role_example") // NestedRole | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AcsRpmAPI.AcsRpmRpmAddRole(context.Background(), rpmRpmAlternateContentSourceHref).NestedRole(nestedRole).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AcsRpmAPI.AcsRpmRpmAddRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AcsRpmRpmAddRole`: NestedRoleResponse
    fmt.Fprintf(os.Stdout, "Response from `AcsRpmAPI.AcsRpmRpmAddRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**rpmRpmAlternateContentSourceHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAcsRpmRpmAddRoleRequest struct via the builder pattern


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


## AcsRpmRpmCreate

> RpmRpmAlternateContentSourceResponse AcsRpmRpmCreate(ctx).RpmRpmAlternateContentSource(rpmRpmAlternateContentSource).Execute()

Create a rpm alternate content source



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
    rpmRpmAlternateContentSource := *openapiclient.NewRpmRpmAlternateContentSource("Name_example", "Remote_example") // RpmRpmAlternateContentSource | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AcsRpmAPI.AcsRpmRpmCreate(context.Background()).RpmRpmAlternateContentSource(rpmRpmAlternateContentSource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AcsRpmAPI.AcsRpmRpmCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AcsRpmRpmCreate`: RpmRpmAlternateContentSourceResponse
    fmt.Fprintf(os.Stdout, "Response from `AcsRpmAPI.AcsRpmRpmCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAcsRpmRpmCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **rpmRpmAlternateContentSource** | [**RpmRpmAlternateContentSource**](RpmRpmAlternateContentSource.md) |  | 

### Return type

[**RpmRpmAlternateContentSourceResponse**](RpmRpmAlternateContentSourceResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AcsRpmRpmDelete

> AsyncOperationResponse AcsRpmRpmDelete(ctx, rpmRpmAlternateContentSourceHref).Execute()

Delete a rpm alternate content source



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
    rpmRpmAlternateContentSourceHref := "rpmRpmAlternateContentSourceHref_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AcsRpmAPI.AcsRpmRpmDelete(context.Background(), rpmRpmAlternateContentSourceHref).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AcsRpmAPI.AcsRpmRpmDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AcsRpmRpmDelete`: AsyncOperationResponse
    fmt.Fprintf(os.Stdout, "Response from `AcsRpmAPI.AcsRpmRpmDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**rpmRpmAlternateContentSourceHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAcsRpmRpmDeleteRequest struct via the builder pattern


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


## AcsRpmRpmList

> PaginatedrpmRpmAlternateContentSourceResponseList AcsRpmRpmList(ctx).Limit(limit).Name(name).NameContains(nameContains).NameIcontains(nameIcontains).NameIn(nameIn).NameStartswith(nameStartswith).Offset(offset).Ordering(ordering).PulpHrefIn(pulpHrefIn).PulpIdIn(pulpIdIn).Fields(fields).ExcludeFields(excludeFields).Execute()

List rpm alternate content sources



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
    ordering := []string{"Ordering_example"} // []string | Ordering  * `pulp_id` - Pulp id * `-pulp_id` - Pulp id (descending) * `pulp_created` - Pulp created * `-pulp_created` - Pulp created (descending) * `pulp_last_updated` - Pulp last updated * `-pulp_last_updated` - Pulp last updated (descending) * `pulp_type` - Pulp type * `-pulp_type` - Pulp type (descending) * `name` - Name * `-name` - Name (descending) * `last_refreshed` - Last refreshed * `-last_refreshed` - Last refreshed (descending) * `pk` - Pk * `-pk` - Pk (descending) (optional)
    pulpHrefIn := []string{"Inner_example"} // []string | Multiple values may be separated by commas. (optional)
    pulpIdIn := []string{"Inner_example"} // []string | Multiple values may be separated by commas. (optional)
    fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
    excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AcsRpmAPI.AcsRpmRpmList(context.Background()).Limit(limit).Name(name).NameContains(nameContains).NameIcontains(nameIcontains).NameIn(nameIn).NameStartswith(nameStartswith).Offset(offset).Ordering(ordering).PulpHrefIn(pulpHrefIn).PulpIdIn(pulpIdIn).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AcsRpmAPI.AcsRpmRpmList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AcsRpmRpmList`: PaginatedrpmRpmAlternateContentSourceResponseList
    fmt.Fprintf(os.Stdout, "Response from `AcsRpmAPI.AcsRpmRpmList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAcsRpmRpmListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Number of results to return per page. | 
 **name** | **string** | Filter results where name matches value | 
 **nameContains** | **string** | Filter results where name contains value | 
 **nameIcontains** | **string** | Filter results where name contains value | 
 **nameIn** | **[]string** | Filter results where name is in a comma-separated list of values | 
 **nameStartswith** | **string** | Filter results where name starts with value | 
 **offset** | **int32** | The initial index from which to return the results. | 
 **ordering** | **[]string** | Ordering  * &#x60;pulp_id&#x60; - Pulp id * &#x60;-pulp_id&#x60; - Pulp id (descending) * &#x60;pulp_created&#x60; - Pulp created * &#x60;-pulp_created&#x60; - Pulp created (descending) * &#x60;pulp_last_updated&#x60; - Pulp last updated * &#x60;-pulp_last_updated&#x60; - Pulp last updated (descending) * &#x60;pulp_type&#x60; - Pulp type * &#x60;-pulp_type&#x60; - Pulp type (descending) * &#x60;name&#x60; - Name * &#x60;-name&#x60; - Name (descending) * &#x60;last_refreshed&#x60; - Last refreshed * &#x60;-last_refreshed&#x60; - Last refreshed (descending) * &#x60;pk&#x60; - Pk * &#x60;-pk&#x60; - Pk (descending) | 
 **pulpHrefIn** | **[]string** | Multiple values may be separated by commas. | 
 **pulpIdIn** | **[]string** | Multiple values may be separated by commas. | 
 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

### Return type

[**PaginatedrpmRpmAlternateContentSourceResponseList**](PaginatedrpmRpmAlternateContentSourceResponseList.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AcsRpmRpmListRoles

> ObjectRolesResponse AcsRpmRpmListRoles(ctx, rpmRpmAlternateContentSourceHref).Fields(fields).ExcludeFields(excludeFields).Execute()





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
    rpmRpmAlternateContentSourceHref := "rpmRpmAlternateContentSourceHref_example" // string | 
    fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
    excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AcsRpmAPI.AcsRpmRpmListRoles(context.Background(), rpmRpmAlternateContentSourceHref).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AcsRpmAPI.AcsRpmRpmListRoles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AcsRpmRpmListRoles`: ObjectRolesResponse
    fmt.Fprintf(os.Stdout, "Response from `AcsRpmAPI.AcsRpmRpmListRoles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**rpmRpmAlternateContentSourceHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAcsRpmRpmListRolesRequest struct via the builder pattern


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


## AcsRpmRpmMyPermissions

> MyPermissionsResponse AcsRpmRpmMyPermissions(ctx, rpmRpmAlternateContentSourceHref).Fields(fields).ExcludeFields(excludeFields).Execute()





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
    rpmRpmAlternateContentSourceHref := "rpmRpmAlternateContentSourceHref_example" // string | 
    fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
    excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AcsRpmAPI.AcsRpmRpmMyPermissions(context.Background(), rpmRpmAlternateContentSourceHref).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AcsRpmAPI.AcsRpmRpmMyPermissions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AcsRpmRpmMyPermissions`: MyPermissionsResponse
    fmt.Fprintf(os.Stdout, "Response from `AcsRpmAPI.AcsRpmRpmMyPermissions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**rpmRpmAlternateContentSourceHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAcsRpmRpmMyPermissionsRequest struct via the builder pattern


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


## AcsRpmRpmPartialUpdate

> AsyncOperationResponse AcsRpmRpmPartialUpdate(ctx, rpmRpmAlternateContentSourceHref).PatchedrpmRpmAlternateContentSource(patchedrpmRpmAlternateContentSource).Execute()

Update a rpm alternate content source



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
    rpmRpmAlternateContentSourceHref := "rpmRpmAlternateContentSourceHref_example" // string | 
    patchedrpmRpmAlternateContentSource := *openapiclient.NewPatchedrpmRpmAlternateContentSource() // PatchedrpmRpmAlternateContentSource | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AcsRpmAPI.AcsRpmRpmPartialUpdate(context.Background(), rpmRpmAlternateContentSourceHref).PatchedrpmRpmAlternateContentSource(patchedrpmRpmAlternateContentSource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AcsRpmAPI.AcsRpmRpmPartialUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AcsRpmRpmPartialUpdate`: AsyncOperationResponse
    fmt.Fprintf(os.Stdout, "Response from `AcsRpmAPI.AcsRpmRpmPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**rpmRpmAlternateContentSourceHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAcsRpmRpmPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedrpmRpmAlternateContentSource** | [**PatchedrpmRpmAlternateContentSource**](PatchedrpmRpmAlternateContentSource.md) |  | 

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


## AcsRpmRpmRead

> RpmRpmAlternateContentSourceResponse AcsRpmRpmRead(ctx, rpmRpmAlternateContentSourceHref).Fields(fields).ExcludeFields(excludeFields).Execute()

Inspect a rpm alternate content source



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
    rpmRpmAlternateContentSourceHref := "rpmRpmAlternateContentSourceHref_example" // string | 
    fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
    excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AcsRpmAPI.AcsRpmRpmRead(context.Background(), rpmRpmAlternateContentSourceHref).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AcsRpmAPI.AcsRpmRpmRead``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AcsRpmRpmRead`: RpmRpmAlternateContentSourceResponse
    fmt.Fprintf(os.Stdout, "Response from `AcsRpmAPI.AcsRpmRpmRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**rpmRpmAlternateContentSourceHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAcsRpmRpmReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

### Return type

[**RpmRpmAlternateContentSourceResponse**](RpmRpmAlternateContentSourceResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AcsRpmRpmRefresh

> TaskGroupOperationResponse AcsRpmRpmRefresh(ctx, rpmRpmAlternateContentSourceHref).Execute()





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
    rpmRpmAlternateContentSourceHref := "rpmRpmAlternateContentSourceHref_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AcsRpmAPI.AcsRpmRpmRefresh(context.Background(), rpmRpmAlternateContentSourceHref).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AcsRpmAPI.AcsRpmRpmRefresh``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AcsRpmRpmRefresh`: TaskGroupOperationResponse
    fmt.Fprintf(os.Stdout, "Response from `AcsRpmAPI.AcsRpmRpmRefresh`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**rpmRpmAlternateContentSourceHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAcsRpmRpmRefreshRequest struct via the builder pattern


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


## AcsRpmRpmRemoveRole

> NestedRoleResponse AcsRpmRpmRemoveRole(ctx, rpmRpmAlternateContentSourceHref).NestedRole(nestedRole).Execute()





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
    rpmRpmAlternateContentSourceHref := "rpmRpmAlternateContentSourceHref_example" // string | 
    nestedRole := *openapiclient.NewNestedRole("Role_example") // NestedRole | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AcsRpmAPI.AcsRpmRpmRemoveRole(context.Background(), rpmRpmAlternateContentSourceHref).NestedRole(nestedRole).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AcsRpmAPI.AcsRpmRpmRemoveRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AcsRpmRpmRemoveRole`: NestedRoleResponse
    fmt.Fprintf(os.Stdout, "Response from `AcsRpmAPI.AcsRpmRpmRemoveRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**rpmRpmAlternateContentSourceHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAcsRpmRpmRemoveRoleRequest struct via the builder pattern


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


## AcsRpmRpmUpdate

> AsyncOperationResponse AcsRpmRpmUpdate(ctx, rpmRpmAlternateContentSourceHref).RpmRpmAlternateContentSource(rpmRpmAlternateContentSource).Execute()

Update a rpm alternate content source



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
    rpmRpmAlternateContentSourceHref := "rpmRpmAlternateContentSourceHref_example" // string | 
    rpmRpmAlternateContentSource := *openapiclient.NewRpmRpmAlternateContentSource("Name_example", "Remote_example") // RpmRpmAlternateContentSource | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AcsRpmAPI.AcsRpmRpmUpdate(context.Background(), rpmRpmAlternateContentSourceHref).RpmRpmAlternateContentSource(rpmRpmAlternateContentSource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AcsRpmAPI.AcsRpmRpmUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AcsRpmRpmUpdate`: AsyncOperationResponse
    fmt.Fprintf(os.Stdout, "Response from `AcsRpmAPI.AcsRpmRpmUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**rpmRpmAlternateContentSourceHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAcsRpmRpmUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **rpmRpmAlternateContentSource** | [**RpmRpmAlternateContentSource**](RpmRpmAlternateContentSource.md) |  | 

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

