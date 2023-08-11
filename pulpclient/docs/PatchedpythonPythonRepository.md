# PatchedpythonPythonRepository

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PulpLabels** | Pointer to **map[string]string** |  | [optional] 
**Name** | Pointer to **string** | A unique name for this repository. | [optional] 
**Description** | Pointer to **NullableString** | An optional description. | [optional] 
**RetainRepoVersions** | Pointer to **NullableInt64** | Retain X versions of the repository. Default is null which retains all versions. | [optional] 
**Remote** | Pointer to **NullableString** | An optional remote to use by default when syncing. | [optional] 
**Autopublish** | Pointer to **bool** | Whether to automatically create publications for new repository versions, and update any distributions pointing to this repository. | [optional] [default to false]

## Methods

### NewPatchedpythonPythonRepository

`func NewPatchedpythonPythonRepository() *PatchedpythonPythonRepository`

NewPatchedpythonPythonRepository instantiates a new PatchedpythonPythonRepository object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedpythonPythonRepositoryWithDefaults

`func NewPatchedpythonPythonRepositoryWithDefaults() *PatchedpythonPythonRepository`

NewPatchedpythonPythonRepositoryWithDefaults instantiates a new PatchedpythonPythonRepository object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPulpLabels

`func (o *PatchedpythonPythonRepository) GetPulpLabels() map[string]string`

GetPulpLabels returns the PulpLabels field if non-nil, zero value otherwise.

### GetPulpLabelsOk

`func (o *PatchedpythonPythonRepository) GetPulpLabelsOk() (*map[string]string, bool)`

GetPulpLabelsOk returns a tuple with the PulpLabels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpLabels

`func (o *PatchedpythonPythonRepository) SetPulpLabels(v map[string]string)`

SetPulpLabels sets PulpLabels field to given value.

### HasPulpLabels

`func (o *PatchedpythonPythonRepository) HasPulpLabels() bool`

HasPulpLabels returns a boolean if a field has been set.

### GetName

`func (o *PatchedpythonPythonRepository) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedpythonPythonRepository) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedpythonPythonRepository) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedpythonPythonRepository) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *PatchedpythonPythonRepository) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchedpythonPythonRepository) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchedpythonPythonRepository) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchedpythonPythonRepository) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *PatchedpythonPythonRepository) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *PatchedpythonPythonRepository) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetRetainRepoVersions

`func (o *PatchedpythonPythonRepository) GetRetainRepoVersions() int64`

GetRetainRepoVersions returns the RetainRepoVersions field if non-nil, zero value otherwise.

### GetRetainRepoVersionsOk

`func (o *PatchedpythonPythonRepository) GetRetainRepoVersionsOk() (*int64, bool)`

GetRetainRepoVersionsOk returns a tuple with the RetainRepoVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetainRepoVersions

`func (o *PatchedpythonPythonRepository) SetRetainRepoVersions(v int64)`

SetRetainRepoVersions sets RetainRepoVersions field to given value.

### HasRetainRepoVersions

`func (o *PatchedpythonPythonRepository) HasRetainRepoVersions() bool`

HasRetainRepoVersions returns a boolean if a field has been set.

### SetRetainRepoVersionsNil

`func (o *PatchedpythonPythonRepository) SetRetainRepoVersionsNil(b bool)`

 SetRetainRepoVersionsNil sets the value for RetainRepoVersions to be an explicit nil

### UnsetRetainRepoVersions
`func (o *PatchedpythonPythonRepository) UnsetRetainRepoVersions()`

UnsetRetainRepoVersions ensures that no value is present for RetainRepoVersions, not even an explicit nil
### GetRemote

`func (o *PatchedpythonPythonRepository) GetRemote() string`

GetRemote returns the Remote field if non-nil, zero value otherwise.

### GetRemoteOk

`func (o *PatchedpythonPythonRepository) GetRemoteOk() (*string, bool)`

GetRemoteOk returns a tuple with the Remote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemote

`func (o *PatchedpythonPythonRepository) SetRemote(v string)`

SetRemote sets Remote field to given value.

### HasRemote

`func (o *PatchedpythonPythonRepository) HasRemote() bool`

HasRemote returns a boolean if a field has been set.

### SetRemoteNil

`func (o *PatchedpythonPythonRepository) SetRemoteNil(b bool)`

 SetRemoteNil sets the value for Remote to be an explicit nil

### UnsetRemote
`func (o *PatchedpythonPythonRepository) UnsetRemote()`

UnsetRemote ensures that no value is present for Remote, not even an explicit nil
### GetAutopublish

`func (o *PatchedpythonPythonRepository) GetAutopublish() bool`

GetAutopublish returns the Autopublish field if non-nil, zero value otherwise.

### GetAutopublishOk

`func (o *PatchedpythonPythonRepository) GetAutopublishOk() (*bool, bool)`

GetAutopublishOk returns a tuple with the Autopublish field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutopublish

`func (o *PatchedpythonPythonRepository) SetAutopublish(v bool)`

SetAutopublish sets Autopublish field to given value.

### HasAutopublish

`func (o *PatchedpythonPythonRepository) HasAutopublish() bool`

HasAutopublish returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


