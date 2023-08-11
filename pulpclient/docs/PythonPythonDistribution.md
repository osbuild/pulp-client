# PythonPythonDistribution

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BasePath** | **string** | The base (relative) path component of the published url. Avoid paths that                     overlap with other distribution base paths (e.g. \&quot;foo\&quot; and \&quot;foo/bar\&quot;) | 
**ContentGuard** | Pointer to **NullableString** | An optional content-guard. | [optional] 
**Hidden** | Pointer to **bool** | Whether this distribution should be shown in the content app. | [optional] [default to false]
**PulpLabels** | Pointer to **map[string]string** |  | [optional] 
**Name** | **string** | A unique name. Ex, &#x60;rawhide&#x60; and &#x60;stable&#x60;. | 
**Repository** | Pointer to **NullableString** | The latest RepositoryVersion for this Repository will be served. | [optional] 
**Publication** | Pointer to **NullableString** | Publication to be served | [optional] 
**AllowUploads** | Pointer to **bool** | Allow packages to be uploaded to this index. | [optional] [default to true]
**Remote** | Pointer to **NullableString** | Remote that can be used to fetch content when using pull-through caching. | [optional] 

## Methods

### NewPythonPythonDistribution

`func NewPythonPythonDistribution(basePath string, name string, ) *PythonPythonDistribution`

NewPythonPythonDistribution instantiates a new PythonPythonDistribution object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPythonPythonDistributionWithDefaults

`func NewPythonPythonDistributionWithDefaults() *PythonPythonDistribution`

NewPythonPythonDistributionWithDefaults instantiates a new PythonPythonDistribution object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBasePath

`func (o *PythonPythonDistribution) GetBasePath() string`

GetBasePath returns the BasePath field if non-nil, zero value otherwise.

### GetBasePathOk

`func (o *PythonPythonDistribution) GetBasePathOk() (*string, bool)`

GetBasePathOk returns a tuple with the BasePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasePath

`func (o *PythonPythonDistribution) SetBasePath(v string)`

SetBasePath sets BasePath field to given value.


### GetContentGuard

`func (o *PythonPythonDistribution) GetContentGuard() string`

GetContentGuard returns the ContentGuard field if non-nil, zero value otherwise.

### GetContentGuardOk

`func (o *PythonPythonDistribution) GetContentGuardOk() (*string, bool)`

GetContentGuardOk returns a tuple with the ContentGuard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentGuard

`func (o *PythonPythonDistribution) SetContentGuard(v string)`

SetContentGuard sets ContentGuard field to given value.

### HasContentGuard

`func (o *PythonPythonDistribution) HasContentGuard() bool`

HasContentGuard returns a boolean if a field has been set.

### SetContentGuardNil

`func (o *PythonPythonDistribution) SetContentGuardNil(b bool)`

 SetContentGuardNil sets the value for ContentGuard to be an explicit nil

### UnsetContentGuard
`func (o *PythonPythonDistribution) UnsetContentGuard()`

UnsetContentGuard ensures that no value is present for ContentGuard, not even an explicit nil
### GetHidden

`func (o *PythonPythonDistribution) GetHidden() bool`

GetHidden returns the Hidden field if non-nil, zero value otherwise.

### GetHiddenOk

`func (o *PythonPythonDistribution) GetHiddenOk() (*bool, bool)`

GetHiddenOk returns a tuple with the Hidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHidden

`func (o *PythonPythonDistribution) SetHidden(v bool)`

SetHidden sets Hidden field to given value.

### HasHidden

`func (o *PythonPythonDistribution) HasHidden() bool`

HasHidden returns a boolean if a field has been set.

### GetPulpLabels

`func (o *PythonPythonDistribution) GetPulpLabels() map[string]string`

GetPulpLabels returns the PulpLabels field if non-nil, zero value otherwise.

### GetPulpLabelsOk

`func (o *PythonPythonDistribution) GetPulpLabelsOk() (*map[string]string, bool)`

GetPulpLabelsOk returns a tuple with the PulpLabels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpLabels

`func (o *PythonPythonDistribution) SetPulpLabels(v map[string]string)`

SetPulpLabels sets PulpLabels field to given value.

### HasPulpLabels

`func (o *PythonPythonDistribution) HasPulpLabels() bool`

HasPulpLabels returns a boolean if a field has been set.

### GetName

`func (o *PythonPythonDistribution) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PythonPythonDistribution) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PythonPythonDistribution) SetName(v string)`

SetName sets Name field to given value.


### GetRepository

`func (o *PythonPythonDistribution) GetRepository() string`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *PythonPythonDistribution) GetRepositoryOk() (*string, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *PythonPythonDistribution) SetRepository(v string)`

SetRepository sets Repository field to given value.

### HasRepository

`func (o *PythonPythonDistribution) HasRepository() bool`

HasRepository returns a boolean if a field has been set.

### SetRepositoryNil

`func (o *PythonPythonDistribution) SetRepositoryNil(b bool)`

 SetRepositoryNil sets the value for Repository to be an explicit nil

### UnsetRepository
`func (o *PythonPythonDistribution) UnsetRepository()`

UnsetRepository ensures that no value is present for Repository, not even an explicit nil
### GetPublication

`func (o *PythonPythonDistribution) GetPublication() string`

GetPublication returns the Publication field if non-nil, zero value otherwise.

### GetPublicationOk

`func (o *PythonPythonDistribution) GetPublicationOk() (*string, bool)`

GetPublicationOk returns a tuple with the Publication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublication

`func (o *PythonPythonDistribution) SetPublication(v string)`

SetPublication sets Publication field to given value.

### HasPublication

`func (o *PythonPythonDistribution) HasPublication() bool`

HasPublication returns a boolean if a field has been set.

### SetPublicationNil

`func (o *PythonPythonDistribution) SetPublicationNil(b bool)`

 SetPublicationNil sets the value for Publication to be an explicit nil

### UnsetPublication
`func (o *PythonPythonDistribution) UnsetPublication()`

UnsetPublication ensures that no value is present for Publication, not even an explicit nil
### GetAllowUploads

`func (o *PythonPythonDistribution) GetAllowUploads() bool`

GetAllowUploads returns the AllowUploads field if non-nil, zero value otherwise.

### GetAllowUploadsOk

`func (o *PythonPythonDistribution) GetAllowUploadsOk() (*bool, bool)`

GetAllowUploadsOk returns a tuple with the AllowUploads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowUploads

`func (o *PythonPythonDistribution) SetAllowUploads(v bool)`

SetAllowUploads sets AllowUploads field to given value.

### HasAllowUploads

`func (o *PythonPythonDistribution) HasAllowUploads() bool`

HasAllowUploads returns a boolean if a field has been set.

### GetRemote

`func (o *PythonPythonDistribution) GetRemote() string`

GetRemote returns the Remote field if non-nil, zero value otherwise.

### GetRemoteOk

`func (o *PythonPythonDistribution) GetRemoteOk() (*string, bool)`

GetRemoteOk returns a tuple with the Remote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemote

`func (o *PythonPythonDistribution) SetRemote(v string)`

SetRemote sets Remote field to given value.

### HasRemote

`func (o *PythonPythonDistribution) HasRemote() bool`

HasRemote returns a boolean if a field has been set.

### SetRemoteNil

`func (o *PythonPythonDistribution) SetRemoteNil(b bool)`

 SetRemoteNil sets the value for Remote to be an explicit nil

### UnsetRemote
`func (o *PythonPythonDistribution) UnsetRemote()`

UnsetRemote ensures that no value is present for Remote, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


