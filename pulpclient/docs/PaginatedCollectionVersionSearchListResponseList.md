# PaginatedCollectionVersionSearchListResponseList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**PaginatedCollectionResponseListMeta**](PaginatedCollectionResponseListMeta.md) |  | [optional] 
**Links** | Pointer to [**PaginatedCollectionResponseListLinks**](PaginatedCollectionResponseListLinks.md) |  | [optional] 
**Data** | Pointer to [**[]CollectionVersionSearchListResponse**](CollectionVersionSearchListResponse.md) |  | [optional] 

## Methods

### NewPaginatedCollectionVersionSearchListResponseList

`func NewPaginatedCollectionVersionSearchListResponseList() *PaginatedCollectionVersionSearchListResponseList`

NewPaginatedCollectionVersionSearchListResponseList instantiates a new PaginatedCollectionVersionSearchListResponseList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginatedCollectionVersionSearchListResponseListWithDefaults

`func NewPaginatedCollectionVersionSearchListResponseListWithDefaults() *PaginatedCollectionVersionSearchListResponseList`

NewPaginatedCollectionVersionSearchListResponseListWithDefaults instantiates a new PaginatedCollectionVersionSearchListResponseList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *PaginatedCollectionVersionSearchListResponseList) GetMeta() PaginatedCollectionResponseListMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *PaginatedCollectionVersionSearchListResponseList) GetMetaOk() (*PaginatedCollectionResponseListMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *PaginatedCollectionVersionSearchListResponseList) SetMeta(v PaginatedCollectionResponseListMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *PaginatedCollectionVersionSearchListResponseList) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetLinks

`func (o *PaginatedCollectionVersionSearchListResponseList) GetLinks() PaginatedCollectionResponseListLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *PaginatedCollectionVersionSearchListResponseList) GetLinksOk() (*PaginatedCollectionResponseListLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *PaginatedCollectionVersionSearchListResponseList) SetLinks(v PaginatedCollectionResponseListLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *PaginatedCollectionVersionSearchListResponseList) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetData

`func (o *PaginatedCollectionVersionSearchListResponseList) GetData() []CollectionVersionSearchListResponse`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *PaginatedCollectionVersionSearchListResponseList) GetDataOk() (*[]CollectionVersionSearchListResponse, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *PaginatedCollectionVersionSearchListResponseList) SetData(v []CollectionVersionSearchListResponse)`

SetData sets Data field to given value.

### HasData

`func (o *PaginatedCollectionVersionSearchListResponseList) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


