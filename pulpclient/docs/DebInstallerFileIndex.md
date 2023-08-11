# DebInstallerFileIndex

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Repository** | Pointer to **string** | A URI of a repository the new content unit should be associated with. | [optional] 
**Artifacts** | **map[string]interface{}** | A dict mapping relative paths inside the Content to the correspondingArtifact URLs. E.g.: {&#39;relative/path&#39;: &#39;/artifacts/1/&#39; | 
**Component** | **string** | Component of the component - architecture combination. | 
**Architecture** | **string** | Architecture of the component - architecture combination. | 
**RelativePath** | Pointer to **string** | Path of directory containing MD5SUMS and SHA256SUMS relative to url. | [optional] 

## Methods

### NewDebInstallerFileIndex

`func NewDebInstallerFileIndex(artifacts map[string]interface{}, component string, architecture string, ) *DebInstallerFileIndex`

NewDebInstallerFileIndex instantiates a new DebInstallerFileIndex object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDebInstallerFileIndexWithDefaults

`func NewDebInstallerFileIndexWithDefaults() *DebInstallerFileIndex`

NewDebInstallerFileIndexWithDefaults instantiates a new DebInstallerFileIndex object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRepository

`func (o *DebInstallerFileIndex) GetRepository() string`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *DebInstallerFileIndex) GetRepositoryOk() (*string, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *DebInstallerFileIndex) SetRepository(v string)`

SetRepository sets Repository field to given value.

### HasRepository

`func (o *DebInstallerFileIndex) HasRepository() bool`

HasRepository returns a boolean if a field has been set.

### GetArtifacts

`func (o *DebInstallerFileIndex) GetArtifacts() map[string]interface{}`

GetArtifacts returns the Artifacts field if non-nil, zero value otherwise.

### GetArtifactsOk

`func (o *DebInstallerFileIndex) GetArtifactsOk() (*map[string]interface{}, bool)`

GetArtifactsOk returns a tuple with the Artifacts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifacts

`func (o *DebInstallerFileIndex) SetArtifacts(v map[string]interface{})`

SetArtifacts sets Artifacts field to given value.


### GetComponent

`func (o *DebInstallerFileIndex) GetComponent() string`

GetComponent returns the Component field if non-nil, zero value otherwise.

### GetComponentOk

`func (o *DebInstallerFileIndex) GetComponentOk() (*string, bool)`

GetComponentOk returns a tuple with the Component field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponent

`func (o *DebInstallerFileIndex) SetComponent(v string)`

SetComponent sets Component field to given value.


### GetArchitecture

`func (o *DebInstallerFileIndex) GetArchitecture() string`

GetArchitecture returns the Architecture field if non-nil, zero value otherwise.

### GetArchitectureOk

`func (o *DebInstallerFileIndex) GetArchitectureOk() (*string, bool)`

GetArchitectureOk returns a tuple with the Architecture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchitecture

`func (o *DebInstallerFileIndex) SetArchitecture(v string)`

SetArchitecture sets Architecture field to given value.


### GetRelativePath

`func (o *DebInstallerFileIndex) GetRelativePath() string`

GetRelativePath returns the RelativePath field if non-nil, zero value otherwise.

### GetRelativePathOk

`func (o *DebInstallerFileIndex) GetRelativePathOk() (*string, bool)`

GetRelativePathOk returns a tuple with the RelativePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelativePath

`func (o *DebInstallerFileIndex) SetRelativePath(v string)`

SetRelativePath sets RelativePath field to given value.

### HasRelativePath

`func (o *DebInstallerFileIndex) HasRelativePath() bool`

HasRelativePath returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


