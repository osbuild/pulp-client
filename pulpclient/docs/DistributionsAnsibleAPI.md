# \DistributionsAnsibleAPI

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DistributionsAnsibleAnsibleAddRole**](DistributionsAnsibleAPI.md#DistributionsAnsibleAnsibleAddRole) | **Post** {ansible_ansible_distribution_href}add_role/ | 
[**DistributionsAnsibleAnsibleCreate**](DistributionsAnsibleAPI.md#DistributionsAnsibleAnsibleCreate) | **Post** /pulp/api/v3/distributions/ansible/ansible/ | Create an ansible distribution
[**DistributionsAnsibleAnsibleDelete**](DistributionsAnsibleAPI.md#DistributionsAnsibleAnsibleDelete) | **Delete** {ansible_ansible_distribution_href} | Delete an ansible distribution
[**DistributionsAnsibleAnsibleList**](DistributionsAnsibleAPI.md#DistributionsAnsibleAnsibleList) | **Get** /pulp/api/v3/distributions/ansible/ansible/ | List ansible distributions
[**DistributionsAnsibleAnsibleListRoles**](DistributionsAnsibleAPI.md#DistributionsAnsibleAnsibleListRoles) | **Get** {ansible_ansible_distribution_href}list_roles/ | 
[**DistributionsAnsibleAnsibleMyPermissions**](DistributionsAnsibleAPI.md#DistributionsAnsibleAnsibleMyPermissions) | **Get** {ansible_ansible_distribution_href}my_permissions/ | 
[**DistributionsAnsibleAnsiblePartialUpdate**](DistributionsAnsibleAPI.md#DistributionsAnsibleAnsiblePartialUpdate) | **Patch** {ansible_ansible_distribution_href} | Update an ansible distribution
[**DistributionsAnsibleAnsibleRead**](DistributionsAnsibleAPI.md#DistributionsAnsibleAnsibleRead) | **Get** {ansible_ansible_distribution_href} | Inspect an ansible distribution
[**DistributionsAnsibleAnsibleRemoveRole**](DistributionsAnsibleAPI.md#DistributionsAnsibleAnsibleRemoveRole) | **Post** {ansible_ansible_distribution_href}remove_role/ | 
[**DistributionsAnsibleAnsibleUpdate**](DistributionsAnsibleAPI.md#DistributionsAnsibleAnsibleUpdate) | **Put** {ansible_ansible_distribution_href} | Update an ansible distribution



## DistributionsAnsibleAnsibleAddRole

> NestedRoleResponse DistributionsAnsibleAnsibleAddRole(ctx, ansibleAnsibleDistributionHref).NestedRole(nestedRole).Execute()





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
    ansibleAnsibleDistributionHref := "ansibleAnsibleDistributionHref_example" // string | 
    nestedRole := *openapiclient.NewNestedRole("Role_example") // NestedRole | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DistributionsAnsibleAPI.DistributionsAnsibleAnsibleAddRole(context.Background(), ansibleAnsibleDistributionHref).NestedRole(nestedRole).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DistributionsAnsibleAPI.DistributionsAnsibleAnsibleAddRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DistributionsAnsibleAnsibleAddRole`: NestedRoleResponse
    fmt.Fprintf(os.Stdout, "Response from `DistributionsAnsibleAPI.DistributionsAnsibleAnsibleAddRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ansibleAnsibleDistributionHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDistributionsAnsibleAnsibleAddRoleRequest struct via the builder pattern


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


## DistributionsAnsibleAnsibleCreate

> AsyncOperationResponse DistributionsAnsibleAnsibleCreate(ctx).AnsibleAnsibleDistribution(ansibleAnsibleDistribution).Execute()

Create an ansible distribution



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
    ansibleAnsibleDistribution := *openapiclient.NewAnsibleAnsibleDistribution("BasePath_example", "Name_example") // AnsibleAnsibleDistribution | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DistributionsAnsibleAPI.DistributionsAnsibleAnsibleCreate(context.Background()).AnsibleAnsibleDistribution(ansibleAnsibleDistribution).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DistributionsAnsibleAPI.DistributionsAnsibleAnsibleCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DistributionsAnsibleAnsibleCreate`: AsyncOperationResponse
    fmt.Fprintf(os.Stdout, "Response from `DistributionsAnsibleAPI.DistributionsAnsibleAnsibleCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDistributionsAnsibleAnsibleCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ansibleAnsibleDistribution** | [**AnsibleAnsibleDistribution**](AnsibleAnsibleDistribution.md) |  | 

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


## DistributionsAnsibleAnsibleDelete

> AsyncOperationResponse DistributionsAnsibleAnsibleDelete(ctx, ansibleAnsibleDistributionHref).Execute()

Delete an ansible distribution



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
    ansibleAnsibleDistributionHref := "ansibleAnsibleDistributionHref_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DistributionsAnsibleAPI.DistributionsAnsibleAnsibleDelete(context.Background(), ansibleAnsibleDistributionHref).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DistributionsAnsibleAPI.DistributionsAnsibleAnsibleDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DistributionsAnsibleAnsibleDelete`: AsyncOperationResponse
    fmt.Fprintf(os.Stdout, "Response from `DistributionsAnsibleAPI.DistributionsAnsibleAnsibleDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ansibleAnsibleDistributionHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDistributionsAnsibleAnsibleDeleteRequest struct via the builder pattern


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


## DistributionsAnsibleAnsibleList

> PaginatedansibleAnsibleDistributionResponseList DistributionsAnsibleAnsibleList(ctx).BasePath(basePath).BasePathContains(basePathContains).BasePathIcontains(basePathIcontains).BasePathIn(basePathIn).Limit(limit).Name(name).NameContains(nameContains).NameIcontains(nameIcontains).NameIn(nameIn).NameStartswith(nameStartswith).Offset(offset).Ordering(ordering).PulpHrefIn(pulpHrefIn).PulpIdIn(pulpIdIn).PulpLabelSelect(pulpLabelSelect).Repository(repository).RepositoryIn(repositoryIn).WithContent(withContent).Fields(fields).ExcludeFields(excludeFields).Execute()

List ansible distributions



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
    basePath := "basePath_example" // string | Filter results where base_path matches value (optional)
    basePathContains := "basePathContains_example" // string | Filter results where base_path contains value (optional)
    basePathIcontains := "basePathIcontains_example" // string | Filter results where base_path contains value (optional)
    basePathIn := []string{"Inner_example"} // []string | Filter results where base_path is in a comma-separated list of values (optional)
    limit := int32(56) // int32 | Number of results to return per page. (optional)
    name := "name_example" // string | Filter results where name matches value (optional)
    nameContains := "nameContains_example" // string | Filter results where name contains value (optional)
    nameIcontains := "nameIcontains_example" // string | Filter results where name contains value (optional)
    nameIn := []string{"Inner_example"} // []string | Filter results where name is in a comma-separated list of values (optional)
    nameStartswith := "nameStartswith_example" // string | Filter results where name starts with value (optional)
    offset := int32(56) // int32 | The initial index from which to return the results. (optional)
    ordering := []string{"Ordering_example"} // []string | Ordering  * `pulp_id` - Pulp id * `-pulp_id` - Pulp id (descending) * `pulp_created` - Pulp created * `-pulp_created` - Pulp created (descending) * `pulp_last_updated` - Pulp last updated * `-pulp_last_updated` - Pulp last updated (descending) * `pulp_type` - Pulp type * `-pulp_type` - Pulp type (descending) * `name` - Name * `-name` - Name (descending) * `pulp_labels` - Pulp labels * `-pulp_labels` - Pulp labels (descending) * `base_path` - Base path * `-base_path` - Base path (descending) * `hidden` - Hidden * `-hidden` - Hidden (descending) * `pk` - Pk * `-pk` - Pk (descending) (optional)
    pulpHrefIn := []string{"Inner_example"} // []string | Multiple values may be separated by commas. (optional)
    pulpIdIn := []string{"Inner_example"} // []string | Multiple values may be separated by commas. (optional)
    pulpLabelSelect := "pulpLabelSelect_example" // string | Filter labels by search string (optional)
    repository := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Filter results where repository matches value (optional)
    repositoryIn := []string{"Inner_example"} // []string | Filter results where repository is in a comma-separated list of values (optional)
    withContent := "withContent_example" // string | Filter distributions based on the content served by them (optional)
    fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
    excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DistributionsAnsibleAPI.DistributionsAnsibleAnsibleList(context.Background()).BasePath(basePath).BasePathContains(basePathContains).BasePathIcontains(basePathIcontains).BasePathIn(basePathIn).Limit(limit).Name(name).NameContains(nameContains).NameIcontains(nameIcontains).NameIn(nameIn).NameStartswith(nameStartswith).Offset(offset).Ordering(ordering).PulpHrefIn(pulpHrefIn).PulpIdIn(pulpIdIn).PulpLabelSelect(pulpLabelSelect).Repository(repository).RepositoryIn(repositoryIn).WithContent(withContent).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DistributionsAnsibleAPI.DistributionsAnsibleAnsibleList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DistributionsAnsibleAnsibleList`: PaginatedansibleAnsibleDistributionResponseList
    fmt.Fprintf(os.Stdout, "Response from `DistributionsAnsibleAPI.DistributionsAnsibleAnsibleList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDistributionsAnsibleAnsibleListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **basePath** | **string** | Filter results where base_path matches value | 
 **basePathContains** | **string** | Filter results where base_path contains value | 
 **basePathIcontains** | **string** | Filter results where base_path contains value | 
 **basePathIn** | **[]string** | Filter results where base_path is in a comma-separated list of values | 
 **limit** | **int32** | Number of results to return per page. | 
 **name** | **string** | Filter results where name matches value | 
 **nameContains** | **string** | Filter results where name contains value | 
 **nameIcontains** | **string** | Filter results where name contains value | 
 **nameIn** | **[]string** | Filter results where name is in a comma-separated list of values | 
 **nameStartswith** | **string** | Filter results where name starts with value | 
 **offset** | **int32** | The initial index from which to return the results. | 
 **ordering** | **[]string** | Ordering  * &#x60;pulp_id&#x60; - Pulp id * &#x60;-pulp_id&#x60; - Pulp id (descending) * &#x60;pulp_created&#x60; - Pulp created * &#x60;-pulp_created&#x60; - Pulp created (descending) * &#x60;pulp_last_updated&#x60; - Pulp last updated * &#x60;-pulp_last_updated&#x60; - Pulp last updated (descending) * &#x60;pulp_type&#x60; - Pulp type * &#x60;-pulp_type&#x60; - Pulp type (descending) * &#x60;name&#x60; - Name * &#x60;-name&#x60; - Name (descending) * &#x60;pulp_labels&#x60; - Pulp labels * &#x60;-pulp_labels&#x60; - Pulp labels (descending) * &#x60;base_path&#x60; - Base path * &#x60;-base_path&#x60; - Base path (descending) * &#x60;hidden&#x60; - Hidden * &#x60;-hidden&#x60; - Hidden (descending) * &#x60;pk&#x60; - Pk * &#x60;-pk&#x60; - Pk (descending) | 
 **pulpHrefIn** | **[]string** | Multiple values may be separated by commas. | 
 **pulpIdIn** | **[]string** | Multiple values may be separated by commas. | 
 **pulpLabelSelect** | **string** | Filter labels by search string | 
 **repository** | **string** | Filter results where repository matches value | 
 **repositoryIn** | **[]string** | Filter results where repository is in a comma-separated list of values | 
 **withContent** | **string** | Filter distributions based on the content served by them | 
 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

### Return type

[**PaginatedansibleAnsibleDistributionResponseList**](PaginatedansibleAnsibleDistributionResponseList.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DistributionsAnsibleAnsibleListRoles

> ObjectRolesResponse DistributionsAnsibleAnsibleListRoles(ctx, ansibleAnsibleDistributionHref).Fields(fields).ExcludeFields(excludeFields).Execute()





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
    ansibleAnsibleDistributionHref := "ansibleAnsibleDistributionHref_example" // string | 
    fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
    excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DistributionsAnsibleAPI.DistributionsAnsibleAnsibleListRoles(context.Background(), ansibleAnsibleDistributionHref).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DistributionsAnsibleAPI.DistributionsAnsibleAnsibleListRoles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DistributionsAnsibleAnsibleListRoles`: ObjectRolesResponse
    fmt.Fprintf(os.Stdout, "Response from `DistributionsAnsibleAPI.DistributionsAnsibleAnsibleListRoles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ansibleAnsibleDistributionHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDistributionsAnsibleAnsibleListRolesRequest struct via the builder pattern


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


## DistributionsAnsibleAnsibleMyPermissions

> MyPermissionsResponse DistributionsAnsibleAnsibleMyPermissions(ctx, ansibleAnsibleDistributionHref).Fields(fields).ExcludeFields(excludeFields).Execute()





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
    ansibleAnsibleDistributionHref := "ansibleAnsibleDistributionHref_example" // string | 
    fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
    excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DistributionsAnsibleAPI.DistributionsAnsibleAnsibleMyPermissions(context.Background(), ansibleAnsibleDistributionHref).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DistributionsAnsibleAPI.DistributionsAnsibleAnsibleMyPermissions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DistributionsAnsibleAnsibleMyPermissions`: MyPermissionsResponse
    fmt.Fprintf(os.Stdout, "Response from `DistributionsAnsibleAPI.DistributionsAnsibleAnsibleMyPermissions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ansibleAnsibleDistributionHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDistributionsAnsibleAnsibleMyPermissionsRequest struct via the builder pattern


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


## DistributionsAnsibleAnsiblePartialUpdate

> AsyncOperationResponse DistributionsAnsibleAnsiblePartialUpdate(ctx, ansibleAnsibleDistributionHref).PatchedansibleAnsibleDistribution(patchedansibleAnsibleDistribution).Execute()

Update an ansible distribution



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
    ansibleAnsibleDistributionHref := "ansibleAnsibleDistributionHref_example" // string | 
    patchedansibleAnsibleDistribution := *openapiclient.NewPatchedansibleAnsibleDistribution() // PatchedansibleAnsibleDistribution | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DistributionsAnsibleAPI.DistributionsAnsibleAnsiblePartialUpdate(context.Background(), ansibleAnsibleDistributionHref).PatchedansibleAnsibleDistribution(patchedansibleAnsibleDistribution).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DistributionsAnsibleAPI.DistributionsAnsibleAnsiblePartialUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DistributionsAnsibleAnsiblePartialUpdate`: AsyncOperationResponse
    fmt.Fprintf(os.Stdout, "Response from `DistributionsAnsibleAPI.DistributionsAnsibleAnsiblePartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ansibleAnsibleDistributionHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDistributionsAnsibleAnsiblePartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedansibleAnsibleDistribution** | [**PatchedansibleAnsibleDistribution**](PatchedansibleAnsibleDistribution.md) |  | 

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


## DistributionsAnsibleAnsibleRead

> AnsibleAnsibleDistributionResponse DistributionsAnsibleAnsibleRead(ctx, ansibleAnsibleDistributionHref).Fields(fields).ExcludeFields(excludeFields).Execute()

Inspect an ansible distribution



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
    ansibleAnsibleDistributionHref := "ansibleAnsibleDistributionHref_example" // string | 
    fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
    excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DistributionsAnsibleAPI.DistributionsAnsibleAnsibleRead(context.Background(), ansibleAnsibleDistributionHref).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DistributionsAnsibleAPI.DistributionsAnsibleAnsibleRead``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DistributionsAnsibleAnsibleRead`: AnsibleAnsibleDistributionResponse
    fmt.Fprintf(os.Stdout, "Response from `DistributionsAnsibleAPI.DistributionsAnsibleAnsibleRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ansibleAnsibleDistributionHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDistributionsAnsibleAnsibleReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

### Return type

[**AnsibleAnsibleDistributionResponse**](AnsibleAnsibleDistributionResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DistributionsAnsibleAnsibleRemoveRole

> NestedRoleResponse DistributionsAnsibleAnsibleRemoveRole(ctx, ansibleAnsibleDistributionHref).NestedRole(nestedRole).Execute()





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
    ansibleAnsibleDistributionHref := "ansibleAnsibleDistributionHref_example" // string | 
    nestedRole := *openapiclient.NewNestedRole("Role_example") // NestedRole | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DistributionsAnsibleAPI.DistributionsAnsibleAnsibleRemoveRole(context.Background(), ansibleAnsibleDistributionHref).NestedRole(nestedRole).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DistributionsAnsibleAPI.DistributionsAnsibleAnsibleRemoveRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DistributionsAnsibleAnsibleRemoveRole`: NestedRoleResponse
    fmt.Fprintf(os.Stdout, "Response from `DistributionsAnsibleAPI.DistributionsAnsibleAnsibleRemoveRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ansibleAnsibleDistributionHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDistributionsAnsibleAnsibleRemoveRoleRequest struct via the builder pattern


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


## DistributionsAnsibleAnsibleUpdate

> AsyncOperationResponse DistributionsAnsibleAnsibleUpdate(ctx, ansibleAnsibleDistributionHref).AnsibleAnsibleDistribution(ansibleAnsibleDistribution).Execute()

Update an ansible distribution



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
    ansibleAnsibleDistributionHref := "ansibleAnsibleDistributionHref_example" // string | 
    ansibleAnsibleDistribution := *openapiclient.NewAnsibleAnsibleDistribution("BasePath_example", "Name_example") // AnsibleAnsibleDistribution | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DistributionsAnsibleAPI.DistributionsAnsibleAnsibleUpdate(context.Background(), ansibleAnsibleDistributionHref).AnsibleAnsibleDistribution(ansibleAnsibleDistribution).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DistributionsAnsibleAPI.DistributionsAnsibleAnsibleUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DistributionsAnsibleAnsibleUpdate`: AsyncOperationResponse
    fmt.Fprintf(os.Stdout, "Response from `DistributionsAnsibleAPI.DistributionsAnsibleAnsibleUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ansibleAnsibleDistributionHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDistributionsAnsibleAnsibleUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ansibleAnsibleDistribution** | [**AnsibleAnsibleDistribution**](AnsibleAnsibleDistribution.md) |  | 

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

