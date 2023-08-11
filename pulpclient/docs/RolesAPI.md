# \RolesAPI

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RolesCreate**](RolesAPI.md#RolesCreate) | **Post** /pulp/api/v3/roles/ | Create a role
[**RolesDelete**](RolesAPI.md#RolesDelete) | **Delete** {role_href} | Delete a role
[**RolesList**](RolesAPI.md#RolesList) | **Get** /pulp/api/v3/roles/ | List roles
[**RolesPartialUpdate**](RolesAPI.md#RolesPartialUpdate) | **Patch** {role_href} | Update a role
[**RolesRead**](RolesAPI.md#RolesRead) | **Get** {role_href} | Inspect a role
[**RolesUpdate**](RolesAPI.md#RolesUpdate) | **Put** {role_href} | Update a role



## RolesCreate

> RoleResponse RolesCreate(ctx).Role(role).Execute()

Create a role



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
    role := *openapiclient.NewRole("Name_example", []string{"Permissions_example"}) // Role | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RolesAPI.RolesCreate(context.Background()).Role(role).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RolesAPI.RolesCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RolesCreate`: RoleResponse
    fmt.Fprintf(os.Stdout, "Response from `RolesAPI.RolesCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRolesCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **role** | [**Role**](Role.md) |  | 

### Return type

[**RoleResponse**](RoleResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RolesDelete

> RolesDelete(ctx, roleHref).Execute()

Delete a role



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
    roleHref := "roleHref_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.RolesAPI.RolesDelete(context.Background(), roleHref).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RolesAPI.RolesDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roleHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRolesDeleteRequest struct via the builder pattern


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


## RolesList

> PaginatedRoleResponseList RolesList(ctx).ContainsPermission(containsPermission).Description(description).DescriptionContains(descriptionContains).DescriptionIcontains(descriptionIcontains).DescriptionIexact(descriptionIexact).ForObjectType(forObjectType).Limit(limit).Locked(locked).Name(name).NameContains(nameContains).NameIcontains(nameIcontains).NameIn(nameIn).NameStartswith(nameStartswith).Offset(offset).Ordering(ordering).PulpHrefIn(pulpHrefIn).PulpIdIn(pulpIdIn).Fields(fields).ExcludeFields(excludeFields).Execute()

List roles



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
    containsPermission := []string{"Inner_example"} // []string | Filter roles that have any of the permissions in the list. (optional)
    description := "description_example" // string | Filter results where description matches value (optional)
    descriptionContains := "descriptionContains_example" // string | Filter results where description contains value (optional)
    descriptionIcontains := "descriptionIcontains_example" // string | Filter results where description contains value (optional)
    descriptionIexact := "descriptionIexact_example" // string | Filter results where description matches value (optional)
    forObjectType := "forObjectType_example" // string | Filter roles that only have permissions for the specified object HREF. (optional)
    limit := int32(56) // int32 | Number of results to return per page. (optional)
    locked := true // bool | Filter results where locked matches value (optional)
    name := "name_example" // string | Filter results where name matches value (optional)
    nameContains := "nameContains_example" // string | Filter results where name contains value (optional)
    nameIcontains := "nameIcontains_example" // string | Filter results where name contains value (optional)
    nameIn := []string{"Inner_example"} // []string | Filter results where name is in a comma-separated list of values (optional)
    nameStartswith := "nameStartswith_example" // string | Filter results where name starts with value (optional)
    offset := int32(56) // int32 | The initial index from which to return the results. (optional)
    ordering := []string{"Ordering_example"} // []string | Ordering  * `pulp_id` - Pulp id * `-pulp_id` - Pulp id (descending) * `pulp_created` - Pulp created * `-pulp_created` - Pulp created (descending) * `pulp_last_updated` - Pulp last updated * `-pulp_last_updated` - Pulp last updated (descending) * `name` - Name * `-name` - Name (descending) * `description` - Description * `-description` - Description (descending) * `locked` - Locked * `-locked` - Locked (descending) * `pk` - Pk * `-pk` - Pk (descending) (optional)
    pulpHrefIn := []string{"Inner_example"} // []string | Multiple values may be separated by commas. (optional)
    pulpIdIn := []string{"Inner_example"} // []string | Multiple values may be separated by commas. (optional)
    fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
    excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RolesAPI.RolesList(context.Background()).ContainsPermission(containsPermission).Description(description).DescriptionContains(descriptionContains).DescriptionIcontains(descriptionIcontains).DescriptionIexact(descriptionIexact).ForObjectType(forObjectType).Limit(limit).Locked(locked).Name(name).NameContains(nameContains).NameIcontains(nameIcontains).NameIn(nameIn).NameStartswith(nameStartswith).Offset(offset).Ordering(ordering).PulpHrefIn(pulpHrefIn).PulpIdIn(pulpIdIn).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RolesAPI.RolesList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RolesList`: PaginatedRoleResponseList
    fmt.Fprintf(os.Stdout, "Response from `RolesAPI.RolesList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRolesListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **containsPermission** | **[]string** | Filter roles that have any of the permissions in the list. | 
 **description** | **string** | Filter results where description matches value | 
 **descriptionContains** | **string** | Filter results where description contains value | 
 **descriptionIcontains** | **string** | Filter results where description contains value | 
 **descriptionIexact** | **string** | Filter results where description matches value | 
 **forObjectType** | **string** | Filter roles that only have permissions for the specified object HREF. | 
 **limit** | **int32** | Number of results to return per page. | 
 **locked** | **bool** | Filter results where locked matches value | 
 **name** | **string** | Filter results where name matches value | 
 **nameContains** | **string** | Filter results where name contains value | 
 **nameIcontains** | **string** | Filter results where name contains value | 
 **nameIn** | **[]string** | Filter results where name is in a comma-separated list of values | 
 **nameStartswith** | **string** | Filter results where name starts with value | 
 **offset** | **int32** | The initial index from which to return the results. | 
 **ordering** | **[]string** | Ordering  * &#x60;pulp_id&#x60; - Pulp id * &#x60;-pulp_id&#x60; - Pulp id (descending) * &#x60;pulp_created&#x60; - Pulp created * &#x60;-pulp_created&#x60; - Pulp created (descending) * &#x60;pulp_last_updated&#x60; - Pulp last updated * &#x60;-pulp_last_updated&#x60; - Pulp last updated (descending) * &#x60;name&#x60; - Name * &#x60;-name&#x60; - Name (descending) * &#x60;description&#x60; - Description * &#x60;-description&#x60; - Description (descending) * &#x60;locked&#x60; - Locked * &#x60;-locked&#x60; - Locked (descending) * &#x60;pk&#x60; - Pk * &#x60;-pk&#x60; - Pk (descending) | 
 **pulpHrefIn** | **[]string** | Multiple values may be separated by commas. | 
 **pulpIdIn** | **[]string** | Multiple values may be separated by commas. | 
 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

### Return type

[**PaginatedRoleResponseList**](PaginatedRoleResponseList.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RolesPartialUpdate

> RoleResponse RolesPartialUpdate(ctx, roleHref).PatchedRole(patchedRole).Execute()

Update a role



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
    roleHref := "roleHref_example" // string | 
    patchedRole := *openapiclient.NewPatchedRole() // PatchedRole | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RolesAPI.RolesPartialUpdate(context.Background(), roleHref).PatchedRole(patchedRole).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RolesAPI.RolesPartialUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RolesPartialUpdate`: RoleResponse
    fmt.Fprintf(os.Stdout, "Response from `RolesAPI.RolesPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roleHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRolesPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedRole** | [**PatchedRole**](PatchedRole.md) |  | 

### Return type

[**RoleResponse**](RoleResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RolesRead

> RoleResponse RolesRead(ctx, roleHref).Fields(fields).ExcludeFields(excludeFields).Execute()

Inspect a role



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
    roleHref := "roleHref_example" // string | 
    fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
    excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RolesAPI.RolesRead(context.Background(), roleHref).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RolesAPI.RolesRead``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RolesRead`: RoleResponse
    fmt.Fprintf(os.Stdout, "Response from `RolesAPI.RolesRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roleHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRolesReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

### Return type

[**RoleResponse**](RoleResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RolesUpdate

> RoleResponse RolesUpdate(ctx, roleHref).Role(role).Execute()

Update a role



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
    roleHref := "roleHref_example" // string | 
    role := *openapiclient.NewRole("Name_example", []string{"Permissions_example"}) // Role | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RolesAPI.RolesUpdate(context.Background(), roleHref).Role(role).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RolesAPI.RolesUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RolesUpdate`: RoleResponse
    fmt.Fprintf(os.Stdout, "Response from `RolesAPI.RolesUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roleHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRolesUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **role** | [**Role**](Role.md) |  | 

### Return type

[**RoleResponse**](RoleResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

