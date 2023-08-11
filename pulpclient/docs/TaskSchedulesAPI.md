# \TaskSchedulesAPI

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TaskSchedulesAddRole**](TaskSchedulesAPI.md#TaskSchedulesAddRole) | **Post** {task_schedule_href}add_role/ | 
[**TaskSchedulesList**](TaskSchedulesAPI.md#TaskSchedulesList) | **Get** /pulp/api/v3/task-schedules/ | List task schedules
[**TaskSchedulesListRoles**](TaskSchedulesAPI.md#TaskSchedulesListRoles) | **Get** {task_schedule_href}list_roles/ | 
[**TaskSchedulesMyPermissions**](TaskSchedulesAPI.md#TaskSchedulesMyPermissions) | **Get** {task_schedule_href}my_permissions/ | 
[**TaskSchedulesRead**](TaskSchedulesAPI.md#TaskSchedulesRead) | **Get** {task_schedule_href} | Inspect a task schedule
[**TaskSchedulesRemoveRole**](TaskSchedulesAPI.md#TaskSchedulesRemoveRole) | **Post** {task_schedule_href}remove_role/ | 



## TaskSchedulesAddRole

> NestedRoleResponse TaskSchedulesAddRole(ctx, taskScheduleHref).NestedRole(nestedRole).Execute()





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
    taskScheduleHref := "taskScheduleHref_example" // string | 
    nestedRole := *openapiclient.NewNestedRole("Role_example") // NestedRole | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TaskSchedulesAPI.TaskSchedulesAddRole(context.Background(), taskScheduleHref).NestedRole(nestedRole).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaskSchedulesAPI.TaskSchedulesAddRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TaskSchedulesAddRole`: NestedRoleResponse
    fmt.Fprintf(os.Stdout, "Response from `TaskSchedulesAPI.TaskSchedulesAddRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskScheduleHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTaskSchedulesAddRoleRequest struct via the builder pattern


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


## TaskSchedulesList

> PaginatedTaskScheduleResponseList TaskSchedulesList(ctx).Limit(limit).Name(name).NameContains(nameContains).Offset(offset).Ordering(ordering).PulpHrefIn(pulpHrefIn).PulpIdIn(pulpIdIn).TaskName(taskName).TaskNameContains(taskNameContains).Fields(fields).ExcludeFields(excludeFields).Execute()

List task schedules



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
    offset := int32(56) // int32 | The initial index from which to return the results. (optional)
    ordering := []string{"Ordering_example"} // []string | Ordering  * `pulp_id` - Pulp id * `-pulp_id` - Pulp id (descending) * `pulp_created` - Pulp created * `-pulp_created` - Pulp created (descending) * `pulp_last_updated` - Pulp last updated * `-pulp_last_updated` - Pulp last updated (descending) * `name` - Name * `-name` - Name (descending) * `next_dispatch` - Next dispatch * `-next_dispatch` - Next dispatch (descending) * `dispatch_interval` - Dispatch interval * `-dispatch_interval` - Dispatch interval (descending) * `task_name` - Task name * `-task_name` - Task name (descending) * `pk` - Pk * `-pk` - Pk (descending) (optional)
    pulpHrefIn := []string{"Inner_example"} // []string | Multiple values may be separated by commas. (optional)
    pulpIdIn := []string{"Inner_example"} // []string | Multiple values may be separated by commas. (optional)
    taskName := "taskName_example" // string | Filter results where task_name matches value (optional)
    taskNameContains := "taskNameContains_example" // string | Filter results where task_name contains value (optional)
    fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
    excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TaskSchedulesAPI.TaskSchedulesList(context.Background()).Limit(limit).Name(name).NameContains(nameContains).Offset(offset).Ordering(ordering).PulpHrefIn(pulpHrefIn).PulpIdIn(pulpIdIn).TaskName(taskName).TaskNameContains(taskNameContains).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaskSchedulesAPI.TaskSchedulesList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TaskSchedulesList`: PaginatedTaskScheduleResponseList
    fmt.Fprintf(os.Stdout, "Response from `TaskSchedulesAPI.TaskSchedulesList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTaskSchedulesListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Number of results to return per page. | 
 **name** | **string** | Filter results where name matches value | 
 **nameContains** | **string** | Filter results where name contains value | 
 **offset** | **int32** | The initial index from which to return the results. | 
 **ordering** | **[]string** | Ordering  * &#x60;pulp_id&#x60; - Pulp id * &#x60;-pulp_id&#x60; - Pulp id (descending) * &#x60;pulp_created&#x60; - Pulp created * &#x60;-pulp_created&#x60; - Pulp created (descending) * &#x60;pulp_last_updated&#x60; - Pulp last updated * &#x60;-pulp_last_updated&#x60; - Pulp last updated (descending) * &#x60;name&#x60; - Name * &#x60;-name&#x60; - Name (descending) * &#x60;next_dispatch&#x60; - Next dispatch * &#x60;-next_dispatch&#x60; - Next dispatch (descending) * &#x60;dispatch_interval&#x60; - Dispatch interval * &#x60;-dispatch_interval&#x60; - Dispatch interval (descending) * &#x60;task_name&#x60; - Task name * &#x60;-task_name&#x60; - Task name (descending) * &#x60;pk&#x60; - Pk * &#x60;-pk&#x60; - Pk (descending) | 
 **pulpHrefIn** | **[]string** | Multiple values may be separated by commas. | 
 **pulpIdIn** | **[]string** | Multiple values may be separated by commas. | 
 **taskName** | **string** | Filter results where task_name matches value | 
 **taskNameContains** | **string** | Filter results where task_name contains value | 
 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

### Return type

[**PaginatedTaskScheduleResponseList**](PaginatedTaskScheduleResponseList.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TaskSchedulesListRoles

> ObjectRolesResponse TaskSchedulesListRoles(ctx, taskScheduleHref).Fields(fields).ExcludeFields(excludeFields).Execute()





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
    taskScheduleHref := "taskScheduleHref_example" // string | 
    fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
    excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TaskSchedulesAPI.TaskSchedulesListRoles(context.Background(), taskScheduleHref).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaskSchedulesAPI.TaskSchedulesListRoles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TaskSchedulesListRoles`: ObjectRolesResponse
    fmt.Fprintf(os.Stdout, "Response from `TaskSchedulesAPI.TaskSchedulesListRoles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskScheduleHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTaskSchedulesListRolesRequest struct via the builder pattern


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


## TaskSchedulesMyPermissions

> MyPermissionsResponse TaskSchedulesMyPermissions(ctx, taskScheduleHref).Fields(fields).ExcludeFields(excludeFields).Execute()





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
    taskScheduleHref := "taskScheduleHref_example" // string | 
    fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
    excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TaskSchedulesAPI.TaskSchedulesMyPermissions(context.Background(), taskScheduleHref).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaskSchedulesAPI.TaskSchedulesMyPermissions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TaskSchedulesMyPermissions`: MyPermissionsResponse
    fmt.Fprintf(os.Stdout, "Response from `TaskSchedulesAPI.TaskSchedulesMyPermissions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskScheduleHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTaskSchedulesMyPermissionsRequest struct via the builder pattern


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


## TaskSchedulesRead

> TaskScheduleResponse TaskSchedulesRead(ctx, taskScheduleHref).Fields(fields).ExcludeFields(excludeFields).Execute()

Inspect a task schedule



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
    taskScheduleHref := "taskScheduleHref_example" // string | 
    fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
    excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TaskSchedulesAPI.TaskSchedulesRead(context.Background(), taskScheduleHref).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaskSchedulesAPI.TaskSchedulesRead``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TaskSchedulesRead`: TaskScheduleResponse
    fmt.Fprintf(os.Stdout, "Response from `TaskSchedulesAPI.TaskSchedulesRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskScheduleHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTaskSchedulesReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

### Return type

[**TaskScheduleResponse**](TaskScheduleResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TaskSchedulesRemoveRole

> NestedRoleResponse TaskSchedulesRemoveRole(ctx, taskScheduleHref).NestedRole(nestedRole).Execute()





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
    taskScheduleHref := "taskScheduleHref_example" // string | 
    nestedRole := *openapiclient.NewNestedRole("Role_example") // NestedRole | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TaskSchedulesAPI.TaskSchedulesRemoveRole(context.Background(), taskScheduleHref).NestedRole(nestedRole).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaskSchedulesAPI.TaskSchedulesRemoveRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TaskSchedulesRemoveRole`: NestedRoleResponse
    fmt.Fprintf(os.Stdout, "Response from `TaskSchedulesAPI.TaskSchedulesRemoveRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskScheduleHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTaskSchedulesRemoveRoleRequest struct via the builder pattern


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

