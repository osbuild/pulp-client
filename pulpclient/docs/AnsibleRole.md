# AnsibleRole

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Repository** | Pointer to **string** | A URI of a repository the new content unit should be associated with. | [optional] 
**Artifact** | **string** | Artifact file representing the physical content | 
**Version** | **string** |  | 
**Name** | **string** |  | 
**Namespace** | **string** |  | 

## Methods

### NewAnsibleRole

`func NewAnsibleRole(artifact string, version string, name string, namespace string, ) *AnsibleRole`

NewAnsibleRole instantiates a new AnsibleRole object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnsibleRoleWithDefaults

`func NewAnsibleRoleWithDefaults() *AnsibleRole`

NewAnsibleRoleWithDefaults instantiates a new AnsibleRole object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRepository

`func (o *AnsibleRole) GetRepository() string`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *AnsibleRole) GetRepositoryOk() (*string, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *AnsibleRole) SetRepository(v string)`

SetRepository sets Repository field to given value.

### HasRepository

`func (o *AnsibleRole) HasRepository() bool`

HasRepository returns a boolean if a field has been set.

### GetArtifact

`func (o *AnsibleRole) GetArtifact() string`

GetArtifact returns the Artifact field if non-nil, zero value otherwise.

### GetArtifactOk

`func (o *AnsibleRole) GetArtifactOk() (*string, bool)`

GetArtifactOk returns a tuple with the Artifact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifact

`func (o *AnsibleRole) SetArtifact(v string)`

SetArtifact sets Artifact field to given value.


### GetVersion

`func (o *AnsibleRole) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *AnsibleRole) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *AnsibleRole) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetName

`func (o *AnsibleRole) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AnsibleRole) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AnsibleRole) SetName(v string)`

SetName sets Name field to given value.


### GetNamespace

`func (o *AnsibleRole) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *AnsibleRole) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *AnsibleRole) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


