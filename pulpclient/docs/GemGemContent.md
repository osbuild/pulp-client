# GemGemContent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Repository** | Pointer to **string** | A URI of a repository the new content unit should be associated with. | [optional] 
**Artifact** | Pointer to **string** | Artifact file representing the physical content | [optional] 
**File** | Pointer to ***os.File** | An uploaded file that should be turned into the artifact of the content unit. | [optional] 

## Methods

### NewGemGemContent

`func NewGemGemContent() *GemGemContent`

NewGemGemContent instantiates a new GemGemContent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGemGemContentWithDefaults

`func NewGemGemContentWithDefaults() *GemGemContent`

NewGemGemContentWithDefaults instantiates a new GemGemContent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRepository

`func (o *GemGemContent) GetRepository() string`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *GemGemContent) GetRepositoryOk() (*string, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *GemGemContent) SetRepository(v string)`

SetRepository sets Repository field to given value.

### HasRepository

`func (o *GemGemContent) HasRepository() bool`

HasRepository returns a boolean if a field has been set.

### GetArtifact

`func (o *GemGemContent) GetArtifact() string`

GetArtifact returns the Artifact field if non-nil, zero value otherwise.

### GetArtifactOk

`func (o *GemGemContent) GetArtifactOk() (*string, bool)`

GetArtifactOk returns a tuple with the Artifact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifact

`func (o *GemGemContent) SetArtifact(v string)`

SetArtifact sets Artifact field to given value.

### HasArtifact

`func (o *GemGemContent) HasArtifact() bool`

HasArtifact returns a boolean if a field has been set.

### GetFile

`func (o *GemGemContent) GetFile() *os.File`

GetFile returns the File field if non-nil, zero value otherwise.

### GetFileOk

`func (o *GemGemContent) GetFileOk() (**os.File, bool)`

GetFileOk returns a tuple with the File field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFile

`func (o *GemGemContent) SetFile(v *os.File)`

SetFile sets File field to given value.

### HasFile

`func (o *GemGemContent) HasFile() bool`

HasFile returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


