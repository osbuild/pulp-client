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

// checks if the DebInstallerFileIndex type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DebInstallerFileIndex{}

// DebInstallerFileIndex A serializer for InstallerFileIndex.
type DebInstallerFileIndex struct {
	// A URI of a repository the new content unit should be associated with.
	Repository *string `json:"repository,omitempty"`
	// A dict mapping relative paths inside the Content to the correspondingArtifact URLs. E.g.: {'relative/path': '/artifacts/1/'
	Artifacts map[string]interface{} `json:"artifacts"`
	// Component of the component - architecture combination.
	Component string `json:"component"`
	// Architecture of the component - architecture combination.
	Architecture string `json:"architecture"`
	// Path of directory containing MD5SUMS and SHA256SUMS relative to url.
	RelativePath *string `json:"relative_path,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DebInstallerFileIndex DebInstallerFileIndex

// NewDebInstallerFileIndex instantiates a new DebInstallerFileIndex object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDebInstallerFileIndex(artifacts map[string]interface{}, component string, architecture string) *DebInstallerFileIndex {
	this := DebInstallerFileIndex{}
	this.Artifacts = artifacts
	this.Component = component
	this.Architecture = architecture
	return &this
}

// NewDebInstallerFileIndexWithDefaults instantiates a new DebInstallerFileIndex object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDebInstallerFileIndexWithDefaults() *DebInstallerFileIndex {
	this := DebInstallerFileIndex{}
	return &this
}

// GetRepository returns the Repository field value if set, zero value otherwise.
func (o *DebInstallerFileIndex) GetRepository() string {
	if o == nil || IsNil(o.Repository) {
		var ret string
		return ret
	}
	return *o.Repository
}

// GetRepositoryOk returns a tuple with the Repository field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DebInstallerFileIndex) GetRepositoryOk() (*string, bool) {
	if o == nil || IsNil(o.Repository) {
		return nil, false
	}
	return o.Repository, true
}

// HasRepository returns a boolean if a field has been set.
func (o *DebInstallerFileIndex) HasRepository() bool {
	if o != nil && !IsNil(o.Repository) {
		return true
	}

	return false
}

// SetRepository gets a reference to the given string and assigns it to the Repository field.
func (o *DebInstallerFileIndex) SetRepository(v string) {
	o.Repository = &v
}

// GetArtifacts returns the Artifacts field value
func (o *DebInstallerFileIndex) GetArtifacts() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Artifacts
}

// GetArtifactsOk returns a tuple with the Artifacts field value
// and a boolean to check if the value has been set.
func (o *DebInstallerFileIndex) GetArtifactsOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Artifacts, true
}

// SetArtifacts sets field value
func (o *DebInstallerFileIndex) SetArtifacts(v map[string]interface{}) {
	o.Artifacts = v
}

// GetComponent returns the Component field value
func (o *DebInstallerFileIndex) GetComponent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Component
}

// GetComponentOk returns a tuple with the Component field value
// and a boolean to check if the value has been set.
func (o *DebInstallerFileIndex) GetComponentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Component, true
}

// SetComponent sets field value
func (o *DebInstallerFileIndex) SetComponent(v string) {
	o.Component = v
}

// GetArchitecture returns the Architecture field value
func (o *DebInstallerFileIndex) GetArchitecture() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Architecture
}

// GetArchitectureOk returns a tuple with the Architecture field value
// and a boolean to check if the value has been set.
func (o *DebInstallerFileIndex) GetArchitectureOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Architecture, true
}

// SetArchitecture sets field value
func (o *DebInstallerFileIndex) SetArchitecture(v string) {
	o.Architecture = v
}

// GetRelativePath returns the RelativePath field value if set, zero value otherwise.
func (o *DebInstallerFileIndex) GetRelativePath() string {
	if o == nil || IsNil(o.RelativePath) {
		var ret string
		return ret
	}
	return *o.RelativePath
}

// GetRelativePathOk returns a tuple with the RelativePath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DebInstallerFileIndex) GetRelativePathOk() (*string, bool) {
	if o == nil || IsNil(o.RelativePath) {
		return nil, false
	}
	return o.RelativePath, true
}

// HasRelativePath returns a boolean if a field has been set.
func (o *DebInstallerFileIndex) HasRelativePath() bool {
	if o != nil && !IsNil(o.RelativePath) {
		return true
	}

	return false
}

// SetRelativePath gets a reference to the given string and assigns it to the RelativePath field.
func (o *DebInstallerFileIndex) SetRelativePath(v string) {
	o.RelativePath = &v
}

func (o DebInstallerFileIndex) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DebInstallerFileIndex) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Repository) {
		toSerialize["repository"] = o.Repository
	}
	toSerialize["artifacts"] = o.Artifacts
	toSerialize["component"] = o.Component
	toSerialize["architecture"] = o.Architecture
	if !IsNil(o.RelativePath) {
		toSerialize["relative_path"] = o.RelativePath
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DebInstallerFileIndex) UnmarshalJSON(bytes []byte) (err error) {
	varDebInstallerFileIndex := _DebInstallerFileIndex{}

	if err = json.Unmarshal(bytes, &varDebInstallerFileIndex); err == nil {
		*o = DebInstallerFileIndex(varDebInstallerFileIndex)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "repository")
		delete(additionalProperties, "artifacts")
		delete(additionalProperties, "component")
		delete(additionalProperties, "architecture")
		delete(additionalProperties, "relative_path")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDebInstallerFileIndex struct {
	value *DebInstallerFileIndex
	isSet bool
}

func (v NullableDebInstallerFileIndex) Get() *DebInstallerFileIndex {
	return v.value
}

func (v *NullableDebInstallerFileIndex) Set(val *DebInstallerFileIndex) {
	v.value = val
	v.isSet = true
}

func (v NullableDebInstallerFileIndex) IsSet() bool {
	return v.isSet
}

func (v *NullableDebInstallerFileIndex) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDebInstallerFileIndex(val *DebInstallerFileIndex) *NullableDebInstallerFileIndex {
	return &NullableDebInstallerFileIndex{value: val, isSet: true}
}

func (v NullableDebInstallerFileIndex) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDebInstallerFileIndex) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


