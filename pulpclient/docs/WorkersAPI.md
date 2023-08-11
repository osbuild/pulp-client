# \WorkersAPI

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**WorkersList**](WorkersAPI.md#WorkersList) | **Get** /pulp/api/v3/workers/ | List workers
[**WorkersRead**](WorkersAPI.md#WorkersRead) | **Get** {worker_href} | Inspect a worker



## WorkersList

> PaginatedWorkerResponseList WorkersList(ctx).LastHeartbeat(lastHeartbeat).LastHeartbeatGt(lastHeartbeatGt).LastHeartbeatGte(lastHeartbeatGte).LastHeartbeatLt(lastHeartbeatLt).LastHeartbeatLte(lastHeartbeatLte).LastHeartbeatRange(lastHeartbeatRange).Limit(limit).Missing(missing).Name(name).NameContains(nameContains).NameIcontains(nameIcontains).NameIn(nameIn).NameStartswith(nameStartswith).Offset(offset).Online(online).Ordering(ordering).PulpHrefIn(pulpHrefIn).PulpIdIn(pulpIdIn).Fields(fields).ExcludeFields(excludeFields).Execute()

List workers



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
    lastHeartbeat := time.Now() // time.Time | Filter results where last_heartbeat matches value (optional)
    lastHeartbeatGt := time.Now() // time.Time | Filter results where last_heartbeat is greater than value (optional)
    lastHeartbeatGte := time.Now() // time.Time | Filter results where last_heartbeat is greater than or equal to value (optional)
    lastHeartbeatLt := time.Now() // time.Time | Filter results where last_heartbeat is less than value (optional)
    lastHeartbeatLte := time.Now() // time.Time | Filter results where last_heartbeat is less than or equal to value (optional)
    lastHeartbeatRange := []time.Time{time.Now()} // []time.Time | Filter results where last_heartbeat is between two comma separated values (optional)
    limit := int32(56) // int32 | Number of results to return per page. (optional)
    missing := true // bool |  (optional)
    name := "name_example" // string | Filter results where name matches value (optional)
    nameContains := "nameContains_example" // string | Filter results where name contains value (optional)
    nameIcontains := "nameIcontains_example" // string | Filter results where name contains value (optional)
    nameIn := []string{"Inner_example"} // []string | Filter results where name is in a comma-separated list of values (optional)
    nameStartswith := "nameStartswith_example" // string | Filter results where name starts with value (optional)
    offset := int32(56) // int32 | The initial index from which to return the results. (optional)
    online := true // bool |  (optional)
    ordering := []string{"Ordering_example"} // []string | Ordering  * `pulp_id` - Pulp id * `-pulp_id` - Pulp id (descending) * `pulp_created` - Pulp created * `-pulp_created` - Pulp created (descending) * `pulp_last_updated` - Pulp last updated * `-pulp_last_updated` - Pulp last updated (descending) * `name` - Name * `-name` - Name (descending) * `last_heartbeat` - Last heartbeat * `-last_heartbeat` - Last heartbeat (descending) * `versions` - Versions * `-versions` - Versions (descending) * `pk` - Pk * `-pk` - Pk (descending) (optional)
    pulpHrefIn := []string{"Inner_example"} // []string | Multiple values may be separated by commas. (optional)
    pulpIdIn := []string{"Inner_example"} // []string | Multiple values may be separated by commas. (optional)
    fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
    excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkersAPI.WorkersList(context.Background()).LastHeartbeat(lastHeartbeat).LastHeartbeatGt(lastHeartbeatGt).LastHeartbeatGte(lastHeartbeatGte).LastHeartbeatLt(lastHeartbeatLt).LastHeartbeatLte(lastHeartbeatLte).LastHeartbeatRange(lastHeartbeatRange).Limit(limit).Missing(missing).Name(name).NameContains(nameContains).NameIcontains(nameIcontains).NameIn(nameIn).NameStartswith(nameStartswith).Offset(offset).Online(online).Ordering(ordering).PulpHrefIn(pulpHrefIn).PulpIdIn(pulpIdIn).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkersAPI.WorkersList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `WorkersList`: PaginatedWorkerResponseList
    fmt.Fprintf(os.Stdout, "Response from `WorkersAPI.WorkersList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWorkersListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **lastHeartbeat** | **time.Time** | Filter results where last_heartbeat matches value | 
 **lastHeartbeatGt** | **time.Time** | Filter results where last_heartbeat is greater than value | 
 **lastHeartbeatGte** | **time.Time** | Filter results where last_heartbeat is greater than or equal to value | 
 **lastHeartbeatLt** | **time.Time** | Filter results where last_heartbeat is less than value | 
 **lastHeartbeatLte** | **time.Time** | Filter results where last_heartbeat is less than or equal to value | 
 **lastHeartbeatRange** | [**[]time.Time**](time.Time.md) | Filter results where last_heartbeat is between two comma separated values | 
 **limit** | **int32** | Number of results to return per page. | 
 **missing** | **bool** |  | 
 **name** | **string** | Filter results where name matches value | 
 **nameContains** | **string** | Filter results where name contains value | 
 **nameIcontains** | **string** | Filter results where name contains value | 
 **nameIn** | **[]string** | Filter results where name is in a comma-separated list of values | 
 **nameStartswith** | **string** | Filter results where name starts with value | 
 **offset** | **int32** | The initial index from which to return the results. | 
 **online** | **bool** |  | 
 **ordering** | **[]string** | Ordering  * &#x60;pulp_id&#x60; - Pulp id * &#x60;-pulp_id&#x60; - Pulp id (descending) * &#x60;pulp_created&#x60; - Pulp created * &#x60;-pulp_created&#x60; - Pulp created (descending) * &#x60;pulp_last_updated&#x60; - Pulp last updated * &#x60;-pulp_last_updated&#x60; - Pulp last updated (descending) * &#x60;name&#x60; - Name * &#x60;-name&#x60; - Name (descending) * &#x60;last_heartbeat&#x60; - Last heartbeat * &#x60;-last_heartbeat&#x60; - Last heartbeat (descending) * &#x60;versions&#x60; - Versions * &#x60;-versions&#x60; - Versions (descending) * &#x60;pk&#x60; - Pk * &#x60;-pk&#x60; - Pk (descending) | 
 **pulpHrefIn** | **[]string** | Multiple values may be separated by commas. | 
 **pulpIdIn** | **[]string** | Multiple values may be separated by commas. | 
 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

### Return type

[**PaginatedWorkerResponseList**](PaginatedWorkerResponseList.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WorkersRead

> WorkerResponse WorkersRead(ctx, workerHref).Fields(fields).ExcludeFields(excludeFields).Execute()

Inspect a worker



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
    workerHref := "workerHref_example" // string | 
    fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
    excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkersAPI.WorkersRead(context.Background(), workerHref).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkersAPI.WorkersRead``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `WorkersRead`: WorkerResponse
    fmt.Fprintf(os.Stdout, "Response from `WorkersAPI.WorkersRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workerHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiWorkersReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

### Return type

[**WorkerResponse**](WorkerResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

