# \PublicationsVerbatimAPI

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PublicationsDebVerbatimCreate**](PublicationsVerbatimAPI.md#PublicationsDebVerbatimCreate) | **Post** /pulp/api/v3/publications/deb/verbatim/ | Create a verbatim publication
[**PublicationsDebVerbatimDelete**](PublicationsVerbatimAPI.md#PublicationsDebVerbatimDelete) | **Delete** {deb_verbatim_publication_href} | Delete a verbatim publication
[**PublicationsDebVerbatimList**](PublicationsVerbatimAPI.md#PublicationsDebVerbatimList) | **Get** /pulp/api/v3/publications/deb/verbatim/ | List verbatim publications
[**PublicationsDebVerbatimRead**](PublicationsVerbatimAPI.md#PublicationsDebVerbatimRead) | **Get** {deb_verbatim_publication_href} | Inspect a verbatim publication



## PublicationsDebVerbatimCreate

> AsyncOperationResponse PublicationsDebVerbatimCreate(ctx).DebVerbatimPublication(debVerbatimPublication).Execute()

Create a verbatim publication



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
    debVerbatimPublication := *openapiclient.NewDebVerbatimPublication() // DebVerbatimPublication | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PublicationsVerbatimAPI.PublicationsDebVerbatimCreate(context.Background()).DebVerbatimPublication(debVerbatimPublication).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PublicationsVerbatimAPI.PublicationsDebVerbatimCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PublicationsDebVerbatimCreate`: AsyncOperationResponse
    fmt.Fprintf(os.Stdout, "Response from `PublicationsVerbatimAPI.PublicationsDebVerbatimCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPublicationsDebVerbatimCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **debVerbatimPublication** | [**DebVerbatimPublication**](DebVerbatimPublication.md) |  | 

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


## PublicationsDebVerbatimDelete

> PublicationsDebVerbatimDelete(ctx, debVerbatimPublicationHref).Execute()

Delete a verbatim publication



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
    debVerbatimPublicationHref := "debVerbatimPublicationHref_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PublicationsVerbatimAPI.PublicationsDebVerbatimDelete(context.Background(), debVerbatimPublicationHref).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PublicationsVerbatimAPI.PublicationsDebVerbatimDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**debVerbatimPublicationHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPublicationsDebVerbatimDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PublicationsDebVerbatimList

> PaginateddebVerbatimPublicationResponseList PublicationsDebVerbatimList(ctx).Content(content).ContentIn(contentIn).Limit(limit).Offset(offset).Ordering(ordering).PulpCreated(pulpCreated).PulpCreatedGt(pulpCreatedGt).PulpCreatedGte(pulpCreatedGte).PulpCreatedLt(pulpCreatedLt).PulpCreatedLte(pulpCreatedLte).PulpCreatedRange(pulpCreatedRange).PulpHrefIn(pulpHrefIn).PulpIdIn(pulpIdIn).Repository(repository).RepositoryVersion(repositoryVersion).Fields(fields).ExcludeFields(excludeFields).Execute()

List verbatim publications



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
    content := "content_example" // string | Content Unit referenced by HREF (optional)
    contentIn := "contentIn_example" // string | Content Unit referenced by HREF (optional)
    limit := int32(56) // int32 | Number of results to return per page. (optional)
    offset := int32(56) // int32 | The initial index from which to return the results. (optional)
    ordering := []string{"Ordering_example"} // []string | Ordering  * `pulp_id` - Pulp id * `-pulp_id` - Pulp id (descending) * `pulp_created` - Pulp created * `-pulp_created` - Pulp created (descending) * `pulp_last_updated` - Pulp last updated * `-pulp_last_updated` - Pulp last updated (descending) * `pulp_type` - Pulp type * `-pulp_type` - Pulp type (descending) * `complete` - Complete * `-complete` - Complete (descending) * `pass_through` - Pass through * `-pass_through` - Pass through (descending) * `pk` - Pk * `-pk` - Pk (descending) (optional)
    pulpCreated := time.Now() // time.Time | Filter results where pulp_created matches value (optional)
    pulpCreatedGt := time.Now() // time.Time | Filter results where pulp_created is greater than value (optional)
    pulpCreatedGte := time.Now() // time.Time | Filter results where pulp_created is greater than or equal to value (optional)
    pulpCreatedLt := time.Now() // time.Time | Filter results where pulp_created is less than value (optional)
    pulpCreatedLte := time.Now() // time.Time | Filter results where pulp_created is less than or equal to value (optional)
    pulpCreatedRange := []time.Time{time.Now()} // []time.Time | Filter results where pulp_created is between two comma separated values (optional)
    pulpHrefIn := []string{"Inner_example"} // []string | Multiple values may be separated by commas. (optional)
    pulpIdIn := []string{"Inner_example"} // []string | Multiple values may be separated by commas. (optional)
    repository := "repository_example" // string | Repository referenced by HREF (optional)
    repositoryVersion := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Repository Version referenced by HREF (optional)
    fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
    excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PublicationsVerbatimAPI.PublicationsDebVerbatimList(context.Background()).Content(content).ContentIn(contentIn).Limit(limit).Offset(offset).Ordering(ordering).PulpCreated(pulpCreated).PulpCreatedGt(pulpCreatedGt).PulpCreatedGte(pulpCreatedGte).PulpCreatedLt(pulpCreatedLt).PulpCreatedLte(pulpCreatedLte).PulpCreatedRange(pulpCreatedRange).PulpHrefIn(pulpHrefIn).PulpIdIn(pulpIdIn).Repository(repository).RepositoryVersion(repositoryVersion).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PublicationsVerbatimAPI.PublicationsDebVerbatimList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PublicationsDebVerbatimList`: PaginateddebVerbatimPublicationResponseList
    fmt.Fprintf(os.Stdout, "Response from `PublicationsVerbatimAPI.PublicationsDebVerbatimList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPublicationsDebVerbatimListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **content** | **string** | Content Unit referenced by HREF | 
 **contentIn** | **string** | Content Unit referenced by HREF | 
 **limit** | **int32** | Number of results to return per page. | 
 **offset** | **int32** | The initial index from which to return the results. | 
 **ordering** | **[]string** | Ordering  * &#x60;pulp_id&#x60; - Pulp id * &#x60;-pulp_id&#x60; - Pulp id (descending) * &#x60;pulp_created&#x60; - Pulp created * &#x60;-pulp_created&#x60; - Pulp created (descending) * &#x60;pulp_last_updated&#x60; - Pulp last updated * &#x60;-pulp_last_updated&#x60; - Pulp last updated (descending) * &#x60;pulp_type&#x60; - Pulp type * &#x60;-pulp_type&#x60; - Pulp type (descending) * &#x60;complete&#x60; - Complete * &#x60;-complete&#x60; - Complete (descending) * &#x60;pass_through&#x60; - Pass through * &#x60;-pass_through&#x60; - Pass through (descending) * &#x60;pk&#x60; - Pk * &#x60;-pk&#x60; - Pk (descending) | 
 **pulpCreated** | **time.Time** | Filter results where pulp_created matches value | 
 **pulpCreatedGt** | **time.Time** | Filter results where pulp_created is greater than value | 
 **pulpCreatedGte** | **time.Time** | Filter results where pulp_created is greater than or equal to value | 
 **pulpCreatedLt** | **time.Time** | Filter results where pulp_created is less than value | 
 **pulpCreatedLte** | **time.Time** | Filter results where pulp_created is less than or equal to value | 
 **pulpCreatedRange** | [**[]time.Time**](time.Time.md) | Filter results where pulp_created is between two comma separated values | 
 **pulpHrefIn** | **[]string** | Multiple values may be separated by commas. | 
 **pulpIdIn** | **[]string** | Multiple values may be separated by commas. | 
 **repository** | **string** | Repository referenced by HREF | 
 **repositoryVersion** | **string** | Repository Version referenced by HREF | 
 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

### Return type

[**PaginateddebVerbatimPublicationResponseList**](PaginateddebVerbatimPublicationResponseList.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PublicationsDebVerbatimRead

> DebVerbatimPublicationResponse PublicationsDebVerbatimRead(ctx, debVerbatimPublicationHref).Fields(fields).ExcludeFields(excludeFields).Execute()

Inspect a verbatim publication



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
    debVerbatimPublicationHref := "debVerbatimPublicationHref_example" // string | 
    fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
    excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PublicationsVerbatimAPI.PublicationsDebVerbatimRead(context.Background(), debVerbatimPublicationHref).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PublicationsVerbatimAPI.PublicationsDebVerbatimRead``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PublicationsDebVerbatimRead`: DebVerbatimPublicationResponse
    fmt.Fprintf(os.Stdout, "Response from `PublicationsVerbatimAPI.PublicationsDebVerbatimRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**debVerbatimPublicationHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPublicationsDebVerbatimReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

### Return type

[**DebVerbatimPublicationResponse**](DebVerbatimPublicationResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

