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

// checks if the CollectionImportDetailResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CollectionImportDetailResponse{}

// CollectionImportDetailResponse A serializer for a CollectionImport detail view.
type CollectionImportDetailResponse struct {
	Id string `json:"id"`
	State string `json:"state"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	StartedAt time.Time `json:"started_at"`
	FinishedAt *time.Time `json:"finished_at,omitempty"`
	Error map[string]interface{} `json:"error,omitempty"`
	Messages map[string]interface{} `json:"messages"`
	AdditionalProperties map[string]interface{}
}

type _CollectionImportDetailResponse CollectionImportDetailResponse

// NewCollectionImportDetailResponse instantiates a new CollectionImportDetailResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCollectionImportDetailResponse(id string, state string, createdAt time.Time, updatedAt time.Time, startedAt time.Time, messages map[string]interface{}) *CollectionImportDetailResponse {
	this := CollectionImportDetailResponse{}
	this.Id = id
	this.State = state
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	this.StartedAt = startedAt
	this.Messages = messages
	return &this
}

// NewCollectionImportDetailResponseWithDefaults instantiates a new CollectionImportDetailResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCollectionImportDetailResponseWithDefaults() *CollectionImportDetailResponse {
	this := CollectionImportDetailResponse{}
	return &this
}

// GetId returns the Id field value
func (o *CollectionImportDetailResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *CollectionImportDetailResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *CollectionImportDetailResponse) SetId(v string) {
	o.Id = v
}

// GetState returns the State field value
func (o *CollectionImportDetailResponse) GetState() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *CollectionImportDetailResponse) GetStateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *CollectionImportDetailResponse) SetState(v string) {
	o.State = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *CollectionImportDetailResponse) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *CollectionImportDetailResponse) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *CollectionImportDetailResponse) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *CollectionImportDetailResponse) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *CollectionImportDetailResponse) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *CollectionImportDetailResponse) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

// GetStartedAt returns the StartedAt field value
func (o *CollectionImportDetailResponse) GetStartedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.StartedAt
}

// GetStartedAtOk returns a tuple with the StartedAt field value
// and a boolean to check if the value has been set.
func (o *CollectionImportDetailResponse) GetStartedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StartedAt, true
}

// SetStartedAt sets field value
func (o *CollectionImportDetailResponse) SetStartedAt(v time.Time) {
	o.StartedAt = v
}

// GetFinishedAt returns the FinishedAt field value if set, zero value otherwise.
func (o *CollectionImportDetailResponse) GetFinishedAt() time.Time {
	if o == nil || IsNil(o.FinishedAt) {
		var ret time.Time
		return ret
	}
	return *o.FinishedAt
}

// GetFinishedAtOk returns a tuple with the FinishedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionImportDetailResponse) GetFinishedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.FinishedAt) {
		return nil, false
	}
	return o.FinishedAt, true
}

// HasFinishedAt returns a boolean if a field has been set.
func (o *CollectionImportDetailResponse) HasFinishedAt() bool {
	if o != nil && !IsNil(o.FinishedAt) {
		return true
	}

	return false
}

// SetFinishedAt gets a reference to the given time.Time and assigns it to the FinishedAt field.
func (o *CollectionImportDetailResponse) SetFinishedAt(v time.Time) {
	o.FinishedAt = &v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *CollectionImportDetailResponse) GetError() map[string]interface{} {
	if o == nil || IsNil(o.Error) {
		var ret map[string]interface{}
		return ret
	}
	return o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionImportDetailResponse) GetErrorOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Error) {
		return map[string]interface{}{}, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *CollectionImportDetailResponse) HasError() bool {
	if o != nil && !IsNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given map[string]interface{} and assigns it to the Error field.
func (o *CollectionImportDetailResponse) SetError(v map[string]interface{}) {
	o.Error = v
}

// GetMessages returns the Messages field value
func (o *CollectionImportDetailResponse) GetMessages() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Messages
}

// GetMessagesOk returns a tuple with the Messages field value
// and a boolean to check if the value has been set.
func (o *CollectionImportDetailResponse) GetMessagesOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Messages, true
}

// SetMessages sets field value
func (o *CollectionImportDetailResponse) SetMessages(v map[string]interface{}) {
	o.Messages = v
}

func (o CollectionImportDetailResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CollectionImportDetailResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["state"] = o.State
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["updated_at"] = o.UpdatedAt
	toSerialize["started_at"] = o.StartedAt
	if !IsNil(o.FinishedAt) {
		toSerialize["finished_at"] = o.FinishedAt
	}
	if !IsNil(o.Error) {
		toSerialize["error"] = o.Error
	}
	toSerialize["messages"] = o.Messages

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CollectionImportDetailResponse) UnmarshalJSON(bytes []byte) (err error) {
	varCollectionImportDetailResponse := _CollectionImportDetailResponse{}

	if err = json.Unmarshal(bytes, &varCollectionImportDetailResponse); err == nil {
		*o = CollectionImportDetailResponse(varCollectionImportDetailResponse)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "state")
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "updated_at")
		delete(additionalProperties, "started_at")
		delete(additionalProperties, "finished_at")
		delete(additionalProperties, "error")
		delete(additionalProperties, "messages")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCollectionImportDetailResponse struct {
	value *CollectionImportDetailResponse
	isSet bool
}

func (v NullableCollectionImportDetailResponse) Get() *CollectionImportDetailResponse {
	return v.value
}

func (v *NullableCollectionImportDetailResponse) Set(val *CollectionImportDetailResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCollectionImportDetailResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCollectionImportDetailResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollectionImportDetailResponse(val *CollectionImportDetailResponse) *NullableCollectionImportDetailResponse {
	return &NullableCollectionImportDetailResponse{value: val, isSet: true}
}

func (v NullableCollectionImportDetailResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollectionImportDetailResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


