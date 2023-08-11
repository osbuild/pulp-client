# DebGenericContent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Repository** | Pointer to **string** | A URI of a repository the new content unit should be associated with. | [optional] 
**Artifact** | Pointer to **string** | Artifact file representing the physical content | [optional] 
**RelativePath** | **string** | Path where the artifact is located relative to distributions base_path | 
**File** | Pointer to ***os.File** | An uploaded file that may be turned into the artifact of the content unit. | [optional] 
**Upload** | Pointer to **string** | An uncommitted upload that may be turned into the artifact of the content unit. | [optional] 

## Methods

### NewDebGenericContent

`func NewDebGenericContent(relativePath string, ) *DebGenericContent`

NewDebGenericContent instantiates a new DebGenericContent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDebGenericContentWithDefaults

`func NewDebGenericContentWithDefaults() *DebGenericContent`

NewDebGenericContentWithDefaults instantiates a new DebGenericContent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRepository

`func (o *DebGenericContent) GetRepository() string`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *DebGenericContent) GetRepositoryOk() (*string, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *DebGenericContent) SetRepository(v string)`

SetRepository sets Repository field to given value.

### HasRepository

`func (o *DebGenericContent) HasRepository() bool`

HasRepository returns a boolean if a field has been set.

### GetArtifact

`func (o *DebGenericContent) GetArtifact() string`

GetArtifact returns the Artifact field if non-nil, zero value otherwise.

### GetArtifactOk

`func (o *DebGenericContent) GetArtifactOk() (*string, bool)`

GetArtifactOk returns a tuple with the Artifact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifact

`func (o *DebGenericContent) SetArtifact(v string)`

SetArtifact sets Artifact field to given value.

### HasArtifact

`func (o *DebGenericContent) HasArtifact() bool`

HasArtifact returns a boolean if a field has been set.

### GetRelativePath

`func (o *DebGenericContent) GetRelativePath() string`

GetRelativePath returns the RelativePath field if non-nil, zero value otherwise.

### GetRelativePathOk

`func (o *DebGenericContent) GetRelativePathOk() (*string, bool)`

GetRelativePathOk returns a tuple with the RelativePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelativePath

`func (o *DebGenericContent) SetRelativePath(v string)`

SetRelativePath sets RelativePath field to given value.


### GetFile

`func (o *DebGenericContent) GetFile() *os.File`

GetFile returns the File field if non-nil, zero value otherwise.

### GetFileOk

`func (o *DebGenericContent) GetFileOk() (**os.File, bool)`

GetFileOk returns a tuple with the File field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFile

`func (o *DebGenericContent) SetFile(v *os.File)`

SetFile sets File field to given value.

### HasFile

`func (o *DebGenericContent) HasFile() bool`

HasFile returns a boolean if a field has been set.

### GetUpload

`func (o *DebGenericContent) GetUpload() string`

GetUpload returns the Upload field if non-nil, zero value otherwise.

### GetUploadOk

`func (o *DebGenericContent) GetUploadOk() (*string, bool)`

GetUploadOk returns a tuple with the Upload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpload

`func (o *DebGenericContent) SetUpload(v string)`

SetUpload sets Upload field to given value.

### HasUpload

`func (o *DebGenericContent) HasUpload() bool`

HasUpload returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


