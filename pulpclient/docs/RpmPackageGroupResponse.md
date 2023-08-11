# RpmPackageGroupResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PulpHref** | Pointer to **string** |  | [optional] [readonly] 
**PulpCreated** | Pointer to **time.Time** | Timestamp of creation. | [optional] [readonly] 
**Id** | **string** | PackageGroup id. | 
**Default** | Pointer to **bool** | PackageGroup default. | [optional] 
**UserVisible** | Pointer to **bool** | PackageGroup user visibility. | [optional] 
**DisplayOrder** | **NullableInt64** | PackageGroup display order. | 
**Name** | **string** | PackageGroup name. | 
**Description** | **string** | PackageGroup description. | 
**Packages** | **map[string]interface{}** | PackageGroup package list. | 
**BiarchOnly** | Pointer to **bool** | PackageGroup biarch only. | [optional] 
**DescByLang** | **map[string]interface{}** | PackageGroup description by language. | 
**NameByLang** | **map[string]interface{}** | PackageGroup name by language. | 
**Digest** | **string** | PackageGroup digest. | 

## Methods

### NewRpmPackageGroupResponse

`func NewRpmPackageGroupResponse(id string, displayOrder NullableInt64, name string, description string, packages map[string]interface{}, descByLang map[string]interface{}, nameByLang map[string]interface{}, digest string, ) *RpmPackageGroupResponse`

NewRpmPackageGroupResponse instantiates a new RpmPackageGroupResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRpmPackageGroupResponseWithDefaults

`func NewRpmPackageGroupResponseWithDefaults() *RpmPackageGroupResponse`

NewRpmPackageGroupResponseWithDefaults instantiates a new RpmPackageGroupResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPulpHref

`func (o *RpmPackageGroupResponse) GetPulpHref() string`

GetPulpHref returns the PulpHref field if non-nil, zero value otherwise.

### GetPulpHrefOk

`func (o *RpmPackageGroupResponse) GetPulpHrefOk() (*string, bool)`

GetPulpHrefOk returns a tuple with the PulpHref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpHref

`func (o *RpmPackageGroupResponse) SetPulpHref(v string)`

SetPulpHref sets PulpHref field to given value.

### HasPulpHref

`func (o *RpmPackageGroupResponse) HasPulpHref() bool`

HasPulpHref returns a boolean if a field has been set.

### GetPulpCreated

`func (o *RpmPackageGroupResponse) GetPulpCreated() time.Time`

GetPulpCreated returns the PulpCreated field if non-nil, zero value otherwise.

### GetPulpCreatedOk

`func (o *RpmPackageGroupResponse) GetPulpCreatedOk() (*time.Time, bool)`

GetPulpCreatedOk returns a tuple with the PulpCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpCreated

`func (o *RpmPackageGroupResponse) SetPulpCreated(v time.Time)`

SetPulpCreated sets PulpCreated field to given value.

### HasPulpCreated

`func (o *RpmPackageGroupResponse) HasPulpCreated() bool`

HasPulpCreated returns a boolean if a field has been set.

### GetId

`func (o *RpmPackageGroupResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RpmPackageGroupResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RpmPackageGroupResponse) SetId(v string)`

SetId sets Id field to given value.


### GetDefault

`func (o *RpmPackageGroupResponse) GetDefault() bool`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *RpmPackageGroupResponse) GetDefaultOk() (*bool, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *RpmPackageGroupResponse) SetDefault(v bool)`

SetDefault sets Default field to given value.

### HasDefault

`func (o *RpmPackageGroupResponse) HasDefault() bool`

HasDefault returns a boolean if a field has been set.

### GetUserVisible

`func (o *RpmPackageGroupResponse) GetUserVisible() bool`

GetUserVisible returns the UserVisible field if non-nil, zero value otherwise.

### GetUserVisibleOk

`func (o *RpmPackageGroupResponse) GetUserVisibleOk() (*bool, bool)`

GetUserVisibleOk returns a tuple with the UserVisible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserVisible

`func (o *RpmPackageGroupResponse) SetUserVisible(v bool)`

SetUserVisible sets UserVisible field to given value.

### HasUserVisible

`func (o *RpmPackageGroupResponse) HasUserVisible() bool`

HasUserVisible returns a boolean if a field has been set.

### GetDisplayOrder

`func (o *RpmPackageGroupResponse) GetDisplayOrder() int64`

GetDisplayOrder returns the DisplayOrder field if non-nil, zero value otherwise.

### GetDisplayOrderOk

`func (o *RpmPackageGroupResponse) GetDisplayOrderOk() (*int64, bool)`

GetDisplayOrderOk returns a tuple with the DisplayOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayOrder

`func (o *RpmPackageGroupResponse) SetDisplayOrder(v int64)`

SetDisplayOrder sets DisplayOrder field to given value.


### SetDisplayOrderNil

`func (o *RpmPackageGroupResponse) SetDisplayOrderNil(b bool)`

 SetDisplayOrderNil sets the value for DisplayOrder to be an explicit nil

### UnsetDisplayOrder
`func (o *RpmPackageGroupResponse) UnsetDisplayOrder()`

UnsetDisplayOrder ensures that no value is present for DisplayOrder, not even an explicit nil
### GetName

`func (o *RpmPackageGroupResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RpmPackageGroupResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RpmPackageGroupResponse) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *RpmPackageGroupResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RpmPackageGroupResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RpmPackageGroupResponse) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetPackages

