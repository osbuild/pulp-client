# CompsXml

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**File** | ***os.File** | Full path of a comps.xml file that may be parsed into comps.xml Content units. | 
**Repository** | Pointer to **string** | URI of an RPM repository the comps.xml content units should be associated to. | [optional] 
**Replace** | Pointer to **bool** | If true, incoming comps.xml replaces existing comps-related ContentUnits in the specified repository. | [optional] 

## Methods

### NewCompsXml

`func NewCompsXml(file *os.File, ) *CompsXml`

NewCompsXml instantiates a new CompsXml object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCompsXmlWithDefaults

`func NewCompsXmlWithDefaults() *CompsXml`

NewCompsXmlWithDefaults instantiates a new CompsXml object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFile

`func (o *CompsXml) GetFile() *os.File`

GetFile returns the File field if non-nil, zero value otherwise.

### GetFileOk

`func (o *CompsXml) GetFileOk() (**os.File, bool)`

GetFileOk returns a tuple with the File field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFile

`func (o *CompsXml) SetFile(v *os.File)`

SetFile sets File field to given value.


### GetRepository

`func (o *CompsXml) GetRepository() string`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *CompsXml) GetRepositoryOk() (*string, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *CompsXml) SetRepository(v string)`

SetRepository sets Repository field to given value.

### HasRepository

`func (o *CompsXml) HasRepository() bool`

HasRepository returns a boolean if a field has been set.

### GetReplace

`func (o *CompsXml) GetReplace() bool`

GetReplace returns the Replace field if non-nil, zero value otherwise.

### GetReplaceOk

`func (o *CompsXml) GetReplaceOk() (*bool, bool)`

GetReplaceOk returns a tuple with the Replace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplace

`func (o *CompsXml) SetReplace(v bool)`

SetReplace sets Replace field to given value.

### HasReplace

`func (o *CompsXml) HasReplace() bool`

HasReplace returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


