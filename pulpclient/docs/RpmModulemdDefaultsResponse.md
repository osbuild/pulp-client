# RpmModulemdDefaultsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PulpHref** | Pointer to **string** |  | [optional] [readonly] 
**PulpCreated** | Pointer to **time.Time** | Timestamp of creation. | [optional] [readonly] 
**Module** | **string** | Modulemd name. | 
**Stream** | **string** | Modulemd default stream. | 
**Profiles** | **map[string]interface{}** | Default profiles for modulemd streams. | 

## Methods

### NewRpmModulemdDefaultsResponse

`func NewRpmModulemdDefaultsResponse(module string, stream string, profiles map[string]interface{}, ) *RpmModulemdDefaultsResponse`

NewRpmModulemdDefaultsResponse instantiates a new RpmModulemdDefaultsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRpmModulemdDefaultsResponseWithDefaults

`func NewRpmModulemdDefaultsResponseWithDefaults() *RpmModulemdDefaultsResponse`

NewRpmModulemdDefaultsResponseWithDefaults instantiates a new RpmModulemdDefaultsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPulpHref

`func (o *RpmModulemdDefaultsResponse) GetPulpHref() string`

GetPulpHref returns the PulpHref field if non-nil, zero value otherwise.

### GetPulpHrefOk

`func (o *RpmModulemdDefaultsResponse) GetPulpHrefOk() (*string, bool)`

GetPulpHrefOk returns a tuple with the PulpHref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpHref

`func (o *RpmModulemdDefaultsResponse) SetPulpHref(v string)`

SetPulpHref sets PulpHref field to given value.

### HasPulpHref

`func (o *RpmModulemdDefaultsResponse) HasPulpHref() bool`

HasPulpHref returns a boolean if a field has been set.

### GetPulpCreated

`func (o *RpmModulemdDefaultsResponse) GetPulpCreated() time.Time`

GetPulpCreated returns the PulpCreated field if non-nil, zero value otherwise.

### GetPulpCreatedOk

`func (o *RpmModulemdDefaultsResponse) GetPulpCreatedOk() (*time.Time, bool)`

GetPulpCreatedOk returns a tuple with the PulpCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpCreated

`func (o *RpmModulemdDefaultsResponse) SetPulpCreated(v time.Time)`

SetPulpCreated sets PulpCreated field to given value.

### HasPulpCreated

`func (o *RpmModulemdDefaultsResponse) HasPulpCreated() bool`

HasPulpCreated returns a boolean if a field has been set.

### GetModule

`func (o *RpmModulemdDefaultsResponse) GetModule() string`

GetModule returns the Module field if non-nil, zero value otherwise.

### GetModuleOk

`func (o *RpmModulemdDefaultsResponse) GetModuleOk() (*string, bool)`

GetModuleOk returns a tuple with the Module field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModule

`func (o *RpmModulemdDefaultsResponse) SetModule(v string)`

SetModule sets Module field to given value.


### GetStream

`func (o *RpmModulemdDefaultsResponse) GetStream() string`

GetStream returns the Stream field if non-nil, zero value otherwise.

### GetStreamOk

`func (o *RpmModulemdDefaultsResponse) GetStreamOk() (*string, bool)`

GetStreamOk returns a tuple with the Stream field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStream

`func (o *RpmModulemdDefaultsResponse) SetStream(v string)`

SetStream sets Stream field to given value.


### GetProfiles

`func (o *RpmModulemdDefaultsResponse) GetProfiles() map[string]interface{}`

GetProfiles returns the Profiles field if non-nil, zero value otherwise.

### GetProfilesOk

`func (o *RpmModulemdDefaultsResponse) GetProfilesOk() (*map[string]interface{}, bool)`

GetProfilesOk returns a tuple with the Profiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfiles

`func (o *RpmModulemdDefaultsResponse) SetProfiles(v map[string]interface{})`

SetProfiles sets Profiles field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


