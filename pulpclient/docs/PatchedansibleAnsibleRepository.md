# PatchedansibleAnsibleRepository

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PulpLabels** | Pointer to **map[string]string** |  | [optional] 
**Name** | Pointer to **string** | A unique name for this repository. | [optional] 
**Description** | Pointer to **NullableString** | An optional description. | [optional] 
**RetainRepoVersions** | Pointer to **NullableInt64** | Retain X versions of the repository. Default is null which retains all versions. | [optional] 
**Remote** | Pointer to **NullableString** | An optional remote to use by default when syncing. | [optional] 
**LastSyncedMetadataTime** | Pointer to **NullableTime** | Last synced metadata time. | [optional] 
**Gpgkey** | Pointer to **NullableString** | Gpg public key to verify collection signatures against | [optional] 
**Private** | Pointer to **bool** |  | [optional] 

## Methods

### NewPatchedansibleAnsibleRepository

`func NewPatchedansibleAnsibleRepository() *PatchedansibleAnsibleRepository`

NewPatchedansibleAnsibleRepository instantiates a new PatchedansibleAnsibleRepository object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedansibleAnsibleRepositoryWithDefaults

`func NewPatchedansibleAnsibleRepositoryWithDefaults() *PatchedansibleAnsibleRepository`

NewPatchedansibleAnsibleRepositoryWithDefaults instantiates a new PatchedansibleAnsibleRepository object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPulpLabels

`func (o *PatchedansibleAnsibleRepository) GetPulpLabels() map[string]string`

GetPulpLabels returns the PulpLabels field if non-nil, zero value otherwise.

### GetPulpLabelsOk

`func (o *PatchedansibleAnsibleRepository) GetPulpLabelsOk() (*map[string]string, bool)`

GetPulpLabelsOk returns a tuple with the PulpLabels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpLabels

`func (o *PatchedansibleAnsibleRepository) SetPulpLabels(v map[string]string)`

SetPulpLabels sets PulpLabels field to given value.

### HasPulpLabels

`func (o *PatchedansibleAnsibleRepository) HasPulpLabels() bool`

HasPulpLabels returns a boolean if a field has been set.

### GetName

