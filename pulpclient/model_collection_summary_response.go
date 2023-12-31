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

// checks if the CollectionSummaryResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CollectionSummaryResponse{}

// CollectionSummaryResponse Collection Version serializer without docs blob.
type CollectionSummaryResponse struct {
	PulpHref *string `json:"pulp_href,omitempty"`
	// The namespace of the collection.
	Namespace *string `json:"namespace,omitempty"`
	// The name of the collection.
	Name *string `json:"name,omitempty"`
	// The version of the collection.
	Version *string `json:"version,omitempty"`
	// The version of Ansible required to use the collection. Multiple versions can be separated with a comma.
	RequiresAnsible NullableString `json:"requires_ansible,omitempty"`
	// Timestamp of creation.
	PulpCreated *time.Time `json:"pulp_created,omitempty"`
	// A JSON field with data about the contents.
	Contents map[string]interface{} `json:"contents,omitempty"`
	// A dict declaring Collections that this collection requires to be installed for it to be usable.
	Dependencies map[string]interface{} `json:"dependencies,omitempty"`
	// A short summary description of the collection.
	Description *string `json:"description,omitempty"`
	Tags []AnsibleTagResponse `json:"tags,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CollectionSummaryResponse CollectionSummaryResponse

// NewCollectionSummaryResponse instantiates a new CollectionSummaryResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCollectionSummaryResponse() *CollectionSummaryResponse {
	this := CollectionSummaryResponse{}
	return &this
}

// NewCollectionSummaryResponseWithDefaults instantiates a new CollectionSummaryResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCollectionSummaryResponseWithDefaults() *CollectionSummaryResponse {
	this := CollectionSummaryResponse{}
	return &this
}

// GetPulpHref returns the PulpHref field value if set, zero value otherwise.
func (o *CollectionSummaryResponse) GetPulpHref() string {
	if o == nil || IsNil(o.PulpHref) {
		var ret string
		return ret
	}
	return *o.PulpHref
}

// GetPulpHrefOk returns a tuple with the PulpHref field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionSummaryResponse) GetPulpHrefOk() (*string, bool) {
	if o == nil || IsNil(o.PulpHref) {
		return nil, false
	}
	return o.PulpHref, true
}

// HasPulpHref returns a boolean if a field has been set.
func (o *CollectionSummaryResponse) HasPulpHref() bool {
	if o != nil && !IsNil(o.PulpHref) {
		return true
	}

	return false
}

// SetPulpHref gets a reference to the given string and assigns it to the PulpHref field.
func (o *CollectionSummaryResponse) SetPulpHref(v string) {
	o.PulpHref = &v
}

// GetNamespace returns the Namespace field value if set, zero value otherwise.
func (o *CollectionSummaryResponse) GetNamespace() string {
	if o == nil || IsNil(o.Namespace) {
		var ret string
		return ret
	}
	return *o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionSummaryResponse) GetNamespaceOk() (*string, bool) {
	if o == nil || IsNil(o.Namespace) {
		return nil, false
	}
	return o.Namespace, true
}

// HasNamespace returns a boolean if a field has been set.
func (o *CollectionSummaryResponse) HasNamespace() bool {
	if o != nil && !IsNil(o.Namespace) {
		return true
	}

	return false
}

// SetNamespace gets a reference to the given string and assigns it to the Namespace field.
func (o *CollectionSummaryResponse) SetNamespace(v string) {
	o.Namespace = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CollectionSummaryResponse) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionSummaryResponse) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CollectionSummaryResponse) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CollectionSummaryResponse) SetName(v string) {
	o.Name = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *CollectionSummaryResponse) GetVersion() string {
	if o == nil || IsNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionSummaryResponse) GetVersionOk() (*string, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *CollectionSummaryResponse) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *CollectionSummaryResponse) SetVersion(v string) {
	o.Version = &v
}

// GetRequiresAnsible returns the RequiresAnsible field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CollectionSummaryResponse) GetRequiresAnsible() string {
	if o == nil || IsNil(o.RequiresAnsible.Get()) {
		var ret string
		return ret
	}
	return *o.RequiresAnsible.Get()
}

// GetRequiresAnsibleOk returns a tuple with the RequiresAnsible field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CollectionSummaryResponse) GetRequiresAnsibleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RequiresAnsible.Get(), o.RequiresAnsible.IsSet()
}

// HasRequiresAnsible returns a boolean if a field has been set.
func (o *CollectionSummaryResponse) HasRequiresAnsible() bool {
	if o != nil && o.RequiresAnsible.IsSet() {
		return true
	}

	return false
}

// SetRequiresAnsible gets a reference to the given NullableString and assigns it to the RequiresAnsible field.
func (o *CollectionSummaryResponse) SetRequiresAnsible(v string) {
	o.RequiresAnsible.Set(&v)
}
// SetRequiresAnsibleNil sets the value for RequiresAnsible to be an explicit nil
func (o *CollectionSummaryResponse) SetRequiresAnsibleNil() {
	o.RequiresAnsible.Set(nil)
}

// UnsetRequiresAnsible ensures that no value is present for RequiresAnsible, not even an explicit nil
func (o *CollectionSummaryResponse) UnsetRequiresAnsible() {
	o.RequiresAnsible.Unset()
}

// GetPulpCreated returns the PulpCreated field value if set, zero value otherwise.
func (o *CollectionSummaryResponse) GetPulpCreated() time.Time {
	if o == nil || IsNil(o.PulpCreated) {
		var ret time.Time
		return ret
	}
	return *o.PulpCreated
}

// GetPulpCreatedOk returns a tuple with the PulpCreated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionSummaryResponse) GetPulpCreatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.PulpCreated) {
		return nil, false
	}
	return o.PulpCreated, true
}

// HasPulpCreated returns a boolean if a field has been set.
func (o *CollectionSummaryResponse) HasPulpCreated() bool {
	if o != nil && !IsNil(o.PulpCreated) {
		return true
	}

	return false
}

// SetPulpCreated gets a reference to the given time.Time and assigns it to the PulpCreated field.
func (o *CollectionSummaryResponse) SetPulpCreated(v time.Time) {
	o.PulpCreated = &v
}

// GetContents returns the Contents field value if set, zero value otherwise.
func (o *CollectionSummaryResponse) GetContents() map[string]interface{} {
	if o == nil || IsNil(o.Contents) {
		var ret map[string]interface{}
		return ret
	}
	return o.Contents
}

// GetContentsOk returns a tuple with the Contents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionSummaryResponse) GetContentsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Contents) {
		return map[string]interface{}{}, false
	}
	return o.Contents, true
}

// HasContents returns a boolean if a field has been set.
func (o *CollectionSummaryResponse) HasContents() bool {
	if o != nil && !IsNil(o.Contents) {
		return true
	}

	return false
}

// SetContents gets a reference to the given map[string]interface{} and assigns it to the Contents field.
func (o *CollectionSummaryResponse) SetContents(v map[string]interface{}) {
	o.Contents = v
}

// GetDependencies returns the Dependencies field value if set, zero value otherwise.
func (o *CollectionSummaryResponse) GetDependencies() map[string]interface{} {
	if o == nil || IsNil(o.Dependencies) {
		var ret map[string]interface{}
		return ret
	}
	return o.Dependencies
}

// GetDependenciesOk returns a tuple with the Dependencies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionSummaryResponse) GetDependenciesOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Dependencies) {
		return map[string]interface{}{}, false
	}
	return o.Dependencies, true
}

// HasDependencies returns a boolean if a field has been set.
func (o *CollectionSummaryResponse) HasDependencies() bool {
	if o != nil && !IsNil(o.Dependencies) {
		return true
	}

	return false
}

// SetDependencies gets a reference to the given map[string]interface{} and assigns it to the Dependencies field.
func (o *CollectionSummaryResponse) SetDependencies(v map[string]interface{}) {
	o.Dependencies = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CollectionSummaryResponse) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionSummaryResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CollectionSummaryResponse) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CollectionSummaryResponse) SetDescription(v string) {
	o.Description = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *CollectionSummaryResponse) GetTags() []AnsibleTagResponse {
	if o == nil || IsNil(o.Tags) {
		var ret []AnsibleTagResponse
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionSummaryResponse) GetTagsOk() ([]AnsibleTagResponse, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *CollectionSummaryResponse) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []AnsibleTagResponse and assigns it to the Tags field.
func (o *CollectionSummaryResponse) SetTags(v []AnsibleTagResponse) {
	o.Tags = v
}

func (o CollectionSummaryResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CollectionSummaryResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PulpHref) {
		toSerialize["pulp_href"] = o.PulpHref
	}
	if !IsNil(o.Namespace) {
		toSerialize["namespace"] = o.Namespace
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	if o.RequiresAnsible.IsSet() {
		toSerialize["requires_ansible"] = o.RequiresAnsible.Get()
	}
	if !IsNil(o.PulpCreated) {
		toSerialize["pulp_created"] = o.PulpCreated
	}
	if !IsNil(o.Contents) {
		toSerialize["contents"] = o.Contents
	}
	if !IsNil(o.Dependencies) {
		toSerialize["dependencies"] = o.Dependencies
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CollectionSummaryResponse) UnmarshalJSON(bytes []byte) (err error) {
	varCollectionSummaryResponse := _CollectionSummaryResponse{}

	if err = json.Unmarshal(bytes, &varCollectionSummaryResponse); err == nil {
		*o = CollectionSummaryResponse(varCollectionSummaryResponse)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "pulp_href")
		delete(additionalProperties, "namespace")
		delete(additionalProperties, "name")
		delete(additionalProperties, "version")
		delete(additionalProperties, "requires_ansible")
		delete(additionalProperties, "pulp_created")
		delete(additionalProperties, "contents")
		delete(additionalProperties, "dependencies")
		delete(additionalProperties, "description")
		delete(additionalProperties, "tags")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCollectionSummaryResponse struct {
	value *CollectionSummaryResponse
	isSet bool
}

func (v NullableCollectionSummaryResponse) Get() *CollectionSummaryResponse {
	return v.value
}

func (v *NullableCollectionSummaryResponse) Set(val *CollectionSummaryResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCollectionSummaryResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCollectionSummaryResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollectionSummaryResponse(val *CollectionSummaryResponse) *NullableCollectionSummaryResponse {
	return &NullableCollectionSummaryResponse{value: val, isSet: true}
}

func (v NullableCollectionSummaryResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollectionSummaryResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


