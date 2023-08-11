# CertguardX509CertGuard

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The unique name. | 
**Description** | Pointer to **NullableString** | An optional description. | [optional] 
**CaCertificate** | **string** | A Certificate Authority (CA) certificate (or a bundle thereof) used to verify client-certificate authenticity. | 

## Methods

### NewCertguardX509CertGuard

`func NewCertguardX509CertGuard(name string, caCertificate string, ) *CertguardX509CertGuard`

NewCertguardX509CertGuard instantiates a new CertguardX509CertGuard object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCertguardX509CertGuardWithDefaults

`func NewCertguardX509CertGuardWithDefaults() *CertguardX509CertGuard`

NewCertguardX509CertGuardWithDefaults instantiates a new CertguardX509CertGuard object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CertguardX509CertGuard) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CertguardX509CertGuard) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CertguardX509CertGuard) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *CertguardX509CertGuard) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CertguardX509CertGuard) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CertguardX509CertGuard) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CertguardX509CertGuard) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *CertguardX509CertGuard) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *CertguardX509CertGuard) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetCaCertificate

`func (o *CertguardX509CertGuard) GetCaCertificate() string`

GetCaCertificate returns the CaCertificate field if non-nil, zero value otherwise.

### GetCaCertificateOk

`func (o *CertguardX509CertGuard) GetCaCertificateOk() (*string, bool)`

GetCaCertificateOk returns a tuple with the CaCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaCertificate

`func (o *CertguardX509CertGuard) SetCaCertificate(v string)`

SetCaCertificate sets CaCertificate field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


