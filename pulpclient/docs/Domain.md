# Domain

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | A name for this domain. | 
**Description** | Pointer to **NullableString** | An optional description. | [optional] 
**StorageClass** | [**StorageClassEnum**](StorageClassEnum.md) |  | 
**StorageSettings** | **map[string]interface{}** | Settings for storage class. | 
**RedirectToObjectStorage** | Pointer to **bool** | Boolean to have the content app redirect to object storage. | [optional] [default to true]
**HideGuardedDistributions** | Pointer to **bool** | Boolean to hide distributions with a content guard in the content app. | [optional] [default to false]

## Methods

### NewDomain

`func NewDomain(name string, storageClass StorageClassEnum, storageSettings map[string]interface{}, ) *Domain`

NewDomain instantiates a new Domain object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDomainWithDefaults

`func NewDomainWithDefaults() *Domain`

NewDomainWithDefaults instantiates a new Domain object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *Domain) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Domain) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Domain) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *Domain) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Domain) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Domain) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Domain) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *Domain) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *Domain) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetStorageClass

`func (o *Domain) GetStorageClass() StorageClassEnum`

GetStorageClass returns the StorageClass field if non-nil, zero value otherwise.

### GetStorageClassOk

`func (o *Domain) GetStorageClassOk() (*StorageClassEnum, bool)`

GetStorageClassOk returns a tuple with the StorageClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageClass

`func (o *Domain) SetStorageClass(v StorageClassEnum)`

SetStorageClass sets StorageClass field to given value.


### GetStorageSettings

`func (o *Domain) GetStorageSettings() map[string]interface{}`

GetStorageSettings returns the StorageSettings field if non-nil, zero value otherwise.

### GetStorageSettingsOk

`func (o *Domain) GetStorageSettingsOk() (*map[string]interface{}, bool)`

GetStorageSettingsOk returns a tuple with the StorageSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageSettings

`func (o *Domain) SetStorageSettings(v map[string]interface{})`

SetStorageSettings sets StorageSettings field to given value.


### GetRedirectToObjectStorage

`func (o *Domain) GetRedirectToObjectStorage() bool`

GetRedirectToObjectStorage returns the RedirectToObjectStorage field if non-nil, zero value otherwise.

### GetRedirectToObjectStorageOk

`func (o *Domain) GetRedirectToObjectStorageOk() (*bool, bool)`

GetRedirectToObjectStorageOk returns a tuple with the RedirectToObjectStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectToObjectStorage

`func (o *Domain) SetRedirectToObjectStorage(v bool)`

SetRedirectToObjectStorage sets RedirectToObjectStorage field to given value.

### HasRedirectToObjectStorage

`func (o *Domain) HasRedirectToObjectStorage() bool`

HasRedirectToObjectStorage returns a boolean if a field has been set.

### GetHideGuardedDistributions

`func (o *Domain) GetHideGuardedDistributions() bool`

GetHideGuardedDistributions returns the HideGuardedDistributions field if non-nil, zero value otherwise.

### GetHideGuardedDistributionsOk

`func (o *Domain) GetHideGuardedDistributionsOk() (*bool, bool)`

GetHideGuardedDistributionsOk returns a tuple with the HideGuardedDistributions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHideGuardedDistributions

`func (o *Domain) SetHideGuardedDistributions(v bool)`

SetHideGuardedDistributions sets HideGuardedDistributions field to given value.

### HasHideGuardedDistributions

`func (o *Domain) HasHideGuardedDistributions() bool`

HasHideGuardedDistributions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


