# FileFilePublication

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RepositoryVersion** | Pointer to **string** |  | [optional] 
**Repository** | Pointer to **string** | A URI of the repository to be published. | [optional] 
**Manifest** | Pointer to **NullableString** | Filename to use for manifest file containing metadata for all the files. | [optional] [default to "PULP_MANIFEST"]

## Methods

### NewFileFilePublication

`func NewFileFilePublication() *FileFilePublication`

NewFileFilePublication instantiates a new FileFilePublication object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFileFilePublicationWithDefaults

`func NewFileFilePublicationWithDefaults() *FileFilePublication`

NewFileFilePublicationWithDefaults instantiates a new FileFilePublication object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRepositoryVersion

`func (o *FileFilePublication) GetRepositoryVersion() string`

GetRepositoryVersion returns the RepositoryVersion field if non-nil, zero value otherwise.

### GetRepositoryVersionOk

`func (o *FileFilePublication) GetRepositoryVersionOk() (*string, bool)`

GetRepositoryVersionOk returns a tuple with the RepositoryVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepositoryVersion

`func (o *FileFilePublication) SetRepositoryVersion(v string)`

SetRepositoryVersion sets RepositoryVersion field to given value.

### HasRepositoryVersion

`func (o *FileFilePublication) HasRepositoryVersion() bool`

HasRepositoryVersion returns a boolean if a field has been set.

### GetRepository

`func (o *FileFilePublication) GetRepository() string`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *FileFilePublication) GetRepositoryOk() (*string, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *FileFilePublication) SetRepository(v string)`

SetRepository sets Repository field to given value.

### HasRepository

`func (o *FileFilePublication) HasRepository() bool`

HasRepository returns a boolean if a field has been set.

### GetManifest

`func (o *FileFilePublication) GetManifest() string`

GetManifest returns the Manifest field if non-nil, zero value otherwise.

### GetManifestOk

`func (o *FileFilePublication) GetManifestOk() (*string, bool)`

GetManifestOk returns a tuple with the Manifest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManifest

`func (o *FileFilePublication) SetManifest(v string)`

SetManifest sets Manifest field to given value.

### HasManifest

`func (o *FileFilePublication) HasManifest() bool`

HasManifest returns a boolean if a field has been set.

### SetManifestNil

`func (o *FileFilePublication) SetManifestNil(b bool)`

 SetManifestNil sets the value for Manifest to be an explicit nil

### UnsetManifest
`func (o *FileFilePublication) UnsetManifest()`

UnsetManifest ensures that no value is present for Manifest, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


