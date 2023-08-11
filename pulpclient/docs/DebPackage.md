# DebPackage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Repository** | Pointer to **string** | A URI of a repository the new content unit should be associated with. | [optional] 
**Artifact** | Pointer to **string** | Artifact file representing the physical content | [optional] 
**RelativePath** | Pointer to **string** | Path where the artifact is located relative to distributions base_path | [optional] 
**File** | Pointer to ***os.File** | An uploaded file that may be turned into the artifact of the content unit. | [optional] 
**Upload** | Pointer to **string** | An uncommitted upload that may be turned into the artifact of the content unit. | [optional] 
**Distribution** | Pointer to **string** | Name of the distribution. | [optional] 
**Component** | Pointer to **string** | Name of the component. | [optional] 

## Methods

### NewDebPackage

`func NewDebPackage() *DebPackage`

NewDebPackage instantiates a new DebPackage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDebPackageWithDefaults

`func NewDebPackageWithDefaults() *DebPackage`

NewDebPackageWithDefaults instantiates a new DebPackage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRepository

`func (o *DebPackage) GetRepository() string`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *DebPackage) GetRepositoryOk() (*string, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *DebPackage) SetRepository(v string)`

SetRepository sets Repository field to given value.

### HasRepository

`func (o *DebPackage) HasRepository() bool`

HasRepository returns a boolean if a field has been set.

### GetArtifact

`func (o *DebPackage) GetArtifact() string`

GetArtifact returns the Artifact field if non-nil, zero value otherwise.

### GetArtifactOk

`func (o *DebPackage) GetArtifactOk() (*string, bool)`

GetArtifactOk returns a tuple with the Artifact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifact

`func (o *DebPackage) SetArtifact(v string)`

SetArtifact sets Artifact field to given value.

### HasArtifact

`func (o *DebPackage) HasArtifact() bool`

HasArtifact returns a boolean if a field has been set.

### GetRelativePath

`func (o *DebPackage) GetRelativePath() string`

GetRelativePath returns the RelativePath field if non-nil, zero value otherwise.

### GetRelativePathOk

`func (o *DebPackage) GetRelativePathOk() (*string, bool)`

GetRelativePathOk returns a tuple with the RelativePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelativePath

`func (o *DebPackage) SetRelativePath(v string)`

SetRelativePath sets RelativePath field to given value.

### HasRelativePath

`func (o *DebPackage) HasRelativePath() bool`

HasRelativePath returns a boolean if a field has been set.

### GetFile

`func (o *DebPackage) GetFile() *os.File`

GetFile returns the File field if non-nil, zero value otherwise.

### GetFileOk

`func (o *DebPackage) GetFileOk() (**os.File, bool)`

GetFileOk returns a tuple with the File field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFile

`func (o *DebPackage) SetFile(v *os.File)`

SetFile sets File field to given value.

### HasFile

`func (o *DebPackage) HasFile() bool`

HasFile returns a boolean if a field has been set.

### GetUpload

`func (o *DebPackage) GetUpload() string`

GetUpload returns the Upload field if non-nil, zero value otherwise.

### GetUploadOk

`func (o *DebPackage) GetUploadOk() (*string, bool)`

GetUploadOk returns a tuple with the Upload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpload

`func (o *DebPackage) SetUpload(v string)`

SetUpload sets Upload field to given value.

### HasUpload

`func (o *DebPackage) HasUpload() bool`

HasUpload returns a boolean if a field has been set.

### GetDistribution

`func (o *DebPackage) GetDistribution() string`

GetDistribution returns the Distribution field if non-nil, zero value otherwise.

### GetDistributionOk

`func (o *DebPackage) GetDistributionOk() (*string, bool)`

GetDistributionOk returns a tuple with the Distribution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistribution

`func (o *DebPackage) SetDistribution(v string)`

SetDistribution sets Distribution field to given value.

### HasDistribution

`func (o *DebPackage) HasDistribution() bool`

HasDistribution returns a boolean if a field has been set.

### GetComponent

`func (o *DebPackage) GetComponent() string`

GetComponent returns the Component field if non-nil, zero value otherwise.

### GetComponentOk

`func (o *DebPackage) GetComponentOk() (*string, bool)`

GetComponentOk returns a tuple with the Component field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponent

`func (o *DebPackage) SetComponent(v string)`

SetComponent sets Component field to given value.

### HasComponent

`func (o *DebPackage) HasComponent() bool`

HasComponent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


