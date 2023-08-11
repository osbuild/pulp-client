# PackageMetadataResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LastSerial** | **int64** | Cache value from last PyPI sync | 
**Info** | **map[string]interface{}** | Core metadata of the package | 
**Releases** | **map[string]interface{}** | List of all the releases of the package | 
**Urls** | **map[string]interface{}** |  | 

## Methods

### NewPackageMetadataResponse

`func NewPackageMetadataResponse(lastSerial int64, info map[string]interface{}, releases map[string]interface{}, urls map[string]interface{}, ) *PackageMetadataResponse`

NewPackageMetadataResponse instantiates a new PackageMetadataResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPackageMetadataResponseWithDefaults

`func NewPackageMetadataResponseWithDefaults() *PackageMetadataResponse`

NewPackageMetadataResponseWithDefaults instantiates a new PackageMetadataResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLastSerial

`func (o *PackageMetadataResponse) GetLastSerial() int64`

GetLastSerial returns the LastSerial field if non-nil, zero value otherwise.

### GetLastSerialOk

`func (o *PackageMetadataResponse) GetLastSerialOk() (*int64, bool)`

GetLastSerialOk returns a tuple with the LastSerial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSerial

`func (o *PackageMetadataResponse) SetLastSerial(v int64)`

SetLastSerial sets LastSerial field to given value.


### GetInfo

`func (o *PackageMetadataResponse) GetInfo() map[string]interface{}`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *PackageMetadataResponse) GetInfoOk() (*map[string]interface{}, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *PackageMetadataResponse) SetInfo(v map[string]interface{})`

SetInfo sets Info field to given value.


### GetReleases

`func (o *PackageMetadataResponse) GetReleases() map[string]interface{}`

GetReleases returns the Releases field if non-nil, zero value otherwise.

### GetReleasesOk

`func (o *PackageMetadataResponse) GetReleasesOk() (*map[string]interface{}, bool)`

GetReleasesOk returns a tuple with the Releases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleases

`func (o *PackageMetadataResponse) SetReleases(v map[string]interface{})`

SetReleases sets Releases field to given value.


### GetUrls

`func (o *PackageMetadataResponse) GetUrls() map[string]interface{}`

GetUrls returns the Urls field if non-nil, zero value otherwise.

### GetUrlsOk

`func (o *PackageMetadataResponse) GetUrlsOk() (*map[string]interface{}, bool)`

GetUrlsOk returns a tuple with the Urls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrls

`func (o *PackageMetadataResponse) SetUrls(v map[string]interface{})`

SetUrls sets Urls field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


