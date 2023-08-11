# \ContentArtifactAPI

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ContentMavenArtifactCreate**](ContentArtifactAPI.md#ContentMavenArtifactCreate) | **Post** /pulp/api/v3/content/maven/artifact/ | Create a maven artifact
[**ContentMavenArtifactList**](ContentArtifactAPI.md#ContentMavenArtifactList) | **Get** /pulp/api/v3/content/maven/artifact/ | List maven artifacts
[**ContentMavenArtifactRead**](ContentArtifactAPI.md#ContentMavenArtifactRead) | **Get** {maven_maven_artifact_href} | Inspect a maven artifact



## ContentMavenArtifactCreate

> MavenMavenArtifactResponse ContentMavenArtifactCreate(ctx).MavenMavenArtifact(mavenMavenArtifact).Execute()

Create a maven artifact



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
    mavenMavenArtifact := *openapiclient.NewMavenMavenArtifact("Artifact_example", "RelativePath_example") // MavenMavenArtifact | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContentArtifactAPI.ContentMavenArtifactCreate(context.Background()).MavenMavenArtifact(mavenMavenArtifact).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContentArtifactAPI.ContentMavenArtifactCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContentMavenArtifactCreate`: MavenMavenArtifactResponse
    fmt.Fprintf(os.Stdout, "Response from `ContentArtifactAPI.ContentMavenArtifactCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiContentMavenArtifactCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **mavenMavenArtifact** | [**MavenMavenArtifact**](MavenMavenArtifact.md) |  | 

### Return type

[**MavenMavenArtifactResponse**](MavenMavenArtifactResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContentMavenArtifactList

> PaginatedmavenMavenArtifactResponseList ContentMavenArtifactList(ctx).ArtifactId(artifactId).Filename(filename).GroupId(groupId).Limit(limit).Offset(offset).Ordering(ordering).PulpHrefIn(pulpHrefIn).PulpIdIn(pulpIdIn).RepositoryVersion(repositoryVersion).RepositoryVersionAdded(repositoryVersionAdded).RepositoryVersionRemoved(repositoryVersionRemoved).Version(version).Fields(fields).ExcludeFields(excludeFields).Execute()

List maven artifacts



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
    artifactId := "artifactId_example" // string | Filter results where artifact_id matches value (optional)
    filename := "filename_example" // string | Filter results where filename matches value (optional)
    groupId := "groupId_example" // string | Filter results where group_id matches value (optional)
    limit := int32(56) // int32 | Number of results to return per page. (optional)
    offset := int32(56) // int32 | The initial index from which to return the results. (optional)
    ordering := []string{"Ordering_example"} // []string | Ordering  * `pulp_id` - Pulp id * `-pulp_id` - Pulp id (descending) * `pulp_created` - Pulp created * `-pulp_created` - Pulp created (descending) * `pulp_last_updated` - Pulp last updated * `-pulp_last_updated` - Pulp last updated (descending) * `pulp_type` - Pulp type * `-pulp_type` - Pulp type (descending) * `upstream_id` - Upstream id * `-upstream_id` - Upstream id (descending) * `timestamp_of_interest` - Timestamp of interest * `-timestamp_of_interest` - Timestamp of interest (descending) * `group_id` - Group id * `-group_id` - Group id (descending) * `artifact_id` - Artifact id * `-artifact_id` - Artifact id (descending) * `version` - Version * `-version` - Version (descending) * `filename` - Filename * `-filename` - Filename (descending) * `pk` - Pk * `-pk` - Pk (descending) (optional)
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
    resp, r, err := apiClient.ContentArtifactAPI.ContentMavenArtifactList(context.Background()).ArtifactId(artifactId).Filename(filename).GroupId(groupId).Limit(limit).Offset(offset).Ordering(ordering).PulpHrefIn(pulpHrefIn).PulpIdIn(pulpIdIn).RepositoryVersion(repositoryVersion).RepositoryVersionAdded(repositoryVersionAdded).RepositoryVersionRemoved(repositoryVersionRemoved).Version(version).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContentArtifactAPI.ContentMavenArtifactList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContentMavenArtifactList`: PaginatedmavenMavenArtifactResponseList
    fmt.Fprintf(os.Stdout, "Response from `ContentArtifactAPI.ContentMavenArtifactList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiContentMavenArtifactListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **artifactId** | **string** | Filter results where artifact_id matches value | 
 **filename** | **string** | Filter results where filename matches value | 
 **groupId** | **string** | Filter results where group_id matches value | 
 **limit** | **int32** | Number of results to return per page. | 
 **offset** | **int32** | The initial index from which to return the results. | 
 **ordering** | **[]string** | Ordering  * &#x60;pulp_id&#x60; - Pulp id * &#x60;-pulp_id&#x60; - Pulp id (descending) * &#x60;pulp_created&#x60; - Pulp created * &#x60;-pulp_created&#x60; - Pulp created (descending) * &#x60;pulp_last_updated&#x60; - Pulp last updated * &#x60;-pulp_last_updated&#x60; - Pulp last updated (descending) * &#x60;pulp_type&#x60; - Pulp type * &#x60;-pulp_type&#x60; - Pulp type (descending) * &#x60;upstream_id&#x60; - Upstream id * &#x60;-upstream_id&#x60; - Upstream id (descending) * &#x60;timestamp_of_interest&#x60; - Timestamp of interest * &#x60;-timestamp_of_interest&#x60; - Timestamp of interest (descending) * &#x60;group_id&#x60; - Group id * &#x60;-group_id&#x60; - Group id (descending) * &#x60;artifact_id&#x60; - Artifact id * &#x60;-artifact_id&#x60; - Artifact id (descending) * &#x60;version&#x60; - Version * &#x60;-version&#x60; - Version (descending) * &#x60;filename&#x60; - Filename * &#x60;-filename&#x60; - Filename (descending) * &#x60;pk&#x60; - Pk * &#x60;-pk&#x60; - Pk (descending) | 
 **pulpHrefIn** | **[]string** | Multiple values may be separated by commas. | 
 **pulpIdIn** | **[]string** | Multiple values may be separated by commas. | 
 **repositoryVersion** | **string** | Repository Version referenced by HREF | 
 **repositoryVersionAdded** | **string** | Repository Version referenced by HREF | 
 **repositoryVersionRemoved** | **string** | Repository Version referenced by HREF | 
 **version** | **string** | Filter results where version matches value | 
 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

### Return type

[**PaginatedmavenMavenArtifactResponseList**](PaginatedmavenMavenArtifactResponseList.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContentMavenArtifactRead

> MavenMavenArtifactResponse ContentMavenArtifactRead(ctx, mavenMavenArtifactHref).Fields(fields).ExcludeFields(excludeFields).Execute()

Inspect a maven artifact



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
    mavenMavenArtifactHref := "mavenMavenArtifactHref_example" // string | 
    fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
    excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContentArtifactAPI.ContentMavenArtifactRead(context.Background(), mavenMavenArtifactHref).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContentArtifactAPI.ContentMavenArtifactRead``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContentMavenArtifactRead`: MavenMavenArtifactResponse
    fmt.Fprintf(os.Stdout, "Response from `ContentArtifactAPI.ContentMavenArtifactRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mavenMavenArtifactHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiContentMavenArtifactReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

### Return type

[**MavenMavenArtifactResponse**](MavenMavenArtifactResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

