# \RemotesGitAPI

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RemotesAnsibleGitAddRole**](RemotesGitAPI.md#RemotesAnsibleGitAddRole) | **Post** {ansible_git_remote_href}add_role/ | 
[**RemotesAnsibleGitCreate**](RemotesGitAPI.md#RemotesAnsibleGitCreate) | **Post** /pulp/api/v3/remotes/ansible/git/ | Create a git remote
[**RemotesAnsibleGitDelete**](RemotesGitAPI.md#RemotesAnsibleGitDelete) | **Delete** {ansible_git_remote_href} | Delete a git remote
[**RemotesAnsibleGitList**](RemotesGitAPI.md#RemotesAnsibleGitList) | **Get** /pulp/api/v3/remotes/ansible/git/ | List git remotes
[**RemotesAnsibleGitListRoles**](RemotesGitAPI.md#RemotesAnsibleGitListRoles) | **Get** {ansible_git_remote_href}list_roles/ | 
[**RemotesAnsibleGitMyPermissions**](RemotesGitAPI.md#RemotesAnsibleGitMyPermissions) | **Get** {ansible_git_remote_href}my_permissions/ | 
[**RemotesAnsibleGitPartialUpdate**](RemotesGitAPI.md#RemotesAnsibleGitPartialUpdate) | **Patch** {ansible_git_remote_href} | Update a git remote
[**RemotesAnsibleGitRead**](RemotesGitAPI.md#RemotesAnsibleGitRead) | **Get** {ansible_git_remote_href} | Inspect a git remote
[**RemotesAnsibleGitRemoveRole**](RemotesGitAPI.md#RemotesAnsibleGitRemoveRole) | **Post** {ansible_git_remote_href}remove_role/ | 
[**RemotesAnsibleGitUpdate**](RemotesGitAPI.md#RemotesAnsibleGitUpdate) | **Put** {ansible_git_remote_href} | Update a git remote



## RemotesAnsibleGitAddRole

> NestedRoleResponse RemotesAnsibleGitAddRole(ctx, ansibleGitRemoteHref).NestedRole(nestedRole).Execute()





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
    ansibleGitRemoteHref := "ansibleGitRemoteHref_example" // string | 
    nestedRole := *openapiclient.NewNestedRole("Role_example") // NestedRole | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RemotesGitAPI.RemotesAnsibleGitAddRole(context.Background(), ansibleGitRemoteHref).NestedRole(nestedRole).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RemotesGitAPI.RemotesAnsibleGitAddRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RemotesAnsibleGitAddRole`: NestedRoleResponse
    fmt.Fprintf(os.Stdout, "Response from `RemotesGitAPI.RemotesAnsibleGitAddRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ansibleGitRemoteHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemotesAnsibleGitAddRoleRequest struct via the builder pattern


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


## RemotesAnsibleGitCreate

> AnsibleGitRemoteResponse RemotesAnsibleGitCreate(ctx).AnsibleGitRemote(ansibleGitRemote).Execute()

Create a git remote



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
    ansibleGitRemote := *openapiclient.NewAnsibleGitRemote("Url_example", "Name_example") // AnsibleGitRemote | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RemotesGitAPI.RemotesAnsibleGitCreate(context.Background()).AnsibleGitRemote(ansibleGitRemote).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RemotesGitAPI.RemotesAnsibleGitCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RemotesAnsibleGitCreate`: AnsibleGitRemoteResponse
    fmt.Fprintf(os.Stdout, "Response from `RemotesGitAPI.RemotesAnsibleGitCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRemotesAnsibleGitCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ansibleGitRemote** | [**AnsibleGitRemote**](AnsibleGitRemote.md) |  | 

### Return type

[**AnsibleGitRemoteResponse**](AnsibleGitRemoteResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemotesAnsibleGitDelete

> AsyncOperationResponse RemotesAnsibleGitDelete(ctx, ansibleGitRemoteHref).Execute()

Delete a git remote



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
    ansibleGitRemoteHref := "ansibleGitRemoteHref_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RemotesGitAPI.RemotesAnsibleGitDelete(context.Background(), ansibleGitRemoteHref).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RemotesGitAPI.RemotesAnsibleGitDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RemotesAnsibleGitDelete`: AsyncOperationResponse
    fmt.Fprintf(os.Stdout, "Response from `RemotesGitAPI.RemotesAnsibleGitDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ansibleGitRemoteHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemotesAnsibleGitDeleteRequest struct via the builder pattern


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


## RemotesAnsibleGitList

> PaginatedansibleGitRemoteResponseList RemotesAnsibleGitList(ctx).Limit(limit).Name(name).NameContains(nameContains).NameIcontains(nameIcontains).NameIn(nameIn).NameStartswith(nameStartswith).Offset(offset).Ordering(ordering).PulpHrefIn(pulpHrefIn).PulpIdIn(pulpIdIn).PulpLabelSelect(pulpLabelSelect).PulpLastUpdated(pulpLastUpdated).PulpLastUpdatedGt(pulpLastUpdatedGt).PulpLastUpdatedGte(pulpLastUpdatedGte).PulpLastUpdatedLt(pulpLastUpdatedLt).PulpLastUpdatedLte(pulpLastUpdatedLte).PulpLastUpdatedRange(pulpLastUpdatedRange).Fields(fields).ExcludeFields(excludeFields).Execute()

List git remotes



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
    resp, r, err := apiClient.RemotesGitAPI.RemotesAnsibleGitList(context.Background()).Limit(limit).Name(name).NameContains(nameContains).NameIcontains(nameIcontains).NameIn(nameIn).NameStartswith(nameStartswith).Offset(offset).Ordering(ordering).PulpHrefIn(pulpHrefIn).PulpIdIn(pulpIdIn).PulpLabelSelect(pulpLabelSelect).PulpLastUpdated(pulpLastUpdated).PulpLastUpdatedGt(pulpLastUpdatedGt).PulpLastUpdatedGte(pulpLastUpdatedGte).PulpLastUpdatedLt(pulpLastUpdatedLt).PulpLastUpdatedLte(pulpLastUpdatedLte).PulpLastUpdatedRange(pulpLastUpdatedRange).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RemotesGitAPI.RemotesAnsibleGitList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RemotesAnsibleGitList`: PaginatedansibleGitRemoteResponseList
    fmt.Fprintf(os.Stdout, "Response from `RemotesGitAPI.RemotesAnsibleGitList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRemotesAnsibleGitListRequest struct via the builder pattern


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

[**PaginatedansibleGitRemoteResponseList**](PaginatedansibleGitRemoteResponseList.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemotesAnsibleGitListRoles

> ObjectRolesResponse RemotesAnsibleGitListRoles(ctx, ansibleGitRemoteHref).Fields(fields).ExcludeFields(excludeFields).Execute()





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
    ansibleGitRemoteHref := "ansibleGitRemoteHref_example" // string | 
    fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
    excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RemotesGitAPI.RemotesAnsibleGitListRoles(context.Background(), ansibleGitRemoteHref).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RemotesGitAPI.RemotesAnsibleGitListRoles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RemotesAnsibleGitListRoles`: ObjectRolesResponse
    fmt.Fprintf(os.Stdout, "Response from `RemotesGitAPI.RemotesAnsibleGitListRoles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ansibleGitRemoteHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemotesAnsibleGitListRolesRequest struct via the builder pattern


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


## RemotesAnsibleGitMyPermissions

> MyPermissionsResponse RemotesAnsibleGitMyPermissions(ctx, ansibleGitRemoteHref).Fields(fields).ExcludeFields(excludeFields).Execute()





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
    ansibleGitRemoteHref := "ansibleGitRemoteHref_example" // string | 
    fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
    excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RemotesGitAPI.RemotesAnsibleGitMyPermissions(context.Background(), ansibleGitRemoteHref).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RemotesGitAPI.RemotesAnsibleGitMyPermissions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RemotesAnsibleGitMyPermissions`: MyPermissionsResponse
    fmt.Fprintf(os.Stdout, "Response from `RemotesGitAPI.RemotesAnsibleGitMyPermissions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ansibleGitRemoteHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemotesAnsibleGitMyPermissionsRequest struct via the builder pattern


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


## RemotesAnsibleGitPartialUpdate

> AsyncOperationResponse RemotesAnsibleGitPartialUpdate(ctx, ansibleGitRemoteHref).PatchedansibleGitRemote(patchedansibleGitRemote).Execute()

Update a git remote



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
    ansibleGitRemoteHref := "ansibleGitRemoteHref_example" // string | 
    patchedansibleGitRemote := *openapiclient.NewPatchedansibleGitRemote() // PatchedansibleGitRemote | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RemotesGitAPI.RemotesAnsibleGitPartialUpdate(context.Background(), ansibleGitRemoteHref).PatchedansibleGitRemote(patchedansibleGitRemote).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RemotesGitAPI.RemotesAnsibleGitPartialUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RemotesAnsibleGitPartialUpdate`: AsyncOperationResponse
    fmt.Fprintf(os.Stdout, "Response from `RemotesGitAPI.RemotesAnsibleGitPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ansibleGitRemoteHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemotesAnsibleGitPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedansibleGitRemote** | [**PatchedansibleGitRemote**](PatchedansibleGitRemote.md) |  | 

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


## RemotesAnsibleGitRead

> AnsibleGitRemoteResponse RemotesAnsibleGitRead(ctx, ansibleGitRemoteHref).Fields(fields).ExcludeFields(excludeFields).Execute()

Inspect a git remote



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
    ansibleGitRemoteHref := "ansibleGitRemoteHref_example" // string | 
    fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
    excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RemotesGitAPI.RemotesAnsibleGitRead(context.Background(), ansibleGitRemoteHref).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RemotesGitAPI.RemotesAnsibleGitRead``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RemotesAnsibleGitRead`: AnsibleGitRemoteResponse
    fmt.Fprintf(os.Stdout, "Response from `RemotesGitAPI.RemotesAnsibleGitRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ansibleGitRemoteHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemotesAnsibleGitReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

### Return type

[**AnsibleGitRemoteResponse**](AnsibleGitRemoteResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemotesAnsibleGitRemoveRole

> NestedRoleResponse RemotesAnsibleGitRemoveRole(ctx, ansibleGitRemoteHref).NestedRole(nestedRole).Execute()





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
    ansibleGitRemoteHref := "ansibleGitRemoteHref_example" // string | 
    nestedRole := *openapiclient.NewNestedRole("Role_example") // NestedRole | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RemotesGitAPI.RemotesAnsibleGitRemoveRole(context.Background(), ansibleGitRemoteHref).NestedRole(nestedRole).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RemotesGitAPI.RemotesAnsibleGitRemoveRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RemotesAnsibleGitRemoveRole`: NestedRoleResponse
    fmt.Fprintf(os.Stdout, "Response from `RemotesGitAPI.RemotesAnsibleGitRemoveRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ansibleGitRemoteHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemotesAnsibleGitRemoveRoleRequest struct via the builder pattern


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


## RemotesAnsibleGitUpdate

> AsyncOperationResponse RemotesAnsibleGitUpdate(ctx, ansibleGitRemoteHref).AnsibleGitRemote(ansibleGitRemote).Execute()

Update a git remote



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
    ansibleGitRemoteHref := "ansibleGitRemoteHref_example" // string | 
    ansibleGitRemote := *openapiclient.NewAnsibleGitRemote("Url_example", "Name_example") // AnsibleGitRemote | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RemotesGitAPI.RemotesAnsibleGitUpdate(context.Background(), ansibleGitRemoteHref).AnsibleGitRemote(ansibleGitRemote).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RemotesGitAPI.RemotesAnsibleGitUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RemotesAnsibleGitUpdate`: AsyncOperationResponse
    fmt.Fprintf(os.Stdout, "Response from `RemotesGitAPI.RemotesAnsibleGitUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ansibleGitRemoteHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemotesAnsibleGitUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ansibleGitRemote** | [**AnsibleGitRemote**](AnsibleGitRemote.md) |  | 

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

