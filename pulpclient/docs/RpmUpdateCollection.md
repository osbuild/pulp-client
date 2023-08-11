# RpmUpdateCollection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **NullableString** | Collection name. | 
**Shortname** | **NullableString** | Collection short name. | 
**Module** | **map[string]interface{}** | Collection modular NSVCA. | 

## Methods

### NewRpmUpdateCollection

`func NewRpmUpdateCollection(name NullableString, shortname NullableString, module map[string]interface{}, ) *RpmUpdateCollection`

NewRpmUpdateCollection instantiates a new RpmUpdateCollection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRpmUpdateCollectionWithDefaults

`func NewRpmUpdateCollectionWithDefaults() *RpmUpdateCollection`

NewRpmUpdateCollectionWithDefaults instantiates a new RpmUpdateCollection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *RpmUpdateCollection) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RpmUpdateCollection) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RpmUpdateCollection) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *RpmUpdateCollection) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *RpmUpdateCollection) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetShortname

`func (o *RpmUpdateCollection) GetShortname() string`

GetShortname returns the Shortname field if non-nil, zero value otherwise.

### GetShortnameOk

`func (o *RpmUpdateCollection) GetShortnameOk() (*string, bool)`

GetShortnameOk returns a tuple with the Shortname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortname

`func (o *RpmUpdateCollection) SetShortname(v string)`

SetShortname sets Shortname field to given value.


### SetShortnameNil

`func (o *RpmUpdateCollection) SetShortnameNil(b bool)`

 SetShortnameNil sets the value for Shortname to be an explicit nil

### UnsetShortname
`func (o *RpmUpdateCollection) UnsetShortname()`

UnsetShortname ensures that no value is present for Shortname, not even an explicit nil
### GetModule

`func (o *RpmUpdateCollection) GetModule() map[string]interface{}`

GetModule returns the Module field if non-nil, zero value otherwise.

### GetModuleOk

`func (o *RpmUpdateCollection) GetModuleOk() (*map[string]interface{}, bool)`

GetModuleOk returns a tuple with the Module field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModule

`func (o *RpmUpdateCollection) SetModule(v map[string]interface{})`

SetModule sets Module field to given value.


### SetModuleNil

`func (o *RpmUpdateCollection) SetModuleNil(b bool)`

 SetModuleNil sets the value for Module to be an explicit nil

### UnsetModule
`func (o *RpmUpdateCollection) UnsetModule()`

UnsetModule ensures that no value is present for Module, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


