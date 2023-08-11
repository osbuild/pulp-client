# RpmUpdateRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Repository** | Pointer to **string** | A URI of a repository the new content unit should be associated with. | [optional] 
**File** | Pointer to ***os.File** | An uploaded file that may be turned into the artifact of the content unit. | [optional] 

## Methods

### NewRpmUpdateRecord

`func NewRpmUpdateRecord() *RpmUpdateRecord`

NewRpmUpdateRecord instantiates a new RpmUpdateRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRpmUpdateRecordWithDefaults

`func NewRpmUpdateRecordWithDefaults() *RpmUpdateRecord`

NewRpmUpdateRecordWithDefaults instantiates a new RpmUpdateRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRepository

`func (o *RpmUpdateRecord) GetRepository() string`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *RpmUpdateRecord) GetRepositoryOk() (*string, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *RpmUpdateRecord) SetRepository(v string)`

SetRepository sets Repository field to given value.

### HasRepository

`func (o *RpmUpdateRecord) HasRepository() bool`

HasRepository returns a boolean if a field has been set.

### GetFile

`func (o *RpmUpdateRecord) GetFile() *os.File`

GetFile returns the File field if non-nil, zero value otherwise.

### GetFileOk

`func (o *RpmUpdateRecord) GetFileOk() (**os.File, bool)`

GetFileOk returns a tuple with the File field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFile

`func (o *RpmUpdateRecord) SetFile(v *os.File)`

SetFile sets File field to given value.

### HasFile

`func (o *RpmUpdateRecord) HasFile() bool`

HasFile returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


