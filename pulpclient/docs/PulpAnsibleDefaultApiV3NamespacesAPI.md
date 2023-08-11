# \PulpAnsibleDefaultApiV3NamespacesAPI

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PulpAnsibleGalaxyDefaultApiV3NamespacesList**](PulpAnsibleDefaultApiV3NamespacesAPI.md#PulpAnsibleGalaxyDefaultApiV3NamespacesList) | **Get** /pulp_ansible/galaxy/default/api/v3/namespaces/ | 
[**PulpAnsibleGalaxyDefaultApiV3NamespacesRead**](PulpAnsibleDefaultApiV3NamespacesAPI.md#PulpAnsibleGalaxyDefaultApiV3NamespacesRead) | **Get** /pulp_ansible/galaxy/default/api/v3/namespaces/{name}/ | 



## PulpAnsibleGalaxyDefaultApiV3NamespacesList

> PaginatedansibleAnsibleNamespaceMetadataResponseList PulpAnsibleGalaxyDefaultApiV3NamespacesList(ctx).Company(company).CompanyContains(companyContains).CompanyIcontains(companyIcontains).CompanyIn(companyIn).CompanyStartswith(companyStartswith).Limit(limit).MetadataSha256(metadataSha256).MetadataSha256In(metadataSha256In).Name(name).NameContains(nameContains).NameIcontains(nameIcontains).NameIn(nameIn).NameStartswith(nameStartswith).Offset(offset).Ordering(ordering).PulpHrefIn(pulpHrefIn).PulpIdIn(pulpIdIn).Fields(fields).ExcludeFields(excludeFields).Execute()





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
    company := "company_example" // string | Filter results where company matches value (optional)
    companyContains := "companyContains_example" // string | Filter results where company contains value (optional)
    companyIcontains := "companyIcontains_example" // string | Filter results where company contains value (optional)
    companyIn := []string{"Inner_example"} // []string | Filter results where company is in a comma-separated list of values (optional)
    companyStartswith := "companyStartswith_example" // string | Filter results where company starts with value (optional)
    limit := int32(56) // int32 | Number of results to return per page. (optional)
    metadataSha256 := "metadataSha256_example" // string | Filter results where metadata_sha256 matches value (optional)
    metadataSha256In := []string{"Inner_example"} // []string | Filter results where metadata_sha256 is in a comma-separated list of values (optional)
    name := "name_example" // string | Filter results where name matches value (optional)
    nameContains := "nameContains_example" // string | Filter results where name contains value (optional)
    nameIcontains := "nameIcontains_example" // string | Filter results where name contains value (optional)
    nameIn := []string{"Inner_example"} // []string | Filter results where name is in a comma-separated list of values (optional)
    nameStartswith := "nameStartswith_example" // string | Filter results where name starts with value (optional)
    offset := int32(56) // int32 | The initial index from which to return the results. (optional)
    ordering := []string{"Ordering_example"} // []string | Ordering  * `pulp_id` - Pulp id * `-pulp_id` - Pulp id (descending) * `pulp_created` - Pulp created * `-pulp_created` - Pulp created (descending) * `pulp_last_updated` - Pulp last updated * `-pulp_last_updated` - Pulp last updated (descending) * `pulp_type` - Pulp type * `-pulp_type` - Pulp type (descending) * `upstream_id` - Upstream id * `-upstream_id` - Upstream id (descending) * `timestamp_of_interest` - Timestamp of interest * `-timestamp_of_interest` - Timestamp of interest (descending) * `name` - Name * `-name` - Name (descending) * `company` - Company * `-company` - Company (descending) * `email` - Email * `-email` - Email (descending) * `description` - Description * `-description` - Description (descending) * `resources` - Resources * `-resources` - Resources (descending) * `links` - Links * `-links` - Links (descending) * `avatar_sha256` - Avatar sha256 * `-avatar_sha256` - Avatar sha256 (descending) * `metadata_sha256` - Metadata sha256 * `-metadata_sha256` - Metadata sha256 (descending) * `pk` - Pk * `-pk` - Pk (descending) (optional)
    pulpHrefIn := []string{"Inner_example"} // []string | Multiple values may be separated by commas. (optional)
    pulpIdIn := []string{"Inner_example"} // []string | Multiple values may be separated by commas. (optional)
    fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
    excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PulpAnsibleDefaultApiV3NamespacesAPI.PulpAnsibleGalaxyDefaultApiV3NamespacesList(context.Background()).Company(company).CompanyContains(companyContains).CompanyIcontains(companyIcontains).CompanyIn(companyIn).CompanyStartswith(companyStartswith).Limit(limit).MetadataSha256(metadataSha256).MetadataSha256In(metadataSha256In).Name(name).NameContains(nameContains).NameIcontains(nameIcontains).NameIn(nameIn).NameStartswith(nameStartswith).Offset(offset).Ordering(ordering).PulpHrefIn(pulpHrefIn).PulpIdIn(pulpIdIn).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PulpAnsibleDefaultApiV3NamespacesAPI.PulpAnsibleGalaxyDefaultApiV3NamespacesList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PulpAnsibleGalaxyDefaultApiV3NamespacesList`: PaginatedansibleAnsibleNamespaceMetadataResponseList
    fmt.Fprintf(os.Stdout, "Response from `PulpAnsibleDefaultApiV3NamespacesAPI.PulpAnsibleGalaxyDefaultApiV3NamespacesList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPulpAnsibleGalaxyDefaultApiV3NamespacesListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **company** | **string** | Filter results where company matches value | 
 **companyContains** | **string** | Filter results where company contains value | 
 **companyIcontains** | **string** | Filter results where company contains value | 
 **companyIn** | **[]string** | Filter results where company is in a comma-separated list of values | 
 **companyStartswith** | **string** | Filter results where company starts with value | 
 **limit** | **int32** | Number of results to return per page. | 
 **metadataSha256** | **string** | Filter results where metadata_sha256 matches value | 
 **metadataSha256In** | **[]string** | Filter results where metadata_sha256 is in a comma-separated list of values | 
 **name** | **string** | Filter results where name matches value | 
 **nameContains** | **string** | Filter results where name contains value | 
 **nameIcontains** | **string** | Filter results where name contains value | 
 **nameIn** | **[]string** | Filter results where name is in a comma-separated list of values | 
 **nameStartswith** | **string** | Filter results where name starts with value | 
 **offset** | **int32** | The initial index from which to return the results. | 
 **ordering** | **[]string** | Ordering  * &#x60;pulp_id&#x60; - Pulp id * &#x60;-pulp_id&#x60; - Pulp id (descending) * &#x60;pulp_created&#x60; - Pulp created * &#x60;-pulp_created&#x60; - Pulp created (descending) * &#x60;pulp_last_updated&#x60; - Pulp last updated * &#x60;-pulp_last_updated&#x60; - Pulp last updated (descending) * &#x60;pulp_type&#x60; - Pulp type * &#x60;-pulp_type&#x60; - Pulp type (descending) * &#x60;upstream_id&#x60; - Upstream id * &#x60;-upstream_id&#x60; - Upstream id (descending) * &#x60;timestamp_of_interest&#x60; - Timestamp of interest * &#x60;-timestamp_of_interest&#x60; - Timestamp of interest (descending) * &#x60;name&#x60; - Name * &#x60;-name&#x60; - Name (descending) * &#x60;company&#x60; - Company * &#x60;-company&#x60; - Company (descending) * &#x60;email&#x60; - Email * &#x60;-email&#x60; - Email (descending) * &#x60;description&#x60; - Description * &#x60;-description&#x60; - Description (descending) * &#x60;resources&#x60; - Resources * &#x60;-resources&#x60; - Resources (descending) * &#x60;links&#x60; - Links * &#x60;-links&#x60; - Links (descending) * &#x60;avatar_sha256&#x60; - Avatar sha256 * &#x60;-avatar_sha256&#x60; - Avatar sha256 (descending) * &#x60;metadata_sha256&#x60; - Metadata sha256 * &#x60;-metadata_sha256&#x60; - Metadata sha256 (descending) * &#x60;pk&#x60; - Pk * &#x60;-pk&#x60; - Pk (descending) | 
 **pulpHrefIn** | **[]string** | Multiple values may be separated by commas. | 
 **pulpIdIn** | **[]string** | Multiple values may be separated by commas. | 
 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

### Return type

[**PaginatedansibleAnsibleNamespaceMetadataResponseList**](PaginatedansibleAnsibleNamespaceMetadataResponseList.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PulpAnsibleGalaxyDefaultApiV3NamespacesRead

> AnsibleAnsibleNamespaceMetadataResponse PulpAnsibleGalaxyDefaultApiV3NamespacesRead(ctx, name).Fields(fields).ExcludeFields(excludeFields).Execute()





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
    fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
    excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PulpAnsibleDefaultApiV3NamespacesAPI.PulpAnsibleGalaxyDefaultApiV3NamespacesRead(context.Background(), name).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PulpAnsibleDefaultApiV3NamespacesAPI.PulpAnsibleGalaxyDefaultApiV3NamespacesRead``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PulpAnsibleGalaxyDefaultApiV3NamespacesRead`: AnsibleAnsibleNamespaceMetadataResponse
    fmt.Fprintf(os.Stdout, "Response from `PulpAnsibleDefaultApiV3NamespacesAPI.PulpAnsibleGalaxyDefaultApiV3NamespacesRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPulpAnsibleGalaxyDefaultApiV3NamespacesReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

### Return type

[**AnsibleAnsibleNamespaceMetadataResponse**](AnsibleAnsibleNamespaceMetadataResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

