# PatchedansibleAnsibleDistribution

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BasePath** | Pointer to **string** | The base (relative) path component of the published url. Avoid paths that                     overlap with other distribution base paths (e.g. \&quot;foo\&quot; and \&quot;foo/bar\&quot;) | [optional] 
**ContentGuard** | Pointer to **NullableString** | An optional content-guard. | [optional] 
**Name** | Pointer to **string** | A unique name. Ex, &#x60;rawhide&#x60; and &#x60;stable&#x60;. | [optional] 
**Repository** | Pointer to **NullableString** | The latest RepositoryVersion for this Repository will be served. | [optional] 
**RepositoryVersion** | Pointer to **NullableString** | RepositoryVersion to be served | [optional] 
**PulpLabels** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewPatchedansibleAnsibleDistribution

`func NewPatchedansibleAnsibleDistribution() *PatchedansibleAnsibleDistribution`

NewPatchedansibleAnsibleDistribution instantiates a new PatchedansibleAnsibleDistribution object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedansibleAnsibleDistributionWithDefaults

`func NewPatchedansibleAnsibleDistributionWithDefaults() *PatchedansibleAnsibleDistribution`

NewPatchedansibleAnsibleDistributionWithDefaults instantiates a new PatchedansibleAnsibleDistribution object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBasePath

`func (o *PatchedansibleAnsibleDistribution) GetBasePath() string`

GetBasePath returns the BasePath field if non-nil, zero value otherwise.

### GetBasePathOk

`func (o *PatchedansibleAnsibleDistribution) GetBasePathOk() (*string, bool)`

GetBasePathOk returns a tuple with the BasePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasePath

`func (o *PatchedansibleAnsibleDistribution) SetBasePath(v string)`

SetBasePath sets BasePath field to given value.

### HasBasePath

`func (o *PatchedansibleAnsibleDistribution) HasBasePath() bool`

HasBasePath returns a boolean if a field has been set.

### GetContentGuard

`func (o *PatchedansibleAnsibleDistribution) GetContentGuard() string`

GetContentGuard returns the ContentGuard field if non-nil, zero value otherwise.

### GetContentGuardOk

`func (o *PatchedansibleAnsibleDistribution) GetContentGuardOk() (*string, bool)`

GetContentGuardOk returns a tuple with the ContentGuard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentGuard

`func (o *PatchedansibleAnsibleDistribution) SetContentGuard(v string)`

SetContentGuard sets ContentGuard field to given value.

### HasContentGuard

`func (o *PatchedansibleAnsibleDistribution) HasContentGuard() bool`

HasContentGuard returns a boolean if a field has been set.

### SetContentGuardNil

`func (o *PatchedansibleAnsibleDistribution) SetContentGuardNil(b bool)`

 SetContentGuardNil sets the value for ContentGuard to be an explicit nil

### UnsetContentGuard
`func (o *PatchedansibleAnsibleDistribution) UnsetContentGuard()`

UnsetContentGuard ensures that no value is present for ContentGuard, not even an explicit nil
### GetName

`func (o *PatchedansibleAnsibleDistribution) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedansibleAnsibleDistribution) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedansibleAnsibleDistribution) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedansibleAnsibleDistribution) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRepository

`func (o *PatchedansibleAnsibleDistribution) GetRepository() string`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *PatchedansibleAnsibleDistribution) GetRepositoryOk() (*string, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *PatchedansibleAnsibleDistribution) SetRepository(v string)`

SetRepository sets Repository field to given value.

### HasRepository

`func (o *PatchedansibleAnsibleDistribution) HasRepository() bool`

HasRepository returns a boolean if a field has been set.

### SetRepositoryNil

`func (o *PatchedansibleAnsibleDistribution) SetRepositoryNil(b bool)`

 SetRepositoryNil sets the value for Repository to be an explicit nil

### UnsetRepository
`func (o *PatchedansibleAnsibleDistribution) UnsetRepository()`

UnsetRepository ensures that no value is present for Repository, not even an explicit nil
### GetRepositoryVersion

`func (o *PatchedansibleAnsibleDistribution) GetRepositoryVersion() string`

GetRepositoryVersion returns the RepositoryVersion field if non-nil, zero value otherwise.

### GetRepositoryVersionOk

`func (o *PatchedansibleAnsibleDistribution) GetRepositoryVersionOk() (*string, bool)`

GetRepositoryVersionOk returns a tuple with the RepositoryVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepositoryVersion

`func (o *PatchedansibleAnsibleDistribution) SetRepositoryVersion(v string)`

SetRepositoryVersion sets RepositoryVersion field to given value.

### HasRepositoryVersion

`func (o *PatchedansibleAnsibleDistribution) HasRepositoryVersion() bool`

HasRepositoryVersion returns a boolean if a field has been set.

### SetRepositoryVersionNil

`func (o *PatchedansibleAnsibleDistribution) SetRepositoryVersionNil(b bool)`

 SetRepositoryVersionNil sets the value for RepositoryVersion to be an explicit nil

### UnsetRepositoryVersion
`func (o *PatchedansibleAnsibleDistribution) UnsetRepositoryVersion()`

UnsetRepositoryVersion ensures that no value is present for RepositoryVersion, not even an explicit nil
### GetPulpLabels

`func (o *PatchedansibleAnsibleDistribution) GetPulpLabels() map[string]string`

GetPulpLabels returns the PulpLabels field if non-nil, zero value otherwise.

### GetPulpLabelsOk

`func (o *PatchedansibleAnsibleDistribution) GetPulpLabelsOk() (*map[string]string, bool)`

GetPulpLabelsOk returns a tuple with the PulpLabels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpLabels

`func (o *PatchedansibleAnsibleDistribution) SetPulpLabels(v map[string]string)`

SetPulpLabels sets PulpLabels field to given value.

### HasPulpLabels

`func (o *PatchedansibleAnsibleDistribution) HasPulpLabels() bool`

HasPulpLabels returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


