# ChecksumResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Path** | **string** | File path. | 
**Checksum** | **string** | Checksum for the file. | 

## Methods

### NewChecksumResponse

`func NewChecksumResponse(path string, checksum string, ) *ChecksumResponse`

NewChecksumResponse instantiates a new ChecksumResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChecksumResponseWithDefaults

`func NewChecksumResponseWithDefaults() *ChecksumResponse`

NewChecksumResponseWithDefaults instantiates a new ChecksumResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPath

`func (o *ChecksumResponse) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *ChecksumResponse) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *ChecksumResponse) SetPath(v string)`

SetPath sets Path field to given value.


### GetChecksum

`func (o *ChecksumResponse) GetChecksum() string`

GetChecksum returns the Checksum field if non-nil, zero value otherwise.

### GetChecksumOk

`func (o *ChecksumResponse) GetChecksumOk() (*string, bool)`

GetChecksumOk returns a tuple with the Checksum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChecksum

`func (o *ChecksumResponse) SetChecksum(v string)`

SetChecksum sets Checksum field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


