# \AcsFileAPI

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AcsFileFileAddRole**](AcsFileAPI.md#AcsFileFileAddRole) | **Post** {file_file_alternate_content_source_href}add_role/ | 
[**AcsFileFileCreate**](AcsFileAPI.md#AcsFileFileCreate) | **Post** /pulp/api/v3/acs/file/file/ | Create a file alternate content source
[**AcsFileFileDelete**](AcsFileAPI.md#AcsFileFileDelete) | **Delete** {file_file_alternate_content_source_href} | Delete a file alternate content source
[**AcsFileFileList**](AcsFileAPI.md#AcsFileFileList) | **Get** /pulp/api/v3/acs/file/file/ | List file alternate content sources
[**AcsFileFileListRoles**](AcsFileAPI.md#AcsFileFileListRoles) | **Get** {file_file_alternate_content_source_href}list_roles/ | 
[**AcsFileFileMyPermissions**](AcsFileAPI.md#AcsFileFileMyPermissions) | **Get** {file_file_alternate_content_source_href}my_permissions/ | 
[**AcsFileFilePartialUpdate**](AcsFileAPI.md#AcsFileFilePartialUpdate) | **Patch** {file_file_alternate_content_source_href} | Update a file alternate content source
[**AcsFileFileRead**](AcsFileAPI.md#AcsFileFileRead) | **Get** {file_file_alternate_content_source_href} | Inspect a file alternate content source
[**AcsFileFileRefresh**](AcsFileAPI.md#AcsFileFileRefresh) | **Post** {file_file_alternate_content_source_href}refresh/ | Refresh metadata
[**AcsFileFileRemoveRole**](AcsFileAPI.md#AcsFileFileRemoveRole) | **Post** {file_file_alternate_content_source_href}remove_role/ | 
[**AcsFileFileUpdate**](AcsFileAPI.md#AcsFileFileUpdate) | **Put** {file_file_alternate_content_source_href} | Update a file alternate content source



## AcsFileFileAddRole

> NestedRoleResponse AcsFileFileAddRole(ctx, fileFileAlternateContentSourceHref).NestedRole(nestedRole).Execute()





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
    fileFileAlternateContentSourceHref := "fileFileAlternateContentSourceHref_example" // string | 
    nestedRole := *openapiclient.NewNestedRole("Role_example") // NestedRole | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AcsFileAPI.AcsFileFileAddRole(context.Background(), fileFileAlternateContentSourceHref).NestedRole(nestedRole).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AcsFileAPI.AcsFileFileAddRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AcsFileFileAddRole`: NestedRoleResponse
    fmt.Fprintf(os.Stdout, "Response from `AcsFileAPI.AcsFileFileAddRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fileFileAlternateContentSourceHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAcsFileFileAddRoleRequest struct via the builder pattern


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


## AcsFileFileCreate

> FileFileAlternateContentSourceResponse AcsFileFileCreate(ctx).FileFileAlternateContentSource(fileFileAlternateContentSource).Execute()

Create a file alternate content source



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
    fileFileAlternateContentSource := *openapiclient.NewFileFileAlternateContentSource("Name_example", "Remote_example") // FileFileAlternateContentSource | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AcsFileAPI.AcsFileFileCreate(context.Background()).FileFileAlternateContentSource(fileFileAlternateContentSource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AcsFileAPI.AcsFileFileCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AcsFileFileCreate`: FileFileAlternateContentSourceResponse
    fmt.Fprintf(os.Stdout, "Response from `AcsFileAPI.AcsFileFileCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAcsFileFileCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fileFileAlternateContentSource** | [**FileFileAlternateContentSource**](FileFileAlternateContentSource.md) |  | 

### Return type

[**FileFileAlternateContentSourceResponse**](FileFileAlternateContentSourceResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AcsFileFileDelete

> AsyncOperationResponse AcsFileFileDelete(ctx, fileFileAlternateContentSourceHref).Execute()

Delete a file alternate content source



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
    fileFileAlternateContentSourceHref := "fileFileAlternateContentSourceHref_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AcsFileAPI.AcsFileFileDelete(context.Background(), fileFileAlternateContentSourceHref).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AcsFileAPI.AcsFileFileDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AcsFileFileDelete`: AsyncOperationResponse
    fmt.Fprintf(os.Stdout, "Response from `AcsFileAPI.AcsFileFileDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fileFileAlternateContentSourceHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAcsFileFileDeleteRequest struct via the builder pattern


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


## AcsFileFileList

> PaginatedfileFileAlternateContentSourceResponseList AcsFileFileList(ctx).Limit(limit).Name(name).NameContains(nameContains).NameIcontains(nameIcontains).NameIn(nameIn).NameStartswith(nameStartswith).Offset(offset).Ordering(ordering).PulpHrefIn(pulpHrefIn).PulpIdIn(pulpIdIn).Fields(fields).ExcludeFields(excludeFields).Execute()

List file alternate content sources



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
    resp, r, err := apiClient.AcsFileAPI.AcsFileFileList(context.Background()).Limit(limit).Name(name).NameContains(nameContains).NameIcontains(nameIcontains).NameIn(nameIn).NameStartswith(nameStartswith).Offset(offset).Ordering(ordering).PulpHrefIn(pulpHrefIn).PulpIdIn(pulpIdIn).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AcsFileAPI.AcsFileFileList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AcsFileFileList`: PaginatedfileFileAlternateContentSourceResponseList
    fmt.Fprintf(os.Stdout, "Response from `AcsFileAPI.AcsFileFileList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAcsFileFileListRequest struct via the builder pattern


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

[**PaginatedfileFileAlternateContentSourceResponseList**](PaginatedfileFileAlternateContentSourceResponseList.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AcsFileFileListRoles

> ObjectRolesResponse AcsFileFileListRoles(ctx, fileFileAlternateContentSourceHref).Fields(fields).ExcludeFields(excludeFields).Execute()





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
    fileFileAlternateContentSourceHref := "fileFileAlternateContentSourceHref_example" // string | 
    fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
    excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AcsFileAPI.AcsFileFileListRoles(context.Background(), fileFileAlternateContentSourceHref).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AcsFileAPI.AcsFileFileListRoles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AcsFileFileListRoles`: ObjectRolesResponse
    fmt.Fprintf(os.Stdout, "Response from `AcsFileAPI.AcsFileFileListRoles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fileFileAlternateContentSourceHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAcsFileFileListRolesRequest struct via the builder pattern


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


## AcsFileFileMyPermissions

> MyPermissionsResponse AcsFileFileMyPermissions(ctx, fileFileAlternateContentSourceHref).Fields(fields).ExcludeFields(excludeFields).Execute()





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
    fileFileAlternateContentSourceHref := "fileFileAlternateContentSourceHref_example" // string | 
    fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
    excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AcsFileAPI.AcsFileFileMyPermissions(context.Background(), fileFileAlternateContentSourceHref).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AcsFileAPI.AcsFileFileMyPermissions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AcsFileFileMyPermissions`: MyPermissionsResponse
    fmt.Fprintf(os.Stdout, "Response from `AcsFileAPI.AcsFileFileMyPermissions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fileFileAlternateContentSourceHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAcsFileFileMyPermissionsRequest struct via the builder pattern


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


## AcsFileFilePartialUpdate

> AsyncOperationResponse AcsFileFilePartialUpdate(ctx, fileFileAlternateContentSourceHref).PatchedfileFileAlternateContentSource(patchedfileFileAlternateContentSource).Execute()

Update a file alternate content source



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
    fileFileAlternateContentSourceHref := "fileFileAlternateContentSourceHref_example" // string | 
    patchedfileFileAlternateContentSource := *openapiclient.NewPatchedfileFileAlternateContentSource() // PatchedfileFileAlternateContentSource | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AcsFileAPI.AcsFileFilePartialUpdate(context.Background(), fileFileAlternateContentSourceHref).PatchedfileFileAlternateContentSource(patchedfileFileAlternateContentSource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AcsFileAPI.AcsFileFilePartialUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AcsFileFilePartialUpdate`: AsyncOperationResponse
    fmt.Fprintf(os.Stdout, "Response from `AcsFileAPI.AcsFileFilePartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fileFileAlternateContentSourceHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAcsFileFilePartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedfileFileAlternateContentSource** | [**PatchedfileFileAlternateContentSource**](PatchedfileFileAlternateContentSource.md) |  | 

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


## AcsFileFileRead

> FileFileAlternateContentSourceResponse AcsFileFileRead(ctx, fileFileAlternateContentSourceHref).Fields(fields).ExcludeFields(excludeFields).Execute()

Inspect a file alternate content source



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
    fileFileAlternateContentSourceHref := "fileFileAlternateContentSourceHref_example" // string | 
    fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
    excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AcsFileAPI.AcsFileFileRead(context.Background(), fileFileAlternateContentSourceHref).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AcsFileAPI.AcsFileFileRead``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AcsFileFileRead`: FileFileAlternateContentSourceResponse
    fmt.Fprintf(os.Stdout, "Response from `AcsFileAPI.AcsFileFileRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fileFileAlternateContentSourceHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAcsFileFileReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

### Return type

[**FileFileAlternateContentSourceResponse**](FileFileAlternateContentSourceResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AcsFileFileRefresh

> TaskGroupOperationResponse AcsFileFileRefresh(ctx, fileFileAlternateContentSourceHref).Execute()

Refresh metadata



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
    fileFileAlternateContentSourceHref := "fileFileAlternateContentSourceHref_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AcsFileAPI.AcsFileFileRefresh(context.Background(), fileFileAlternateContentSourceHref).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AcsFileAPI.AcsFileFileRefresh``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AcsFileFileRefresh`: TaskGroupOperationResponse
    fmt.Fprintf(os.Stdout, "Response from `AcsFileAPI.AcsFileFileRefresh`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fileFileAlternateContentSourceHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAcsFileFileRefreshRequest struct via the builder pattern


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


## AcsFileFileRemoveRole

> NestedRoleResponse AcsFileFileRemoveRole(ctx, fileFileAlternateContentSourceHref).NestedRole(nestedRole).Execute()





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
    fileFileAlternateContentSourceHref := "fileFileAlternateContentSourceHref_example" // string | 
    nestedRole := *openapiclient.NewNestedRole("Role_example") // NestedRole | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AcsFileAPI.AcsFileFileRemoveRole(context.Background(), fileFileAlternateContentSourceHref).NestedRole(nestedRole).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AcsFileAPI.AcsFileFileRemoveRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AcsFileFileRemoveRole`: NestedRoleResponse
    fmt.Fprintf(os.Stdout, "Response from `AcsFileAPI.AcsFileFileRemoveRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fileFileAlternateContentSourceHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAcsFileFileRemoveRoleRequest struct via the builder pattern


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


## AcsFileFileUpdate

> AsyncOperationResponse AcsFileFileUpdate(ctx, fileFileAlternateContentSourceHref).FileFileAlternateContentSource(fileFileAlternateContentSource).Execute()

Update a file alternate content source



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
    fileFileAlternateContentSourceHref := "fileFileAlternateContentSourceHref_example" // string | 
    fileFileAlternateContentSource := *openapiclient.NewFileFileAlternateContentSource("Name_example", "Remote_example") // FileFileAlternateContentSource | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AcsFileAPI.AcsFileFileUpdate(context.Background(), fileFileAlternateContentSourceHref).FileFileAlternateContentSource(fileFileAlternateContentSource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AcsFileAPI.AcsFileFileUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AcsFileFileUpdate`: AsyncOperationResponse
    fmt.Fprintf(os.Stdout, "Response from `AcsFileAPI.AcsFileFileUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fileFileAlternateContentSourceHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAcsFileFileUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fileFileAlternateContentSource** | [**FileFileAlternateContentSource**](FileFileAlternateContentSource.md) |  | 

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

