# \ContentAPI

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ContentList**](ContentAPI.md#ContentList) | **Get** /pulp/api/v3/content/ | List content



## ContentList

> PaginatedMultipleArtifactContentResponseList ContentList(ctx).Limit(limit).Offset(offset).Ordering(ordering).PulpHrefIn(pulpHrefIn).PulpIdIn(pulpIdIn).PulpTypeIn(pulpTypeIn).RepositoryVersion(repositoryVersion).RepositoryVersionAdded(repositoryVersionAdded).RepositoryVersionRemoved(repositoryVersionRemoved).Fields(fields).ExcludeFields(excludeFields).Execute()

List content



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
    ordering := []string{"Ordering_example"} // []string | Ordering  * `pk` - Pk * `-pk` - Pk (descending) (optional)
    pulpHrefIn := []string{"Inner_example"} // []string | Multiple values may be separated by commas. (optional)
    pulpIdIn := []string{"Inner_example"} // []string | Multiple values may be separated by commas. (optional)
    pulpTypeIn := []string{"PulpTypeIn_example"} // []string | Pulp type is in  * `core.publishedmetadata` - core.publishedmetadata * `ansible.role` - ansible.role * `ansible.collection_version` - ansible.collection_version * `ansible.collection_mark` - ansible.collection_mark * `ansible.collection_signature` - ansible.collection_signature * `ansible.namespace` - ansible.namespace * `ansible.collection_deprecation` - ansible.collection_deprecation * `container.blob` - container.blob * `container.manifest` - container.manifest * `container.tag` - container.tag * `container.signature` - container.signature * `deb.package` - deb.package * `deb.installer_package` - deb.installer_package * `deb.generic` - deb.generic * `deb.release` - deb.release * `deb.release_architecture` - deb.release_architecture * `deb.release_component` - deb.release_component * `deb.package_release_component` - deb.package_release_component * `deb.release_file` - deb.release_file * `deb.package_index` - deb.package_index * `deb.installer_file_index` - deb.installer_file_index * `file.file` - file.file * `gem.shallow-gem` - gem.shallow-gem * `gem.gem` - gem.gem * `maven.artifact` - maven.artifact * `maven.metadata` - maven.metadata * `ostree.object` - ostree.object * `ostree.commit` - ostree.commit * `ostree.refs` - ostree.refs * `ostree.content` - ostree.content * `ostree.config` - ostree.config * `ostree.summary` - ostree.summary * `python.python` - python.python * `rpm.advisory` - rpm.advisory * `rpm.packagegroup` - rpm.packagegroup * `rpm.packagecategory` - rpm.packagecategory * `rpm.packageenvironment` - rpm.packageenvironment * `rpm.packagelangpacks` - rpm.packagelangpacks * `rpm.repo_metadata_file` - rpm.repo_metadata_file * `rpm.distribution_tree` - rpm.distribution_tree * `rpm.package` - rpm.package * `rpm.modulemd` - rpm.modulemd * `rpm.modulemd_defaults` - rpm.modulemd_defaults * `rpm.modulemd_obsolete` - rpm.modulemd_obsolete (optional)
    repositoryVersion := "repositoryVersion_example" // string | Repository Version referenced by HREF (optional)
    repositoryVersionAdded := "repositoryVersionAdded_example" // string | Repository Version referenced by HREF (optional)
    repositoryVersionRemoved := "repositoryVersionRemoved_example" // string | Repository Version referenced by HREF (optional)
    fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
    excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContentAPI.ContentList(context.Background()).Limit(limit).Offset(offset).Ordering(ordering).PulpHrefIn(pulpHrefIn).PulpIdIn(pulpIdIn).PulpTypeIn(pulpTypeIn).RepositoryVersion(repositoryVersion).RepositoryVersionAdded(repositoryVersionAdded).RepositoryVersionRemoved(repositoryVersionRemoved).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContentAPI.ContentList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContentList`: PaginatedMultipleArtifactContentResponseList
    fmt.Fprintf(os.Stdout, "Response from `ContentAPI.ContentList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiContentListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Number of results to return per page. | 
 **offset** | **int32** | The initial index from which to return the results. | 
 **ordering** | **[]string** | Ordering  * &#x60;pk&#x60; - Pk * &#x60;-pk&#x60; - Pk (descending) | 
 **pulpHrefIn** | **[]string** | Multiple values may be separated by commas. | 
 **pulpIdIn** | **[]string** | Multiple values may be separated by commas. | 
 **pulpTypeIn** | **[]string** | Pulp type is in  * &#x60;core.publishedmetadata&#x60; - core.publishedmetadata * &#x60;ansible.role&#x60; - ansible.role * &#x60;ansible.collection_version&#x60; - ansible.collection_version * &#x60;ansible.collection_mark&#x60; - ansible.collection_mark * &#x60;ansible.collection_signature&#x60; - ansible.collection_signature * &#x60;ansible.namespace&#x60; - ansible.namespace * &#x60;ansible.collection_deprecation&#x60; - ansible.collection_deprecation * &#x60;container.blob&#x60; - container.blob * &#x60;container.manifest&#x60; - container.manifest * &#x60;container.tag&#x60; - container.tag * &#x60;container.signature&#x60; - container.signature * &#x60;deb.package&#x60; - deb.package * &#x60;deb.installer_package&#x60; - deb.installer_package * &#x60;deb.generic&#x60; - deb.generic * &#x60;deb.release&#x60; - deb.release * &#x60;deb.release_architecture&#x60; - deb.release_architecture * &#x60;deb.release_component&#x60; - deb.release_component * &#x60;deb.package_release_component&#x60; - deb.package_release_component * &#x60;deb.release_file&#x60; - deb.release_file * &#x60;deb.package_index&#x60; - deb.package_index * &#x60;deb.installer_file_index&#x60; - deb.installer_file_index * &#x60;file.file&#x60; - file.file * &#x60;gem.shallow-gem&#x60; - gem.shallow-gem * &#x60;gem.gem&#x60; - gem.gem * &#x60;maven.artifact&#x60; - maven.artifact * &#x60;maven.metadata&#x60; - maven.metadata * &#x60;ostree.object&#x60; - ostree.object * &#x60;ostree.commit&#x60; - ostree.commit * &#x60;ostree.refs&#x60; - ostree.refs * &#x60;ostree.content&#x60; - ostree.content * &#x60;ostree.config&#x60; - ostree.config * &#x60;ostree.summary&#x60; - ostree.summary * &#x60;python.python&#x60; - python.python * &#x60;rpm.advisory&#x60; - rpm.advisory * &#x60;rpm.packagegroup&#x60; - rpm.packagegroup * &#x60;rpm.packagecategory&#x60; - rpm.packagecategory * &#x60;rpm.packageenvironment&#x60; - rpm.packageenvironment * &#x60;rpm.packagelangpacks&#x60; - rpm.packagelangpacks * &#x60;rpm.repo_metadata_file&#x60; - rpm.repo_metadata_file * &#x60;rpm.distribution_tree&#x60; - rpm.distribution_tree * &#x60;rpm.package&#x60; - rpm.package * &#x60;rpm.modulemd&#x60; - rpm.modulemd * &#x60;rpm.modulemd_defaults&#x60; - rpm.modulemd_defaults * &#x60;rpm.modulemd_obsolete&#x60; - rpm.modulemd_obsolete | 
 **repositoryVersion** | **string** | Repository Version referenced by HREF | 
 **repositoryVersionAdded** | **string** | Repository Version referenced by HREF | 
 **repositoryVersionRemoved** | **string** | Repository Version referenced by HREF | 
 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

### Return type

[**PaginatedMultipleArtifactContentResponseList**](PaginatedMultipleArtifactContentResponseList.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

