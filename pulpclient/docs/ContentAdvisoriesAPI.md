# \ContentAdvisoriesAPI

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ContentRpmAdvisoriesCreate**](ContentAdvisoriesAPI.md#ContentRpmAdvisoriesCreate) | **Post** /pulp/api/v3/content/rpm/advisories/ | Create an update record
[**ContentRpmAdvisoriesList**](ContentAdvisoriesAPI.md#ContentRpmAdvisoriesList) | **Get** /pulp/api/v3/content/rpm/advisories/ | List update records
[**ContentRpmAdvisoriesRead**](ContentAdvisoriesAPI.md#ContentRpmAdvisoriesRead) | **Get** {rpm_update_record_href} | Inspect an update record



## ContentRpmAdvisoriesCreate

> AsyncOperationResponse ContentRpmAdvisoriesCreate(ctx).Repository(repository).File(file).Execute()

Create an update record



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
    repository := "repository_example" // string | A URI of a repository the new content unit should be associated with. (optional)
    file := os.NewFile(1234, "some_file") // *os.File | An uploaded file that may be turned into the artifact of the content unit. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContentAdvisoriesAPI.ContentRpmAdvisoriesCreate(context.Background()).Repository(repository).File(file).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContentAdvisoriesAPI.ContentRpmAdvisoriesCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContentRpmAdvisoriesCreate`: AsyncOperationResponse
    fmt.Fprintf(os.Stdout, "Response from `ContentAdvisoriesAPI.ContentRpmAdvisoriesCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiContentRpmAdvisoriesCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **repository** | **string** | A URI of a repository the new content unit should be associated with. | 
 **file** | ***os.File** | An uploaded file that may be turned into the artifact of the content unit. | 

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


## ContentRpmAdvisoriesList

> PaginatedrpmUpdateRecordResponseList ContentRpmAdvisoriesList(ctx).Id(id).IdIn(idIn).Limit(limit).Offset(offset).Ordering(ordering).PulpHrefIn(pulpHrefIn).PulpIdIn(pulpIdIn).RepositoryVersion(repositoryVersion).RepositoryVersionAdded(repositoryVersionAdded).RepositoryVersionRemoved(repositoryVersionRemoved).Severity(severity).SeverityIn(severityIn).SeverityNe(severityNe).Status(status).StatusIn(statusIn).StatusNe(statusNe).Type_(type_).TypeIn(typeIn).TypeNe(typeNe).Fields(fields).ExcludeFields(excludeFields).Execute()

List update records



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
    id := "id_example" // string | Filter results where id matches value (optional)
    idIn := []string{"Inner_example"} // []string | Filter results where id is in a comma-separated list of values (optional)
    limit := int32(56) // int32 | Number of results to return per page. (optional)
    offset := int32(56) // int32 | The initial index from which to return the results. (optional)
    ordering := []string{"Ordering_example"} // []string | Ordering  * `pulp_id` - Pulp id * `-pulp_id` - Pulp id (descending) * `pulp_created` - Pulp created * `-pulp_created` - Pulp created (descending) * `pulp_last_updated` - Pulp last updated * `-pulp_last_updated` - Pulp last updated (descending) * `pulp_type` - Pulp type * `-pulp_type` - Pulp type (descending) * `upstream_id` - Upstream id * `-upstream_id` - Upstream id (descending) * `timestamp_of_interest` - Timestamp of interest * `-timestamp_of_interest` - Timestamp of interest (descending) * `id` - Id * `-id` - Id (descending) * `updated_date` - Updated date * `-updated_date` - Updated date (descending) * `description` - Description * `-description` - Description (descending) * `issued_date` - Issued date * `-issued_date` - Issued date (descending) * `fromstr` - Fromstr * `-fromstr` - Fromstr (descending) * `status` - Status * `-status` - Status (descending) * `title` - Title * `-title` - Title (descending) * `summary` - Summary * `-summary` - Summary (descending) * `version` - Version * `-version` - Version (descending) * `type` - Type * `-type` - Type (descending) * `severity` - Severity * `-severity` - Severity (descending) * `solution` - Solution * `-solution` - Solution (descending) * `release` - Release * `-release` - Release (descending) * `rights` - Rights * `-rights` - Rights (descending) * `reboot_suggested` - Reboot suggested * `-reboot_suggested` - Reboot suggested (descending) * `pushcount` - Pushcount * `-pushcount` - Pushcount (descending) * `digest` - Digest * `-digest` - Digest (descending) * `pk` - Pk * `-pk` - Pk (descending) (optional)
    pulpHrefIn := []string{"Inner_example"} // []string | Multiple values may be separated by commas. (optional)
    pulpIdIn := []string{"Inner_example"} // []string | Multiple values may be separated by commas. (optional)
    repositoryVersion := "repositoryVersion_example" // string | Repository Version referenced by HREF (optional)
    repositoryVersionAdded := "repositoryVersionAdded_example" // string | Repository Version referenced by HREF (optional)
    repositoryVersionRemoved := "repositoryVersionRemoved_example" // string | Repository Version referenced by HREF (optional)
    severity := "severity_example" // string | Filter results where severity matches value (optional)
    severityIn := []string{"Inner_example"} // []string | Filter results where severity is in a comma-separated list of values (optional)
    severityNe := "severityNe_example" // string | Filter results where severity not equal to value (optional)
    status := "status_example" // string | Filter results where status matches value (optional)
    statusIn := []string{"Inner_example"} // []string | Filter results where status is in a comma-separated list of values (optional)
    statusNe := "statusNe_example" // string | Filter results where status not equal to value (optional)
    type_ := "type__example" // string | Filter results where type matches value (optional)
    typeIn := []string{"Inner_example"} // []string | Filter results where type is in a comma-separated list of values (optional)
    typeNe := "typeNe_example" // string | Filter results where type not equal to value (optional)
    fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
    excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContentAdvisoriesAPI.ContentRpmAdvisoriesList(context.Background()).Id(id).IdIn(idIn).Limit(limit).Offset(offset).Ordering(ordering).PulpHrefIn(pulpHrefIn).PulpIdIn(pulpIdIn).RepositoryVersion(repositoryVersion).RepositoryVersionAdded(repositoryVersionAdded).RepositoryVersionRemoved(repositoryVersionRemoved).Severity(severity).SeverityIn(severityIn).SeverityNe(severityNe).Status(status).StatusIn(statusIn).StatusNe(statusNe).Type_(type_).TypeIn(typeIn).TypeNe(typeNe).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContentAdvisoriesAPI.ContentRpmAdvisoriesList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContentRpmAdvisoriesList`: PaginatedrpmUpdateRecordResponseList
    fmt.Fprintf(os.Stdout, "Response from `ContentAdvisoriesAPI.ContentRpmAdvisoriesList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiContentRpmAdvisoriesListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** | Filter results where id matches value | 
 **idIn** | **[]string** | Filter results where id is in a comma-separated list of values | 
 **limit** | **int32** | Number of results to return per page. | 
 **offset** | **int32** | The initial index from which to return the results. | 
 **ordering** | **[]string** | Ordering  * &#x60;pulp_id&#x60; - Pulp id * &#x60;-pulp_id&#x60; - Pulp id (descending) * &#x60;pulp_created&#x60; - Pulp created * &#x60;-pulp_created&#x60; - Pulp created (descending) * &#x60;pulp_last_updated&#x60; - Pulp last updated * &#x60;-pulp_last_updated&#x60; - Pulp last updated (descending) * &#x60;pulp_type&#x60; - Pulp type * &#x60;-pulp_type&#x60; - Pulp type (descending) * &#x60;upstream_id&#x60; - Upstream id * &#x60;-upstream_id&#x60; - Upstream id (descending) * &#x60;timestamp_of_interest&#x60; - Timestamp of interest * &#x60;-timestamp_of_interest&#x60; - Timestamp of interest (descending) * &#x60;id&#x60; - Id * &#x60;-id&#x60; - Id (descending) * &#x60;updated_date&#x60; - Updated date * &#x60;-updated_date&#x60; - Updated date (descending) * &#x60;description&#x60; - Description * &#x60;-description&#x60; - Description (descending) * &#x60;issued_date&#x60; - Issued date * &#x60;-issued_date&#x60; - Issued date (descending) * &#x60;fromstr&#x60; - Fromstr * &#x60;-fromstr&#x60; - Fromstr (descending) * &#x60;status&#x60; - Status * &#x60;-status&#x60; - Status (descending) * &#x60;title&#x60; - Title * &#x60;-title&#x60; - Title (descending) * &#x60;summary&#x60; - Summary * &#x60;-summary&#x60; - Summary (descending) * &#x60;version&#x60; - Version * &#x60;-version&#x60; - Version (descending) * &#x60;type&#x60; - Type * &#x60;-type&#x60; - Type (descending) * &#x60;severity&#x60; - Severity * &#x60;-severity&#x60; - Severity (descending) * &#x60;solution&#x60; - Solution * &#x60;-solution&#x60; - Solution (descending) * &#x60;release&#x60; - Release * &#x60;-release&#x60; - Release (descending) * &#x60;rights&#x60; - Rights * &#x60;-rights&#x60; - Rights (descending) * &#x60;reboot_suggested&#x60; - Reboot suggested * &#x60;-reboot_suggested&#x60; - Reboot suggested (descending) * &#x60;pushcount&#x60; - Pushcount * &#x60;-pushcount&#x60; - Pushcount (descending) * &#x60;digest&#x60; - Digest * &#x60;-digest&#x60; - Digest (descending) * &#x60;pk&#x60; - Pk * &#x60;-pk&#x60; - Pk (descending) | 
 **pulpHrefIn** | **[]string** | Multiple values may be separated by commas. | 
 **pulpIdIn** | **[]string** | Multiple values may be separated by commas. | 
 **repositoryVersion** | **string** | Repository Version referenced by HREF | 
 **repositoryVersionAdded** | **string** | Repository Version referenced by HREF | 
 **repositoryVersionRemoved** | **string** | Repository Version referenced by HREF | 
 **severity** | **string** | Filter results where severity matches value | 
 **severityIn** | **[]string** | Filter results where severity is in a comma-separated list of values | 
 **severityNe** | **string** | Filter results where severity not equal to value | 
 **status** | **string** | Filter results where status matches value | 
 **statusIn** | **[]string** | Filter results where status is in a comma-separated list of values | 
 **statusNe** | **string** | Filter results where status not equal to value | 
 **type_** | **string** | Filter results where type matches value | 
 **typeIn** | **[]string** | Filter results where type is in a comma-separated list of values | 
 **typeNe** | **string** | Filter results where type not equal to value | 
 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

### Return type

[**PaginatedrpmUpdateRecordResponseList**](PaginatedrpmUpdateRecordResponseList.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContentRpmAdvisoriesRead

> RpmUpdateRecordResponse ContentRpmAdvisoriesRead(ctx, rpmUpdateRecordHref).Fields(fields).ExcludeFields(excludeFields).Execute()

Inspect an update record



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
    rpmUpdateRecordHref := "rpmUpdateRecordHref_example" // string | 
    fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
    excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContentAdvisoriesAPI.ContentRpmAdvisoriesRead(context.Background(), rpmUpdateRecordHref).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContentAdvisoriesAPI.ContentRpmAdvisoriesRead``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContentRpmAdvisoriesRead`: RpmUpdateRecordResponse
    fmt.Fprintf(os.Stdout, "Response from `ContentAdvisoriesAPI.ContentRpmAdvisoriesRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**rpmUpdateRecordHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiContentRpmAdvisoriesReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

### Return type

[**RpmUpdateRecordResponse**](RpmUpdateRecordResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

