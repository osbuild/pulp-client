# \PulpAnsibleDefaultApiV3CollectionsVersionsAPI

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PulpAnsibleGalaxyDefaultApiV3CollectionsVersionsDelete**](PulpAnsibleDefaultApiV3CollectionsVersionsAPI.md#PulpAnsibleGalaxyDefaultApiV3CollectionsVersionsDelete) | **Delete** /pulp_ansible/galaxy/default/api/v3/collections/{namespace}/{name}/versions/{version}/ | 
[**PulpAnsibleGalaxyDefaultApiV3CollectionsVersionsList**](PulpAnsibleDefaultApiV3CollectionsVersionsAPI.md#PulpAnsibleGalaxyDefaultApiV3CollectionsVersionsList) | **Get** /pulp_ansible/galaxy/default/api/v3/collections/{namespace}/{name}/versions/ | 
[**PulpAnsibleGalaxyDefaultApiV3CollectionsVersionsRead**](PulpAnsibleDefaultApiV3CollectionsVersionsAPI.md#PulpAnsibleGalaxyDefaultApiV3CollectionsVersionsRead) | **Get** /pulp_ansible/galaxy/default/api/v3/collections/{namespace}/{name}/versions/{version}/ | 



## PulpAnsibleGalaxyDefaultApiV3CollectionsVersionsDelete

> AsyncOperationResponse PulpAnsibleGalaxyDefaultApiV3CollectionsVersionsDelete(ctx, name, namespace, version).Execute()





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
    name := "name_example" // string | 
    namespace := "namespace_example" // string | 
    version := "version_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PulpAnsibleDefaultApiV3CollectionsVersionsAPI.PulpAnsibleGalaxyDefaultApiV3CollectionsVersionsDelete(context.Background(), name, namespace, version).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PulpAnsibleDefaultApiV3CollectionsVersionsAPI.PulpAnsibleGalaxyDefaultApiV3CollectionsVersionsDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PulpAnsibleGalaxyDefaultApiV3CollectionsVersionsDelete`: AsyncOperationResponse
    fmt.Fprintf(os.Stdout, "Response from `PulpAnsibleDefaultApiV3CollectionsVersionsAPI.PulpAnsibleGalaxyDefaultApiV3CollectionsVersionsDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 
**namespace** | **string** |  | 
**version** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPulpAnsibleGalaxyDefaultApiV3CollectionsVersionsDeleteRequest struct via the builder pattern


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


## PulpAnsibleGalaxyDefaultApiV3CollectionsVersionsList

> PaginatedCollectionVersionListResponseList PulpAnsibleGalaxyDefaultApiV3CollectionsVersionsList(ctx, name, namespace).IsHighest(isHighest).Limit(limit).Name2(name2).Namespace2(namespace2).Offset(offset).Ordering(ordering).PulpHrefIn(pulpHrefIn).PulpIdIn(pulpIdIn).Q(q).RepositoryVersion(repositoryVersion).RepositoryVersionAdded(repositoryVersionAdded).RepositoryVersionRemoved(repositoryVersionRemoved).Tags(tags).Version(version).Fields(fields).ExcludeFields(excludeFields).Execute()





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
    name := "name_example" // string | 
    namespace := "namespace_example" // string | 
    isHighest := true // bool |  (optional)
    limit := int32(56) // int32 | Number of results to return per page. (optional)
    name2 := "name_example" // string |  (optional)
    namespace2 := "namespace_example" // string |  (optional)
    offset := int32(56) // int32 | The initial index from which to return the results. (optional)
    ordering := []string{"Ordering_example"} // []string | Ordering  * `pulp_id` - Pulp id * `-pulp_id` - Pulp id (descending) * `pulp_created` - Pulp created * `-pulp_created` - Pulp created (descending) * `pulp_last_updated` - Pulp last updated * `-pulp_last_updated` - Pulp last updated (descending) * `pulp_type` - Pulp type * `-pulp_type` - Pulp type (descending) * `upstream_id` - Upstream id * `-upstream_id` - Upstream id (descending) * `timestamp_of_interest` - Timestamp of interest * `-timestamp_of_interest` - Timestamp of interest (descending) * `authors` - Authors * `-authors` - Authors (descending) * `contents` - Contents * `-contents` - Contents (descending) * `dependencies` - Dependencies * `-dependencies` - Dependencies (descending) * `description` - Description * `-description` - Description (descending) * `docs_blob` - Docs blob * `-docs_blob` - Docs blob (descending) * `manifest` - Manifest * `-manifest` - Manifest (descending) * `files` - Files * `-files` - Files (descending) * `documentation` - Documentation * `-documentation` - Documentation (descending) * `homepage` - Homepage * `-homepage` - Homepage (descending) * `issues` - Issues * `-issues` - Issues (descending) * `license` - License * `-license` - License (descending) * `name` - Name * `-name` - Name (descending) * `namespace` - Namespace * `-namespace` - Namespace (descending) * `repository` - Repository * `-repository` - Repository (descending) * `requires_ansible` - Requires ansible * `-requires_ansible` - Requires ansible (descending) * `version` - Version * `-version` - Version (descending) * `version_major` - Version major * `-version_major` - Version major (descending) * `version_minor` - Version minor * `-version_minor` - Version minor (descending) * `version_patch` - Version patch * `-version_patch` - Version patch (descending) * `version_prerelease` - Version prerelease * `-version_prerelease` - Version prerelease (descending) * `is_highest` - Is highest * `-is_highest` - Is highest (descending) * `search_vector` - Search vector * `-search_vector` - Search vector (descending) * `pk` - Pk * `-pk` - Pk (descending) (optional)
    pulpHrefIn := []string{"Inner_example"} // []string | Multiple values may be separated by commas. (optional)
    pulpIdIn := []string{"Inner_example"} // []string | Multiple values may be separated by commas. (optional)
    q := "q_example" // string |  (optional)
    repositoryVersion := "repositoryVersion_example" // string | Repository Version referenced by HREF (optional)
    repositoryVersionAdded := "repositoryVersionAdded_example" // string | Repository Version referenced by HREF (optional)
    repositoryVersionRemoved := "repositoryVersionRemoved_example" // string | Repository Version referenced by HREF (optional)
    tags := "tags_example" // string | Filter by comma separate list of tags that must all be matched (optional)
    version := "version_example" // string | Filter results where version matches value (optional)
    fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
    excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PulpAnsibleDefaultApiV3CollectionsVersionsAPI.PulpAnsibleGalaxyDefaultApiV3CollectionsVersionsList(context.Background(), name, namespace).IsHighest(isHighest).Limit(limit).Name2(name2).Namespace2(namespace2).Offset(offset).Ordering(ordering).PulpHrefIn(pulpHrefIn).PulpIdIn(pulpIdIn).Q(q).RepositoryVersion(repositoryVersion).RepositoryVersionAdded(repositoryVersionAdded).RepositoryVersionRemoved(repositoryVersionRemoved).Tags(tags).Version(version).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PulpAnsibleDefaultApiV3CollectionsVersionsAPI.PulpAnsibleGalaxyDefaultApiV3CollectionsVersionsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PulpAnsibleGalaxyDefaultApiV3CollectionsVersionsList`: PaginatedCollectionVersionListResponseList
    fmt.Fprintf(os.Stdout, "Response from `PulpAnsibleDefaultApiV3CollectionsVersionsAPI.PulpAnsibleGalaxyDefaultApiV3CollectionsVersionsList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 
**namespace** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPulpAnsibleGalaxyDefaultApiV3CollectionsVersionsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **isHighest** | **bool** |  | 
 **limit** | **int32** | Number of results to return per page. | 
 **name2** | **string** |  | 
 **namespace2** | **string** |  | 
 **offset** | **int32** | The initial index from which to return the results. | 
 **ordering** | **[]string** | Ordering  * &#x60;pulp_id&#x60; - Pulp id * &#x60;-pulp_id&#x60; - Pulp id (descending) * &#x60;pulp_created&#x60; - Pulp created * &#x60;-pulp_created&#x60; - Pulp created (descending) * &#x60;pulp_last_updated&#x60; - Pulp last updated * &#x60;-pulp_last_updated&#x60; - Pulp last updated (descending) * &#x60;pulp_type&#x60; - Pulp type * &#x60;-pulp_type&#x60; - Pulp type (descending) * &#x60;upstream_id&#x60; - Upstream id * &#x60;-upstream_id&#x60; - Upstream id (descending) * &#x60;timestamp_of_interest&#x60; - Timestamp of interest * &#x60;-timestamp_of_interest&#x60; - Timestamp of interest (descending) * &#x60;authors&#x60; - Authors * &#x60;-authors&#x60; - Authors (descending) * &#x60;contents&#x60; - Contents * &#x60;-contents&#x60; - Contents (descending) * &#x60;dependencies&#x60; - Dependencies * &#x60;-dependencies&#x60; - Dependencies (descending) * &#x60;description&#x60; - Description * &#x60;-description&#x60; - Description (descending) * &#x60;docs_blob&#x60; - Docs blob * &#x60;-docs_blob&#x60; - Docs blob (descending) * &#x60;manifest&#x60; - Manifest * &#x60;-manifest&#x60; - Manifest (descending) * &#x60;files&#x60; - Files * &#x60;-files&#x60; - Files (descending) * &#x60;documentation&#x60; - Documentation * &#x60;-documentation&#x60; - Documentation (descending) * &#x60;homepage&#x60; - Homepage * &#x60;-homepage&#x60; - Homepage (descending) * &#x60;issues&#x60; - Issues * &#x60;-issues&#x60; - Issues (descending) * &#x60;license&#x60; - License * &#x60;-license&#x60; - License (descending) * &#x60;name&#x60; - Name * &#x60;-name&#x60; - Name (descending) * &#x60;namespace&#x60; - Namespace * &#x60;-namespace&#x60; - Namespace (descending) * &#x60;repository&#x60; - Repository * &#x60;-repository&#x60; - Repository (descending) * &#x60;requires_ansible&#x60; - Requires ansible * &#x60;-requires_ansible&#x60; - Requires ansible (descending) * &#x60;version&#x60; - Version * &#x60;-version&#x60; - Version (descending) * &#x60;version_major&#x60; - Version major * &#x60;-version_major&#x60; - Version major (descending) * &#x60;version_minor&#x60; - Version minor * &#x60;-version_minor&#x60; - Version minor (descending) * &#x60;version_patch&#x60; - Version patch * &#x60;-version_patch&#x60; - Version patch (descending) * &#x60;version_prerelease&#x60; - Version prerelease * &#x60;-version_prerelease&#x60; - Version prerelease (descending) * &#x60;is_highest&#x60; - Is highest * &#x60;-is_highest&#x60; - Is highest (descending) * &#x60;search_vector&#x60; - Search vector * &#x60;-search_vector&#x60; - Search vector (descending) * &#x60;pk&#x60; - Pk * &#x60;-pk&#x60; - Pk (descending) | 
 **pulpHrefIn** | **[]string** | Multiple values may be separated by commas. | 
 **pulpIdIn** | **[]string** | Multiple values may be separated by commas. | 
 **q** | **string** |  | 
 **repositoryVersion** | **string** | Repository Version referenced by HREF | 
 **repositoryVersionAdded** | **string** | Repository Version referenced by HREF | 
 **repositoryVersionRemoved** | **string** | Repository Version referenced by HREF | 
 **tags** | **string** | Filter by comma separate list of tags that must all be matched | 
 **version** | **string** | Filter results where version matches value | 
 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

### Return type

[**PaginatedCollectionVersionListResponseList**](PaginatedCollectionVersionListResponseList.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PulpAnsibleGalaxyDefaultApiV3CollectionsVersionsRead

> CollectionVersionResponse PulpAnsibleGalaxyDefaultApiV3CollectionsVersionsRead(ctx, name, namespace, version).Fields(fields).ExcludeFields(excludeFields).Execute()





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
    name := "name_example" // string | 
    namespace := "namespace_example" // string | 
    version := "version_example" // string | 
    fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
    excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PulpAnsibleDefaultApiV3CollectionsVersionsAPI.PulpAnsibleGalaxyDefaultApiV3CollectionsVersionsRead(context.Background(), name, namespace, version).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PulpAnsibleDefaultApiV3CollectionsVersionsAPI.PulpAnsibleGalaxyDefaultApiV3CollectionsVersionsRead``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PulpAnsibleGalaxyDefaultApiV3CollectionsVersionsRead`: CollectionVersionResponse
    fmt.Fprintf(os.Stdout, "Response from `PulpAnsibleDefaultApiV3CollectionsVersionsAPI.PulpAnsibleGalaxyDefaultApiV3CollectionsVersionsRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 
**namespace** | **string** |  | 
**version** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPulpAnsibleGalaxyDefaultApiV3CollectionsVersionsReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

### Return type

[**CollectionVersionResponse**](CollectionVersionResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

