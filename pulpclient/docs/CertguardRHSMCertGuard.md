# CertguardRHSMCertGuard

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The unique name. | 
**Description** | Pointer to **NullableString** | An optional description. | [optional] 
**CaCertificate** | **string** | A Certificate Authority (CA) certificate (or a bundle thereof) used to verify client-certificate authenticity. | 

## Methods

### NewCertguardRHSMCertGuard

`func NewCertguardRHSMCertGuard(name string, caCertificate string, ) *CertguardRHSMCertGuard`

NewCertguardRHSMCertGuard instantiates a new CertguardRHSMCertGuard object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCertguardRHSMCertGuardWithDefaults

`func NewCertguardRHSMCertGuardWithDefaults() *CertguardRHSMCertGuard`

NewCertguardRHSMCertGuardWithDefaults instantiates a new CertguardRHSMCertGuard object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CertguardRHSMCertGuard) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CertguardRHSMCertGuard) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CertguardRHSMCertGuard) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *CertguardRHSMCertGuard) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CertguardRHSMCertGuard) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CertguardRHSMCertGuard) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CertguardRHSMCertGuard) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *CertguardRHSMCertGuard) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *CertguardRHSMCertGuard) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetCaCertificate

`func (o *CertguardRHSMCertGuard) GetCaCertificate() string`

GetCaCertificate returns the CaCertificate field if non-nil, zero value otherwise.

### GetCaCertificateOk

`func (o *CertguardRHSMCertGuard) GetCaCertificateOk() (*string, bool)`

GetCaCertificateOk returns a tuple with the CaCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaCertificate

`func (o *CertguardRHSMCertGuard) SetCaCertificate(v string)`

SetCaCertificate sets CaCertificate field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


