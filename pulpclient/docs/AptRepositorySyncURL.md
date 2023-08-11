# AptRepositorySyncURL

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Remote** | Pointer to **string** | A remote to sync from. This will override a remote set on repository. | [optional] 
**Mirror** | Pointer to **bool** | If &#x60;&#x60;True&#x60;&#x60;, synchronization will remove all content that is not present in the remote repository. If &#x60;&#x60;False&#x60;&#x60;, sync will be additive only. | [optional] [default to false]
**Optimize** | Pointer to **bool** | Using optimize sync, will skip the processing of metadata if the checksum has not changed since the last sync. This greately improves re-sync performance in such situations. If you feel the sync is missing something that has changed about the remote repository you are syncing, try using optimize&#x3D;False for a full re-sync. Consider opening an issue on why we should not optimize in your use case. | [optional] [default to true]

## Methods

### NewAptRepositorySyncURL

`func NewAptRepositorySyncURL() *AptRepositorySyncURL`

NewAptRepositorySyncURL instantiates a new AptRepositorySyncURL object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAptRepositorySyncURLWithDefaults

`func NewAptRepositorySyncURLWithDefaults() *AptRepositorySyncURL`

NewAptRepositorySyncURLWithDefaults instantiates a new AptRepositorySyncURL object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRemote

`func (o *AptRepositorySyncURL) GetRemote() string`

GetRemote returns the Remote field if non-nil, zero value otherwise.

### GetRemoteOk

`func (o *AptRepositorySyncURL) GetRemoteOk() (*string, bool)`

GetRemoteOk returns a tuple with the Remote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemote

`func (o *AptRepositorySyncURL) SetRemote(v string)`

SetRemote sets Remote field to given value.

### HasRemote

`func (o *AptRepositorySyncURL) HasRemote() bool`

HasRemote returns a boolean if a field has been set.

### GetMirror

`func (o *AptRepositorySyncURL) GetMirror() bool`

GetMirror returns the Mirror field if non-nil, zero value otherwise.

### GetMirrorOk

`func (o *AptRepositorySyncURL) GetMirrorOk() (*bool, bool)`

GetMirrorOk returns a tuple with the Mirror field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMirror

`func (o *AptRepositorySyncURL) SetMirror(v bool)`

SetMirror sets Mirror field to given value.

### HasMirror

`func (o *AptRepositorySyncURL) HasMirror() bool`

HasMirror returns a boolean if a field has been set.

### GetOptimize

`func (o *AptRepositorySyncURL) GetOptimize() bool`

GetOptimize returns the Optimize field if non-nil, zero value otherwise.

### GetOptimizeOk

`func (o *AptRepositorySyncURL) GetOptimizeOk() (*bool, bool)`

GetOptimizeOk returns a tuple with the Optimize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptimize

`func (o *AptRepositorySyncURL) SetOptimize(v bool)`

SetOptimize sets Optimize field to given value.

### HasOptimize

`func (o *AptRepositorySyncURL) HasOptimize() bool`

HasOptimize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


