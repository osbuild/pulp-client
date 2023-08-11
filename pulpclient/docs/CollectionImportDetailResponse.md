# CollectionImportDetailResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**State** | **string** |  | 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 
**StartedAt** | **time.Time** |  | 
**FinishedAt** | Pointer to **time.Time** |  | [optional] 
**Error** | Pointer to **map[string]interface{}** |  | [optional] 
**Messages** | **map[string]interface{}** |  | 

## Methods

### NewCollectionImportDetailResponse

`func NewCollectionImportDetailResponse(id string, state string, createdAt time.Time, updatedAt time.Time, startedAt time.Time, messages map[string]interface{}, ) *CollectionImportDetailResponse`

NewCollectionImportDetailResponse instantiates a new CollectionImportDetailResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCollectionImportDetailResponseWithDefaults

`func NewCollectionImportDetailResponseWithDefaults() *CollectionImportDetailResponse`

NewCollectionImportDetailResponseWithDefaults instantiates a new CollectionImportDetailResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CollectionImportDetailResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CollectionImportDetailResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CollectionImportDetailResponse) SetId(v string)`

SetId sets Id field to given value.


### GetState

`func (o *CollectionImportDetailResponse) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *CollectionImportDetailResponse) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *CollectionImportDetailResponse) SetState(v string)`

SetState sets State field to given value.


### GetCreatedAt

`func (o *CollectionImportDetailResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CollectionImportDetailResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CollectionImportDetailResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *CollectionImportDetailResponse) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *CollectionImportDetailResponse) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *CollectionImportDetailResponse) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetStartedAt

`func (o *CollectionImportDetailResponse) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *CollectionImportDetailResponse) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *CollectionImportDetailResponse) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.


### GetFinishedAt

`func (o *CollectionImportDetailResponse) GetFinishedAt() time.Time`

GetFinishedAt returns the FinishedAt field if non-nil, zero value otherwise.

### GetFinishedAtOk

`func (o *CollectionImportDetailResponse) GetFinishedAtOk() (*time.Time, bool)`

GetFinishedAtOk returns a tuple with the FinishedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinishedAt

`func (o *CollectionImportDetailResponse) SetFinishedAt(v time.Time)`

SetFinishedAt sets FinishedAt field to given value.

### HasFinishedAt

`func (o *CollectionImportDetailResponse) HasFinishedAt() bool`

HasFinishedAt returns a boolean if a field has been set.

### GetError

`func (o *CollectionImportDetailResponse) GetError() map[string]interface{}`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *CollectionImportDetailResponse) GetErrorOk() (*map[string]interface{}, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *CollectionImportDetailResponse) SetError(v map[string]interface{})`

SetError sets Error field to given value.

### HasError

`func (o *CollectionImportDetailResponse) HasError() bool`

HasError returns a boolean if a field has been set.

### GetMessages

`func (o *CollectionImportDetailResponse) GetMessages() map[string]interface{}`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *CollectionImportDetailResponse) GetMessagesOk() (*map[string]interface{}, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *CollectionImportDetailResponse) SetMessages(v map[string]interface{})`

SetMessages sets Messages field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


