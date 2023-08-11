# UploadChunk

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**File** | ***os.File** | A chunk of the uploaded file. | 
**Sha256** | Pointer to **NullableString** | The SHA-256 checksum of the chunk if available. | [optional] 

## Methods

### NewUploadChunk

`func NewUploadChunk(file *os.File, ) *UploadChunk`

NewUploadChunk instantiates a new UploadChunk object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUploadChunkWithDefaults

`func NewUploadChunkWithDefaults() *UploadChunk`

NewUploadChunkWithDefaults instantiates a new UploadChunk object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFile

`func (o *UploadChunk) GetFile() *os.File`

GetFile returns the File field if non-nil, zero value otherwise.

### GetFileOk

`func (o *UploadChunk) GetFileOk() (**os.File, bool)`

GetFileOk returns a tuple with the File field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFile

`func (o *UploadChunk) SetFile(v *os.File)`

SetFile sets File field to given value.


### GetSha256

`func (o *UploadChunk) GetSha256() string`

GetSha256 returns the Sha256 field if non-nil, zero value otherwise.

### GetSha256Ok

`func (o *UploadChunk) GetSha256Ok() (*string, bool)`

GetSha256Ok returns a tuple with the Sha256 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha256

`func (o *UploadChunk) SetSha256(v string)`

SetSha256 sets Sha256 field to given value.

### HasSha256

`func (o *UploadChunk) HasSha256() bool`

HasSha256 returns a boolean if a field has been set.

### SetSha256Nil

`func (o *UploadChunk) SetSha256Nil(b bool)`

 SetSha256Nil sets the value for Sha256 to be an explicit nil

### UnsetSha256
`func (o *UploadChunk) UnsetSha256()`

UnsetSha256 ensures that no value is present for Sha256, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


