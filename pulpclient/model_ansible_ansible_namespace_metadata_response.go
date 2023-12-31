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

// checks if the AnsibleAnsibleNamespaceMetadataResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AnsibleAnsibleNamespaceMetadataResponse{}

// AnsibleAnsibleNamespaceMetadataResponse A serializer for Namespaces.
type AnsibleAnsibleNamespaceMetadataResponse struct {
	PulpHref *string `json:"pulp_href,omitempty"`
	// Required named, only accepts lowercase, numbers and underscores.
	Name string `json:"name"`
	// Optional namespace company owner.
	Company *string `json:"company,omitempty"`
	// Optional namespace contact email.
	Email *string `json:"email,omitempty"`
	// Optional short description.
	Description *string `json:"description,omitempty"`
	// Optional resource page in markdown format.
	Resources *string `json:"resources,omitempty"`
	// Labeled related links.
	Links []NamespaceLinkResponse `json:"links,omitempty"`
	// SHA256 digest of avatar image if present.
	AvatarSha256 *string `json:"avatar_sha256,omitempty"`
	// Download link for avatar image if present.
	AvatarUrl *string `json:"avatar_url,omitempty"`
	MetadataSha256 *string `json:"metadata_sha256,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AnsibleAnsibleNamespaceMetadataResponse AnsibleAnsibleNamespaceMetadataResponse

// NewAnsibleAnsibleNamespaceMetadataResponse instantiates a new AnsibleAnsibleNamespaceMetadataResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAnsibleAnsibleNamespaceMetadataResponse(name string) *AnsibleAnsibleNamespaceMetadataResponse {
	this := AnsibleAnsibleNamespaceMetadataResponse{}
	this.Name = name
	return &this
}

// NewAnsibleAnsibleNamespaceMetadataResponseWithDefaults instantiates a new AnsibleAnsibleNamespaceMetadataResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAnsibleAnsibleNamespaceMetadataResponseWithDefaults() *AnsibleAnsibleNamespaceMetadataResponse {
	this := AnsibleAnsibleNamespaceMetadataResponse{}
	return &this
}

// GetPulpHref returns the PulpHref field value if set, zero value otherwise.
func (o *AnsibleAnsibleNamespaceMetadataResponse) GetPulpHref() string {
	if o == nil || IsNil(o.PulpHref) {
		var ret string
		return ret
	}
	return *o.PulpHref
}

// GetPulpHrefOk returns a tuple with the PulpHref field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnsibleAnsibleNamespaceMetadataResponse) GetPulpHrefOk() (*string, bool) {
	if o == nil || IsNil(o.PulpHref) {
		return nil, false
	}
	return o.PulpHref, true
}

// HasPulpHref returns a boolean if a field has been set.
func (o *AnsibleAnsibleNamespaceMetadataResponse) HasPulpHref() bool {
	if o != nil && !IsNil(o.PulpHref) {
		return true
	}

	return false
}

// SetPulpHref gets a reference to the given string and assigns it to the PulpHref field.
func (o *AnsibleAnsibleNamespaceMetadataResponse) SetPulpHref(v string) {
	o.PulpHref = &v
}

// GetName returns the Name field value
func (o *AnsibleAnsibleNamespaceMetadataResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AnsibleAnsibleNamespaceMetadataResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AnsibleAnsibleNamespaceMetadataResponse) SetName(v string) {
	o.Name = v
}

// GetCompany returns the Company field value if set, zero value otherwise.
func (o *AnsibleAnsibleNamespaceMetadataResponse) GetCompany() string {
	if o == nil || IsNil(o.Company) {
		var ret string
		return ret
	}
	return *o.Company
}

// GetCompanyOk returns a tuple with the Company field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnsibleAnsibleNamespaceMetadataResponse) GetCompanyOk() (*string, bool) {
	if o == nil || IsNil(o.Company) {
		return nil, false
	}
	return o.Company, true
}

// HasCompany returns a boolean if a field has been set.
func (o *AnsibleAnsibleNamespaceMetadataResponse) HasCompany() bool {
	if o != nil && !IsNil(o.Company) {
		return true
	}

	return false
}

// SetCompany gets a reference to the given string and assigns it to the Company field.
func (o *AnsibleAnsibleNamespaceMetadataResponse) SetCompany(v string) {
	o.Company = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *AnsibleAnsibleNamespaceMetadataResponse) GetEmail() string {
	if o == nil || IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnsibleAnsibleNamespaceMetadataResponse) GetEmailOk() (*string, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *AnsibleAnsibleNamespaceMetadataResponse) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *AnsibleAnsibleNamespaceMetadataResponse) SetEmail(v string) {
	o.Email = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AnsibleAnsibleNamespaceMetadataResponse) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnsibleAnsibleNamespaceMetadataResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AnsibleAnsibleNamespaceMetadataResponse) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AnsibleAnsibleNamespaceMetadataResponse) SetDescription(v string) {
	o.Description = &v
}

// GetResources returns the Resources field value if set, zero value otherwise.
func (o *AnsibleAnsibleNamespaceMetadataResponse) GetResources() string {
	if o == nil || IsNil(o.Resources) {
		var ret string
		return ret
	}
	return *o.Resources
}

// GetResourcesOk returns a tuple with the Resources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnsibleAnsibleNamespaceMetadataResponse) GetResourcesOk() (*string, bool) {
	if o == nil || IsNil(o.Resources) {
		return nil, false
	}
	return o.Resources, true
}

// HasResources returns a boolean if a field has been set.
func (o *AnsibleAnsibleNamespaceMetadataResponse) HasResources() bool {
	if o != nil && !IsNil(o.Resources) {
		return true
	}

	return false
}

// SetResources gets a reference to the given string and assigns it to the Resources field.
func (o *AnsibleAnsibleNamespaceMetadataResponse) SetResources(v string) {
	o.Resources = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *AnsibleAnsibleNamespaceMetadataResponse) GetLinks() []NamespaceLinkResponse {
	if o == nil || IsNil(o.Links) {
		var ret []NamespaceLinkResponse
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnsibleAnsibleNamespaceMetadataResponse) GetLinksOk() ([]NamespaceLinkResponse, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *AnsibleAnsibleNamespaceMetadataResponse) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []NamespaceLinkResponse and assigns it to the Links field.
func (o *AnsibleAnsibleNamespaceMetadataResponse) SetLinks(v []NamespaceLinkResponse) {
	o.Links = v
}

// GetAvatarSha256 returns the AvatarSha256 field value if set, zero value otherwise.
func (o *AnsibleAnsibleNamespaceMetadataResponse) GetAvatarSha256() string {
	if o == nil || IsNil(o.AvatarSha256) {
		var ret string
		return ret
	}
	return *o.AvatarSha256
}

// GetAvatarSha256Ok returns a tuple with the AvatarSha256 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnsibleAnsibleNamespaceMetadataResponse) GetAvatarSha256Ok() (*string, bool) {
	if o == nil || IsNil(o.AvatarSha256) {
		return nil, false
	}
	return o.AvatarSha256, true
}

// HasAvatarSha256 returns a boolean if a field has been set.
func (o *AnsibleAnsibleNamespaceMetadataResponse) HasAvatarSha256() bool {
	if o != nil && !IsNil(o.AvatarSha256) {
		return true
	}

	return false
}

// SetAvatarSha256 gets a reference to the given string and assigns it to the AvatarSha256 field.
func (o *AnsibleAnsibleNamespaceMetadataResponse) SetAvatarSha256(v string) {
	o.AvatarSha256 = &v
}

// GetAvatarUrl returns the AvatarUrl field value if set, zero value otherwise.
func (o *AnsibleAnsibleNamespaceMetadataResponse) GetAvatarUrl() string {
	if o == nil || IsNil(o.AvatarUrl) {
		var ret string
		return ret
	}
	return *o.AvatarUrl
}

// GetAvatarUrlOk returns a tuple with the AvatarUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnsibleAnsibleNamespaceMetadataResponse) GetAvatarUrlOk() (*string, bool) {
	if o == nil || IsNil(o.AvatarUrl) {
		return nil, false
	}
	return o.AvatarUrl, true
}

// HasAvatarUrl returns a boolean if a field has been set.
func (o *AnsibleAnsibleNamespaceMetadataResponse) HasAvatarUrl() bool {
	if o != nil && !IsNil(o.AvatarUrl) {
		return true
	}

	return false
}

// SetAvatarUrl gets a reference to the given string and assigns it to the AvatarUrl field.
func (o *AnsibleAnsibleNamespaceMetadataResponse) SetAvatarUrl(v string) {
	o.AvatarUrl = &v
}

// GetMetadataSha256 returns the MetadataSha256 field value if set, zero value otherwise.
func (o *AnsibleAnsibleNamespaceMetadataResponse) GetMetadataSha256() string {
	if o == nil || IsNil(o.MetadataSha256) {
		var ret string
		return ret
	}
	return *o.MetadataSha256
}

// GetMetadataSha256Ok returns a tuple with the MetadataSha256 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnsibleAnsibleNamespaceMetadataResponse) GetMetadataSha256Ok() (*string, bool) {
	if o == nil || IsNil(o.MetadataSha256) {
		return nil, false
	}
	return o.MetadataSha256, true
}

// HasMetadataSha256 returns a boolean if a field has been set.
func (o *AnsibleAnsibleNamespaceMetadataResponse) HasMetadataSha256() bool {
	if o != nil && !IsNil(o.MetadataSha256) {
		return true
	}

	return false
}

// SetMetadataSha256 gets a reference to the given string and assigns it to the MetadataSha256 field.
func (o *AnsibleAnsibleNamespaceMetadataResponse) SetMetadataSha256(v string) {
	o.MetadataSha256 = &v
}

func (o AnsibleAnsibleNamespaceMetadataResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AnsibleAnsibleNamespaceMetadataResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PulpHref) {
		toSerialize["pulp_href"] = o.PulpHref
	}
	toSerialize["name"] = o.Name
	if !IsNil(o.Company) {
		toSerialize["company"] = o.Company
	}
	if !IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Resources) {
		toSerialize["resources"] = o.Resources
	}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.AvatarSha256) {
		toSerialize["avatar_sha256"] = o.AvatarSha256
	}
	if !IsNil(o.AvatarUrl) {
		toSerialize["avatar_url"] = o.AvatarUrl
	}
	if !IsNil(o.MetadataSha256) {
		toSerialize["metadata_sha256"] = o.MetadataSha256
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AnsibleAnsibleNamespaceMetadataResponse) UnmarshalJSON(bytes []byte) (err error) {
	varAnsibleAnsibleNamespaceMetadataResponse := _AnsibleAnsibleNamespaceMetadataResponse{}

	if err = json.Unmarshal(bytes, &varAnsibleAnsibleNamespaceMetadataResponse); err == nil {
		*o = AnsibleAnsibleNamespaceMetadataResponse(varAnsibleAnsibleNamespaceMetadataResponse)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "pulp_href")
		delete(additionalProperties, "name")
		delete(additionalProperties, "company")
		delete(additionalProperties, "email")
		delete(additionalProperties, "description")
		delete(additionalProperties, "resources")
		delete(additionalProperties, "links")
		delete(additionalProperties, "avatar_sha256")
		delete(additionalProperties, "avatar_url")
		delete(additionalProperties, "metadata_sha256")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAnsibleAnsibleNamespaceMetadataResponse struct {
	value *AnsibleAnsibleNamespaceMetadataResponse
	isSet bool
}

func (v NullableAnsibleAnsibleNamespaceMetadataResponse) Get() *AnsibleAnsibleNamespaceMetadataResponse {
	return v.value
}

func (v *NullableAnsibleAnsibleNamespaceMetadataResponse) Set(val *AnsibleAnsibleNamespaceMetadataResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAnsibleAnsibleNamespaceMetadataResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAnsibleAnsibleNamespaceMetadataResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAnsibleAnsibleNamespaceMetadataResponse(val *AnsibleAnsibleNamespaceMetadataResponse) *NullableAnsibleAnsibleNamespaceMetadataResponse {
	return &NullableAnsibleAnsibleNamespaceMetadataResponse{value: val, isSet: true}
}

func (v NullableAnsibleAnsibleNamespaceMetadataResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAnsibleAnsibleNamespaceMetadataResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


