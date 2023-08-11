# PatchedpythonPythonDistribution

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
**AllowUploads** | Pointer to **bool** | Allow packages to be uploaded to this index. | [optional] [default to true]
**Remote** | Pointer to **NullableString** | Remote that can be used to fetch content when using pull-through caching. | [optional] 

## Methods

### NewPatchedpythonPythonDistribution

`func NewPatchedpythonPythonDistribution() *PatchedpythonPythonDistribution`

NewPatchedpythonPythonDistribution instantiates a new PatchedpythonPythonDistribution object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedpythonPythonDistributionWithDefaults

`func NewPatchedpythonPythonDistributionWithDefaults() *PatchedpythonPythonDistribution`

NewPatchedpythonPythonDistributionWithDefaults instantiates a new PatchedpythonPythonDistribution object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBasePath

`func (o *PatchedpythonPythonDistribution) GetBasePath() string`

GetBasePath returns the BasePath field if non-nil, zero value otherwise.

### GetBasePathOk

`func (o *PatchedpythonPythonDistribution) GetBasePathOk() (*string, bool)`

GetBasePathOk returns a tuple with the BasePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasePath

`func (o *PatchedpythonPythonDistribution) SetBasePath(v string)`

SetBasePath sets BasePath field to given value.

### HasBasePath

`func (o *PatchedpythonPythonDistribution) HasBasePath() bool`

HasBasePath returns a boolean if a field has been set.

### GetContentGuard

`func (o *PatchedpythonPythonDistribution) GetContentGuard() string`

GetContentGuard returns the ContentGuard field if non-nil, zero value otherwise.

### GetContentGuardOk

`func (o *PatchedpythonPythonDistribution) GetContentGuardOk() (*string, bool)`

GetContentGuardOk returns a tuple with the ContentGuard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentGuard

`func (o *PatchedpythonPythonDistribution) SetContentGuard(v string)`

SetContentGuard sets ContentGuard field to given value.

### HasContentGuard

`func (o *PatchedpythonPythonDistribution) HasContentGuard() bool`

HasContentGuard returns a boolean if a field has been set.

### SetContentGuardNil

`func (o *PatchedpythonPythonDistribution) SetContentGuardNil(b bool)`

 SetContentGuardNil sets the value for ContentGuard to be an explicit nil

### UnsetContentGuard
`func (o *PatchedpythonPythonDistribution) UnsetContentGuard()`

UnsetContentGuard ensures that no value is present for ContentGuard, not even an explicit nil
### GetHidden

`func (o *PatchedpythonPythonDistribution) GetHidden() bool`

GetHidden returns the Hidden field if non-nil, zero value otherwise.

### GetHiddenOk

`func (o *PatchedpythonPythonDistribution) GetHiddenOk() (*bool, bool)`

GetHiddenOk returns a tuple with the Hidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHidden

`func (o *PatchedpythonPythonDistribution) SetHidden(v bool)`

SetHidden sets Hidden field to given value.

### HasHidden

`func (o *PatchedpythonPythonDistribution) HasHidden() bool`

HasHidden returns a boolean if a field has been set.

### GetPulpLabels

`func (o *PatchedpythonPythonDistribution) GetPulpLabels() map[string]string`

GetPulpLabels returns the PulpLabels field if non-nil, zero value otherwise.

### GetPulpLabelsOk

`func (o *PatchedpythonPythonDistribution) GetPulpLabelsOk() (*map[string]string, bool)`

GetPulpLabelsOk returns a tuple with the PulpLabels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpLabels

`func (o *PatchedpythonPythonDistribution) SetPulpLabels(v map[string]string)`

SetPulpLabels sets PulpLabels field to given value.

### HasPulpLabels

`func (o *PatchedpythonPythonDistribution) HasPulpLabels() bool`

HasPulpLabels returns a boolean if a field has been set.

### GetName

`func (o *PatchedpythonPythonDistribution) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedpythonPythonDistribution) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedpythonPythonDistribution) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedpythonPythonDistribution) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRepository

`func (o *PatchedpythonPythonDistribution) GetRepository() string`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *PatchedpythonPythonDistribution) GetRepositoryOk() (*string, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *PatchedpythonPythonDistribution) SetRepository(v string)`

SetRepository sets Repository field to given value.

### HasRepository

`func (o *PatchedpythonPythonDistribution) HasRepository() bool`

HasRepository returns a boolean if a field has been set.

### SetRepositoryNil

`func (o *PatchedpythonPythonDistribution) SetRepositoryNil(b bool)`

 SetRepositoryNil sets the value for Repository to be an explicit nil

### UnsetRepository
`func (o *PatchedpythonPythonDistribution) UnsetRepository()`

UnsetRepository ensures that no value is present for Repository, not even an explicit nil
### GetPublication

`func (o *PatchedpythonPythonDistribution) GetPublication() string`

GetPublication returns the Publication field if non-nil, zero value otherwise.

### GetPublicationOk

`func (o *PatchedpythonPythonDistribution) GetPublicationOk() (*string, bool)`

GetPublicationOk returns a tuple with the Publication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublication

`func (o *PatchedpythonPythonDistribution) SetPublication(v string)`

SetPublication sets Publication field to given value.

### HasPublication

`func (o *PatchedpythonPythonDistribution) HasPublication() bool`

HasPublication returns a boolean if a field has been set.

### SetPublicationNil

`func (o *PatchedpythonPythonDistribution) SetPublicationNil(b bool)`

 SetPublicationNil sets the value for Publication to be an explicit nil

### UnsetPublication
`func (o *PatchedpythonPythonDistribution) UnsetPublication()`

UnsetPublication ensures that no value is present for Publication, not even an explicit nil
### GetAllowUploads

`func (o *PatchedpythonPythonDistribution) GetAllowUploads() bool`

GetAllowUploads returns the AllowUploads field if non-nil, zero value otherwise.

### GetAllowUploadsOk

`func (o *PatchedpythonPythonDistribution) GetAllowUploadsOk() (*bool, bool)`

GetAllowUploadsOk returns a tuple with the AllowUploads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowUploads

`func (o *PatchedpythonPythonDistribution) SetAllowUploads(v bool)`

SetAllowUploads sets AllowUploads field to given value.

### HasAllowUploads

`func (o *PatchedpythonPythonDistribution) HasAllowUploads() bool`

HasAllowUploads returns a boolean if a field has been set.

### GetRemote

`func (o *PatchedpythonPythonDistribution) GetRemote() string`

GetRemote returns the Remote field if non-nil, zero value otherwise.

### GetRemoteOk

`func (o *PatchedpythonPythonDistribution) GetRemoteOk() (*string, bool)`

GetRemoteOk returns a tuple with the Remote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemote

`func (o *PatchedpythonPythonDistribution) SetRemote(v string)`

SetRemote sets Remote field to given value.

### HasRemote

`func (o *PatchedpythonPythonDistribution) HasRemote() bool`

HasRemote returns a boolean if a field has been set.

### SetRemoteNil

`func (o *PatchedpythonPythonDistribution) SetRemoteNil(b bool)`

 SetRemoteNil sets the value for Remote to be an explicit nil

### UnsetRemote
`func (o *PatchedpythonPythonDistribution) UnsetRemote()`

UnsetRemote ensures that no value is present for Remote, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


