# RpmRpmRepository

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PulpLabels** | Pointer to **map[string]string** |  | [optional] 
**Name** | **string** | A unique name for this repository. | 
**Description** | Pointer to **NullableString** | An optional description. | [optional] 
**RetainRepoVersions** | Pointer to **NullableInt64** | Retain X versions of the repository. Default is null which retains all versions. | [optional] 
**Remote** | Pointer to **NullableString** | An optional remote to use by default when syncing. | [optional] 
**Autopublish** | Pointer to **bool** | Whether to automatically create publications for new repository versions, and update any distributions pointing to this repository. | [optional] [default to false]
**MetadataSigningService** | Pointer to **NullableString** | A reference to an associated signing service. | [optional] 
**RetainPackageVersions** | Pointer to **int64** | The number of versions of each package to keep in the repository; older versions will be purged. The default is &#39;0&#39;, which will disable this feature and keep all versions of each package. | [optional] 
**MetadataChecksumType** | Pointer to [**NullableMetadataChecksumTypeEnum**](MetadataChecksumTypeEnum.md) |  | [optional] 
**PackageChecksumType** | Pointer to [**NullablePackageChecksumTypeEnum**](PackageChecksumTypeEnum.md) |  | [optional] 
**Gpgcheck** | Pointer to **int64** | An option specifying whether a client should perform a GPG signature check on packages. | [optional] [default to 0]
**RepoGpgcheck** | Pointer to **int64** | An option specifying whether a client should perform a GPG signature check on the repodata. | [optional] [default to 0]
**SqliteMetadata** | Pointer to **bool** | DEPRECATED: An option specifying whether Pulp should generate SQLite metadata. | [optional] [default to false]

## Methods

### NewRpmRpmRepository

`func NewRpmRpmRepository(name string, ) *RpmRpmRepository`

NewRpmRpmRepository instantiates a new RpmRpmRepository object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRpmRpmRepositoryWithDefaults

`func NewRpmRpmRepositoryWithDefaults() *RpmRpmRepository`

NewRpmRpmRepositoryWithDefaults instantiates a new RpmRpmRepository object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPulpLabels

`func (o *RpmRpmRepository) GetPulpLabels() map[string]string`

GetPulpLabels returns the PulpLabels field if non-nil, zero value otherwise.

### GetPulpLabelsOk

`func (o *RpmRpmRepository) GetPulpLabelsOk() (*map[string]string, bool)`

GetPulpLabelsOk returns a tuple with the PulpLabels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpLabels

`func (o *RpmRpmRepository) SetPulpLabels(v map[string]string)`

SetPulpLabels sets PulpLabels field to given value.

### HasPulpLabels

`func (o *RpmRpmRepository) HasPulpLabels() bool`

HasPulpLabels returns a boolean if a field has been set.

### GetName