`func (o *PatchedansibleAnsibleRepository) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedansibleAnsibleRepository) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedansibleAnsibleRepository) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedansibleAnsibleRepository) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *PatchedansibleAnsibleRepository) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchedansibleAnsibleRepository) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchedansibleAnsibleRepository) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchedansibleAnsibleRepository) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *PatchedansibleAnsibleRepository) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *PatchedansibleAnsibleRepository) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetRetainRepoVersions

`func (o *PatchedansibleAnsibleRepository) GetRetainRepoVersions() int64`

GetRetainRepoVersions returns the RetainRepoVersions field if non-nil, zero value otherwise.

### GetRetainRepoVersionsOk

`func (o *PatchedansibleAnsibleRepository) GetRetainRepoVersionsOk() (*int64, bool)`

GetRetainRepoVersionsOk returns a tuple with the RetainRepoVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetainRepoVersions

`func (o *PatchedansibleAnsibleRepository) SetRetainRepoVersions(v int64)`

SetRetainRepoVersions sets RetainRepoVersions field to given value.

### HasRetainRepoVersions

`func (o *PatchedansibleAnsibleRepository) HasRetainRepoVersions() bool`

HasRetainRepoVersions returns a boolean if a field has been set.

### SetRetainRepoVersionsNil

`func (o *PatchedansibleAnsibleRepository) SetRetainRepoVersionsNil(b bool)`

 SetRetainRepoVersionsNil sets the value for RetainRepoVersions to be an explicit nil

### UnsetRetainRepoVersions
`func (o *PatchedansibleAnsibleRepository) UnsetRetainRepoVersions()`

UnsetRetainRepoVersions ensures that no value is present for RetainRepoVersions, not even an explicit nil
### GetRemote

`func (o *PatchedansibleAnsibleRepository) GetRemote() string`

GetRemote returns the Remote field if non-nil, zero value otherwise.

### GetRemoteOk

`func (o *PatchedansibleAnsibleRepository) GetRemoteOk() (*string, bool)`

GetRemoteOk returns a tuple with the Remote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemote

`func (o *PatchedansibleAnsibleRepository) SetRemote(v string)`

SetRemote sets Remote field to given value.

### HasRemote

`func (o *PatchedansibleAnsibleRepository) HasRemote() bool`

HasRemote returns a boolean if a field has been set.

### SetRemoteNil

`func (o *PatchedansibleAnsibleRepository) SetRemoteNil(b bool)`

 SetRemoteNil sets the value for Remote to be an explicit nil

### UnsetRemote
`func (o *PatchedansibleAnsibleRepository) UnsetRemote()`

UnsetRemote ensures that no value is present for Remote, not even an explicit nil
### GetLastSyncedMetadataTime

`func (o *PatchedansibleAnsibleRepository) GetLastSyncedMetadataTime() time.Time`

GetLastSyncedMetadataTime returns the LastSyncedMetadataTime field if non-nil, zero value otherwise.

### GetLastSyncedMetadataTimeOk

`func (o *PatchedansibleAnsibleRepository) GetLastSyncedMetadataTimeOk() (*time.Time, bool)`

GetLastSyncedMetadataTimeOk returns a tuple with the LastSyncedMetadataTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSyncedMetadataTime

`func (o *PatchedansibleAnsibleRepository) SetLastSyncedMetadataTime(v time.Time)`

SetLastSyncedMetadataTime sets LastSyncedMetadataTime field to given value.

### HasLastSyncedMetadataTime

`func (o *PatchedansibleAnsibleRepository) HasLastSyncedMetadataTime() bool`

HasLastSyncedMetadataTime returns a boolean if a field has been set.

### SetLastSyncedMetadataTimeNil

`func (o *PatchedansibleAnsibleRepository) SetLastSyncedMetadataTimeNil(b bool)`

 SetLastSyncedMetadataTimeNil sets the value for LastSyncedMetadataTime to be an explicit nil

### UnsetLastSyncedMetadataTime
`func (o *PatchedansibleAnsibleRepository) UnsetLastSyncedMetadataTime()`

UnsetLastSyncedMetadataTime ensures that no value is present for LastSyncedMetadataTime, not even an explicit nil
### GetGpgkey

`func (o *PatchedansibleAnsibleRepository) GetGpgkey() string`

GetGpgkey returns the Gpgkey field if non-nil, zero value otherwise.

### GetGpgkeyOk

`func (o *PatchedansibleAnsibleRepository) GetGpgkeyOk() (*string, bool)`

GetGpgkeyOk returns a tuple with the Gpgkey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpgkey

`func (o *PatchedansibleAnsibleRepository) SetGpgkey(v string)`

SetGpgkey sets Gpgkey field to given value.

### HasGpgkey

`func (o *PatchedansibleAnsibleRepository) HasGpgkey() bool`

HasGpgkey returns a boolean if a field has been set.

### SetGpgkeyNil

`func (o *PatchedansibleAnsibleRepository) SetGpgkeyNil(b bool)`

 SetGpgkeyNil sets the value for Gpgkey to be an explicit nil

### UnsetGpgkey
`func (o *PatchedansibleAnsibleRepository) UnsetGpgkey()`

UnsetGpgkey ensures that no value is present for Gpgkey, not even an explicit nil
### GetPrivate

`func (o *PatchedansibleAnsibleRepository) GetPrivate() bool`

GetPrivate returns the Private field if non-nil, zero value otherwise.

### GetPrivateOk

`func (o *PatchedansibleAnsibleRepository) GetPrivateOk() (*bool, bool)`

GetPrivateOk returns a tuple with the Private field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivate

`func (o *PatchedansibleAnsibleRepository) SetPrivate(v bool)`

SetPrivate sets Private field to given value.

### HasPrivate

`func (o *PatchedansibleAnsibleRepository) HasPrivate() bool`

HasPrivate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


