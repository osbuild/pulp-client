# PatchedrpmRpmDistribution

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BasePath** | Pointer to **string** | The base (relative) path component of the published url. Avoid paths that                     overlap with other distribution base paths (e.g. \&quot;foo\&quot; and \&quot;foo/bar\&quot;) | [optional] 
**ContentGuard** | Pointer to **NullableString** | An optional content-guard. | [optional] 
**Hidden** | Pointer to **bool** | Whether this distribution should be shown in the content app. | [optional] [default to false]
**PulpLabels** | Pointer to **map[string]string** |  | [optional] 
**Name** | Pointer to **string** | A unique name. Ex, &#x60;rawhide&#x60; and &#x60;stable&#x60;. | [optional] 
**Repository** | Pointer to **NullableString** | The latest RepositoryVersion for this Repository will be served. | [optional] 
**Publication** | Pointer to **NullableString** | Publication to be served | [optional] 

## Methods

### NewPatchedrpmRpmDistribution

`func NewPatchedrpmRpmDistribution() *PatchedrpmRpmDistribution`

NewPatchedrpmRpmDistribution instantiates a new PatchedrpmRpmDistribution object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedrpmRpmDistributionWithDefaults

`func NewPatchedrpmRpmDistributionWithDefaults() *PatchedrpmRpmDistribution`

NewPatchedrpmRpmDistributionWithDefaults instantiates a new PatchedrpmRpmDistribution object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBasePath

`func (o *PatchedrpmRpmDistribution) GetBasePath() string`

GetBasePath returns the BasePath field if non-nil, zero value otherwise.

### GetBasePathOk

`func (o *PatchedrpmRpmDistribution) GetBasePathOk() (*string, bool)`

GetBasePathOk returns a tuple with the BasePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasePath

`func (o *PatchedrpmRpmDistribution) SetBasePath(v string)`

SetBasePath sets BasePath field to given value.

### HasBasePath

`func (o *PatchedrpmRpmDistribution) HasBasePath() bool`

HasBasePath returns a boolean if a field has been set.

### GetContentGuard

`func (o *PatchedrpmRpmDistribution) GetContentGuard() string`

GetContentGuard returns the ContentGuard field if non-nil, zero value otherwise.

### GetContentGuardOk

`func (o *PatchedrpmRpmDistribution) GetContentGuardOk() (*string, bool)`

GetContentGuardOk returns a tuple with the ContentGuard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentGuard

`func (o *PatchedrpmRpmDistribution) SetContentGuard(v string)`

SetContentGuard sets ContentGuard field to given value.

### HasContentGuard

`func (o *PatchedrpmRpmDistribution) HasContentGuard() bool`

HasContentGuard returns a boolean if a field has been set.

### SetContentGuardNil

`func (o *PatchedrpmRpmDistribution) SetContentGuardNil(b bool)`

 SetContentGuardNil sets the value for ContentGuard to be an explicit nil

### UnsetContentGuard
`func (o *PatchedrpmRpmDistribution) UnsetContentGuard()`

UnsetContentGuard ensures that no value is present for ContentGuard, not even an explicit nil
### GetHidden

`func (o *PatchedrpmRpmDistribution) GetHidden() bool`

GetHidden returns the Hidden field if non-nil, zero value otherwise.

### GetHiddenOk

`func (o *PatchedrpmRpmDistribution) GetHiddenOk() (*bool, bool)`

GetHiddenOk returns a tuple with the Hidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHidden

`func (o *PatchedrpmRpmDistribution) SetHidden(v bool)`

SetHidden sets Hidden field to given value.

### HasHidden

`func (o *PatchedrpmRpmDistribution) HasHidden() bool`

HasHidden returns a boolean if a field has been set.

### GetPulpLabels

`func (o *PatchedrpmRpmDistribution) GetPulpLabels() map[string]string`

GetPulpLabels returns the PulpLabels field if non-nil, zero value otherwise.

### GetPulpLabelsOk

`func (o *PatchedrpmRpmDistribution) GetPulpLabelsOk() (*map[string]string, bool)`

GetPulpLabelsOk returns a tuple with the PulpLabels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpLabels

`func (o *PatchedrpmRpmDistribution) SetPulpLabels(v map[string]string)`

SetPulpLabels sets PulpLabels field to given value.

### HasPulpLabels

`func (o *PatchedrpmRpmDistribution) HasPulpLabels() bool`

HasPulpLabels returns a boolean if a field has been set.

### GetName

`func (o *PatchedrpmRpmDistribution) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedrpmRpmDistribution) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedrpmRpmDistribution) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedrpmRpmDistribution) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRepository

`func (o *PatchedrpmRpmDistribution) GetRepository() string`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *PatchedrpmRpmDistribution) GetRepositoryOk() (*string, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *PatchedrpmRpmDistribution) SetRepository(v string)`

SetRepository sets Repository field to given value.

### HasRepository

`func (o *PatchedrpmRpmDistribution) HasRepository() bool`

HasRepository returns a boolean if a field has been set.

### SetRepositoryNil

`func (o *PatchedrpmRpmDistribution) SetRepositoryNil(b bool)`

 SetRepositoryNil sets the value for Repository to be an explicit nil

### UnsetRepository
`func (o *PatchedrpmRpmDistribution) UnsetRepository()`

UnsetRepository ensures that no value is present for Repository, not even an explicit nil
### GetPublication

`func (o *PatchedrpmRpmDistribution) GetPublication() string`

GetPublication returns the Publication field if non-nil, zero value otherwise.

### GetPublicationOk

`func (o *PatchedrpmRpmDistribution) GetPublicationOk() (*string, bool)`

GetPublicationOk returns a tuple with the Publication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublication

`func (o *PatchedrpmRpmDistribution) SetPublication(v string)`

SetPublication sets Publication field to given value.

### HasPublication

`func (o *PatchedrpmRpmDistribution) HasPublication() bool`

HasPublication returns a boolean if a field has been set.

### SetPublicationNil

`func (o *PatchedrpmRpmDistribution) SetPublicationNil(b bool)`

 SetPublicationNil sets the value for Publication to be an explicit nil

### UnsetPublication
`func (o *PatchedrpmRpmDistribution) UnsetPublication()`

UnsetPublication ensures that no value is present for Publication, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


