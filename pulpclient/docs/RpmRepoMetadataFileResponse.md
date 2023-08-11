# RpmRepoMetadataFileResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PulpHref** | Pointer to **string** |  | [optional] [readonly] 
**PulpCreated** | Pointer to **time.Time** | Timestamp of creation. | [optional] [readonly] 
**Md5** | Pointer to **string** | The MD5 checksum if available. | [optional] [readonly] 
**Sha1** | Pointer to **string** | The SHA-1 checksum if available. | [optional] [readonly] 
**Sha224** | Pointer to **string** | The SHA-224 checksum if available. | [optional] [readonly] 
**Sha256** | Pointer to **string** | The SHA-256 checksum if available. | [optional] [readonly] 
**Sha384** | Pointer to **string** | The SHA-384 checksum if available. | [optional] [readonly] 
**Sha512** | Pointer to **string** | The SHA-512 checksum if available. | [optional] [readonly] 
**Artifact** | Pointer to **string** | Artifact file representing the physical content | [optional] 
**RelativePath** | **string** | Relative path of the file. | 
**DataType** | **string** | Metadata type. | 
**ChecksumType** | **string** | Checksum type for the file. | 
**Checksum** | **string** | Checksum for the file. | 

## Methods

### NewRpmRepoMetadataFileResponse

`func NewRpmRepoMetadataFileResponse(relativePath string, dataType string, checksumType string, checksum string, ) *RpmRepoMetadataFileResponse`

NewRpmRepoMetadataFileResponse instantiates a new RpmRepoMetadataFileResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRpmRepoMetadataFileResponseWithDefaults

`func NewRpmRepoMetadataFileResponseWithDefaults() *RpmRepoMetadataFileResponse`

NewRpmRepoMetadataFileResponseWithDefaults instantiates a new RpmRepoMetadataFileResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPulpHref

`func (o *RpmRepoMetadataFileResponse) GetPulpHref() string`

GetPulpHref returns the PulpHref field if non-nil, zero value otherwise.

### GetPulpHrefOk

`func (o *RpmRepoMetadataFileResponse) GetPulpHrefOk() (*string, bool)`

GetPulpHrefOk returns a tuple with the PulpHref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpHref

`func (o *RpmRepoMetadataFileResponse) SetPulpHref(v string)`

SetPulpHref sets PulpHref field to given value.

### HasPulpHref

`func (o *RpmRepoMetadataFileResponse) HasPulpHref() bool`

HasPulpHref returns a boolean if a field has been set.

### GetPulpCreated

`func (o *RpmRepoMetadataFileResponse) GetPulpCreated() time.Time`

GetPulpCreated returns the PulpCreated field if non-nil, zero value otherwise.

### GetPulpCreatedOk

`func (o *RpmRepoMetadataFileResponse) GetPulpCreatedOk() (*time.Time, bool)`

GetPulpCreatedOk returns a tuple with the PulpCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpCreated

`func (o *RpmRepoMetadataFileResponse) SetPulpCreated(v time.Time)`

SetPulpCreated sets PulpCreated field to given value.

### HasPulpCreated

`func (o *RpmRepoMetadataFileResponse) HasPulpCreated() bool`

HasPulpCreated returns a boolean if a field has been set.

### GetMd5

`func (o *RpmRepoMetadataFileResponse) GetMd5() string`

GetMd5 returns the Md5 field if non-nil, zero value otherwise.

### GetMd5Ok

`func (o *RpmRepoMetadataFileResponse) GetMd5Ok() (*string, bool)`

GetMd5Ok returns a tuple with the Md5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMd5

`func (o *RpmRepoMetadataFileResponse) SetMd5(v string)`

SetMd5 sets Md5 field to given value.

### HasMd5

`func (o *RpmRepoMetadataFileResponse) HasMd5() bool`

HasMd5 returns a boolean if a field has been set.

### GetSha1

`func (o *RpmRepoMetadataFileResponse) GetSha1() string`

GetSha1 returns the Sha1 field if non-nil, zero value otherwise.

### GetSha1Ok

`func (o *RpmRepoMetadataFileResponse) GetSha1Ok() (*string, bool)`

GetSha1Ok returns a tuple with the Sha1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha1

`func (o *RpmRepoMetadataFileResponse) SetSha1(v string)`

SetSha1 sets Sha1 field to given value.

### HasSha1

`func (o *RpmRepoMetadataFileResponse) HasSha1() bool`

HasSha1 returns a boolean if a field has been set.

### GetSha224

`func (o *RpmRepoMetadataFileResponse) GetSha224() string`

GetSha224 returns the Sha224 field if non-nil, zero value otherwise.

### GetSha224Ok

`func (o *RpmRepoMetadataFileResponse) GetSha224Ok() (*string, bool)`

GetSha224Ok returns a tuple with the Sha224 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha224

