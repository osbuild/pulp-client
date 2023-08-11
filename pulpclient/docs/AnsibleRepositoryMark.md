# AnsibleRepositoryMark

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContentUnits** | **[]interface{}** | List of collection version hrefs to mark, use * to mark all content in repository | 
**Value** | **string** | The string value of this mark. | 

## Methods

### NewAnsibleRepositoryMark

`func NewAnsibleRepositoryMark(contentUnits []interface{}, value string, ) *AnsibleRepositoryMark`

NewAnsibleRepositoryMark instantiates a new AnsibleRepositoryMark object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnsibleRepositoryMarkWithDefaults

`func NewAnsibleRepositoryMarkWithDefaults() *AnsibleRepositoryMark`

NewAnsibleRepositoryMarkWithDefaults instantiates a new AnsibleRepositoryMark object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContentUnits

`func (o *AnsibleRepositoryMark) GetContentUnits() []interface{}`

GetContentUnits returns the ContentUnits field if non-nil, zero value otherwise.

### GetContentUnitsOk

`func (o *AnsibleRepositoryMark) GetContentUnitsOk() (*[]interface{}, bool)`

GetContentUnitsOk returns a tuple with the ContentUnits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentUnits

`func (o *AnsibleRepositoryMark) SetContentUnits(v []interface{})`

SetContentUnits sets ContentUnits field to given value.


### GetValue

`func (o *AnsibleRepositoryMark) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *AnsibleRepositoryMark) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *AnsibleRepositoryMark) SetValue(v string)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


