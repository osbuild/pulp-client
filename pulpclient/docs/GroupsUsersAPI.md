# \GroupsUsersAPI

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GroupsUsersCreate**](GroupsUsersAPI.md#GroupsUsersCreate) | **Post** {group_href}users/ | Create an user
[**GroupsUsersDelete**](GroupsUsersAPI.md#GroupsUsersDelete) | **Delete** {groups_user_href} | Delete an user
[**GroupsUsersList**](GroupsUsersAPI.md#GroupsUsersList) | **Get** {group_href}users/ | List users



## GroupsUsersCreate

> GroupUserResponse GroupsUsersCreate(ctx, groupHref).GroupUser(groupUser).Execute()

Create an user



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
    groupHref := "groupHref_example" // string | 
    groupUser := *openapiclient.NewGroupUser("Username_example") // GroupUser | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GroupsUsersAPI.GroupsUsersCreate(context.Background(), groupHref).GroupUser(groupUser).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsUsersAPI.GroupsUsersCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GroupsUsersCreate`: GroupUserResponse
    fmt.Fprintf(os.Stdout, "Response from `GroupsUsersAPI.GroupsUsersCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGroupsUsersCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **groupUser** | [**GroupUser**](GroupUser.md) |  | 

### Return type

[**GroupUserResponse**](GroupUserResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GroupsUsersDelete

> GroupsUsersDelete(ctx, groupsUserHref).Execute()

Delete an user



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
    groupsUserHref := "groupsUserHref_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.GroupsUsersAPI.GroupsUsersDelete(context.Background(), groupsUserHref).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsUsersAPI.GroupsUsersDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupsUserHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGroupsUsersDeleteRequest struct via the builder pattern


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


## GroupsUsersList

> PaginatedGroupUserResponseList GroupsUsersList(ctx, groupHref).Limit(limit).Offset(offset).Fields(fields).ExcludeFields(excludeFields).Execute()

List users



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
    groupHref := "groupHref_example" // string | 
    limit := int32(56) // int32 | Number of results to return per page. (optional)
    offset := int32(56) // int32 | The initial index from which to return the results. (optional)
    fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
    excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GroupsUsersAPI.GroupsUsersList(context.Background(), groupHref).Limit(limit).Offset(offset).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsUsersAPI.GroupsUsersList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GroupsUsersList`: PaginatedGroupUserResponseList
    fmt.Fprintf(os.Stdout, "Response from `GroupsUsersAPI.GroupsUsersList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGroupsUsersListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Number of results to return per page. | 
 **offset** | **int32** | The initial index from which to return the results. | 
 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

### Return type

[**PaginatedGroupUserResponseList**](PaginatedGroupUserResponseList.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

