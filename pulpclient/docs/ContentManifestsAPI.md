# \ContentManifestsAPI

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ContentContainerManifestsList**](ContentManifestsAPI.md#ContentContainerManifestsList) | **Get** /pulp/api/v3/content/container/manifests/ | List manifests
[**ContentContainerManifestsRead**](ContentManifestsAPI.md#ContentContainerManifestsRead) | **Get** {container_manifest_href} | Inspect a manifest



## ContentContainerManifestsList

> PaginatedcontainerManifestResponseList ContentContainerManifestsList(ctx).Digest(digest).DigestIn(digestIn).Limit(limit).MediaType(mediaType).Offset(offset).Ordering(ordering).PulpHrefIn(pulpHrefIn).PulpIdIn(pulpIdIn).RepositoryVersion(repositoryVersion).RepositoryVersionAdded(repositoryVersionAdded).RepositoryVersionRemoved(repositoryVersionRemoved).Fields(fields).ExcludeFields(excludeFields).Execute()

List manifests



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
    digest := "digest_example" // string | Filter results where digest matches value (optional)
    digestIn := []string{"Inner_example"} // []string | Filter results where digest is in a comma-separated list of values (optional)
    limit := int32(56) // int32 | Number of results to return per page. (optional)
    mediaType := []string{"MediaType_example"} // []string | * `application/vnd.docker.distribution.manifest.v1+json` - application/vnd.docker.distribution.manifest.v1+json * `application/vnd.docker.distribution.manifest.v2+json` - application/vnd.docker.distribution.manifest.v2+json * `application/vnd.docker.distribution.manifest.list.v2+json` - application/vnd.docker.distribution.manifest.list.v2+json * `application/vnd.oci.image.manifest.v1+json` - application/vnd.oci.image.manifest.v1+json * `application/vnd.oci.image.index.v1+json` - application/vnd.oci.image.index.v1+json (optional)
    offset := int32(56) // int32 | The initial index from which to return the results. (optional)
    ordering := []string{"Ordering_example"} // []string | Ordering  * `pulp_id` - Pulp id * `-pulp_id` - Pulp id (descending) * `pulp_created` - Pulp created * `-pulp_created` - Pulp created (descending) * `pulp_last_updated` - Pulp last updated * `-pulp_last_updated` - Pulp last updated (descending) * `pulp_type` - Pulp type * `-pulp_type` - Pulp type (descending) * `upstream_id` - Upstream id * `-upstream_id` - Upstream id (descending) * `timestamp_of_interest` - Timestamp of interest * `-timestamp_of_interest` - Timestamp of interest (descending) * `digest` - Digest * `-digest` - Digest (descending) * `schema_version` - Schema version * `-schema_version` - Schema version (descending) * `media_type` - Media type * `-media_type` - Media type (descending) * `pk` - Pk * `-pk` - Pk (descending) (optional)
    pulpHrefIn := []string{"Inner_example"} // []string | Multiple values may be separated by commas. (optional)
    pulpIdIn := []string{"Inner_example"} // []string | Multiple values may be separated by commas. (optional)
    repositoryVersion := "repositoryVersion_example" // string | Repository Version referenced by HREF (optional)
    repositoryVersionAdded := "repositoryVersionAdded_example" // string | Repository Version referenced by HREF (optional)
    repositoryVersionRemoved := "repositoryVersionRemoved_example" // string | Repository Version referenced by HREF (optional)
    fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
    excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContentManifestsAPI.ContentContainerManifestsList(context.Background()).Digest(digest).DigestIn(digestIn).Limit(limit).MediaType(mediaType).Offset(offset).Ordering(ordering).PulpHrefIn(pulpHrefIn).PulpIdIn(pulpIdIn).RepositoryVersion(repositoryVersion).RepositoryVersionAdded(repositoryVersionAdded).RepositoryVersionRemoved(repositoryVersionRemoved).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContentManifestsAPI.ContentContainerManifestsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContentContainerManifestsList`: PaginatedcontainerManifestResponseList
    fmt.Fprintf(os.Stdout, "Response from `ContentManifestsAPI.ContentContainerManifestsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiContentContainerManifestsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **digest** | **string** | Filter results where digest matches value | 
 **digestIn** | **[]string** | Filter results where digest is in a comma-separated list of values | 
 **limit** | **int32** | Number of results to return per page. | 
 **mediaType** | **[]string** | * &#x60;application/vnd.docker.distribution.manifest.v1+json&#x60; - application/vnd.docker.distribution.manifest.v1+json * &#x60;application/vnd.docker.distribution.manifest.v2+json&#x60; - application/vnd.docker.distribution.manifest.v2+json * &#x60;application/vnd.docker.distribution.manifest.list.v2+json&#x60; - application/vnd.docker.distribution.manifest.list.v2+json * &#x60;application/vnd.oci.image.manifest.v1+json&#x60; - application/vnd.oci.image.manifest.v1+json * &#x60;application/vnd.oci.image.index.v1+json&#x60; - application/vnd.oci.image.index.v1+json | 
 **offset** | **int32** | The initial index from which to return the results. | 
 **ordering** | **[]string** | Ordering  * &#x60;pulp_id&#x60; - Pulp id * &#x60;-pulp_id&#x60; - Pulp id (descending) * &#x60;pulp_created&#x60; - Pulp created * &#x60;-pulp_created&#x60; - Pulp created (descending) * &#x60;pulp_last_updated&#x60; - Pulp last updated * &#x60;-pulp_last_updated&#x60; - Pulp last updated (descending) * &#x60;pulp_type&#x60; - Pulp type * &#x60;-pulp_type&#x60; - Pulp type (descending) * &#x60;upstream_id&#x60; - Upstream id * &#x60;-upstream_id&#x60; - Upstream id (descending) * &#x60;timestamp_of_interest&#x60; - Timestamp of interest * &#x60;-timestamp_of_interest&#x60; - Timestamp of interest (descending) * &#x60;digest&#x60; - Digest * &#x60;-digest&#x60; - Digest (descending) * &#x60;schema_version&#x60; - Schema version * &#x60;-schema_version&#x60; - Schema version (descending) * &#x60;media_type&#x60; - Media type * &#x60;-media_type&#x60; - Media type (descending) * &#x60;pk&#x60; - Pk * &#x60;-pk&#x60; - Pk (descending) | 
 **pulpHrefIn** | **[]string** | Multiple values may be separated by commas. | 
 **pulpIdIn** | **[]string** | Multiple values may be separated by commas. | 
 **repositoryVersion** | **string** | Repository Version referenced by HREF | 
 **repositoryVersionAdded** | **string** | Repository Version referenced by HREF | 
 **repositoryVersionRemoved** | **string** | Repository Version referenced by HREF | 
 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

### Return type

[**PaginatedcontainerManifestResponseList**](PaginatedcontainerManifestResponseList.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContentContainerManifestsRead

> ContainerManifestResponse ContentContainerManifestsRead(ctx, containerManifestHref).Fields(fields).ExcludeFields(excludeFields).Execute()

Inspect a manifest



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
    containerManifestHref := "containerManifestHref_example" // string | 
    fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
    excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContentManifestsAPI.ContentContainerManifestsRead(context.Background(), containerManifestHref).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContentManifestsAPI.ContentContainerManifestsRead``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContentContainerManifestsRead`: ContainerManifestResponse
    fmt.Fprintf(os.Stdout, "Response from `ContentManifestsAPI.ContentContainerManifestsRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**containerManifestHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiContentContainerManifestsReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

### Return type

[**ContainerManifestResponse**](ContainerManifestResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

