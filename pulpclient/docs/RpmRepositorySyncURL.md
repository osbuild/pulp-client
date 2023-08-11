# RpmRepositorySyncURL

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Remote** | Pointer to **string** | A remote to sync from. This will override a remote set on repository. | [optional] 
**Mirror** | Pointer to **NullableBool** | DEPRECATED: If &#x60;&#x60;True&#x60;&#x60;, &#x60;&#x60;sync_policy&#x60;&#x60; will default to &#39;mirror_complete&#39; instead of &#39;additive&#39;. | [optional] 
**SyncPolicy** | Pointer to [**NullableSyncPolicyEnum**](SyncPolicyEnum.md) |  | [optional] 
**SkipTypes** | Pointer to [**[]SkipTypesEnum**](SkipTypesEnum.md) | List of content types to skip during sync. | [optional] [default to []]
**Optimize** | Pointer to **bool** | Whether or not to optimize sync. | [optional] [default to true]

## Methods

### NewRpmRepositorySyncURL

`func NewRpmRepositorySyncURL() *RpmRepositorySyncURL`

NewRpmRepositorySyncURL instantiates a new RpmRepositorySyncURL object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRpmRepositorySyncURLWithDefaults

`func NewRpmRepositorySyncURLWithDefaults() *RpmRepositorySyncURL`

NewRpmRepositorySyncURLWithDefaults instantiates a new RpmRepositorySyncURL object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRemote

`func (o *RpmRepositorySyncURL) GetRemote() string`

GetRemote returns the Remote field if non-nil, zero value otherwise.

### GetRemoteOk

`func (o *RpmRepositorySyncURL) GetRemoteOk() (*string, bool)`

GetRemoteOk returns a tuple with the Remote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemote

`func (o *RpmRepositorySyncURL) SetRemote(v string)`

SetRemote sets Remote field to given value.

### HasRemote

`func (o *RpmRepositorySyncURL) HasRemote() bool`

HasRemote returns a boolean if a field has been set.

### GetMirror

`func (o *RpmRepositorySyncURL) GetMirror() bool`

GetMirror returns the Mirror field if non-nil, zero value otherwise.

### GetMirrorOk

`func (o *RpmRepositorySyncURL) GetMirrorOk() (*bool, bool)`

GetMirrorOk returns a tuple with the Mirror field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMirror

`func (o *RpmRepositorySyncURL) SetMirror(v bool)`

SetMirror sets Mirror field to given value.

### HasMirror

`func (o *RpmRepositorySyncURL) HasMirror() bool`

HasMirror returns a boolean if a field has been set.

### SetMirrorNil

`func (o *RpmRepositorySyncURL) SetMirrorNil(b bool)`

 SetMirrorNil sets the value for Mirror to be an explicit nil

### UnsetMirror
`func (o *RpmRepositorySyncURL) UnsetMirror()`

UnsetMirror ensures that no value is present for Mirror, not even an explicit nil
### GetSyncPolicy

`func (o *RpmRepositorySyncURL) GetSyncPolicy() SyncPolicyEnum`

GetSyncPolicy returns the SyncPolicy field if non-nil, zero value otherwise.

### GetSyncPolicyOk

`func (o *RpmRepositorySyncURL) GetSyncPolicyOk() (*SyncPolicyEnum, bool)`

GetSyncPolicyOk returns a tuple with the SyncPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncPolicy

`func (o *RpmRepositorySyncURL) SetSyncPolicy(v SyncPolicyEnum)`

SetSyncPolicy sets SyncPolicy field to given value.

### HasSyncPolicy

`func (o *RpmRepositorySyncURL) HasSyncPolicy() bool`

HasSyncPolicy returns a boolean if a field has been set.

### SetSyncPolicyNil

`func (o *RpmRepositorySyncURL) SetSyncPolicyNil(b bool)`

 SetSyncPolicyNil sets the value for SyncPolicy to be an explicit nil

### UnsetSyncPolicy
`func (o *RpmRepositorySyncURL) UnsetSyncPolicy()`

UnsetSyncPolicy ensures that no value is present for SyncPolicy, not even an explicit nil
### GetSkipTypes

`func (o *RpmRepositorySyncURL) GetSkipTypes() []SkipTypesEnum`

GetSkipTypes returns the SkipTypes field if non-nil, zero value otherwise.

### GetSkipTypesOk

`func (o *RpmRepositorySyncURL) GetSkipTypesOk() (*[]SkipTypesEnum, bool)`

GetSkipTypesOk returns a tuple with the SkipTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipTypes

`func (o *RpmRepositorySyncURL) SetSkipTypes(v []SkipTypesEnum)`

SetSkipTypes sets SkipTypes field to given value.

### HasSkipTypes

`func (o *RpmRepositorySyncURL) HasSkipTypes() bool`

HasSkipTypes returns a boolean if a field has been set.

### GetOptimize

`func (o *RpmRepositorySyncURL) GetOptimize() bool`

GetOptimize returns the Optimize field if non-nil, zero value otherwise.

### GetOptimizeOk

`func (o *RpmRepositorySyncURL) GetOptimizeOk() (*bool, bool)`

GetOptimizeOk returns a tuple with the Optimize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptimize

`func (o *RpmRepositorySyncURL) SetOptimize(v bool)`

SetOptimize sets Optimize field to given value.

### HasOptimize

`func (o *RpmRepositorySyncURL) HasOptimize() bool`

HasOptimize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


