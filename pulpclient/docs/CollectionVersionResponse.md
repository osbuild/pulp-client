# CollectionVersionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Version** | Pointer to **string** |  | [optional] [readonly] 
**Href** | Pointer to **string** | Get href. | [optional] [readonly] 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 
**RequiresAnsible** | Pointer to **NullableString** |  | [optional] 
**Marks** | Pointer to **[]string** | Get a list of mark values filtering only those in the current repo. | [optional] [readonly] 
**Artifact** | Pointer to [**ArtifactRefResponse**](ArtifactRefResponse.md) |  | [optional] [readonly] 
**Collection** | Pointer to [**CollectionRefResponse**](CollectionRefResponse.md) |  | [optional] [readonly] 
**DownloadUrl** | Pointer to **string** | Get artifact download URL. | [optional] [readonly] 
**Name** | Pointer to **string** |  | [optional] [readonly] 
**Namespace** | Pointer to [**CollectionNamespaceResponse**](CollectionNamespaceResponse.md) |  | [optional] [readonly] 
**Signatures** | [**[]CollectionVersionSignatureResponse**](CollectionVersionSignatureResponse.md) |  | 
**Metadata** | Pointer to [**CollectionMetadataResponse**](CollectionMetadataResponse.md) |  | [optional] [readonly] 
**GitUrl** | Pointer to **string** | Get the git URL. | [optional] [readonly] 
**GitCommitSha** | Pointer to **string** | Get the git commit sha. | [optional] [readonly] 
**Manifest** | Pointer to **map[string]interface{}** | A JSON field holding MANIFEST.json data. | [optional] [readonly] 
**Files** | Pointer to **map[string]interface{}** | A JSON field holding FILES.json data. | [optional] [readonly] 

## Methods

### NewCollectionVersionResponse

`func NewCollectionVersionResponse(createdAt time.Time, updatedAt time.Time, signatures []CollectionVersionSignatureResponse, ) *CollectionVersionResponse`

NewCollectionVersionResponse instantiates a new CollectionVersionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCollectionVersionResponseWithDefaults

`func NewCollectionVersionResponseWithDefaults() *CollectionVersionResponse`

NewCollectionVersionResponseWithDefaults instantiates a new CollectionVersionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersion

`func (o *CollectionVersionResponse) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *CollectionVersionResponse) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *CollectionVersionResponse) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *CollectionVersionResponse) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetHref

`func (o *CollectionVersionResponse) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *CollectionVersionResponse) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *CollectionVersionResponse) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *CollectionVersionResponse) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetCreatedAt

`func (o *CollectionVersionResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CollectionVersionResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CollectionVersionResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *CollectionVersionResponse) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *CollectionVersionResponse) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *CollectionVersionResponse) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetRequiresAnsible

`func (o *CollectionVersionResponse) GetRequiresAnsible() string`

GetRequiresAnsible returns the RequiresAnsible field if non-nil, zero value otherwise.

### GetRequiresAnsibleOk

`func (o *CollectionVersionResponse) GetRequiresAnsibleOk() (*string, bool)`

GetRequiresAnsibleOk returns a tuple with the RequiresAnsible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiresAnsible

`func (o *CollectionVersionResponse) SetRequiresAnsible(v string)`

SetRequiresAnsible sets RequiresAnsible field to given value.

### HasRequiresAnsible

`func (o *CollectionVersionResponse) HasRequiresAnsible() bool`

HasRequiresAnsible returns a boolean if a field has been set.

### SetRequiresAnsibleNil

`func (o *CollectionVersionResponse) SetRequiresAnsibleNil(b bool)`

 SetRequiresAnsibleNil sets the value for RequiresAnsible to be an explicit nil

### UnsetRequiresAnsible
`func (o *CollectionVersionResponse) UnsetRequiresAnsible()`

UnsetRequiresAnsible ensures that no value is present for RequiresAnsible, not even an explicit nil
### GetMarks

`func (o *CollectionVersionResponse) GetMarks() []string`

GetMarks returns the Marks field if non-nil, zero value otherwise.

### GetMarksOk

`func (o *CollectionVersionResponse) GetMarksOk() (*[]string, bool)`

GetMarksOk returns a tuple with the Marks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarks

`func (o *CollectionVersionResponse) SetMarks(v []string)`

SetMarks sets Marks field to given value.

### HasMarks

`func (o *CollectionVersionResponse) HasMarks() bool`

HasMarks returns a boolean if a field has been set.

### GetArtifact

`func (o *CollectionVersionResponse) GetArtifact() ArtifactRefResponse`

GetArtifact returns the Artifact field if non-nil, zero value otherwise.

### GetArtifactOk

`func (o *CollectionVersionResponse) GetArtifactOk() (*ArtifactRefResponse, bool)`

GetArtifactOk returns a tuple with the Artifact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifact

`func (o *CollectionVersionResponse) SetArtifact(v ArtifactRefResponse)`

SetArtifact sets Artifact field to given value.

### HasArtifact

`func (o *CollectionVersionResponse) HasArtifact() bool`

HasArtifact returns a boolean if a field has been set.

### GetCollection

`func (o *CollectionVersionResponse) GetCollection() CollectionRefResponse`

GetCollection returns the Collection field if non-nil, zero value otherwise.

### GetCollectionOk

`func (o *CollectionVersionResponse) GetCollectionOk() (*CollectionRefResponse, bool)`

GetCollectionOk returns a tuple with the Collection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollection

`func (o *CollectionVersionResponse) SetCollection(v CollectionRefResponse)`

SetCollection sets Collection field to given value.

### HasCollection

`func (o *CollectionVersionResponse) HasCollection() bool`

HasCollection returns a boolean if a field has been set.

### GetDownloadUrl

`func (o *CollectionVersionResponse) GetDownloadUrl() string`

GetDownloadUrl returns the DownloadUrl field if non-nil, zero value otherwise.

### GetDownloadUrlOk

`func (o *CollectionVersionResponse) GetDownloadUrlOk() (*string, bool)`

GetDownloadUrlOk returns a tuple with the DownloadUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadUrl

`func (o *CollectionVersionResponse) SetDownloadUrl(v string)`

SetDownloadUrl sets DownloadUrl field to given value.

### HasDownloadUrl

`func (o *CollectionVersionResponse) HasDownloadUrl() bool`

HasDownloadUrl returns a boolean if a field has been set.

### GetName

`func (o *CollectionVersionResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CollectionVersionResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CollectionVersionResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CollectionVersionResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNamespace

`func (o *CollectionVersionResponse) GetNamespace() CollectionNamespaceResponse`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *CollectionVersionResponse) GetNamespaceOk() (*CollectionNamespaceResponse, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *CollectionVersionResponse) SetNamespace(v CollectionNamespaceResponse)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *CollectionVersionResponse) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetSignatures

