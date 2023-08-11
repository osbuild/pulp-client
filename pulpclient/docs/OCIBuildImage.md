# OCIBuildImage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContainerfileArtifact** | Pointer to **string** | Artifact representing the Containerfile that should be used to run podman-build. | [optional] 
**Containerfile** | Pointer to ***os.File** | An uploaded Containerfile that should be used to run podman-build. | [optional] 
**Tag** | Pointer to **string** | A tag name for the new image being built. | [optional] [default to "latest"]
**Artifacts** | Pointer to **map[string]interface{}** | A JSON string where each key is an artifact href and the value is it&#39;s relative path (name) inside the /pulp_working_directory of the build container executing the Containerfile. | [optional] 

## Methods

### NewOCIBuildImage

`func NewOCIBuildImage() *OCIBuildImage`

NewOCIBuildImage instantiates a new OCIBuildImage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOCIBuildImageWithDefaults

`func NewOCIBuildImageWithDefaults() *OCIBuildImage`

NewOCIBuildImageWithDefaults instantiates a new OCIBuildImage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContainerfileArtifact

`func (o *OCIBuildImage) GetContainerfileArtifact() string`

GetContainerfileArtifact returns the ContainerfileArtifact field if non-nil, zero value otherwise.

### GetContainerfileArtifactOk

`func (o *OCIBuildImage) GetContainerfileArtifactOk() (*string, bool)`

GetContainerfileArtifactOk returns a tuple with the ContainerfileArtifact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerfileArtifact

`func (o *OCIBuildImage) SetContainerfileArtifact(v string)`

SetContainerfileArtifact sets ContainerfileArtifact field to given value.

### HasContainerfileArtifact

`func (o *OCIBuildImage) HasContainerfileArtifact() bool`

HasContainerfileArtifact returns a boolean if a field has been set.

### GetContainerfile

`func (o *OCIBuildImage) GetContainerfile() *os.File`

GetContainerfile returns the Containerfile field if non-nil, zero value otherwise.

### GetContainerfileOk

`func (o *OCIBuildImage) GetContainerfileOk() (**os.File, bool)`

GetContainerfileOk returns a tuple with the Containerfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerfile

`func (o *OCIBuildImage) SetContainerfile(v *os.File)`

SetContainerfile sets Containerfile field to given value.

### HasContainerfile

`func (o *OCIBuildImage) HasContainerfile() bool`

HasContainerfile returns a boolean if a field has been set.

### GetTag

`func (o *OCIBuildImage) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *OCIBuildImage) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *OCIBuildImage) SetTag(v string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *OCIBuildImage) HasTag() bool`

HasTag returns a boolean if a field has been set.

### GetArtifacts

`func (o *OCIBuildImage) GetArtifacts() map[string]interface{}`

GetArtifacts returns the Artifacts field if non-nil, zero value otherwise.

### GetArtifactsOk

`func (o *OCIBuildImage) GetArtifactsOk() (*map[string]interface{}, bool)`

GetArtifactsOk returns a tuple with the Artifacts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifacts

`func (o *OCIBuildImage) SetArtifacts(v map[string]interface{})`

SetArtifacts sets Artifacts field to given value.

### HasArtifacts

`func (o *OCIBuildImage) HasArtifacts() bool`

HasArtifacts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


