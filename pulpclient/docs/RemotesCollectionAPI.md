# \RemotesCollectionAPI

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RemotesAnsibleCollectionAddRole**](RemotesCollectionAPI.md#RemotesAnsibleCollectionAddRole) | **Post** {ansible_collection_remote_href}add_role/ | 
[**RemotesAnsibleCollectionCreate**](RemotesCollectionAPI.md#RemotesAnsibleCollectionCreate) | **Post** /pulp/api/v3/remotes/ansible/collection/ | Create a collection remote
[**RemotesAnsibleCollectionDelete**](RemotesCollectionAPI.md#RemotesAnsibleCollectionDelete) | **Delete** {ansible_collection_remote_href} | Delete a collection remote
[**RemotesAnsibleCollectionList**](RemotesCollectionAPI.md#RemotesAnsibleCollectionList) | **Get** /pulp/api/v3/remotes/ansible/collection/ | List collection remotes
[**RemotesAnsibleCollectionListRoles**](RemotesCollectionAPI.md#RemotesAnsibleCollectionListRoles) | **Get** {ansible_collection_remote_href}list_roles/ | 
[**RemotesAnsibleCollectionMyPermissions**](RemotesCollectionAPI.md#RemotesAnsibleCollectionMyPermissions) | **Get** {ansible_collection_remote_href}my_permissions/ | 
[**RemotesAnsibleCollectionPartialUpdate**](RemotesCollectionAPI.md#RemotesAnsibleCollectionPartialUpdate) | **Patch** {ansible_collection_remote_href} | Update a collection remote
[**RemotesAnsibleCollectionRead**](RemotesCollectionAPI.md#RemotesAnsibleCollectionRead) | **Get** {ansible_collection_remote_href} | Inspect a collection remote
[**RemotesAnsibleCollectionRemoveRole**](RemotesCollectionAPI.md#RemotesAnsibleCollectionRemoveRole) | **Post** {ansible_collection_remote_href}remove_role/ | 
[**RemotesAnsibleCollectionUpdate**](RemotesCollectionAPI.md#RemotesAnsibleCollectionUpdate) | **Put** {ansible_collection_remote_href} | Update a collection remote



## RemotesAnsibleCollectionAddRole

> NestedRoleResponse RemotesAnsibleCollectionAddRole(ctx, ansibleCollectionRemoteHref).NestedRole(nestedRole).Execute()





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
    ansibleCollectionRemoteHref := "ansibleCollectionRemoteHref_example" // string | 
    nestedRole := *openapiclient.NewNestedRole("Role_example") // NestedRole | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RemotesCollectionAPI.RemotesAnsibleCollectionAddRole(context.Background(), ansibleCollectionRemoteHref).NestedRole(nestedRole).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RemotesCollectionAPI.RemotesAnsibleCollectionAddRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RemotesAnsibleCollectionAddRole`: NestedRoleResponse
    fmt.Fprintf(os.Stdout, "Response from `RemotesCollectionAPI.RemotesAnsibleCollectionAddRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ansibleCollectionRemoteHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemotesAnsibleCollectionAddRoleRequest struct via the builder pattern


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


## RemotesAnsibleCollectionCreate

> AnsibleCollectionRemoteResponse RemotesAnsibleCollectionCreate(ctx).AnsibleCollectionRemote(ansibleCollectionRemote).Execute()

Create a collection remote



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
    ansibleCollectionRemote := *openapiclient.NewAnsibleCollectionRemote("Name_example", "Url_example") // AnsibleCollectionRemote | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RemotesCollectionAPI.RemotesAnsibleCollectionCreate(context.Background()).AnsibleCollectionRemote(ansibleCollectionRemote).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RemotesCollectionAPI.RemotesAnsibleCollectionCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RemotesAnsibleCollectionCreate`: AnsibleCollectionRemoteResponse
    fmt.Fprintf(os.Stdout, "Response from `RemotesCollectionAPI.RemotesAnsibleCollectionCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRemotesAnsibleCollectionCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ansibleCollectionRemote** | [**AnsibleCollectionRemote**](AnsibleCollectionRemote.md) |  | 

### Return type

[**AnsibleCollectionRemoteResponse**](AnsibleCollectionRemoteResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemotesAnsibleCollectionDelete

> AsyncOperationResponse RemotesAnsibleCollectionDelete(ctx, ansibleCollectionRemoteHref).Execute()

Delete a collection remote



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
    ansibleCollectionRemoteHref := "ansibleCollectionRemoteHref_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RemotesCollectionAPI.RemotesAnsibleCollectionDelete(context.Background(), ansibleCollectionRemoteHref).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RemotesCollectionAPI.RemotesAnsibleCollectionDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RemotesAnsibleCollectionDelete`: AsyncOperationResponse
    fmt.Fprintf(os.Stdout, "Response from `RemotesCollectionAPI.RemotesAnsibleCollectionDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ansibleCollectionRemoteHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemotesAnsibleCollectionDeleteRequest struct via the builder pattern


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


## RemotesAnsibleCollectionList

> PaginatedansibleCollectionRemoteResponseList RemotesAnsibleCollectionList(ctx).Limit(limit).Name(name).NameContains(nameContains).NameIcontains(nameIcontains).NameIn(nameIn).NameStartswith(nameStartswith).Offset(offset).Ordering(ordering).PulpHrefIn(pulpHrefIn).PulpIdIn(pulpIdIn).PulpLabelSelect(pulpLabelSelect).PulpLastUpdated(pulpLastUpdated).PulpLastUpdatedGt(pulpLastUpdatedGt).PulpLastUpdatedGte(pulpLastUpdatedGte).PulpLastUpdatedLt(pulpLastUpdatedLt).PulpLastUpdatedLte(pulpLastUpdatedLte).PulpLastUpdatedRange(pulpLastUpdatedRange).Url(url).UrlContains(urlContains).UrlIcontains(urlIcontains).UrlIn(urlIn).Fields(fields).ExcludeFields(excludeFields).Execute()

List collection remotes



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
    ordering := []string{"Ordering_example"} // []string | Ordering  * `pulp_id` - Pulp id * `-pulp_id` - Pulp id (descending) * `pulp_created` - Pulp created * `-pulp_created` - Pulp created (descending) * `pulp_last_updated` - Pulp last updated * `-pulp_last_updated` - Pulp last updated (descending) * `pulp_type` - Pulp type * `-pulp_type` - Pulp type (descending) * `name` - Name * `-name` - Name (descending) * `pulp_labels` - Pulp labels * `-pulp_labels` - Pulp labels (descending) * `url` - Url * `-url` - Url (descending) * `ca_cert` - Ca cert * `-ca_cert` - Ca cert (descending) * `client_cert` - Client cert * `-client_cert` - Client cert (descending) * `client_key` - Client key * `-client_key` - Client key (descending) * `tls_validation` - Tls validation * `-tls_validation` - Tls validation (descending) * `username` - Username * `-username` - Username (descending) * `password` - Password * `-password` - Password (descending) * `proxy_url` - Proxy url * `-proxy_url` - Proxy url (descending) * `proxy_username` - Proxy username * `-proxy_username` - Proxy username (descending) * `proxy_password` - Proxy password * `-proxy_password` - Proxy password (descending) * `download_concurrency` - Download concurrency * `-download_concurrency` - Download concurrency (descending) * `max_retries` - Max retries * `-max_retries` - Max retries (descending) * `policy` - Policy * `-policy` - Policy (descending) * `total_timeout` - Total timeout * `-total_timeout` - Total timeout (descending) * `connect_timeout` - Connect timeout * `-connect_timeout` - Connect timeout (descending) * `sock_connect_timeout` - Sock connect timeout * `-sock_connect_timeout` - Sock connect timeout (descending) * `sock_read_timeout` - Sock read timeout * `-sock_read_timeout` - Sock read timeout (descending) * `headers` - Headers * `-headers` - Headers (descending) * `rate_limit` - Rate limit * `-rate_limit` - Rate limit (descending) * `requirements_file` - Requirements file * `-requirements_file` - Requirements file (descending) * `auth_url` - Auth url * `-auth_url` - Auth url (descending) * `token` - Token * `-token` - Token (descending) * `sync_dependencies` - Sync dependencies * `-sync_dependencies` - Sync dependencies (descending) * `signed_only` - Signed only * `-signed_only` - Signed only (descending) * `pk` - Pk * `-pk` - Pk (descending) (optional)
    pulpHrefIn := []string{"Inner_example"} // []string | Multiple values may be separated by commas. (optional)
    pulpIdIn := []string{"Inner_example"} // []string | Multiple values may be separated by commas. (optional)
    pulpLabelSelect := "pulpLabelSelect_example" // string | Filter labels by search string (optional)
    pulpLastUpdated := time.Now() // time.Time | Filter results where pulp_last_updated matches value (optional)
    pulpLastUpdatedGt := time.Now() // time.Time | Filter results where pulp_last_updated is greater than value (optional)
    pulpLastUpdatedGte := time.Now() // time.Time | Filter results where pulp_last_updated is greater than or equal to value (optional)
    pulpLastUpdatedLt := time.Now() // time.Time | Filter results where pulp_last_updated is less than value (optional)
    pulpLastUpdatedLte := time.Now() // time.Time | Filter results where pulp_last_updated is less than or equal to value (optional)
    pulpLastUpdatedRange := []time.Time{time.Now()} // []time.Time | Filter results where pulp_last_updated is between two comma separated values (optional)
    url := "url_example" // string | Filter results where url matches value (optional)
    urlContains := "urlContains_example" // string | Filter results where url contains value (optional)
    urlIcontains := "urlIcontains_example" // string | Filter results where url contains value (optional)
    urlIn := []string{"Inner_example"} // []string | Filter results where url is in a comma-separated list of values (optional)
    fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
    excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RemotesCollectionAPI.RemotesAnsibleCollectionList(context.Background()).Limit(limit).Name(name).NameContains(nameContains).NameIcontains(nameIcontains).NameIn(nameIn).NameStartswith(nameStartswith).Offset(offset).Ordering(ordering).PulpHrefIn(pulpHrefIn).PulpIdIn(pulpIdIn).PulpLabelSelect(pulpLabelSelect).PulpLastUpdated(pulpLastUpdated).PulpLastUpdatedGt(pulpLastUpdatedGt).PulpLastUpdatedGte(pulpLastUpdatedGte).PulpLastUpdatedLt(pulpLastUpdatedLt).PulpLastUpdatedLte(pulpLastUpdatedLte).PulpLastUpdatedRange(pulpLastUpdatedRange).Url(url).UrlContains(urlContains).UrlIcontains(urlIcontains).UrlIn(urlIn).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RemotesCollectionAPI.RemotesAnsibleCollectionList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RemotesAnsibleCollectionList`: PaginatedansibleCollectionRemoteResponseList
    fmt.Fprintf(os.Stdout, "Response from `RemotesCollectionAPI.RemotesAnsibleCollectionList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRemotesAnsibleCollectionListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Number of results to return per page. | 
 **name** | **string** | Filter results where name matches value | 
 **nameContains** | **string** | Filter results where name contains value | 
 **nameIcontains** | **string** | Filter results where name contains value | 
 **nameIn** | **[]string** | Filter results where name is in a comma-separated list of values | 
 **nameStartswith** | **string** | Filter results where name starts with value | 
 **offset** | **int32** | The initial index from which to return the results. | 
 **ordering** | **[]string** | Ordering  * &#x60;pulp_id&#x60; - Pulp id * &#x60;-pulp_id&#x60; - Pulp id (descending) * &#x60;pulp_created&#x60; - Pulp created * &#x60;-pulp_created&#x60; - Pulp created (descending) * &#x60;pulp_last_updated&#x60; - Pulp last updated * &#x60;-pulp_last_updated&#x60; - Pulp last updated (descending) * &#x60;pulp_type&#x60; - Pulp type * &#x60;-pulp_type&#x60; - Pulp type (descending) * &#x60;name&#x60; - Name * &#x60;-name&#x60; - Name (descending) * &#x60;pulp_labels&#x60; - Pulp labels * &#x60;-pulp_labels&#x60; - Pulp labels (descending) * &#x60;url&#x60; - Url * &#x60;-url&#x60; - Url (descending) * &#x60;ca_cert&#x60; - Ca cert * &#x60;-ca_cert&#x60; - Ca cert (descending) * &#x60;client_cert&#x60; - Client cert * &#x60;-client_cert&#x60; - Client cert (descending) * &#x60;client_key&#x60; - Client key * &#x60;-client_key&#x60; - Client key (descending) * &#x60;tls_validation&#x60; - Tls validation * &#x60;-tls_validation&#x60; - Tls validation (descending) * &#x60;username&#x60; - Username * &#x60;-username&#x60; - Username (descending) * &#x60;password&#x60; - Password * &#x60;-password&#x60; - Password (descending) * &#x60;proxy_url&#x60; - Proxy url * &#x60;-proxy_url&#x60; - Proxy url (descending) * &#x60;proxy_username&#x60; - Proxy username * &#x60;-proxy_username&#x60; - Proxy username (descending) * &#x60;proxy_password&#x60; - Proxy password * &#x60;-proxy_password&#x60; - Proxy password (descending) * &#x60;download_concurrency&#x60; - Download concurrency * &#x60;-download_concurrency&#x60; - Download concurrency (descending) * &#x60;max_retries&#x60; - Max retries * &#x60;-max_retries&#x60; - Max retries (descending) * &#x60;policy&#x60; - Policy * &#x60;-policy&#x60; - Policy (descending) * &#x60;total_timeout&#x60; - Total timeout * &#x60;-total_timeout&#x60; - Total timeout (descending) * &#x60;connect_timeout&#x60; - Connect timeout * &#x60;-connect_timeout&#x60; - Connect timeout (descending) * &#x60;sock_connect_timeout&#x60; - Sock connect timeout * &#x60;-sock_connect_timeout&#x60; - Sock connect timeout (descending) * &#x60;sock_read_timeout&#x60; - Sock read timeout * &#x60;-sock_read_timeout&#x60; - Sock read timeout (descending) * &#x60;headers&#x60; - Headers * &#x60;-headers&#x60; - Headers (descending) * &#x60;rate_limit&#x60; - Rate limit * &#x60;-rate_limit&#x60; - Rate limit (descending) * &#x60;requirements_file&#x60; - Requirements file * &#x60;-requirements_file&#x60; - Requirements file (descending) * &#x60;auth_url&#x60; - Auth url * &#x60;-auth_url&#x60; - Auth url (descending) * &#x60;token&#x60; - Token * &#x60;-token&#x60; - Token (descending) * &#x60;sync_dependencies&#x60; - Sync dependencies * &#x60;-sync_dependencies&#x60; - Sync dependencies (descending) * &#x60;signed_only&#x60; - Signed only * &#x60;-signed_only&#x60; - Signed only (descending) * &#x60;pk&#x60; - Pk * &#x60;-pk&#x60; - Pk (descending) | 
 **pulpHrefIn** | **[]string** | Multiple values may be separated by commas. | 
 **pulpIdIn** | **[]string** | Multiple values may be separated by commas. | 
 **pulpLabelSelect** | **string** | Filter labels by search string | 
 **pulpLastUpdated** | **time.Time** | Filter results where pulp_last_updated matches value | 
 **pulpLastUpdatedGt** | **time.Time** | Filter results where pulp_last_updated is greater than value | 
 **pulpLastUpdatedGte** | **time.Time** | Filter results where pulp_last_updated is greater than or equal to value | 
 **pulpLastUpdatedLt** | **time.Time** | Filter results where pulp_last_updated is less than value | 
 **pulpLastUpdatedLte** | **time.Time** | Filter results where pulp_last_updated is less than or equal to value | 
 **pulpLastUpdatedRange** | [**[]time.Time**](time.Time.md) | Filter results where pulp_last_updated is between two comma separated values | 
 **url** | **string** | Filter results where url matches value | 
 **urlContains** | **string** | Filter results where url contains value | 
 **urlIcontains** | **string** | Filter results where url contains value | 
 **urlIn** | **[]string** | Filter results where url is in a comma-separated list of values | 
 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

### Return type

[**PaginatedansibleCollectionRemoteResponseList**](PaginatedansibleCollectionRemoteResponseList.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemotesAnsibleCollectionListRoles

> ObjectRolesResponse RemotesAnsibleCollectionListRoles(ctx, ansibleCollectionRemoteHref).Fields(fields).ExcludeFields(excludeFields).Execute()





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
    ansibleCollectionRemoteHref := "ansibleCollectionRemoteHref_example" // string | 
    fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
    excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RemotesCollectionAPI.RemotesAnsibleCollectionListRoles(context.Background(), ansibleCollectionRemoteHref).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RemotesCollectionAPI.RemotesAnsibleCollectionListRoles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RemotesAnsibleCollectionListRoles`: ObjectRolesResponse
    fmt.Fprintf(os.Stdout, "Response from `RemotesCollectionAPI.RemotesAnsibleCollectionListRoles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ansibleCollectionRemoteHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemotesAnsibleCollectionListRolesRequest struct via the builder pattern


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


## RemotesAnsibleCollectionMyPermissions

> MyPermissionsResponse RemotesAnsibleCollectionMyPermissions(ctx, ansibleCollectionRemoteHref).Fields(fields).ExcludeFields(excludeFields).Execute()





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
    ansibleCollectionRemoteHref := "ansibleCollectionRemoteHref_example" // string | 
    fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
    excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RemotesCollectionAPI.RemotesAnsibleCollectionMyPermissions(context.Background(), ansibleCollectionRemoteHref).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RemotesCollectionAPI.RemotesAnsibleCollectionMyPermissions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RemotesAnsibleCollectionMyPermissions`: MyPermissionsResponse
    fmt.Fprintf(os.Stdout, "Response from `RemotesCollectionAPI.RemotesAnsibleCollectionMyPermissions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ansibleCollectionRemoteHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemotesAnsibleCollectionMyPermissionsRequest struct via the builder pattern


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


## RemotesAnsibleCollectionPartialUpdate

> AsyncOperationResponse RemotesAnsibleCollectionPartialUpdate(ctx, ansibleCollectionRemoteHref).PatchedansibleCollectionRemote(patchedansibleCollectionRemote).Execute()

Update a collection remote



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
    ansibleCollectionRemoteHref := "ansibleCollectionRemoteHref_example" // string | 
    patchedansibleCollectionRemote := *openapiclient.NewPatchedansibleCollectionRemote() // PatchedansibleCollectionRemote | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RemotesCollectionAPI.RemotesAnsibleCollectionPartialUpdate(context.Background(), ansibleCollectionRemoteHref).PatchedansibleCollectionRemote(patchedansibleCollectionRemote).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RemotesCollectionAPI.RemotesAnsibleCollectionPartialUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RemotesAnsibleCollectionPartialUpdate`: AsyncOperationResponse
    fmt.Fprintf(os.Stdout, "Response from `RemotesCollectionAPI.RemotesAnsibleCollectionPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ansibleCollectionRemoteHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemotesAnsibleCollectionPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedansibleCollectionRemote** | [**PatchedansibleCollectionRemote**](PatchedansibleCollectionRemote.md) |  | 

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


## RemotesAnsibleCollectionRead

> AnsibleCollectionRemoteResponse RemotesAnsibleCollectionRead(ctx, ansibleCollectionRemoteHref).Fields(fields).ExcludeFields(excludeFields).Execute()

Inspect a collection remote



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
    ansibleCollectionRemoteHref := "ansibleCollectionRemoteHref_example" // string | 
    fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
    excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RemotesCollectionAPI.RemotesAnsibleCollectionRead(context.Background(), ansibleCollectionRemoteHref).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RemotesCollectionAPI.RemotesAnsibleCollectionRead``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RemotesAnsibleCollectionRead`: AnsibleCollectionRemoteResponse
    fmt.Fprintf(os.Stdout, "Response from `RemotesCollectionAPI.RemotesAnsibleCollectionRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ansibleCollectionRemoteHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemotesAnsibleCollectionReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

### Return type

[**AnsibleCollectionRemoteResponse**](AnsibleCollectionRemoteResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemotesAnsibleCollectionRemoveRole

> NestedRoleResponse RemotesAnsibleCollectionRemoveRole(ctx, ansibleCollectionRemoteHref).NestedRole(nestedRole).Execute()





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
    ansibleCollectionRemoteHref := "ansibleCollectionRemoteHref_example" // string | 
    nestedRole := *openapiclient.NewNestedRole("Role_example") // NestedRole | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RemotesCollectionAPI.RemotesAnsibleCollectionRemoveRole(context.Background(), ansibleCollectionRemoteHref).NestedRole(nestedRole).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RemotesCollectionAPI.RemotesAnsibleCollectionRemoveRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RemotesAnsibleCollectionRemoveRole`: NestedRoleResponse
    fmt.Fprintf(os.Stdout, "Response from `RemotesCollectionAPI.RemotesAnsibleCollectionRemoveRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ansibleCollectionRemoteHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemotesAnsibleCollectionRemoveRoleRequest struct via the builder pattern


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


## RemotesAnsibleCollectionUpdate

> AsyncOperationResponse RemotesAnsibleCollectionUpdate(ctx, ansibleCollectionRemoteHref).AnsibleCollectionRemote(ansibleCollectionRemote).Execute()

Update a collection remote



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
    ansibleCollectionRemoteHref := "ansibleCollectionRemoteHref_example" // string | 
    ansibleCollectionRemote := *openapiclient.NewAnsibleCollectionRemote("Name_example", "Url_example") // AnsibleCollectionRemote | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RemotesCollectionAPI.RemotesAnsibleCollectionUpdate(context.Background(), ansibleCollectionRemoteHref).AnsibleCollectionRemote(ansibleCollectionRemote).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RemotesCollectionAPI.RemotesAnsibleCollectionUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RemotesAnsibleCollectionUpdate`: AsyncOperationResponse
    fmt.Fprintf(os.Stdout, "Response from `RemotesCollectionAPI.RemotesAnsibleCollectionUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ansibleCollectionRemoteHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemotesAnsibleCollectionUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ansibleCollectionRemote** | [**AnsibleCollectionRemote**](AnsibleCollectionRemote.md) |  | 

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

