# CertguardRHSMCertGuardResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PulpHref** | Pointer to **string** |  | [optional] [readonly] 
**PulpCreated** | Pointer to **time.Time** | Timestamp of creation. | [optional] [readonly] 
**Name** | **string** | The unique name. | 
**Description** | Pointer to **NullableString** | An optional description. | [optional] 
**CaCertificate** | **string** | A Certificate Authority (CA) certificate (or a bundle thereof) used to verify client-certificate authenticity. | 

## Methods

### NewCertguardRHSMCertGuardResponse

`func NewCertguardRHSMCertGuardResponse(name string, caCertificate string, ) *CertguardRHSMCertGuardResponse`

NewCertguardRHSMCertGuardResponse instantiates a new CertguardRHSMCertGuardResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCertguardRHSMCertGuardResponseWithDefaults

`func NewCertguardRHSMCertGuardResponseWithDefaults() *CertguardRHSMCertGuardResponse`

NewCertguardRHSMCertGuardResponseWithDefaults instantiates a new CertguardRHSMCertGuardResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPulpHref

`func (o *CertguardRHSMCertGuardResponse) GetPulpHref() string`

GetPulpHref returns the PulpHref field if non-nil, zero value otherwise.

### GetPulpHrefOk

`func (o *CertguardRHSMCertGuardResponse) GetPulpHrefOk() (*string, bool)`

GetPulpHrefOk returns a tuple with the PulpHref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpHref

`func (o *CertguardRHSMCertGuardResponse) SetPulpHref(v string)`

SetPulpHref sets PulpHref field to given value.

### HasPulpHref

`func (o *CertguardRHSMCertGuardResponse) HasPulpHref() bool`

HasPulpHref returns a boolean if a field has been set.

### GetPulpCreated

`func (o *CertguardRHSMCertGuardResponse) GetPulpCreated() time.Time`

GetPulpCreated returns the PulpCreated field if non-nil, zero value otherwise.

### GetPulpCreatedOk

`func (o *CertguardRHSMCertGuardResponse) GetPulpCreatedOk() (*time.Time, bool)`

GetPulpCreatedOk returns a tuple with the PulpCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpCreated

`func (o *CertguardRHSMCertGuardResponse) SetPulpCreated(v time.Time)`

SetPulpCreated sets PulpCreated field to given value.

### HasPulpCreated

`func (o *CertguardRHSMCertGuardResponse) HasPulpCreated() bool`

HasPulpCreated returns a boolean if a field has been set.

### GetName

`func (o *CertguardRHSMCertGuardResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CertguardRHSMCertGuardResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CertguardRHSMCertGuardResponse) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *CertguardRHSMCertGuardResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CertguardRHSMCertGuardResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CertguardRHSMCertGuardResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CertguardRHSMCertGuardResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *CertguardRHSMCertGuardResponse) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *CertguardRHSMCertGuardResponse) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetCaCertificate

`func (o *CertguardRHSMCertGuardResponse) GetCaCertificate() string`

GetCaCertificate returns the CaCertificate field if non-nil, zero value otherwise.

### GetCaCertificateOk

`func (o *CertguardRHSMCertGuardResponse) GetCaCertificateOk() (*string, bool)`

GetCaCertificateOk returns a tuple with the CaCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaCertificate

`func (o *CertguardRHSMCertGuardResponse) SetCaCertificate(v string)`

SetCaCertificate sets CaCertificate field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


