# \PulpAnsibleDefaultApiV3CollectionVersionsAllAPI

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PulpAnsibleGalaxyDefaultApiV3CollectionVersionsAllList**](PulpAnsibleDefaultApiV3CollectionVersionsAllAPI.md#PulpAnsibleGalaxyDefaultApiV3CollectionVersionsAllList) | **Get** /pulp_ansible/galaxy/default/api/v3/collection_versions/all/ | 



## PulpAnsibleGalaxyDefaultApiV3CollectionVersionsAllList

> []UnpaginatedCollectionVersionResponse PulpAnsibleGalaxyDefaultApiV3CollectionVersionsAllList(ctx).Execute()





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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PulpAnsibleDefaultApiV3CollectionVersionsAllAPI.PulpAnsibleGalaxyDefaultApiV3CollectionVersionsAllList(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PulpAnsibleDefaultApiV3CollectionVersionsAllAPI.PulpAnsibleGalaxyDefaultApiV3CollectionVersionsAllList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PulpAnsibleGalaxyDefaultApiV3CollectionVersionsAllList`: []UnpaginatedCollectionVersionResponse
    fmt.Fprintf(os.Stdout, "Response from `PulpAnsibleDefaultApiV3CollectionVersionsAllAPI.PulpAnsibleGalaxyDefaultApiV3CollectionVersionsAllList`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPulpAnsibleGalaxyDefaultApiV3CollectionVersionsAllListRequest struct via the builder pattern


### Return type

[**[]UnpaginatedCollectionVersionResponse**](UnpaginatedCollectionVersionResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

