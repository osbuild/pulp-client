# \RepositoriesFileVersionsAPI

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RepositoriesFileFileVersionsDelete**](RepositoriesFileVersionsAPI.md#RepositoriesFileFileVersionsDelete) | **Delete** {file_file_repository_version_href} | Delete a repository version
[**RepositoriesFileFileVersionsList**](RepositoriesFileVersionsAPI.md#RepositoriesFileFileVersionsList) | **Get** {file_file_repository_href}versions/ | List repository versions
[**RepositoriesFileFileVersionsRead**](RepositoriesFileVersionsAPI.md#RepositoriesFileFileVersionsRead) | **Get** {file_file_repository_version_href} | Inspect a repository version
[**RepositoriesFileFileVersionsRepair**](RepositoriesFileVersionsAPI.md#RepositoriesFileFileVersionsRepair) | **Post** {file_file_repository_version_href}repair/ | 



## RepositoriesFileFileVersionsDelete

> AsyncOperationResponse RepositoriesFileFileVersionsDelete(ctx, fileFileRepositoryVersionHref).Execute()

Delete a repository version



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
    fileFileRepositoryVersionHref := "fileFileRepositoryVersionHref_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RepositoriesFileVersionsAPI.RepositoriesFileFileVersionsDelete(context.Background(), fileFileRepositoryVersionHref).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesFileVersionsAPI.RepositoriesFileFileVersionsDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RepositoriesFileFileVersionsDelete`: AsyncOperationResponse
    fmt.Fprintf(os.Stdout, "Response from `RepositoriesFileVersionsAPI.RepositoriesFileFileVersionsDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fileFileRepositoryVersionHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesFileFileVersionsDeleteRequest struct via the builder pattern


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


## RepositoriesFileFileVersionsList

> PaginatedRepositoryVersionResponseList RepositoriesFileFileVersionsList(ctx, fileFileRepositoryHref).Content(content).ContentIn(contentIn).Limit(limit).Number(number).NumberGt(numberGt).NumberGte(numberGte).NumberLt(numberLt).NumberLte(numberLte).NumberRange(numberRange).Offset(offset).Ordering(ordering).PulpCreated(pulpCreated).PulpCreatedGt(pulpCreatedGt).PulpCreatedGte(pulpCreatedGte).PulpCreatedLt(pulpCreatedLt).PulpCreatedLte(pulpCreatedLte).PulpCreatedRange(pulpCreatedRange).PulpHrefIn(pulpHrefIn).Fields(fields).ExcludeFields(excludeFields).Execute()

List repository versions



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
    fileFileRepositoryHref := "fileFileRepositoryHref_example" // string | 
    content := "content_example" // string | Content Unit referenced by HREF (optional)
    contentIn := "contentIn_example" // string | Content Unit referenced by HREF (optional)
    limit := int32(56) // int32 | Number of results to return per page. (optional)
    number := int32(56) // int32 | Filter results where number matches value (optional)
    numberGt := int32(56) // int32 | Filter results where number is greater than value (optional)
    numberGte := int32(56) // int32 | Filter results where number is greater than or equal to value (optional)
    numberLt := int32(56) // int32 | Filter results where number is less than value (optional)
    numberLte := int32(56) // int32 | Filter results where number is less than or equal to value (optional)
    numberRange := []int32{int32(123)} // []int32 | Filter results where number is between two comma separated values (optional)
    offset := int32(56) // int32 | The initial index from which to return the results. (optional)
    ordering := []string{"Ordering_example"} // []string | Ordering  * `pulp_id` - Pulp id * `-pulp_id` - Pulp id (descending) * `pulp_created` - Pulp created * `-pulp_created` - Pulp created (descending) * `pulp_last_updated` - Pulp last updated * `-pulp_last_updated` - Pulp last updated (descending) * `number` - Number * `-number` - Number (descending) * `complete` - Complete * `-complete` - Complete (descending) * `info` - Info * `-info` - Info (descending) * `pk` - Pk * `-pk` - Pk (descending) (optional)
    pulpCreated := time.Now() // time.Time | Filter results where pulp_created matches value (optional)
    pulpCreatedGt := time.Now() // time.Time | Filter results where pulp_created is greater than value (optional)
    pulpCreatedGte := time.Now() // time.Time | Filter results where pulp_created is greater than or equal to value (optional)
    pulpCreatedLt := time.Now() // time.Time | Filter results where pulp_created is less than value (optional)
    pulpCreatedLte := time.Now() // time.Time | Filter results where pulp_created is less than or equal to value (optional)
    pulpCreatedRange := []time.Time{time.Now()} // []time.Time | Filter results where pulp_created is between two comma separated values (optional)
    pulpHrefIn := []string{"Inner_example"} // []string | Multiple values may be separated by commas. (optional)
    fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
    excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RepositoriesFileVersionsAPI.RepositoriesFileFileVersionsList(context.Background(), fileFileRepositoryHref).Content(content).ContentIn(contentIn).Limit(limit).Number(number).NumberGt(numberGt).NumberGte(numberGte).NumberLt(numberLt).NumberLte(numberLte).NumberRange(numberRange).Offset(offset).Ordering(ordering).PulpCreated(pulpCreated).PulpCreatedGt(pulpCreatedGt).PulpCreatedGte(pulpCreatedGte).PulpCreatedLt(pulpCreatedLt).PulpCreatedLte(pulpCreatedLte).PulpCreatedRange(pulpCreatedRange).PulpHrefIn(pulpHrefIn).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesFileVersionsAPI.RepositoriesFileFileVersionsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RepositoriesFileFileVersionsList`: PaginatedRepositoryVersionResponseList
    fmt.Fprintf(os.Stdout, "Response from `RepositoriesFileVersionsAPI.RepositoriesFileFileVersionsList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fileFileRepositoryHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesFileFileVersionsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **content** | **string** | Content Unit referenced by HREF | 
 **contentIn** | **string** | Content Unit referenced by HREF | 
 **limit** | **int32** | Number of results to return per page. | 
 **number** | **int32** | Filter results where number matches value | 
 **numberGt** | **int32** | Filter results where number is greater than value | 
 **numberGte** | **int32** | Filter results where number is greater than or equal to value | 
 **numberLt** | **int32** | Filter results where number is less than value | 
 **numberLte** | **int32** | Filter results where number is less than or equal to value | 
 **numberRange** | **[]int32** | Filter results where number is between two comma separated values | 
 **offset** | **int32** | The initial index from which to return the results. | 
 **ordering** | **[]string** | Ordering  * &#x60;pulp_id&#x60; - Pulp id * &#x60;-pulp_id&#x60; - Pulp id (descending) * &#x60;pulp_created&#x60; - Pulp created * &#x60;-pulp_created&#x60; - Pulp created (descending) * &#x60;pulp_last_updated&#x60; - Pulp last updated * &#x60;-pulp_last_updated&#x60; - Pulp last updated (descending) * &#x60;number&#x60; - Number * &#x60;-number&#x60; - Number (descending) * &#x60;complete&#x60; - Complete * &#x60;-complete&#x60; - Complete (descending) * &#x60;info&#x60; - Info * &#x60;-info&#x60; - Info (descending) * &#x60;pk&#x60; - Pk * &#x60;-pk&#x60; - Pk (descending) | 
 **pulpCreated** | **time.Time** | Filter results where pulp_created matches value | 
 **pulpCreatedGt** | **time.Time** | Filter results where pulp_created is greater than value | 
 **pulpCreatedGte** | **time.Time** | Filter results where pulp_created is greater than or equal to value | 
 **pulpCreatedLt** | **time.Time** | Filter results where pulp_created is less than value | 
 **pulpCreatedLte** | **time.Time** | Filter results where pulp_created is less than or equal to value | 
 **pulpCreatedRange** | [**[]time.Time**](time.Time.md) | Filter results where pulp_created is between two comma separated values | 
 **pulpHrefIn** | **[]string** | Multiple values may be separated by commas. | 
 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

### Return type

[**PaginatedRepositoryVersionResponseList**](PaginatedRepositoryVersionResponseList.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RepositoriesFileFileVersionsRead

> RepositoryVersionResponse RepositoriesFileFileVersionsRead(ctx, fileFileRepositoryVersionHref).Fields(fields).ExcludeFields(excludeFields).Execute()

Inspect a repository version



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
    fileFileRepositoryVersionHref := "fileFileRepositoryVersionHref_example" // string | 
    fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
    excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RepositoriesFileVersionsAPI.RepositoriesFileFileVersionsRead(context.Background(), fileFileRepositoryVersionHref).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesFileVersionsAPI.RepositoriesFileFileVersionsRead``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RepositoriesFileFileVersionsRead`: RepositoryVersionResponse
    fmt.Fprintf(os.Stdout, "Response from `RepositoriesFileVersionsAPI.RepositoriesFileFileVersionsRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fileFileRepositoryVersionHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesFileFileVersionsReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

### Return type

[**RepositoryVersionResponse**](RepositoryVersionResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RepositoriesFileFileVersionsRepair

> AsyncOperationResponse RepositoriesFileFileVersionsRepair(ctx, fileFileRepositoryVersionHref).Repair(repair).Execute()





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
    fileFileRepositoryVersionHref := "fileFileRepositoryVersionHref_example" // string | 
    repair := *openapiclient.NewRepair() // Repair | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RepositoriesFileVersionsAPI.RepositoriesFileFileVersionsRepair(context.Background(), fileFileRepositoryVersionHref).Repair(repair).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesFileVersionsAPI.RepositoriesFileFileVersionsRepair``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RepositoriesFileFileVersionsRepair`: AsyncOperationResponse
    fmt.Fprintf(os.Stdout, "Response from `RepositoriesFileVersionsAPI.RepositoriesFileFileVersionsRepair`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fileFileRepositoryVersionHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesFileFileVersionsRepairRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **repair** | [**Repair**](Repair.md) |  | 

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

