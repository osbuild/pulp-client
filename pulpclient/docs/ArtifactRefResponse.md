# ArtifactRefResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filename** | **string** |  | 
**Sha256** | **string** |  | 
**Size** | **int64** |  | 

## Methods

### NewArtifactRefResponse

`func NewArtifactRefResponse(filename string, sha256 string, size int64, ) *ArtifactRefResponse`

NewArtifactRefResponse instantiates a new ArtifactRefResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArtifactRefResponseWithDefaults

`func NewArtifactRefResponseWithDefaults() *ArtifactRefResponse`

NewArtifactRefResponseWithDefaults instantiates a new ArtifactRefResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilename

`func (o *ArtifactRefResponse) GetFilename() string`

GetFilename returns the Filename field if non-nil, zero value otherwise.

### GetFilenameOk

`func (o *ArtifactRefResponse) GetFilenameOk() (*string, bool)`

GetFilenameOk returns a tuple with the Filename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilename

`func (o *ArtifactRefResponse) SetFilename(v string)`

SetFilename sets Filename field to given value.


### GetSha256

`func (o *ArtifactRefResponse) GetSha256() string`

GetSha256 returns the Sha256 field if non-nil, zero value otherwise.

### GetSha256Ok

`func (o *ArtifactRefResponse) GetSha256Ok() (*string, bool)`

GetSha256Ok returns a tuple with the Sha256 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha256

`func (o *ArtifactRefResponse) SetSha256(v string)`

SetSha256 sets Sha256 field to given value.


### GetSize

`func (o *ArtifactRefResponse) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *ArtifactRefResponse) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *ArtifactRefResponse) SetSize(v int64)`

SetSize sets Size field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


