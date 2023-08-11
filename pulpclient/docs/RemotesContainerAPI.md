# \RemotesContainerAPI

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RemotesContainerContainerAddRole**](RemotesContainerAPI.md#RemotesContainerContainerAddRole) | **Post** {container_container_remote_href}add_role/ | 
[**RemotesContainerContainerCreate**](RemotesContainerAPI.md#RemotesContainerContainerCreate) | **Post** /pulp/api/v3/remotes/container/container/ | Create a container remote
[**RemotesContainerContainerDelete**](RemotesContainerAPI.md#RemotesContainerContainerDelete) | **Delete** {container_container_remote_href} | Delete a container remote
[**RemotesContainerContainerList**](RemotesContainerAPI.md#RemotesContainerContainerList) | **Get** /pulp/api/v3/remotes/container/container/ | List container remotes
[**RemotesContainerContainerListRoles**](RemotesContainerAPI.md#RemotesContainerContainerListRoles) | **Get** {container_container_remote_href}list_roles/ | 
[**RemotesContainerContainerMyPermissions**](RemotesContainerAPI.md#RemotesContainerContainerMyPermissions) | **Get** {container_container_remote_href}my_permissions/ | 
[**RemotesContainerContainerPartialUpdate**](RemotesContainerAPI.md#RemotesContainerContainerPartialUpdate) | **Patch** {container_container_remote_href} | Update a container remote
[**RemotesContainerContainerRead**](RemotesContainerAPI.md#RemotesContainerContainerRead) | **Get** {container_container_remote_href} | Inspect a container remote
[**RemotesContainerContainerRemoveRole**](RemotesContainerAPI.md#RemotesContainerContainerRemoveRole) | **Post** {container_container_remote_href}remove_role/ | 
[**RemotesContainerContainerUpdate**](RemotesContainerAPI.md#RemotesContainerContainerUpdate) | **Put** {container_container_remote_href} | Update a container remote



## RemotesContainerContainerAddRole

> NestedRoleResponse RemotesContainerContainerAddRole(ctx, containerContainerRemoteHref).NestedRole(nestedRole).Execute()





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
    containerContainerRemoteHref := "containerContainerRemoteHref_example" // string | 
    nestedRole := *openapiclient.NewNestedRole("Role_example") // NestedRole | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RemotesContainerAPI.RemotesContainerContainerAddRole(context.Background(), containerContainerRemoteHref).NestedRole(nestedRole).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RemotesContainerAPI.RemotesContainerContainerAddRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RemotesContainerContainerAddRole`: NestedRoleResponse
    fmt.Fprintf(os.Stdout, "Response from `RemotesContainerAPI.RemotesContainerContainerAddRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**containerContainerRemoteHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemotesContainerContainerAddRoleRequest struct via the builder pattern


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


## RemotesContainerContainerCreate

> ContainerContainerRemoteResponse RemotesContainerContainerCreate(ctx).ContainerContainerRemote(containerContainerRemote).Execute()

Create a container remote



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
    containerContainerRemote := *openapiclient.NewContainerContainerRemote("Name_example", "Url_example", "UpstreamName_example") // ContainerContainerRemote | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RemotesContainerAPI.RemotesContainerContainerCreate(context.Background()).ContainerContainerRemote(containerContainerRemote).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RemotesContainerAPI.RemotesContainerContainerCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RemotesContainerContainerCreate`: ContainerContainerRemoteResponse
    fmt.Fprintf(os.Stdout, "Response from `RemotesContainerAPI.RemotesContainerContainerCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRemotesContainerContainerCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **containerContainerRemote** | [**ContainerContainerRemote**](ContainerContainerRemote.md) |  | 

### Return type

[**ContainerContainerRemoteResponse**](ContainerContainerRemoteResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemotesContainerContainerDelete

> AsyncOperationResponse RemotesContainerContainerDelete(ctx, containerContainerRemoteHref).Execute()

Delete a container remote



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
    containerContainerRemoteHref := "containerContainerRemoteHref_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RemotesContainerAPI.RemotesContainerContainerDelete(context.Background(), containerContainerRemoteHref).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RemotesContainerAPI.RemotesContainerContainerDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RemotesContainerContainerDelete`: AsyncOperationResponse
    fmt.Fprintf(os.Stdout, "Response from `RemotesContainerAPI.RemotesContainerContainerDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**containerContainerRemoteHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemotesContainerContainerDeleteRequest struct via the builder pattern


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


## RemotesContainerContainerList

> PaginatedcontainerContainerRemoteResponseList RemotesContainerContainerList(ctx).Limit(limit).Name(name).NameContains(nameContains).NameIcontains(nameIcontains).NameIn(nameIn).NameStartswith(nameStartswith).Offset(offset).Ordering(ordering).PulpHrefIn(pulpHrefIn).PulpIdIn(pulpIdIn).PulpLabelSelect(pulpLabelSelect).PulpLastUpdated(pulpLastUpdated).PulpLastUpdatedGt(pulpLastUpdatedGt).PulpLastUpdatedGte(pulpLastUpdatedGte).PulpLastUpdatedLt(pulpLastUpdatedLt).PulpLastUpdatedLte(pulpLastUpdatedLte).PulpLastUpdatedRange(pulpLastUpdatedRange).Fields(fields).ExcludeFields(excludeFields).Execute()

List container remotes



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
    limit := int32(56) // int32 | Number of results to return per page. (optional)
    name := "name_example" // string | Filter results where name matches value (optional)
    nameContains := "nameContains_example" // string | Filter results where name contains value (optional)
    nameIcontains := "nameIcontains_example" // string | Filter results where name contains value (optional)
    nameIn := []string{"Inner_example"} // []string | Filter results where name is in a comma-separated list of values (optional)
    nameStartswith := "nameStartswith_example" // string | Filter results where name starts with value (optional)
    offset := int32(56) // int32 | The initial index from which to return the results. (optional)
    ordering := []string{"Ordering_example"} // []string | Ordering  * `pulp_id` - Pulp id * `-pulp_id` - Pulp id (descending) * `pulp_created` - Pulp created * `-pulp_created` - Pulp created (descending) * `pulp_last_updated` - Pulp last updated * `-pulp_last_updated` - Pulp last updated (descending) * `pulp_type` - Pulp type * `-pulp_type` - Pulp type (descending) * `name` - Name * `-name` - Name (descending) * `pulp_labels` - Pulp labels * `-pulp_labels` - Pulp labels (descending) * `url` - Url * `-url` - Url (descending) * `ca_cert` - Ca cert * `-ca_cert` - Ca cert (descending) * `client_cert` - Client cert * `-client_cert` - Client cert (descending) * `client_key` - Client key * `-client_key` - Client key (descending) * `tls_validation` - Tls validation * `-tls_validation` - Tls validation (descending) * `username` - Username * `-username` - Username (descending) * `password` - Password * `-password` - Password (descending) * `proxy_url` - Proxy url * `-proxy_url` - Proxy url (descending) * `proxy_username` - Proxy username * `-proxy_username` - Proxy username (descending) * `proxy_password` - Proxy password * `-proxy_password` - Proxy password (descending) * `download_concurrency` - Download concurrency * `-download_concurrency` - Download concurrency (descending) * `max_retries` - Max retries * `-max_retries` - Max retries (descending) * `policy` - Policy * `-policy` - Policy (descending) * `total_timeout` - Total timeout * `-total_timeout` - Total timeout (descending) * `connect_timeout` - Connect timeout * `-connect_timeout` - Connect timeout (descending) * `sock_connect_timeout` - Sock connect timeout * `-sock_connect_timeout` - Sock connect timeout (descending) * `sock_read_timeout` - Sock read timeout * `-sock_read_timeout` - Sock read timeout (descending) * `headers` - Headers * `-headers` - Headers (descending) * `rate_limit` - Rate limit * `-rate_limit` - Rate limit (descending) * `pk` - Pk * `-pk` - Pk (descending) (optional)
    pulpHrefIn := []string{"Inner_example"} // []string | Multiple values may be separated by commas. (optional)
    pulpIdIn := []string{"Inner_example"} // []string | Multiple values may be separated by commas. (optional)
    pulpLabelSelect := "pulpLabelSelect_example" // string | Filter labels by search string (optional)
    pulpLastUpdated := time.Now() // time.Time | Filter results where pulp_last_updated matches value (optional)
    pulpLastUpdatedGt := time.Now() // time.Time | Filter results where pulp_last_updated is greater than value (optional)
    pulpLastUpdatedGte := time.Now() // time.Time | Filter results where pulp_last_updated is greater than or equal to value (optional)
    pulpLastUpdatedLt := time.Now() // time.Time | Filter results where pulp_last_updated is less than value (optional)
    pulpLastUpdatedLte := time.Now() // time.Time | Filter results where pulp_last_updated is less than or equal to value (optional)
    pulpLastUpdatedRange := []time.Time{time.Now()} // []time.Time | Filter results where pulp_last_updated is between two comma separated values (optional)
    fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
    excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RemotesContainerAPI.RemotesContainerContainerList(context.Background()).Limit(limit).Name(name).NameContains(nameContains).NameIcontains(nameIcontains).NameIn(nameIn).NameStartswith(nameStartswith).Offset(offset).Ordering(ordering).PulpHrefIn(pulpHrefIn).PulpIdIn(pulpIdIn).PulpLabelSelect(pulpLabelSelect).PulpLastUpdated(pulpLastUpdated).PulpLastUpdatedGt(pulpLastUpdatedGt).PulpLastUpdatedGte(pulpLastUpdatedGte).PulpLastUpdatedLt(pulpLastUpdatedLt).PulpLastUpdatedLte(pulpLastUpdatedLte).PulpLastUpdatedRange(pulpLastUpdatedRange).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RemotesContainerAPI.RemotesContainerContainerList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RemotesContainerContainerList`: PaginatedcontainerContainerRemoteResponseList
    fmt.Fprintf(os.Stdout, "Response from `RemotesContainerAPI.RemotesContainerContainerList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRemotesContainerContainerListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Number of results to return per page. | 
 **name** | **string** | Filter results where name matches value | 
 **nameContains** | **string** | Filter results where name contains value | 
 **nameIcontains** | **string** | Filter results where name contains value | 
 **nameIn** | **[]string** | Filter results where name is in a comma-separated list of values | 
 **nameStartswith** | **string** | Filter results where name starts with value | 
 **offset** | **int32** | The initial index from which to return the results. | 
 **ordering** | **[]string** | Ordering  * &#x60;pulp_id&#x60; - Pulp id * &#x60;-pulp_id&#x60; - Pulp id (descending) * &#x60;pulp_created&#x60; - Pulp created * &#x60;-pulp_created&#x60; - Pulp created (descending) * &#x60;pulp_last_updated&#x60; - Pulp last updated * &#x60;-pulp_last_updated&#x60; - Pulp last updated (descending) * &#x60;pulp_type&#x60; - Pulp type * &#x60;-pulp_type&#x60; - Pulp type (descending) * &#x60;name&#x60; - Name * &#x60;-name&#x60; - Name (descending) * &#x60;pulp_labels&#x60; - Pulp labels * &#x60;-pulp_labels&#x60; - Pulp labels (descending) * &#x60;url&#x60; - Url * &#x60;-url&#x60; - Url (descending) * &#x60;ca_cert&#x60; - Ca cert * &#x60;-ca_cert&#x60; - Ca cert (descending) * &#x60;client_cert&#x60; - Client cert * &#x60;-client_cert&#x60; - Client cert (descending) * &#x60;client_key&#x60; - Client key * &#x60;-client_key&#x60; - Client key (descending) * &#x60;tls_validation&#x60; - Tls validation * &#x60;-tls_validation&#x60; - Tls validation (descending) * &#x60;username&#x60; - Username * &#x60;-username&#x60; - Username (descending) * &#x60;password&#x60; - Password * &#x60;-password&#x60; - Password (descending) * &#x60;proxy_url&#x60; - Proxy url * &#x60;-proxy_url&#x60; - Proxy url (descending) * &#x60;proxy_username&#x60; - Proxy username * &#x60;-proxy_username&#x60; - Proxy username (descending) * &#x60;proxy_password&#x60; - Proxy password * &#x60;-proxy_password&#x60; - Proxy password (descending) * &#x60;download_concurrency&#x60; - Download concurrency * &#x60;-download_concurrency&#x60; - Download concurrency (descending) * &#x60;max_retries&#x60; - Max retries * &#x60;-max_retries&#x60; - Max retries (descending) * &#x60;policy&#x60; - Policy * &#x60;-policy&#x60; - Policy (descending) * &#x60;total_timeout&#x60; - Total timeout * &#x60;-total_timeout&#x60; - Total timeout (descending) * &#x60;connect_timeout&#x60; - Connect timeout * &#x60;-connect_timeout&#x60; - Connect timeout (descending) * &#x60;sock_connect_timeout&#x60; - Sock connect timeout * &#x60;-sock_connect_timeout&#x60; - Sock connect timeout (descending) * &#x60;sock_read_timeout&#x60; - Sock read timeout * &#x60;-sock_read_timeout&#x60; - Sock read timeout (descending) * &#x60;headers&#x60; - Headers * &#x60;-headers&#x60; - Headers (descending) * &#x60;rate_limit&#x60; - Rate limit * &#x60;-rate_limit&#x60; - Rate limit (descending) * &#x60;pk&#x60; - Pk * &#x60;-pk&#x60; - Pk (descending) | 
 **pulpHrefIn** | **[]string** | Multiple values may be separated by commas. | 
 **pulpIdIn** | **[]string** | Multiple values may be separated by commas. | 
 **pulpLabelSelect** | **string** | Filter labels by search string | 
 **pulpLastUpdated** | **time.Time** | Filter results where pulp_last_updated matches value | 
 **pulpLastUpdatedGt** | **time.Time** | Filter results where pulp_last_updated is greater than value | 
 **pulpLastUpdatedGte** | **time.Time** | Filter results where pulp_last_updated is greater than or equal to value | 
 **pulpLastUpdatedLt** | **time.Time** | Filter results where pulp_last_updated is less than value | 
 **pulpLastUpdatedLte** | **time.Time** | Filter results where pulp_last_updated is less than or equal to value | 
 **pulpLastUpdatedRange** | [**[]time.Time**](time.Time.md) | Filter results where pulp_last_updated is between two comma separated values | 
 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

### Return type

[**PaginatedcontainerContainerRemoteResponseList**](PaginatedcontainerContainerRemoteResponseList.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemotesContainerContainerListRoles

> ObjectRolesResponse RemotesContainerContainerListRoles(ctx, containerContainerRemoteHref).Fields(fields).ExcludeFields(excludeFields).Execute()





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
    containerContainerRemoteHref := "containerContainerRemoteHref_example" // string | 
    fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
    excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RemotesContainerAPI.RemotesContainerContainerListRoles(context.Background(), containerContainerRemoteHref).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RemotesContainerAPI.RemotesContainerContainerListRoles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RemotesContainerContainerListRoles`: ObjectRolesResponse
    fmt.Fprintf(os.Stdout, "Response from `RemotesContainerAPI.RemotesContainerContainerListRoles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**containerContainerRemoteHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemotesContainerContainerListRolesRequest struct via the builder pattern


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


## RemotesContainerContainerMyPermissions

> MyPermissionsResponse RemotesContainerContainerMyPermissions(ctx, containerContainerRemoteHref).Fields(fields).ExcludeFields(excludeFields).Execute()





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
    containerContainerRemoteHref := "containerContainerRemoteHref_example" // string | 
    fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
    excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RemotesContainerAPI.RemotesContainerContainerMyPermissions(context.Background(), containerContainerRemoteHref).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RemotesContainerAPI.RemotesContainerContainerMyPermissions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RemotesContainerContainerMyPermissions`: MyPermissionsResponse
    fmt.Fprintf(os.Stdout, "Response from `RemotesContainerAPI.RemotesContainerContainerMyPermissions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**containerContainerRemoteHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemotesContainerContainerMyPermissionsRequest struct via the builder pattern


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


## RemotesContainerContainerPartialUpdate

> AsyncOperationResponse RemotesContainerContainerPartialUpdate(ctx, containerContainerRemoteHref).PatchedcontainerContainerRemote(patchedcontainerContainerRemote).Execute()

Update a container remote



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
    containerContainerRemoteHref := "containerContainerRemoteHref_example" // string | 
    patchedcontainerContainerRemote := *openapiclient.NewPatchedcontainerContainerRemote() // PatchedcontainerContainerRemote | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RemotesContainerAPI.RemotesContainerContainerPartialUpdate(context.Background(), containerContainerRemoteHref).PatchedcontainerContainerRemote(patchedcontainerContainerRemote).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RemotesContainerAPI.RemotesContainerContainerPartialUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RemotesContainerContainerPartialUpdate`: AsyncOperationResponse
    fmt.Fprintf(os.Stdout, "Response from `RemotesContainerAPI.RemotesContainerContainerPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**containerContainerRemoteHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemotesContainerContainerPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedcontainerContainerRemote** | [**PatchedcontainerContainerRemote**](PatchedcontainerContainerRemote.md) |  | 

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


## RemotesContainerContainerRead

> ContainerContainerRemoteResponse RemotesContainerContainerRead(ctx, containerContainerRemoteHref).Fields(fields).ExcludeFields(excludeFields).Execute()

Inspect a container remote



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
    containerContainerRemoteHref := "containerContainerRemoteHref_example" // string | 
    fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
    excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RemotesContainerAPI.RemotesContainerContainerRead(context.Background(), containerContainerRemoteHref).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RemotesContainerAPI.RemotesContainerContainerRead``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RemotesContainerContainerRead`: ContainerContainerRemoteResponse
    fmt.Fprintf(os.Stdout, "Response from `RemotesContainerAPI.RemotesContainerContainerRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**containerContainerRemoteHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemotesContainerContainerReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

### Return type

[**ContainerContainerRemoteResponse**](ContainerContainerRemoteResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemotesContainerContainerRemoveRole

> NestedRoleResponse RemotesContainerContainerRemoveRole(ctx, containerContainerRemoteHref).NestedRole(nestedRole).Execute()





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
    containerContainerRemoteHref := "containerContainerRemoteHref_example" // string | 
    nestedRole := *openapiclient.NewNestedRole("Role_example") // NestedRole | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RemotesContainerAPI.RemotesContainerContainerRemoveRole(context.Background(), containerContainerRemoteHref).NestedRole(nestedRole).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RemotesContainerAPI.RemotesContainerContainerRemoveRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RemotesContainerContainerRemoveRole`: NestedRoleResponse
    fmt.Fprintf(os.Stdout, "Response from `RemotesContainerAPI.RemotesContainerContainerRemoveRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**containerContainerRemoteHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemotesContainerContainerRemoveRoleRequest struct via the builder pattern


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


## RemotesContainerContainerUpdate

> AsyncOperationResponse RemotesContainerContainerUpdate(ctx, containerContainerRemoteHref).ContainerContainerRemote(containerContainerRemote).Execute()

Update a container remote



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
    containerContainerRemoteHref := "containerContainerRemoteHref_example" // string | 
    containerContainerRemote := *openapiclient.NewContainerContainerRemote("Name_example", "Url_example", "UpstreamName_example") // ContainerContainerRemote | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RemotesContainerAPI.RemotesContainerContainerUpdate(context.Background(), containerContainerRemoteHref).ContainerContainerRemote(containerContainerRemote).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RemotesContainerAPI.RemotesContainerContainerUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RemotesContainerContainerUpdate`: AsyncOperationResponse
    fmt.Fprintf(os.Stdout, "Response from `RemotesContainerAPI.RemotesContainerContainerUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**containerContainerRemoteHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemotesContainerContainerUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **containerContainerRemote** | [**ContainerContainerRemote**](ContainerContainerRemote.md) |  | 

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

