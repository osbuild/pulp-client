# RpmPackageCategoryResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PulpHref** | Pointer to **string** |  | [optional] [readonly] 
**PulpCreated** | Pointer to **time.Time** | Timestamp of creation. | [optional] [readonly] 
**Id** | **string** | Category id. | 
**Name** | **string** | Category name. | 
**Description** | **string** | Category description. | 
**DisplayOrder** | **NullableInt64** | Category display order. | 
**GroupIds** | **map[string]interface{}** | Category group list. | 
**DescByLang** | **map[string]interface{}** | Category description by language. | 
**NameByLang** | **map[string]interface{}** | Category name by language. | 
**Digest** | **string** | Category digest. | 

## Methods

### NewRpmPackageCategoryResponse

`func NewRpmPackageCategoryResponse(id string, name string, description string, displayOrder NullableInt64, groupIds map[string]interface{}, descByLang map[string]interface{}, nameByLang map[string]interface{}, digest string, ) *RpmPackageCategoryResponse`

NewRpmPackageCategoryResponse instantiates a new RpmPackageCategoryResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRpmPackageCategoryResponseWithDefaults

`func NewRpmPackageCategoryResponseWithDefaults() *RpmPackageCategoryResponse`

NewRpmPackageCategoryResponseWithDefaults instantiates a new RpmPackageCategoryResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPulpHref

`func (o *RpmPackageCategoryResponse) GetPulpHref() string`

GetPulpHref returns the PulpHref field if non-nil, zero value otherwise.

### GetPulpHrefOk

`func (o *RpmPackageCategoryResponse) GetPulpHrefOk() (*string, bool)`

GetPulpHrefOk returns a tuple with the PulpHref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpHref

`func (o *RpmPackageCategoryResponse) SetPulpHref(v string)`

SetPulpHref sets PulpHref field to given value.

### HasPulpHref

`func (o *RpmPackageCategoryResponse) HasPulpHref() bool`

HasPulpHref returns a boolean if a field has been set.

### GetPulpCreated

`func (o *RpmPackageCategoryResponse) GetPulpCreated() time.Time`

GetPulpCreated returns the PulpCreated field if non-nil, zero value otherwise.

### GetPulpCreatedOk

`func (o *RpmPackageCategoryResponse) GetPulpCreatedOk() (*time.Time, bool)`

GetPulpCreatedOk returns a tuple with the PulpCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpCreated

`func (o *RpmPackageCategoryResponse) SetPulpCreated(v time.Time)`

SetPulpCreated sets PulpCreated field to given value.

### HasPulpCreated

`func (o *RpmPackageCategoryResponse) HasPulpCreated() bool`

HasPulpCreated returns a boolean if a field has been set.

### GetId

`func (o *RpmPackageCategoryResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RpmPackageCategoryResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RpmPackageCategoryResponse) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *RpmPackageCategoryResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RpmPackageCategoryResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RpmPackageCategoryResponse) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *RpmPackageCategoryResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RpmPackageCategoryResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RpmPackageCategoryResponse) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetDisplayOrder

`func (o *RpmPackageCategoryResponse) GetDisplayOrder() int64`

GetDisplayOrder returns the DisplayOrder field if non-nil, zero value otherwise.

### GetDisplayOrderOk

`func (o *RpmPackageCategoryResponse) GetDisplayOrderOk() (*int64, bool)`

GetDisplayOrderOk returns a tuple with the DisplayOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayOrder

`func (o *RpmPackageCategoryResponse) SetDisplayOrder(v int64)`

SetDisplayOrder sets DisplayOrder field to given value.


### SetDisplayOrderNil

`func (o *RpmPackageCategoryResponse) SetDisplayOrderNil(b bool)`

 SetDisplayOrderNil sets the value for DisplayOrder to be an explicit nil

### UnsetDisplayOrder
`func (o *RpmPackageCategoryResponse) UnsetDisplayOrder()`

UnsetDisplayOrder ensures that no value is present for DisplayOrder, not even an explicit nil
### GetGroupIds

`func (o *RpmPackageCategoryResponse) GetGroupIds() map[string]interface{}`

GetGroupIds returns the GroupIds field if non-nil, zero value otherwise.

### GetGroupIdsOk

`func (o *RpmPackageCategoryResponse) GetGroupIdsOk() (*map[string]interface{}, bool)`

GetGroupIdsOk returns a tuple with the GroupIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupIds

`func (o *RpmPackageCategoryResponse) SetGroupIds(v map[string]interface{})`

SetGroupIds sets GroupIds field to given value.


### SetGroupIdsNil

`func (o *RpmPackageCategoryResponse) SetGroupIdsNil(b bool)`

 SetGroupIdsNil sets the value for GroupIds to be an explicit nil

### UnsetGroupIds
`func (o *RpmPackageCategoryResponse) UnsetGroupIds()`

UnsetGroupIds ensures that no value is present for GroupIds, not even an explicit nil
### GetDescByLang

`func (o *RpmPackageCategoryResponse) GetDescByLang() map[string]interface{}`

GetDescByLang returns the DescByLang field if non-nil, zero value otherwise.

### GetDescByLangOk

`func (o *RpmPackageCategoryResponse) GetDescByLangOk() (*map[string]interface{}, bool)`

GetDescByLangOk returns a tuple with the DescByLang field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescByLang

`func (o *RpmPackageCategoryResponse) SetDescByLang(v map[string]interface{})`

SetDescByLang sets DescByLang field to given value.


### SetDescByLangNil

`func (o *RpmPackageCategoryResponse) SetDescByLangNil(b bool)`

 SetDescByLangNil sets the value for DescByLang to be an explicit nil

### UnsetDescByLang
`func (o *RpmPackageCategoryResponse) UnsetDescByLang()`

UnsetDescByLang ensures that no value is present for DescByLang, not even an explicit nil
### GetNameByLang

`func (o *RpmPackageCategoryResponse) GetNameByLang() map[string]interface{}`

GetNameByLang returns the NameByLang field if non-nil, zero value otherwise.

### GetNameByLangOk

`func (o *RpmPackageCategoryResponse) GetNameByLangOk() (*map[string]interface{}, bool)`

GetNameByLangOk returns a tuple with the NameByLang field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameByLang

`func (o *RpmPackageCategoryResponse) SetNameByLang(v map[string]interface{})`

SetNameByLang sets NameByLang field to given value.


### SetNameByLangNil

`func (o *RpmPackageCategoryResponse) SetNameByLangNil(b bool)`

 SetNameByLangNil sets the value for NameByLang to be an explicit nil

### UnsetNameByLang
`func (o *RpmPackageCategoryResponse) UnsetNameByLang()`

UnsetNameByLang ensures that no value is present for NameByLang, not even an explicit nil
### GetDigest

`func (o *RpmPackageCategoryResponse) GetDigest() string`

GetDigest returns the Digest field if non-nil, zero value otherwise.

### GetDigestOk

`func (o *RpmPackageCategoryResponse) GetDigestOk() (*string, bool)`

GetDigestOk returns a tuple with the Digest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigest

`func (o *RpmPackageCategoryResponse) SetDigest(v string)`

SetDigest sets Digest field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


