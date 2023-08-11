# \RepositoriesAnsibleAPI

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RepositoriesAnsibleAnsibleAddRole**](RepositoriesAnsibleAPI.md#RepositoriesAnsibleAnsibleAddRole) | **Post** {ansible_ansible_repository_href}add_role/ | 
[**RepositoriesAnsibleAnsibleCopyCollectionVersion**](RepositoriesAnsibleAPI.md#RepositoriesAnsibleAnsibleCopyCollectionVersion) | **Post** {ansible_ansible_repository_href}copy_collection_version/ | 
[**RepositoriesAnsibleAnsibleCreate**](RepositoriesAnsibleAPI.md#RepositoriesAnsibleAnsibleCreate) | **Post** /pulp/api/v3/repositories/ansible/ansible/ | Create an ansible repository
[**RepositoriesAnsibleAnsibleDelete**](RepositoriesAnsibleAPI.md#RepositoriesAnsibleAnsibleDelete) | **Delete** {ansible_ansible_repository_href} | Delete an ansible repository
[**RepositoriesAnsibleAnsibleList**](RepositoriesAnsibleAPI.md#RepositoriesAnsibleAnsibleList) | **Get** /pulp/api/v3/repositories/ansible/ansible/ | List ansible repositorys
[**RepositoriesAnsibleAnsibleListRoles**](RepositoriesAnsibleAPI.md#RepositoriesAnsibleAnsibleListRoles) | **Get** {ansible_ansible_repository_href}list_roles/ | 
[**RepositoriesAnsibleAnsibleMark**](RepositoriesAnsibleAPI.md#RepositoriesAnsibleAnsibleMark) | **Post** {ansible_ansible_repository_href}mark/ | 
[**RepositoriesAnsibleAnsibleModify**](RepositoriesAnsibleAPI.md#RepositoriesAnsibleAnsibleModify) | **Post** {ansible_ansible_repository_href}modify/ | Modify Repository Content
[**RepositoriesAnsibleAnsibleMoveCollectionVersion**](RepositoriesAnsibleAPI.md#RepositoriesAnsibleAnsibleMoveCollectionVersion) | **Post** {ansible_ansible_repository_href}move_collection_version/ | 
[**RepositoriesAnsibleAnsibleMyPermissions**](RepositoriesAnsibleAPI.md#RepositoriesAnsibleAnsibleMyPermissions) | **Get** {ansible_ansible_repository_href}my_permissions/ | 
[**RepositoriesAnsibleAnsiblePartialUpdate**](RepositoriesAnsibleAPI.md#RepositoriesAnsibleAnsiblePartialUpdate) | **Patch** {ansible_ansible_repository_href} | Update an ansible repository
[**RepositoriesAnsibleAnsibleRead**](RepositoriesAnsibleAPI.md#RepositoriesAnsibleAnsibleRead) | **Get** {ansible_ansible_repository_href} | Inspect an ansible repository
[**RepositoriesAnsibleAnsibleRebuildMetadata**](RepositoriesAnsibleAPI.md#RepositoriesAnsibleAnsibleRebuildMetadata) | **Post** {ansible_ansible_repository_href}rebuild_metadata/ | 
[**RepositoriesAnsibleAnsibleRemoveRole**](RepositoriesAnsibleAPI.md#RepositoriesAnsibleAnsibleRemoveRole) | **Post** {ansible_ansible_repository_href}remove_role/ | 
[**RepositoriesAnsibleAnsibleSign**](RepositoriesAnsibleAPI.md#RepositoriesAnsibleAnsibleSign) | **Post** {ansible_ansible_repository_href}sign/ | 
[**RepositoriesAnsibleAnsibleSync**](RepositoriesAnsibleAPI.md#RepositoriesAnsibleAnsibleSync) | **Post** {ansible_ansible_repository_href}sync/ | 
[**RepositoriesAnsibleAnsibleUnmark**](RepositoriesAnsibleAPI.md#RepositoriesAnsibleAnsibleUnmark) | **Post** {ansible_ansible_repository_href}unmark/ | 
[**RepositoriesAnsibleAnsibleUpdate**](RepositoriesAnsibleAPI.md#RepositoriesAnsibleAnsibleUpdate) | **Put** {ansible_ansible_repository_href} | Update an ansible repository



## RepositoriesAnsibleAnsibleAddRole

> NestedRoleResponse RepositoriesAnsibleAnsibleAddRole(ctx, ansibleAnsibleRepositoryHref).NestedRole(nestedRole).Execute()





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
    ansibleAnsibleRepositoryHref := "ansibleAnsibleRepositoryHref_example" // string | 
    nestedRole := *openapiclient.NewNestedRole("Role_example") // NestedRole | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RepositoriesAnsibleAPI.RepositoriesAnsibleAnsibleAddRole(context.Background(), ansibleAnsibleRepositoryHref).NestedRole(nestedRole).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesAnsibleAPI.RepositoriesAnsibleAnsibleAddRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RepositoriesAnsibleAnsibleAddRole`: NestedRoleResponse
    fmt.Fprintf(os.Stdout, "Response from `RepositoriesAnsibleAPI.RepositoriesAnsibleAnsibleAddRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ansibleAnsibleRepositoryHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesAnsibleAnsibleAddRoleRequest struct via the builder pattern


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


## RepositoriesAnsibleAnsibleCopyCollectionVersion

> AsyncOperationResponse RepositoriesAnsibleAnsibleCopyCollectionVersion(ctx, ansibleAnsibleRepositoryHref).CollectionVersionCopyMove(collectionVersionCopyMove).Execute()





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
    ansibleAnsibleRepositoryHref := "ansibleAnsibleRepositoryHref_example" // string | 
    collectionVersionCopyMove := *openapiclient.NewCollectionVersionCopyMove([]string{"CollectionVersions_example"}, []string{"DestinationRepositories_example"}) // CollectionVersionCopyMove | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RepositoriesAnsibleAPI.RepositoriesAnsibleAnsibleCopyCollectionVersion(context.Background(), ansibleAnsibleRepositoryHref).CollectionVersionCopyMove(collectionVersionCopyMove).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesAnsibleAPI.RepositoriesAnsibleAnsibleCopyCollectionVersion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RepositoriesAnsibleAnsibleCopyCollectionVersion`: AsyncOperationResponse
    fmt.Fprintf(os.Stdout, "Response from `RepositoriesAnsibleAPI.RepositoriesAnsibleAnsibleCopyCollectionVersion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ansibleAnsibleRepositoryHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesAnsibleAnsibleCopyCollectionVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **collectionVersionCopyMove** | [**CollectionVersionCopyMove**](CollectionVersionCopyMove.md) |  | 

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


## RepositoriesAnsibleAnsibleCreate

> AnsibleAnsibleRepositoryResponse RepositoriesAnsibleAnsibleCreate(ctx).AnsibleAnsibleRepository(ansibleAnsibleRepository).Execute()

Create an ansible repository



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
    ansibleAnsibleRepository := *openapiclient.NewAnsibleAnsibleRepository("Name_example") // AnsibleAnsibleRepository | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RepositoriesAnsibleAPI.RepositoriesAnsibleAnsibleCreate(context.Background()).AnsibleAnsibleRepository(ansibleAnsibleRepository).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesAnsibleAPI.RepositoriesAnsibleAnsibleCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RepositoriesAnsibleAnsibleCreate`: AnsibleAnsibleRepositoryResponse
    fmt.Fprintf(os.Stdout, "Response from `RepositoriesAnsibleAPI.RepositoriesAnsibleAnsibleCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesAnsibleAnsibleCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ansibleAnsibleRepository** | [**AnsibleAnsibleRepository**](AnsibleAnsibleRepository.md) |  | 

### Return type

[**AnsibleAnsibleRepositoryResponse**](AnsibleAnsibleRepositoryResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RepositoriesAnsibleAnsibleDelete

> AsyncOperationResponse RepositoriesAnsibleAnsibleDelete(ctx, ansibleAnsibleRepositoryHref).Execute()

Delete an ansible repository



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
    ansibleAnsibleRepositoryHref := "ansibleAnsibleRepositoryHref_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RepositoriesAnsibleAPI.RepositoriesAnsibleAnsibleDelete(context.Background(), ansibleAnsibleRepositoryHref).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesAnsibleAPI.RepositoriesAnsibleAnsibleDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RepositoriesAnsibleAnsibleDelete`: AsyncOperationResponse
    fmt.Fprintf(os.Stdout, "Response from `RepositoriesAnsibleAPI.RepositoriesAnsibleAnsibleDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ansibleAnsibleRepositoryHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesAnsibleAnsibleDeleteRequest struct via the builder pattern


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


## RepositoriesAnsibleAnsibleList

> PaginatedansibleAnsibleRepositoryResponseList RepositoriesAnsibleAnsibleList(ctx).LatestWithContent(latestWithContent).Limit(limit).Name(name).NameContains(nameContains).NameIcontains(nameIcontains).NameIn(nameIn).NameStartswith(nameStartswith).Offset(offset).Ordering(ordering).PulpHrefIn(pulpHrefIn).PulpIdIn(pulpIdIn).PulpLabelSelect(pulpLabelSelect).Remote(remote).RetainRepoVersions(retainRepoVersions).RetainRepoVersionsGt(retainRepoVersionsGt).RetainRepoVersionsGte(retainRepoVersionsGte).RetainRepoVersionsIsnull(retainRepoVersionsIsnull).RetainRepoVersionsLt(retainRepoVersionsLt).RetainRepoVersionsLte(retainRepoVersionsLte).RetainRepoVersionsNe(retainRepoVersionsNe).RetainRepoVersionsRange(retainRepoVersionsRange).WithContent(withContent).Fields(fields).ExcludeFields(excludeFields).Execute()

List ansible repositorys



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
    latestWithContent := "latestWithContent_example" // string | Content Unit referenced by HREF (optional)
    limit := int32(56) // int32 | Number of results to return per page. (optional)
    name := "name_example" // string | Filter results where name matches value (optional)
    nameContains := "nameContains_example" // string | Filter results where name contains value (optional)
    nameIcontains := "nameIcontains_example" // string | Filter results where name contains value (optional)
    nameIn := []string{"Inner_example"} // []string | Filter results where name is in a comma-separated list of values (optional)
    nameStartswith := "nameStartswith_example" // string | Filter results where name starts with value (optional)
    offset := int32(56) // int32 | The initial index from which to return the results. (optional)
    ordering := []string{"Ordering_example"} // []string | Ordering  * `pulp_id` - Pulp id * `-pulp_id` - Pulp id (descending) * `pulp_created` - Pulp created * `-pulp_created` - Pulp created (descending) * `pulp_last_updated` - Pulp last updated * `-pulp_last_updated` - Pulp last updated (descending) * `pulp_type` - Pulp type * `-pulp_type` - Pulp type (descending) * `name` - Name * `-name` - Name (descending) * `pulp_labels` - Pulp labels * `-pulp_labels` - Pulp labels (descending) * `description` - Description * `-description` - Description (descending) * `next_version` - Next version * `-next_version` - Next version (descending) * `retain_repo_versions` - Retain repo versions * `-retain_repo_versions` - Retain repo versions (descending) * `user_hidden` - User hidden * `-user_hidden` - User hidden (descending) * `pk` - Pk * `-pk` - Pk (descending) (optional)
    pulpHrefIn := []string{"Inner_example"} // []string | Multiple values may be separated by commas. (optional)
    pulpIdIn := []string{"Inner_example"} // []string | Multiple values may be separated by commas. (optional)
    pulpLabelSelect := "pulpLabelSelect_example" // string | Filter labels by search string (optional)
    remote := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Foreign Key referenced by HREF (optional)
    retainRepoVersions := int32(56) // int32 | Filter results where retain_repo_versions matches value (optional)
    retainRepoVersionsGt := int32(56) // int32 | Filter results where retain_repo_versions is greater than value (optional)
    retainRepoVersionsGte := int32(56) // int32 | Filter results where retain_repo_versions is greater than or equal to value (optional)
    retainRepoVersionsIsnull := true // bool | Filter results where retain_repo_versions has a null value (optional)
    retainRepoVersionsLt := int32(56) // int32 | Filter results where retain_repo_versions is less than value (optional)
    retainRepoVersionsLte := int32(56) // int32 | Filter results where retain_repo_versions is less than or equal to value (optional)
    retainRepoVersionsNe := int32(56) // int32 | Filter results where retain_repo_versions not equal to value (optional)
    retainRepoVersionsRange := []int32{int32(123)} // []int32 | Filter results where retain_repo_versions is between two comma separated values (optional)
    withContent := "withContent_example" // string | Content Unit referenced by HREF (optional)
    fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
    excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RepositoriesAnsibleAPI.RepositoriesAnsibleAnsibleList(context.Background()).LatestWithContent(latestWithContent).Limit(limit).Name(name).NameContains(nameContains).NameIcontains(nameIcontains).NameIn(nameIn).NameStartswith(nameStartswith).Offset(offset).Ordering(ordering).PulpHrefIn(pulpHrefIn).PulpIdIn(pulpIdIn).PulpLabelSelect(pulpLabelSelect).Remote(remote).RetainRepoVersions(retainRepoVersions).RetainRepoVersionsGt(retainRepoVersionsGt).RetainRepoVersionsGte(retainRepoVersionsGte).RetainRepoVersionsIsnull(retainRepoVersionsIsnull).RetainRepoVersionsLt(retainRepoVersionsLt).RetainRepoVersionsLte(retainRepoVersionsLte).RetainRepoVersionsNe(retainRepoVersionsNe).RetainRepoVersionsRange(retainRepoVersionsRange).WithContent(withContent).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesAnsibleAPI.RepositoriesAnsibleAnsibleList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RepositoriesAnsibleAnsibleList`: PaginatedansibleAnsibleRepositoryResponseList
    fmt.Fprintf(os.Stdout, "Response from `RepositoriesAnsibleAPI.RepositoriesAnsibleAnsibleList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesAnsibleAnsibleListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **latestWithContent** | **string** | Content Unit referenced by HREF | 
 **limit** | **int32** | Number of results to return per page. | 
 **name** | **string** | Filter results where name matches value | 
 **nameContains** | **string** | Filter results where name contains value | 
 **nameIcontains** | **string** | Filter results where name contains value | 
 **nameIn** | **[]string** | Filter results where name is in a comma-separated list of values | 
 **nameStartswith** | **string** | Filter results where name starts with value | 
 **offset** | **int32** | The initial index from which to return the results. | 
 **ordering** | **[]string** | Ordering  * &#x60;pulp_id&#x60; - Pulp id * &#x60;-pulp_id&#x60; - Pulp id (descending) * &#x60;pulp_created&#x60; - Pulp created * &#x60;-pulp_created&#x60; - Pulp created (descending) * &#x60;pulp_last_updated&#x60; - Pulp last updated * &#x60;-pulp_last_updated&#x60; - Pulp last updated (descending) * &#x60;pulp_type&#x60; - Pulp type * &#x60;-pulp_type&#x60; - Pulp type (descending) * &#x60;name&#x60; - Name * &#x60;-name&#x60; - Name (descending) * &#x60;pulp_labels&#x60; - Pulp labels * &#x60;-pulp_labels&#x60; - Pulp labels (descending) * &#x60;description&#x60; - Description * &#x60;-description&#x60; - Description (descending) * &#x60;next_version&#x60; - Next version * &#x60;-next_version&#x60; - Next version (descending) * &#x60;retain_repo_versions&#x60; - Retain repo versions * &#x60;-retain_repo_versions&#x60; - Retain repo versions (descending) * &#x60;user_hidden&#x60; - User hidden * &#x60;-user_hidden&#x60; - User hidden (descending) * &#x60;pk&#x60; - Pk * &#x60;-pk&#x60; - Pk (descending) | 
 **pulpHrefIn** | **[]string** | Multiple values may be separated by commas. | 
 **pulpIdIn** | **[]string** | Multiple values may be separated by commas. | 
 **pulpLabelSelect** | **string** | Filter labels by search string | 
 **remote** | **string** | Foreign Key referenced by HREF | 
 **retainRepoVersions** | **int32** | Filter results where retain_repo_versions matches value | 
 **retainRepoVersionsGt** | **int32** | Filter results where retain_repo_versions is greater than value | 
 **retainRepoVersionsGte** | **int32** | Filter results where retain_repo_versions is greater than or equal to value | 
 **retainRepoVersionsIsnull** | **bool** | Filter results where retain_repo_versions has a null value | 
 **retainRepoVersionsLt** | **int32** | Filter results where retain_repo_versions is less than value | 
 **retainRepoVersionsLte** | **int32** | Filter results where retain_repo_versions is less than or equal to value | 
 **retainRepoVersionsNe** | **int32** | Filter results where retain_repo_versions not equal to value | 
 **retainRepoVersionsRange** | **[]int32** | Filter results where retain_repo_versions is between two comma separated values | 
 **withContent** | **string** | Content Unit referenced by HREF | 
 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

### Return type

[**PaginatedansibleAnsibleRepositoryResponseList**](PaginatedansibleAnsibleRepositoryResponseList.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RepositoriesAnsibleAnsibleListRoles

> ObjectRolesResponse RepositoriesAnsibleAnsibleListRoles(ctx, ansibleAnsibleRepositoryHref).Fields(fields).ExcludeFields(excludeFields).Execute()





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
    ansibleAnsibleRepositoryHref := "ansibleAnsibleRepositoryHref_example" // string | 
    fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
    excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RepositoriesAnsibleAPI.RepositoriesAnsibleAnsibleListRoles(context.Background(), ansibleAnsibleRepositoryHref).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesAnsibleAPI.RepositoriesAnsibleAnsibleListRoles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RepositoriesAnsibleAnsibleListRoles`: ObjectRolesResponse
    fmt.Fprintf(os.Stdout, "Response from `RepositoriesAnsibleAPI.RepositoriesAnsibleAnsibleListRoles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ansibleAnsibleRepositoryHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesAnsibleAnsibleListRolesRequest struct via the builder pattern


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


## RepositoriesAnsibleAnsibleMark

> AsyncOperationResponse RepositoriesAnsibleAnsibleMark(ctx, ansibleAnsibleRepositoryHref).AnsibleRepositoryMark(ansibleRepositoryMark).Execute()





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
    ansibleAnsibleRepositoryHref := "ansibleAnsibleRepositoryHref_example" // string | 
    ansibleRepositoryMark := *openapiclient.NewAnsibleRepositoryMark([]interface{}{nil}, "Value_example") // AnsibleRepositoryMark | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RepositoriesAnsibleAPI.RepositoriesAnsibleAnsibleMark(context.Background(), ansibleAnsibleRepositoryHref).AnsibleRepositoryMark(ansibleRepositoryMark).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesAnsibleAPI.RepositoriesAnsibleAnsibleMark``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RepositoriesAnsibleAnsibleMark`: AsyncOperationResponse
    fmt.Fprintf(os.Stdout, "Response from `RepositoriesAnsibleAPI.RepositoriesAnsibleAnsibleMark`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ansibleAnsibleRepositoryHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesAnsibleAnsibleMarkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ansibleRepositoryMark** | [**AnsibleRepositoryMark**](AnsibleRepositoryMark.md) |  | 

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


## RepositoriesAnsibleAnsibleModify

> AsyncOperationResponse RepositoriesAnsibleAnsibleModify(ctx, ansibleAnsibleRepositoryHref).RepositoryAddRemoveContent(repositoryAddRemoveContent).Execute()

Modify Repository Content



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
    ansibleAnsibleRepositoryHref := "ansibleAnsibleRepositoryHref_example" // string | 
    repositoryAddRemoveContent := *openapiclient.NewRepositoryAddRemoveContent() // RepositoryAddRemoveContent | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RepositoriesAnsibleAPI.RepositoriesAnsibleAnsibleModify(context.Background(), ansibleAnsibleRepositoryHref).RepositoryAddRemoveContent(repositoryAddRemoveContent).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesAnsibleAPI.RepositoriesAnsibleAnsibleModify``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RepositoriesAnsibleAnsibleModify`: AsyncOperationResponse
    fmt.Fprintf(os.Stdout, "Response from `RepositoriesAnsibleAPI.RepositoriesAnsibleAnsibleModify`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ansibleAnsibleRepositoryHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesAnsibleAnsibleModifyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **repositoryAddRemoveContent** | [**RepositoryAddRemoveContent**](RepositoryAddRemoveContent.md) |  | 

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


## RepositoriesAnsibleAnsibleMoveCollectionVersion

> AsyncOperationResponse RepositoriesAnsibleAnsibleMoveCollectionVersion(ctx, ansibleAnsibleRepositoryHref).CollectionVersionCopyMove(collectionVersionCopyMove).Execute()





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
    ansibleAnsibleRepositoryHref := "ansibleAnsibleRepositoryHref_example" // string | 
    collectionVersionCopyMove := *openapiclient.NewCollectionVersionCopyMove([]string{"CollectionVersions_example"}, []string{"DestinationRepositories_example"}) // CollectionVersionCopyMove | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RepositoriesAnsibleAPI.RepositoriesAnsibleAnsibleMoveCollectionVersion(context.Background(), ansibleAnsibleRepositoryHref).CollectionVersionCopyMove(collectionVersionCopyMove).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesAnsibleAPI.RepositoriesAnsibleAnsibleMoveCollectionVersion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RepositoriesAnsibleAnsibleMoveCollectionVersion`: AsyncOperationResponse
    fmt.Fprintf(os.Stdout, "Response from `RepositoriesAnsibleAPI.RepositoriesAnsibleAnsibleMoveCollectionVersion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ansibleAnsibleRepositoryHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesAnsibleAnsibleMoveCollectionVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **collectionVersionCopyMove** | [**CollectionVersionCopyMove**](CollectionVersionCopyMove.md) |  | 

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


## RepositoriesAnsibleAnsibleMyPermissions

> MyPermissionsResponse RepositoriesAnsibleAnsibleMyPermissions(ctx, ansibleAnsibleRepositoryHref).Fields(fields).ExcludeFields(excludeFields).Execute()





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
    ansibleAnsibleRepositoryHref := "ansibleAnsibleRepositoryHref_example" // string | 
    fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
    excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RepositoriesAnsibleAPI.RepositoriesAnsibleAnsibleMyPermissions(context.Background(), ansibleAnsibleRepositoryHref).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesAnsibleAPI.RepositoriesAnsibleAnsibleMyPermissions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RepositoriesAnsibleAnsibleMyPermissions`: MyPermissionsResponse
    fmt.Fprintf(os.Stdout, "Response from `RepositoriesAnsibleAPI.RepositoriesAnsibleAnsibleMyPermissions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ansibleAnsibleRepositoryHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesAnsibleAnsibleMyPermissionsRequest struct via the builder pattern


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


## RepositoriesAnsibleAnsiblePartialUpdate

> AsyncOperationResponse RepositoriesAnsibleAnsiblePartialUpdate(ctx, ansibleAnsibleRepositoryHref).PatchedansibleAnsibleRepository(patchedansibleAnsibleRepository).Execute()

Update an ansible repository



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
    ansibleAnsibleRepositoryHref := "ansibleAnsibleRepositoryHref_example" // string | 
    patchedansibleAnsibleRepository := *openapiclient.NewPatchedansibleAnsibleRepository() // PatchedansibleAnsibleRepository | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RepositoriesAnsibleAPI.RepositoriesAnsibleAnsiblePartialUpdate(context.Background(), ansibleAnsibleRepositoryHref).PatchedansibleAnsibleRepository(patchedansibleAnsibleRepository).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesAnsibleAPI.RepositoriesAnsibleAnsiblePartialUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RepositoriesAnsibleAnsiblePartialUpdate`: AsyncOperationResponse
    fmt.Fprintf(os.Stdout, "Response from `RepositoriesAnsibleAPI.RepositoriesAnsibleAnsiblePartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ansibleAnsibleRepositoryHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesAnsibleAnsiblePartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedansibleAnsibleRepository** | [**PatchedansibleAnsibleRepository**](PatchedansibleAnsibleRepository.md) |  | 

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


## RepositoriesAnsibleAnsibleRead

> AnsibleAnsibleRepositoryResponse RepositoriesAnsibleAnsibleRead(ctx, ansibleAnsibleRepositoryHref).Fields(fields).ExcludeFields(excludeFields).Execute()

Inspect an ansible repository



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
    ansibleAnsibleRepositoryHref := "ansibleAnsibleRepositoryHref_example" // string | 
    fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
    excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RepositoriesAnsibleAPI.RepositoriesAnsibleAnsibleRead(context.Background(), ansibleAnsibleRepositoryHref).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesAnsibleAPI.RepositoriesAnsibleAnsibleRead``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RepositoriesAnsibleAnsibleRead`: AnsibleAnsibleRepositoryResponse
    fmt.Fprintf(os.Stdout, "Response from `RepositoriesAnsibleAPI.RepositoriesAnsibleAnsibleRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ansibleAnsibleRepositoryHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesAnsibleAnsibleReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

### Return type

[**AnsibleAnsibleRepositoryResponse**](AnsibleAnsibleRepositoryResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RepositoriesAnsibleAnsibleRebuildMetadata

> AsyncOperationResponse RepositoriesAnsibleAnsibleRebuildMetadata(ctx, ansibleAnsibleRepositoryHref).AnsibleRepositoryRebuild(ansibleRepositoryRebuild).Execute()





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
    ansibleAnsibleRepositoryHref := "ansibleAnsibleRepositoryHref_example" // string | 
    ansibleRepositoryRebuild := *openapiclient.NewAnsibleRepositoryRebuild() // AnsibleRepositoryRebuild | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RepositoriesAnsibleAPI.RepositoriesAnsibleAnsibleRebuildMetadata(context.Background(), ansibleAnsibleRepositoryHref).AnsibleRepositoryRebuild(ansibleRepositoryRebuild).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesAnsibleAPI.RepositoriesAnsibleAnsibleRebuildMetadata``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RepositoriesAnsibleAnsibleRebuildMetadata`: AsyncOperationResponse
    fmt.Fprintf(os.Stdout, "Response from `RepositoriesAnsibleAPI.RepositoriesAnsibleAnsibleRebuildMetadata`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ansibleAnsibleRepositoryHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesAnsibleAnsibleRebuildMetadataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ansibleRepositoryRebuild** | [**AnsibleRepositoryRebuild**](AnsibleRepositoryRebuild.md) |  | 

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


## RepositoriesAnsibleAnsibleRemoveRole

> NestedRoleResponse RepositoriesAnsibleAnsibleRemoveRole(ctx, ansibleAnsibleRepositoryHref).NestedRole(nestedRole).Execute()





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
    ansibleAnsibleRepositoryHref := "ansibleAnsibleRepositoryHref_example" // string | 
    nestedRole := *openapiclient.NewNestedRole("Role_example") // NestedRole | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RepositoriesAnsibleAPI.RepositoriesAnsibleAnsibleRemoveRole(context.Background(), ansibleAnsibleRepositoryHref).NestedRole(nestedRole).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesAnsibleAPI.RepositoriesAnsibleAnsibleRemoveRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RepositoriesAnsibleAnsibleRemoveRole`: NestedRoleResponse
    fmt.Fprintf(os.Stdout, "Response from `RepositoriesAnsibleAPI.RepositoriesAnsibleAnsibleRemoveRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ansibleAnsibleRepositoryHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesAnsibleAnsibleRemoveRoleRequest struct via the builder pattern


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


## RepositoriesAnsibleAnsibleSign

> AsyncOperationResponse RepositoriesAnsibleAnsibleSign(ctx, ansibleAnsibleRepositoryHref).AnsibleRepositorySignature(ansibleRepositorySignature).Execute()





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
    ansibleAnsibleRepositoryHref := "ansibleAnsibleRepositoryHref_example" // string | 
    ansibleRepositorySignature := *openapiclient.NewAnsibleRepositorySignature([]interface{}{nil}, "SigningService_example") // AnsibleRepositorySignature | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RepositoriesAnsibleAPI.RepositoriesAnsibleAnsibleSign(context.Background(), ansibleAnsibleRepositoryHref).AnsibleRepositorySignature(ansibleRepositorySignature).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesAnsibleAPI.RepositoriesAnsibleAnsibleSign``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RepositoriesAnsibleAnsibleSign`: AsyncOperationResponse
    fmt.Fprintf(os.Stdout, "Response from `RepositoriesAnsibleAPI.RepositoriesAnsibleAnsibleSign`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ansibleAnsibleRepositoryHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesAnsibleAnsibleSignRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ansibleRepositorySignature** | [**AnsibleRepositorySignature**](AnsibleRepositorySignature.md) |  | 

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


## RepositoriesAnsibleAnsibleSync

> AsyncOperationResponse RepositoriesAnsibleAnsibleSync(ctx, ansibleAnsibleRepositoryHref).AnsibleRepositorySyncURL(ansibleRepositorySyncURL).Execute()





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
    ansibleAnsibleRepositoryHref := "ansibleAnsibleRepositoryHref_example" // string | 
    ansibleRepositorySyncURL := *openapiclient.NewAnsibleRepositorySyncURL() // AnsibleRepositorySyncURL | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RepositoriesAnsibleAPI.RepositoriesAnsibleAnsibleSync(context.Background(), ansibleAnsibleRepositoryHref).AnsibleRepositorySyncURL(ansibleRepositorySyncURL).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesAnsibleAPI.RepositoriesAnsibleAnsibleSync``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RepositoriesAnsibleAnsibleSync`: AsyncOperationResponse
    fmt.Fprintf(os.Stdout, "Response from `RepositoriesAnsibleAPI.RepositoriesAnsibleAnsibleSync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ansibleAnsibleRepositoryHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesAnsibleAnsibleSyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ansibleRepositorySyncURL** | [**AnsibleRepositorySyncURL**](AnsibleRepositorySyncURL.md) |  | 

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


## RepositoriesAnsibleAnsibleUnmark

> AsyncOperationResponse RepositoriesAnsibleAnsibleUnmark(ctx, ansibleAnsibleRepositoryHref).AnsibleRepositoryMark(ansibleRepositoryMark).Execute()





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
    ansibleAnsibleRepositoryHref := "ansibleAnsibleRepositoryHref_example" // string | 
    ansibleRepositoryMark := *openapiclient.NewAnsibleRepositoryMark([]interface{}{nil}, "Value_example") // AnsibleRepositoryMark | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RepositoriesAnsibleAPI.RepositoriesAnsibleAnsibleUnmark(context.Background(), ansibleAnsibleRepositoryHref).AnsibleRepositoryMark(ansibleRepositoryMark).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesAnsibleAPI.RepositoriesAnsibleAnsibleUnmark``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RepositoriesAnsibleAnsibleUnmark`: AsyncOperationResponse
    fmt.Fprintf(os.Stdout, "Response from `RepositoriesAnsibleAPI.RepositoriesAnsibleAnsibleUnmark`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ansibleAnsibleRepositoryHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesAnsibleAnsibleUnmarkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ansibleRepositoryMark** | [**AnsibleRepositoryMark**](AnsibleRepositoryMark.md) |  | 

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


## RepositoriesAnsibleAnsibleUpdate

> AsyncOperationResponse RepositoriesAnsibleAnsibleUpdate(ctx, ansibleAnsibleRepositoryHref).AnsibleAnsibleRepository(ansibleAnsibleRepository).Execute()

Update an ansible repository



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
    ansibleAnsibleRepositoryHref := "ansibleAnsibleRepositoryHref_example" // string | 
    ansibleAnsibleRepository := *openapiclient.NewAnsibleAnsibleRepository("Name_example") // AnsibleAnsibleRepository | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RepositoriesAnsibleAPI.RepositoriesAnsibleAnsibleUpdate(context.Background(), ansibleAnsibleRepositoryHref).AnsibleAnsibleRepository(ansibleAnsibleRepository).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesAnsibleAPI.RepositoriesAnsibleAnsibleUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RepositoriesAnsibleAnsibleUpdate`: AsyncOperationResponse
    fmt.Fprintf(os.Stdout, "Response from `RepositoriesAnsibleAPI.RepositoriesAnsibleAnsibleUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ansibleAnsibleRepositoryHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesAnsibleAnsibleUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ansibleAnsibleRepository** | [**AnsibleAnsibleRepository**](AnsibleAnsibleRepository.md) |  | 

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

