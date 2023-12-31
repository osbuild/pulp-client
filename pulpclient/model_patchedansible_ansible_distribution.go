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

// checks if the PatchedansibleAnsibleDistribution type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchedansibleAnsibleDistribution{}

// PatchedansibleAnsibleDistribution Serializer for Ansible Distributions.
type PatchedansibleAnsibleDistribution struct {
	// The base (relative) path component of the published url. Avoid paths that                     overlap with other distribution base paths (e.g. \"foo\" and \"foo/bar\")
	BasePath *string `json:"base_path,omitempty"`
	// An optional content-guard.
	ContentGuard NullableString `json:"content_guard,omitempty"`
	// A unique name. Ex, `rawhide` and `stable`.
	Name *string `json:"name,omitempty"`
	// The latest RepositoryVersion for this Repository will be served.
	Repository NullableString `json:"repository,omitempty"`
	// RepositoryVersion to be served
	RepositoryVersion NullableString `json:"repository_version,omitempty"`
	PulpLabels *map[string]string `json:"pulp_labels,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PatchedansibleAnsibleDistribution PatchedansibleAnsibleDistribution

// NewPatchedansibleAnsibleDistribution instantiates a new PatchedansibleAnsibleDistribution object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedansibleAnsibleDistribution() *PatchedansibleAnsibleDistribution {
	this := PatchedansibleAnsibleDistribution{}
	return &this
}

// NewPatchedansibleAnsibleDistributionWithDefaults instantiates a new PatchedansibleAnsibleDistribution object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedansibleAnsibleDistributionWithDefaults() *PatchedansibleAnsibleDistribution {
	this := PatchedansibleAnsibleDistribution{}
	return &this
}

// GetBasePath returns the BasePath field value if set, zero value otherwise.
func (o *PatchedansibleAnsibleDistribution) GetBasePath() string {
	if o == nil || IsNil(o.BasePath) {
		var ret string
		return ret
	}
	return *o.BasePath
}

// GetBasePathOk returns a tuple with the BasePath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedansibleAnsibleDistribution) GetBasePathOk() (*string, bool) {
	if o == nil || IsNil(o.BasePath) {
		return nil, false
	}
	return o.BasePath, true
}

// HasBasePath returns a boolean if a field has been set.
func (o *PatchedansibleAnsibleDistribution) HasBasePath() bool {
	if o != nil && !IsNil(o.BasePath) {
		return true
	}

	return false
}

// SetBasePath gets a reference to the given string and assigns it to the BasePath field.
func (o *PatchedansibleAnsibleDistribution) SetBasePath(v string) {
	o.BasePath = &v
}

// GetContentGuard returns the ContentGuard field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedansibleAnsibleDistribution) GetContentGuard() string {
	if o == nil || IsNil(o.ContentGuard.Get()) {
		var ret string
		return ret
	}
	return *o.ContentGuard.Get()
}

// GetContentGuardOk returns a tuple with the ContentGuard field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedansibleAnsibleDistribution) GetContentGuardOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ContentGuard.Get(), o.ContentGuard.IsSet()
}

// HasContentGuard returns a boolean if a field has been set.
func (o *PatchedansibleAnsibleDistribution) HasContentGuard() bool {
	if o != nil && o.ContentGuard.IsSet() {
		return true
	}

	return false
}

// SetContentGuard gets a reference to the given NullableString and assigns it to the ContentGuard field.
func (o *PatchedansibleAnsibleDistribution) SetContentGuard(v string) {
	o.ContentGuard.Set(&v)
}
// SetContentGuardNil sets the value for ContentGuard to be an explicit nil
func (o *PatchedansibleAnsibleDistribution) SetContentGuardNil() {
	o.ContentGuard.Set(nil)
}

// UnsetContentGuard ensures that no value is present for ContentGuard, not even an explicit nil
func (o *PatchedansibleAnsibleDistribution) UnsetContentGuard() {
	o.ContentGuard.Unset()
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PatchedansibleAnsibleDistribution) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedansibleAnsibleDistribution) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PatchedansibleAnsibleDistribution) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PatchedansibleAnsibleDistribution) SetName(v string) {
	o.Name = &v
}

// GetRepository returns the Repository field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedansibleAnsibleDistribution) GetRepository() string {
	if o == nil || IsNil(o.Repository.Get()) {
		var ret string
		return ret
	}
	return *o.Repository.Get()
}

// GetRepositoryOk returns a tuple with the Repository field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedansibleAnsibleDistribution) GetRepositoryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Repository.Get(), o.Repository.IsSet()
}

// HasRepository returns a boolean if a field has been set.
func (o *PatchedansibleAnsibleDistribution) HasRepository() bool {
	if o != nil && o.Repository.IsSet() {
		return true
	}

	return false
}

// SetRepository gets a reference to the given NullableString and assigns it to the Repository field.
func (o *PatchedansibleAnsibleDistribution) SetRepository(v string) {
	o.Repository.Set(&v)
}
// SetRepositoryNil sets the value for Repository to be an explicit nil
func (o *PatchedansibleAnsibleDistribution) SetRepositoryNil() {
	o.Repository.Set(nil)
}

// UnsetRepository ensures that no value is present for Repository, not even an explicit nil
func (o *PatchedansibleAnsibleDistribution) UnsetRepository() {
	o.Repository.Unset()
}

// GetRepositoryVersion returns the RepositoryVersion field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedansibleAnsibleDistribution) GetRepositoryVersion() string {
	if o == nil || IsNil(o.RepositoryVersion.Get()) {
		var ret string
		return ret
	}
	return *o.RepositoryVersion.Get()
}

// GetRepositoryVersionOk returns a tuple with the RepositoryVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedansibleAnsibleDistribution) GetRepositoryVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RepositoryVersion.Get(), o.RepositoryVersion.IsSet()
}

// HasRepositoryVersion returns a boolean if a field has been set.
func (o *PatchedansibleAnsibleDistribution) HasRepositoryVersion() bool {
	if o != nil && o.RepositoryVersion.IsSet() {
		return true
	}

	return false
}

// SetRepositoryVersion gets a reference to the given NullableString and assigns it to the RepositoryVersion field.
func (o *PatchedansibleAnsibleDistribution) SetRepositoryVersion(v string) {
	o.RepositoryVersion.Set(&v)
}
// SetRepositoryVersionNil sets the value for RepositoryVersion to be an explicit nil
func (o *PatchedansibleAnsibleDistribution) SetRepositoryVersionNil() {
	o.RepositoryVersion.Set(nil)
}

// UnsetRepositoryVersion ensures that no value is present for RepositoryVersion, not even an explicit nil
func (o *PatchedansibleAnsibleDistribution) UnsetRepositoryVersion() {
	o.RepositoryVersion.Unset()
}

// GetPulpLabels returns the PulpLabels field value if set, zero value otherwise.
func (o *PatchedansibleAnsibleDistribution) GetPulpLabels() map[string]string {
	if o == nil || IsNil(o.PulpLabels) {
		var ret map[string]string
		return ret
	}
	return *o.PulpLabels
}

// GetPulpLabelsOk returns a tuple with the PulpLabels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedansibleAnsibleDistribution) GetPulpLabelsOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.PulpLabels) {
		return nil, false
	}
	return o.PulpLabels, true
}

// HasPulpLabels returns a boolean if a field has been set.
func (o *PatchedansibleAnsibleDistribution) HasPulpLabels() bool {
	if o != nil && !IsNil(o.PulpLabels) {
		return true
	}

	return false
}

// SetPulpLabels gets a reference to the given map[string]string and assigns it to the PulpLabels field.
func (o *PatchedansibleAnsibleDistribution) SetPulpLabels(v map[string]string) {
	o.PulpLabels = &v
}

func (o PatchedansibleAnsibleDistribution) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchedansibleAnsibleDistribution) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BasePath) {
		toSerialize["base_path"] = o.BasePath
	}
	if o.ContentGuard.IsSet() {
		toSerialize["content_guard"] = o.ContentGuard.Get()
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if o.Repository.IsSet() {
		toSerialize["repository"] = o.Repository.Get()
	}
	if o.RepositoryVersion.IsSet() {
		toSerialize["repository_version"] = o.RepositoryVersion.Get()
	}
	if !IsNil(o.PulpLabels) {
		toSerialize["pulp_labels"] = o.PulpLabels
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PatchedansibleAnsibleDistribution) UnmarshalJSON(bytes []byte) (err error) {
	varPatchedansibleAnsibleDistribution := _PatchedansibleAnsibleDistribution{}

	if err = json.Unmarshal(bytes, &varPatchedansibleAnsibleDistribution); err == nil {
		*o = PatchedansibleAnsibleDistribution(varPatchedansibleAnsibleDistribution)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "base_path")
		delete(additionalProperties, "content_guard")
		delete(additionalProperties, "name")
		delete(additionalProperties, "repository")
		delete(additionalProperties, "repository_version")
		delete(additionalProperties, "pulp_labels")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePatchedansibleAnsibleDistribution struct {
	value *PatchedansibleAnsibleDistribution
	isSet bool
}

func (v NullablePatchedansibleAnsibleDistribution) Get() *PatchedansibleAnsibleDistribution {
	return v.value
}

func (v *NullablePatchedansibleAnsibleDistribution) Set(val *PatchedansibleAnsibleDistribution) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedansibleAnsibleDistribution) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedansibleAnsibleDistribution) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedansibleAnsibleDistribution(val *PatchedansibleAnsibleDistribution) *NullablePatchedansibleAnsibleDistribution {
	return &NullablePatchedansibleAnsibleDistribution{value: val, isSet: true}
}

func (v NullablePatchedansibleAnsibleDistribution) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedansibleAnsibleDistribution) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


