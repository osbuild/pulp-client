# \ContentCollectionSignaturesAPI

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ContentAnsibleCollectionSignaturesCreate**](ContentCollectionSignaturesAPI.md#ContentAnsibleCollectionSignaturesCreate) | **Post** /pulp/api/v3/content/ansible/collection_signatures/ | Create a collection version signature
[**ContentAnsibleCollectionSignaturesList**](ContentCollectionSignaturesAPI.md#ContentAnsibleCollectionSignaturesList) | **Get** /pulp/api/v3/content/ansible/collection_signatures/ | List collection version signatures
[**ContentAnsibleCollectionSignaturesRead**](ContentCollectionSignaturesAPI.md#ContentAnsibleCollectionSignaturesRead) | **Get** {ansible_collection_version_signature_href} | Inspect a collection version signature



## ContentAnsibleCollectionSignaturesCreate

> AsyncOperationResponse ContentAnsibleCollectionSignaturesCreate(ctx).File(file).SignedCollection(signedCollection).Repository(repository).Execute()

Create a collection version signature



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
    file := os.NewFile(1234, "some_file") // *os.File | An uploaded file that may be turned into the artifact of the content unit.
    signedCollection := "signedCollection_example" // string | The content this signature is pointing to.
    repository := "repository_example" // string | A URI of a repository the new content unit should be associated with. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContentCollectionSignaturesAPI.ContentAnsibleCollectionSignaturesCreate(context.Background()).File(file).SignedCollection(signedCollection).Repository(repository).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContentCollectionSignaturesAPI.ContentAnsibleCollectionSignaturesCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContentAnsibleCollectionSignaturesCreate`: AsyncOperationResponse
    fmt.Fprintf(os.Stdout, "Response from `ContentCollectionSignaturesAPI.ContentAnsibleCollectionSignaturesCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiContentAnsibleCollectionSignaturesCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **file** | ***os.File** | An uploaded file that may be turned into the artifact of the content unit. | 
 **signedCollection** | **string** | The content this signature is pointing to. | 
 **repository** | **string** | A URI of a repository the new content unit should be associated with. | 

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


## ContentAnsibleCollectionSignaturesList

> PaginatedansibleCollectionVersionSignatureResponseList ContentAnsibleCollectionSignaturesList(ctx).Limit(limit).Offset(offset).Ordering(ordering).PubkeyFingerprint(pubkeyFingerprint).PubkeyFingerprintIn(pubkeyFingerprintIn).PulpHrefIn(pulpHrefIn).PulpIdIn(pulpIdIn).RepositoryVersion(repositoryVersion).RepositoryVersionAdded(repositoryVersionAdded).RepositoryVersionRemoved(repositoryVersionRemoved).SignedCollection(signedCollection).SigningService(signingService).Fields(fields).ExcludeFields(excludeFields).Execute()

List collection version signatures



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
    ordering := []string{"Ordering_example"} // []string | Ordering  * `pulp_id` - Pulp id * `-pulp_id` - Pulp id (descending) * `pulp_created` - Pulp created * `-pulp_created` - Pulp created (descending) * `pulp_last_updated` - Pulp last updated * `-pulp_last_updated` - Pulp last updated (descending) * `pulp_type` - Pulp type * `-pulp_type` - Pulp type (descending) * `upstream_id` - Upstream id * `-upstream_id` - Upstream id (descending) * `timestamp_of_interest` - Timestamp of interest * `-timestamp_of_interest` - Timestamp of interest (descending) * `data` - Data * `-data` - Data (descending) * `digest` - Digest * `-digest` - Digest (descending) * `pubkey_fingerprint` - Pubkey fingerprint * `-pubkey_fingerprint` - Pubkey fingerprint (descending) * `pk` - Pk * `-pk` - Pk (descending) (optional)
    pubkeyFingerprint := "pubkeyFingerprint_example" // string | Filter results where pubkey_fingerprint matches value (optional)
    pubkeyFingerprintIn := []string{"Inner_example"} // []string | Filter results where pubkey_fingerprint is in a comma-separated list of values (optional)
    pulpHrefIn := []string{"Inner_example"} // []string | Multiple values may be separated by commas. (optional)
    pulpIdIn := []string{"Inner_example"} // []string | Multiple values may be separated by commas. (optional)
    repositoryVersion := "repositoryVersion_example" // string | Repository Version referenced by HREF (optional)
    repositoryVersionAdded := "repositoryVersionAdded_example" // string | Repository Version referenced by HREF (optional)
    repositoryVersionRemoved := "repositoryVersionRemoved_example" // string | Repository Version referenced by HREF (optional)
    signedCollection := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Filter signatures for collection version (optional)
    signingService := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Filter signatures produced by signature service (optional)
    fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
    excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContentCollectionSignaturesAPI.ContentAnsibleCollectionSignaturesList(context.Background()).Limit(limit).Offset(offset).Ordering(ordering).PubkeyFingerprint(pubkeyFingerprint).PubkeyFingerprintIn(pubkeyFingerprintIn).PulpHrefIn(pulpHrefIn).PulpIdIn(pulpIdIn).RepositoryVersion(repositoryVersion).RepositoryVersionAdded(repositoryVersionAdded).RepositoryVersionRemoved(repositoryVersionRemoved).SignedCollection(signedCollection).SigningService(signingService).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContentCollectionSignaturesAPI.ContentAnsibleCollectionSignaturesList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContentAnsibleCollectionSignaturesList`: PaginatedansibleCollectionVersionSignatureResponseList
    fmt.Fprintf(os.Stdout, "Response from `ContentCollectionSignaturesAPI.ContentAnsibleCollectionSignaturesList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiContentAnsibleCollectionSignaturesListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Number of results to return per page. | 
 **offset** | **int32** | The initial index from which to return the results. | 
 **ordering** | **[]string** | Ordering  * &#x60;pulp_id&#x60; - Pulp id * &#x60;-pulp_id&#x60; - Pulp id (descending) * &#x60;pulp_created&#x60; - Pulp created * &#x60;-pulp_created&#x60; - Pulp created (descending) * &#x60;pulp_last_updated&#x60; - Pulp last updated * &#x60;-pulp_last_updated&#x60; - Pulp last updated (descending) * &#x60;pulp_type&#x60; - Pulp type * &#x60;-pulp_type&#x60; - Pulp type (descending) * &#x60;upstream_id&#x60; - Upstream id * &#x60;-upstream_id&#x60; - Upstream id (descending) * &#x60;timestamp_of_interest&#x60; - Timestamp of interest * &#x60;-timestamp_of_interest&#x60; - Timestamp of interest (descending) * &#x60;data&#x60; - Data * &#x60;-data&#x60; - Data (descending) * &#x60;digest&#x60; - Digest * &#x60;-digest&#x60; - Digest (descending) * &#x60;pubkey_fingerprint&#x60; - Pubkey fingerprint * &#x60;-pubkey_fingerprint&#x60; - Pubkey fingerprint (descending) * &#x60;pk&#x60; - Pk * &#x60;-pk&#x60; - Pk (descending) | 
 **pubkeyFingerprint** | **string** | Filter results where pubkey_fingerprint matches value | 
 **pubkeyFingerprintIn** | **[]string** | Filter results where pubkey_fingerprint is in a comma-separated list of values | 
 **pulpHrefIn** | **[]string** | Multiple values may be separated by commas. | 
 **pulpIdIn** | **[]string** | Multiple values may be separated by commas. | 
 **repositoryVersion** | **string** | Repository Version referenced by HREF | 
 **repositoryVersionAdded** | **string** | Repository Version referenced by HREF | 
 **repositoryVersionRemoved** | **string** | Repository Version referenced by HREF | 
 **signedCollection** | **string** | Filter signatures for collection version | 
 **signingService** | **string** | Filter signatures produced by signature service | 
 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

### Return type

[**PaginatedansibleCollectionVersionSignatureResponseList**](PaginatedansibleCollectionVersionSignatureResponseList.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContentAnsibleCollectionSignaturesRead

> AnsibleCollectionVersionSignatureResponse ContentAnsibleCollectionSignaturesRead(ctx, ansibleCollectionVersionSignatureHref).Fields(fields).ExcludeFields(excludeFields).Execute()

Inspect a collection version signature



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
    ansibleCollectionVersionSignatureHref := "ansibleCollectionVersionSignatureHref_example" // string | 
    fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
    excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContentCollectionSignaturesAPI.ContentAnsibleCollectionSignaturesRead(context.Background(), ansibleCollectionVersionSignatureHref).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContentCollectionSignaturesAPI.ContentAnsibleCollectionSignaturesRead``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContentAnsibleCollectionSignaturesRead`: AnsibleCollectionVersionSignatureResponse
    fmt.Fprintf(os.Stdout, "Response from `ContentCollectionSignaturesAPI.ContentAnsibleCollectionSignaturesRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ansibleCollectionVersionSignatureHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiContentAnsibleCollectionSignaturesReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

### Return type

[**AnsibleCollectionVersionSignatureResponse**](AnsibleCollectionVersionSignatureResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

