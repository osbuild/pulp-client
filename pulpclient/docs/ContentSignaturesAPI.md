# \ContentSignaturesAPI

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ContentContainerSignaturesList**](ContentSignaturesAPI.md#ContentContainerSignaturesList) | **Get** /pulp/api/v3/content/container/signatures/ | List manifest signatures
[**ContentContainerSignaturesRead**](ContentSignaturesAPI.md#ContentContainerSignaturesRead) | **Get** {container_manifest_signature_href} | Inspect a manifest signature



## ContentContainerSignaturesList

> PaginatedcontainerManifestSignatureResponseList ContentContainerSignaturesList(ctx).Digest(digest).DigestIn(digestIn).KeyId(keyId).KeyIdIn(keyIdIn).Limit(limit).Manifest(manifest).Name(name).NameContains(nameContains).NameIcontains(nameIcontains).NameIn(nameIn).NameStartswith(nameStartswith).Offset(offset).Ordering(ordering).PulpHrefIn(pulpHrefIn).PulpIdIn(pulpIdIn).RepositoryVersion(repositoryVersion).RepositoryVersionAdded(repositoryVersionAdded).RepositoryVersionRemoved(repositoryVersionRemoved).Fields(fields).ExcludeFields(excludeFields).Execute()

List manifest signatures



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
    keyId := "keyId_example" // string | Filter results where key_id matches value (optional)
    keyIdIn := []string{"Inner_example"} // []string | Filter results where key_id is in a comma-separated list of values (optional)
    limit := int32(56) // int32 | Number of results to return per page. (optional)
    manifest := []string{"Inner_example"} // []string | Multiple values may be separated by commas. (optional)
    name := "name_example" // string | Filter results where name matches value (optional)
    nameContains := "nameContains_example" // string | Filter results where name contains value (optional)
    nameIcontains := "nameIcontains_example" // string | Filter results where name contains value (optional)
    nameIn := []string{"Inner_example"} // []string | Filter results where name is in a comma-separated list of values (optional)
    nameStartswith := "nameStartswith_example" // string | Filter results where name starts with value (optional)
    offset := int32(56) // int32 | The initial index from which to return the results. (optional)
    ordering := []string{"Ordering_example"} // []string | Ordering  * `pulp_id` - Pulp id * `-pulp_id` - Pulp id (descending) * `pulp_created` - Pulp created * `-pulp_created` - Pulp created (descending) * `pulp_last_updated` - Pulp last updated * `-pulp_last_updated` - Pulp last updated (descending) * `pulp_type` - Pulp type * `-pulp_type` - Pulp type (descending) * `upstream_id` - Upstream id * `-upstream_id` - Upstream id (descending) * `timestamp_of_interest` - Timestamp of interest * `-timestamp_of_interest` - Timestamp of interest (descending) * `name` - Name * `-name` - Name (descending) * `digest` - Digest * `-digest` - Digest (descending) * `type` - Type * `-type` - Type (descending) * `key_id` - Key id * `-key_id` - Key id (descending) * `timestamp` - Timestamp * `-timestamp` - Timestamp (descending) * `creator` - Creator * `-creator` - Creator (descending) * `data` - Data * `-data` - Data (descending) * `pk` - Pk * `-pk` - Pk (descending) (optional)
    pulpHrefIn := []string{"Inner_example"} // []string | Multiple values may be separated by commas. (optional)
    pulpIdIn := []string{"Inner_example"} // []string | Multiple values may be separated by commas. (optional)
    repositoryVersion := "repositoryVersion_example" // string | Repository Version referenced by HREF (optional)
    repositoryVersionAdded := "repositoryVersionAdded_example" // string | Repository Version referenced by HREF (optional)
    repositoryVersionRemoved := "repositoryVersionRemoved_example" // string | Repository Version referenced by HREF (optional)
    fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
    excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContentSignaturesAPI.ContentContainerSignaturesList(context.Background()).Digest(digest).DigestIn(digestIn).KeyId(keyId).KeyIdIn(keyIdIn).Limit(limit).Manifest(manifest).Name(name).NameContains(nameContains).NameIcontains(nameIcontains).NameIn(nameIn).NameStartswith(nameStartswith).Offset(offset).Ordering(ordering).PulpHrefIn(pulpHrefIn).PulpIdIn(pulpIdIn).RepositoryVersion(repositoryVersion).RepositoryVersionAdded(repositoryVersionAdded).RepositoryVersionRemoved(repositoryVersionRemoved).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContentSignaturesAPI.ContentContainerSignaturesList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContentContainerSignaturesList`: PaginatedcontainerManifestSignatureResponseList
    fmt.Fprintf(os.Stdout, "Response from `ContentSignaturesAPI.ContentContainerSignaturesList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiContentContainerSignaturesListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **digest** | **string** | Filter results where digest matches value | 
 **digestIn** | **[]string** | Filter results where digest is in a comma-separated list of values | 
 **keyId** | **string** | Filter results where key_id matches value | 
 **keyIdIn** | **[]string** | Filter results where key_id is in a comma-separated list of values | 
 **limit** | **int32** | Number of results to return per page. | 
 **manifest** | **[]string** | Multiple values may be separated by commas. | 
 **name** | **string** | Filter results where name matches value | 
 **nameContains** | **string** | Filter results where name contains value | 
 **nameIcontains** | **string** | Filter results where name contains value | 
 **nameIn** | **[]string** | Filter results where name is in a comma-separated list of values | 
 **nameStartswith** | **string** | Filter results where name starts with value | 
 **offset** | **int32** | The initial index from which to return the results. | 
 **ordering** | **[]string** | Ordering  * &#x60;pulp_id&#x60; - Pulp id * &#x60;-pulp_id&#x60; - Pulp id (descending) * &#x60;pulp_created&#x60; - Pulp created * &#x60;-pulp_created&#x60; - Pulp created (descending) * &#x60;pulp_last_updated&#x60; - Pulp last updated * &#x60;-pulp_last_updated&#x60; - Pulp last updated (descending) * &#x60;pulp_type&#x60; - Pulp type * &#x60;-pulp_type&#x60; - Pulp type (descending) * &#x60;upstream_id&#x60; - Upstream id * &#x60;-upstream_id&#x60; - Upstream id (descending) * &#x60;timestamp_of_interest&#x60; - Timestamp of interest * &#x60;-timestamp_of_interest&#x60; - Timestamp of interest (descending) * &#x60;name&#x60; - Name * &#x60;-name&#x60; - Name (descending) * &#x60;digest&#x60; - Digest * &#x60;-digest&#x60; - Digest (descending) * &#x60;type&#x60; - Type * &#x60;-type&#x60; - Type (descending) * &#x60;key_id&#x60; - Key id * &#x60;-key_id&#x60; - Key id (descending) * &#x60;timestamp&#x60; - Timestamp * &#x60;-timestamp&#x60; - Timestamp (descending) * &#x60;creator&#x60; - Creator * &#x60;-creator&#x60; - Creator (descending) * &#x60;data&#x60; - Data * &#x60;-data&#x60; - Data (descending) * &#x60;pk&#x60; - Pk * &#x60;-pk&#x60; - Pk (descending) | 
 **pulpHrefIn** | **[]string** | Multiple values may be separated by commas. | 
 **pulpIdIn** | **[]string** | Multiple values may be separated by commas. | 
 **repositoryVersion** | **string** | Repository Version referenced by HREF | 
 **repositoryVersionAdded** | **string** | Repository Version referenced by HREF | 
 **repositoryVersionRemoved** | **string** | Repository Version referenced by HREF | 
 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

### Return type

[**PaginatedcontainerManifestSignatureResponseList**](PaginatedcontainerManifestSignatureResponseList.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContentContainerSignaturesRead

> ContainerManifestSignatureResponse ContentContainerSignaturesRead(ctx, containerManifestSignatureHref).Fields(fields).ExcludeFields(excludeFields).Execute()

Inspect a manifest signature



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
    containerManifestSignatureHref := "containerManifestSignatureHref_example" // string | 
    fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
    excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContentSignaturesAPI.ContentContainerSignaturesRead(context.Background(), containerManifestSignatureHref).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContentSignaturesAPI.ContentContainerSignaturesRead``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContentContainerSignaturesRead`: ContainerManifestSignatureResponse
    fmt.Fprintf(os.Stdout, "Response from `ContentSignaturesAPI.ContentContainerSignaturesRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**containerManifestSignatureHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiContentContainerSignaturesReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

### Return type

[**ContainerManifestSignatureResponse**](ContainerManifestSignatureResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

