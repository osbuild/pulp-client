# \ContentCollectionDeprecationsAPI

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ContentAnsibleCollectionDeprecationsCreate**](ContentCollectionDeprecationsAPI.md#ContentAnsibleCollectionDeprecationsCreate) | **Post** /pulp/api/v3/content/ansible/collection_deprecations/ | Create an ansible collection deprecated
[**ContentAnsibleCollectionDeprecationsList**](ContentCollectionDeprecationsAPI.md#ContentAnsibleCollectionDeprecationsList) | **Get** /pulp/api/v3/content/ansible/collection_deprecations/ | List ansible collection deprecateds
[**ContentAnsibleCollectionDeprecationsRead**](ContentCollectionDeprecationsAPI.md#ContentAnsibleCollectionDeprecationsRead) | **Get** {ansible_ansible_collection_deprecated_href} | Inspect an ansible collection deprecated



## ContentAnsibleCollectionDeprecationsCreate

> AnsibleCollectionResponse ContentAnsibleCollectionDeprecationsCreate(ctx).AnsibleCollection(ansibleCollection).Execute()

Create an ansible collection deprecated



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
    ansibleCollection := *openapiclient.NewAnsibleCollection("Name_example", "Namespace_example") // AnsibleCollection | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContentCollectionDeprecationsAPI.ContentAnsibleCollectionDeprecationsCreate(context.Background()).AnsibleCollection(ansibleCollection).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContentCollectionDeprecationsAPI.ContentAnsibleCollectionDeprecationsCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContentAnsibleCollectionDeprecationsCreate`: AnsibleCollectionResponse
    fmt.Fprintf(os.Stdout, "Response from `ContentCollectionDeprecationsAPI.ContentAnsibleCollectionDeprecationsCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiContentAnsibleCollectionDeprecationsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ansibleCollection** | [**AnsibleCollection**](AnsibleCollection.md) |  | 

### Return type

[**AnsibleCollectionResponse**](AnsibleCollectionResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContentAnsibleCollectionDeprecationsList

> PaginatedansibleCollectionResponseList ContentAnsibleCollectionDeprecationsList(ctx).Limit(limit).Offset(offset).Ordering(ordering).PulpHrefIn(pulpHrefIn).PulpIdIn(pulpIdIn).RepositoryVersion(repositoryVersion).RepositoryVersionAdded(repositoryVersionAdded).RepositoryVersionRemoved(repositoryVersionRemoved).Fields(fields).ExcludeFields(excludeFields).Execute()

List ansible collection deprecateds



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
    offset := int32(56) // int32 | The initial index from which to return the results. (optional)
    ordering := []string{"Ordering_example"} // []string | Ordering  * `pk` - Pk * `-pk` - Pk (descending) (optional)
    pulpHrefIn := []string{"Inner_example"} // []string | Multiple values may be separated by commas. (optional)
    pulpIdIn := []string{"Inner_example"} // []string | Multiple values may be separated by commas. (optional)
    repositoryVersion := "repositoryVersion_example" // string | Repository Version referenced by HREF (optional)
    repositoryVersionAdded := "repositoryVersionAdded_example" // string | Repository Version referenced by HREF (optional)
    repositoryVersionRemoved := "repositoryVersionRemoved_example" // string | Repository Version referenced by HREF (optional)
    fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
    excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContentCollectionDeprecationsAPI.ContentAnsibleCollectionDeprecationsList(context.Background()).Limit(limit).Offset(offset).Ordering(ordering).PulpHrefIn(pulpHrefIn).PulpIdIn(pulpIdIn).RepositoryVersion(repositoryVersion).RepositoryVersionAdded(repositoryVersionAdded).RepositoryVersionRemoved(repositoryVersionRemoved).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContentCollectionDeprecationsAPI.ContentAnsibleCollectionDeprecationsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContentAnsibleCollectionDeprecationsList`: PaginatedansibleCollectionResponseList
    fmt.Fprintf(os.Stdout, "Response from `ContentCollectionDeprecationsAPI.ContentAnsibleCollectionDeprecationsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiContentAnsibleCollectionDeprecationsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Number of results to return per page. | 
 **offset** | **int32** | The initial index from which to return the results. | 
 **ordering** | **[]string** | Ordering  * &#x60;pk&#x60; - Pk * &#x60;-pk&#x60; - Pk (descending) | 
 **pulpHrefIn** | **[]string** | Multiple values may be separated by commas. | 
 **pulpIdIn** | **[]string** | Multiple values may be separated by commas. | 
 **repositoryVersion** | **string** | Repository Version referenced by HREF | 
 **repositoryVersionAdded** | **string** | Repository Version referenced by HREF | 
 **repositoryVersionRemoved** | **string** | Repository Version referenced by HREF | 
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


## ContentAnsibleCollectionDeprecationsRead

> AnsibleCollectionResponse ContentAnsibleCollectionDeprecationsRead(ctx, ansibleAnsibleCollectionDeprecatedHref).Fields(fields).ExcludeFields(excludeFields).Execute()

Inspect an ansible collection deprecated



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
    ansibleAnsibleCollectionDeprecatedHref := "ansibleAnsibleCollectionDeprecatedHref_example" // string | 
    fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
    excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContentCollectionDeprecationsAPI.ContentAnsibleCollectionDeprecationsRead(context.Background(), ansibleAnsibleCollectionDeprecatedHref).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContentCollectionDeprecationsAPI.ContentAnsibleCollectionDeprecationsRead``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContentAnsibleCollectionDeprecationsRead`: AnsibleCollectionResponse
    fmt.Fprintf(os.Stdout, "Response from `ContentCollectionDeprecationsAPI.ContentAnsibleCollectionDeprecationsRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ansibleAnsibleCollectionDeprecatedHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiContentAnsibleCollectionDeprecationsReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

### Return type

[**AnsibleCollectionResponse**](AnsibleCollectionResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

