# FileFileRepository

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PulpLabels** | Pointer to **map[string]string** |  | [optional] 
**Name** | **string** | A unique name for this repository. | 
**Description** | Pointer to **NullableString** | An optional description. | [optional] 
**RetainRepoVersions** | Pointer to **NullableInt64** | Retain X versions of the repository. Default is null which retains all versions. | [optional] 
**Remote** | Pointer to **NullableString** | An optional remote to use by default when syncing. | [optional] 
**Autopublish** | Pointer to **bool** | Whether to automatically create publications for new repository versions, and update any distributions pointing to this repository. | [optional] [default to false]
**Manifest** | Pointer to **NullableString** | Filename to use for manifest file containing metadata for all the files. | [optional] [default to "PULP_MANIFEST"]

## Methods

### NewFileFileRepository

`func NewFileFileRepository(name string, ) *FileFileRepository`

NewFileFileRepository instantiates a new FileFileRepository object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFileFileRepositoryWithDefaults

`func NewFileFileRepositoryWithDefaults() *FileFileRepository`

NewFileFileRepositoryWithDefaults instantiates a new FileFileRepository object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPulpLabels

`func (o *FileFileRepository) GetPulpLabels() map[string]string`

GetPulpLabels returns the PulpLabels field if non-nil, zero value otherwise.

### GetPulpLabelsOk

`func (o *FileFileRepository) GetPulpLabelsOk() (*map[string]string, bool)`

GetPulpLabelsOk returns a tuple with the PulpLabels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpLabels

`func (o *FileFileRepository) SetPulpLabels(v map[string]string)`

SetPulpLabels sets PulpLabels field to given value.

### HasPulpLabels

`func (o *FileFileRepository) HasPulpLabels() bool`

HasPulpLabels returns a boolean if a field has been set.

### GetName

`func (o *FileFileRepository) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FileFileRepository) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FileFileRepository) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *FileFileRepository) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FileFileRepository) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FileFileRepository) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *FileFileRepository) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *FileFileRepository) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *FileFileRepository) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetRetainRepoVersions

`func (o *FileFileRepository) GetRetainRepoVersions() int64`

GetRetainRepoVersions returns the RetainRepoVersions field if non-nil, zero value otherwise.

### GetRetainRepoVersionsOk

`func (o *FileFileRepository) GetRetainRepoVersionsOk() (*int64, bool)`

GetRetainRepoVersionsOk returns a tuple with the RetainRepoVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetainRepoVersions

`func (o *FileFileRepository) SetRetainRepoVersions(v int64)`

SetRetainRepoVersions sets RetainRepoVersions field to given value.

### HasRetainRepoVersions

`func (o *FileFileRepository) HasRetainRepoVersions() bool`

HasRetainRepoVersions returns a boolean if a field has been set.

### SetRetainRepoVersionsNil

`func (o *FileFileRepository) SetRetainRepoVersionsNil(b bool)`

 SetRetainRepoVersionsNil sets the value for RetainRepoVersions to be an explicit nil

### UnsetRetainRepoVersions
`func (o *FileFileRepository) UnsetRetainRepoVersions()`

UnsetRetainRepoVersions ensures that no value is present for RetainRepoVersions, not even an explicit nil
### GetRemote

`func (o *FileFileRepository) GetRemote() string`

GetRemote returns the Remote field if non-nil, zero value otherwise.

### GetRemoteOk

`func (o *FileFileRepository) GetRemoteOk() (*string, bool)`

GetRemoteOk returns a tuple with the Remote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemote

`func (o *FileFileRepository) SetRemote(v string)`

SetRemote sets Remote field to given value.

### HasRemote

`func (o *FileFileRepository) HasRemote() bool`

HasRemote returns a boolean if a field has been set.

### SetRemoteNil

`func (o *FileFileRepository) SetRemoteNil(b bool)`

 SetRemoteNil sets the value for Remote to be an explicit nil

### UnsetRemote
`func (o *FileFileRepository) UnsetRemote()`

UnsetRemote ensures that no value is present for Remote, not even an explicit nil
### GetAutopublish

`func (o *FileFileRepository) GetAutopublish() bool`

GetAutopublish returns the Autopublish field if non-nil, zero value otherwise.

### GetAutopublishOk

`func (o *FileFileRepository) GetAutopublishOk() (*bool, bool)`

GetAutopublishOk returns a tuple with the Autopublish field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutopublish

`func (o *FileFileRepository) SetAutopublish(v bool)`

SetAutopublish sets Autopublish field to given value.

### HasAutopublish

`func (o *FileFileRepository) HasAutopublish() bool`

HasAutopublish returns a boolean if a field has been set.

### GetManifest

`func (o *FileFileRepository) GetManifest() string`

GetManifest returns the Manifest field if non-nil, zero value otherwise.

### GetManifestOk

`func (o *FileFileRepository) GetManifestOk() (*string, bool)`

GetManifestOk returns a tuple with the Manifest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManifest

`func (o *FileFileRepository) SetManifest(v string)`

SetManifest sets Manifest field to given value.

### HasManifest

`func (o *FileFileRepository) HasManifest() bool`

HasManifest returns a boolean if a field has been set.

### SetManifestNil

`func (o *FileFileRepository) SetManifestNil(b bool)`

 SetManifestNil sets the value for Manifest to be an explicit nil

### UnsetManifest
`func (o *FileFileRepository) UnsetManifest()`

UnsetManifest ensures that no value is present for Manifest, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


