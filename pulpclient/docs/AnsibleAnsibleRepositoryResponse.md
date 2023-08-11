# AnsibleAnsibleRepositoryResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PulpHref** | Pointer to **string** |  | [optional] [readonly] 
**PulpCreated** | Pointer to **time.Time** | Timestamp of creation. | [optional] [readonly] 
**VersionsHref** | Pointer to **string** |  | [optional] [readonly] 
**PulpLabels** | Pointer to **map[string]string** |  | [optional] 
**LatestVersionHref** | Pointer to **string** |  | [optional] [readonly] 
**Name** | **string** | A unique name for this repository. | 
**Description** | Pointer to **NullableString** | An optional description. | [optional] 
**RetainRepoVersions** | Pointer to **NullableInt64** | Retain X versions of the repository. Default is null which retains all versions. | [optional] 
**Remote** | Pointer to **NullableString** | An optional remote to use by default when syncing. | [optional] 
**LastSyncedMetadataTime** | Pointer to **NullableTime** | Last synced metadata time. | [optional] 
**Gpgkey** | Pointer to **NullableString** | Gpg public key to verify collection signatures against | [optional] 
**LastSyncTask** | Pointer to [**TaskResponse**](TaskResponse.md) |  | [optional] [readonly] 
**Private** | Pointer to **bool** |  | [optional] 

## Methods

### NewAnsibleAnsibleRepositoryResponse

`func NewAnsibleAnsibleRepositoryResponse(name string, ) *AnsibleAnsibleRepositoryResponse`

NewAnsibleAnsibleRepositoryResponse instantiates a new AnsibleAnsibleRepositoryResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnsibleAnsibleRepositoryResponseWithDefaults

`func NewAnsibleAnsibleRepositoryResponseWithDefaults() *AnsibleAnsibleRepositoryResponse`

NewAnsibleAnsibleRepositoryResponseWithDefaults instantiates a new AnsibleAnsibleRepositoryResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPulpHref

`func (o *AnsibleAnsibleRepositoryResponse) GetPulpHref() string`

GetPulpHref returns the PulpHref field if non-nil, zero value otherwise.

### GetPulpHrefOk

`func (o *AnsibleAnsibleRepositoryResponse) GetPulpHrefOk() (*string, bool)`

GetPulpHrefOk returns a tuple with the PulpHref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpHref

`func (o *AnsibleAnsibleRepositoryResponse) SetPulpHref(v string)`

SetPulpHref sets PulpHref field to given value.

### HasPulpHref

`func (o *AnsibleAnsibleRepositoryResponse) HasPulpHref() bool`

HasPulpHref returns a boolean if a field has been set.

### GetPulpCreated

`func (o *AnsibleAnsibleRepositoryResponse) GetPulpCreated() time.Time`

GetPulpCreated returns the PulpCreated field if non-nil, zero value otherwise.

### GetPulpCreatedOk

`func (o *AnsibleAnsibleRepositoryResponse) GetPulpCreatedOk() (*time.Time, bool)`

GetPulpCreatedOk returns a tuple with the PulpCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpCreated

`func (o *AnsibleAnsibleRepositoryResponse) SetPulpCreated(v time.Time)`

SetPulpCreated sets PulpCreated field to given value.

### HasPulpCreated

`func (o *AnsibleAnsibleRepositoryResponse) HasPulpCreated() bool`

HasPulpCreated returns a boolean if a field has been set.

### GetVersionsHref

`func (o *AnsibleAnsibleRepositoryResponse) GetVersionsHref() string`

GetVersionsHref returns the VersionsHref field if non-nil, zero value otherwise.

### GetVersionsHrefOk

`func (o *AnsibleAnsibleRepositoryResponse) GetVersionsHrefOk() (*string, bool)`

GetVersionsHrefOk returns a tuple with the VersionsHref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionsHref

`func (o *AnsibleAnsibleRepositoryResponse) SetVersionsHref(v string)`

SetVersionsHref sets VersionsHref field to given value.

### HasVersionsHref

`func (o *AnsibleAnsibleRepositoryResponse) HasVersionsHref() bool`

HasVersionsHref returns a boolean if a field has been set.

### GetPulpLabels

`func (o *AnsibleAnsibleRepositoryResponse) GetPulpLabels() map[string]string`

