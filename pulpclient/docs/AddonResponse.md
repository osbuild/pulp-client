# AddonResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddonId** | **string** | Addon id. | 
**Uid** | **string** | Addon uid. | 
**Name** | **string** | Addon name. | 
**Type** | **string** | Addon type. | 
**Packages** | **string** | Relative path to directory with binary RPMs. | 

## Methods

### NewAddonResponse

`func NewAddonResponse(addonId string, uid string, name string, type_ string, packages string, ) *AddonResponse`

NewAddonResponse instantiates a new AddonResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddonResponseWithDefaults

`func NewAddonResponseWithDefaults() *AddonResponse`

NewAddonResponseWithDefaults instantiates a new AddonResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddonId

`func (o *AddonResponse) GetAddonId() string`

GetAddonId returns the AddonId field if non-nil, zero value otherwise.

### GetAddonIdOk

`func (o *AddonResponse) GetAddonIdOk() (*string, bool)`

GetAddonIdOk returns a tuple with the AddonId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddonId

`func (o *AddonResponse) SetAddonId(v string)`

SetAddonId sets AddonId field to given value.


### GetUid

`func (o *AddonResponse) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *AddonResponse) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *AddonResponse) SetUid(v string)`

SetUid sets Uid field to given value.


### GetName

`func (o *AddonResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AddonResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AddonResponse) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *AddonResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AddonResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AddonResponse) SetType(v string)`

SetType sets Type field to given value.


### GetPackages

`func (o *AddonResponse) GetPackages() string`

GetPackages returns the Packages field if non-nil, zero value otherwise.

### GetPackagesOk

`func (o *AddonResponse) GetPackagesOk() (*string, bool)`

GetPackagesOk returns a tuple with the Packages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackages

`func (o *AddonResponse) SetPackages(v string)`

SetPackages sets Packages field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


