# FileFileRepositoryResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PulpHref** | Pointer to **string** |  | [optional] [readonly] 
**PulpCreated** | Pointer to **time.Time** | Timestamp of creation. | [optional] [readonly] 
**VersionsHref** | Pointer to **string** |  | [optional] [readonly] 
**PulpLabels** | Pointer to **map[string]string** |  | [optional] 
**LatestVersionHref** | Pointer to **string** |  | [optional] [readonly] 
**Name** | **string** | A unique name for this repository. | 
**Description** | Pointer to **NullableString** | An optional description. | [optional] 
**RetainRepoVersions** | Pointer to **NullableInt64** | Retain X versions of the repository. Default is null which retains all versions. | [optional] 
**Remote** | Pointer to **NullableString** | An optional remote to use by default when syncing. | [optional] 
**Autopublish** | Pointer to **bool** | Whether to automatically create publications for new repository versions, and update any distributions pointing to this repository. | [optional] [default to false]
**Manifest** | Pointer to **NullableString** | Filename to use for manifest file containing metadata for all the files. | [optional] [default to "PULP_MANIFEST"]

## Methods

### NewFileFileRepositoryResponse

`func NewFileFileRepositoryResponse(name string, ) *FileFileRepositoryResponse`

NewFileFileRepositoryResponse instantiates a new FileFileRepositoryResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFileFileRepositoryResponseWithDefaults

`func NewFileFileRepositoryResponseWithDefaults() *FileFileRepositoryResponse`

NewFileFileRepositoryResponseWithDefaults instantiates a new FileFileRepositoryResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPulpHref

`func (o *FileFileRepositoryResponse) GetPulpHref() string`

GetPulpHref returns the PulpHref field if non-nil, zero value otherwise.

### GetPulpHrefOk

`func (o *FileFileRepositoryResponse) GetPulpHrefOk() (*string, bool)`

GetPulpHrefOk returns a tuple with the PulpHref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpHref

`func (o *FileFileRepositoryResponse) SetPulpHref(v string)`

SetPulpHref sets PulpHref field to given value.

### HasPulpHref

`func (o *FileFileRepositoryResponse) HasPulpHref() bool`

HasPulpHref returns a boolean if a field has been set.

### GetPulpCreated

`func (o *FileFileRepositoryResponse) GetPulpCreated() time.Time`

GetPulpCreated returns the PulpCreated field if non-nil, zero value otherwise.

### GetPulpCreatedOk

`func (o *FileFileRepositoryResponse) GetPulpCreatedOk() (*time.Time, bool)`

GetPulpCreatedOk returns a tuple with the PulpCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpCreated

`func (o *FileFileRepositoryResponse) SetPulpCreated(v time.Time)`

SetPulpCreated sets PulpCreated field to given value.

### HasPulpCreated

`func (o *FileFileRepositoryResponse) HasPulpCreated() bool`

HasPulpCreated returns a boolean if a field has been set.

### GetVersionsHref

`func (o *FileFileRepositoryResponse) GetVersionsHref() string`

GetVersionsHref returns the VersionsHref field if non-nil, zero value otherwise.

### GetVersionsHrefOk

`func (o *FileFileRepositoryResponse) GetVersionsHrefOk() (*string, bool)`

GetVersionsHrefOk returns a tuple with the VersionsHref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionsHref

`func (o *FileFileRepositoryResponse) SetVersionsHref(v string)`

SetVersionsHref sets VersionsHref field to given value.

### HasVersionsHref

`func (o *FileFileRepositoryResponse) HasVersionsHref() bool`

HasVersionsHref returns a boolean if a field has been set.

### GetPulpLabels

`func (o *FileFileRepositoryResponse) GetPulpLabels() map[string]string`

GetPulpLabels returns the PulpLabels field if non-nil, zero value otherwise.

### GetPulpLabelsOk

`func (o *FileFileRepositoryResponse) GetPulpLabelsOk() (*map[string]string, bool)`

GetPulpLabelsOk returns a tuple with the PulpLabels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpLabels

`func (o *FileFileRepositoryResponse) SetPulpLabels(v map[string]string)`

SetPulpLabels sets PulpLabels field to given value.

### HasPulpLabels

`func (o *FileFileRepositoryResponse) HasPulpLabels() bool`

HasPulpLabels returns a boolean if a field has been set.

### GetLatestVersionHref

`func (o *FileFileRepositoryResponse) GetLatestVersionHref() string`

GetLatestVersionHref returns the LatestVersionHref field if non-nil, zero value otherwise.

### GetLatestVersionHrefOk