`func (o *RpmPackageGroupResponse) GetPackages() map[string]interface{}`

GetPackages returns the Packages field if non-nil, zero value otherwise.

### GetPackagesOk

`func (o *RpmPackageGroupResponse) GetPackagesOk() (*map[string]interface{}, bool)`

GetPackagesOk returns a tuple with the Packages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackages

`func (o *RpmPackageGroupResponse) SetPackages(v map[string]interface{})`

SetPackages sets Packages field to given value.


### SetPackagesNil

`func (o *RpmPackageGroupResponse) SetPackagesNil(b bool)`

 SetPackagesNil sets the value for Packages to be an explicit nil

### UnsetPackages
`func (o *RpmPackageGroupResponse) UnsetPackages()`

UnsetPackages ensures that no value is present for Packages, not even an explicit nil
### GetBiarchOnly

`func (o *RpmPackageGroupResponse) GetBiarchOnly() bool`

GetBiarchOnly returns the BiarchOnly field if non-nil, zero value otherwise.

### GetBiarchOnlyOk

`func (o *RpmPackageGroupResponse) GetBiarchOnlyOk() (*bool, bool)`

GetBiarchOnlyOk returns a tuple with the BiarchOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBiarchOnly

`func (o *RpmPackageGroupResponse) SetBiarchOnly(v bool)`

SetBiarchOnly sets BiarchOnly field to given value.

### HasBiarchOnly

`func (o *RpmPackageGroupResponse) HasBiarchOnly() bool`

HasBiarchOnly returns a boolean if a field has been set.

### GetDescByLang

`func (o *RpmPackageGroupResponse) GetDescByLang() map[string]interface{}`

GetDescByLang returns the DescByLang field if non-nil, zero value otherwise.

### GetDescByLangOk

`func (o *RpmPackageGroupResponse) GetDescByLangOk() (*map[string]interface{}, bool)`

GetDescByLangOk returns a tuple with the DescByLang field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescByLang

`func (o *RpmPackageGroupResponse) SetDescByLang(v map[string]interface{})`

SetDescByLang sets DescByLang field to given value.


### SetDescByLangNil

`func (o *RpmPackageGroupResponse) SetDescByLangNil(b bool)`

 SetDescByLangNil sets the value for DescByLang to be an explicit nil

### UnsetDescByLang
`func (o *RpmPackageGroupResponse) UnsetDescByLang()`

UnsetDescByLang ensures that no value is present for DescByLang, not even an explicit nil
### GetNameByLang

`func (o *RpmPackageGroupResponse) GetNameByLang() map[string]interface{}`

GetNameByLang returns the NameByLang field if non-nil, zero value otherwise.

### GetNameByLangOk

`func (o *RpmPackageGroupResponse) GetNameByLangOk() (*map[string]interface{}, bool)`

GetNameByLangOk returns a tuple with the NameByLang field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameByLang

`func (o *RpmPackageGroupResponse) SetNameByLang(v map[string]interface{})`

SetNameByLang sets NameByLang field to given value.


### SetNameByLangNil

`func (o *RpmPackageGroupResponse) SetNameByLangNil(b bool)`

 SetNameByLangNil sets the value for NameByLang to be an explicit nil

### UnsetNameByLang
`func (o *RpmPackageGroupResponse) UnsetNameByLang()`

UnsetNameByLang ensures that no value is present for NameByLang, not even an explicit nil
### GetDigest

`func (o *RpmPackageGroupResponse) GetDigest() string`

GetDigest returns the Digest field if non-nil, zero value otherwise.

### GetDigestOk

`func (o *RpmPackageGroupResponse) GetDigestOk() (*string, bool)`

GetDigestOk returns a tuple with the Digest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigest

`func (o *RpmPackageGroupResponse) SetDigest(v string)`

SetDigest sets Digest field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