`func (o *RpmRpmRepository) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RpmRpmRepository) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RpmRpmRepository) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *RpmRpmRepository) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RpmRpmRepository) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RpmRpmRepository) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RpmRpmRepository) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *RpmRpmRepository) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *RpmRpmRepository) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetRetainRepoVersions

`func (o *RpmRpmRepository) GetRetainRepoVersions() int64`

GetRetainRepoVersions returns the RetainRepoVersions field if non-nil, zero value otherwise.

### GetRetainRepoVersionsOk

`func (o *RpmRpmRepository) GetRetainRepoVersionsOk() (*int64, bool)`

GetRetainRepoVersionsOk returns a tuple with the RetainRepoVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetainRepoVersions

`func (o *RpmRpmRepository) SetRetainRepoVersions(v int64)`

SetRetainRepoVersions sets RetainRepoVersions field to given value.

### HasRetainRepoVersions

`func (o *RpmRpmRepository) HasRetainRepoVersions() bool`

HasRetainRepoVersions returns a boolean if a field has been set.

### SetRetainRepoVersionsNil

`func (o *RpmRpmRepository) SetRetainRepoVersionsNil(b bool)`

 SetRetainRepoVersionsNil sets the value for RetainRepoVersions to be an explicit nil

### UnsetRetainRepoVersions
`func (o *RpmRpmRepository) UnsetRetainRepoVersions()`

UnsetRetainRepoVersions ensures that no value is present for RetainRepoVersions, not even an explicit nil
### GetRemote

`func (o *RpmRpmRepository) GetRemote() string`

GetRemote returns the Remote field if non-nil, zero value otherwise.

### GetRemoteOk

`func (o *RpmRpmRepository) GetRemoteOk() (*string, bool)`

GetRemoteOk returns a tuple with the Remote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemote

`func (o *RpmRpmRepository) SetRemote(v string)`

SetRemote sets Remote field to given value.

### HasRemote

`func (o *RpmRpmRepository) HasRemote() bool`

HasRemote returns a boolean if a field has been set.

### SetRemoteNil

`func (o *RpmRpmRepository) SetRemoteNil(b bool)`

 SetRemoteNil sets the value for Remote to be an explicit nil

### UnsetRemote
`func (o *RpmRpmRepository) UnsetRemote()`

UnsetRemote ensures that no value is present for Remote, not even an explicit nil
### GetAutopublish

`func (o *RpmRpmRepository) GetAutopublish() bool`

GetAutopublish returns the Autopublish field if non-nil, zero value otherwise.

### GetAutopublishOk

`func (o *RpmRpmRepository) GetAutopublishOk() (*bool, bool)`

GetAutopublishOk returns a tuple with the Autopublish field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutopublish

`func (o *RpmRpmRepository) SetAutopublish(v bool)`

SetAutopublish sets Autopublish field to given value.

### HasAutopublish

`func (o *RpmRpmRepository) HasAutopublish() bool`

HasAutopublish returns a boolean if a field has been set.

### GetMetadataSigningService

`func (o *RpmRpmRepository) GetMetadataSigningService() string`

GetMetadataSigningService returns the MetadataSigningService field if non-nil, zero value otherwise.

### GetMetadataSigningServiceOk

`func (o *RpmRpmRepository) GetMetadataSigningServiceOk() (*string, bool)`

GetMetadataSigningServiceOk returns a tuple with the MetadataSigningService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataSigningService

`func (o *RpmRpmRepository) SetMetadataSigningService(v string)`

SetMetadataSigningService sets MetadataSigningService field to given value.

### HasMetadataSigningService

`func (o *RpmRpmRepository) HasMetadataSigningService() bool`

HasMetadataSigningService returns a boolean if a field has been set.

### SetMetadataSigningServiceNil

`func (o *RpmRpmRepository) SetMetadataSigningServiceNil(b bool)`

 SetMetadataSigningServiceNil sets the value for MetadataSigningService to be an explicit nil

### UnsetMetadataSigningService
`func (o *RpmRpmRepository) UnsetMetadataSigningService()`

UnsetMetadataSigningService ensures that no value is present for MetadataSigningService, not even an explicit nil
### GetRetainPackageVersions

`func (o *RpmRpmRepository) GetRetainPackageVersions() int64`

GetRetainPackageVersions returns the RetainPackageVersions field if non-nil, zero value otherwise.

### GetRetainPackageVersionsOk

`func (o *RpmRpmRepository) GetRetainPackageVersionsOk() (*int64, bool)`

GetRetainPackageVersionsOk returns a tuple with the RetainPackageVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetainPackageVersions

`func (o *RpmRpmRepository) SetRetainPackageVersions(v int64)`

SetRetainPackageVersions sets RetainPackageVersions field to given value.

### HasRetainPackageVersions

`func (o *RpmRpmRepository) HasRetainPackageVersions() bool`

HasRetainPackageVersions returns a boolean if a field has been set.

### GetMetadataChecksumType

`func (o *RpmRpmRepository) GetMetadataChecksumType() MetadataChecksumTypeEnum`

GetMetadataChecksumType returns the MetadataChecksumType field if non-nil, zero value otherwise.

### GetMetadataChecksumTypeOk

`func (o *RpmRpmRepository) GetMetadataChecksumTypeOk() (*MetadataChecksumTypeEnum, bool)`

GetMetadataChecksumTypeOk returns a tuple with the MetadataChecksumType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataChecksumType

`func (o *RpmRpmRepository) SetMetadataChecksumType(v MetadataChecksumTypeEnum)`

SetMetadataChecksumType sets MetadataChecksumType field to given value.

### HasMetadataChecksumType

`func (o *RpmRpmRepository) HasMetadataChecksumType() bool`

HasMetadataChecksumType returns a boolean if a field has been set.

### SetMetadataChecksumTypeNil

`func (o *RpmRpmRepository) SetMetadataChecksumTypeNil(b bool)`

 SetMetadataChecksumTypeNil sets the value for MetadataChecksumType to be an explicit nil

### UnsetMetadataChecksumType
`func (o *RpmRpmRepository) UnsetMetadataChecksumType()`

UnsetMetadataChecksumType ensures that no value is present for MetadataChecksumType, not even an explicit nil
### GetPackageChecksumType

`func (o *RpmRpmRepository) GetPackageChecksumType() PackageChecksumTypeEnum`

GetPackageChecksumType returns the PackageChecksumType field if non-nil, zero value otherwise.

### GetPackageChecksumTypeOk

`func (o *RpmRpmRepository) GetPackageChecksumTypeOk() (*PackageChecksumTypeEnum, bool)`

GetPackageChecksumTypeOk returns a tuple with the PackageChecksumType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageChecksumType

`func (o *RpmRpmRepository) SetPackageChecksumType(v PackageChecksumTypeEnum)`

SetPackageChecksumType sets PackageChecksumType field to given value.

### HasPackageChecksumType

`func (o *RpmRpmRepository) HasPackageChecksumType() bool`

HasPackageChecksumType returns a boolean if a field has been set.

### SetPackageChecksumTypeNil

`func (o *RpmRpmRepository) SetPackageChecksumTypeNil(b bool)`

 SetPackageChecksumTypeNil sets the value for PackageChecksumType to be an explicit nil

### UnsetPackageChecksumType
`func (o *RpmRpmRepository) UnsetPackageChecksumType()`

UnsetPackageChecksumType ensures that no value is present for PackageChecksumType, not even an explicit nil
### GetGpgcheck

`func (o *RpmRpmRepository) GetGpgcheck() int64`

GetGpgcheck returns the Gpgcheck field if non-nil, zero value otherwise.

### GetGpgcheckOk

`func (o *RpmRpmRepository) GetGpgcheckOk() (*int64, bool)`

GetGpgcheckOk returns a tuple with the Gpgcheck field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpgcheck

`func (o *RpmRpmRepository) SetGpgcheck(v int64)`

SetGpgcheck sets Gpgcheck field to given value.

### HasGpgcheck

`func (o *RpmRpmRepository) HasGpgcheck() bool`

HasGpgcheck returns a boolean if a field has been set.

### GetRepoGpgcheck

`func (o *RpmRpmRepository) GetRepoGpgcheck() int64`

GetRepoGpgcheck returns the RepoGpgcheck field if non-nil, zero value otherwise.

### GetRepoGpgcheckOk

`func (o *RpmRpmRepository) GetRepoGpgcheckOk() (*int64, bool)`

GetRepoGpgcheckOk returns a tuple with the RepoGpgcheck field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepoGpgcheck

`func (o *RpmRpmRepository) SetRepoGpgcheck(v int64)`

SetRepoGpgcheck sets RepoGpgcheck field to given value.

### HasRepoGpgcheck

`func (o *RpmRpmRepository) HasRepoGpgcheck() bool`

HasRepoGpgcheck returns a boolean if a field has been set.

### GetSqliteMetadata

`func (o *RpmRpmRepository) GetSqliteMetadata() bool`

GetSqliteMetadata returns the SqliteMetadata field if non-nil, zero value otherwise.

### GetSqliteMetadataOk

`func (o *RpmRpmRepository) GetSqliteMetadataOk() (*bool, bool)`

GetSqliteMetadataOk returns a tuple with the SqliteMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSqliteMetadata

`func (o *RpmRpmRepository) SetSqliteMetadata(v bool)`

SetSqliteMetadata sets SqliteMetadata field to given value.

### HasSqliteMetadata

`func (o *RpmRpmRepository) HasSqliteMetadata() bool`

HasSqliteMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


