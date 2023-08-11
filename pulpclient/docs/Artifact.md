# Artifact

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**File** | ***os.File** | The stored file. | 
**Size** | Pointer to **int64** | The size of the file in bytes. | [optional] 
**Md5** | Pointer to **NullableString** | The MD5 checksum of the file if available. | [optional] 
**Sha1** | Pointer to **NullableString** | The SHA-1 checksum of the file if available. | [optional] 
**Sha224** | Pointer to **NullableString** | The SHA-224 checksum of the file if available. | [optional] 
**Sha256** | Pointer to **NullableString** | The SHA-256 checksum of the file if available. | [optional] 
**Sha384** | Pointer to **NullableString** | The SHA-384 checksum of the file if available. | [optional] 
**Sha512** | Pointer to **NullableString** | The SHA-512 checksum of the file if available. | [optional] 

## Methods

### NewArtifact

`func NewArtifact(file *os.File, ) *Artifact`

NewArtifact instantiates a new Artifact object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArtifactWithDefaults

`func NewArtifactWithDefaults() *Artifact`

NewArtifactWithDefaults instantiates a new Artifact object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFile

`func (o *Artifact) GetFile() *os.File`

GetFile returns the File field if non-nil, zero value otherwise.

### GetFileOk

`func (o *Artifact) GetFileOk() (**os.File, bool)`

GetFileOk returns a tuple with the File field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFile

`func (o *Artifact) SetFile(v *os.File)`

SetFile sets File field to given value.


### GetSize

`func (o *Artifact) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *Artifact) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *Artifact) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *Artifact) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetMd5

`func (o *Artifact) GetMd5() string`

GetMd5 returns the Md5 field if non-nil, zero value otherwise.

### GetMd5Ok

`func (o *Artifact) GetMd5Ok() (*string, bool)`

GetMd5Ok returns a tuple with the Md5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMd5

`func (o *Artifact) SetMd5(v string)`

SetMd5 sets Md5 field to given value.

### HasMd5

`func (o *Artifact) HasMd5() bool`

HasMd5 returns a boolean if a field has been set.

### SetMd5Nil

`func (o *Artifact) SetMd5Nil(b bool)`

 SetMd5Nil sets the value for Md5 to be an explicit nil

### UnsetMd5
`func (o *Artifact) UnsetMd5()`

UnsetMd5 ensures that no value is present for Md5, not even an explicit nil
### GetSha1

`func (o *Artifact) GetSha1() string`

GetSha1 returns the Sha1 field if non-nil, zero value otherwise.

### GetSha1Ok

`func (o *Artifact) GetSha1Ok() (*string, bool)`

GetSha1Ok returns a tuple with the Sha1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha1

`func (o *Artifact) SetSha1(v string)`

SetSha1 sets Sha1 field to given value.

### HasSha1

`func (o *Artifact) HasSha1() bool`

HasSha1 returns a boolean if a field has been set.

### SetSha1Nil

`func (o *Artifact) SetSha1Nil(b bool)`

 SetSha1Nil sets the value for Sha1 to be an explicit nil

### UnsetSha1
`func (o *Artifact) UnsetSha1()`

UnsetSha1 ensures that no value is present for Sha1, not even an explicit nil
### GetSha224

`func (o *Artifact) GetSha224() string`

GetSha224 returns the Sha224 field if non-nil, zero value otherwise.

### GetSha224Ok

`func (o *Artifact) GetSha224Ok() (*string, bool)`

GetSha224Ok returns a tuple with the Sha224 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha224

`func (o *Artifact) SetSha224(v string)`

SetSha224 sets Sha224 field to given value.

### HasSha224

`func (o *Artifact) HasSha224() bool`

HasSha224 returns a boolean if a field has been set.

### SetSha224Nil

`func (o *Artifact) SetSha224Nil(b bool)`

 SetSha224Nil sets the value for Sha224 to be an explicit nil

### UnsetSha224
`func (o *Artifact) UnsetSha224()`

UnsetSha224 ensures that no value is present for Sha224, not even an explicit nil
### GetSha256

`func (o *Artifact) GetSha256() string`

GetSha256 returns the Sha256 field if non-nil, zero value otherwise.

### GetSha256Ok

`func (o *Artifact) GetSha256Ok() (*string, bool)`

GetSha256Ok returns a tuple with the Sha256 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha256

`func (o *Artifact) SetSha256(v string)`

SetSha256 sets Sha256 field to given value.

### HasSha256

`func (o *Artifact) HasSha256() bool`

HasSha256 returns a boolean if a field has been set.

### SetSha256Nil

`func (o *Artifact) SetSha256Nil(b bool)`

 SetSha256Nil sets the value for Sha256 to be an explicit nil

### UnsetSha256
`func (o *Artifact) UnsetSha256()`

UnsetSha256 ensures that no value is present for Sha256, not even an explicit nil
### GetSha384

`func (o *Artifact) GetSha384() string`

GetSha384 returns the Sha384 field if non-nil, zero value otherwise.

### GetSha384Ok

`func (o *Artifact) GetSha384Ok() (*string, bool)`

GetSha384Ok returns a tuple with the Sha384 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha384

`func (o *Artifact) SetSha384(v string)`

SetSha384 sets Sha384 field to given value.

### HasSha384

`func (o *Artifact) HasSha384() bool`

HasSha384 returns a boolean if a field has been set.

### SetSha384Nil

`func (o *Artifact) SetSha384Nil(b bool)`

 SetSha384Nil sets the value for Sha384 to be an explicit nil

### UnsetSha384
`func (o *Artifact) UnsetSha384()`

UnsetSha384 ensures that no value is present for Sha384, not even an explicit nil
### GetSha512

`func (o *Artifact) GetSha512() string`

GetSha512 returns the Sha512 field if non-nil, zero value otherwise.

### GetSha512Ok

`func (o *Artifact) GetSha512Ok() (*string, bool)`

GetSha512Ok returns a tuple with the Sha512 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha512

`func (o *Artifact) SetSha512(v string)`

SetSha512 sets Sha512 field to given value.

### HasSha512

`func (o *Artifact) HasSha512() bool`

HasSha512 returns a boolean if a field has been set.

### SetSha512Nil

`func (o *Artifact) SetSha512Nil(b bool)`

 SetSha512Nil sets the value for Sha512 to be an explicit nil

### UnsetSha512
`func (o *Artifact) UnsetSha512()`

UnsetSha512 ensures that no value is present for Sha512, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


