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

// checks if the DebRelease type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DebRelease{}

// DebRelease A Serializer for Release.
type DebRelease struct {
	// A URI of a repository the new content unit should be associated with.
	Repository *string `json:"repository,omitempty"`
	Codename string `json:"codename"`
	Suite string `json:"suite"`
	Distribution string `json:"distribution"`
	AdditionalProperties map[string]interface{}
}

type _DebRelease DebRelease

// NewDebRelease instantiates a new DebRelease object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDebRelease(codename string, suite string, distribution string) *DebRelease {
	this := DebRelease{}
	this.Codename = codename
	this.Suite = suite
	this.Distribution = distribution
	return &this
}

// NewDebReleaseWithDefaults instantiates a new DebRelease object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDebReleaseWithDefaults() *DebRelease {
	this := DebRelease{}
	return &this
}

// GetRepository returns the Repository field value if set, zero value otherwise.
func (o *DebRelease) GetRepository() string {
	if o == nil || IsNil(o.Repository) {
		var ret string
		return ret
	}
	return *o.Repository
}

// GetRepositoryOk returns a tuple with the Repository field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DebRelease) GetRepositoryOk() (*string, bool) {
	if o == nil || IsNil(o.Repository) {
		return nil, false
	}
	return o.Repository, true
}

// HasRepository returns a boolean if a field has been set.
func (o *DebRelease) HasRepository() bool {
	if o != nil && !IsNil(o.Repository) {
		return true
	}

	return false
}

// SetRepository gets a reference to the given string and assigns it to the Repository field.
func (o *DebRelease) SetRepository(v string) {
	o.Repository = &v
}

// GetCodename returns the Codename field value
func (o *DebRelease) GetCodename() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Codename
}

// GetCodenameOk returns a tuple with the Codename field value
// and a boolean to check if the value has been set.
func (o *DebRelease) GetCodenameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Codename, true
}

// SetCodename sets field value
func (o *DebRelease) SetCodename(v string) {
	o.Codename = v
}

// GetSuite returns the Suite field value
func (o *DebRelease) GetSuite() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Suite
}

// GetSuiteOk returns a tuple with the Suite field value
// and a boolean to check if the value has been set.
func (o *DebRelease) GetSuiteOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Suite, true
}

// SetSuite sets field value
func (o *DebRelease) SetSuite(v string) {
	o.Suite = v
}

// GetDistribution returns the Distribution field value
func (o *DebRelease) GetDistribution() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Distribution
}

// GetDistributionOk returns a tuple with the Distribution field value
// and a boolean to check if the value has been set.
func (o *DebRelease) GetDistributionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Distribution, true
}

// SetDistribution sets field value
func (o *DebRelease) SetDistribution(v string) {
	o.Distribution = v
}

func (o DebRelease) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DebRelease) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Repository) {
		toSerialize["repository"] = o.Repository
	}
	toSerialize["codename"] = o.Codename
	toSerialize["suite"] = o.Suite
	toSerialize["distribution"] = o.Distribution

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DebRelease) UnmarshalJSON(bytes []byte) (err error) {
	varDebRelease := _DebRelease{}

	if err = json.Unmarshal(bytes, &varDebRelease); err == nil {
		*o = DebRelease(varDebRelease)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "repository")
		delete(additionalProperties, "codename")
		delete(additionalProperties, "suite")
		delete(additionalProperties, "distribution")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDebRelease struct {
	value *DebRelease
	isSet bool
}

func (v NullableDebRelease) Get() *DebRelease {
	return v.value
}

func (v *NullableDebRelease) Set(val *DebRelease) {
	v.value = val
	v.isSet = true
}

func (v NullableDebRelease) IsSet() bool {
	return v.isSet
}

func (v *NullableDebRelease) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDebRelease(val *DebRelease) *NullableDebRelease {
	return &NullableDebRelease{value: val, isSet: true}
}

func (v NullableDebRelease) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDebRelease) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


