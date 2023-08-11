# AnsibleRepositorySyncURL

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Remote** | Pointer to **string** | A remote to sync from. This will override a remote set on repository. | [optional] 
**Mirror** | Pointer to **bool** | If &#x60;&#x60;True&#x60;&#x60;, synchronization will remove all content that is not present in the remote repository. If &#x60;&#x60;False&#x60;&#x60;, sync will be additive only. | [optional] [default to false]
**Optimize** | Pointer to **bool** | Whether to optimize sync or not. | [optional] [default to true]

## Methods

### NewAnsibleRepositorySyncURL

`func NewAnsibleRepositorySyncURL() *AnsibleRepositorySyncURL`

NewAnsibleRepositorySyncURL instantiates a new AnsibleRepositorySyncURL object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnsibleRepositorySyncURLWithDefaults

`func NewAnsibleRepositorySyncURLWithDefaults() *AnsibleRepositorySyncURL`

NewAnsibleRepositorySyncURLWithDefaults instantiates a new AnsibleRepositorySyncURL object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRemote

`func (o *AnsibleRepositorySyncURL) GetRemote() string`

GetRemote returns the Remote field if non-nil, zero value otherwise.

### GetRemoteOk

`func (o *AnsibleRepositorySyncURL) GetRemoteOk() (*string, bool)`

GetRemoteOk returns a tuple with the Remote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemote

`func (o *AnsibleRepositorySyncURL) SetRemote(v string)`

SetRemote sets Remote field to given value.

### HasRemote

`func (o *AnsibleRepositorySyncURL) HasRemote() bool`

HasRemote returns a boolean if a field has been set.

### GetMirror

`func (o *AnsibleRepositorySyncURL) GetMirror() bool`

GetMirror returns the Mirror field if non-nil, zero value otherwise.

### GetMirrorOk

`func (o *AnsibleRepositorySyncURL) GetMirrorOk() (*bool, bool)`

GetMirrorOk returns a tuple with the Mirror field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMirror

`func (o *AnsibleRepositorySyncURL) SetMirror(v bool)`

SetMirror sets Mirror field to given value.

### HasMirror

`func (o *AnsibleRepositorySyncURL) HasMirror() bool`

HasMirror returns a boolean if a field has been set.

### GetOptimize

`func (o *AnsibleRepositorySyncURL) GetOptimize() bool`

GetOptimize returns the Optimize field if non-nil, zero value otherwise.

### GetOptimizeOk

`func (o *AnsibleRepositorySyncURL) GetOptimizeOk() (*bool, bool)`

GetOptimizeOk returns a tuple with the Optimize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptimize

`func (o *AnsibleRepositorySyncURL) SetOptimize(v bool)`

SetOptimize sets Optimize field to given value.

### HasOptimize

`func (o *AnsibleRepositorySyncURL) HasOptimize() bool`

HasOptimize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


