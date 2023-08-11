# \UsersRolesAPI

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UsersRolesCreate**](UsersRolesAPI.md#UsersRolesCreate) | **Post** {auth_user_href}roles/ | Create an user role
[**UsersRolesDelete**](UsersRolesAPI.md#UsersRolesDelete) | **Delete** {auth_users_user_role_href} | Delete an user role
[**UsersRolesList**](UsersRolesAPI.md#UsersRolesList) | **Get** {auth_user_href}roles/ | List user roles
[**UsersRolesRead**](UsersRolesAPI.md#UsersRolesRead) | **Get** {auth_users_user_role_href} | Inspect an user role



## UsersRolesCreate

> UserRoleResponse UsersRolesCreate(ctx, authUserHref).UserRole(userRole).Execute()

Create an user role



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
    authUserHref := "authUserHref_example" // string | 
    userRole := *openapiclient.NewUserRole("Role_example", "ContentObject_example") // UserRole | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersRolesAPI.UsersRolesCreate(context.Background(), authUserHref).UserRole(userRole).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersRolesAPI.UsersRolesCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersRolesCreate`: UserRoleResponse
    fmt.Fprintf(os.Stdout, "Response from `UsersRolesAPI.UsersRolesCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authUserHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersRolesCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userRole** | [**UserRole**](UserRole.md) |  | 

### Return type

[**UserRoleResponse**](UserRoleResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersRolesDelete

> UsersRolesDelete(ctx, authUsersUserRoleHref).Execute()

Delete an user role



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
    authUsersUserRoleHref := "authUsersUserRoleHref_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.UsersRolesAPI.UsersRolesDelete(context.Background(), authUsersUserRoleHref).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersRolesAPI.UsersRolesDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authUsersUserRoleHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersRolesDeleteRequest struct via the builder pattern


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


## UsersRolesList

> PaginatedUserRoleResponseList UsersRolesList(ctx, authUserHref).ContentObject(contentObject).Domain(domain).Limit(limit).Offset(offset).Ordering(ordering).PulpHrefIn(pulpHrefIn).PulpIdIn(pulpIdIn).Role(role).RoleContains(roleContains).RoleIcontains(roleIcontains).RoleIn(roleIn).RoleStartswith(roleStartswith).Fields(fields).ExcludeFields(excludeFields).Execute()

List user roles



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
    authUserHref := "authUserHref_example" // string | 
    contentObject := "contentObject_example" // string | content_object (optional)
    domain := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Foreign Key referenced by HREF (optional)
    limit := int32(56) // int32 | Number of results to return per page. (optional)
    offset := int32(56) // int32 | The initial index from which to return the results. (optional)
    ordering := []string{"Ordering_example"} // []string | Ordering  * `role` - Role * `-role` - Role (descending) * `description` - Description * `-description` - Description (descending) * `pulp_created` - Pulp created * `-pulp_created` - Pulp created (descending) * `pk` - Pk * `-pk` - Pk (descending) (optional)
    pulpHrefIn := []string{"Inner_example"} // []string | Multiple values may be separated by commas. (optional)
    pulpIdIn := []string{"Inner_example"} // []string | Multiple values may be separated by commas. (optional)
    role := "role_example" // string |  (optional)
    roleContains := "roleContains_example" // string |  (optional)
    roleIcontains := "roleIcontains_example" // string |  (optional)
    roleIn := []string{"Inner_example"} // []string | Multiple values may be separated by commas. (optional)
    roleStartswith := "roleStartswith_example" // string |  (optional)
    fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
    excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersRolesAPI.UsersRolesList(context.Background(), authUserHref).ContentObject(contentObject).Domain(domain).Limit(limit).Offset(offset).Ordering(ordering).PulpHrefIn(pulpHrefIn).PulpIdIn(pulpIdIn).Role(role).RoleContains(roleContains).RoleIcontains(roleIcontains).RoleIn(roleIn).RoleStartswith(roleStartswith).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersRolesAPI.UsersRolesList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersRolesList`: PaginatedUserRoleResponseList
    fmt.Fprintf(os.Stdout, "Response from `UsersRolesAPI.UsersRolesList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authUserHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersRolesListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **contentObject** | **string** | content_object | 
 **domain** | **string** | Foreign Key referenced by HREF | 
 **limit** | **int32** | Number of results to return per page. | 
 **offset** | **int32** | The initial index from which to return the results. | 
 **ordering** | **[]string** | Ordering  * &#x60;role&#x60; - Role * &#x60;-role&#x60; - Role (descending) * &#x60;description&#x60; - Description * &#x60;-description&#x60; - Description (descending) * &#x60;pulp_created&#x60; - Pulp created * &#x60;-pulp_created&#x60; - Pulp created (descending) * &#x60;pk&#x60; - Pk * &#x60;-pk&#x60; - Pk (descending) | 
 **pulpHrefIn** | **[]string** | Multiple values may be separated by commas. | 
 **pulpIdIn** | **[]string** | Multiple values may be separated by commas. | 
 **role** | **string** |  | 
 **roleContains** | **string** |  | 
 **roleIcontains** | **string** |  | 
 **roleIn** | **[]string** | Multiple values may be separated by commas. | 
 **roleStartswith** | **string** |  | 
 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

### Return type

[**PaginatedUserRoleResponseList**](PaginatedUserRoleResponseList.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersRolesRead

> UserRoleResponse UsersRolesRead(ctx, authUsersUserRoleHref).Fields(fields).ExcludeFields(excludeFields).Execute()

Inspect an user role



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
    authUsersUserRoleHref := "authUsersUserRoleHref_example" // string | 
    fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
    excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersRolesAPI.UsersRolesRead(context.Background(), authUsersUserRoleHref).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersRolesAPI.UsersRolesRead``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersRolesRead`: UserRoleResponse
    fmt.Fprintf(os.Stdout, "Response from `UsersRolesAPI.UsersRolesRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authUsersUserRoleHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersRolesReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

### Return type

[**UserRoleResponse**](UserRoleResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

