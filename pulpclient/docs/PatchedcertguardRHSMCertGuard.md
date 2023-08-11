# PatchedcertguardRHSMCertGuard

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The unique name. | [optional] 
**Description** | Pointer to **NullableString** | An optional description. | [optional] 
**CaCertificate** | Pointer to **string** | A Certificate Authority (CA) certificate (or a bundle thereof) used to verify client-certificate authenticity. | [optional] 

## Methods

### NewPatchedcertguardRHSMCertGuard

`func NewPatchedcertguardRHSMCertGuard() *PatchedcertguardRHSMCertGuard`

NewPatchedcertguardRHSMCertGuard instantiates a new PatchedcertguardRHSMCertGuard object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedcertguardRHSMCertGuardWithDefaults

`func NewPatchedcertguardRHSMCertGuardWithDefaults() *PatchedcertguardRHSMCertGuard`

NewPatchedcertguardRHSMCertGuardWithDefaults instantiates a new PatchedcertguardRHSMCertGuard object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PatchedcertguardRHSMCertGuard) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedcertguardRHSMCertGuard) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedcertguardRHSMCertGuard) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedcertguardRHSMCertGuard) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *PatchedcertguardRHSMCertGuard) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchedcertguardRHSMCertGuard) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchedcertguardRHSMCertGuard) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchedcertguardRHSMCertGuard) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *PatchedcertguardRHSMCertGuard) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *PatchedcertguardRHSMCertGuard) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetCaCertificate

`func (o *PatchedcertguardRHSMCertGuard) GetCaCertificate() string`

GetCaCertificate returns the CaCertificate field if non-nil, zero value otherwise.

### GetCaCertificateOk

`func (o *PatchedcertguardRHSMCertGuard) GetCaCertificateOk() (*string, bool)`

GetCaCertificateOk returns a tuple with the CaCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaCertificate

`func (o *PatchedcertguardRHSMCertGuard) SetCaCertificate(v string)`

SetCaCertificate sets CaCertificate field to given value.

### HasCaCertificate

`func (o *PatchedcertguardRHSMCertGuard) HasCaCertificate() bool`

HasCaCertificate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


