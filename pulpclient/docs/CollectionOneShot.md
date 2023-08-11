# CollectionOneShot

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**File** | ***os.File** | The Collection tarball. | 
**Sha256** | Pointer to **string** | An optional sha256 checksum of the uploaded file. | [optional] 
**ExpectedNamespace** | Pointer to **string** | The expected &#39;namespace&#39; of the Collection to be verified against the metadata during import. | [optional] 
**ExpectedName** | Pointer to **string** | The expected &#39;name&#39; of the Collection to be verified against the metadata during import. | [optional] 
**ExpectedVersion** | Pointer to **string** | The expected version of the Collection to be verified against the metadata during import. | [optional] 

## Methods

### NewCollectionOneShot

`func NewCollectionOneShot(file *os.File, ) *CollectionOneShot`

NewCollectionOneShot instantiates a new CollectionOneShot object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCollectionOneShotWithDefaults

`func NewCollectionOneShotWithDefaults() *CollectionOneShot`

NewCollectionOneShotWithDefaults instantiates a new CollectionOneShot object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFile

`func (o *CollectionOneShot) GetFile() *os.File`

GetFile returns the File field if non-nil, zero value otherwise.

### GetFileOk

`func (o *CollectionOneShot) GetFileOk() (**os.File, bool)`

GetFileOk returns a tuple with the File field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFile

`func (o *CollectionOneShot) SetFile(v *os.File)`

SetFile sets File field to given value.


### GetSha256

`func (o *CollectionOneShot) GetSha256() string`

GetSha256 returns the Sha256 field if non-nil, zero value otherwise.

### GetSha256Ok

`func (o *CollectionOneShot) GetSha256Ok() (*string, bool)`

GetSha256Ok returns a tuple with the Sha256 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha256

`func (o *CollectionOneShot) SetSha256(v string)`

SetSha256 sets Sha256 field to given value.

### HasSha256

`func (o *CollectionOneShot) HasSha256() bool`

HasSha256 returns a boolean if a field has been set.

### GetExpectedNamespace

`func (o *CollectionOneShot) GetExpectedNamespace() string`

GetExpectedNamespace returns the ExpectedNamespace field if non-nil, zero value otherwise.

### GetExpectedNamespaceOk

`func (o *CollectionOneShot) GetExpectedNamespaceOk() (*string, bool)`

GetExpectedNamespaceOk returns a tuple with the ExpectedNamespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedNamespace

`func (o *CollectionOneShot) SetExpectedNamespace(v string)`

SetExpectedNamespace sets ExpectedNamespace field to given value.

### HasExpectedNamespace

`func (o *CollectionOneShot) HasExpectedNamespace() bool`

HasExpectedNamespace returns a boolean if a field has been set.

### GetExpectedName

`func (o *CollectionOneShot) GetExpectedName() string`

GetExpectedName returns the ExpectedName field if non-nil, zero value otherwise.

### GetExpectedNameOk

`func (o *CollectionOneShot) GetExpectedNameOk() (*string, bool)`

GetExpectedNameOk returns a tuple with the ExpectedName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedName

`func (o *CollectionOneShot) SetExpectedName(v string)`

SetExpectedName sets ExpectedName field to given value.

### HasExpectedName

`func (o *CollectionOneShot) HasExpectedName() bool`

HasExpectedName returns a boolean if a field has been set.

### GetExpectedVersion

`func (o *CollectionOneShot) GetExpectedVersion() string`

GetExpectedVersion returns the ExpectedVersion field if non-nil, zero value otherwise.

### GetExpectedVersionOk

`func (o *CollectionOneShot) GetExpectedVersionOk() (*string, bool)`

GetExpectedVersionOk returns a tuple with the ExpectedVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedVersion

`func (o *CollectionOneShot) SetExpectedVersion(v string)`

SetExpectedVersion sets ExpectedVersion field to given value.

### HasExpectedVersion

`func (o *CollectionOneShot) HasExpectedVersion() bool`

HasExpectedVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


