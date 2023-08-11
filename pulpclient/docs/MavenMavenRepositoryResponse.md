# MavenMavenRepositoryResponse

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

## Methods

### NewMavenMavenRepositoryResponse

`func NewMavenMavenRepositoryResponse(name string, ) *MavenMavenRepositoryResponse`

NewMavenMavenRepositoryResponse instantiates a new MavenMavenRepositoryResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMavenMavenRepositoryResponseWithDefaults

`func NewMavenMavenRepositoryResponseWithDefaults() *MavenMavenRepositoryResponse`

NewMavenMavenRepositoryResponseWithDefaults instantiates a new MavenMavenRepositoryResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPulpHref

`func (o *MavenMavenRepositoryResponse) GetPulpHref() string`

GetPulpHref returns the PulpHref field if non-nil, zero value otherwise.

### GetPulpHrefOk

`func (o *MavenMavenRepositoryResponse) GetPulpHrefOk() (*string, bool)`

GetPulpHrefOk returns a tuple with the PulpHref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpHref

`func (o *MavenMavenRepositoryResponse) SetPulpHref(v string)`

SetPulpHref sets PulpHref field to given value.

### HasPulpHref

`func (o *MavenMavenRepositoryResponse) HasPulpHref() bool`

HasPulpHref returns a boolean if a field has been set.

### GetPulpCreated

`func (o *MavenMavenRepositoryResponse) GetPulpCreated() time.Time`

GetPulpCreated returns the PulpCreated field if non-nil, zero value otherwise.

### GetPulpCreatedOk

`func (o *MavenMavenRepositoryResponse) GetPulpCreatedOk() (*time.Time, bool)`

GetPulpCreatedOk returns a tuple with the PulpCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpCreated

`func (o *MavenMavenRepositoryResponse) SetPulpCreated(v time.Time)`

SetPulpCreated sets PulpCreated field to given value.

### HasPulpCreated

`func (o *MavenMavenRepositoryResponse) HasPulpCreated() bool`

HasPulpCreated returns a boolean if a field has been set.

### GetVersionsHref

`func (o *MavenMavenRepositoryResponse) GetVersionsHref() string`

GetVersionsHref returns the VersionsHref field if non-nil, zero value otherwise.

### GetVersionsHrefOk

`func (o *MavenMavenRepositoryResponse) GetVersionsHrefOk() (*string, bool)`

GetVersionsHrefOk returns a tuple with the VersionsHref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionsHref

`func (o *MavenMavenRepositoryResponse) SetVersionsHref(v string)`

SetVersionsHref sets VersionsHref field to given value.

### HasVersionsHref

`func (o *MavenMavenRepositoryResponse) HasVersionsHref() bool`

HasVersionsHref returns a boolean if a field has been set.

### GetPulpLabels

`func (o *MavenMavenRepositoryResponse) GetPulpLabels() map[string]string`

GetPulpLabels returns the PulpLabels field if non-nil, zero value otherwise.

### GetPulpLabelsOk

`func (o *MavenMavenRepositoryResponse) GetPulpLabelsOk() (*map[string]string, bool)`

GetPulpLabelsOk returns a tuple with the PulpLabels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpLabels

`func (o *MavenMavenRepositoryResponse) SetPulpLabels(v map[string]string)`

SetPulpLabels sets PulpLabels field to given value.

### HasPulpLabels

`func (o *MavenMavenRepositoryResponse) HasPulpLabels() bool`

HasPulpLabels returns a boolean if a field has been set.

### GetLatestVersionHref

`func (o *MavenMavenRepositoryResponse) GetLatestVersionHref() string`

GetLatestVersionHref returns the LatestVersionHref field if non-nil, zero value otherwise.

### GetLatestVersionHrefOk

`func (o *MavenMavenRepositoryResponse) GetLatestVersionHrefOk() (*string, bool)`

GetLatestVersionHrefOk returns a tuple with the LatestVersionHref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestVersionHref

`func (o *MavenMavenRepositoryResponse) SetLatestVersionHref(v string)`

SetLatestVersionHref sets LatestVersionHref field to given value.

### HasLatestVersionHref

`func (o *MavenMavenRepositoryResponse) HasLatestVersionHref() bool`

HasLatestVersionHref returns a boolean if a field has been set.

### GetName

`func (o *MavenMavenRepositoryResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MavenMavenRepositoryResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MavenMavenRepositoryResponse) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *MavenMavenRepositoryResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MavenMavenRepositoryResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MavenMavenRepositoryResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MavenMavenRepositoryResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *MavenMavenRepositoryResponse) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *MavenMavenRepositoryResponse) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetRetainRepoVersions

`func (o *MavenMavenRepositoryResponse) GetRetainRepoVersions() int64`

GetRetainRepoVersions returns the RetainRepoVersions field if non-nil, zero value otherwise.

### GetRetainRepoVersionsOk

`func (o *MavenMavenRepositoryResponse) GetRetainRepoVersionsOk() (*int64, bool)`

GetRetainRepoVersionsOk returns a tuple with the RetainRepoVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetainRepoVersions

`func (o *MavenMavenRepositoryResponse) SetRetainRepoVersions(v int64)`

SetRetainRepoVersions sets RetainRepoVersions field to given value.

### HasRetainRepoVersions

`func (o *MavenMavenRepositoryResponse) HasRetainRepoVersions() bool`

HasRetainRepoVersions returns a boolean if a field has been set.

### SetRetainRepoVersionsNil

`func (o *MavenMavenRepositoryResponse) SetRetainRepoVersionsNil(b bool)`

 SetRetainRepoVersionsNil sets the value for RetainRepoVersions to be an explicit nil

### UnsetRetainRepoVersions
`func (o *MavenMavenRepositoryResponse) UnsetRetainRepoVersions()`

UnsetRetainRepoVersions ensures that no value is present for RetainRepoVersions, not even an explicit nil
### GetRemote

`func (o *MavenMavenRepositoryResponse) GetRemote() string`

GetRemote returns the Remote field if non-nil, zero value otherwise.

### GetRemoteOk

`func (o *MavenMavenRepositoryResponse) GetRemoteOk() (*string, bool)`

GetRemoteOk returns a tuple with the Remote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemote

`func (o *MavenMavenRepositoryResponse) SetRemote(v string)`

SetRemote sets Remote field to given value.

### HasRemote

`func (o *MavenMavenRepositoryResponse) HasRemote() bool`

HasRemote returns a boolean if a field has been set.

### SetRemoteNil

`func (o *MavenMavenRepositoryResponse) SetRemoteNil(b bool)`

 SetRemoteNil sets the value for Remote to be an explicit nil

### UnsetRemote
`func (o *MavenMavenRepositoryResponse) UnsetRemote()`

UnsetRemote ensures that no value is present for Remote, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


