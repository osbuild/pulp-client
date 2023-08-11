# Task

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of task. | 
**LoggingCid** | **string** | The logging correlation id associated with this task | 

## Methods

### NewTask

`func NewTask(name string, loggingCid string, ) *Task`

NewTask instantiates a new Task object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskWithDefaults

`func NewTaskWithDefaults() *Task`

NewTaskWithDefaults instantiates a new Task object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *Task) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Task) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Task) SetName(v string)`

SetName sets Name field to given value.


### GetLoggingCid

`func (o *Task) GetLoggingCid() string`

GetLoggingCid returns the LoggingCid field if non-nil, zero value otherwise.

### GetLoggingCidOk

`func (o *Task) GetLoggingCidOk() (*string, bool)`

GetLoggingCidOk returns a tuple with the LoggingCid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoggingCid

`func (o *Task) SetLoggingCid(v string)`

SetLoggingCid sets LoggingCid field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


