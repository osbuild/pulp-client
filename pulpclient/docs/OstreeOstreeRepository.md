# OstreeOstreeRepository

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PulpLabels** | Pointer to **map[string]string** |  | [optional] 
**Name** | **string** | A unique name for this repository. | 
**Description** | Pointer to **NullableString** | An optional description. | [optional] 
**RetainRepoVersions** | Pointer to **NullableInt64** | Retain X versions of the repository. Default is null which retains all versions. | [optional] 
**Remote** | Pointer to **NullableString** | An optional remote to use by default when syncing. | [optional] 
**ComputeDelta** | Pointer to **bool** |  | [optional] [default to true]

## Methods

### NewOstreeOstreeRepository

`func NewOstreeOstreeRepository(name string, ) *OstreeOstreeRepository`

NewOstreeOstreeRepository instantiates a new OstreeOstreeRepository object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOstreeOstreeRepositoryWithDefaults

`func NewOstreeOstreeRepositoryWithDefaults() *OstreeOstreeRepository`

NewOstreeOstreeRepositoryWithDefaults instantiates a new OstreeOstreeRepository object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPulpLabels

`func (o *OstreeOstreeRepository) GetPulpLabels() map[string]string`

GetPulpLabels returns the PulpLabels field if non-nil, zero value otherwise.

### GetPulpLabelsOk

`func (o *OstreeOstreeRepository) GetPulpLabelsOk() (*map[string]string, bool)`

GetPulpLabelsOk returns a tuple with the PulpLabels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpLabels

`func (o *OstreeOstreeRepository) SetPulpLabels(v map[string]string)`

SetPulpLabels sets PulpLabels field to given value.

### HasPulpLabels

`func (o *OstreeOstreeRepository) HasPulpLabels() bool`

HasPulpLabels returns a boolean if a field has been set.

### GetName

`func (o *OstreeOstreeRepository) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OstreeOstreeRepository) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OstreeOstreeRepository) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *OstreeOstreeRepository) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *OstreeOstreeRepository) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *OstreeOstreeRepository) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *OstreeOstreeRepository) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *OstreeOstreeRepository) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *OstreeOstreeRepository) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetRetainRepoVersions

`func (o *OstreeOstreeRepository) GetRetainRepoVersions() int64`

GetRetainRepoVersions returns the RetainRepoVersions field if non-nil, zero value otherwise.

### GetRetainRepoVersionsOk

`func (o *OstreeOstreeRepository) GetRetainRepoVersionsOk() (*int64, bool)`

GetRetainRepoVersionsOk returns a tuple with the RetainRepoVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetainRepoVersions

`func (o *OstreeOstreeRepository) SetRetainRepoVersions(v int64)`

SetRetainRepoVersions sets RetainRepoVersions field to given value.

### HasRetainRepoVersions

`func (o *OstreeOstreeRepository) HasRetainRepoVersions() bool`

HasRetainRepoVersions returns a boolean if a field has been set.

### SetRetainRepoVersionsNil

`func (o *OstreeOstreeRepository) SetRetainRepoVersionsNil(b bool)`

 SetRetainRepoVersionsNil sets the value for RetainRepoVersions to be an explicit nil

### UnsetRetainRepoVersions
`func (o *OstreeOstreeRepository) UnsetRetainRepoVersions()`

UnsetRetainRepoVersions ensures that no value is present for RetainRepoVersions, not even an explicit nil
### GetRemote

`func (o *OstreeOstreeRepository) GetRemote() string`

GetRemote returns the Remote field if non-nil, zero value otherwise.

### GetRemoteOk

`func (o *OstreeOstreeRepository) GetRemoteOk() (*string, bool)`

GetRemoteOk returns a tuple with the Remote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemote

`func (o *OstreeOstreeRepository) SetRemote(v string)`

SetRemote sets Remote field to given value.

### HasRemote

`func (o *OstreeOstreeRepository) HasRemote() bool`

HasRemote returns a boolean if a field has been set.

### SetRemoteNil

`func (o *OstreeOstreeRepository) SetRemoteNil(b bool)`

 SetRemoteNil sets the value for Remote to be an explicit nil

### UnsetRemote
`func (o *OstreeOstreeRepository) UnsetRemote()`

UnsetRemote ensures that no value is present for Remote, not even an explicit nil
### GetComputeDelta

`func (o *OstreeOstreeRepository) GetComputeDelta() bool`

GetComputeDelta returns the ComputeDelta field if non-nil, zero value otherwise.

### GetComputeDeltaOk

`func (o *OstreeOstreeRepository) GetComputeDeltaOk() (*bool, bool)`

GetComputeDeltaOk returns a tuple with the ComputeDelta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeDelta

`func (o *OstreeOstreeRepository) SetComputeDelta(v bool)`

SetComputeDelta sets ComputeDelta field to given value.

### HasComputeDelta

`func (o *OstreeOstreeRepository) HasComputeDelta() bool`

HasComputeDelta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


