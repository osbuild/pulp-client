# \GalaxyDetailAPI

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GalaxyCollectionDetailGet**](GalaxyDetailAPI.md#GalaxyCollectionDetailGet) | **Get** {ansible_collection_href} | 



## GalaxyCollectionDetailGet

> GalaxyCollectionResponse GalaxyCollectionDetailGet(ctx, ansibleCollectionHref).Fields(fields).ExcludeFields(excludeFields).Execute()





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
    ansibleCollectionHref := "ansibleCollectionHref_example" // string | 
    fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
    excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GalaxyDetailAPI.GalaxyCollectionDetailGet(context.Background(), ansibleCollectionHref).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GalaxyDetailAPI.GalaxyCollectionDetailGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GalaxyCollectionDetailGet`: GalaxyCollectionResponse
    fmt.Fprintf(os.Stdout, "Response from `GalaxyDetailAPI.GalaxyCollectionDetailGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ansibleCollectionHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGalaxyCollectionDetailGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

### Return type

[**GalaxyCollectionResponse**](GalaxyCollectionResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