`func (o *FileFileRepositoryResponse) GetLatestVersionHrefOk() (*string, bool)`

GetLatestVersionHrefOk returns a tuple with the LatestVersionHref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestVersionHref

`func (o *FileFileRepositoryResponse) SetLatestVersionHref(v string)`

SetLatestVersionHref sets LatestVersionHref field to given value.

### HasLatestVersionHref

`func (o *FileFileRepositoryResponse) HasLatestVersionHref() bool`

HasLatestVersionHref returns a boolean if a field has been set.

### GetName

`func (o *FileFileRepositoryResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FileFileRepositoryResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FileFileRepositoryResponse) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *FileFileRepositoryResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FileFileRepositoryResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FileFileRepositoryResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *FileFileRepositoryResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *FileFileRepositoryResponse) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *FileFileRepositoryResponse) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetRetainRepoVersions

`func (o *FileFileRepositoryResponse) GetRetainRepoVersions() int64`

GetRetainRepoVersions returns the RetainRepoVersions field if non-nil, zero value otherwise.

### GetRetainRepoVersionsOk

`func (o *FileFileRepositoryResponse) GetRetainRepoVersionsOk() (*int64, bool)`

GetRetainRepoVersionsOk returns a tuple with the RetainRepoVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetainRepoVersions

`func (o *FileFileRepositoryResponse) SetRetainRepoVersions(v int64)`

SetRetainRepoVersions sets RetainRepoVersions field to given value.

### HasRetainRepoVersions

`func (o *FileFileRepositoryResponse) HasRetainRepoVersions() bool`

HasRetainRepoVersions returns a boolean if a field has been set.

### SetRetainRepoVersionsNil

`func (o *FileFileRepositoryResponse) SetRetainRepoVersionsNil(b bool)`

 SetRetainRepoVersionsNil sets the value for RetainRepoVersions to be an explicit nil

### UnsetRetainRepoVersions
`func (o *FileFileRepositoryResponse) UnsetRetainRepoVersions()`

UnsetRetainRepoVersions ensures that no value is present for RetainRepoVersions, not even an explicit nil
### GetRemote

`func (o *FileFileRepositoryResponse) GetRemote() string`

GetRemote returns the Remote field if non-nil, zero value otherwise.

### GetRemoteOk

`func (o *FileFileRepositoryResponse) GetRemoteOk() (*string, bool)`

GetRemoteOk returns a tuple with the Remote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemote

`func (o *FileFileRepositoryResponse) SetRemote(v string)`

SetRemote sets Remote field to given value.

### HasRemote

`func (o *FileFileRepositoryResponse) HasRemote() bool`

HasRemote returns a boolean if a field has been set.

### SetRemoteNil

`func (o *FileFileRepositoryResponse) SetRemoteNil(b bool)`

 SetRemoteNil sets the value for Remote to be an explicit nil

### UnsetRemote
`func (o *FileFileRepositoryResponse) UnsetRemote()`

UnsetRemote ensures that no value is present for Remote, not even an explicit nil
### GetAutopublish

`func (o *FileFileRepositoryResponse) GetAutopublish() bool`

GetAutopublish returns the Autopublish field if non-nil, zero value otherwise.

### GetAutopublishOk

`func (o *FileFileRepositoryResponse) GetAutopublishOk() (*bool, bool)`

GetAutopublishOk returns a tuple with the Autopublish field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutopublish

`func (o *FileFileRepositoryResponse) SetAutopublish(v bool)`

SetAutopublish sets Autopublish field to given value.

### HasAutopublish

`func (o *FileFileRepositoryResponse) HasAutopublish() bool`

HasAutopublish returns a boolean if a field has been set.

### GetManifest

`func (o *FileFileRepositoryResponse) GetManifest() string`

GetManifest returns the Manifest field if non-nil, zero value otherwise.

### GetManifestOk

`func (o *FileFileRepositoryResponse) GetManifestOk() (*string, bool)`

GetManifestOk returns a tuple with the Manifest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManifest

`func (o *FileFileRepositoryResponse) SetManifest(v string)`

SetManifest sets Manifest field to given value.

### HasManifest

`func (o *FileFileRepositoryResponse) HasManifest() bool`

HasManifest returns a boolean if a field has been set.

### SetManifestNil

`func (o *FileFileRepositoryResponse) SetManifestNil(b bool)`

 SetManifestNil sets the value for Manifest to be an explicit nil

### UnsetManifest
`func (o *FileFileRepositoryResponse) UnsetManifest()`

UnsetManifest ensures that no value is present for Manifest, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


