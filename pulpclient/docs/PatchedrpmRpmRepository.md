# PatchedrpmRpmRepository

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PulpLabels** | Pointer to **map[string]string** |  | [optional] 
**Name** | Pointer to **string** | A unique name for this repository. | [optional] 
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

### NewPatchedrpmRpmRepository

`func NewPatchedrpmRpmRepository() *PatchedrpmRpmRepository`

NewPatchedrpmRpmRepository instantiates a new PatchedrpmRpmRepository object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedrpmRpmRepositoryWithDefaults

`func NewPatchedrpmRpmRepositoryWithDefaults() *PatchedrpmRpmRepository`

NewPatchedrpmRpmRepositoryWithDefaults instantiates a new PatchedrpmRpmRepository object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPulpLabels

`func (o *PatchedrpmRpmRepository) GetPulpLabels() map[string]string`

GetPulpLabels returns the PulpLabels field if non-nil, zero value otherwise.

### GetPulpLabelsOk

`func (o *PatchedrpmRpmRepository) GetPulpLabelsOk() (*map[string]string, bool)`

GetPulpLabelsOk returns a tuple with the PulpLabels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpLabels

`func (o *PatchedrpmRpmRepository) SetPulpLabels(v map[string]string)`

SetPulpLabels sets PulpLabels field to given value.

### HasPulpLabels

`func (o *PatchedrpmRpmRepository) HasPulpLabels() bool`

HasPulpLabels returns a boolean if a field has been set.

### GetName

