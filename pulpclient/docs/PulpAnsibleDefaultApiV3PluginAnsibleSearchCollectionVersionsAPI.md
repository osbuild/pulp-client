# \PulpAnsibleDefaultApiV3PluginAnsibleSearchCollectionVersionsAPI

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PulpAnsibleGalaxyDefaultApiV3PluginAnsibleSearchCollectionVersionsList**](PulpAnsibleDefaultApiV3PluginAnsibleSearchCollectionVersionsAPI.md#PulpAnsibleGalaxyDefaultApiV3PluginAnsibleSearchCollectionVersionsList) | **Get** /pulp_ansible/galaxy/default/api/v3/plugin/ansible/search/collection-versions/ | 
[**PulpAnsibleGalaxyDefaultApiV3PluginAnsibleSearchCollectionVersionsRebuild**](PulpAnsibleDefaultApiV3PluginAnsibleSearchCollectionVersionsAPI.md#PulpAnsibleGalaxyDefaultApiV3PluginAnsibleSearchCollectionVersionsRebuild) | **Post** /pulp_ansible/galaxy/default/api/v3/plugin/ansible/search/collection-versions/ | 



## PulpAnsibleGalaxyDefaultApiV3PluginAnsibleSearchCollectionVersionsList

> PaginatedCollectionVersionSearchListResponseList PulpAnsibleGalaxyDefaultApiV3PluginAnsibleSearchCollectionVersionsList(ctx).Dependency(dependency).Deprecated(deprecated).Distribution(distribution).DistributionBasePath(distributionBasePath).Highest(highest).IsDeprecated(isDeprecated).IsHighest(isHighest).IsSigned(isSigned).Keywords(keywords).Limit(limit).Name(name).Namespace(namespace).Offset(offset).OrderBy(orderBy).Q(q).Repository(repository).RepositoryLabel(repositoryLabel).RepositoryName(repositoryName).RepositoryVersion(repositoryVersion).Signed(signed).Tags(tags).Version(version).VersionRange(versionRange).Fields(fields).ExcludeFields(excludeFields).Execute()





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
    dependency := "dependency_example" // string |  (optional)
    deprecated := true // bool |  (optional)
    distribution := []string{"Inner_example"} // []string | Filter collectionversions that are in these distrubtion ids. (optional)
    distributionBasePath := []string{"Inner_example"} // []string | Filter collectionversions that are in these base paths. (optional)
    highest := true // bool |  (optional)
    isDeprecated := true // bool |  (optional)
    isHighest := true // bool |  (optional)
    isSigned := true // bool |  (optional)
    keywords := "keywords_example" // string |  (optional)
    limit := int32(56) // int32 | Number of results to return per page. (optional)
    name := "name_example" // string |  (optional)
    namespace := "namespace_example" // string |  (optional)
    offset := int32(56) // int32 | The initial index from which to return the results. (optional)
    orderBy := []string{"OrderBy_example"} // []string | Ordering  * `pulp_created` - by CV created * `-pulp_created` - by CV created (descending) * `namespace` - by CV namespace * `-namespace` - by CV namespace (descending) * `name` - by CV name * `-name` - by CV name (descending) * `version` - by CV version * `-version` - by CV version (descending) (optional)
    q := "q_example" // string |  (optional)
    repository := []string{"Inner_example"} // []string | Filter collectionversions that are in these repository ids. (optional)
    repositoryLabel := "repositoryLabel_example" // string | Filter labels by search string (optional)
    repositoryName := []string{"Inner_example"} // []string | Filter collectionversions that are in these repositories. (optional)
    repositoryVersion := "repositoryVersion_example" // string |  (optional)
    signed := true // bool |  (optional)
    tags := "tags_example" // string | Filter by comma separate list of tags that must all be matched (optional)
    version := "version_example" // string |  (optional)
    versionRange := "versionRange_example" // string |  (optional)
    fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
    excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PulpAnsibleDefaultApiV3PluginAnsibleSearchCollectionVersionsAPI.PulpAnsibleGalaxyDefaultApiV3PluginAnsibleSearchCollectionVersionsList(context.Background()).Dependency(dependency).Deprecated(deprecated).Distribution(distribution).DistributionBasePath(distributionBasePath).Highest(highest).IsDeprecated(isDeprecated).IsHighest(isHighest).IsSigned(isSigned).Keywords(keywords).Limit(limit).Name(name).Namespace(namespace).Offset(offset).OrderBy(orderBy).Q(q).Repository(repository).RepositoryLabel(repositoryLabel).RepositoryName(repositoryName).RepositoryVersion(repositoryVersion).Signed(signed).Tags(tags).Version(version).VersionRange(versionRange).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PulpAnsibleDefaultApiV3PluginAnsibleSearchCollectionVersionsAPI.PulpAnsibleGalaxyDefaultApiV3PluginAnsibleSearchCollectionVersionsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PulpAnsibleGalaxyDefaultApiV3PluginAnsibleSearchCollectionVersionsList`: PaginatedCollectionVersionSearchListResponseList
    fmt.Fprintf(os.Stdout, "Response from `PulpAnsibleDefaultApiV3PluginAnsibleSearchCollectionVersionsAPI.PulpAnsibleGalaxyDefaultApiV3PluginAnsibleSearchCollectionVersionsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPulpAnsibleGalaxyDefaultApiV3PluginAnsibleSearchCollectionVersionsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dependency** | **string** |  | 
 **deprecated** | **bool** |  | 
 **distribution** | **[]string** | Filter collectionversions that are in these distrubtion ids. | 
 **distributionBasePath** | **[]string** | Filter collectionversions that are in these base paths. | 
 **highest** | **bool** |  | 
 **isDeprecated** | **bool** |  | 
 **isHighest** | **bool** |  | 
 **isSigned** | **bool** |  | 
 **keywords** | **string** |  | 
 **limit** | **int32** | Number of results to return per page. | 
 **name** | **string** |  | 
 **namespace** | **string** |  | 
 **offset** | **int32** | The initial index from which to return the results. | 
 **orderBy** | **[]string** | Ordering  * &#x60;pulp_created&#x60; - by CV created * &#x60;-pulp_created&#x60; - by CV created (descending) * &#x60;namespace&#x60; - by CV namespace * &#x60;-namespace&#x60; - by CV namespace (descending) * &#x60;name&#x60; - by CV name * &#x60;-name&#x60; - by CV name (descending) * &#x60;version&#x60; - by CV version * &#x60;-version&#x60; - by CV version (descending) | 
 **q** | **string** |  | 
 **repository** | **[]string** | Filter collectionversions that are in these repository ids. | 
 **repositoryLabel** | **string** | Filter labels by search string | 
 **repositoryName** | **[]string** | Filter collectionversions that are in these repositories. | 
 **repositoryVersion** | **string** |  | 
 **signed** | **bool** |  | 
 **tags** | **string** | Filter by comma separate list of tags that must all be matched | 
 **version** | **string** |  | 
 **versionRange** | **string** |  | 
 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

### Return type

[**PaginatedCollectionVersionSearchListResponseList**](PaginatedCollectionVersionSearchListResponseList.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PulpAnsibleGalaxyDefaultApiV3PluginAnsibleSearchCollectionVersionsRebuild

> CollectionVersionSearchListResponse PulpAnsibleGalaxyDefaultApiV3PluginAnsibleSearchCollectionVersionsRebuild(ctx).CollectionVersionSearchList(collectionVersionSearchList).Execute()





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
    collectionVersionSearchList := *openapiclient.NewCollectionVersionSearchList(*openapiclient.NewRepository("Name_example"), map[string]interface{}(123), "TODO", false, false, false) // CollectionVersionSearchList | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PulpAnsibleDefaultApiV3PluginAnsibleSearchCollectionVersionsAPI.PulpAnsibleGalaxyDefaultApiV3PluginAnsibleSearchCollectionVersionsRebuild(context.Background()).CollectionVersionSearchList(collectionVersionSearchList).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PulpAnsibleDefaultApiV3PluginAnsibleSearchCollectionVersionsAPI.PulpAnsibleGalaxyDefaultApiV3PluginAnsibleSearchCollectionVersionsRebuild``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PulpAnsibleGalaxyDefaultApiV3PluginAnsibleSearchCollectionVersionsRebuild`: CollectionVersionSearchListResponse
    fmt.Fprintf(os.Stdout, "Response from `PulpAnsibleDefaultApiV3PluginAnsibleSearchCollectionVersionsAPI.PulpAnsibleGalaxyDefaultApiV3PluginAnsibleSearchCollectionVersionsRebuild`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPulpAnsibleGalaxyDefaultApiV3PluginAnsibleSearchCollectionVersionsRebuildRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **collectionVersionSearchList** | [**CollectionVersionSearchList**](CollectionVersionSearchList.md) |  | 

### Return type

[**CollectionVersionSearchListResponse**](CollectionVersionSearchListResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

