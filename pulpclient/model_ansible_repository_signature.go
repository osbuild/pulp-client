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
)

// checks if the AnsibleRepositorySignature type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AnsibleRepositorySignature{}

// AnsibleRepositorySignature A serializer for the signing action.
type AnsibleRepositorySignature struct {
	// List of collection version hrefs to sign, use * to sign all content in repository
	ContentUnits []interface{} `json:"content_units"`
	// A signing service to use to sign the collections
	SigningService string `json:"signing_service"`
	AdditionalProperties map[string]interface{}
}

type _AnsibleRepositorySignature AnsibleRepositorySignature

// NewAnsibleRepositorySignature instantiates a new AnsibleRepositorySignature object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAnsibleRepositorySignature(contentUnits []interface{}, signingService string) *AnsibleRepositorySignature {
	this := AnsibleRepositorySignature{}
	this.ContentUnits = contentUnits
	this.SigningService = signingService
	return &this
}

// NewAnsibleRepositorySignatureWithDefaults instantiates a new AnsibleRepositorySignature object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAnsibleRepositorySignatureWithDefaults() *AnsibleRepositorySignature {
	this := AnsibleRepositorySignature{}
	return &this
}

// GetContentUnits returns the ContentUnits field value
func (o *AnsibleRepositorySignature) GetContentUnits() []interface{} {
	if o == nil {
		var ret []interface{}
		return ret
	}

	return o.ContentUnits
}

// GetContentUnitsOk returns a tuple with the ContentUnits field value
// and a boolean to check if the value has been set.
func (o *AnsibleRepositorySignature) GetContentUnitsOk() ([]interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return o.ContentUnits, true
}

// SetContentUnits sets field value
func (o *AnsibleRepositorySignature) SetContentUnits(v []interface{}) {
	o.ContentUnits = v
}

// GetSigningService returns the SigningService field value
func (o *AnsibleRepositorySignature) GetSigningService() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SigningService
}

// GetSigningServiceOk returns a tuple with the SigningService field value
// and a boolean to check if the value has been set.
func (o *AnsibleRepositorySignature) GetSigningServiceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SigningService, true
}

// SetSigningService sets field value
func (o *AnsibleRepositorySignature) SetSigningService(v string) {
	o.SigningService = v
}

func (o AnsibleRepositorySignature) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AnsibleRepositorySignature) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["content_units"] = o.ContentUnits
	toSerialize["signing_service"] = o.SigningService

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AnsibleRepositorySignature) UnmarshalJSON(bytes []byte) (err error) {
	varAnsibleRepositorySignature := _AnsibleRepositorySignature{}

	if err = json.Unmarshal(bytes, &varAnsibleRepositorySignature); err == nil {
		*o = AnsibleRepositorySignature(varAnsibleRepositorySignature)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "content_units")
		delete(additionalProperties, "signing_service")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAnsibleRepositorySignature struct {
	value *AnsibleRepositorySignature
	isSet bool
}

func (v NullableAnsibleRepositorySignature) Get() *AnsibleRepositorySignature {
	return v.value
}

func (v *NullableAnsibleRepositorySignature) Set(val *AnsibleRepositorySignature) {
	v.value = val
	v.isSet = true
}

func (v NullableAnsibleRepositorySignature) IsSet() bool {
	return v.isSet
}

func (v *NullableAnsibleRepositorySignature) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAnsibleRepositorySignature(val *AnsibleRepositorySignature) *NullableAnsibleRepositorySignature {
	return &NullableAnsibleRepositorySignature{value: val, isSet: true}
}

func (v NullableAnsibleRepositorySignature) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAnsibleRepositorySignature) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


