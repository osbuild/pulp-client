# RpmModulemd

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Repository** | Pointer to **string** | A URI of a repository the new content unit should be associated with. | [optional] 
**Name** | **string** | Modulemd name. | 
**Stream** | **string** | Stream name. | 
**Version** | **string** | Modulemd version. | 
**StaticContext** | Pointer to **bool** | Modulemd static-context flag. | [optional] 
**Context** | **string** | Modulemd context. | 
**Arch** | **string** | Modulemd architecture. | 
**Artifacts** | **map[string]interface{}** | Modulemd artifacts. | 
**Dependencies** | **map[string]interface{}** | Modulemd dependencies. | 
**Packages** | Pointer to **[]string** | Modulemd artifacts&#39; packages. | [optional] 
**Snippet** | **string** | Modulemd snippet | 
**Profiles** | **map[string]interface{}** | Modulemd profiles. | 
**Description** | **string** | Description of module. | 

## Methods

### NewRpmModulemd

`func NewRpmModulemd(name string, stream string, version string, context string, arch string, artifacts map[string]interface{}, dependencies map[string]interface{}, snippet string, profiles map[string]interface{}, description string, ) *RpmModulemd`

NewRpmModulemd instantiates a new RpmModulemd object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRpmModulemdWithDefaults

`func NewRpmModulemdWithDefaults() *RpmModulemd`

NewRpmModulemdWithDefaults instantiates a new RpmModulemd object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRepository

`func (o *RpmModulemd) GetRepository() string`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *RpmModulemd) GetRepositoryOk() (*string, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *RpmModulemd) SetRepository(v string)`

SetRepository sets Repository field to given value.

### HasRepository

`func (o *RpmModulemd) HasRepository() bool`

HasRepository returns a boolean if a field has been set.

### GetName

`func (o *RpmModulemd) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RpmModulemd) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RpmModulemd) SetName(v string)`

SetName sets Name field to given value.


### GetStream

`func (o *RpmModulemd) GetStream() string`

GetStream returns the Stream field if non-nil, zero value otherwise.

### GetStreamOk

`func (o *RpmModulemd) GetStreamOk() (*string, bool)`

GetStreamOk returns a tuple with the Stream field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStream

`func (o *RpmModulemd) SetStream(v string)`

SetStream sets Stream field to given value.


### GetVersion

`func (o *RpmModulemd) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *RpmModulemd) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *RpmModulemd) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetStaticContext

`func (o *RpmModulemd) GetStaticContext() bool`

GetStaticContext returns the StaticContext field if non-nil, zero value otherwise.

### GetStaticContextOk

`func (o *RpmModulemd) GetStaticContextOk() (*bool, bool)`

GetStaticContextOk returns a tuple with the StaticContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaticContext

`func (o *RpmModulemd) SetStaticContext(v bool)`

SetStaticContext sets StaticContext field to given value.

### HasStaticContext

`func (o *RpmModulemd) HasStaticContext() bool`

HasStaticContext returns a boolean if a field has been set.

### GetContext

`func (o *RpmModulemd) GetContext() string`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *RpmModulemd) GetContextOk() (*string, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *RpmModulemd) SetContext(v string)`

SetContext sets Context field to given value.


### GetArch

`func (o *RpmModulemd) GetArch() string`

GetArch returns the Arch field if non-nil, zero value otherwise.

### GetArchOk

`func (o *RpmModulemd) GetArchOk() (*string, bool)`

GetArchOk returns a tuple with the Arch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArch

`func (o *RpmModulemd) SetArch(v string)`

SetArch sets Arch field to given value.


### GetArtifacts

`func (o *RpmModulemd) GetArtifacts() map[string]interface{}`

GetArtifacts returns the Artifacts field if non-nil, zero value otherwise.

### GetArtifactsOk

`func (o *RpmModulemd) GetArtifactsOk() (*map[string]interface{}, bool)`

GetArtifactsOk returns a tuple with the Artifacts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifacts

`func (o *RpmModulemd) SetArtifacts(v map[string]interface{})`

SetArtifacts sets Artifacts field to given value.


### SetArtifactsNil

`func (o *RpmModulemd) SetArtifactsNil(b bool)`

 SetArtifactsNil sets the value for Artifacts to be an explicit nil

### UnsetArtifacts
`func (o *RpmModulemd) UnsetArtifacts()`

UnsetArtifacts ensures that no value is present for Artifacts, not even an explicit nil
### GetDependencies

`func (o *RpmModulemd) GetDependencies() map[string]interface{}`

GetDependencies returns the Dependencies field if non-nil, zero value otherwise.

### GetDependenciesOk

`func (o *RpmModulemd) GetDependenciesOk() (*map[string]interface{}, bool)`

GetDependenciesOk returns a tuple with the Dependencies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependencies

`func (o *RpmModulemd) SetDependencies(v map[string]interface{})`

SetDependencies sets Dependencies field to given value.


### SetDependenciesNil

`func (o *RpmModulemd) SetDependenciesNil(b bool)`

 SetDependenciesNil sets the value for Dependencies to be an explicit nil

### UnsetDependencies
`func (o *RpmModulemd) UnsetDependencies()`

UnsetDependencies ensures that no value is present for Dependencies, not even an explicit nil
### GetPackages

`func (o *RpmModulemd) GetPackages() []*string`

GetPackages returns the Packages field if non-nil, zero value otherwise.

### GetPackagesOk

`func (o *RpmModulemd) GetPackagesOk() (*[]*string, bool)`

GetPackagesOk returns a tuple with the Packages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackages

`func (o *RpmModulemd) SetPackages(v []*string)`

SetPackages sets Packages field to given value.

### HasPackages

`func (o *RpmModulemd) HasPackages() bool`

HasPackages returns a boolean if a field has been set.

### GetSnippet

`func (o *RpmModulemd) GetSnippet() string`

GetSnippet returns the Snippet field if non-nil, zero value otherwise.

### GetSnippetOk

`func (o *RpmModulemd) GetSnippetOk() (*string, bool)`

GetSnippetOk returns a tuple with the Snippet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnippet

`func (o *RpmModulemd) SetSnippet(v string)`

SetSnippet sets Snippet field to given value.


### GetProfiles

`func (o *RpmModulemd) GetProfiles() map[string]interface{}`

GetProfiles returns the Profiles field if non-nil, zero value otherwise.

### GetProfilesOk

`func (o *RpmModulemd) GetProfilesOk() (*map[string]interface{}, bool)`

GetProfilesOk returns a tuple with the Profiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfiles

`func (o *RpmModulemd) SetProfiles(v map[string]interface{})`

SetProfiles sets Profiles field to given value.


### SetProfilesNil

`func (o *RpmModulemd) SetProfilesNil(b bool)`

 SetProfilesNil sets the value for Profiles to be an explicit nil

### UnsetProfiles
`func (o *RpmModulemd) UnsetProfiles()`

UnsetProfiles ensures that no value is present for Profiles, not even an explicit nil
### GetDescription

`func (o *RpmModulemd) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RpmModulemd) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RpmModulemd) SetDescription(v string)`

SetDescription sets Description field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


