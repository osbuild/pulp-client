# RpmModulemdObsolete

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Repository** | Pointer to **string** | A URI of a repository the new content unit should be associated with. | [optional] 
**Modified** | **string** | Obsolete modified time. | 
**ModuleName** | **string** | Modulemd name. | 
**ModuleStream** | **string** | Modulemd&#39;s stream. | 
**Message** | **string** | Obsolete description. | 
**OverridePrevious** | **NullableString** | Reset previous obsoletes. | 
**ModuleContext** | **NullableString** | Modulemd&#39;s context. | 
**EolDate** | **NullableString** | End of Life date. | 
**ObsoletedByModuleName** | **NullableString** | Obsolete by module name. | 
**ObsoletedByModuleStream** | **NullableString** | Obsolete by module stream. | 
**Snippet** | **string** | Module Obsolete snippet. | 

## Methods

### NewRpmModulemdObsolete

`func NewRpmModulemdObsolete(modified string, moduleName string, moduleStream string, message string, overridePrevious NullableString, moduleContext NullableString, eolDate NullableString, obsoletedByModuleName NullableString, obsoletedByModuleStream NullableString, snippet string, ) *RpmModulemdObsolete`

NewRpmModulemdObsolete instantiates a new RpmModulemdObsolete object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRpmModulemdObsoleteWithDefaults

`func NewRpmModulemdObsoleteWithDefaults() *RpmModulemdObsolete`

NewRpmModulemdObsoleteWithDefaults instantiates a new RpmModulemdObsolete object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRepository

`func (o *RpmModulemdObsolete) GetRepository() string`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *RpmModulemdObsolete) GetRepositoryOk() (*string, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *RpmModulemdObsolete) SetRepository(v string)`

SetRepository sets Repository field to given value.

### HasRepository

`func (o *RpmModulemdObsolete) HasRepository() bool`

HasRepository returns a boolean if a field has been set.

### GetModified

`func (o *RpmModulemdObsolete) GetModified() string`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *RpmModulemdObsolete) GetModifiedOk() (*string, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *RpmModulemdObsolete) SetModified(v string)`

SetModified sets Modified field to given value.


### GetModuleName

`func (o *RpmModulemdObsolete) GetModuleName() string`

GetModuleName returns the ModuleName field if non-nil, zero value otherwise.

### GetModuleNameOk

`func (o *RpmModulemdObsolete) GetModuleNameOk() (*string, bool)`

GetModuleNameOk returns a tuple with the ModuleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModuleName

`func (o *RpmModulemdObsolete) SetModuleName(v string)`

SetModuleName sets ModuleName field to given value.


### GetModuleStream

`func (o *RpmModulemdObsolete) GetModuleStream() string`

GetModuleStream returns the ModuleStream field if non-nil, zero value otherwise.

### GetModuleStreamOk

`func (o *RpmModulemdObsolete) GetModuleStreamOk() (*string, bool)`

GetModuleStreamOk returns a tuple with the ModuleStream field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModuleStream

`func (o *RpmModulemdObsolete) SetModuleStream(v string)`

SetModuleStream sets ModuleStream field to given value.


### GetMessage

