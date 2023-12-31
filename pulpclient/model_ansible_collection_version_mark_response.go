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

// checks if the AnsibleCollectionVersionMarkResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AnsibleCollectionVersionMarkResponse{}

// AnsibleCollectionVersionMarkResponse A serializer for mark models.
type AnsibleCollectionVersionMarkResponse struct {
	// Timestamp of creation.
	PulpCreated *time.Time `json:"pulp_created,omitempty"`
	PulpHref *string `json:"pulp_href,omitempty"`
	// The content this mark is pointing to.
	MarkedCollection string `json:"marked_collection"`
	// The string value of this mark.
	Value string `json:"value"`
	AdditionalProperties map[string]interface{}
}

type _AnsibleCollectionVersionMarkResponse AnsibleCollectionVersionMarkResponse

// NewAnsibleCollectionVersionMarkResponse instantiates a new AnsibleCollectionVersionMarkResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAnsibleCollectionVersionMarkResponse(markedCollection string, value string) *AnsibleCollectionVersionMarkResponse {
	this := AnsibleCollectionVersionMarkResponse{}
	this.MarkedCollection = markedCollection
	this.Value = value
	return &this
}

// NewAnsibleCollectionVersionMarkResponseWithDefaults instantiates a new AnsibleCollectionVersionMarkResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAnsibleCollectionVersionMarkResponseWithDefaults() *AnsibleCollectionVersionMarkResponse {
	this := AnsibleCollectionVersionMarkResponse{}
	return &this
}

// GetPulpCreated returns the PulpCreated field value if set, zero value otherwise.
func (o *AnsibleCollectionVersionMarkResponse) GetPulpCreated() time.Time {
	if o == nil || IsNil(o.PulpCreated) {
		var ret time.Time
		return ret
	}
	return *o.PulpCreated
}

// GetPulpCreatedOk returns a tuple with the PulpCreated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnsibleCollectionVersionMarkResponse) GetPulpCreatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.PulpCreated) {
		return nil, false
	}
	return o.PulpCreated, true
}

// HasPulpCreated returns a boolean if a field has been set.
func (o *AnsibleCollectionVersionMarkResponse) HasPulpCreated() bool {
	if o != nil && !IsNil(o.PulpCreated) {
		return true
	}

	return false
}

// SetPulpCreated gets a reference to the given time.Time and assigns it to the PulpCreated field.
func (o *AnsibleCollectionVersionMarkResponse) SetPulpCreated(v time.Time) {
	o.PulpCreated = &v
}

// GetPulpHref returns the PulpHref field value if set, zero value otherwise.
func (o *AnsibleCollectionVersionMarkResponse) GetPulpHref() string {
	if o == nil || IsNil(o.PulpHref) {
		var ret string
		return ret
	}
	return *o.PulpHref
}

// GetPulpHrefOk returns a tuple with the PulpHref field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnsibleCollectionVersionMarkResponse) GetPulpHrefOk() (*string, bool) {
	if o == nil || IsNil(o.PulpHref) {
		return nil, false
	}
	return o.PulpHref, true
}

// HasPulpHref returns a boolean if a field has been set.
func (o *AnsibleCollectionVersionMarkResponse) HasPulpHref() bool {
	if o != nil && !IsNil(o.PulpHref) {
		return true
	}

	return false
}

// SetPulpHref gets a reference to the given string and assigns it to the PulpHref field.
func (o *AnsibleCollectionVersionMarkResponse) SetPulpHref(v string) {
	o.PulpHref = &v
}

// GetMarkedCollection returns the MarkedCollection field value
func (o *AnsibleCollectionVersionMarkResponse) GetMarkedCollection() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MarkedCollection
}

// GetMarkedCollectionOk returns a tuple with the MarkedCollection field value
// and a boolean to check if the value has been set.
func (o *AnsibleCollectionVersionMarkResponse) GetMarkedCollectionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MarkedCollection, true
}

// SetMarkedCollection sets field value
func (o *AnsibleCollectionVersionMarkResponse) SetMarkedCollection(v string) {
	o.MarkedCollection = v
}

// GetValue returns the Value field value
func (o *AnsibleCollectionVersionMarkResponse) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *AnsibleCollectionVersionMarkResponse) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *AnsibleCollectionVersionMarkResponse) SetValue(v string) {
	o.Value = v
}

func (o AnsibleCollectionVersionMarkResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AnsibleCollectionVersionMarkResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PulpCreated) {
		toSerialize["pulp_created"] = o.PulpCreated
	}
	if !IsNil(o.PulpHref) {
		toSerialize["pulp_href"] = o.PulpHref
	}
	toSerialize["marked_collection"] = o.MarkedCollection
	toSerialize["value"] = o.Value

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AnsibleCollectionVersionMarkResponse) UnmarshalJSON(bytes []byte) (err error) {
	varAnsibleCollectionVersionMarkResponse := _AnsibleCollectionVersionMarkResponse{}

	if err = json.Unmarshal(bytes, &varAnsibleCollectionVersionMarkResponse); err == nil {
		*o = AnsibleCollectionVersionMarkResponse(varAnsibleCollectionVersionMarkResponse)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "pulp_created")
		delete(additionalProperties, "pulp_href")
		delete(additionalProperties, "marked_collection")
		delete(additionalProperties, "value")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAnsibleCollectionVersionMarkResponse struct {
	value *AnsibleCollectionVersionMarkResponse
	isSet bool
}

func (v NullableAnsibleCollectionVersionMarkResponse) Get() *AnsibleCollectionVersionMarkResponse {
	return v.value
}

func (v *NullableAnsibleCollectionVersionMarkResponse) Set(val *AnsibleCollectionVersionMarkResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAnsibleCollectionVersionMarkResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAnsibleCollectionVersionMarkResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAnsibleCollectionVersionMarkResponse(val *AnsibleCollectionVersionMarkResponse) *NullableAnsibleCollectionVersionMarkResponse {
	return &NullableAnsibleCollectionVersionMarkResponse{value: val, isSet: true}
}

func (v NullableAnsibleCollectionVersionMarkResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAnsibleCollectionVersionMarkResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


