# PackageUploadTaskResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Session** | **string** |  | 
**Task** | **string** |  | 
**TaskStartTime** | **time.Time** |  | 

## Methods

### NewPackageUploadTaskResponse

`func NewPackageUploadTaskResponse(session string, task string, taskStartTime time.Time, ) *PackageUploadTaskResponse`

NewPackageUploadTaskResponse instantiates a new PackageUploadTaskResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPackageUploadTaskResponseWithDefaults

`func NewPackageUploadTaskResponseWithDefaults() *PackageUploadTaskResponse`

NewPackageUploadTaskResponseWithDefaults instantiates a new PackageUploadTaskResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSession

`func (o *PackageUploadTaskResponse) GetSession() string`

GetSession returns the Session field if non-nil, zero value otherwise.

### GetSessionOk

`func (o *PackageUploadTaskResponse) GetSessionOk() (*string, bool)`

GetSessionOk returns a tuple with the Session field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSession

`func (o *PackageUploadTaskResponse) SetSession(v string)`

SetSession sets Session field to given value.


### GetTask

`func (o *PackageUploadTaskResponse) GetTask() string`

GetTask returns the Task field if non-nil, zero value otherwise.

### GetTaskOk

`func (o *PackageUploadTaskResponse) GetTaskOk() (*string, bool)`

GetTaskOk returns a tuple with the Task field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTask

`func (o *PackageUploadTaskResponse) SetTask(v string)`

SetTask sets Task field to given value.


### GetTaskStartTime

`func (o *PackageUploadTaskResponse) GetTaskStartTime() time.Time`

GetTaskStartTime returns the TaskStartTime field if non-nil, zero value otherwise.

### GetTaskStartTimeOk

`func (o *PackageUploadTaskResponse) GetTaskStartTimeOk() (*time.Time, bool)`

GetTaskStartTimeOk returns a tuple with the TaskStartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskStartTime

`func (o *PackageUploadTaskResponse) SetTaskStartTime(v time.Time)`

SetTaskStartTime sets TaskStartTime field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


