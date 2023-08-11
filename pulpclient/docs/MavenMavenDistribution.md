# MavenMavenDistribution

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BasePath** | **string** | The base (relative) path component of the published url. Avoid paths that                     overlap with other distribution base paths (e.g. \&quot;foo\&quot; and \&quot;foo/bar\&quot;) | 
**ContentGuard** | Pointer to **NullableString** | An optional content-guard. | [optional] 
**Hidden** | Pointer to **bool** | Whether this distribution should be shown in the content app. | [optional] [default to false]
**PulpLabels** | Pointer to **map[string]string** |  | [optional] 
**Name** | **string** | A unique name. Ex, &#x60;rawhide&#x60; and &#x60;stable&#x60;. | 
**Repository** | Pointer to **NullableString** | The latest RepositoryVersion for this Repository will be served. | [optional] 
**Remote** | Pointer to **NullableString** | Remote that can be used to fetch content when using pull-through caching. | [optional] 

## Methods

### NewMavenMavenDistribution

`func NewMavenMavenDistribution(basePath string, name string, ) *MavenMavenDistribution`

NewMavenMavenDistribution instantiates a new MavenMavenDistribution object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMavenMavenDistributionWithDefaults

`func NewMavenMavenDistributionWithDefaults() *MavenMavenDistribution`

NewMavenMavenDistributionWithDefaults instantiates a new MavenMavenDistribution object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBasePath

`func (o *MavenMavenDistribution) GetBasePath() string`

GetBasePath returns the BasePath field if non-nil, zero value otherwise.

### GetBasePathOk

`func (o *MavenMavenDistribution) GetBasePathOk() (*string, bool)`

GetBasePathOk returns a tuple with the BasePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasePath

`func (o *MavenMavenDistribution) SetBasePath(v string)`

SetBasePath sets BasePath field to given value.


### GetContentGuard

`func (o *MavenMavenDistribution) GetContentGuard() string`

GetContentGuard returns the ContentGuard field if non-nil, zero value otherwise.

### GetContentGuardOk

`func (o *MavenMavenDistribution) GetContentGuardOk() (*string, bool)`

GetContentGuardOk returns a tuple with the ContentGuard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentGuard

`func (o *MavenMavenDistribution) SetContentGuard(v string)`

SetContentGuard sets ContentGuard field to given value.

### HasContentGuard

`func (o *MavenMavenDistribution) HasContentGuard() bool`

HasContentGuard returns a boolean if a field has been set.

### SetContentGuardNil

`func (o *MavenMavenDistribution) SetContentGuardNil(b bool)`

 SetContentGuardNil sets the value for ContentGuard to be an explicit nil

### UnsetContentGuard
`func (o *MavenMavenDistribution) UnsetContentGuard()`

UnsetContentGuard ensures that no value is present for ContentGuard, not even an explicit nil
### GetHidden

`func (o *MavenMavenDistribution) GetHidden() bool`

GetHidden returns the Hidden field if non-nil, zero value otherwise.

### GetHiddenOk

`func (o *MavenMavenDistribution) GetHiddenOk() (*bool, bool)`

GetHiddenOk returns a tuple with the Hidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHidden

`func (o *MavenMavenDistribution) SetHidden(v bool)`

SetHidden sets Hidden field to given value.

### HasHidden

`func (o *MavenMavenDistribution) HasHidden() bool`

HasHidden returns a boolean if a field has been set.

### GetPulpLabels

`func (o *MavenMavenDistribution) GetPulpLabels() map[string]string`

GetPulpLabels returns the PulpLabels field if non-nil, zero value otherwise.

### GetPulpLabelsOk

`func (o *MavenMavenDistribution) GetPulpLabelsOk() (*map[string]string, bool)`

GetPulpLabelsOk returns a tuple with the PulpLabels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpLabels

`func (o *MavenMavenDistribution) SetPulpLabels(v map[string]string)`

SetPulpLabels sets PulpLabels field to given value.

### HasPulpLabels

`func (o *MavenMavenDistribution) HasPulpLabels() bool`

HasPulpLabels returns a boolean if a field has been set.

### GetName

`func (o *MavenMavenDistribution) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MavenMavenDistribution) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MavenMavenDistribution) SetName(v string)`

SetName sets Name field to given value.


### GetRepository

`func (o *MavenMavenDistribution) GetRepository() string`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *MavenMavenDistribution) GetRepositoryOk() (*string, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *MavenMavenDistribution) SetRepository(v string)`

SetRepository sets Repository field to given value.

### HasRepository

`func (o *MavenMavenDistribution) HasRepository() bool`

HasRepository returns a boolean if a field has been set.

### SetRepositoryNil

`func (o *MavenMavenDistribution) SetRepositoryNil(b bool)`

 SetRepositoryNil sets the value for Repository to be an explicit nil

### UnsetRepository
`func (o *MavenMavenDistribution) UnsetRepository()`

UnsetRepository ensures that no value is present for Repository, not even an explicit nil
### GetRemote

`func (o *MavenMavenDistribution) GetRemote() string`

GetRemote returns the Remote field if non-nil, zero value otherwise.

### GetRemoteOk

`func (o *MavenMavenDistribution) GetRemoteOk() (*string, bool)`

GetRemoteOk returns a tuple with the Remote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemote

`func (o *MavenMavenDistribution) SetRemote(v string)`

SetRemote sets Remote field to given value.

### HasRemote

`func (o *MavenMavenDistribution) HasRemote() bool`

HasRemote returns a boolean if a field has been set.

### SetRemoteNil

`func (o *MavenMavenDistribution) SetRemoteNil(b bool)`

 SetRemoteNil sets the value for Remote to be an explicit nil

### UnsetRemote
`func (o *MavenMavenDistribution) UnsetRemote()`

UnsetRemote ensures that no value is present for Remote, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


