# DebInstallerPackage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Repository** | Pointer to **string** | A URI of a repository the new content unit should be associated with. | [optional] 
**Artifact** | Pointer to **string** | Artifact file representing the physical content | [optional] 
**RelativePath** | Pointer to **string** | Path where the artifact is located relative to distributions base_path | [optional] 
**File** | Pointer to ***os.File** | An uploaded file that may be turned into the artifact of the content unit. | [optional] 
**Upload** | Pointer to **string** | An uncommitted upload that may be turned into the artifact of the content unit. | [optional] 

## Methods

### NewDebInstallerPackage

`func NewDebInstallerPackage() *DebInstallerPackage`

NewDebInstallerPackage instantiates a new DebInstallerPackage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDebInstallerPackageWithDefaults

`func NewDebInstallerPackageWithDefaults() *DebInstallerPackage`

NewDebInstallerPackageWithDefaults instantiates a new DebInstallerPackage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRepository

`func (o *DebInstallerPackage) GetRepository() string`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *DebInstallerPackage) GetRepositoryOk() (*string, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *DebInstallerPackage) SetRepository(v string)`

SetRepository sets Repository field to given value.

### HasRepository

`func (o *DebInstallerPackage) HasRepository() bool`

HasRepository returns a boolean if a field has been set.

### GetArtifact

`func (o *DebInstallerPackage) GetArtifact() string`

GetArtifact returns the Artifact field if non-nil, zero value otherwise.

### GetArtifactOk

`func (o *DebInstallerPackage) GetArtifactOk() (*string, bool)`

GetArtifactOk returns a tuple with the Artifact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifact

`func (o *DebInstallerPackage) SetArtifact(v string)`

SetArtifact sets Artifact field to given value.

### HasArtifact

`func (o *DebInstallerPackage) HasArtifact() bool`

HasArtifact returns a boolean if a field has been set.

### GetRelativePath

`func (o *DebInstallerPackage) GetRelativePath() string`

GetRelativePath returns the RelativePath field if non-nil, zero value otherwise.

### GetRelativePathOk

`func (o *DebInstallerPackage) GetRelativePathOk() (*string, bool)`

GetRelativePathOk returns a tuple with the RelativePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelativePath

`func (o *DebInstallerPackage) SetRelativePath(v string)`

SetRelativePath sets RelativePath field to given value.

### HasRelativePath

`func (o *DebInstallerPackage) HasRelativePath() bool`

HasRelativePath returns a boolean if a field has been set.

### GetFile

`func (o *DebInstallerPackage) GetFile() *os.File`

GetFile returns the File field if non-nil, zero value otherwise.

### GetFileOk

`func (o *DebInstallerPackage) GetFileOk() (**os.File, bool)`

GetFileOk returns a tuple with the File field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFile

`func (o *DebInstallerPackage) SetFile(v *os.File)`

SetFile sets File field to given value.

### HasFile

`func (o *DebInstallerPackage) HasFile() bool`

HasFile returns a boolean if a field has been set.

### GetUpload

`func (o *DebInstallerPackage) GetUpload() string`

GetUpload returns the Upload field if non-nil, zero value otherwise.

### GetUploadOk

`func (o *DebInstallerPackage) GetUploadOk() (*string, bool)`

GetUploadOk returns a tuple with the Upload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpload

`func (o *DebInstallerPackage) SetUpload(v string)`

SetUpload sets Upload field to given value.

### HasUpload

`func (o *DebInstallerPackage) HasUpload() bool`

HasUpload returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


