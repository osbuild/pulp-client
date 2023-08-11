# PatchedrpmRpmAlternateContentSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name of Alternate Content Source. | [optional] 
**LastRefreshed** | Pointer to **NullableTime** | Date of last refresh of AlternateContentSource. | [optional] 
**Paths** | Pointer to **[]string** | List of paths that will be appended to the Remote url when searching for content. | [optional] 
**Remote** | Pointer to **string** | The remote to provide alternate content source. | [optional] 

## Methods

### NewPatchedrpmRpmAlternateContentSource

`func NewPatchedrpmRpmAlternateContentSource() *PatchedrpmRpmAlternateContentSource`

NewPatchedrpmRpmAlternateContentSource instantiates a new PatchedrpmRpmAlternateContentSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedrpmRpmAlternateContentSourceWithDefaults

`func NewPatchedrpmRpmAlternateContentSourceWithDefaults() *PatchedrpmRpmAlternateContentSource`

NewPatchedrpmRpmAlternateContentSourceWithDefaults instantiates a new PatchedrpmRpmAlternateContentSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PatchedrpmRpmAlternateContentSource) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedrpmRpmAlternateContentSource) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedrpmRpmAlternateContentSource) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedrpmRpmAlternateContentSource) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLastRefreshed

`func (o *PatchedrpmRpmAlternateContentSource) GetLastRefreshed() time.Time`

GetLastRefreshed returns the LastRefreshed field if non-nil, zero value otherwise.

### GetLastRefreshedOk

`func (o *PatchedrpmRpmAlternateContentSource) GetLastRefreshedOk() (*time.Time, bool)`

GetLastRefreshedOk returns a tuple with the LastRefreshed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastRefreshed

`func (o *PatchedrpmRpmAlternateContentSource) SetLastRefreshed(v time.Time)`

SetLastRefreshed sets LastRefreshed field to given value.

### HasLastRefreshed

`func (o *PatchedrpmRpmAlternateContentSource) HasLastRefreshed() bool`

HasLastRefreshed returns a boolean if a field has been set.

### SetLastRefreshedNil

`func (o *PatchedrpmRpmAlternateContentSource) SetLastRefreshedNil(b bool)`

 SetLastRefreshedNil sets the value for LastRefreshed to be an explicit nil

### UnsetLastRefreshed
`func (o *PatchedrpmRpmAlternateContentSource) UnsetLastRefreshed()`

UnsetLastRefreshed ensures that no value is present for LastRefreshed, not even an explicit nil
### GetPaths

`func (o *PatchedrpmRpmAlternateContentSource) GetPaths() []string`

GetPaths returns the Paths field if non-nil, zero value otherwise.

### GetPathsOk

`func (o *PatchedrpmRpmAlternateContentSource) GetPathsOk() (*[]string, bool)`

GetPathsOk returns a tuple with the Paths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaths

`func (o *PatchedrpmRpmAlternateContentSource) SetPaths(v []string)`

SetPaths sets Paths field to given value.

### HasPaths

`func (o *PatchedrpmRpmAlternateContentSource) HasPaths() bool`

HasPaths returns a boolean if a field has been set.

### GetRemote

`func (o *PatchedrpmRpmAlternateContentSource) GetRemote() string`

GetRemote returns the Remote field if non-nil, zero value otherwise.

### GetRemoteOk

`func (o *PatchedrpmRpmAlternateContentSource) GetRemoteOk() (*string, bool)`

GetRemoteOk returns a tuple with the Remote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemote

`func (o *PatchedrpmRpmAlternateContentSource) SetRemote(v string)`

SetRemote sets Remote field to given value.

### HasRemote

`func (o *PatchedrpmRpmAlternateContentSource) HasRemote() bool`

HasRemote returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


