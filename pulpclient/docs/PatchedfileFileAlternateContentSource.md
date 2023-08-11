# PatchedfileFileAlternateContentSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name of Alternate Content Source. | [optional] 
**LastRefreshed** | Pointer to **NullableTime** | Date of last refresh of AlternateContentSource. | [optional] 
**Paths** | Pointer to **[]string** | List of paths that will be appended to the Remote url when searching for content. | [optional] 
**Remote** | Pointer to **string** | The remote to provide alternate content source. | [optional] 

## Methods

### NewPatchedfileFileAlternateContentSource

`func NewPatchedfileFileAlternateContentSource() *PatchedfileFileAlternateContentSource`

NewPatchedfileFileAlternateContentSource instantiates a new PatchedfileFileAlternateContentSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedfileFileAlternateContentSourceWithDefaults

`func NewPatchedfileFileAlternateContentSourceWithDefaults() *PatchedfileFileAlternateContentSource`

NewPatchedfileFileAlternateContentSourceWithDefaults instantiates a new PatchedfileFileAlternateContentSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PatchedfileFileAlternateContentSource) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedfileFileAlternateContentSource) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedfileFileAlternateContentSource) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedfileFileAlternateContentSource) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLastRefreshed

`func (o *PatchedfileFileAlternateContentSource) GetLastRefreshed() time.Time`

GetLastRefreshed returns the LastRefreshed field if non-nil, zero value otherwise.

### GetLastRefreshedOk

`func (o *PatchedfileFileAlternateContentSource) GetLastRefreshedOk() (*time.Time, bool)`

GetLastRefreshedOk returns a tuple with the LastRefreshed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastRefreshed

`func (o *PatchedfileFileAlternateContentSource) SetLastRefreshed(v time.Time)`

SetLastRefreshed sets LastRefreshed field to given value.

### HasLastRefreshed

`func (o *PatchedfileFileAlternateContentSource) HasLastRefreshed() bool`

HasLastRefreshed returns a boolean if a field has been set.

### SetLastRefreshedNil

`func (o *PatchedfileFileAlternateContentSource) SetLastRefreshedNil(b bool)`

 SetLastRefreshedNil sets the value for LastRefreshed to be an explicit nil

### UnsetLastRefreshed
`func (o *PatchedfileFileAlternateContentSource) UnsetLastRefreshed()`

UnsetLastRefreshed ensures that no value is present for LastRefreshed, not even an explicit nil
### GetPaths

`func (o *PatchedfileFileAlternateContentSource) GetPaths() []string`

GetPaths returns the Paths field if non-nil, zero value otherwise.

### GetPathsOk

`func (o *PatchedfileFileAlternateContentSource) GetPathsOk() (*[]string, bool)`

GetPathsOk returns a tuple with the Paths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaths

`func (o *PatchedfileFileAlternateContentSource) SetPaths(v []string)`

SetPaths sets Paths field to given value.

### HasPaths

`func (o *PatchedfileFileAlternateContentSource) HasPaths() bool`

HasPaths returns a boolean if a field has been set.

### GetRemote

`func (o *PatchedfileFileAlternateContentSource) GetRemote() string`

GetRemote returns the Remote field if non-nil, zero value otherwise.

### GetRemoteOk

`func (o *PatchedfileFileAlternateContentSource) GetRemoteOk() (*string, bool)`

GetRemoteOk returns a tuple with the Remote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemote

`func (o *PatchedfileFileAlternateContentSource) SetRemote(v string)`

SetRemote sets Remote field to given value.

### HasRemote

`func (o *PatchedfileFileAlternateContentSource) HasRemote() bool`

HasRemote returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


