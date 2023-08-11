# \ContentGemAPI

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ContentGemGemCreate**](ContentGemAPI.md#ContentGemGemCreate) | **Post** /pulp/api/v3/content/gem/gem/ | Create a gem content
[**ContentGemGemList**](ContentGemAPI.md#ContentGemGemList) | **Get** /pulp/api/v3/content/gem/gem/ | List gem contents
[**ContentGemGemRead**](ContentGemAPI.md#ContentGemGemRead) | **Get** {gem_gem_content_href} | Inspect a gem content



## ContentGemGemCreate

> AsyncOperationResponse ContentGemGemCreate(ctx).Repository(repository).Artifact(artifact).File(file).Execute()

Create a gem content



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
    artifact := "artifact_example" // string | Artifact file representing the physical content (optional)
    file := os.NewFile(1234, "some_file") // *os.File | An uploaded file that should be turned into the artifact of the content unit. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContentGemAPI.ContentGemGemCreate(context.Background()).Repository(repository).Artifact(artifact).File(file).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContentGemAPI.ContentGemGemCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContentGemGemCreate`: AsyncOperationResponse
    fmt.Fprintf(os.Stdout, "Response from `ContentGemAPI.ContentGemGemCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiContentGemGemCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **repository** | **string** | A URI of a repository the new content unit should be associated with. | 
 **artifact** | **string** | Artifact file representing the physical content | 
 **file** | ***os.File** | An uploaded file that should be turned into the artifact of the content unit. | 

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


## ContentGemGemList

> PaginatedgemGemContentResponseList ContentGemGemList(ctx).Checksum(checksum).Limit(limit).Name(name).Offset(offset).Ordering(ordering).Prerelease(prerelease).PulpHrefIn(pulpHrefIn).PulpIdIn(pulpIdIn).RepositoryVersion(repositoryVersion).RepositoryVersionAdded(repositoryVersionAdded).RepositoryVersionRemoved(repositoryVersionRemoved).Version(version).Fields(fields).ExcludeFields(excludeFields).Execute()

List gem contents



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
    checksum := "checksum_example" // string | Filter results where checksum matches value (optional)
    limit := int32(56) // int32 | Number of results to return per page. (optional)
    name := "name_example" // string | Filter results where name matches value (optional)
    offset := int32(56) // int32 | The initial index from which to return the results. (optional)
    ordering := []string{"Ordering_example"} // []string | Ordering  * `pulp_id` - Pulp id * `-pulp_id` - Pulp id (descending) * `pulp_created` - Pulp created * `-pulp_created` - Pulp created (descending) * `pulp_last_updated` - Pulp last updated * `-pulp_last_updated` - Pulp last updated (descending) * `pulp_type` - Pulp type * `-pulp_type` - Pulp type (descending) * `upstream_id` - Upstream id * `-upstream_id` - Upstream id (descending) * `timestamp_of_interest` - Timestamp of interest * `-timestamp_of_interest` - Timestamp of interest (descending) * `name` - Name * `-name` - Name (descending) * `version` - Version * `-version` - Version (descending) * `checksum` - Checksum * `-checksum` - Checksum (descending) * `prerelease` - Prerelease * `-prerelease` - Prerelease (descending) * `dependencies` - Dependencies * `-dependencies` - Dependencies (descending) * `required_ruby_version` - Required ruby version * `-required_ruby_version` - Required ruby version (descending) * `required_rubygems_version` - Required rubygems version * `-required_rubygems_version` - Required rubygems version (descending) * `pk` - Pk * `-pk` - Pk (descending) (optional)
    prerelease := true // bool | Filter results where prerelease matches value (optional)
    pulpHrefIn := []string{"Inner_example"} // []string | Multiple values may be separated by commas. (optional)
    pulpIdIn := []string{"Inner_example"} // []string | Multiple values may be separated by commas. (optional)
    repositoryVersion := "repositoryVersion_example" // string | Repository Version referenced by HREF (optional)
    repositoryVersionAdded := "repositoryVersionAdded_example" // string | Repository Version referenced by HREF (optional)
    repositoryVersionRemoved := "repositoryVersionRemoved_example" // string | Repository Version referenced by HREF (optional)
    version := "version_example" // string | Filter results where version matches value (optional)
    fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
    excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContentGemAPI.ContentGemGemList(context.Background()).Checksum(checksum).Limit(limit).Name(name).Offset(offset).Ordering(ordering).Prerelease(prerelease).PulpHrefIn(pulpHrefIn).PulpIdIn(pulpIdIn).RepositoryVersion(repositoryVersion).RepositoryVersionAdded(repositoryVersionAdded).RepositoryVersionRemoved(repositoryVersionRemoved).Version(version).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContentGemAPI.ContentGemGemList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContentGemGemList`: PaginatedgemGemContentResponseList
    fmt.Fprintf(os.Stdout, "Response from `ContentGemAPI.ContentGemGemList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiContentGemGemListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **checksum** | **string** | Filter results where checksum matches value | 
 **limit** | **int32** | Number of results to return per page. | 
 **name** | **string** | Filter results where name matches value | 
 **offset** | **int32** | The initial index from which to return the results. | 
 **ordering** | **[]string** | Ordering  * &#x60;pulp_id&#x60; - Pulp id * &#x60;-pulp_id&#x60; - Pulp id (descending) * &#x60;pulp_created&#x60; - Pulp created * &#x60;-pulp_created&#x60; - Pulp created (descending) * &#x60;pulp_last_updated&#x60; - Pulp last updated * &#x60;-pulp_last_updated&#x60; - Pulp last updated (descending) * &#x60;pulp_type&#x60; - Pulp type * &#x60;-pulp_type&#x60; - Pulp type (descending) * &#x60;upstream_id&#x60; - Upstream id * &#x60;-upstream_id&#x60; - Upstream id (descending) * &#x60;timestamp_of_interest&#x60; - Timestamp of interest * &#x60;-timestamp_of_interest&#x60; - Timestamp of interest (descending) * &#x60;name&#x60; - Name * &#x60;-name&#x60; - Name (descending) * &#x60;version&#x60; - Version * &#x60;-version&#x60; - Version (descending) * &#x60;checksum&#x60; - Checksum * &#x60;-checksum&#x60; - Checksum (descending) * &#x60;prerelease&#x60; - Prerelease * &#x60;-prerelease&#x60; - Prerelease (descending) * &#x60;dependencies&#x60; - Dependencies * &#x60;-dependencies&#x60; - Dependencies (descending) * &#x60;required_ruby_version&#x60; - Required ruby version * &#x60;-required_ruby_version&#x60; - Required ruby version (descending) * &#x60;required_rubygems_version&#x60; - Required rubygems version * &#x60;-required_rubygems_version&#x60; - Required rubygems version (descending) * &#x60;pk&#x60; - Pk * &#x60;-pk&#x60; - Pk (descending) | 
 **prerelease** | **bool** | Filter results where prerelease matches value | 
 **pulpHrefIn** | **[]string** | Multiple values may be separated by commas. | 
 **pulpIdIn** | **[]string** | Multiple values may be separated by commas. | 
 **repositoryVersion** | **string** | Repository Version referenced by HREF | 
 **repositoryVersionAdded** | **string** | Repository Version referenced by HREF | 
 **repositoryVersionRemoved** | **string** | Repository Version referenced by HREF | 
 **version** | **string** | Filter results where version matches value | 
 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

### Return type

[**PaginatedgemGemContentResponseList**](PaginatedgemGemContentResponseList.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContentGemGemRead

> GemGemContentResponse ContentGemGemRead(ctx, gemGemContentHref).Fields(fields).ExcludeFields(excludeFields).Execute()

Inspect a gem content



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
    gemGemContentHref := "gemGemContentHref_example" // string | 
    fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
    excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContentGemAPI.ContentGemGemRead(context.Background(), gemGemContentHref).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContentGemAPI.ContentGemGemRead``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContentGemGemRead`: GemGemContentResponse
    fmt.Fprintf(os.Stdout, "Response from `ContentGemAPI.ContentGemGemRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**gemGemContentHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiContentGemGemReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

### Return type

[**GemGemContentResponse**](GemGemContentResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