GetPulpLabels returns the PulpLabels field if non-nil, zero value otherwise.

### GetPulpLabelsOk

`func (o *AnsibleAnsibleRepositoryResponse) GetPulpLabelsOk() (*map[string]string, bool)`

GetPulpLabelsOk returns a tuple with the PulpLabels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpLabels

`func (o *AnsibleAnsibleRepositoryResponse) SetPulpLabels(v map[string]string)`

SetPulpLabels sets PulpLabels field to given value.

### HasPulpLabels

`func (o *AnsibleAnsibleRepositoryResponse) HasPulpLabels() bool`

HasPulpLabels returns a boolean if a field has been set.

### GetLatestVersionHref

`func (o *AnsibleAnsibleRepositoryResponse) GetLatestVersionHref() string`

GetLatestVersionHref returns the LatestVersionHref field if non-nil, zero value otherwise.

### GetLatestVersionHrefOk

`func (o *AnsibleAnsibleRepositoryResponse) GetLatestVersionHrefOk() (*string, bool)`

GetLatestVersionHrefOk returns a tuple with the LatestVersionHref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestVersionHref

`func (o *AnsibleAnsibleRepositoryResponse) SetLatestVersionHref(v string)`

SetLatestVersionHref sets LatestVersionHref field to given value.

### HasLatestVersionHref

`func (o *AnsibleAnsibleRepositoryResponse) HasLatestVersionHref() bool`

HasLatestVersionHref returns a boolean if a field has been set.

### GetName

`func (o *AnsibleAnsibleRepositoryResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AnsibleAnsibleRepositoryResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AnsibleAnsibleRepositoryResponse) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *AnsibleAnsibleRepositoryResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AnsibleAnsibleRepositoryResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AnsibleAnsibleRepositoryResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AnsibleAnsibleRepositoryResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *AnsibleAnsibleRepositoryResponse) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *AnsibleAnsibleRepositoryResponse) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetRetainRepoVersions

`func (o *AnsibleAnsibleRepositoryResponse) GetRetainRepoVersions() int64`

GetRetainRepoVersions returns the RetainRepoVersions field if non-nil, zero value otherwise.

### GetRetainRepoVersionsOk

`func (o *AnsibleAnsibleRepositoryResponse) GetRetainRepoVersionsOk() (*int64, bool)`

GetRetainRepoVersionsOk returns a tuple with the RetainRepoVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetainRepoVersions

`func (o *AnsibleAnsibleRepositoryResponse) SetRetainRepoVersions(v int64)`

SetRetainRepoVersions sets RetainRepoVersions field to given value.

### HasRetainRepoVersions

`func (o *AnsibleAnsibleRepositoryResponse) HasRetainRepoVersions() bool`

HasRetainRepoVersions returns a boolean if a field has been set.

### SetRetainRepoVersionsNil

`func (o *AnsibleAnsibleRepositoryResponse) SetRetainRepoVersionsNil(b bool)`

 SetRetainRepoVersionsNil sets the value for RetainRepoVersions to be an explicit nil

### UnsetRetainRepoVersions
`func (o *AnsibleAnsibleRepositoryResponse) UnsetRetainRepoVersions()`

UnsetRetainRepoVersions ensures that no value is present for RetainRepoVersions, not even an explicit nil
### GetRemote

`func (o *AnsibleAnsibleRepositoryResponse) GetRemote() string`

GetRemote returns the Remote field if non-nil, zero value otherwise.

### GetRemoteOk

`func (o *AnsibleAnsibleRepositoryResponse) GetRemoteOk() (*string, bool)`

GetRemoteOk returns a tuple with the Remote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemote

`func (o *AnsibleAnsibleRepositoryResponse) SetRemote(v string)`

SetRemote sets Remote field to given value.

### HasRemote

`func (o *AnsibleAnsibleRepositoryResponse) HasRemote() bool`

HasRemote returns a boolean if a field has been set.

### SetRemoteNil

`func (o *AnsibleAnsibleRepositoryResponse) SetRemoteNil(b bool)`

 SetRemoteNil sets the value for Remote to be an explicit nil

### UnsetRemote
`func (o *AnsibleAnsibleRepositoryResponse) UnsetRemote()`

UnsetRemote ensures that no value is present for Remote, not even an explicit nil
### GetLastSyncedMetadataTime

`func (o *AnsibleAnsibleRepositoryResponse) GetLastSyncedMetadataTime() time.Time`

GetLastSyncedMetadataTime returns the LastSyncedMetadataTime field if non-nil, zero value otherwise.

### GetLastSyncedMetadataTimeOk

`func (o *AnsibleAnsibleRepositoryResponse) GetLastSyncedMetadataTimeOk() (*time.Time, bool)`

GetLastSyncedMetadataTimeOk returns a tuple with the LastSyncedMetadataTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSyncedMetadataTime

`func (o *AnsibleAnsibleRepositoryResponse) SetLastSyncedMetadataTime(v time.Time)`

SetLastSyncedMetadataTime sets LastSyncedMetadataTime field to given value.

### HasLastSyncedMetadataTime

`func (o *AnsibleAnsibleRepositoryResponse) HasLastSyncedMetadataTime() bool`

HasLastSyncedMetadataTime returns a boolean if a field has been set.

### SetLastSyncedMetadataTimeNil

`func (o *AnsibleAnsibleRepositoryResponse) SetLastSyncedMetadataTimeNil(b bool)`

 SetLastSyncedMetadataTimeNil sets the value for LastSyncedMetadataTime to be an explicit nil

### UnsetLastSyncedMetadataTime
`func (o *AnsibleAnsibleRepositoryResponse) UnsetLastSyncedMetadataTime()`

UnsetLastSyncedMetadataTime ensures that no value is present for LastSyncedMetadataTime, not even an explicit nil
### GetGpgkey

`func (o *AnsibleAnsibleRepositoryResponse) GetGpgkey() string`

GetGpgkey returns the Gpgkey field if non-nil, zero value otherwise.

### GetGpgkeyOk

`func (o *AnsibleAnsibleRepositoryResponse) GetGpgkeyOk() (*string, bool)`

GetGpgkeyOk returns a tuple with the Gpgkey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpgkey

`func (o *AnsibleAnsibleRepositoryResponse) SetGpgkey(v string)`

SetGpgkey sets Gpgkey field to given value.

### HasGpgkey

`func (o *AnsibleAnsibleRepositoryResponse) HasGpgkey() bool`

HasGpgkey returns a boolean if a field has been set.

### SetGpgkeyNil

`func (o *AnsibleAnsibleRepositoryResponse) SetGpgkeyNil(b bool)`

 SetGpgkeyNil sets the value for Gpgkey to be an explicit nil

### UnsetGpgkey
`func (o *AnsibleAnsibleRepositoryResponse) UnsetGpgkey()`

UnsetGpgkey ensures that no value is present for Gpgkey, not even an explicit nil
### GetLastSyncTask

`func (o *AnsibleAnsibleRepositoryResponse) GetLastSyncTask() TaskResponse`

GetLastSyncTask returns the LastSyncTask field if non-nil, zero value otherwise.

### GetLastSyncTaskOk

`func (o *AnsibleAnsibleRepositoryResponse) GetLastSyncTaskOk() (*TaskResponse, bool)`

GetLastSyncTaskOk returns a tuple with the LastSyncTask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSyncTask

`func (o *AnsibleAnsibleRepositoryResponse) SetLastSyncTask(v TaskResponse)`

SetLastSyncTask sets LastSyncTask field to given value.

### HasLastSyncTask

`func (o *AnsibleAnsibleRepositoryResponse) HasLastSyncTask() bool`

HasLastSyncTask returns a boolean if a field has been set.

### GetPrivate

`func (o *AnsibleAnsibleRepositoryResponse) GetPrivate() bool`

GetPrivate returns the Private field if non-nil, zero value otherwise.

### GetPrivateOk

`func (o *AnsibleAnsibleRepositoryResponse) GetPrivateOk() (*bool, bool)`

GetPrivateOk returns a tuple with the Private field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivate

`func (o *AnsibleAnsibleRepositoryResponse) SetPrivate(v bool)`

SetPrivate sets Private field to given value.

### HasPrivate

`func (o *AnsibleAnsibleRepositoryResponse) HasPrivate() bool`

HasPrivate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


