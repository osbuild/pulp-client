# RpmRpmAlternateContentSourceResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PulpHref** | Pointer to **string** |  | [optional] [readonly] 
**PulpCreated** | Pointer to **time.Time** | Timestamp of creation. | [optional] [readonly] 
**Name** | **string** | Name of Alternate Content Source. | 
**LastRefreshed** | Pointer to **NullableTime** | Date of last refresh of AlternateContentSource. | [optional] 
**Paths** | Pointer to **[]string** | List of paths that will be appended to the Remote url when searching for content. | [optional] 
**Remote** | **string** | The remote to provide alternate content source. | 

## Methods

### NewRpmRpmAlternateContentSourceResponse

`func NewRpmRpmAlternateContentSourceResponse(name string, remote string, ) *RpmRpmAlternateContentSourceResponse`

NewRpmRpmAlternateContentSourceResponse instantiates a new RpmRpmAlternateContentSourceResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRpmRpmAlternateContentSourceResponseWithDefaults

`func NewRpmRpmAlternateContentSourceResponseWithDefaults() *RpmRpmAlternateContentSourceResponse`

NewRpmRpmAlternateContentSourceResponseWithDefaults instantiates a new RpmRpmAlternateContentSourceResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPulpHref

`func (o *RpmRpmAlternateContentSourceResponse) GetPulpHref() string`

GetPulpHref returns the PulpHref field if non-nil, zero value otherwise.

### GetPulpHrefOk

`func (o *RpmRpmAlternateContentSourceResponse) GetPulpHrefOk() (*string, bool)`

GetPulpHrefOk returns a tuple with the PulpHref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpHref

`func (o *RpmRpmAlternateContentSourceResponse) SetPulpHref(v string)`

SetPulpHref sets PulpHref field to given value.

### HasPulpHref

`func (o *RpmRpmAlternateContentSourceResponse) HasPulpHref() bool`

HasPulpHref returns a boolean if a field has been set.

### GetPulpCreated

`func (o *RpmRpmAlternateContentSourceResponse) GetPulpCreated() time.Time`

GetPulpCreated returns the PulpCreated field if non-nil, zero value otherwise.

### GetPulpCreatedOk

`func (o *RpmRpmAlternateContentSourceResponse) GetPulpCreatedOk() (*time.Time, bool)`

GetPulpCreatedOk returns a tuple with the PulpCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpCreated

`func (o *RpmRpmAlternateContentSourceResponse) SetPulpCreated(v time.Time)`

SetPulpCreated sets PulpCreated field to given value.

### HasPulpCreated

`func (o *RpmRpmAlternateContentSourceResponse) HasPulpCreated() bool`

HasPulpCreated returns a boolean if a field has been set.

### GetName

`func (o *RpmRpmAlternateContentSourceResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RpmRpmAlternateContentSourceResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RpmRpmAlternateContentSourceResponse) SetName(v string)`

SetName sets Name field to given value.


### GetLastRefreshed

`func (o *RpmRpmAlternateContentSourceResponse) GetLastRefreshed() time.Time`

GetLastRefreshed returns the LastRefreshed field if non-nil, zero value otherwise.

### GetLastRefreshedOk

`func (o *RpmRpmAlternateContentSourceResponse) GetLastRefreshedOk() (*time.Time, bool)`

GetLastRefreshedOk returns a tuple with the LastRefreshed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastRefreshed

`func (o *RpmRpmAlternateContentSourceResponse) SetLastRefreshed(v time.Time)`

SetLastRefreshed sets LastRefreshed field to given value.

### HasLastRefreshed

`func (o *RpmRpmAlternateContentSourceResponse) HasLastRefreshed() bool`

HasLastRefreshed returns a boolean if a field has been set.

### SetLastRefreshedNil

`func (o *RpmRpmAlternateContentSourceResponse) SetLastRefreshedNil(b bool)`

 SetLastRefreshedNil sets the value for LastRefreshed to be an explicit nil

### UnsetLastRefreshed
`func (o *RpmRpmAlternateContentSourceResponse) UnsetLastRefreshed()`

UnsetLastRefreshed ensures that no value is present for LastRefreshed, not even an explicit nil
### GetPaths

`func (o *RpmRpmAlternateContentSourceResponse) GetPaths() []string`

GetPaths returns the Paths field if non-nil, zero value otherwise.

### GetPathsOk

`func (o *RpmRpmAlternateContentSourceResponse) GetPathsOk() (*[]string, bool)`

GetPathsOk returns a tuple with the Paths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaths

`func (o *RpmRpmAlternateContentSourceResponse) SetPaths(v []string)`

SetPaths sets Paths field to given value.

### HasPaths

`func (o *RpmRpmAlternateContentSourceResponse) HasPaths() bool`

HasPaths returns a boolean if a field has been set.

### GetRemote

`func (o *RpmRpmAlternateContentSourceResponse) GetRemote() string`

GetRemote returns the Remote field if non-nil, zero value otherwise.

### GetRemoteOk

`func (o *RpmRpmAlternateContentSourceResponse) GetRemoteOk() (*string, bool)`

GetRemoteOk returns a tuple with the Remote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemote

`func (o *RpmRpmAlternateContentSourceResponse) SetRemote(v string)`

SetRemote sets Remote field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


