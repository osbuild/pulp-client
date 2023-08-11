# DebPackageIndex

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Repository** | Pointer to **string** | A URI of a repository the new content unit should be associated with. | [optional] 
**Artifacts** | **map[string]interface{}** | A dict mapping relative paths inside the Content to the correspondingArtifact URLs. E.g.: {&#39;relative/path&#39;: &#39;/artifacts/1/&#39; | 
**Component** | Pointer to **string** | Component of the component - architecture combination. | [optional] 
**Architecture** | Pointer to **string** | Architecture of the component - architecture combination. | [optional] 
**RelativePath** | Pointer to **string** | Path of file relative to url. | [optional] 

## Methods

### NewDebPackageIndex

`func NewDebPackageIndex(artifacts map[string]interface{}, ) *DebPackageIndex`

NewDebPackageIndex instantiates a new DebPackageIndex object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDebPackageIndexWithDefaults

`func NewDebPackageIndexWithDefaults() *DebPackageIndex`

NewDebPackageIndexWithDefaults instantiates a new DebPackageIndex object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRepository

`func (o *DebPackageIndex) GetRepository() string`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *DebPackageIndex) GetRepositoryOk() (*string, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *DebPackageIndex) SetRepository(v string)`

SetRepository sets Repository field to given value.

### HasRepository

`func (o *DebPackageIndex) HasRepository() bool`

HasRepository returns a boolean if a field has been set.

### GetArtifacts

`func (o *DebPackageIndex) GetArtifacts() map[string]interface{}`

GetArtifacts returns the Artifacts field if non-nil, zero value otherwise.

### GetArtifactsOk

`func (o *DebPackageIndex) GetArtifactsOk() (*map[string]interface{}, bool)`

GetArtifactsOk returns a tuple with the Artifacts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifacts

`func (o *DebPackageIndex) SetArtifacts(v map[string]interface{})`

SetArtifacts sets Artifacts field to given value.


### GetComponent

`func (o *DebPackageIndex) GetComponent() string`

GetComponent returns the Component field if non-nil, zero value otherwise.

### GetComponentOk

`func (o *DebPackageIndex) GetComponentOk() (*string, bool)`

GetComponentOk returns a tuple with the Component field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponent

`func (o *DebPackageIndex) SetComponent(v string)`

SetComponent sets Component field to given value.

### HasComponent

`func (o *DebPackageIndex) HasComponent() bool`

HasComponent returns a boolean if a field has been set.

### GetArchitecture

`func (o *DebPackageIndex) GetArchitecture() string`

GetArchitecture returns the Architecture field if non-nil, zero value otherwise.

### GetArchitectureOk

`func (o *DebPackageIndex) GetArchitectureOk() (*string, bool)`

GetArchitectureOk returns a tuple with the Architecture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchitecture

`func (o *DebPackageIndex) SetArchitecture(v string)`

SetArchitecture sets Architecture field to given value.

### HasArchitecture

`func (o *DebPackageIndex) HasArchitecture() bool`

HasArchitecture returns a boolean if a field has been set.

### GetRelativePath

`func (o *DebPackageIndex) GetRelativePath() string`

GetRelativePath returns the RelativePath field if non-nil, zero value otherwise.

### GetRelativePathOk

`func (o *DebPackageIndex) GetRelativePathOk() (*string, bool)`

GetRelativePathOk returns a tuple with the RelativePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelativePath

`func (o *DebPackageIndex) SetRelativePath(v string)`

SetRelativePath sets RelativePath field to given value.

### HasRelativePath

`func (o *DebPackageIndex) HasRelativePath() bool`

HasRelativePath returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


