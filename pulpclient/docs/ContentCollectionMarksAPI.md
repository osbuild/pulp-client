# \ContentCollectionMarksAPI

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ContentAnsibleCollectionMarksCreate**](ContentCollectionMarksAPI.md#ContentAnsibleCollectionMarksCreate) | **Post** /pulp/api/v3/content/ansible/collection_marks/ | Create a collection version mark
[**ContentAnsibleCollectionMarksList**](ContentCollectionMarksAPI.md#ContentAnsibleCollectionMarksList) | **Get** /pulp/api/v3/content/ansible/collection_marks/ | List collection version marks
[**ContentAnsibleCollectionMarksRead**](ContentCollectionMarksAPI.md#ContentAnsibleCollectionMarksRead) | **Get** {ansible_collection_version_mark_href} | Inspect a collection version mark



## ContentAnsibleCollectionMarksCreate

> AnsibleCollectionVersionMarkResponse ContentAnsibleCollectionMarksCreate(ctx).AnsibleCollectionVersionMark(ansibleCollectionVersionMark).Execute()

Create a collection version mark



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
    ansibleCollectionVersionMark := *openapiclient.NewAnsibleCollectionVersionMark("MarkedCollection_example", "Value_example") // AnsibleCollectionVersionMark | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContentCollectionMarksAPI.ContentAnsibleCollectionMarksCreate(context.Background()).AnsibleCollectionVersionMark(ansibleCollectionVersionMark).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContentCollectionMarksAPI.ContentAnsibleCollectionMarksCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContentAnsibleCollectionMarksCreate`: AnsibleCollectionVersionMarkResponse
    fmt.Fprintf(os.Stdout, "Response from `ContentCollectionMarksAPI.ContentAnsibleCollectionMarksCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiContentAnsibleCollectionMarksCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ansibleCollectionVersionMark** | [**AnsibleCollectionVersionMark**](AnsibleCollectionVersionMark.md) |  | 

### Return type

[**AnsibleCollectionVersionMarkResponse**](AnsibleCollectionVersionMarkResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContentAnsibleCollectionMarksList

> PaginatedansibleCollectionVersionMarkResponseList ContentAnsibleCollectionMarksList(ctx).Limit(limit).MarkedCollection(markedCollection).Offset(offset).Ordering(ordering).PulpHrefIn(pulpHrefIn).PulpIdIn(pulpIdIn).RepositoryVersion(repositoryVersion).RepositoryVersionAdded(repositoryVersionAdded).RepositoryVersionRemoved(repositoryVersionRemoved).Value(value).ValueIn(valueIn).Fields(fields).ExcludeFields(excludeFields).Execute()

List collection version marks



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
    markedCollection := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Filter marks for collection version (optional)
    offset := int32(56) // int32 | The initial index from which to return the results. (optional)
    ordering := []string{"Ordering_example"} // []string | Ordering  * `pulp_id` - Pulp id * `-pulp_id` - Pulp id (descending) * `pulp_created` - Pulp created * `-pulp_created` - Pulp created (descending) * `pulp_last_updated` - Pulp last updated * `-pulp_last_updated` - Pulp last updated (descending) * `pulp_type` - Pulp type * `-pulp_type` - Pulp type (descending) * `upstream_id` - Upstream id * `-upstream_id` - Upstream id (descending) * `timestamp_of_interest` - Timestamp of interest * `-timestamp_of_interest` - Timestamp of interest (descending) * `value` - Value * `-value` - Value (descending) * `pk` - Pk * `-pk` - Pk (descending) (optional)
    pulpHrefIn := []string{"Inner_example"} // []string | Multiple values may be separated by commas. (optional)
    pulpIdIn := []string{"Inner_example"} // []string | Multiple values may be separated by commas. (optional)
    repositoryVersion := "repositoryVersion_example" // string | Repository Version referenced by HREF (optional)
    repositoryVersionAdded := "repositoryVersionAdded_example" // string | Repository Version referenced by HREF (optional)
    repositoryVersionRemoved := "repositoryVersionRemoved_example" // string | Repository Version referenced by HREF (optional)
    value := "value_example" // string | Filter marks by value (optional)
    valueIn := []string{"Inner_example"} // []string | Filter results where value is in a comma-separated list of values (optional)
    fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
    excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContentCollectionMarksAPI.ContentAnsibleCollectionMarksList(context.Background()).Limit(limit).MarkedCollection(markedCollection).Offset(offset).Ordering(ordering).PulpHrefIn(pulpHrefIn).PulpIdIn(pulpIdIn).RepositoryVersion(repositoryVersion).RepositoryVersionAdded(repositoryVersionAdded).RepositoryVersionRemoved(repositoryVersionRemoved).Value(value).ValueIn(valueIn).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContentCollectionMarksAPI.ContentAnsibleCollectionMarksList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContentAnsibleCollectionMarksList`: PaginatedansibleCollectionVersionMarkResponseList
    fmt.Fprintf(os.Stdout, "Response from `ContentCollectionMarksAPI.ContentAnsibleCollectionMarksList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiContentAnsibleCollectionMarksListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Number of results to return per page. | 
 **markedCollection** | **string** | Filter marks for collection version | 
 **offset** | **int32** | The initial index from which to return the results. | 
 **ordering** | **[]string** | Ordering  * &#x60;pulp_id&#x60; - Pulp id * &#x60;-pulp_id&#x60; - Pulp id (descending) * &#x60;pulp_created&#x60; - Pulp created * &#x60;-pulp_created&#x60; - Pulp created (descending) * &#x60;pulp_last_updated&#x60; - Pulp last updated * &#x60;-pulp_last_updated&#x60; - Pulp last updated (descending) * &#x60;pulp_type&#x60; - Pulp type * &#x60;-pulp_type&#x60; - Pulp type (descending) * &#x60;upstream_id&#x60; - Upstream id * &#x60;-upstream_id&#x60; - Upstream id (descending) * &#x60;timestamp_of_interest&#x60; - Timestamp of interest * &#x60;-timestamp_of_interest&#x60; - Timestamp of interest (descending) * &#x60;value&#x60; - Value * &#x60;-value&#x60; - Value (descending) * &#x60;pk&#x60; - Pk * &#x60;-pk&#x60; - Pk (descending) | 
 **pulpHrefIn** | **[]string** | Multiple values may be separated by commas. | 
 **pulpIdIn** | **[]string** | Multiple values may be separated by commas. | 
 **repositoryVersion** | **string** | Repository Version referenced by HREF | 
 **repositoryVersionAdded** | **string** | Repository Version referenced by HREF | 
 **repositoryVersionRemoved** | **string** | Repository Version referenced by HREF | 
 **value** | **string** | Filter marks by value | 
 **valueIn** | **[]string** | Filter results where value is in a comma-separated list of values | 
 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

### Return type

[**PaginatedansibleCollectionVersionMarkResponseList**](PaginatedansibleCollectionVersionMarkResponseList.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContentAnsibleCollectionMarksRead

> AnsibleCollectionVersionMarkResponse ContentAnsibleCollectionMarksRead(ctx, ansibleCollectionVersionMarkHref).Fields(fields).ExcludeFields(excludeFields).Execute()

Inspect a collection version mark



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
    ansibleCollectionVersionMarkHref := "ansibleCollectionVersionMarkHref_example" // string | 
    fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
    excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContentCollectionMarksAPI.ContentAnsibleCollectionMarksRead(context.Background(), ansibleCollectionVersionMarkHref).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContentCollectionMarksAPI.ContentAnsibleCollectionMarksRead``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContentAnsibleCollectionMarksRead`: AnsibleCollectionVersionMarkResponse
    fmt.Fprintf(os.Stdout, "Response from `ContentCollectionMarksAPI.ContentAnsibleCollectionMarksRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ansibleCollectionVersionMarkHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiContentAnsibleCollectionMarksReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

### Return type

[**AnsibleCollectionVersionMarkResponse**](AnsibleCollectionVersionMarkResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

