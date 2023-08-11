# CollectionVersionSignatureResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Signature** | Pointer to **string** |  | [optional] [readonly] 
**PubkeyFingerprint** | **string** |  | 
**SigningService** | Pointer to **NullableString** |  | [optional] [readonly] 
**PulpCreated** | Pointer to **time.Time** |  | [optional] [readonly] 

## Methods

### NewCollectionVersionSignatureResponse

`func NewCollectionVersionSignatureResponse(pubkeyFingerprint string, ) *CollectionVersionSignatureResponse`

NewCollectionVersionSignatureResponse instantiates a new CollectionVersionSignatureResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCollectionVersionSignatureResponseWithDefaults

`func NewCollectionVersionSignatureResponseWithDefaults() *CollectionVersionSignatureResponse`

NewCollectionVersionSignatureResponseWithDefaults instantiates a new CollectionVersionSignatureResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSignature

`func (o *CollectionVersionSignatureResponse) GetSignature() string`

GetSignature returns the Signature field if non-nil, zero value otherwise.

### GetSignatureOk

`func (o *CollectionVersionSignatureResponse) GetSignatureOk() (*string, bool)`

GetSignatureOk returns a tuple with the Signature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignature

`func (o *CollectionVersionSignatureResponse) SetSignature(v string)`

SetSignature sets Signature field to given value.

### HasSignature

`func (o *CollectionVersionSignatureResponse) HasSignature() bool`

HasSignature returns a boolean if a field has been set.

### GetPubkeyFingerprint

`func (o *CollectionVersionSignatureResponse) GetPubkeyFingerprint() string`

GetPubkeyFingerprint returns the PubkeyFingerprint field if non-nil, zero value otherwise.

### GetPubkeyFingerprintOk

`func (o *CollectionVersionSignatureResponse) GetPubkeyFingerprintOk() (*string, bool)`

GetPubkeyFingerprintOk returns a tuple with the PubkeyFingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPubkeyFingerprint

`func (o *CollectionVersionSignatureResponse) SetPubkeyFingerprint(v string)`

SetPubkeyFingerprint sets PubkeyFingerprint field to given value.


### GetSigningService

`func (o *CollectionVersionSignatureResponse) GetSigningService() string`

GetSigningService returns the SigningService field if non-nil, zero value otherwise.

### GetSigningServiceOk

`func (o *CollectionVersionSignatureResponse) GetSigningServiceOk() (*string, bool)`

GetSigningServiceOk returns a tuple with the SigningService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigningService

`func (o *CollectionVersionSignatureResponse) SetSigningService(v string)`

SetSigningService sets SigningService field to given value.

### HasSigningService

`func (o *CollectionVersionSignatureResponse) HasSigningService() bool`

HasSigningService returns a boolean if a field has been set.

### SetSigningServiceNil

`func (o *CollectionVersionSignatureResponse) SetSigningServiceNil(b bool)`

 SetSigningServiceNil sets the value for SigningService to be an explicit nil

### UnsetSigningService
`func (o *CollectionVersionSignatureResponse) UnsetSigningService()`

UnsetSigningService ensures that no value is present for SigningService, not even an explicit nil
### GetPulpCreated

`func (o *CollectionVersionSignatureResponse) GetPulpCreated() time.Time`

GetPulpCreated returns the PulpCreated field if non-nil, zero value otherwise.

### GetPulpCreatedOk

`func (o *CollectionVersionSignatureResponse) GetPulpCreatedOk() (*time.Time, bool)`

GetPulpCreatedOk returns a tuple with the PulpCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpCreated

`func (o *CollectionVersionSignatureResponse) SetPulpCreated(v time.Time)`

SetPulpCreated sets PulpCreated field to given value.

### HasPulpCreated

`func (o *CollectionVersionSignatureResponse) HasPulpCreated() bool`

HasPulpCreated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