`func (o *CollectionVersionResponse) GetSignatures() []CollectionVersionSignatureResponse`

GetSignatures returns the Signatures field if non-nil, zero value otherwise.

### GetSignaturesOk

`func (o *CollectionVersionResponse) GetSignaturesOk() (*[]CollectionVersionSignatureResponse, bool)`

GetSignaturesOk returns a tuple with the Signatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignatures

`func (o *CollectionVersionResponse) SetSignatures(v []CollectionVersionSignatureResponse)`

SetSignatures sets Signatures field to given value.


### GetMetadata

`func (o *CollectionVersionResponse) GetMetadata() CollectionMetadataResponse`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CollectionVersionResponse) GetMetadataOk() (*CollectionMetadataResponse, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CollectionVersionResponse) SetMetadata(v CollectionMetadataResponse)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CollectionVersionResponse) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetGitUrl

`func (o *CollectionVersionResponse) GetGitUrl() string`

GetGitUrl returns the GitUrl field if non-nil, zero value otherwise.

### GetGitUrlOk

`func (o *CollectionVersionResponse) GetGitUrlOk() (*string, bool)`

GetGitUrlOk returns a tuple with the GitUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitUrl

`func (o *CollectionVersionResponse) SetGitUrl(v string)`

SetGitUrl sets GitUrl field to given value.

### HasGitUrl

`func (o *CollectionVersionResponse) HasGitUrl() bool`

HasGitUrl returns a boolean if a field has been set.

### GetGitCommitSha

`func (o *CollectionVersionResponse) GetGitCommitSha() string`

GetGitCommitSha returns the GitCommitSha field if non-nil, zero value otherwise.

### GetGitCommitShaOk

`func (o *CollectionVersionResponse) GetGitCommitShaOk() (*string, bool)`

GetGitCommitShaOk returns a tuple with the GitCommitSha field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitCommitSha

`func (o *CollectionVersionResponse) SetGitCommitSha(v string)`

SetGitCommitSha sets GitCommitSha field to given value.

### HasGitCommitSha

`func (o *CollectionVersionResponse) HasGitCommitSha() bool`

HasGitCommitSha returns a boolean if a field has been set.

### GetManifest

`func (o *CollectionVersionResponse) GetManifest() map[string]interface{}`

GetManifest returns the Manifest field if non-nil, zero value otherwise.

### GetManifestOk

`func (o *CollectionVersionResponse) GetManifestOk() (*map[string]interface{}, bool)`

GetManifestOk returns a tuple with the Manifest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManifest

`func (o *CollectionVersionResponse) SetManifest(v map[string]interface{})`

SetManifest sets Manifest field to given value.

### HasManifest

`func (o *CollectionVersionResponse) HasManifest() bool`

HasManifest returns a boolean if a field has been set.

### GetFiles

`func (o *CollectionVersionResponse) GetFiles() map[string]interface{}`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *CollectionVersionResponse) GetFilesOk() (*map[string]interface{}, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *CollectionVersionResponse) SetFiles(v map[string]interface{})`

SetFiles sets Files field to given value.

### HasFiles

`func (o *CollectionVersionResponse) HasFiles() bool`

HasFiles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


