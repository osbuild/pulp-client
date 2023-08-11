# AnsibleCollectionVersion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Upload** | Pointer to **string** | An uncommitted upload that may be turned into the artifact of the content unit. | [optional] 
**Artifact** | Pointer to **string** | Artifact file representing the physical content | [optional] 
**Repository** | Pointer to **string** | A URI of a repository the new content unit should be associated with. | [optional] 
**File** | Pointer to ***os.File** | An uploaded file that may be turned into the artifact of the content unit. | [optional] 
**ExpectedName** | Pointer to **string** | The name of the collection. | [optional] 
**ExpectedNamespace** | Pointer to **string** | The namespace of the collection. | [optional] 
**ExpectedVersion** | Pointer to **string** | The version of the collection. | [optional] 

## Methods

### NewAnsibleCollectionVersion

`func NewAnsibleCollectionVersion() *AnsibleCollectionVersion`

NewAnsibleCollectionVersion instantiates a new AnsibleCollectionVersion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnsibleCollectionVersionWithDefaults

`func NewAnsibleCollectionVersionWithDefaults() *AnsibleCollectionVersion`

NewAnsibleCollectionVersionWithDefaults instantiates a new AnsibleCollectionVersion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUpload

`func (o *AnsibleCollectionVersion) GetUpload() string`

GetUpload returns the Upload field if non-nil, zero value otherwise.

### GetUploadOk

`func (o *AnsibleCollectionVersion) GetUploadOk() (*string, bool)`

GetUploadOk returns a tuple with the Upload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpload

`func (o *AnsibleCollectionVersion) SetUpload(v string)`

SetUpload sets Upload field to given value.

### HasUpload

`func (o *AnsibleCollectionVersion) HasUpload() bool`

HasUpload returns a boolean if a field has been set.

### GetArtifact

`func (o *AnsibleCollectionVersion) GetArtifact() string`

GetArtifact returns the Artifact field if non-nil, zero value otherwise.

### GetArtifactOk

`func (o *AnsibleCollectionVersion) GetArtifactOk() (*string, bool)`

GetArtifactOk returns a tuple with the Artifact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifact

`func (o *AnsibleCollectionVersion) SetArtifact(v string)`

SetArtifact sets Artifact field to given value.

### HasArtifact

`func (o *AnsibleCollectionVersion) HasArtifact() bool`

HasArtifact returns a boolean if a field has been set.

### GetRepository

`func (o *AnsibleCollectionVersion) GetRepository() string`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *AnsibleCollectionVersion) GetRepositoryOk() (*string, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *AnsibleCollectionVersion) SetRepository(v string)`

SetRepository sets Repository field to given value.

### HasRepository

`func (o *AnsibleCollectionVersion) HasRepository() bool`

HasRepository returns a boolean if a field has been set.

### GetFile

`func (o *AnsibleCollectionVersion) GetFile() *os.File`

GetFile returns the File field if non-nil, zero value otherwise.

### GetFileOk

`func (o *AnsibleCollectionVersion) GetFileOk() (**os.File, bool)`

GetFileOk returns a tuple with the File field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFile

`func (o *AnsibleCollectionVersion) SetFile(v *os.File)`

SetFile sets File field to given value.

### HasFile

`func (o *AnsibleCollectionVersion) HasFile() bool`

HasFile returns a boolean if a field has been set.

### GetExpectedName

`func (o *AnsibleCollectionVersion) GetExpectedName() string`

GetExpectedName returns the ExpectedName field if non-nil, zero value otherwise.

### GetExpectedNameOk

`func (o *AnsibleCollectionVersion) GetExpectedNameOk() (*string, bool)`

GetExpectedNameOk returns a tuple with the ExpectedName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedName

`func (o *AnsibleCollectionVersion) SetExpectedName(v string)`

SetExpectedName sets ExpectedName field to given value.

### HasExpectedName

`func (o *AnsibleCollectionVersion) HasExpectedName() bool`

HasExpectedName returns a boolean if a field has been set.

### GetExpectedNamespace

`func (o *AnsibleCollectionVersion) GetExpectedNamespace() string`

GetExpectedNamespace returns the ExpectedNamespace field if non-nil, zero value otherwise.

### GetExpectedNamespaceOk

`func (o *AnsibleCollectionVersion) GetExpectedNamespaceOk() (*string, bool)`

GetExpectedNamespaceOk returns a tuple with the ExpectedNamespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedNamespace

`func (o *AnsibleCollectionVersion) SetExpectedNamespace(v string)`

SetExpectedNamespace sets ExpectedNamespace field to given value.

### HasExpectedNamespace

`func (o *AnsibleCollectionVersion) HasExpectedNamespace() bool`

HasExpectedNamespace returns a boolean if a field has been set.

### GetExpectedVersion

`func (o *AnsibleCollectionVersion) GetExpectedVersion() string`

GetExpectedVersion returns the ExpectedVersion field if non-nil, zero value otherwise.

### GetExpectedVersionOk

`func (o *AnsibleCollectionVersion) GetExpectedVersionOk() (*string, bool)`

GetExpectedVersionOk returns a tuple with the ExpectedVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedVersion

`func (o *AnsibleCollectionVersion) SetExpectedVersion(v string)`

SetExpectedVersion sets ExpectedVersion field to given value.

### HasExpectedVersion

`func (o *AnsibleCollectionVersion) HasExpectedVersion() bool`

HasExpectedVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


