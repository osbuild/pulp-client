# \AnsibleCollectionsAPI

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AnsibleCollectionsAddRole**](AnsibleCollectionsAPI.md#AnsibleCollectionsAddRole) | **Post** {ansible_collection_href}add_role/ | 
[**AnsibleCollectionsList**](AnsibleCollectionsAPI.md#AnsibleCollectionsList) | **Get** /pulp/api/v3/ansible/collections/ | List collections
[**AnsibleCollectionsListRoles**](AnsibleCollectionsAPI.md#AnsibleCollectionsListRoles) | **Get** {ansible_collection_href}list_roles/ | 
[**AnsibleCollectionsMyPermissions**](AnsibleCollectionsAPI.md#AnsibleCollectionsMyPermissions) | **Get** {ansible_collection_href}my_permissions/ | 
[**AnsibleCollectionsRemoveRole**](AnsibleCollectionsAPI.md#AnsibleCollectionsRemoveRole) | **Post** {ansible_collection_href}remove_role/ | 
[**UploadCollection**](AnsibleCollectionsAPI.md#UploadCollection) | **Post** /ansible/collections/ | Upload a collection



## AnsibleCollectionsAddRole

> NestedRoleResponse AnsibleCollectionsAddRole(ctx, ansibleCollectionHref).NestedRole(nestedRole).Execute()





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
    ansibleCollectionHref := "ansibleCollectionHref_example" // string | 
    nestedRole := *openapiclient.NewNestedRole("Role_example") // NestedRole | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AnsibleCollectionsAPI.AnsibleCollectionsAddRole(context.Background(), ansibleCollectionHref).NestedRole(nestedRole).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AnsibleCollectionsAPI.AnsibleCollectionsAddRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AnsibleCollectionsAddRole`: NestedRoleResponse
    fmt.Fprintf(os.Stdout, "Response from `AnsibleCollectionsAPI.AnsibleCollectionsAddRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ansibleCollectionHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAnsibleCollectionsAddRoleRequest struct via the builder pattern


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


## AnsibleCollectionsList

> PaginatedansibleCollectionResponseList AnsibleCollectionsList(ctx).Limit(limit).Name(name).Namespace(namespace).Offset(offset).Ordering(ordering).PulpHrefIn(pulpHrefIn).PulpIdIn(pulpIdIn).Fields(fields).ExcludeFields(excludeFields).Execute()

List collections



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
    name := "name_example" // string |  (optional)
    namespace := "namespace_example" // string |  (optional)
    offset := int32(56) // int32 | The initial index from which to return the results. (optional)
    ordering := []string{"Ordering_example"} // []string | Ordering  * `pulp_id` - Pulp id * `-pulp_id` - Pulp id (descending) * `pulp_created` - Pulp created * `-pulp_created` - Pulp created (descending) * `pulp_last_updated` - Pulp last updated * `-pulp_last_updated` - Pulp last updated (descending) * `namespace` - Namespace * `-namespace` - Namespace (descending) * `name` - Name * `-name` - Name (descending) * `pk` - Pk * `-pk` - Pk (descending) (optional)
    pulpHrefIn := []string{"Inner_example"} // []string | Multiple values may be separated by commas. (optional)
    pulpIdIn := []string{"Inner_example"} // []string | Multiple values may be separated by commas. (optional)
    fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
    excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AnsibleCollectionsAPI.AnsibleCollectionsList(context.Background()).Limit(limit).Name(name).Namespace(namespace).Offset(offset).Ordering(ordering).PulpHrefIn(pulpHrefIn).PulpIdIn(pulpIdIn).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AnsibleCollectionsAPI.AnsibleCollectionsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AnsibleCollectionsList`: PaginatedansibleCollectionResponseList
    fmt.Fprintf(os.Stdout, "Response from `AnsibleCollectionsAPI.AnsibleCollectionsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAnsibleCollectionsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Number of results to return per page. | 
 **name** | **string** |  | 
 **namespace** | **string** |  | 
 **offset** | **int32** | The initial index from which to return the results. | 
 **ordering** | **[]string** | Ordering  * &#x60;pulp_id&#x60; - Pulp id * &#x60;-pulp_id&#x60; - Pulp id (descending) * &#x60;pulp_created&#x60; - Pulp created * &#x60;-pulp_created&#x60; - Pulp created (descending) * &#x60;pulp_last_updated&#x60; - Pulp last updated * &#x60;-pulp_last_updated&#x60; - Pulp last updated (descending) * &#x60;namespace&#x60; - Namespace * &#x60;-namespace&#x60; - Namespace (descending) * &#x60;name&#x60; - Name * &#x60;-name&#x60; - Name (descending) * &#x60;pk&#x60; - Pk * &#x60;-pk&#x60; - Pk (descending) | 
 **pulpHrefIn** | **[]string** | Multiple values may be separated by commas. | 
 **pulpIdIn** | **[]string** | Multiple values may be separated by commas. | 
 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

### Return type

[**PaginatedansibleCollectionResponseList**](PaginatedansibleCollectionResponseList.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AnsibleCollectionsListRoles

> ObjectRolesResponse AnsibleCollectionsListRoles(ctx, ansibleCollectionHref).Fields(fields).ExcludeFields(excludeFields).Execute()





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
    ansibleCollectionHref := "ansibleCollectionHref_example" // string | 
    fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
    excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AnsibleCollectionsAPI.AnsibleCollectionsListRoles(context.Background(), ansibleCollectionHref).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AnsibleCollectionsAPI.AnsibleCollectionsListRoles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AnsibleCollectionsListRoles`: ObjectRolesResponse
    fmt.Fprintf(os.Stdout, "Response from `AnsibleCollectionsAPI.AnsibleCollectionsListRoles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ansibleCollectionHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAnsibleCollectionsListRolesRequest struct via the builder pattern


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


## AnsibleCollectionsMyPermissions

> MyPermissionsResponse AnsibleCollectionsMyPermissions(ctx, ansibleCollectionHref).Fields(fields).ExcludeFields(excludeFields).Execute()





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
    ansibleCollectionHref := "ansibleCollectionHref_example" // string | 
    fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
    excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AnsibleCollectionsAPI.AnsibleCollectionsMyPermissions(context.Background(), ansibleCollectionHref).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AnsibleCollectionsAPI.AnsibleCollectionsMyPermissions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AnsibleCollectionsMyPermissions`: MyPermissionsResponse
    fmt.Fprintf(os.Stdout, "Response from `AnsibleCollectionsAPI.AnsibleCollectionsMyPermissions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ansibleCollectionHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAnsibleCollectionsMyPermissionsRequest struct via the builder pattern


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


## AnsibleCollectionsRemoveRole

> NestedRoleResponse AnsibleCollectionsRemoveRole(ctx, ansibleCollectionHref).NestedRole(nestedRole).Execute()





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
    ansibleCollectionHref := "ansibleCollectionHref_example" // string | 
    nestedRole := *openapiclient.NewNestedRole("Role_example") // NestedRole | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AnsibleCollectionsAPI.AnsibleCollectionsRemoveRole(context.Background(), ansibleCollectionHref).NestedRole(nestedRole).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AnsibleCollectionsAPI.AnsibleCollectionsRemoveRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AnsibleCollectionsRemoveRole`: NestedRoleResponse
    fmt.Fprintf(os.Stdout, "Response from `AnsibleCollectionsAPI.AnsibleCollectionsRemoveRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ansibleCollectionHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAnsibleCollectionsRemoveRoleRequest struct via the builder pattern


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


## UploadCollection

> AsyncOperationResponse UploadCollection(ctx).File(file).Sha256(sha256).ExpectedNamespace(expectedNamespace).ExpectedName(expectedName).ExpectedVersion(expectedVersion).Execute()

Upload a collection



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
    file := os.NewFile(1234, "some_file") // *os.File | The Collection tarball.
    sha256 := "sha256_example" // string | An optional sha256 checksum of the uploaded file. (optional)
    expectedNamespace := "expectedNamespace_example" // string | The expected 'namespace' of the Collection to be verified against the metadata during import. (optional)
    expectedName := "expectedName_example" // string | The expected 'name' of the Collection to be verified against the metadata during import. (optional)
    expectedVersion := "expectedVersion_example" // string | The expected version of the Collection to be verified against the metadata during import. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AnsibleCollectionsAPI.UploadCollection(context.Background()).File(file).Sha256(sha256).ExpectedNamespace(expectedNamespace).ExpectedName(expectedName).ExpectedVersion(expectedVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AnsibleCollectionsAPI.UploadCollection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UploadCollection`: AsyncOperationResponse
    fmt.Fprintf(os.Stdout, "Response from `AnsibleCollectionsAPI.UploadCollection`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUploadCollectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **file** | ***os.File** | The Collection tarball. | 
 **sha256** | **string** | An optional sha256 checksum of the uploaded file. | 
 **expectedNamespace** | **string** | The expected &#39;namespace&#39; of the Collection to be verified against the metadata during import. | 
 **expectedName** | **string** | The expected &#39;name&#39; of the Collection to be verified against the metadata during import. | 
 **expectedVersion** | **string** | The expected version of the Collection to be verified against the metadata during import. | 

### Return type

[**AsyncOperationResponse**](AsyncOperationResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: multipart/form-data, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

