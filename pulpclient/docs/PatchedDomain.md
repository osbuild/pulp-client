# PatchedDomain

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | A name for this domain. | [optional] 
**Description** | Pointer to **NullableString** | An optional description. | [optional] 
**StorageClass** | Pointer to [**StorageClassEnum**](StorageClassEnum.md) |  | [optional] 
**StorageSettings** | Pointer to **map[string]interface{}** | Settings for storage class. | [optional] 
**RedirectToObjectStorage** | Pointer to **bool** | Boolean to have the content app redirect to object storage. | [optional] [default to true]
**HideGuardedDistributions** | Pointer to **bool** | Boolean to hide distributions with a content guard in the content app. | [optional] [default to false]

## Methods

### NewPatchedDomain

`func NewPatchedDomain() *PatchedDomain`

NewPatchedDomain instantiates a new PatchedDomain object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedDomainWithDefaults

`func NewPatchedDomainWithDefaults() *PatchedDomain`

NewPatchedDomainWithDefaults instantiates a new PatchedDomain object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PatchedDomain) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedDomain) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedDomain) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedDomain) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *PatchedDomain) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchedDomain) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchedDomain) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchedDomain) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *PatchedDomain) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *PatchedDomain) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetStorageClass

`func (o *PatchedDomain) GetStorageClass() StorageClassEnum`

GetStorageClass returns the StorageClass field if non-nil, zero value otherwise.

### GetStorageClassOk

`func (o *PatchedDomain) GetStorageClassOk() (*StorageClassEnum, bool)`

GetStorageClassOk returns a tuple with the StorageClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageClass

`func (o *PatchedDomain) SetStorageClass(v StorageClassEnum)`

SetStorageClass sets StorageClass field to given value.

### HasStorageClass

`func (o *PatchedDomain) HasStorageClass() bool`

HasStorageClass returns a boolean if a field has been set.

### GetStorageSettings

`func (o *PatchedDomain) GetStorageSettings() map[string]interface{}`

GetStorageSettings returns the StorageSettings field if non-nil, zero value otherwise.

### GetStorageSettingsOk

`func (o *PatchedDomain) GetStorageSettingsOk() (*map[string]interface{}, bool)`

GetStorageSettingsOk returns a tuple with the StorageSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageSettings

`func (o *PatchedDomain) SetStorageSettings(v map[string]interface{})`

SetStorageSettings sets StorageSettings field to given value.

### HasStorageSettings

`func (o *PatchedDomain) HasStorageSettings() bool`

HasStorageSettings returns a boolean if a field has been set.

### GetRedirectToObjectStorage

`func (o *PatchedDomain) GetRedirectToObjectStorage() bool`

GetRedirectToObjectStorage returns the RedirectToObjectStorage field if non-nil, zero value otherwise.

### GetRedirectToObjectStorageOk

`func (o *PatchedDomain) GetRedirectToObjectStorageOk() (*bool, bool)`

GetRedirectToObjectStorageOk returns a tuple with the RedirectToObjectStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectToObjectStorage

`func (o *PatchedDomain) SetRedirectToObjectStorage(v bool)`

SetRedirectToObjectStorage sets RedirectToObjectStorage field to given value.

### HasRedirectToObjectStorage

`func (o *PatchedDomain) HasRedirectToObjectStorage() bool`

HasRedirectToObjectStorage returns a boolean if a field has been set.

### GetHideGuardedDistributions

`func (o *PatchedDomain) GetHideGuardedDistributions() bool`

GetHideGuardedDistributions returns the HideGuardedDistributions field if non-nil, zero value otherwise.

### GetHideGuardedDistributionsOk

`func (o *PatchedDomain) GetHideGuardedDistributionsOk() (*bool, bool)`

GetHideGuardedDistributionsOk returns a tuple with the HideGuardedDistributions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHideGuardedDistributions

`func (o *PatchedDomain) SetHideGuardedDistributions(v bool)`

SetHideGuardedDistributions sets HideGuardedDistributions field to given value.

### HasHideGuardedDistributions

`func (o *PatchedDomain) HasHideGuardedDistributions() bool`

HasHideGuardedDistributions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


