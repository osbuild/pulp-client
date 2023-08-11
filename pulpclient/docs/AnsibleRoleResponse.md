# AnsibleRoleResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PulpHref** | Pointer to **string** |  | [optional] [readonly] 
**Artifact** | **string** | Artifact file representing the physical content | 
**PulpCreated** | Pointer to **time.Time** | Timestamp of creation. | [optional] [readonly] 
**Version** | **string** |  | 
**Name** | **string** |  | 
**Namespace** | **string** |  | 

## Methods

### NewAnsibleRoleResponse

`func NewAnsibleRoleResponse(artifact string, version string, name string, namespace string, ) *AnsibleRoleResponse`

NewAnsibleRoleResponse instantiates a new AnsibleRoleResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnsibleRoleResponseWithDefaults

`func NewAnsibleRoleResponseWithDefaults() *AnsibleRoleResponse`

NewAnsibleRoleResponseWithDefaults instantiates a new AnsibleRoleResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPulpHref

`func (o *AnsibleRoleResponse) GetPulpHref() string`

GetPulpHref returns the PulpHref field if non-nil, zero value otherwise.

### GetPulpHrefOk

`func (o *AnsibleRoleResponse) GetPulpHrefOk() (*string, bool)`

GetPulpHrefOk returns a tuple with the PulpHref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpHref

`func (o *AnsibleRoleResponse) SetPulpHref(v string)`

SetPulpHref sets PulpHref field to given value.

### HasPulpHref

`func (o *AnsibleRoleResponse) HasPulpHref() bool`

HasPulpHref returns a boolean if a field has been set.

### GetArtifact

`func (o *AnsibleRoleResponse) GetArtifact() string`

GetArtifact returns the Artifact field if non-nil, zero value otherwise.

### GetArtifactOk

`func (o *AnsibleRoleResponse) GetArtifactOk() (*string, bool)`

GetArtifactOk returns a tuple with the Artifact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifact

`func (o *AnsibleRoleResponse) SetArtifact(v string)`

SetArtifact sets Artifact field to given value.


### GetPulpCreated

`func (o *AnsibleRoleResponse) GetPulpCreated() time.Time`

GetPulpCreated returns the PulpCreated field if non-nil, zero value otherwise.

### GetPulpCreatedOk

`func (o *AnsibleRoleResponse) GetPulpCreatedOk() (*time.Time, bool)`

GetPulpCreatedOk returns a tuple with the PulpCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpCreated

`func (o *AnsibleRoleResponse) SetPulpCreated(v time.Time)`

SetPulpCreated sets PulpCreated field to given value.

### HasPulpCreated

`func (o *AnsibleRoleResponse) HasPulpCreated() bool`

HasPulpCreated returns a boolean if a field has been set.

### GetVersion

`func (o *AnsibleRoleResponse) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *AnsibleRoleResponse) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *AnsibleRoleResponse) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetName

`func (o *AnsibleRoleResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AnsibleRoleResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AnsibleRoleResponse) SetName(v string)`

SetName sets Name field to given value.


### GetNamespace

`func (o *AnsibleRoleResponse) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *AnsibleRoleResponse) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *AnsibleRoleResponse) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