`func (o *RpmModulemdObsolete) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *RpmModulemdObsolete) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *RpmModulemdObsolete) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetOverridePrevious

`func (o *RpmModulemdObsolete) GetOverridePrevious() string`

GetOverridePrevious returns the OverridePrevious field if non-nil, zero value otherwise.

### GetOverridePreviousOk

`func (o *RpmModulemdObsolete) GetOverridePreviousOk() (*string, bool)`

GetOverridePreviousOk returns a tuple with the OverridePrevious field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverridePrevious

`func (o *RpmModulemdObsolete) SetOverridePrevious(v string)`

SetOverridePrevious sets OverridePrevious field to given value.


### SetOverridePreviousNil

`func (o *RpmModulemdObsolete) SetOverridePreviousNil(b bool)`

 SetOverridePreviousNil sets the value for OverridePrevious to be an explicit nil

### UnsetOverridePrevious
`func (o *RpmModulemdObsolete) UnsetOverridePrevious()`

UnsetOverridePrevious ensures that no value is present for OverridePrevious, not even an explicit nil
### GetModuleContext

`func (o *RpmModulemdObsolete) GetModuleContext() string`

GetModuleContext returns the ModuleContext field if non-nil, zero value otherwise.

### GetModuleContextOk

`func (o *RpmModulemdObsolete) GetModuleContextOk() (*string, bool)`

GetModuleContextOk returns a tuple with the ModuleContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModuleContext

`func (o *RpmModulemdObsolete) SetModuleContext(v string)`

SetModuleContext sets ModuleContext field to given value.


### SetModuleContextNil

`func (o *RpmModulemdObsolete) SetModuleContextNil(b bool)`

 SetModuleContextNil sets the value for ModuleContext to be an explicit nil

### UnsetModuleContext
`func (o *RpmModulemdObsolete) UnsetModuleContext()`

UnsetModuleContext ensures that no value is present for ModuleContext, not even an explicit nil
### GetEolDate

`func (o *RpmModulemdObsolete) GetEolDate() string`

GetEolDate returns the EolDate field if non-nil, zero value otherwise.

### GetEolDateOk

`func (o *RpmModulemdObsolete) GetEolDateOk() (*string, bool)`

GetEolDateOk returns a tuple with the EolDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEolDate

`func (o *RpmModulemdObsolete) SetEolDate(v string)`

SetEolDate sets EolDate field to given value.


### SetEolDateNil

`func (o *RpmModulemdObsolete) SetEolDateNil(b bool)`

 SetEolDateNil sets the value for EolDate to be an explicit nil

### UnsetEolDate
`func (o *RpmModulemdObsolete) UnsetEolDate()`

UnsetEolDate ensures that no value is present for EolDate, not even an explicit nil
### GetObsoletedByModuleName

`func (o *RpmModulemdObsolete) GetObsoletedByModuleName() string`

GetObsoletedByModuleName returns the ObsoletedByModuleName field if non-nil, zero value otherwise.

### GetObsoletedByModuleNameOk

`func (o *RpmModulemdObsolete) GetObsoletedByModuleNameOk() (*string, bool)`

GetObsoletedByModuleNameOk returns a tuple with the ObsoletedByModuleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObsoletedByModuleName

`func (o *RpmModulemdObsolete) SetObsoletedByModuleName(v string)`

SetObsoletedByModuleName sets ObsoletedByModuleName field to given value.


### SetObsoletedByModuleNameNil

`func (o *RpmModulemdObsolete) SetObsoletedByModuleNameNil(b bool)`

 SetObsoletedByModuleNameNil sets the value for ObsoletedByModuleName to be an explicit nil

### UnsetObsoletedByModuleName
`func (o *RpmModulemdObsolete) UnsetObsoletedByModuleName()`

UnsetObsoletedByModuleName ensures that no value is present for ObsoletedByModuleName, not even an explicit nil
### GetObsoletedByModuleStream

`func (o *RpmModulemdObsolete) GetObsoletedByModuleStream() string`

GetObsoletedByModuleStream returns the ObsoletedByModuleStream field if non-nil, zero value otherwise.

### GetObsoletedByModuleStreamOk

`func (o *RpmModulemdObsolete) GetObsoletedByModuleStreamOk() (*string, bool)`

GetObsoletedByModuleStreamOk returns a tuple with the ObsoletedByModuleStream field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObsoletedByModuleStream

`func (o *RpmModulemdObsolete) SetObsoletedByModuleStream(v string)`

SetObsoletedByModuleStream sets ObsoletedByModuleStream field to given value.


### SetObsoletedByModuleStreamNil

`func (o *RpmModulemdObsolete) SetObsoletedByModuleStreamNil(b bool)`

 SetObsoletedByModuleStreamNil sets the value for ObsoletedByModuleStream to be an explicit nil

### UnsetObsoletedByModuleStream
`func (o *RpmModulemdObsolete) UnsetObsoletedByModuleStream()`

UnsetObsoletedByModuleStream ensures that no value is present for ObsoletedByModuleStream, not even an explicit nil
### GetSnippet

`func (o *RpmModulemdObsolete) GetSnippet() string`

GetSnippet returns the Snippet field if non-nil, zero value otherwise.

### GetSnippetOk

`func (o *RpmModulemdObsolete) GetSnippetOk() (*string, bool)`

GetSnippetOk returns a tuple with the Snippet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnippet

`func (o *RpmModulemdObsolete) SetSnippet(v string)`

SetSnippet sets Snippet field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