`func (o *RpmRepoMetadataFileResponse) SetSha224(v string)`

SetSha224 sets Sha224 field to given value.

### HasSha224

`func (o *RpmRepoMetadataFileResponse) HasSha224() bool`

HasSha224 returns a boolean if a field has been set.

### GetSha256

`func (o *RpmRepoMetadataFileResponse) GetSha256() string`

GetSha256 returns the Sha256 field if non-nil, zero value otherwise.

### GetSha256Ok

`func (o *RpmRepoMetadataFileResponse) GetSha256Ok() (*string, bool)`

GetSha256Ok returns a tuple with the Sha256 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha256

`func (o *RpmRepoMetadataFileResponse) SetSha256(v string)`

SetSha256 sets Sha256 field to given value.

### HasSha256

`func (o *RpmRepoMetadataFileResponse) HasSha256() bool`

HasSha256 returns a boolean if a field has been set.

### GetSha384

`func (o *RpmRepoMetadataFileResponse) GetSha384() string`

GetSha384 returns the Sha384 field if non-nil, zero value otherwise.

### GetSha384Ok

`func (o *RpmRepoMetadataFileResponse) GetSha384Ok() (*string, bool)`

GetSha384Ok returns a tuple with the Sha384 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha384

`func (o *RpmRepoMetadataFileResponse) SetSha384(v string)`

SetSha384 sets Sha384 field to given value.

### HasSha384

`func (o *RpmRepoMetadataFileResponse) HasSha384() bool`

HasSha384 returns a boolean if a field has been set.

### GetSha512

`func (o *RpmRepoMetadataFileResponse) GetSha512() string`

GetSha512 returns the Sha512 field if non-nil, zero value otherwise.

### GetSha512Ok

`func (o *RpmRepoMetadataFileResponse) GetSha512Ok() (*string, bool)`

GetSha512Ok returns a tuple with the Sha512 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha512

`func (o *RpmRepoMetadataFileResponse) SetSha512(v string)`

SetSha512 sets Sha512 field to given value.

### HasSha512

`func (o *RpmRepoMetadataFileResponse) HasSha512() bool`

HasSha512 returns a boolean if a field has been set.

### GetArtifact

`func (o *RpmRepoMetadataFileResponse) GetArtifact() string`

GetArtifact returns the Artifact field if non-nil, zero value otherwise.

### GetArtifactOk

`func (o *RpmRepoMetadataFileResponse) GetArtifactOk() (*string, bool)`

GetArtifactOk returns a tuple with the Artifact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifact

`func (o *RpmRepoMetadataFileResponse) SetArtifact(v string)`

SetArtifact sets Artifact field to given value.

### HasArtifact

`func (o *RpmRepoMetadataFileResponse) HasArtifact() bool`

HasArtifact returns a boolean if a field has been set.

### GetRelativePath

`func (o *RpmRepoMetadataFileResponse) GetRelativePath() string`

GetRelativePath returns the RelativePath field if non-nil, zero value otherwise.

### GetRelativePathOk

`func (o *RpmRepoMetadataFileResponse) GetRelativePathOk() (*string, bool)`

GetRelativePathOk returns a tuple with the RelativePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelativePath

`func (o *RpmRepoMetadataFileResponse) SetRelativePath(v string)`

SetRelativePath sets RelativePath field to given value.


### GetDataType

`func (o *RpmRepoMetadataFileResponse) GetDataType() string`

GetDataType returns the DataType field if non-nil, zero value otherwise.

### GetDataTypeOk

`func (o *RpmRepoMetadataFileResponse) GetDataTypeOk() (*string, bool)`

GetDataTypeOk returns a tuple with the DataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataType

`func (o *RpmRepoMetadataFileResponse) SetDataType(v string)`

SetDataType sets DataType field to given value.


### GetChecksumType

`func (o *RpmRepoMetadataFileResponse) GetChecksumType() string`

GetChecksumType returns the ChecksumType field if non-nil, zero value otherwise.

### GetChecksumTypeOk

`func (o *RpmRepoMetadataFileResponse) GetChecksumTypeOk() (*string, bool)`

GetChecksumTypeOk returns a tuple with the ChecksumType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChecksumType

`func (o *RpmRepoMetadataFileResponse) SetChecksumType(v string)`

SetChecksumType sets ChecksumType field to given value.


### GetChecksum

`func (o *RpmRepoMetadataFileResponse) GetChecksum() string`

GetChecksum returns the Checksum field if non-nil, zero value otherwise.

### GetChecksumOk

`func (o *RpmRepoMetadataFileResponse) GetChecksumOk() (*string, bool)`

GetChecksumOk returns a tuple with the Checksum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChecksum

`func (o *RpmRepoMetadataFileResponse) SetChecksum(v string)`

SetChecksum sets Checksum field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


