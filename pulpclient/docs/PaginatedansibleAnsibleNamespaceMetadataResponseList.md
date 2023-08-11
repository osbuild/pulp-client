# PaginatedansibleAnsibleNamespaceMetadataResponseList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int32** |  | [optional] 
**Next** | Pointer to **NullableString** |  | [optional] 
**Previous** | Pointer to **NullableString** |  | [optional] 
**Results** | Pointer to [**[]AnsibleAnsibleNamespaceMetadataResponse**](AnsibleAnsibleNamespaceMetadataResponse.md) |  | [optional] 

## Methods

### NewPaginatedansibleAnsibleNamespaceMetadataResponseList

`func NewPaginatedansibleAnsibleNamespaceMetadataResponseList() *PaginatedansibleAnsibleNamespaceMetadataResponseList`

NewPaginatedansibleAnsibleNamespaceMetadataResponseList instantiates a new PaginatedansibleAnsibleNamespaceMetadataResponseList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginatedansibleAnsibleNamespaceMetadataResponseListWithDefaults

`func NewPaginatedansibleAnsibleNamespaceMetadataResponseListWithDefaults() *PaginatedansibleAnsibleNamespaceMetadataResponseList`

NewPaginatedansibleAnsibleNamespaceMetadataResponseListWithDefaults instantiates a new PaginatedansibleAnsibleNamespaceMetadataResponseList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *PaginatedansibleAnsibleNamespaceMetadataResponseList) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *PaginatedansibleAnsibleNamespaceMetadataResponseList) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *PaginatedansibleAnsibleNamespaceMetadataResponseList) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *PaginatedansibleAnsibleNamespaceMetadataResponseList) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetNext

`func (o *PaginatedansibleAnsibleNamespaceMetadataResponseList) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *PaginatedansibleAnsibleNamespaceMetadataResponseList) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *PaginatedansibleAnsibleNamespaceMetadataResponseList) SetNext(v string)`

SetNext sets Next field to given value.

### HasNext

`func (o *PaginatedansibleAnsibleNamespaceMetadataResponseList) HasNext() bool`

HasNext returns a boolean if a field has been set.

### SetNextNil

`func (o *PaginatedansibleAnsibleNamespaceMetadataResponseList) SetNextNil(b bool)`

 SetNextNil sets the value for Next to be an explicit nil

### UnsetNext
`func (o *PaginatedansibleAnsibleNamespaceMetadataResponseList) UnsetNext()`

UnsetNext ensures that no value is present for Next, not even an explicit nil
### GetPrevious

`func (o *PaginatedansibleAnsibleNamespaceMetadataResponseList) GetPrevious() string`

GetPrevious returns the Previous field if non-nil, zero value otherwise.

### GetPreviousOk

`func (o *PaginatedansibleAnsibleNamespaceMetadataResponseList) GetPreviousOk() (*string, bool)`

GetPreviousOk returns a tuple with the Previous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevious

`func (o *PaginatedansibleAnsibleNamespaceMetadataResponseList) SetPrevious(v string)`

SetPrevious sets Previous field to given value.

### HasPrevious

`func (o *PaginatedansibleAnsibleNamespaceMetadataResponseList) HasPrevious() bool`

HasPrevious returns a boolean if a field has been set.

### SetPreviousNil

`func (o *PaginatedansibleAnsibleNamespaceMetadataResponseList) SetPreviousNil(b bool)`

 SetPreviousNil sets the value for Previous to be an explicit nil

### UnsetPrevious
`func (o *PaginatedansibleAnsibleNamespaceMetadataResponseList) UnsetPrevious()`

UnsetPrevious ensures that no value is present for Previous, not even an explicit nil
### GetResults

`func (o *PaginatedansibleAnsibleNamespaceMetadataResponseList) GetResults() []AnsibleAnsibleNamespaceMetadataResponse`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *PaginatedansibleAnsibleNamespaceMetadataResponseList) GetResultsOk() (*[]AnsibleAnsibleNamespaceMetadataResponse, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *PaginatedansibleAnsibleNamespaceMetadataResponseList) SetResults(v []AnsibleAnsibleNamespaceMetadataResponse)`

SetResults sets Results field to given value.

### HasResults

`func (o *PaginatedansibleAnsibleNamespaceMetadataResponseList) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


