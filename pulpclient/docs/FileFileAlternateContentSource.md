# FileFileAlternateContentSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of Alternate Content Source. | 
**LastRefreshed** | Pointer to **NullableTime** | Date of last refresh of AlternateContentSource. | [optional] 
**Paths** | Pointer to **[]string** | List of paths that will be appended to the Remote url when searching for content. | [optional] 
**Remote** | **string** | The remote to provide alternate content source. | 

## Methods

### NewFileFileAlternateContentSource

`func NewFileFileAlternateContentSource(name string, remote string, ) *FileFileAlternateContentSource`

NewFileFileAlternateContentSource instantiates a new FileFileAlternateContentSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFileFileAlternateContentSourceWithDefaults

`func NewFileFileAlternateContentSourceWithDefaults() *FileFileAlternateContentSource`

NewFileFileAlternateContentSourceWithDefaults instantiates a new FileFileAlternateContentSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *FileFileAlternateContentSource) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FileFileAlternateContentSource) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FileFileAlternateContentSource) SetName(v string)`

SetName sets Name field to given value.


### GetLastRefreshed

`func (o *FileFileAlternateContentSource) GetLastRefreshed() time.Time`

GetLastRefreshed returns the LastRefreshed field if non-nil, zero value otherwise.

### GetLastRefreshedOk

`func (o *FileFileAlternateContentSource) GetLastRefreshedOk() (*time.Time, bool)`

GetLastRefreshedOk returns a tuple with the LastRefreshed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastRefreshed

`func (o *FileFileAlternateContentSource) SetLastRefreshed(v time.Time)`

SetLastRefreshed sets LastRefreshed field to given value.

### HasLastRefreshed

`func (o *FileFileAlternateContentSource) HasLastRefreshed() bool`

HasLastRefreshed returns a boolean if a field has been set.

### SetLastRefreshedNil

`func (o *FileFileAlternateContentSource) SetLastRefreshedNil(b bool)`

 SetLastRefreshedNil sets the value for LastRefreshed to be an explicit nil

### UnsetLastRefreshed
`func (o *FileFileAlternateContentSource) UnsetLastRefreshed()`

UnsetLastRefreshed ensures that no value is present for LastRefreshed, not even an explicit nil
### GetPaths

`func (o *FileFileAlternateContentSource) GetPaths() []string`

GetPaths returns the Paths field if non-nil, zero value otherwise.

### GetPathsOk

`func (o *FileFileAlternateContentSource) GetPathsOk() (*[]string, bool)`

GetPathsOk returns a tuple with the Paths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaths

`func (o *FileFileAlternateContentSource) SetPaths(v []string)`

SetPaths sets Paths field to given value.

### HasPaths

`func (o *FileFileAlternateContentSource) HasPaths() bool`

HasPaths returns a boolean if a field has been set.

### GetRemote

`func (o *FileFileAlternateContentSource) GetRemote() string`

GetRemote returns the Remote field if non-nil, zero value otherwise.

### GetRemoteOk

`func (o *FileFileAlternateContentSource) GetRemoteOk() (*string, bool)`

GetRemoteOk returns a tuple with the Remote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemote

`func (o *FileFileAlternateContentSource) SetRemote(v string)`

SetRemote sets Remote field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


