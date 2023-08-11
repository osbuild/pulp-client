# PatchedcertguardX509CertGuard

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The unique name. | [optional] 
**Description** | Pointer to **NullableString** | An optional description. | [optional] 
**CaCertificate** | Pointer to **string** | A Certificate Authority (CA) certificate (or a bundle thereof) used to verify client-certificate authenticity. | [optional] 

## Methods

### NewPatchedcertguardX509CertGuard

`func NewPatchedcertguardX509CertGuard() *PatchedcertguardX509CertGuard`

NewPatchedcertguardX509CertGuard instantiates a new PatchedcertguardX509CertGuard object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedcertguardX509CertGuardWithDefaults

`func NewPatchedcertguardX509CertGuardWithDefaults() *PatchedcertguardX509CertGuard`

NewPatchedcertguardX509CertGuardWithDefaults instantiates a new PatchedcertguardX509CertGuard object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PatchedcertguardX509CertGuard) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedcertguardX509CertGuard) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedcertguardX509CertGuard) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedcertguardX509CertGuard) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *PatchedcertguardX509CertGuard) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchedcertguardX509CertGuard) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchedcertguardX509CertGuard) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchedcertguardX509CertGuard) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *PatchedcertguardX509CertGuard) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *PatchedcertguardX509CertGuard) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetCaCertificate

`func (o *PatchedcertguardX509CertGuard) GetCaCertificate() string`

GetCaCertificate returns the CaCertificate field if non-nil, zero value otherwise.

### GetCaCertificateOk

`func (o *PatchedcertguardX509CertGuard) GetCaCertificateOk() (*string, bool)`

GetCaCertificateOk returns a tuple with the CaCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaCertificate

`func (o *PatchedcertguardX509CertGuard) SetCaCertificate(v string)`

SetCaCertificate sets CaCertificate field to given value.

### HasCaCertificate

`func (o *PatchedcertguardX509CertGuard) HasCaCertificate() bool`

HasCaCertificate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


