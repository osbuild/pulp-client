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

// checks if the NamespaceLinkResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NamespaceLinkResponse{}

// NamespaceLinkResponse Provides backwards compatible interface for links with the legacy GalaxyNG API.
type NamespaceLinkResponse struct {
	Url string `json:"url"`
	Name string `json:"name"`
	AdditionalProperties map[string]interface{}
}

type _NamespaceLinkResponse NamespaceLinkResponse

// NewNamespaceLinkResponse instantiates a new NamespaceLinkResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNamespaceLinkResponse(url string, name string) *NamespaceLinkResponse {
	this := NamespaceLinkResponse{}
	this.Url = url
	this.Name = name
	return &this
}

// NewNamespaceLinkResponseWithDefaults instantiates a new NamespaceLinkResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNamespaceLinkResponseWithDefaults() *NamespaceLinkResponse {
	this := NamespaceLinkResponse{}
	return &this
}

// GetUrl returns the Url field value
func (o *NamespaceLinkResponse) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *NamespaceLinkResponse) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *NamespaceLinkResponse) SetUrl(v string) {
	o.Url = v
}

// GetName returns the Name field value
func (o *NamespaceLinkResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *NamespaceLinkResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *NamespaceLinkResponse) SetName(v string) {
	o.Name = v
}

func (o NamespaceLinkResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NamespaceLinkResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["url"] = o.Url
	toSerialize["name"] = o.Name

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *NamespaceLinkResponse) UnmarshalJSON(bytes []byte) (err error) {
	varNamespaceLinkResponse := _NamespaceLinkResponse{}

	if err = json.Unmarshal(bytes, &varNamespaceLinkResponse); err == nil {
		*o = NamespaceLinkResponse(varNamespaceLinkResponse)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "url")
		delete(additionalProperties, "name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNamespaceLinkResponse struct {
	value *NamespaceLinkResponse
	isSet bool
}

func (v NullableNamespaceLinkResponse) Get() *NamespaceLinkResponse {
	return v.value
}

func (v *NullableNamespaceLinkResponse) Set(val *NamespaceLinkResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableNamespaceLinkResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableNamespaceLinkResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNamespaceLinkResponse(val *NamespaceLinkResponse) *NullableNamespaceLinkResponse {
	return &NullableNamespaceLinkResponse{value: val, isSet: true}
}

func (v NullableNamespaceLinkResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNamespaceLinkResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


