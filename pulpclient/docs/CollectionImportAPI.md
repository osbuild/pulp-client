# \CollectionImportAPI

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CollectionImportRead**](CollectionImportAPI.md#CollectionImportRead) | **Get** {ansible_collection_import_href} | Inspect a collection import



## CollectionImportRead

> CollectionImportDetailResponse CollectionImportRead(ctx, ansibleCollectionImportHref).Since(since).Fields(fields).ExcludeFields(excludeFields).Execute()

Inspect a collection import



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
    ansibleCollectionImportHref := "ansibleCollectionImportHref_example" // string | 
    since := "since_example" // string | Filter messages since a given timestamp (optional)
    fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
    excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CollectionImportAPI.CollectionImportRead(context.Background(), ansibleCollectionImportHref).Since(since).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CollectionImportAPI.CollectionImportRead``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CollectionImportRead`: CollectionImportDetailResponse
    fmt.Fprintf(os.Stdout, "Response from `CollectionImportAPI.CollectionImportRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ansibleCollectionImportHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCollectionImportReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **since** | **string** | Filter messages since a given timestamp | 
 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

### Return type

[**CollectionImportDetailResponse**](CollectionImportDetailResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

