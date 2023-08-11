# RpmRpmPublication

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RepositoryVersion** | Pointer to **string** |  | [optional] 
**Repository** | Pointer to **string** | A URI of the repository to be published. | [optional] 
**MetadataChecksumType** | Pointer to [**MetadataChecksumTypeEnum**](MetadataChecksumTypeEnum.md) |  | [optional] 
**PackageChecksumType** | Pointer to [**PackageChecksumTypeEnum**](PackageChecksumTypeEnum.md) |  | [optional] 
**Gpgcheck** | Pointer to **int64** | An option specifying whether a client should perform a GPG signature check on packages. | [optional] 
**RepoGpgcheck** | Pointer to **int64** | An option specifying whether a client should perform a GPG signature check on the repodata. | [optional] 
**SqliteMetadata** | Pointer to **bool** | DEPRECATED: An option specifying whether Pulp should generate SQLite metadata. | [optional] [default to false]

## Methods

### NewRpmRpmPublication

`func NewRpmRpmPublication() *RpmRpmPublication`

NewRpmRpmPublication instantiates a new RpmRpmPublication object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRpmRpmPublicationWithDefaults

`func NewRpmRpmPublicationWithDefaults() *RpmRpmPublication`

NewRpmRpmPublicationWithDefaults instantiates a new RpmRpmPublication object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRepositoryVersion

`func (o *RpmRpmPublication) GetRepositoryVersion() string`

GetRepositoryVersion returns the RepositoryVersion field if non-nil, zero value otherwise.

### GetRepositoryVersionOk

`func (o *RpmRpmPublication) GetRepositoryVersionOk() (*string, bool)`

GetRepositoryVersionOk returns a tuple with the RepositoryVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepositoryVersion

`func (o *RpmRpmPublication) SetRepositoryVersion(v string)`

SetRepositoryVersion sets RepositoryVersion field to given value.

### HasRepositoryVersion

`func (o *RpmRpmPublication) HasRepositoryVersion() bool`

HasRepositoryVersion returns a boolean if a field has been set.

### GetRepository

`func (o *RpmRpmPublication) GetRepository() string`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *RpmRpmPublication) GetRepositoryOk() (*string, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *RpmRpmPublication) SetRepository(v string)`

SetRepository sets Repository field to given value.

### HasRepository

`func (o *RpmRpmPublication) HasRepository() bool`

HasRepository returns a boolean if a field has been set.

### GetMetadataChecksumType

`func (o *RpmRpmPublication) GetMetadataChecksumType() MetadataChecksumTypeEnum`

GetMetadataChecksumType returns the MetadataChecksumType field if non-nil, zero value otherwise.

### GetMetadataChecksumTypeOk

`func (o *RpmRpmPublication) GetMetadataChecksumTypeOk() (*MetadataChecksumTypeEnum, bool)`

GetMetadataChecksumTypeOk returns a tuple with the MetadataChecksumType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataChecksumType

`func (o *RpmRpmPublication) SetMetadataChecksumType(v MetadataChecksumTypeEnum)`

SetMetadataChecksumType sets MetadataChecksumType field to given value.

### HasMetadataChecksumType

`func (o *RpmRpmPublication) HasMetadataChecksumType() bool`

HasMetadataChecksumType returns a boolean if a field has been set.

### GetPackageChecksumType

`func (o *RpmRpmPublication) GetPackageChecksumType() PackageChecksumTypeEnum`

GetPackageChecksumType returns the PackageChecksumType field if non-nil, zero value otherwise.

### GetPackageChecksumTypeOk

`func (o *RpmRpmPublication) GetPackageChecksumTypeOk() (*PackageChecksumTypeEnum, bool)`

GetPackageChecksumTypeOk returns a tuple with the PackageChecksumType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageChecksumType

`func (o *RpmRpmPublication) SetPackageChecksumType(v PackageChecksumTypeEnum)`

SetPackageChecksumType sets PackageChecksumType field to given value.

### HasPackageChecksumType

`func (o *RpmRpmPublication) HasPackageChecksumType() bool`

HasPackageChecksumType returns a boolean if a field has been set.

### GetGpgcheck

`func (o *RpmRpmPublication) GetGpgcheck() int64`

GetGpgcheck returns the Gpgcheck field if non-nil, zero value otherwise.

### GetGpgcheckOk

`func (o *RpmRpmPublication) GetGpgcheckOk() (*int64, bool)`

GetGpgcheckOk returns a tuple with the Gpgcheck field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpgcheck

`func (o *RpmRpmPublication) SetGpgcheck(v int64)`

SetGpgcheck sets Gpgcheck field to given value.

### HasGpgcheck

`func (o *RpmRpmPublication) HasGpgcheck() bool`

HasGpgcheck returns a boolean if a field has been set.

### GetRepoGpgcheck

`func (o *RpmRpmPublication) GetRepoGpgcheck() int64`

GetRepoGpgcheck returns the RepoGpgcheck field if non-nil, zero value otherwise.

### GetRepoGpgcheckOk

`func (o *RpmRpmPublication) GetRepoGpgcheckOk() (*int64, bool)`

GetRepoGpgcheckOk returns a tuple with the RepoGpgcheck field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepoGpgcheck

`func (o *RpmRpmPublication) SetRepoGpgcheck(v int64)`

SetRepoGpgcheck sets RepoGpgcheck field to given value.

### HasRepoGpgcheck

`func (o *RpmRpmPublication) HasRepoGpgcheck() bool`

HasRepoGpgcheck returns a boolean if a field has been set.

### GetSqliteMetadata

`func (o *RpmRpmPublication) GetSqliteMetadata() bool`

GetSqliteMetadata returns the SqliteMetadata field if non-nil, zero value otherwise.

### GetSqliteMetadataOk

`func (o *RpmRpmPublication) GetSqliteMetadataOk() (*bool, bool)`

GetSqliteMetadataOk returns a tuple with the SqliteMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSqliteMetadata

`func (o *RpmRpmPublication) SetSqliteMetadata(v bool)`

SetSqliteMetadata sets SqliteMetadata field to given value.

### HasSqliteMetadata

`func (o *RpmRpmPublication) HasSqliteMetadata() bool`

HasSqliteMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