`func (o *PatchedrpmRpmRepository) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedrpmRpmRepository) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedrpmRpmRepository) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedrpmRpmRepository) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *PatchedrpmRpmRepository) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchedrpmRpmRepository) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchedrpmRpmRepository) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchedrpmRpmRepository) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *PatchedrpmRpmRepository) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *PatchedrpmRpmRepository) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetRetainRepoVersions

`func (o *PatchedrpmRpmRepository) GetRetainRepoVersions() int64`

GetRetainRepoVersions returns the RetainRepoVersions field if non-nil, zero value otherwise.

### GetRetainRepoVersionsOk

`func (o *PatchedrpmRpmRepository) GetRetainRepoVersionsOk() (*int64, bool)`

GetRetainRepoVersionsOk returns a tuple with the RetainRepoVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetainRepoVersions

`func (o *PatchedrpmRpmRepository) SetRetainRepoVersions(v int64)`

SetRetainRepoVersions sets RetainRepoVersions field to given value.

### HasRetainRepoVersions

`func (o *PatchedrpmRpmRepository) HasRetainRepoVersions() bool`

HasRetainRepoVersions returns a boolean if a field has been set.

### SetRetainRepoVersionsNil

`func (o *PatchedrpmRpmRepository) SetRetainRepoVersionsNil(b bool)`

 SetRetainRepoVersionsNil sets the value for RetainRepoVersions to be an explicit nil

### UnsetRetainRepoVersions
`func (o *PatchedrpmRpmRepository) UnsetRetainRepoVersions()`

UnsetRetainRepoVersions ensures that no value is present for RetainRepoVersions, not even an explicit nil
### GetRemote

`func (o *PatchedrpmRpmRepository) GetRemote() string`

GetRemote returns the Remote field if non-nil, zero value otherwise.

### GetRemoteOk

`func (o *PatchedrpmRpmRepository) GetRemoteOk() (*string, bool)`

GetRemoteOk returns a tuple with the Remote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemote

`func (o *PatchedrpmRpmRepository) SetRemote(v string)`

SetRemote sets Remote field to given value.

### HasRemote

`func (o *PatchedrpmRpmRepository) HasRemote() bool`

HasRemote returns a boolean if a field has been set.

### SetRemoteNil

`func (o *PatchedrpmRpmRepository) SetRemoteNil(b bool)`

 SetRemoteNil sets the value for Remote to be an explicit nil

### UnsetRemote
`func (o *PatchedrpmRpmRepository) UnsetRemote()`

UnsetRemote ensures that no value is present for Remote, not even an explicit nil
### GetAutopublish

`func (o *PatchedrpmRpmRepository) GetAutopublish() bool`

GetAutopublish returns the Autopublish field if non-nil, zero value otherwise.

### GetAutopublishOk

`func (o *PatchedrpmRpmRepository) GetAutopublishOk() (*bool, bool)`

GetAutopublishOk returns a tuple with the Autopublish field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutopublish

`func (o *PatchedrpmRpmRepository) SetAutopublish(v bool)`

SetAutopublish sets Autopublish field to given value.

### HasAutopublish

`func (o *PatchedrpmRpmRepository) HasAutopublish() bool`

HasAutopublish returns a boolean if a field has been set.

### GetMetadataSigningService

`func (o *PatchedrpmRpmRepository) GetMetadataSigningService() string`

GetMetadataSigningService returns the MetadataSigningService field if non-nil, zero value otherwise.

### GetMetadataSigningServiceOk

`func (o *PatchedrpmRpmRepository) GetMetadataSigningServiceOk() (*string, bool)`

GetMetadataSigningServiceOk returns a tuple with the MetadataSigningService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataSigningService

`func (o *PatchedrpmRpmRepository) SetMetadataSigningService(v string)`

SetMetadataSigningService sets MetadataSigningService field to given value.

### HasMetadataSigningService

`func (o *PatchedrpmRpmRepository) HasMetadataSigningService() bool`

HasMetadataSigningService returns a boolean if a field has been set.

### SetMetadataSigningServiceNil

`func (o *PatchedrpmRpmRepository) SetMetadataSigningServiceNil(b bool)`

 SetMetadataSigningServiceNil sets the value for MetadataSigningService to be an explicit nil

### UnsetMetadataSigningService
`func (o *PatchedrpmRpmRepository) UnsetMetadataSigningService()`

UnsetMetadataSigningService ensures that no value is present for MetadataSigningService, not even an explicit nil
### GetRetainPackageVersions

`func (o *PatchedrpmRpmRepository) GetRetainPackageVersions() int64`

GetRetainPackageVersions returns the RetainPackageVersions field if non-nil, zero value otherwise.

### GetRetainPackageVersionsOk

`func (o *PatchedrpmRpmRepository) GetRetainPackageVersionsOk() (*int64, bool)`

GetRetainPackageVersionsOk returns a tuple with the RetainPackageVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetainPackageVersions

`func (o *PatchedrpmRpmRepository) SetRetainPackageVersions(v int64)`

SetRetainPackageVersions sets RetainPackageVersions field to given value.

### HasRetainPackageVersions

`func (o *PatchedrpmRpmRepository) HasRetainPackageVersions() bool`

HasRetainPackageVersions returns a boolean if a field has been set.

### GetMetadataChecksumType

`func (o *PatchedrpmRpmRepository) GetMetadataChecksumType() MetadataChecksumTypeEnum`

GetMetadataChecksumType returns the MetadataChecksumType field if non-nil, zero value otherwise.

### GetMetadataChecksumTypeOk

`func (o *PatchedrpmRpmRepository) GetMetadataChecksumTypeOk() (*MetadataChecksumTypeEnum, bool)`

GetMetadataChecksumTypeOk returns a tuple with the MetadataChecksumType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataChecksumType

`func (o *PatchedrpmRpmRepository) SetMetadataChecksumType(v MetadataChecksumTypeEnum)`

SetMetadataChecksumType sets MetadataChecksumType field to given value.

### HasMetadataChecksumType

`func (o *PatchedrpmRpmRepository) HasMetadataChecksumType() bool`

HasMetadataChecksumType returns a boolean if a field has been set.

### SetMetadataChecksumTypeNil

`func (o *PatchedrpmRpmRepository) SetMetadataChecksumTypeNil(b bool)`

 SetMetadataChecksumTypeNil sets the value for MetadataChecksumType to be an explicit nil

### UnsetMetadataChecksumType
`func (o *PatchedrpmRpmRepository) UnsetMetadataChecksumType()`

UnsetMetadataChecksumType ensures that no value is present for MetadataChecksumType, not even an explicit nil
### GetPackageChecksumType

`func (o *PatchedrpmRpmRepository) GetPackageChecksumType() PackageChecksumTypeEnum`

GetPackageChecksumType returns the PackageChecksumType field if non-nil, zero value otherwise.

### GetPackageChecksumTypeOk

`func (o *PatchedrpmRpmRepository) GetPackageChecksumTypeOk() (*PackageChecksumTypeEnum, bool)`

GetPackageChecksumTypeOk returns a tuple with the PackageChecksumType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageChecksumType

`func (o *PatchedrpmRpmRepository) SetPackageChecksumType(v PackageChecksumTypeEnum)`

SetPackageChecksumType sets PackageChecksumType field to given value.

### HasPackageChecksumType

`func (o *PatchedrpmRpmRepository) HasPackageChecksumType() bool`

HasPackageChecksumType returns a boolean if a field has been set.

### SetPackageChecksumTypeNil

`func (o *PatchedrpmRpmRepository) SetPackageChecksumTypeNil(b bool)`

 SetPackageChecksumTypeNil sets the value for PackageChecksumType to be an explicit nil

### UnsetPackageChecksumType
`func (o *PatchedrpmRpmRepository) UnsetPackageChecksumType()`

UnsetPackageChecksumType ensures that no value is present for PackageChecksumType, not even an explicit nil
### GetGpgcheck

`func (o *PatchedrpmRpmRepository) GetGpgcheck() int64`

GetGpgcheck returns the Gpgcheck field if non-nil, zero value otherwise.

### GetGpgcheckOk

`func (o *PatchedrpmRpmRepository) GetGpgcheckOk() (*int64, bool)`

GetGpgcheckOk returns a tuple with the Gpgcheck field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpgcheck

`func (o *PatchedrpmRpmRepository) SetGpgcheck(v int64)`

SetGpgcheck sets Gpgcheck field to given value.

### HasGpgcheck

`func (o *PatchedrpmRpmRepository) HasGpgcheck() bool`

HasGpgcheck returns a boolean if a field has been set.

### GetRepoGpgcheck

`func (o *PatchedrpmRpmRepository) GetRepoGpgcheck() int64`

GetRepoGpgcheck returns the RepoGpgcheck field if non-nil, zero value otherwise.

### GetRepoGpgcheckOk

`func (o *PatchedrpmRpmRepository) GetRepoGpgcheckOk() (*int64, bool)`

GetRepoGpgcheckOk returns a tuple with the RepoGpgcheck field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepoGpgcheck

`func (o *PatchedrpmRpmRepository) SetRepoGpgcheck(v int64)`

SetRepoGpgcheck sets RepoGpgcheck field to given value.

### HasRepoGpgcheck

`func (o *PatchedrpmRpmRepository) HasRepoGpgcheck() bool`

HasRepoGpgcheck returns a boolean if a field has been set.

### GetSqliteMetadata

`func (o *PatchedrpmRpmRepository) GetSqliteMetadata() bool`

GetSqliteMetadata returns the SqliteMetadata field if non-nil, zero value otherwise.

### GetSqliteMetadataOk

`func (o *PatchedrpmRpmRepository) GetSqliteMetadataOk() (*bool, bool)`

GetSqliteMetadataOk returns a tuple with the SqliteMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSqliteMetadata

`func (o *PatchedrpmRpmRepository) SetSqliteMetadata(v bool)`

SetSqliteMetadata sets SqliteMetadata field to given value.

### HasSqliteMetadata

`func (o *PatchedrpmRpmRepository) HasSqliteMetadata() bool`

HasSqliteMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


