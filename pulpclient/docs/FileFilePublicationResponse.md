# FileFilePublicationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PulpHref** | Pointer to **string** |  | [optional] [readonly] 
**PulpCreated** | Pointer to **time.Time** | Timestamp of creation. | [optional] [readonly] 
**RepositoryVersion** | Pointer to **string** |  | [optional] 
**Repository** | Pointer to **string** | A URI of the repository to be published. | [optional] 
**Distributions** | Pointer to **[]string** | This publication is currently hosted as defined by these distributions. | [optional] [readonly] 
**Manifest** | Pointer to **NullableString** | Filename to use for manifest file containing metadata for all the files. | [optional] [default to "PULP_MANIFEST"]

## Methods

### NewFileFilePublicationResponse

`func NewFileFilePublicationResponse() *FileFilePublicationResponse`

NewFileFilePublicationResponse instantiates a new FileFilePublicationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFileFilePublicationResponseWithDefaults

`func NewFileFilePublicationResponseWithDefaults() *FileFilePublicationResponse`

NewFileFilePublicationResponseWithDefaults instantiates a new FileFilePublicationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPulpHref

`func (o *FileFilePublicationResponse) GetPulpHref() string`

GetPulpHref returns the PulpHref field if non-nil, zero value otherwise.

### GetPulpHrefOk

`func (o *FileFilePublicationResponse) GetPulpHrefOk() (*string, bool)`

GetPulpHrefOk returns a tuple with the PulpHref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpHref

`func (o *FileFilePublicationResponse) SetPulpHref(v string)`

SetPulpHref sets PulpHref field to given value.

### HasPulpHref

`func (o *FileFilePublicationResponse) HasPulpHref() bool`

HasPulpHref returns a boolean if a field has been set.

### GetPulpCreated

`func (o *FileFilePublicationResponse) GetPulpCreated() time.Time`

GetPulpCreated returns the PulpCreated field if non-nil, zero value otherwise.

### GetPulpCreatedOk

`func (o *FileFilePublicationResponse) GetPulpCreatedOk() (*time.Time, bool)`

GetPulpCreatedOk returns a tuple with the PulpCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpCreated

`func (o *FileFilePublicationResponse) SetPulpCreated(v time.Time)`

SetPulpCreated sets PulpCreated field to given value.

### HasPulpCreated

`func (o *FileFilePublicationResponse) HasPulpCreated() bool`

HasPulpCreated returns a boolean if a field has been set.

### GetRepositoryVersion

`func (o *FileFilePublicationResponse) GetRepositoryVersion() string`

GetRepositoryVersion returns the RepositoryVersion field if non-nil, zero value otherwise.

### GetRepositoryVersionOk

`func (o *FileFilePublicationResponse) GetRepositoryVersionOk() (*string, bool)`

GetRepositoryVersionOk returns a tuple with the RepositoryVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepositoryVersion

`func (o *FileFilePublicationResponse) SetRepositoryVersion(v string)`

SetRepositoryVersion sets RepositoryVersion field to given value.

### HasRepositoryVersion

`func (o *FileFilePublicationResponse) HasRepositoryVersion() bool`

HasRepositoryVersion returns a boolean if a field has been set.

### GetRepository

`func (o *FileFilePublicationResponse) GetRepository() string`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *FileFilePublicationResponse) GetRepositoryOk() (*string, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *FileFilePublicationResponse) SetRepository(v string)`

SetRepository sets Repository field to given value.

### HasRepository

`func (o *FileFilePublicationResponse) HasRepository() bool`

HasRepository returns a boolean if a field has been set.

### GetDistributions

`func (o *FileFilePublicationResponse) GetDistributions() []string`

GetDistributions returns the Distributions field if non-nil, zero value otherwise.

### GetDistributionsOk

`func (o *FileFilePublicationResponse) GetDistributionsOk() (*[]string, bool)`

GetDistributionsOk returns a tuple with the Distributions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistributions

`func (o *FileFilePublicationResponse) SetDistributions(v []string)`

SetDistributions sets Distributions field to given value.

### HasDistributions

`func (o *FileFilePublicationResponse) HasDistributions() bool`

HasDistributions returns a boolean if a field has been set.

### GetManifest

`func (o *FileFilePublicationResponse) GetManifest() string`

GetManifest returns the Manifest field if non-nil, zero value otherwise.

### GetManifestOk

`func (o *FileFilePublicationResponse) GetManifestOk() (*string, bool)`

GetManifestOk returns a tuple with the Manifest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManifest

`func (o *FileFilePublicationResponse) SetManifest(v string)`

SetManifest sets Manifest field to given value.

### HasManifest

`func (o *FileFilePublicationResponse) HasManifest() bool`

HasManifest returns a boolean if a field has been set.

### SetManifestNil

`func (o *FileFilePublicationResponse) SetManifestNil(b bool)`

 SetManifestNil sets the value for Manifest to be an explicit nil

### UnsetManifest
`func (o *FileFilePublicationResponse) UnsetManifest()`

UnsetManifest ensures that no value is present for Manifest, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


