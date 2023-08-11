# \RpmCompsAPI

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RpmCompsUpload**](RpmCompsAPI.md#RpmCompsUpload) | **Post** /pulp/api/v3/rpm/comps/ | Upload comps.xml



## RpmCompsUpload

> AsyncOperationResponse RpmCompsUpload(ctx).File(file).Repository(repository).Replace(replace).Execute()

Upload comps.xml



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
    file := os.NewFile(1234, "some_file") // *os.File | Full path of a comps.xml file that may be parsed into comps.xml Content units.
    repository := "repository_example" // string | URI of an RPM repository the comps.xml content units should be associated to. (optional)
    replace := true // bool | If true, incoming comps.xml replaces existing comps-related ContentUnits in the specified repository. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RpmCompsAPI.RpmCompsUpload(context.Background()).File(file).Repository(repository).Replace(replace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RpmCompsAPI.RpmCompsUpload``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RpmCompsUpload`: AsyncOperationResponse
    fmt.Fprintf(os.Stdout, "Response from `RpmCompsAPI.RpmCompsUpload`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRpmCompsUploadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **file** | ***os.File** | Full path of a comps.xml file that may be parsed into comps.xml Content units. | 
 **repository** | **string** | URI of an RPM repository the comps.xml content units should be associated to. | 
 **replace** | **bool** | If true, incoming comps.xml replaces existing comps-related ContentUnits in the specified repository. | 

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

