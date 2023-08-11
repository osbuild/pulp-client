# CollectionVersionSearchList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Repository** | [**Repository**](Repository.md) |  | 
**CollectionVersion** | **map[string]interface{}** | Collection Version serializer without docs blob. | 
**NamespaceMetadata** | [**NullableAnsibleAnsibleNamespaceMetadata**](AnsibleAnsibleNamespaceMetadata.md) |  | 
**IsHighest** | **bool** |  | 
**IsDeprecated** | **bool** |  | 
**IsSigned** | **bool** |  | 

## Methods

### NewCollectionVersionSearchList

`func NewCollectionVersionSearchList(repository Repository, collectionVersion map[string]interface{}, namespaceMetadata NullableAnsibleAnsibleNamespaceMetadata, isHighest bool, isDeprecated bool, isSigned bool, ) *CollectionVersionSearchList`

NewCollectionVersionSearchList instantiates a new CollectionVersionSearchList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCollectionVersionSearchListWithDefaults

`func NewCollectionVersionSearchListWithDefaults() *CollectionVersionSearchList`

NewCollectionVersionSearchListWithDefaults instantiates a new CollectionVersionSearchList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRepository

`func (o *CollectionVersionSearchList) GetRepository() Repository`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *CollectionVersionSearchList) GetRepositoryOk() (*Repository, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *CollectionVersionSearchList) SetRepository(v Repository)`

SetRepository sets Repository field to given value.


### GetCollectionVersion

`func (o *CollectionVersionSearchList) GetCollectionVersion() map[string]interface{}`

GetCollectionVersion returns the CollectionVersion field if non-nil, zero value otherwise.

### GetCollectionVersionOk

`func (o *CollectionVersionSearchList) GetCollectionVersionOk() (*map[string]interface{}, bool)`

GetCollectionVersionOk returns a tuple with the CollectionVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionVersion

`func (o *CollectionVersionSearchList) SetCollectionVersion(v map[string]interface{})`

SetCollectionVersion sets CollectionVersion field to given value.


### GetNamespaceMetadata

`func (o *CollectionVersionSearchList) GetNamespaceMetadata() AnsibleAnsibleNamespaceMetadata`

GetNamespaceMetadata returns the NamespaceMetadata field if non-nil, zero value otherwise.

### GetNamespaceMetadataOk

`func (o *CollectionVersionSearchList) GetNamespaceMetadataOk() (*AnsibleAnsibleNamespaceMetadata, bool)`

GetNamespaceMetadataOk returns a tuple with the NamespaceMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespaceMetadata

`func (o *CollectionVersionSearchList) SetNamespaceMetadata(v AnsibleAnsibleNamespaceMetadata)`

SetNamespaceMetadata sets NamespaceMetadata field to given value.


### SetNamespaceMetadataNil

`func (o *CollectionVersionSearchList) SetNamespaceMetadataNil(b bool)`

 SetNamespaceMetadataNil sets the value for NamespaceMetadata to be an explicit nil

### UnsetNamespaceMetadata
`func (o *CollectionVersionSearchList) UnsetNamespaceMetadata()`

UnsetNamespaceMetadata ensures that no value is present for NamespaceMetadata, not even an explicit nil
### GetIsHighest

`func (o *CollectionVersionSearchList) GetIsHighest() bool`

GetIsHighest returns the IsHighest field if non-nil, zero value otherwise.

### GetIsHighestOk

`func (o *CollectionVersionSearchList) GetIsHighestOk() (*bool, bool)`

GetIsHighestOk returns a tuple with the IsHighest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsHighest

`func (o *CollectionVersionSearchList) SetIsHighest(v bool)`

SetIsHighest sets IsHighest field to given value.


### GetIsDeprecated

`func (o *CollectionVersionSearchList) GetIsDeprecated() bool`

GetIsDeprecated returns the IsDeprecated field if non-nil, zero value otherwise.

### GetIsDeprecatedOk

`func (o *CollectionVersionSearchList) GetIsDeprecatedOk() (*bool, bool)`

GetIsDeprecatedOk returns a tuple with the IsDeprecated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeprecated

`func (o *CollectionVersionSearchList) SetIsDeprecated(v bool)`

SetIsDeprecated sets IsDeprecated field to given value.


### GetIsSigned

`func (o *CollectionVersionSearchList) GetIsSigned() bool`

GetIsSigned returns the IsSigned field if non-nil, zero value otherwise.

### GetIsSignedOk

`func (o *CollectionVersionSearchList) GetIsSignedOk() (*bool, bool)`

GetIsSignedOk returns a tuple with the IsSigned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSigned

`func (o *CollectionVersionSearchList) SetIsSigned(v bool)`

SetIsSigned sets IsSigned field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


