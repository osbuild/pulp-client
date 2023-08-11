# \PulpAnsibleArtifactsCollectionsV3API

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PulpAnsibleGalaxyApiV3ArtifactsCollectionsCreate**](PulpAnsibleArtifactsCollectionsV3API.md#PulpAnsibleGalaxyApiV3ArtifactsCollectionsCreate) | **Post** /pulp_ansible/galaxy/{path}/api/v3/artifacts/collections/ | Upload a collection
[**PulpAnsibleGalaxyApiV3PluginAnsibleContentCollectionsArtifactsCreate**](PulpAnsibleArtifactsCollectionsV3API.md#PulpAnsibleGalaxyApiV3PluginAnsibleContentCollectionsArtifactsCreate) | **Post** /pulp_ansible/galaxy/{path}/api/v3/plugin/ansible/content/{distro_base_path}/collections/artifacts/ | Upload a collection
[**PulpAnsibleGalaxyDefaultApiV3ArtifactsCollectionsCreate**](PulpAnsibleArtifactsCollectionsV3API.md#PulpAnsibleGalaxyDefaultApiV3ArtifactsCollectionsCreate) | **Post** /pulp_ansible/galaxy/default/api/v3/artifacts/collections/ | Upload a collection
[**PulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsArtifactsCreate**](PulpAnsibleArtifactsCollectionsV3API.md#PulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsArtifactsCreate) | **Post** /pulp_ansible/galaxy/default/api/v3/plugin/ansible/content/{distro_base_path}/collections/artifacts/ | Upload a collection



## PulpAnsibleGalaxyApiV3ArtifactsCollectionsCreate

> AsyncOperationResponse PulpAnsibleGalaxyApiV3ArtifactsCollectionsCreate(ctx, path).File(file).Sha256(sha256).ExpectedNamespace(expectedNamespace).ExpectedName(expectedName).ExpectedVersion(expectedVersion).Execute()

Upload a collection



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
    path := "path_example" // string | 
    file := os.NewFile(1234, "some_file") // *os.File | The Collection tarball.
    sha256 := "sha256_example" // string | An optional sha256 checksum of the uploaded file. (optional)
    expectedNamespace := "expectedNamespace_example" // string | The expected 'namespace' of the Collection to be verified against the metadata during import. (optional)
    expectedName := "expectedName_example" // string | The expected 'name' of the Collection to be verified against the metadata during import. (optional)
    expectedVersion := "expectedVersion_example" // string | The expected version of the Collection to be verified against the metadata during import. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PulpAnsibleArtifactsCollectionsV3API.PulpAnsibleGalaxyApiV3ArtifactsCollectionsCreate(context.Background(), path).File(file).Sha256(sha256).ExpectedNamespace(expectedNamespace).ExpectedName(expectedName).ExpectedVersion(expectedVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PulpAnsibleArtifactsCollectionsV3API.PulpAnsibleGalaxyApiV3ArtifactsCollectionsCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PulpAnsibleGalaxyApiV3ArtifactsCollectionsCreate`: AsyncOperationResponse
    fmt.Fprintf(os.Stdout, "Response from `PulpAnsibleArtifactsCollectionsV3API.PulpAnsibleGalaxyApiV3ArtifactsCollectionsCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**path** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPulpAnsibleGalaxyApiV3ArtifactsCollectionsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **file** | ***os.File** | The Collection tarball. | 
 **sha256** | **string** | An optional sha256 checksum of the uploaded file. | 
 **expectedNamespace** | **string** | The expected &#39;namespace&#39; of the Collection to be verified against the metadata during import. | 
 **expectedName** | **string** | The expected &#39;name&#39; of the Collection to be verified against the metadata during import. | 
 **expectedVersion** | **string** | The expected version of the Collection to be verified against the metadata during import. | 

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


## PulpAnsibleGalaxyApiV3PluginAnsibleContentCollectionsArtifactsCreate

> AsyncOperationResponse PulpAnsibleGalaxyApiV3PluginAnsibleContentCollectionsArtifactsCreate(ctx, distroBasePath, path).File(file).Sha256(sha256).ExpectedNamespace(expectedNamespace).ExpectedName(expectedName).ExpectedVersion(expectedVersion).Execute()

Upload a collection



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
    distroBasePath := "distroBasePath_example" // string | 
    path := "path_example" // string | 
    file := os.NewFile(1234, "some_file") // *os.File | The Collection tarball.
    sha256 := "sha256_example" // string | An optional sha256 checksum of the uploaded file. (optional)
    expectedNamespace := "expectedNamespace_example" // string | The expected 'namespace' of the Collection to be verified against the metadata during import. (optional)
    expectedName := "expectedName_example" // string | The expected 'name' of the Collection to be verified against the metadata during import. (optional)
    expectedVersion := "expectedVersion_example" // string | The expected version of the Collection to be verified against the metadata during import. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PulpAnsibleArtifactsCollectionsV3API.PulpAnsibleGalaxyApiV3PluginAnsibleContentCollectionsArtifactsCreate(context.Background(), distroBasePath, path).File(file).Sha256(sha256).ExpectedNamespace(expectedNamespace).ExpectedName(expectedName).ExpectedVersion(expectedVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PulpAnsibleArtifactsCollectionsV3API.PulpAnsibleGalaxyApiV3PluginAnsibleContentCollectionsArtifactsCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PulpAnsibleGalaxyApiV3PluginAnsibleContentCollectionsArtifactsCreate`: AsyncOperationResponse
    fmt.Fprintf(os.Stdout, "Response from `PulpAnsibleArtifactsCollectionsV3API.PulpAnsibleGalaxyApiV3PluginAnsibleContentCollectionsArtifactsCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**distroBasePath** | **string** |  | 
**path** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPulpAnsibleGalaxyApiV3PluginAnsibleContentCollectionsArtifactsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **file** | ***os.File** | The Collection tarball. | 
 **sha256** | **string** | An optional sha256 checksum of the uploaded file. | 
 **expectedNamespace** | **string** | The expected &#39;namespace&#39; of the Collection to be verified against the metadata during import. | 
 **expectedName** | **string** | The expected &#39;name&#39; of the Collection to be verified against the metadata during import. | 
 **expectedVersion** | **string** | The expected version of the Collection to be verified against the metadata during import. | 

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


## PulpAnsibleGalaxyDefaultApiV3ArtifactsCollectionsCreate

> AsyncOperationResponse PulpAnsibleGalaxyDefaultApiV3ArtifactsCollectionsCreate(ctx).File(file).Sha256(sha256).ExpectedNamespace(expectedNamespace).ExpectedName(expectedName).ExpectedVersion(expectedVersion).Execute()

Upload a collection



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
    file := os.NewFile(1234, "some_file") // *os.File | The Collection tarball.
    sha256 := "sha256_example" // string | An optional sha256 checksum of the uploaded file. (optional)
    expectedNamespace := "expectedNamespace_example" // string | The expected 'namespace' of the Collection to be verified against the metadata during import. (optional)
    expectedName := "expectedName_example" // string | The expected 'name' of the Collection to be verified against the metadata during import. (optional)
    expectedVersion := "expectedVersion_example" // string | The expected version of the Collection to be verified against the metadata during import. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PulpAnsibleArtifactsCollectionsV3API.PulpAnsibleGalaxyDefaultApiV3ArtifactsCollectionsCreate(context.Background()).File(file).Sha256(sha256).ExpectedNamespace(expectedNamespace).ExpectedName(expectedName).ExpectedVersion(expectedVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PulpAnsibleArtifactsCollectionsV3API.PulpAnsibleGalaxyDefaultApiV3ArtifactsCollectionsCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PulpAnsibleGalaxyDefaultApiV3ArtifactsCollectionsCreate`: AsyncOperationResponse
    fmt.Fprintf(os.Stdout, "Response from `PulpAnsibleArtifactsCollectionsV3API.PulpAnsibleGalaxyDefaultApiV3ArtifactsCollectionsCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPulpAnsibleGalaxyDefaultApiV3ArtifactsCollectionsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **file** | ***os.File** | The Collection tarball. | 
 **sha256** | **string** | An optional sha256 checksum of the uploaded file. | 
 **expectedNamespace** | **string** | The expected &#39;namespace&#39; of the Collection to be verified against the metadata during import. | 
 **expectedName** | **string** | The expected &#39;name&#39; of the Collection to be verified against the metadata during import. | 
 **expectedVersion** | **string** | The expected version of the Collection to be verified against the metadata during import. | 

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


## PulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsArtifactsCreate

> AsyncOperationResponse PulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsArtifactsCreate(ctx, distroBasePath).File(file).Sha256(sha256).ExpectedNamespace(expectedNamespace).ExpectedName(expectedName).ExpectedVersion(expectedVersion).Execute()

Upload a collection



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
    distroBasePath := "distroBasePath_example" // string | 
    file := os.NewFile(1234, "some_file") // *os.File | The Collection tarball.
    sha256 := "sha256_example" // string | An optional sha256 checksum of the uploaded file. (optional)
    expectedNamespace := "expectedNamespace_example" // string | The expected 'namespace' of the Collection to be verified against the metadata during import. (optional)
    expectedName := "expectedName_example" // string | The expected 'name' of the Collection to be verified against the metadata during import. (optional)
    expectedVersion := "expectedVersion_example" // string | The expected version of the Collection to be verified against the metadata during import. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PulpAnsibleArtifactsCollectionsV3API.PulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsArtifactsCreate(context.Background(), distroBasePath).File(file).Sha256(sha256).ExpectedNamespace(expectedNamespace).ExpectedName(expectedName).ExpectedVersion(expectedVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PulpAnsibleArtifactsCollectionsV3API.PulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsArtifactsCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsArtifactsCreate`: AsyncOperationResponse
    fmt.Fprintf(os.Stdout, "Response from `PulpAnsibleArtifactsCollectionsV3API.PulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsArtifactsCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**distroBasePath** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsArtifactsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **file** | ***os.File** | The Collection tarball. | 
 **sha256** | **string** | An optional sha256 checksum of the uploaded file. | 
 **expectedNamespace** | **string** | The expected &#39;namespace&#39; of the Collection to be verified against the metadata during import. | 
 **expectedName** | **string** | The expected &#39;name&#39; of the Collection to be verified against the metadata during import. | 
 **expectedVersion** | **string** | The expected version of the Collection to be verified against the metadata during import. | 

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

