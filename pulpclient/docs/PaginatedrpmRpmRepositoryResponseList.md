# PaginatedrpmRpmRepositoryResponseList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int32** |  | [optional] 
**Next** | Pointer to **NullableString** |  | [optional] 
**Previous** | Pointer to **NullableString** |  | [optional] 
**Results** | Pointer to [**[]RpmRpmRepositoryResponse**](RpmRpmRepositoryResponse.md) |  | [optional] 

## Methods

### NewPaginatedrpmRpmRepositoryResponseList

`func NewPaginatedrpmRpmRepositoryResponseList() *PaginatedrpmRpmRepositoryResponseList`

NewPaginatedrpmRpmRepositoryResponseList instantiates a new PaginatedrpmRpmRepositoryResponseList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginatedrpmRpmRepositoryResponseListWithDefaults

`func NewPaginatedrpmRpmRepositoryResponseListWithDefaults() *PaginatedrpmRpmRepositoryResponseList`

NewPaginatedrpmRpmRepositoryResponseListWithDefaults instantiates a new PaginatedrpmRpmRepositoryResponseList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *PaginatedrpmRpmRepositoryResponseList) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *PaginatedrpmRpmRepositoryResponseList) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *PaginatedrpmRpmRepositoryResponseList) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *PaginatedrpmRpmRepositoryResponseList) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetNext

`func (o *PaginatedrpmRpmRepositoryResponseList) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *PaginatedrpmRpmRepositoryResponseList) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *PaginatedrpmRpmRepositoryResponseList) SetNext(v string)`

SetNext sets Next field to given value.

### HasNext

`func (o *PaginatedrpmRpmRepositoryResponseList) HasNext() bool`

HasNext returns a boolean if a field has been set.

### SetNextNil

`func (o *PaginatedrpmRpmRepositoryResponseList) SetNextNil(b bool)`

 SetNextNil sets the value for Next to be an explicit nil

### UnsetNext
`func (o *PaginatedrpmRpmRepositoryResponseList) UnsetNext()`

UnsetNext ensures that no value is present for Next, not even an explicit nil
### GetPrevious

`func (o *PaginatedrpmRpmRepositoryResponseList) GetPrevious() string`

GetPrevious returns the Previous field if non-nil, zero value otherwise.

### GetPreviousOk

`func (o *PaginatedrpmRpmRepositoryResponseList) GetPreviousOk() (*string, bool)`

GetPreviousOk returns a tuple with the Previous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevious

`func (o *PaginatedrpmRpmRepositoryResponseList) SetPrevious(v string)`

SetPrevious sets Previous field to given value.

### HasPrevious

`func (o *PaginatedrpmRpmRepositoryResponseList) HasPrevious() bool`

HasPrevious returns a boolean if a field has been set.

### SetPreviousNil

`func (o *PaginatedrpmRpmRepositoryResponseList) SetPreviousNil(b bool)`

 SetPreviousNil sets the value for Previous to be an explicit nil

### UnsetPrevious
`func (o *PaginatedrpmRpmRepositoryResponseList) UnsetPrevious()`

UnsetPrevious ensures that no value is present for Previous, not even an explicit nil
### GetResults

`func (o *PaginatedrpmRpmRepositoryResponseList) GetResults() []RpmRpmRepositoryResponse`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *PaginatedrpmRpmRepositoryResponseList) GetResultsOk() (*[]RpmRpmRepositoryResponse, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *PaginatedrpmRpmRepositoryResponseList) SetResults(v []RpmRpmRepositoryResponse)`

SetResults sets Results field to given value.

### HasResults

`func (o *PaginatedrpmRpmRepositoryResponseList) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


