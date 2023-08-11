# CollectionNamespaceResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**MetadataSha256** | **NullableString** |  | 

## Methods

### NewCollectionNamespaceResponse

`func NewCollectionNamespaceResponse(name string, metadataSha256 NullableString, ) *CollectionNamespaceResponse`

NewCollectionNamespaceResponse instantiates a new CollectionNamespaceResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCollectionNamespaceResponseWithDefaults

`func NewCollectionNamespaceResponseWithDefaults() *CollectionNamespaceResponse`

NewCollectionNamespaceResponseWithDefaults instantiates a new CollectionNamespaceResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CollectionNamespaceResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CollectionNamespaceResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CollectionNamespaceResponse) SetName(v string)`

SetName sets Name field to given value.


### GetMetadataSha256

`func (o *CollectionNamespaceResponse) GetMetadataSha256() string`

GetMetadataSha256 returns the MetadataSha256 field if non-nil, zero value otherwise.

### GetMetadataSha256Ok

`func (o *CollectionNamespaceResponse) GetMetadataSha256Ok() (*string, bool)`

GetMetadataSha256Ok returns a tuple with the MetadataSha256 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataSha256

`func (o *CollectionNamespaceResponse) SetMetadataSha256(v string)`

SetMetadataSha256 sets MetadataSha256 field to given value.


### SetMetadataSha256Nil

`func (o *CollectionNamespaceResponse) SetMetadataSha256Nil(b bool)`

 SetMetadataSha256Nil sets the value for MetadataSha256 to be an explicit nil

### UnsetMetadataSha256
`func (o *CollectionNamespaceResponse) UnsetMetadataSha256()`

UnsetMetadataSha256 ensures that no value is present for MetadataSha256, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


