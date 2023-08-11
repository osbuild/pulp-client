# CollectionVersionSearchListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Repository** | [**RepositoryResponse**](RepositoryResponse.md) |  | 
**CollectionVersion** | [**CollectionSummaryResponse**](CollectionSummaryResponse.md) |  | 
**RepositoryVersion** | Pointer to **string** |  | [optional] [readonly] 
**NamespaceMetadata** | [**NullableAnsibleAnsibleNamespaceMetadataResponse**](AnsibleAnsibleNamespaceMetadataResponse.md) |  | 
**IsHighest** | **bool** |  | 
**IsDeprecated** | **bool** |  | 
**IsSigned** | **bool** |  | 

## Methods

### NewCollectionVersionSearchListResponse

`func NewCollectionVersionSearchListResponse(repository RepositoryResponse, collectionVersion CollectionSummaryResponse, namespaceMetadata NullableAnsibleAnsibleNamespaceMetadataResponse, isHighest bool, isDeprecated bool, isSigned bool, ) *CollectionVersionSearchListResponse`

NewCollectionVersionSearchListResponse instantiates a new CollectionVersionSearchListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCollectionVersionSearchListResponseWithDefaults

`func NewCollectionVersionSearchListResponseWithDefaults() *CollectionVersionSearchListResponse`

NewCollectionVersionSearchListResponseWithDefaults instantiates a new CollectionVersionSearchListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRepository

`func (o *CollectionVersionSearchListResponse) GetRepository() RepositoryResponse`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *CollectionVersionSearchListResponse) GetRepositoryOk() (*RepositoryResponse, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *CollectionVersionSearchListResponse) SetRepository(v RepositoryResponse)`

SetRepository sets Repository field to given value.


### GetCollectionVersion

`func (o *CollectionVersionSearchListResponse) GetCollectionVersion() CollectionSummaryResponse`

GetCollectionVersion returns the CollectionVersion field if non-nil, zero value otherwise.

### GetCollectionVersionOk

`func (o *CollectionVersionSearchListResponse) GetCollectionVersionOk() (*CollectionSummaryResponse, bool)`

GetCollectionVersionOk returns a tuple with the CollectionVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionVersion

`func (o *CollectionVersionSearchListResponse) SetCollectionVersion(v CollectionSummaryResponse)`

SetCollectionVersion sets CollectionVersion field to given value.


### GetRepositoryVersion

`func (o *CollectionVersionSearchListResponse) GetRepositoryVersion() string`

GetRepositoryVersion returns the RepositoryVersion field if non-nil, zero value otherwise.

### GetRepositoryVersionOk

`func (o *CollectionVersionSearchListResponse) GetRepositoryVersionOk() (*string, bool)`

GetRepositoryVersionOk returns a tuple with the RepositoryVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepositoryVersion

`func (o *CollectionVersionSearchListResponse) SetRepositoryVersion(v string)`

SetRepositoryVersion sets RepositoryVersion field to given value.

### HasRepositoryVersion

`func (o *CollectionVersionSearchListResponse) HasRepositoryVersion() bool`

HasRepositoryVersion returns a boolean if a field has been set.

### GetNamespaceMetadata

`func (o *CollectionVersionSearchListResponse) GetNamespaceMetadata() AnsibleAnsibleNamespaceMetadataResponse`

GetNamespaceMetadata returns the NamespaceMetadata field if non-nil, zero value otherwise.

### GetNamespaceMetadataOk

`func (o *CollectionVersionSearchListResponse) GetNamespaceMetadataOk() (*AnsibleAnsibleNamespaceMetadataResponse, bool)`

GetNamespaceMetadataOk returns a tuple with the NamespaceMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespaceMetadata

`func (o *CollectionVersionSearchListResponse) SetNamespaceMetadata(v AnsibleAnsibleNamespaceMetadataResponse)`

SetNamespaceMetadata sets NamespaceMetadata field to given value.


### SetNamespaceMetadataNil

`func (o *CollectionVersionSearchListResponse) SetNamespaceMetadataNil(b bool)`

 SetNamespaceMetadataNil sets the value for NamespaceMetadata to be an explicit nil

### UnsetNamespaceMetadata
`func (o *CollectionVersionSearchListResponse) UnsetNamespaceMetadata()`

UnsetNamespaceMetadata ensures that no value is present for NamespaceMetadata, not even an explicit nil
### GetIsHighest

`func (o *CollectionVersionSearchListResponse) GetIsHighest() bool`

GetIsHighest returns the IsHighest field if non-nil, zero value otherwise.

### GetIsHighestOk

`func (o *CollectionVersionSearchListResponse) GetIsHighestOk() (*bool, bool)`

GetIsHighestOk returns a tuple with the IsHighest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsHighest

`func (o *CollectionVersionSearchListResponse) SetIsHighest(v bool)`

SetIsHighest sets IsHighest field to given value.


### GetIsDeprecated

`func (o *CollectionVersionSearchListResponse) GetIsDeprecated() bool`

GetIsDeprecated returns the IsDeprecated field if non-nil, zero value otherwise.

### GetIsDeprecatedOk

`func (o *CollectionVersionSearchListResponse) GetIsDeprecatedOk() (*bool, bool)`

GetIsDeprecatedOk returns a tuple with the IsDeprecated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeprecated

`func (o *CollectionVersionSearchListResponse) SetIsDeprecated(v bool)`

SetIsDeprecated sets IsDeprecated field to given value.


### GetIsSigned

`func (o *CollectionVersionSearchListResponse) GetIsSigned() bool`

GetIsSigned returns the IsSigned field if non-nil, zero value otherwise.

### GetIsSignedOk

`func (o *CollectionVersionSearchListResponse) GetIsSignedOk() (*bool, bool)`

GetIsSignedOk returns a tuple with the IsSigned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSigned

`func (o *CollectionVersionSearchListResponse) SetIsSigned(v bool)`

SetIsSigned sets IsSigned field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


