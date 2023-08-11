# AnsibleCollectionVersionSignature

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Repository** | Pointer to **string** | A URI of a repository the new content unit should be associated with. | [optional] 
**File** | ***os.File** | An uploaded file that may be turned into the artifact of the content unit. | 
**SignedCollection** | **string** | The content this signature is pointing to. | 

## Methods

### NewAnsibleCollectionVersionSignature

`func NewAnsibleCollectionVersionSignature(file *os.File, signedCollection string, ) *AnsibleCollectionVersionSignature`

NewAnsibleCollectionVersionSignature instantiates a new AnsibleCollectionVersionSignature object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnsibleCollectionVersionSignatureWithDefaults

`func NewAnsibleCollectionVersionSignatureWithDefaults() *AnsibleCollectionVersionSignature`

NewAnsibleCollectionVersionSignatureWithDefaults instantiates a new AnsibleCollectionVersionSignature object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRepository

`func (o *AnsibleCollectionVersionSignature) GetRepository() string`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *AnsibleCollectionVersionSignature) GetRepositoryOk() (*string, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *AnsibleCollectionVersionSignature) SetRepository(v string)`

SetRepository sets Repository field to given value.

### HasRepository

`func (o *AnsibleCollectionVersionSignature) HasRepository() bool`

HasRepository returns a boolean if a field has been set.

### GetFile

`func (o *AnsibleCollectionVersionSignature) GetFile() *os.File`

GetFile returns the File field if non-nil, zero value otherwise.

### GetFileOk

`func (o *AnsibleCollectionVersionSignature) GetFileOk() (**os.File, bool)`

GetFileOk returns a tuple with the File field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFile

`func (o *AnsibleCollectionVersionSignature) SetFile(v *os.File)`

SetFile sets File field to given value.


### GetSignedCollection

`func (o *AnsibleCollectionVersionSignature) GetSignedCollection() string`

GetSignedCollection returns the SignedCollection field if non-nil, zero value otherwise.

### GetSignedCollectionOk

`func (o *AnsibleCollectionVersionSignature) GetSignedCollectionOk() (*string, bool)`

GetSignedCollectionOk returns a tuple with the SignedCollection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignedCollection

`func (o *AnsibleCollectionVersionSignature) SetSignedCollection(v string)`

SetSignedCollection sets SignedCollection field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


