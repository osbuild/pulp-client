# CollectionVersionListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Version** | Pointer to **string** |  | [optional] [readonly] 
**Href** | Pointer to **string** | Get href. | [optional] [readonly] 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 
**RequiresAnsible** | Pointer to **NullableString** |  | [optional] 
**Marks** | Pointer to **[]string** | Get a list of mark values filtering only those in the current repo. | [optional] [readonly] 

## Methods

### NewCollectionVersionListResponse

`func NewCollectionVersionListResponse(createdAt time.Time, updatedAt time.Time, ) *CollectionVersionListResponse`

NewCollectionVersionListResponse instantiates a new CollectionVersionListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCollectionVersionListResponseWithDefaults

`func NewCollectionVersionListResponseWithDefaults() *CollectionVersionListResponse`

NewCollectionVersionListResponseWithDefaults instantiates a new CollectionVersionListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersion

`func (o *CollectionVersionListResponse) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *CollectionVersionListResponse) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *CollectionVersionListResponse) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *CollectionVersionListResponse) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetHref

`func (o *CollectionVersionListResponse) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *CollectionVersionListResponse) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *CollectionVersionListResponse) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *CollectionVersionListResponse) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetCreatedAt

`func (o *CollectionVersionListResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CollectionVersionListResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CollectionVersionListResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *CollectionVersionListResponse) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *CollectionVersionListResponse) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *CollectionVersionListResponse) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetRequiresAnsible

`func (o *CollectionVersionListResponse) GetRequiresAnsible() string`

GetRequiresAnsible returns the RequiresAnsible field if non-nil, zero value otherwise.

### GetRequiresAnsibleOk

`func (o *CollectionVersionListResponse) GetRequiresAnsibleOk() (*string, bool)`

GetRequiresAnsibleOk returns a tuple with the RequiresAnsible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiresAnsible

`func (o *CollectionVersionListResponse) SetRequiresAnsible(v string)`

SetRequiresAnsible sets RequiresAnsible field to given value.

### HasRequiresAnsible

`func (o *CollectionVersionListResponse) HasRequiresAnsible() bool`

HasRequiresAnsible returns a boolean if a field has been set.

### SetRequiresAnsibleNil

`func (o *CollectionVersionListResponse) SetRequiresAnsibleNil(b bool)`

 SetRequiresAnsibleNil sets the value for RequiresAnsible to be an explicit nil

### UnsetRequiresAnsible
`func (o *CollectionVersionListResponse) UnsetRequiresAnsible()`

UnsetRequiresAnsible ensures that no value is present for RequiresAnsible, not even an explicit nil
### GetMarks

`func (o *CollectionVersionListResponse) GetMarks() []string`

GetMarks returns the Marks field if non-nil, zero value otherwise.

### GetMarksOk

`func (o *CollectionVersionListResponse) GetMarksOk() (*[]string, bool)`

GetMarksOk returns a tuple with the Marks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarks

`func (o *CollectionVersionListResponse) SetMarks(v []string)`

SetMarks sets Marks field to given value.

### HasMarks

`func (o *CollectionVersionListResponse) HasMarks() bool`

HasMarks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


