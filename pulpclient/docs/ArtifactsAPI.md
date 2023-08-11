# \ArtifactsAPI

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ArtifactsCreate**](ArtifactsAPI.md#ArtifactsCreate) | **Post** /pulp/api/v3/artifacts/ | Create an artifact
[**ArtifactsDelete**](ArtifactsAPI.md#ArtifactsDelete) | **Delete** {artifact_href} | Delete an artifact
[**ArtifactsList**](ArtifactsAPI.md#ArtifactsList) | **Get** /pulp/api/v3/artifacts/ | List artifacts
[**ArtifactsRead**](ArtifactsAPI.md#ArtifactsRead) | **Get** {artifact_href} | Inspect an artifact



## ArtifactsCreate

> ArtifactResponse ArtifactsCreate(ctx).File(file).Size(size).Md5(md5).Sha1(sha1).Sha224(sha224).Sha256(sha256).Sha384(sha384).Sha512(sha512).Execute()

Create an artifact



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
    file := os.NewFile(1234, "some_file") // *os.File | The stored file.
    size := int64(789) // int64 | The size of the file in bytes. (optional)
    md5 := "md5_example" // string | The MD5 checksum of the file if available. (optional)
    sha1 := "sha1_example" // string | The SHA-1 checksum of the file if available. (optional)
    sha224 := "sha224_example" // string | The SHA-224 checksum of the file if available. (optional)
    sha256 := "sha256_example" // string | The SHA-256 checksum of the file if available. (optional)
    sha384 := "sha384_example" // string | The SHA-384 checksum of the file if available. (optional)
    sha512 := "sha512_example" // string | The SHA-512 checksum of the file if available. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ArtifactsAPI.ArtifactsCreate(context.Background()).File(file).Size(size).Md5(md5).Sha1(sha1).Sha224(sha224).Sha256(sha256).Sha384(sha384).Sha512(sha512).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ArtifactsAPI.ArtifactsCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ArtifactsCreate`: ArtifactResponse
    fmt.Fprintf(os.Stdout, "Response from `ArtifactsAPI.ArtifactsCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiArtifactsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **file** | ***os.File** | The stored file. | 
 **size** | **int64** | The size of the file in bytes. | 
 **md5** | **string** | The MD5 checksum of the file if available. | 
 **sha1** | **string** | The SHA-1 checksum of the file if available. | 
 **sha224** | **string** | The SHA-224 checksum of the file if available. | 
 **sha256** | **string** | The SHA-256 checksum of the file if available. | 
 **sha384** | **string** | The SHA-384 checksum of the file if available. | 
 **sha512** | **string** | The SHA-512 checksum of the file if available. | 

### Return type

[**ArtifactResponse**](ArtifactResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: multipart/form-data, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ArtifactsDelete

> ArtifactsDelete(ctx, artifactHref).Execute()

Delete an artifact



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
    artifactHref := "artifactHref_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ArtifactsAPI.ArtifactsDelete(context.Background(), artifactHref).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ArtifactsAPI.ArtifactsDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**artifactHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiArtifactsDeleteRequest struct via the builder pattern


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


## ArtifactsList

> PaginatedArtifactResponseList ArtifactsList(ctx).Limit(limit).Md5(md5).Offset(offset).Ordering(ordering).PulpHrefIn(pulpHrefIn).PulpIdIn(pulpIdIn).RepositoryVersion(repositoryVersion).Sha1(sha1).Sha224(sha224).Sha256(sha256).Sha384(sha384).Sha512(sha512).Fields(fields).ExcludeFields(excludeFields).Execute()

List artifacts



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
    md5 := "md5_example" // string | Filter results where md5 matches value (optional)
    offset := int32(56) // int32 | The initial index from which to return the results. (optional)
    ordering := []string{"Ordering_example"} // []string | Ordering  * `pulp_id` - Pulp id * `-pulp_id` - Pulp id (descending) * `pulp_created` - Pulp created * `-pulp_created` - Pulp created (descending) * `pulp_last_updated` - Pulp last updated * `-pulp_last_updated` - Pulp last updated (descending) * `file` - File * `-file` - File (descending) * `size` - Size * `-size` - Size (descending) * `md5` - Md5 * `-md5` - Md5 (descending) * `sha1` - Sha1 * `-sha1` - Sha1 (descending) * `sha224` - Sha224 * `-sha224` - Sha224 (descending) * `sha256` - Sha256 * `-sha256` - Sha256 (descending) * `sha384` - Sha384 * `-sha384` - Sha384 (descending) * `sha512` - Sha512 * `-sha512` - Sha512 (descending) * `timestamp_of_interest` - Timestamp of interest * `-timestamp_of_interest` - Timestamp of interest (descending) * `pk` - Pk * `-pk` - Pk (descending) (optional)
    pulpHrefIn := []string{"Inner_example"} // []string | Multiple values may be separated by commas. (optional)
    pulpIdIn := []string{"Inner_example"} // []string | Multiple values may be separated by commas. (optional)
    repositoryVersion := "repositoryVersion_example" // string | Repository Version referenced by HREF (optional)
    sha1 := "sha1_example" // string | Filter results where sha1 matches value (optional)
    sha224 := "sha224_example" // string | Filter results where sha224 matches value (optional)
    sha256 := "sha256_example" // string | Filter results where sha256 matches value (optional)
    sha384 := "sha384_example" // string | Filter results where sha384 matches value (optional)
    sha512 := "sha512_example" // string | Filter results where sha512 matches value (optional)
    fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
    excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ArtifactsAPI.ArtifactsList(context.Background()).Limit(limit).Md5(md5).Offset(offset).Ordering(ordering).PulpHrefIn(pulpHrefIn).PulpIdIn(pulpIdIn).RepositoryVersion(repositoryVersion).Sha1(sha1).Sha224(sha224).Sha256(sha256).Sha384(sha384).Sha512(sha512).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ArtifactsAPI.ArtifactsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ArtifactsList`: PaginatedArtifactResponseList
    fmt.Fprintf(os.Stdout, "Response from `ArtifactsAPI.ArtifactsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiArtifactsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Number of results to return per page. | 
 **md5** | **string** | Filter results where md5 matches value | 
 **offset** | **int32** | The initial index from which to return the results. | 
 **ordering** | **[]string** | Ordering  * &#x60;pulp_id&#x60; - Pulp id * &#x60;-pulp_id&#x60; - Pulp id (descending) * &#x60;pulp_created&#x60; - Pulp created * &#x60;-pulp_created&#x60; - Pulp created (descending) * &#x60;pulp_last_updated&#x60; - Pulp last updated * &#x60;-pulp_last_updated&#x60; - Pulp last updated (descending) * &#x60;file&#x60; - File * &#x60;-file&#x60; - File (descending) * &#x60;size&#x60; - Size * &#x60;-size&#x60; - Size (descending) * &#x60;md5&#x60; - Md5 * &#x60;-md5&#x60; - Md5 (descending) * &#x60;sha1&#x60; - Sha1 * &#x60;-sha1&#x60; - Sha1 (descending) * &#x60;sha224&#x60; - Sha224 * &#x60;-sha224&#x60; - Sha224 (descending) * &#x60;sha256&#x60; - Sha256 * &#x60;-sha256&#x60; - Sha256 (descending) * &#x60;sha384&#x60; - Sha384 * &#x60;-sha384&#x60; - Sha384 (descending) * &#x60;sha512&#x60; - Sha512 * &#x60;-sha512&#x60; - Sha512 (descending) * &#x60;timestamp_of_interest&#x60; - Timestamp of interest * &#x60;-timestamp_of_interest&#x60; - Timestamp of interest (descending) * &#x60;pk&#x60; - Pk * &#x60;-pk&#x60; - Pk (descending) | 
 **pulpHrefIn** | **[]string** | Multiple values may be separated by commas. | 
 **pulpIdIn** | **[]string** | Multiple values may be separated by commas. | 
 **repositoryVersion** | **string** | Repository Version referenced by HREF | 
 **sha1** | **string** | Filter results where sha1 matches value | 
 **sha224** | **string** | Filter results where sha224 matches value | 
 **sha256** | **string** | Filter results where sha256 matches value | 
 **sha384** | **string** | Filter results where sha384 matches value | 
 **sha512** | **string** | Filter results where sha512 matches value | 
 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

### Return type

[**PaginatedArtifactResponseList**](PaginatedArtifactResponseList.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ArtifactsRead

> ArtifactResponse ArtifactsRead(ctx, artifactHref).Fields(fields).ExcludeFields(excludeFields).Execute()

Inspect an artifact



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
    artifactHref := "artifactHref_example" // string | 
    fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
    excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ArtifactsAPI.ArtifactsRead(context.Background(), artifactHref).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ArtifactsAPI.ArtifactsRead``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ArtifactsRead`: ArtifactResponse
    fmt.Fprintf(os.Stdout, "Response from `ArtifactsAPI.ArtifactsRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**artifactHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiArtifactsReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

### Return type

[**ArtifactResponse**](ArtifactResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

