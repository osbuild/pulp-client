# RpmModulemdObsoleteResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PulpHref** | Pointer to **string** |  | [optional] [readonly] 
**PulpCreated** | Pointer to **time.Time** | Timestamp of creation. | [optional] [readonly] 
**Modified** | **string** | Obsolete modified time. | 
**ModuleName** | **string** | Modulemd name. | 
**ModuleStream** | **string** | Modulemd&#39;s stream. | 
**Message** | **string** | Obsolete description. | 
**OverridePrevious** | **NullableString** | Reset previous obsoletes. | 
**ModuleContext** | **NullableString** | Modulemd&#39;s context. | 
**EolDate** | **NullableString** | End of Life date. | 
**ObsoletedByModuleName** | **NullableString** | Obsolete by module name. | 
**ObsoletedByModuleStream** | **NullableString** | Obsolete by module stream. | 

## Methods

### NewRpmModulemdObsoleteResponse

`func NewRpmModulemdObsoleteResponse(modified string, moduleName string, moduleStream string, message string, overridePrevious NullableString, moduleContext NullableString, eolDate NullableString, obsoletedByModuleName NullableString, obsoletedByModuleStream NullableString, ) *RpmModulemdObsoleteResponse`

NewRpmModulemdObsoleteResponse instantiates a new RpmModulemdObsoleteResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRpmModulemdObsoleteResponseWithDefaults

`func NewRpmModulemdObsoleteResponseWithDefaults() *RpmModulemdObsoleteResponse`

NewRpmModulemdObsoleteResponseWithDefaults instantiates a new RpmModulemdObsoleteResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPulpHref

`func (o *RpmModulemdObsoleteResponse) GetPulpHref() string`

GetPulpHref returns the PulpHref field if non-nil, zero value otherwise.

### GetPulpHrefOk

`func (o *RpmModulemdObsoleteResponse) GetPulpHrefOk() (*string, bool)`

GetPulpHrefOk returns a tuple with the PulpHref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpHref

`func (o *RpmModulemdObsoleteResponse) SetPulpHref(v string)`

SetPulpHref sets PulpHref field to given value.

### HasPulpHref

`func (o *RpmModulemdObsoleteResponse) HasPulpHref() bool`

HasPulpHref returns a boolean if a field has been set.

### GetPulpCreated

`func (o *RpmModulemdObsoleteResponse) GetPulpCreated() time.Time`

GetPulpCreated returns the PulpCreated field if non-nil, zero value otherwise.

### GetPulpCreatedOk

`func (o *RpmModulemdObsoleteResponse) GetPulpCreatedOk() (*time.Time, bool)`

GetPulpCreatedOk returns a tuple with the PulpCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpCreated

`func (o *RpmModulemdObsoleteResponse) SetPulpCreated(v time.Time)`

SetPulpCreated sets PulpCreated field to given value.

### HasPulpCreated

`func (o *RpmModulemdObsoleteResponse) HasPulpCreated() bool`

HasPulpCreated returns a boolean if a field has been set.

### GetModified

`func (o *RpmModulemdObsoleteResponse) GetModified() string`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *RpmModulemdObsoleteResponse) GetModifiedOk() (*string, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *RpmModulemdObsoleteResponse) SetModified(v string)`

SetModified sets Modified field to given value.


### GetModuleName

`func (o *RpmModulemdObsoleteResponse) GetModuleName() string`

GetModuleName returns the ModuleName field if non-nil, zero value otherwise.

### GetModuleNameOk

`func (o *RpmModulemdObsoleteResponse) GetModuleNameOk() (*string, bool)`

GetModuleNameOk returns a tuple with the ModuleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModuleName

`func (o *RpmModulemdObsoleteResponse) SetModuleName(v string)`

SetModuleName sets ModuleName field to given value.


### GetModuleStream

`func (o *RpmModulemdObsoleteResponse) GetModuleStream() string`

GetModuleStream returns the ModuleStream field if non-nil, zero value otherwise.

### GetModuleStreamOk

`func (o *RpmModulemdObsoleteResponse) GetModuleStreamOk() (*string, bool)`

GetModuleStreamOk returns a tuple with the ModuleStream field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModuleStream

`func (o *RpmModulemdObsoleteResponse) SetModuleStream(v string)`

SetModuleStream sets ModuleStream field to given value.


### GetMessage

`func (o *RpmModulemdObsoleteResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *RpmModulemdObsoleteResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *RpmModulemdObsoleteResponse) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetOverridePrevious

`func (o *RpmModulemdObsoleteResponse) GetOverridePrevious() string`

GetOverridePrevious returns the OverridePrevious field if non-nil, zero value otherwise.

### GetOverridePreviousOk

`func (o *RpmModulemdObsoleteResponse) GetOverridePreviousOk() (*string, bool)`

