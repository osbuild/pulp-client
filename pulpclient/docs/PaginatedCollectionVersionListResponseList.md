# PaginatedCollectionVersionListResponseList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**PaginatedCollectionResponseListMeta**](PaginatedCollectionResponseListMeta.md) |  | [optional] 
**Links** | Pointer to [**PaginatedCollectionResponseListLinks**](PaginatedCollectionResponseListLinks.md) |  | [optional] 
**Data** | Pointer to [**[]CollectionVersionListResponse**](CollectionVersionListResponse.md) |  | [optional] 

## Methods

### NewPaginatedCollectionVersionListResponseList

`func NewPaginatedCollectionVersionListResponseList() *PaginatedCollectionVersionListResponseList`

NewPaginatedCollectionVersionListResponseList instantiates a new PaginatedCollectionVersionListResponseList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginatedCollectionVersionListResponseListWithDefaults

`func NewPaginatedCollectionVersionListResponseListWithDefaults() *PaginatedCollectionVersionListResponseList`

NewPaginatedCollectionVersionListResponseListWithDefaults instantiates a new PaginatedCollectionVersionListResponseList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *PaginatedCollectionVersionListResponseList) GetMeta() PaginatedCollectionResponseListMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *PaginatedCollectionVersionListResponseList) GetMetaOk() (*PaginatedCollectionResponseListMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *PaginatedCollectionVersionListResponseList) SetMeta(v PaginatedCollectionResponseListMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *PaginatedCollectionVersionListResponseList) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetLinks

`func (o *PaginatedCollectionVersionListResponseList) GetLinks() PaginatedCollectionResponseListLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *PaginatedCollectionVersionListResponseList) GetLinksOk() (*PaginatedCollectionResponseListLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *PaginatedCollectionVersionListResponseList) SetLinks(v PaginatedCollectionResponseListLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *PaginatedCollectionVersionListResponseList) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetData

`func (o *PaginatedCollectionVersionListResponseList) GetData() []CollectionVersionListResponse`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *PaginatedCollectionVersionListResponseList) GetDataOk() (*[]CollectionVersionListResponse, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *PaginatedCollectionVersionListResponseList) SetData(v []CollectionVersionListResponse)`

SetData sets Data field to given value.

### HasData

`func (o *PaginatedCollectionVersionListResponseList) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


