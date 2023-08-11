# MavenMavenArtifact

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Repository** | Pointer to **string** | A URI of a repository the new content unit should be associated with. | [optional] 
**Artifact** | **string** | Artifact file representing the physical content | 
**RelativePath** | **string** | Path where the artifact is located relative to distributions base_path | 

## Methods

### NewMavenMavenArtifact

`func NewMavenMavenArtifact(artifact string, relativePath string, ) *MavenMavenArtifact`

NewMavenMavenArtifact instantiates a new MavenMavenArtifact object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMavenMavenArtifactWithDefaults

`func NewMavenMavenArtifactWithDefaults() *MavenMavenArtifact`

NewMavenMavenArtifactWithDefaults instantiates a new MavenMavenArtifact object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRepository

`func (o *MavenMavenArtifact) GetRepository() string`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *MavenMavenArtifact) GetRepositoryOk() (*string, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *MavenMavenArtifact) SetRepository(v string)`

SetRepository sets Repository field to given value.

### HasRepository

`func (o *MavenMavenArtifact) HasRepository() bool`

HasRepository returns a boolean if a field has been set.

### GetArtifact

`func (o *MavenMavenArtifact) GetArtifact() string`

GetArtifact returns the Artifact field if non-nil, zero value otherwise.

### GetArtifactOk

`func (o *MavenMavenArtifact) GetArtifactOk() (*string, bool)`

GetArtifactOk returns a tuple with the Artifact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifact

`func (o *MavenMavenArtifact) SetArtifact(v string)`

SetArtifact sets Artifact field to given value.


### GetRelativePath

`func (o *MavenMavenArtifact) GetRelativePath() string`

GetRelativePath returns the RelativePath field if non-nil, zero value otherwise.

### GetRelativePathOk

`func (o *MavenMavenArtifact) GetRelativePathOk() (*string, bool)`

GetRelativePathOk returns a tuple with the RelativePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelativePath

`func (o *MavenMavenArtifact) SetRelativePath(v string)`

SetRelativePath sets RelativePath field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