GetOverridePreviousOk returns a tuple with the OverridePrevious field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverridePrevious

`func (o *RpmModulemdObsoleteResponse) SetOverridePrevious(v string)`

SetOverridePrevious sets OverridePrevious field to given value.


### SetOverridePreviousNil

`func (o *RpmModulemdObsoleteResponse) SetOverridePreviousNil(b bool)`

 SetOverridePreviousNil sets the value for OverridePrevious to be an explicit nil

### UnsetOverridePrevious
`func (o *RpmModulemdObsoleteResponse) UnsetOverridePrevious()`

UnsetOverridePrevious ensures that no value is present for OverridePrevious, not even an explicit nil
### GetModuleContext

`func (o *RpmModulemdObsoleteResponse) GetModuleContext() string`

GetModuleContext returns the ModuleContext field if non-nil, zero value otherwise.

### GetModuleContextOk

`func (o *RpmModulemdObsoleteResponse) GetModuleContextOk() (*string, bool)`

GetModuleContextOk returns a tuple with the ModuleContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModuleContext

`func (o *RpmModulemdObsoleteResponse) SetModuleContext(v string)`

SetModuleContext sets ModuleContext field to given value.


### SetModuleContextNil

`func (o *RpmModulemdObsoleteResponse) SetModuleContextNil(b bool)`

 SetModuleContextNil sets the value for ModuleContext to be an explicit nil

### UnsetModuleContext
`func (o *RpmModulemdObsoleteResponse) UnsetModuleContext()`

UnsetModuleContext ensures that no value is present for ModuleContext, not even an explicit nil
### GetEolDate

`func (o *RpmModulemdObsoleteResponse) GetEolDate() string`

GetEolDate returns the EolDate field if non-nil, zero value otherwise.

### GetEolDateOk

`func (o *RpmModulemdObsoleteResponse) GetEolDateOk() (*string, bool)`

GetEolDateOk returns a tuple with the EolDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEolDate

`func (o *RpmModulemdObsoleteResponse) SetEolDate(v string)`

SetEolDate sets EolDate field to given value.


### SetEolDateNil

`func (o *RpmModulemdObsoleteResponse) SetEolDateNil(b bool)`

 SetEolDateNil sets the value for EolDate to be an explicit nil

### UnsetEolDate
`func (o *RpmModulemdObsoleteResponse) UnsetEolDate()`

UnsetEolDate ensures that no value is present for EolDate, not even an explicit nil
### GetObsoletedByModuleName

`func (o *RpmModulemdObsoleteResponse) GetObsoletedByModuleName() string`

GetObsoletedByModuleName returns the ObsoletedByModuleName field if non-nil, zero value otherwise.

### GetObsoletedByModuleNameOk

`func (o *RpmModulemdObsoleteResponse) GetObsoletedByModuleNameOk() (*string, bool)`

GetObsoletedByModuleNameOk returns a tuple with the ObsoletedByModuleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObsoletedByModuleName

`func (o *RpmModulemdObsoleteResponse) SetObsoletedByModuleName(v string)`

SetObsoletedByModuleName sets ObsoletedByModuleName field to given value.


### SetObsoletedByModuleNameNil

`func (o *RpmModulemdObsoleteResponse) SetObsoletedByModuleNameNil(b bool)`

 SetObsoletedByModuleNameNil sets the value for ObsoletedByModuleName to be an explicit nil

### UnsetObsoletedByModuleName
`func (o *RpmModulemdObsoleteResponse) UnsetObsoletedByModuleName()`

UnsetObsoletedByModuleName ensures that no value is present for ObsoletedByModuleName, not even an explicit nil
### GetObsoletedByModuleStream

`func (o *RpmModulemdObsoleteResponse) GetObsoletedByModuleStream() string`

GetObsoletedByModuleStream returns the ObsoletedByModuleStream field if non-nil, zero value otherwise.

### GetObsoletedByModuleStreamOk

`func (o *RpmModulemdObsoleteResponse) GetObsoletedByModuleStreamOk() (*string, bool)`

GetObsoletedByModuleStreamOk returns a tuple with the ObsoletedByModuleStream field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObsoletedByModuleStream

`func (o *RpmModulemdObsoleteResponse) SetObsoletedByModuleStream(v string)`

SetObsoletedByModuleStream sets ObsoletedByModuleStream field to given value.


### SetObsoletedByModuleStreamNil

`func (o *RpmModulemdObsoleteResponse) SetObsoletedByModuleStreamNil(b bool)`

 SetObsoletedByModuleStreamNil sets the value for ObsoletedByModuleStream to be an explicit nil

### UnsetObsoletedByModuleStream
`func (o *RpmModulemdObsoleteResponse) UnsetObsoletedByModuleStream()`

UnsetObsoletedByModuleStream ensures that no value is present for ObsoletedByModuleStream, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


