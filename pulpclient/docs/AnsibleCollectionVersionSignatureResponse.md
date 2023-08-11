# AnsibleCollectionVersionSignatureResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PulpHref** | Pointer to **string** |  | [optional] [readonly] 
**PulpCreated** | Pointer to **time.Time** | Timestamp of creation. | [optional] [readonly] 
**SignedCollection** | **string** | The content this signature is pointing to. | 
**PubkeyFingerprint** | Pointer to **string** | The fingerprint of the public key. | [optional] [readonly] 
**SigningService** | Pointer to **NullableString** | The signing service used to create the signature. | [optional] [readonly] 

## Methods

### NewAnsibleCollectionVersionSignatureResponse

`func NewAnsibleCollectionVersionSignatureResponse(signedCollection string, ) *AnsibleCollectionVersionSignatureResponse`

NewAnsibleCollectionVersionSignatureResponse instantiates a new AnsibleCollectionVersionSignatureResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnsibleCollectionVersionSignatureResponseWithDefaults

`func NewAnsibleCollectionVersionSignatureResponseWithDefaults() *AnsibleCollectionVersionSignatureResponse`

NewAnsibleCollectionVersionSignatureResponseWithDefaults instantiates a new AnsibleCollectionVersionSignatureResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPulpHref

`func (o *AnsibleCollectionVersionSignatureResponse) GetPulpHref() string`

GetPulpHref returns the PulpHref field if non-nil, zero value otherwise.

### GetPulpHrefOk

`func (o *AnsibleCollectionVersionSignatureResponse) GetPulpHrefOk() (*string, bool)`

GetPulpHrefOk returns a tuple with the PulpHref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpHref

`func (o *AnsibleCollectionVersionSignatureResponse) SetPulpHref(v string)`

SetPulpHref sets PulpHref field to given value.

### HasPulpHref

`func (o *AnsibleCollectionVersionSignatureResponse) HasPulpHref() bool`

HasPulpHref returns a boolean if a field has been set.

### GetPulpCreated

`func (o *AnsibleCollectionVersionSignatureResponse) GetPulpCreated() time.Time`

GetPulpCreated returns the PulpCreated field if non-nil, zero value otherwise.

### GetPulpCreatedOk

`func (o *AnsibleCollectionVersionSignatureResponse) GetPulpCreatedOk() (*time.Time, bool)`

GetPulpCreatedOk returns a tuple with the PulpCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpCreated

`func (o *AnsibleCollectionVersionSignatureResponse) SetPulpCreated(v time.Time)`

SetPulpCreated sets PulpCreated field to given value.

### HasPulpCreated

`func (o *AnsibleCollectionVersionSignatureResponse) HasPulpCreated() bool`

HasPulpCreated returns a boolean if a field has been set.

### GetSignedCollection

`func (o *AnsibleCollectionVersionSignatureResponse) GetSignedCollection() string`

GetSignedCollection returns the SignedCollection field if non-nil, zero value otherwise.

### GetSignedCollectionOk

`func (o *AnsibleCollectionVersionSignatureResponse) GetSignedCollectionOk() (*string, bool)`

GetSignedCollectionOk returns a tuple with the SignedCollection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignedCollection

`func (o *AnsibleCollectionVersionSignatureResponse) SetSignedCollection(v string)`

SetSignedCollection sets SignedCollection field to given value.


### GetPubkeyFingerprint

`func (o *AnsibleCollectionVersionSignatureResponse) GetPubkeyFingerprint() string`

GetPubkeyFingerprint returns the PubkeyFingerprint field if non-nil, zero value otherwise.

### GetPubkeyFingerprintOk

`func (o *AnsibleCollectionVersionSignatureResponse) GetPubkeyFingerprintOk() (*string, bool)`

GetPubkeyFingerprintOk returns a tuple with the PubkeyFingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPubkeyFingerprint

`func (o *AnsibleCollectionVersionSignatureResponse) SetPubkeyFingerprint(v string)`

SetPubkeyFingerprint sets PubkeyFingerprint field to given value.

### HasPubkeyFingerprint

`func (o *AnsibleCollectionVersionSignatureResponse) HasPubkeyFingerprint() bool`

HasPubkeyFingerprint returns a boolean if a field has been set.

### GetSigningService

`func (o *AnsibleCollectionVersionSignatureResponse) GetSigningService() string`

GetSigningService returns the SigningService field if non-nil, zero value otherwise.

### GetSigningServiceOk

`func (o *AnsibleCollectionVersionSignatureResponse) GetSigningServiceOk() (*string, bool)`

GetSigningServiceOk returns a tuple with the SigningService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigningService

`func (o *AnsibleCollectionVersionSignatureResponse) SetSigningService(v string)`

SetSigningService sets SigningService field to given value.

### HasSigningService

`func (o *AnsibleCollectionVersionSignatureResponse) HasSigningService() bool`

HasSigningService returns a boolean if a field has been set.

### SetSigningServiceNil

`func (o *AnsibleCollectionVersionSignatureResponse) SetSigningServiceNil(b bool)`

 SetSigningServiceNil sets the value for SigningService to be an explicit nil

### UnsetSigningService
`func (o *AnsibleCollectionVersionSignatureResponse) UnsetSigningService()`

UnsetSigningService ensures that no value is present for SigningService, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


