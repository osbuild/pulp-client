# CollectionSummaryResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PulpHref** | Pointer to **string** |  | [optional] [readonly] 
**Namespace** | Pointer to **string** | The namespace of the collection. | [optional] [readonly] 
**Name** | Pointer to **string** | The name of the collection. | [optional] [readonly] 
**Version** | Pointer to **string** | The version of the collection. | [optional] [readonly] 
**RequiresAnsible** | Pointer to **NullableString** | The version of Ansible required to use the collection. Multiple versions can be separated with a comma. | [optional] [readonly] 
**PulpCreated** | Pointer to **time.Time** | Timestamp of creation. | [optional] [readonly] 
**Contents** | Pointer to **map[string]interface{}** | A JSON field with data about the contents. | [optional] [readonly] 
**Dependencies** | Pointer to **map[string]interface{}** | A dict declaring Collections that this collection requires to be installed for it to be usable. | [optional] [readonly] 
**Description** | Pointer to **string** | A short summary description of the collection. | [optional] [readonly] 
**Tags** | Pointer to [**[]AnsibleTagResponse**](AnsibleTagResponse.md) |  | [optional] [readonly] 

## Methods

### NewCollectionSummaryResponse

`func NewCollectionSummaryResponse() *CollectionSummaryResponse`

NewCollectionSummaryResponse instantiates a new CollectionSummaryResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCollectionSummaryResponseWithDefaults

`func NewCollectionSummaryResponseWithDefaults() *CollectionSummaryResponse`

NewCollectionSummaryResponseWithDefaults instantiates a new CollectionSummaryResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPulpHref

`func (o *CollectionSummaryResponse) GetPulpHref() string`

GetPulpHref returns the PulpHref field if non-nil, zero value otherwise.

### GetPulpHrefOk

`func (o *CollectionSummaryResponse) GetPulpHrefOk() (*string, bool)`

GetPulpHrefOk returns a tuple with the PulpHref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpHref

`func (o *CollectionSummaryResponse) SetPulpHref(v string)`

SetPulpHref sets PulpHref field to given value.

### HasPulpHref

`func (o *CollectionSummaryResponse) HasPulpHref() bool`

HasPulpHref returns a boolean if a field has been set.

### GetNamespace

`func (o *CollectionSummaryResponse) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *CollectionSummaryResponse) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *CollectionSummaryResponse) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *CollectionSummaryResponse) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetName

`func (o *CollectionSummaryResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CollectionSummaryResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CollectionSummaryResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CollectionSummaryResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetVersion

`func (o *CollectionSummaryResponse) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *CollectionSummaryResponse) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *CollectionSummaryResponse) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *CollectionSummaryResponse) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetRequiresAnsible

`func (o *CollectionSummaryResponse) GetRequiresAnsible() string`

GetRequiresAnsible returns the RequiresAnsible field if non-nil, zero value otherwise.

### GetRequiresAnsibleOk

`func (o *CollectionSummaryResponse) GetRequiresAnsibleOk() (*string, bool)`

GetRequiresAnsibleOk returns a tuple with the RequiresAnsible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiresAnsible

`func (o *CollectionSummaryResponse) SetRequiresAnsible(v string)`

SetRequiresAnsible sets RequiresAnsible field to given value.

### HasRequiresAnsible

`func (o *CollectionSummaryResponse) HasRequiresAnsible() bool`

HasRequiresAnsible returns a boolean if a field has been set.

### SetRequiresAnsibleNil

`func (o *CollectionSummaryResponse) SetRequiresAnsibleNil(b bool)`

 SetRequiresAnsibleNil sets the value for RequiresAnsible to be an explicit nil

### UnsetRequiresAnsible
`func (o *CollectionSummaryResponse) UnsetRequiresAnsible()`

UnsetRequiresAnsible ensures that no value is present for RequiresAnsible, not even an explicit nil
### GetPulpCreated

`func (o *CollectionSummaryResponse) GetPulpCreated() time.Time`

GetPulpCreated returns the PulpCreated field if non-nil, zero value otherwise.

### GetPulpCreatedOk

`func (o *CollectionSummaryResponse) GetPulpCreatedOk() (*time.Time, bool)`

GetPulpCreatedOk returns a tuple with the PulpCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpCreated

`func (o *CollectionSummaryResponse) SetPulpCreated(v time.Time)`

SetPulpCreated sets PulpCreated field to given value.

### HasPulpCreated

`func (o *CollectionSummaryResponse) HasPulpCreated() bool`

HasPulpCreated returns a boolean if a field has been set.

### GetContents

`func (o *CollectionSummaryResponse) GetContents() map[string]interface{}`

GetContents returns the Contents field if non-nil, zero value otherwise.

### GetContentsOk

`func (o *CollectionSummaryResponse) GetContentsOk() (*map[string]interface{}, bool)`

GetContentsOk returns a tuple with the Contents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContents

`func (o *CollectionSummaryResponse) SetContents(v map[string]interface{})`

SetContents sets Contents field to given value.

### HasContents

`func (o *CollectionSummaryResponse) HasContents() bool`

HasContents returns a boolean if a field has been set.

### GetDependencies

`func (o *CollectionSummaryResponse) GetDependencies() map[string]interface{}`

GetDependencies returns the Dependencies field if non-nil, zero value otherwise.

### GetDependenciesOk

`func (o *CollectionSummaryResponse) GetDependenciesOk() (*map[string]interface{}, bool)`

GetDependenciesOk returns a tuple with the Dependencies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependencies

`func (o *CollectionSummaryResponse) SetDependencies(v map[string]interface{})`

SetDependencies sets Dependencies field to given value.

### HasDependencies

`func (o *CollectionSummaryResponse) HasDependencies() bool`

HasDependencies returns a boolean if a field has been set.

### GetDescription

`func (o *CollectionSummaryResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CollectionSummaryResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CollectionSummaryResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CollectionSummaryResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetTags

`func (o *CollectionSummaryResponse) GetTags() []AnsibleTagResponse`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CollectionSummaryResponse) GetTagsOk() (*[]AnsibleTagResponse, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CollectionSummaryResponse) SetTags(v []AnsibleTagResponse)`

SetTags sets Tags field to given value.

### HasTags

`func (o *CollectionSummaryResponse) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


