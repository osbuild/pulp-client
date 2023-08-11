# AnsibleCollectionVersionMarkResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PulpCreated** | Pointer to **time.Time** | Timestamp of creation. | [optional] [readonly] 
**PulpHref** | Pointer to **string** |  | [optional] [readonly] 
**MarkedCollection** | **string** | The content this mark is pointing to. | 
**Value** | **string** | The string value of this mark. | 

## Methods

### NewAnsibleCollectionVersionMarkResponse

`func NewAnsibleCollectionVersionMarkResponse(markedCollection string, value string, ) *AnsibleCollectionVersionMarkResponse`

NewAnsibleCollectionVersionMarkResponse instantiates a new AnsibleCollectionVersionMarkResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnsibleCollectionVersionMarkResponseWithDefaults

`func NewAnsibleCollectionVersionMarkResponseWithDefaults() *AnsibleCollectionVersionMarkResponse`

NewAnsibleCollectionVersionMarkResponseWithDefaults instantiates a new AnsibleCollectionVersionMarkResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPulpCreated

`func (o *AnsibleCollectionVersionMarkResponse) GetPulpCreated() time.Time`

GetPulpCreated returns the PulpCreated field if non-nil, zero value otherwise.

### GetPulpCreatedOk

`func (o *AnsibleCollectionVersionMarkResponse) GetPulpCreatedOk() (*time.Time, bool)`

GetPulpCreatedOk returns a tuple with the PulpCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpCreated

`func (o *AnsibleCollectionVersionMarkResponse) SetPulpCreated(v time.Time)`

SetPulpCreated sets PulpCreated field to given value.

### HasPulpCreated

`func (o *AnsibleCollectionVersionMarkResponse) HasPulpCreated() bool`

HasPulpCreated returns a boolean if a field has been set.

### GetPulpHref

`func (o *AnsibleCollectionVersionMarkResponse) GetPulpHref() string`

GetPulpHref returns the PulpHref field if non-nil, zero value otherwise.

### GetPulpHrefOk

`func (o *AnsibleCollectionVersionMarkResponse) GetPulpHrefOk() (*string, bool)`

GetPulpHrefOk returns a tuple with the PulpHref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpHref

`func (o *AnsibleCollectionVersionMarkResponse) SetPulpHref(v string)`

SetPulpHref sets PulpHref field to given value.

### HasPulpHref

`func (o *AnsibleCollectionVersionMarkResponse) HasPulpHref() bool`

HasPulpHref returns a boolean if a field has been set.

### GetMarkedCollection

`func (o *AnsibleCollectionVersionMarkResponse) GetMarkedCollection() string`

GetMarkedCollection returns the MarkedCollection field if non-nil, zero value otherwise.

### GetMarkedCollectionOk

`func (o *AnsibleCollectionVersionMarkResponse) GetMarkedCollectionOk() (*string, bool)`

GetMarkedCollectionOk returns a tuple with the MarkedCollection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarkedCollection

`func (o *AnsibleCollectionVersionMarkResponse) SetMarkedCollection(v string)`

SetMarkedCollection sets MarkedCollection field to given value.


### GetValue

`func (o *AnsibleCollectionVersionMarkResponse) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *AnsibleCollectionVersionMarkResponse) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *AnsibleCollectionVersionMarkResponse) SetValue(v string)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


