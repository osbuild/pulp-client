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

// checks if the DebVerbatimPublication type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DebVerbatimPublication{}

// DebVerbatimPublication A Serializer for VerbatimPublication.
type DebVerbatimPublication struct {
	RepositoryVersion *string `json:"repository_version,omitempty"`
	// A URI of the repository to be published.
	Repository *string `json:"repository,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DebVerbatimPublication DebVerbatimPublication

// NewDebVerbatimPublication instantiates a new DebVerbatimPublication object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDebVerbatimPublication() *DebVerbatimPublication {
	this := DebVerbatimPublication{}
	return &this
}

// NewDebVerbatimPublicationWithDefaults instantiates a new DebVerbatimPublication object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDebVerbatimPublicationWithDefaults() *DebVerbatimPublication {
	this := DebVerbatimPublication{}
	return &this
}

// GetRepositoryVersion returns the RepositoryVersion field value if set, zero value otherwise.
func (o *DebVerbatimPublication) GetRepositoryVersion() string {
	if o == nil || IsNil(o.RepositoryVersion) {
		var ret string
		return ret
	}
	return *o.RepositoryVersion
}

// GetRepositoryVersionOk returns a tuple with the RepositoryVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DebVerbatimPublication) GetRepositoryVersionOk() (*string, bool) {
	if o == nil || IsNil(o.RepositoryVersion) {
		return nil, false
	}
	return o.RepositoryVersion, true
}

// HasRepositoryVersion returns a boolean if a field has been set.
func (o *DebVerbatimPublication) HasRepositoryVersion() bool {
	if o != nil && !IsNil(o.RepositoryVersion) {
		return true
	}

	return false
}

// SetRepositoryVersion gets a reference to the given string and assigns it to the RepositoryVersion field.
func (o *DebVerbatimPublication) SetRepositoryVersion(v string) {
	o.RepositoryVersion = &v
}

// GetRepository returns the Repository field value if set, zero value otherwise.
func (o *DebVerbatimPublication) GetRepository() string {
	if o == nil || IsNil(o.Repository) {
		var ret string
		return ret
	}
	return *o.Repository
}

// GetRepositoryOk returns a tuple with the Repository field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DebVerbatimPublication) GetRepositoryOk() (*string, bool) {
	if o == nil || IsNil(o.Repository) {
		return nil, false
	}
	return o.Repository, true
}

// HasRepository returns a boolean if a field has been set.
func (o *DebVerbatimPublication) HasRepository() bool {
	if o != nil && !IsNil(o.Repository) {
		return true
	}

	return false
}

// SetRepository gets a reference to the given string and assigns it to the Repository field.
func (o *DebVerbatimPublication) SetRepository(v string) {
	o.Repository = &v
}

func (o DebVerbatimPublication) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DebVerbatimPublication) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.RepositoryVersion) {
		toSerialize["repository_version"] = o.RepositoryVersion
	}
	if !IsNil(o.Repository) {
		toSerialize["repository"] = o.Repository
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DebVerbatimPublication) UnmarshalJSON(bytes []byte) (err error) {
	varDebVerbatimPublication := _DebVerbatimPublication{}

	if err = json.Unmarshal(bytes, &varDebVerbatimPublication); err == nil {
		*o = DebVerbatimPublication(varDebVerbatimPublication)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "repository_version")
		delete(additionalProperties, "repository")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDebVerbatimPublication struct {
	value *DebVerbatimPublication
	isSet bool
}

func (v NullableDebVerbatimPublication) Get() *DebVerbatimPublication {
	return v.value
}

func (v *NullableDebVerbatimPublication) Set(val *DebVerbatimPublication) {
	v.value = val
	v.isSet = true
}

func (v NullableDebVerbatimPublication) IsSet() bool {
	return v.isSet
}

func (v *NullableDebVerbatimPublication) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDebVerbatimPublication(val *DebVerbatimPublication) *NullableDebVerbatimPublication {
	return &NullableDebVerbatimPublication{value: val, isSet: true}
}

func (v NullableDebVerbatimPublication) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDebVerbatimPublication) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


