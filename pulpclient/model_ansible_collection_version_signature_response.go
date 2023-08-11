/*
Pulp 3 API

Fetch, Upload, Organize, and Distribute Software Packages

API version: v3
Contact: pulp-list@redhat.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package pulpclient

import (
	"encoding/json"
	"time"
)

// checks if the AnsibleCollectionVersionSignatureResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AnsibleCollectionVersionSignatureResponse{}

// AnsibleCollectionVersionSignatureResponse A serializer for signature models.
type AnsibleCollectionVersionSignatureResponse struct {
	PulpHref *string `json:"pulp_href,omitempty"`
	// Timestamp of creation.
	PulpCreated *time.Time `json:"pulp_created,omitempty"`
	// The content this signature is pointing to.
	SignedCollection string `json:"signed_collection"`
	// The fingerprint of the public key.
	PubkeyFingerprint *string `json:"pubkey_fingerprint,omitempty"`
	// The signing service used to create the signature.
	SigningService NullableString `json:"signing_service,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AnsibleCollectionVersionSignatureResponse AnsibleCollectionVersionSignatureResponse

// NewAnsibleCollectionVersionSignatureResponse instantiates a new AnsibleCollectionVersionSignatureResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAnsibleCollectionVersionSignatureResponse(signedCollection string) *AnsibleCollectionVersionSignatureResponse {
	this := AnsibleCollectionVersionSignatureResponse{}
	this.SignedCollection = signedCollection
	return &this
}

// NewAnsibleCollectionVersionSignatureResponseWithDefaults instantiates a new AnsibleCollectionVersionSignatureResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAnsibleCollectionVersionSignatureResponseWithDefaults() *AnsibleCollectionVersionSignatureResponse {
	this := AnsibleCollectionVersionSignatureResponse{}
	return &this
}

// GetPulpHref returns the PulpHref field value if set, zero value otherwise.
func (o *AnsibleCollectionVersionSignatureResponse) GetPulpHref() string {
	if o == nil || IsNil(o.PulpHref) {
		var ret string
		return ret
	}
	return *o.PulpHref
}

// GetPulpHrefOk returns a tuple with the PulpHref field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnsibleCollectionVersionSignatureResponse) GetPulpHrefOk() (*string, bool) {
	if o == nil || IsNil(o.PulpHref) {
		return nil, false
	}
	return o.PulpHref, true
}

// HasPulpHref returns a boolean if a field has been set.
func (o *AnsibleCollectionVersionSignatureResponse) HasPulpHref() bool {
	if o != nil && !IsNil(o.PulpHref) {
		return true
	}

	return false
}

// SetPulpHref gets a reference to the given string and assigns it to the PulpHref field.
func (o *AnsibleCollectionVersionSignatureResponse) SetPulpHref(v string) {
	o.PulpHref = &v
}

// GetPulpCreated returns the PulpCreated field value if set, zero value otherwise.
func (o *AnsibleCollectionVersionSignatureResponse) GetPulpCreated() time.Time {
	if o == nil || IsNil(o.PulpCreated) {
		var ret time.Time
		return ret
	}
	return *o.PulpCreated
}

// GetPulpCreatedOk returns a tuple with the PulpCreated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnsibleCollectionVersionSignatureResponse) GetPulpCreatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.PulpCreated) {
		return nil, false
	}
	return o.PulpCreated, true
}

// HasPulpCreated returns a boolean if a field has been set.
func (o *AnsibleCollectionVersionSignatureResponse) HasPulpCreated() bool {
	if o != nil && !IsNil(o.PulpCreated) {
		return true
	}

	return false
}

// SetPulpCreated gets a reference to the given time.Time and assigns it to the PulpCreated field.
func (o *AnsibleCollectionVersionSignatureResponse) SetPulpCreated(v time.Time) {
	o.PulpCreated = &v
}

// GetSignedCollection returns the SignedCollection field value
func (o *AnsibleCollectionVersionSignatureResponse) GetSignedCollection() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SignedCollection
}

// GetSignedCollectionOk returns a tuple with the SignedCollection field value
// and a boolean to check if the value has been set.
func (o *AnsibleCollectionVersionSignatureResponse) GetSignedCollectionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SignedCollection, true
}

// SetSignedCollection sets field value
func (o *AnsibleCollectionVersionSignatureResponse) SetSignedCollection(v string) {
	o.SignedCollection = v
}

// GetPubkeyFingerprint returns the PubkeyFingerprint field value if set, zero value otherwise.
func (o *AnsibleCollectionVersionSignatureResponse) GetPubkeyFingerprint() string {
	if o == nil || IsNil(o.PubkeyFingerprint) {
		var ret string
		return ret
	}
	return *o.PubkeyFingerprint
}

// GetPubkeyFingerprintOk returns a tuple with the PubkeyFingerprint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnsibleCollectionVersionSignatureResponse) GetPubkeyFingerprintOk() (*string, bool) {
	if o == nil || IsNil(o.PubkeyFingerprint) {
		return nil, false
	}
	return o.PubkeyFingerprint, true
}

// HasPubkeyFingerprint returns a boolean if a field has been set.
func (o *AnsibleCollectionVersionSignatureResponse) HasPubkeyFingerprint() bool {
	if o != nil && !IsNil(o.PubkeyFingerprint) {
		return true
	}

	return false
}

// SetPubkeyFingerprint gets a reference to the given string and assigns it to the PubkeyFingerprint field.
func (o *AnsibleCollectionVersionSignatureResponse) SetPubkeyFingerprint(v string) {
	o.PubkeyFingerprint = &v
}

// GetSigningService returns the SigningService field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AnsibleCollectionVersionSignatureResponse) GetSigningService() string {
	if o == nil || IsNil(o.SigningService.Get()) {
		var ret string
		return ret
	}
	return *o.SigningService.Get()
}

// GetSigningServiceOk returns a tuple with the SigningService field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AnsibleCollectionVersionSignatureResponse) GetSigningServiceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SigningService.Get(), o.SigningService.IsSet()
}

// HasSigningService returns a boolean if a field has been set.
func (o *AnsibleCollectionVersionSignatureResponse) HasSigningService() bool {
	if o != nil && o.SigningService.IsSet() {
		return true
	}

	return false
}

// SetSigningService gets a reference to the given NullableString and assigns it to the SigningService field.
func (o *AnsibleCollectionVersionSignatureResponse) SetSigningService(v string) {
	o.SigningService.Set(&v)
}
// SetSigningServiceNil sets the value for SigningService to be an explicit nil
func (o *AnsibleCollectionVersionSignatureResponse) SetSigningServiceNil() {
	o.SigningService.Set(nil)
}

// UnsetSigningService ensures that no value is present for SigningService, not even an explicit nil
func (o *AnsibleCollectionVersionSignatureResponse) UnsetSigningService() {
	o.SigningService.Unset()
}

func (o AnsibleCollectionVersionSignatureResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AnsibleCollectionVersionSignatureResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PulpHref) {
		toSerialize["pulp_href"] = o.PulpHref
	}
	if !IsNil(o.PulpCreated) {
		toSerialize["pulp_created"] = o.PulpCreated
	}
	toSerialize["signed_collection"] = o.SignedCollection
	if !IsNil(o.PubkeyFingerprint) {
		toSerialize["pubkey_fingerprint"] = o.PubkeyFingerprint
	}
	if o.SigningService.IsSet() {
		toSerialize["signing_service"] = o.SigningService.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AnsibleCollectionVersionSignatureResponse) UnmarshalJSON(bytes []byte) (err error) {
	varAnsibleCollectionVersionSignatureResponse := _AnsibleCollectionVersionSignatureResponse{}

	if err = json.Unmarshal(bytes, &varAnsibleCollectionVersionSignatureResponse); err == nil {
		*o = AnsibleCollectionVersionSignatureResponse(varAnsibleCollectionVersionSignatureResponse)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "pulp_href")
		delete(additionalProperties, "pulp_created")
		delete(additionalProperties, "signed_collection")
		delete(additionalProperties, "pubkey_fingerprint")
		delete(additionalProperties, "signing_service")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAnsibleCollectionVersionSignatureResponse struct {
	value *AnsibleCollectionVersionSignatureResponse
	isSet bool
}

func (v NullableAnsibleCollectionVersionSignatureResponse) Get() *AnsibleCollectionVersionSignatureResponse {
	return v.value
}

func (v *NullableAnsibleCollectionVersionSignatureResponse) Set(val *AnsibleCollectionVersionSignatureResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAnsibleCollectionVersionSignatureResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAnsibleCollectionVersionSignatureResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAnsibleCollectionVersionSignatureResponse(val *AnsibleCollectionVersionSignatureResponse) *NullableAnsibleCollectionVersionSignatureResponse {
	return &NullableAnsibleCollectionVersionSignatureResponse{value: val, isSet: true}
}

func (v NullableAnsibleCollectionVersionSignatureResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAnsibleCollectionVersionSignatureResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

