# AnsibleAnsibleRepository

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PulpLabels** | Pointer to **map[string]string** |  | [optional] 
**Name** | **string** | A unique name for this repository. | 
**Description** | Pointer to **NullableString** | An optional description. | [optional] 
**RetainRepoVersions** | Pointer to **NullableInt64** | Retain X versions of the repository. Default is null which retains all versions. | [optional] 
**Remote** | Pointer to **NullableString** | An optional remote to use by default when syncing. | [optional] 
**LastSyncedMetadataTime** | Pointer to **NullableTime** | Last synced metadata time. | [optional] 
**Gpgkey** | Pointer to **NullableString** | Gpg public key to verify collection signatures against | [optional] 
**Private** | Pointer to **bool** |  | [optional] 

## Methods

### NewAnsibleAnsibleRepository

`func NewAnsibleAnsibleRepository(name string, ) *AnsibleAnsibleRepository`

NewAnsibleAnsibleRepository instantiates a new AnsibleAnsibleRepository object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnsibleAnsibleRepositoryWithDefaults

`func NewAnsibleAnsibleRepositoryWithDefaults() *AnsibleAnsibleRepository`

NewAnsibleAnsibleRepositoryWithDefaults instantiates a new AnsibleAnsibleRepository object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPulpLabels

`func (o *AnsibleAnsibleRepository) GetPulpLabels() map[string]string`

GetPulpLabels returns the PulpLabels field if non-nil, zero value otherwise.

### GetPulpLabelsOk

`func (o *AnsibleAnsibleRepository) GetPulpLabelsOk() (*map[string]string, bool)`

GetPulpLabelsOk returns a tuple with the PulpLabels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpLabels

`func (o *AnsibleAnsibleRepository) SetPulpLabels(v map[string]string)`

SetPulpLabels sets PulpLabels field to given value.

### HasPulpLabels

`func (o *AnsibleAnsibleRepository) HasPulpLabels() bool`

HasPulpLabels returns a boolean if a field has been set.

### GetName

`func (o *AnsibleAnsibleRepository) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AnsibleAnsibleRepository) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AnsibleAnsibleRepository) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *AnsibleAnsibleRepository) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AnsibleAnsibleRepository) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AnsibleAnsibleRepository) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AnsibleAnsibleRepository) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *AnsibleAnsibleRepository) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *AnsibleAnsibleRepository) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetRetainRepoVersions

`func (o *AnsibleAnsibleRepository) GetRetainRepoVersions() int64`

GetRetainRepoVersions returns the RetainRepoVersions field if non-nil, zero value otherwise.

### GetRetainRepoVersionsOk

`func (o *AnsibleAnsibleRepository) GetRetainRepoVersionsOk() (*int64, bool)`

GetRetainRepoVersionsOk returns a tuple with the RetainRepoVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetainRepoVersions

`func (o *AnsibleAnsibleRepository) SetRetainRepoVersions(v int64)`

SetRetainRepoVersions sets RetainRepoVersions field to given value.

### HasRetainRepoVersions

`func (o *AnsibleAnsibleRepository) HasRetainRepoVersions() bool`

HasRetainRepoVersions returns a boolean if a field has been set.

### SetRetainRepoVersionsNil

`func (o *AnsibleAnsibleRepository) SetRetainRepoVersionsNil(b bool)`

 SetRetainRepoVersionsNil sets the value for RetainRepoVersions to be an explicit nil

### UnsetRetainRepoVersions
`func (o *AnsibleAnsibleRepository) UnsetRetainRepoVersions()`

UnsetRetainRepoVersions ensures that no value is present for RetainRepoVersions, not even an explicit nil
### GetRemote

`func (o *AnsibleAnsibleRepository) GetRemote() string`

GetRemote returns the Remote field if non-nil, zero value otherwise.

### GetRemoteOk

`func (o *AnsibleAnsibleRepository) GetRemoteOk() (*string, bool)`

GetRemoteOk returns a tuple with the Remote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemote

`func (o *AnsibleAnsibleRepository) SetRemote(v string)`

SetRemote sets Remote field to given value.

### HasRemote

`func (o *AnsibleAnsibleRepository) HasRemote() bool`

HasRemote returns a boolean if a field has been set.

### SetRemoteNil

`func (o *AnsibleAnsibleRepository) SetRemoteNil(b bool)`

 SetRemoteNil sets the value for Remote to be an explicit nil

### UnsetRemote
`func (o *AnsibleAnsibleRepository) UnsetRemote()`

UnsetRemote ensures that no value is present for Remote, not even an explicit nil
### GetLastSyncedMetadataTime

`func (o *AnsibleAnsibleRepository) GetLastSyncedMetadataTime() time.Time`

GetLastSyncedMetadataTime returns the LastSyncedMetadataTime field if non-nil, zero value otherwise.

### GetLastSyncedMetadataTimeOk

`func (o *AnsibleAnsibleRepository) GetLastSyncedMetadataTimeOk() (*time.Time, bool)`

GetLastSyncedMetadataTimeOk returns a tuple with the LastSyncedMetadataTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSyncedMetadataTime

`func (o *AnsibleAnsibleRepository) SetLastSyncedMetadataTime(v time.Time)`

SetLastSyncedMetadataTime sets LastSyncedMetadataTime field to given value.

### HasLastSyncedMetadataTime

`func (o *AnsibleAnsibleRepository) HasLastSyncedMetadataTime() bool`

HasLastSyncedMetadataTime returns a boolean if a field has been set.

### SetLastSyncedMetadataTimeNil

`func (o *AnsibleAnsibleRepository) SetLastSyncedMetadataTimeNil(b bool)`

 SetLastSyncedMetadataTimeNil sets the value for LastSyncedMetadataTime to be an explicit nil

### UnsetLastSyncedMetadataTime
`func (o *AnsibleAnsibleRepository) UnsetLastSyncedMetadataTime()`

UnsetLastSyncedMetadataTime ensures that no value is present for LastSyncedMetadataTime, not even an explicit nil
### GetGpgkey

`func (o *AnsibleAnsibleRepository) GetGpgkey() string`

GetGpgkey returns the Gpgkey field if non-nil, zero value otherwise.

### GetGpgkeyOk

`func (o *AnsibleAnsibleRepository) GetGpgkeyOk() (*string, bool)`

GetGpgkeyOk returns a tuple with the Gpgkey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpgkey

`func (o *AnsibleAnsibleRepository) SetGpgkey(v string)`

SetGpgkey sets Gpgkey field to given value.

### HasGpgkey

`func (o *AnsibleAnsibleRepository) HasGpgkey() bool`

HasGpgkey returns a boolean if a field has been set.

### SetGpgkeyNil

`func (o *AnsibleAnsibleRepository) SetGpgkeyNil(b bool)`

 SetGpgkeyNil sets the value for Gpgkey to be an explicit nil

### UnsetGpgkey
`func (o *AnsibleAnsibleRepository) UnsetGpgkey()`

UnsetGpgkey ensures that no value is present for Gpgkey, not even an explicit nil
### GetPrivate

`func (o *AnsibleAnsibleRepository) GetPrivate() bool`

GetPrivate returns the Private field if non-nil, zero value otherwise.

### GetPrivateOk

`func (o *AnsibleAnsibleRepository) GetPrivateOk() (*bool, bool)`

GetPrivateOk returns a tuple with the Private field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivate

`func (o *AnsibleAnsibleRepository) SetPrivate(v bool)`

SetPrivate sets Private field to given value.

### HasPrivate

`func (o *AnsibleAnsibleRepository) HasPrivate() bool`

HasPrivate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


