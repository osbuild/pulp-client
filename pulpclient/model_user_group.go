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

// checks if the UserGroup type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserGroup{}

// UserGroup Serializer for Groups that belong to an User.
type UserGroup struct {
	// Name.
	Name string `json:"name"`
	AdditionalProperties map[string]interface{}
}

type _UserGroup UserGroup

// NewUserGroup instantiates a new UserGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserGroup(name string) *UserGroup {
	this := UserGroup{}
	this.Name = name
	return &this
}

// NewUserGroupWithDefaults instantiates a new UserGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserGroupWithDefaults() *UserGroup {
	this := UserGroup{}
	return &this
}

// GetName returns the Name field value
func (o *UserGroup) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *UserGroup) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *UserGroup) SetName(v string) {
	o.Name = v
}

func (o UserGroup) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserGroup) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UserGroup) UnmarshalJSON(bytes []byte) (err error) {
	varUserGroup := _UserGroup{}

	if err = json.Unmarshal(bytes, &varUserGroup); err == nil {
		*o = UserGroup(varUserGroup)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUserGroup struct {
	value *UserGroup
	isSet bool
}

func (v NullableUserGroup) Get() *UserGroup {
	return v.value
}

func (v *NullableUserGroup) Set(val *UserGroup) {
	v.value = val
	v.isSet = true
}

func (v NullableUserGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableUserGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserGroup(val *UserGroup) *NullableUserGroup {
	return &NullableUserGroup{value: val, isSet: true}
}

func (v NullableUserGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


