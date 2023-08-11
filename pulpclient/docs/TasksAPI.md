# \TasksAPI

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TasksAddRole**](TasksAPI.md#TasksAddRole) | **Post** {task_href}add_role/ | 
[**TasksCancel**](TasksAPI.md#TasksCancel) | **Patch** {task_href} | Cancel a task
[**TasksDelete**](TasksAPI.md#TasksDelete) | **Delete** {task_href} | Delete a task
[**TasksList**](TasksAPI.md#TasksList) | **Get** /pulp/api/v3/tasks/ | List tasks
[**TasksListRoles**](TasksAPI.md#TasksListRoles) | **Get** {task_href}list_roles/ | 
[**TasksMyPermissions**](TasksAPI.md#TasksMyPermissions) | **Get** {task_href}my_permissions/ | 
[**TasksPurge**](TasksAPI.md#TasksPurge) | **Post** /pulp/api/v3/tasks/purge/ | Purge Completed Tasks
[**TasksRead**](TasksAPI.md#TasksRead) | **Get** {task_href} | Inspect a task
[**TasksRemoveRole**](TasksAPI.md#TasksRemoveRole) | **Post** {task_href}remove_role/ | 



## TasksAddRole

> NestedRoleResponse TasksAddRole(ctx, taskHref).NestedRole(nestedRole).Execute()





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
    taskHref := "taskHref_example" // string | 
    nestedRole := *openapiclient.NewNestedRole("Role_example") // NestedRole | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TasksAPI.TasksAddRole(context.Background(), taskHref).NestedRole(nestedRole).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TasksAPI.TasksAddRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TasksAddRole`: NestedRoleResponse
    fmt.Fprintf(os.Stdout, "Response from `TasksAPI.TasksAddRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTasksAddRoleRequest struct via the builder pattern


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


## TasksCancel

> TaskResponse TasksCancel(ctx, taskHref).PatchedTaskCancel(patchedTaskCancel).Execute()

Cancel a task



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
    taskHref := "taskHref_example" // string | 
    patchedTaskCancel := *openapiclient.NewPatchedTaskCancel() // PatchedTaskCancel | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TasksAPI.TasksCancel(context.Background(), taskHref).PatchedTaskCancel(patchedTaskCancel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TasksAPI.TasksCancel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TasksCancel`: TaskResponse
    fmt.Fprintf(os.Stdout, "Response from `TasksAPI.TasksCancel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTasksCancelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedTaskCancel** | [**PatchedTaskCancel**](PatchedTaskCancel.md) |  | 

### Return type

[**TaskResponse**](TaskResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TasksDelete

> TasksDelete(ctx, taskHref).Execute()

Delete a task



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
    taskHref := "taskHref_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.TasksAPI.TasksDelete(context.Background(), taskHref).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TasksAPI.TasksDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTasksDeleteRequest struct via the builder pattern


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


## TasksList

> PaginatedTaskResponseList TasksList(ctx).ChildTasks(childTasks).CreatedResources(createdResources).ExclusiveResources(exclusiveResources).ExclusiveResourcesIn(exclusiveResourcesIn).FinishedAt(finishedAt).FinishedAtGt(finishedAtGt).FinishedAtGte(finishedAtGte).FinishedAtLt(finishedAtLt).FinishedAtLte(finishedAtLte).FinishedAtRange(finishedAtRange).Limit(limit).LoggingCid(loggingCid).LoggingCidContains(loggingCidContains).Name(name).NameContains(nameContains).NameIn(nameIn).Offset(offset).Ordering(ordering).ParentTask(parentTask).PulpHrefIn(pulpHrefIn).PulpIdIn(pulpIdIn).ReservedResources(reservedResources).ReservedResourcesIn(reservedResourcesIn).ReservedResourcesRecord(reservedResourcesRecord).SharedResources(sharedResources).SharedResourcesIn(sharedResourcesIn).StartedAt(startedAt).StartedAtGt(startedAtGt).StartedAtGte(startedAtGte).StartedAtLt(startedAtLt).StartedAtLte(startedAtLte).StartedAtRange(startedAtRange).State(state).StateIn(stateIn).TaskGroup(taskGroup).Worker(worker).WorkerIn(workerIn).Fields(fields).ExcludeFields(excludeFields).Execute()

List tasks



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "github.com/osbuild/pulp-client/pulpclient"
)

func main() {
    childTasks := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Filter results where child_tasks matches value (optional)
    createdResources := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
    exclusiveResources := "exclusiveResources_example" // string |  (optional)
    exclusiveResourcesIn := []string{"Inner_example"} // []string | Multiple values may be separated by commas. (optional)
    finishedAt := time.Now() // time.Time | Filter results where finished_at matches value (optional)
    finishedAtGt := time.Now() // time.Time | Filter results where finished_at is greater than value (optional)
    finishedAtGte := time.Now() // time.Time | Filter results where finished_at is greater than or equal to value (optional)
    finishedAtLt := time.Now() // time.Time | Filter results where finished_at is less than value (optional)
    finishedAtLte := time.Now() // time.Time | Filter results where finished_at is less than or equal to value (optional)
    finishedAtRange := []time.Time{time.Now()} // []time.Time | Filter results where finished_at is between two comma separated values (optional)
    limit := int32(56) // int32 | Number of results to return per page. (optional)
    loggingCid := "loggingCid_example" // string | Filter results where logging_cid matches value (optional)
    loggingCidContains := "loggingCidContains_example" // string | Filter results where logging_cid contains value (optional)
    name := "name_example" // string | Filter results where name matches value (optional)
    nameContains := "nameContains_example" // string | Filter results where name contains value (optional)
    nameIn := []string{"Inner_example"} // []string | Filter results where name is in a comma-separated list of values (optional)
    offset := int32(56) // int32 | The initial index from which to return the results. (optional)
    ordering := []string{"Ordering_example"} // []string | Ordering  * `pulp_id` - Pulp id * `-pulp_id` - Pulp id (descending) * `pulp_created` - Pulp created * `-pulp_created` - Pulp created (descending) * `pulp_last_updated` - Pulp last updated * `-pulp_last_updated` - Pulp last updated (descending) * `state` - State * `-state` - State (descending) * `name` - Name * `-name` - Name (descending) * `logging_cid` - Logging cid * `-logging_cid` - Logging cid (descending) * `started_at` - Started at * `-started_at` - Started at (descending) * `finished_at` - Finished at * `-finished_at` - Finished at (descending) * `error` - Error * `-error` - Error (descending) * `args` - Args * `-args` - Args (descending) * `kwargs` - Kwargs * `-kwargs` - Kwargs (descending) * `reserved_resources_record` - Reserved resources record * `-reserved_resources_record` - Reserved resources record (descending) * `versions` - Versions * `-versions` - Versions (descending) * `pk` - Pk * `-pk` - Pk (descending) (optional)
    parentTask := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Filter results where parent_task matches value (optional)
    pulpHrefIn := []string{"Inner_example"} // []string | Multiple values may be separated by commas. (optional)
    pulpIdIn := []string{"Inner_example"} // []string | Multiple values may be separated by commas. (optional)
    reservedResources := "reservedResources_example" // string |  (optional)
    reservedResourcesIn := []string{"Inner_example"} // []string | Multiple values may be separated by commas. (optional)
    reservedResourcesRecord := []string{"Inner_example"} // []string |  (optional)
    sharedResources := "sharedResources_example" // string |  (optional)
    sharedResourcesIn := []string{"Inner_example"} // []string | Multiple values may be separated by commas. (optional)
    startedAt := time.Now() // time.Time | Filter results where started_at matches value (optional)
    startedAtGt := time.Now() // time.Time | Filter results where started_at is greater than value (optional)
    startedAtGte := time.Now() // time.Time | Filter results where started_at is greater than or equal to value (optional)
    startedAtLt := time.Now() // time.Time | Filter results where started_at is less than value (optional)
    startedAtLte := time.Now() // time.Time | Filter results where started_at is less than or equal to value (optional)
    startedAtRange := []time.Time{time.Now()} // []time.Time | Filter results where started_at is between two comma separated values (optional)
    state := "state_example" // string | Filter results where state matches value  * `waiting` - Waiting * `skipped` - Skipped * `running` - Running * `completed` - Completed * `failed` - Failed * `canceled` - Canceled * `canceling` - Canceling (optional)
    stateIn := []string{"Inner_example"} // []string | Filter results where state is in a comma-separated list of values (optional)
    taskGroup := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Filter results where task_group matches value (optional)
    worker := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Filter results where worker matches value (optional)
    workerIn := []string{"Inner_example"} // []string | Filter results where worker is in a comma-separated list of values (optional)
    fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
    excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TasksAPI.TasksList(context.Background()).ChildTasks(childTasks).CreatedResources(createdResources).ExclusiveResources(exclusiveResources).ExclusiveResourcesIn(exclusiveResourcesIn).FinishedAt(finishedAt).FinishedAtGt(finishedAtGt).FinishedAtGte(finishedAtGte).FinishedAtLt(finishedAtLt).FinishedAtLte(finishedAtLte).FinishedAtRange(finishedAtRange).Limit(limit).LoggingCid(loggingCid).LoggingCidContains(loggingCidContains).Name(name).NameContains(nameContains).NameIn(nameIn).Offset(offset).Ordering(ordering).ParentTask(parentTask).PulpHrefIn(pulpHrefIn).PulpIdIn(pulpIdIn).ReservedResources(reservedResources).ReservedResourcesIn(reservedResourcesIn).ReservedResourcesRecord(reservedResourcesRecord).SharedResources(sharedResources).SharedResourcesIn(sharedResourcesIn).StartedAt(startedAt).StartedAtGt(startedAtGt).StartedAtGte(startedAtGte).StartedAtLt(startedAtLt).StartedAtLte(startedAtLte).StartedAtRange(startedAtRange).State(state).StateIn(stateIn).TaskGroup(taskGroup).Worker(worker).WorkerIn(workerIn).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TasksAPI.TasksList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TasksList`: PaginatedTaskResponseList
    fmt.Fprintf(os.Stdout, "Response from `TasksAPI.TasksList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTasksListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **childTasks** | **string** | Filter results where child_tasks matches value | 
 **createdResources** | **string** |  | 
 **exclusiveResources** | **string** |  | 
 **exclusiveResourcesIn** | **[]string** | Multiple values may be separated by commas. | 
 **finishedAt** | **time.Time** | Filter results where finished_at matches value | 
 **finishedAtGt** | **time.Time** | Filter results where finished_at is greater than value | 
 **finishedAtGte** | **time.Time** | Filter results where finished_at is greater than or equal to value | 
 **finishedAtLt** | **time.Time** | Filter results where finished_at is less than value | 
 **finishedAtLte** | **time.Time** | Filter results where finished_at is less than or equal to value | 
 **finishedAtRange** | [**[]time.Time**](time.Time.md) | Filter results where finished_at is between two comma separated values | 
 **limit** | **int32** | Number of results to return per page. | 
 **loggingCid** | **string** | Filter results where logging_cid matches value | 
 **loggingCidContains** | **string** | Filter results where logging_cid contains value | 
 **name** | **string** | Filter results where name matches value | 
 **nameContains** | **string** | Filter results where name contains value | 
 **nameIn** | **[]string** | Filter results where name is in a comma-separated list of values | 
 **offset** | **int32** | The initial index from which to return the results. | 
 **ordering** | **[]string** | Ordering  * &#x60;pulp_id&#x60; - Pulp id * &#x60;-pulp_id&#x60; - Pulp id (descending) * &#x60;pulp_created&#x60; - Pulp created * &#x60;-pulp_created&#x60; - Pulp created (descending) * &#x60;pulp_last_updated&#x60; - Pulp last updated * &#x60;-pulp_last_updated&#x60; - Pulp last updated (descending) * &#x60;state&#x60; - State * &#x60;-state&#x60; - State (descending) * &#x60;name&#x60; - Name * &#x60;-name&#x60; - Name (descending) * &#x60;logging_cid&#x60; - Logging cid * &#x60;-logging_cid&#x60; - Logging cid (descending) * &#x60;started_at&#x60; - Started at * &#x60;-started_at&#x60; - Started at (descending) * &#x60;finished_at&#x60; - Finished at * &#x60;-finished_at&#x60; - Finished at (descending) * &#x60;error&#x60; - Error * &#x60;-error&#x60; - Error (descending) * &#x60;args&#x60; - Args * &#x60;-args&#x60; - Args (descending) * &#x60;kwargs&#x60; - Kwargs * &#x60;-kwargs&#x60; - Kwargs (descending) * &#x60;reserved_resources_record&#x60; - Reserved resources record * &#x60;-reserved_resources_record&#x60; - Reserved resources record (descending) * &#x60;versions&#x60; - Versions * &#x60;-versions&#x60; - Versions (descending) * &#x60;pk&#x60; - Pk * &#x60;-pk&#x60; - Pk (descending) | 
 **parentTask** | **string** | Filter results where parent_task matches value | 
 **pulpHrefIn** | **[]string** | Multiple values may be separated by commas. | 
 **pulpIdIn** | **[]string** | Multiple values may be separated by commas. | 
 **reservedResources** | **string** |  | 
 **reservedResourcesIn** | **[]string** | Multiple values may be separated by commas. | 
 **reservedResourcesRecord** | **[]string** |  | 
 **sharedResources** | **string** |  | 
 **sharedResourcesIn** | **[]string** | Multiple values may be separated by commas. | 
 **startedAt** | **time.Time** | Filter results where started_at matches value | 
 **startedAtGt** | **time.Time** | Filter results where started_at is greater than value | 
 **startedAtGte** | **time.Time** | Filter results where started_at is greater than or equal to value | 
 **startedAtLt** | **time.Time** | Filter results where started_at is less than value | 
 **startedAtLte** | **time.Time** | Filter results where started_at is less than or equal to value | 
 **startedAtRange** | [**[]time.Time**](time.Time.md) | Filter results where started_at is between two comma separated values | 
 **state** | **string** | Filter results where state matches value  * &#x60;waiting&#x60; - Waiting * &#x60;skipped&#x60; - Skipped * &#x60;running&#x60; - Running * &#x60;completed&#x60; - Completed * &#x60;failed&#x60; - Failed * &#x60;canceled&#x60; - Canceled * &#x60;canceling&#x60; - Canceling | 
 **stateIn** | **[]string** | Filter results where state is in a comma-separated list of values | 
 **taskGroup** | **string** | Filter results where task_group matches value | 
 **worker** | **string** | Filter results where worker matches value | 
 **workerIn** | **[]string** | Filter results where worker is in a comma-separated list of values | 
 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

### Return type

[**PaginatedTaskResponseList**](PaginatedTaskResponseList.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TasksListRoles

> ObjectRolesResponse TasksListRoles(ctx, taskHref).Fields(fields).ExcludeFields(excludeFields).Execute()





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
    taskHref := "taskHref_example" // string | 
    fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
    excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TasksAPI.TasksListRoles(context.Background(), taskHref).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TasksAPI.TasksListRoles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TasksListRoles`: ObjectRolesResponse
    fmt.Fprintf(os.Stdout, "Response from `TasksAPI.TasksListRoles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTasksListRolesRequest struct via the builder pattern


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


## TasksMyPermissions

> MyPermissionsResponse TasksMyPermissions(ctx, taskHref).Fields(fields).ExcludeFields(excludeFields).Execute()





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
    taskHref := "taskHref_example" // string | 
    fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
    excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TasksAPI.TasksMyPermissions(context.Background(), taskHref).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TasksAPI.TasksMyPermissions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TasksMyPermissions`: MyPermissionsResponse
    fmt.Fprintf(os.Stdout, "Response from `TasksAPI.TasksMyPermissions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTasksMyPermissionsRequest struct via the builder pattern


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


## TasksPurge

> AsyncOperationResponse TasksPurge(ctx).Purge(purge).Execute()

Purge Completed Tasks



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
    purge := *openapiclient.NewPurge() // Purge | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TasksAPI.TasksPurge(context.Background()).Purge(purge).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TasksAPI.TasksPurge``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TasksPurge`: AsyncOperationResponse
    fmt.Fprintf(os.Stdout, "Response from `TasksAPI.TasksPurge`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTasksPurgeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **purge** | [**Purge**](Purge.md) |  | 

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


## TasksRead

> TaskResponse TasksRead(ctx, taskHref).Fields(fields).ExcludeFields(excludeFields).Execute()

Inspect a task



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
    taskHref := "taskHref_example" // string | 
    fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
    excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TasksAPI.TasksRead(context.Background(), taskHref).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TasksAPI.TasksRead``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TasksRead`: TaskResponse
    fmt.Fprintf(os.Stdout, "Response from `TasksAPI.TasksRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTasksReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

### Return type

[**TaskResponse**](TaskResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TasksRemoveRole

> NestedRoleResponse TasksRemoveRole(ctx, taskHref).NestedRole(nestedRole).Execute()





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
    taskHref := "taskHref_example" // string | 
    nestedRole := *openapiclient.NewNestedRole("Role_example") // NestedRole | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TasksAPI.TasksRemoveRole(context.Background(), taskHref).NestedRole(nestedRole).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TasksAPI.TasksRemoveRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TasksRemoveRole`: NestedRoleResponse
    fmt.Fprintf(os.Stdout, "Response from `TasksAPI.TasksRemoveRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTasksRemoveRoleRequest struct via the builder pattern


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

