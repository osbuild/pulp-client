# UnpaginatedCollectionVersionResponse

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

## Methods

### NewUnpaginatedCollectionVersionResponse

`func NewUnpaginatedCollectionVersionResponse(createdAt time.Time, updatedAt time.Time, signatures []CollectionVersionSignatureResponse, ) *UnpaginatedCollectionVersionResponse`

NewUnpaginatedCollectionVersionResponse instantiates a new UnpaginatedCollectionVersionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnpaginatedCollectionVersionResponseWithDefaults

`func NewUnpaginatedCollectionVersionResponseWithDefaults() *UnpaginatedCollectionVersionResponse`

NewUnpaginatedCollectionVersionResponseWithDefaults instantiates a new UnpaginatedCollectionVersionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersion

`func (o *UnpaginatedCollectionVersionResponse) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *UnpaginatedCollectionVersionResponse) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *UnpaginatedCollectionVersionResponse) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *UnpaginatedCollectionVersionResponse) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetHref

`func (o *UnpaginatedCollectionVersionResponse) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *UnpaginatedCollectionVersionResponse) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *UnpaginatedCollectionVersionResponse) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *UnpaginatedCollectionVersionResponse) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetCreatedAt

`func (o *UnpaginatedCollectionVersionResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *UnpaginatedCollectionVersionResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *UnpaginatedCollectionVersionResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *UnpaginatedCollectionVersionResponse) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *UnpaginatedCollectionVersionResponse) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *UnpaginatedCollectionVersionResponse) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetRequiresAnsible

`func (o *UnpaginatedCollectionVersionResponse) GetRequiresAnsible() string`

GetRequiresAnsible returns the RequiresAnsible field if non-nil, zero value otherwise.

### GetRequiresAnsibleOk

`func (o *UnpaginatedCollectionVersionResponse) GetRequiresAnsibleOk() (*string, bool)`

GetRequiresAnsibleOk returns a tuple with the RequiresAnsible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiresAnsible

`func (o *UnpaginatedCollectionVersionResponse) SetRequiresAnsible(v string)`

SetRequiresAnsible sets RequiresAnsible field to given value.

### HasRequiresAnsible

`func (o *UnpaginatedCollectionVersionResponse) HasRequiresAnsible() bool`

HasRequiresAnsible returns a boolean if a field has been set.

### SetRequiresAnsibleNil

`func (o *UnpaginatedCollectionVersionResponse) SetRequiresAnsibleNil(b bool)`

 SetRequiresAnsibleNil sets the value for RequiresAnsible to be an explicit nil

### UnsetRequiresAnsible
`func (o *UnpaginatedCollectionVersionResponse) UnsetRequiresAnsible()`

UnsetRequiresAnsible ensures that no value is present for RequiresAnsible, not even an explicit nil
### GetMarks

`func (o *UnpaginatedCollectionVersionResponse) GetMarks() []string`

GetMarks returns the Marks field if non-nil, zero value otherwise.

### GetMarksOk

`func (o *UnpaginatedCollectionVersionResponse) GetMarksOk() (*[]string, bool)`

GetMarksOk returns a tuple with the Marks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarks

`func (o *UnpaginatedCollectionVersionResponse) SetMarks(v []string)`

SetMarks sets Marks field to given value.

### HasMarks

`func (o *UnpaginatedCollectionVersionResponse) HasMarks() bool`

HasMarks returns a boolean if a field has been set.

### GetArtifact

`func (o *UnpaginatedCollectionVersionResponse) GetArtifact() ArtifactRefResponse`

GetArtifact returns the Artifact field if non-nil, zero value otherwise.

### GetArtifactOk

`func (o *UnpaginatedCollectionVersionResponse) GetArtifactOk() (*ArtifactRefResponse, bool)`

GetArtifactOk returns a tuple with the Artifact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifact

`func (o *UnpaginatedCollectionVersionResponse) SetArtifact(v ArtifactRefResponse)`

SetArtifact sets Artifact field to given value.

### HasArtifact

`func (o *UnpaginatedCollectionVersionResponse) HasArtifact() bool`

HasArtifact returns a boolean if a field has been set.

### GetCollection

`func (o *UnpaginatedCollectionVersionResponse) GetCollection() CollectionRefResponse`

GetCollection returns the Collection field if non-nil, zero value otherwise.

### GetCollectionOk

`func (o *UnpaginatedCollectionVersionResponse) GetCollectionOk() (*CollectionRefResponse, bool)`

GetCollectionOk returns a tuple with the Collection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollection

`func (o *UnpaginatedCollectionVersionResponse) SetCollection(v CollectionRefResponse)`

SetCollection sets Collection field to given value.

### HasCollection

`func (o *UnpaginatedCollectionVersionResponse) HasCollection() bool`

HasCollection returns a boolean if a field has been set.

### GetDownloadUrl

`func (o *UnpaginatedCollectionVersionResponse) GetDownloadUrl() string`

GetDownloadUrl returns the DownloadUrl field if non-nil, zero value otherwise.

### GetDownloadUrlOk

`func (o *UnpaginatedCollectionVersionResponse) GetDownloadUrlOk() (*string, bool)`

GetDownloadUrlOk returns a tuple with the DownloadUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadUrl

`func (o *UnpaginatedCollectionVersionResponse) SetDownloadUrl(v string)`

SetDownloadUrl sets DownloadUrl field to given value.

### HasDownloadUrl

`func (o *UnpaginatedCollectionVersionResponse) HasDownloadUrl() bool`

HasDownloadUrl returns a boolean if a field has been set.

### GetName

`func (o *UnpaginatedCollectionVersionResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UnpaginatedCollectionVersionResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UnpaginatedCollectionVersionResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UnpaginatedCollectionVersionResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNamespace

`func (o *UnpaginatedCollectionVersionResponse) GetNamespace() CollectionNamespaceResponse`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *UnpaginatedCollectionVersionResponse) GetNamespaceOk() (*CollectionNamespaceResponse, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *UnpaginatedCollectionVersionResponse) SetNamespace(v CollectionNamespaceResponse)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *UnpaginatedCollectionVersionResponse) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetSignatures

`func (o *UnpaginatedCollectionVersionResponse) GetSignatures() []CollectionVersionSignatureResponse`

GetSignatures returns the Signatures field if non-nil, zero value otherwise.

### GetSignaturesOk

`func (o *UnpaginatedCollectionVersionResponse) GetSignaturesOk() (*[]CollectionVersionSignatureResponse, bool)`

GetSignaturesOk returns a tuple with the Signatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignatures

`func (o *UnpaginatedCollectionVersionResponse) SetSignatures(v []CollectionVersionSignatureResponse)`

SetSignatures sets Signatures field to given value.


### GetMetadata

`func (o *UnpaginatedCollectionVersionResponse) GetMetadata() CollectionMetadataResponse`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *UnpaginatedCollectionVersionResponse) GetMetadataOk() (*CollectionMetadataResponse, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *UnpaginatedCollectionVersionResponse) SetMetadata(v CollectionMetadataResponse)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *UnpaginatedCollectionVersionResponse) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetGitUrl

`func (o *UnpaginatedCollectionVersionResponse) GetGitUrl() string`

GetGitUrl returns the GitUrl field if non-nil, zero value otherwise.

### GetGitUrlOk

`func (o *UnpaginatedCollectionVersionResponse) GetGitUrlOk() (*string, bool)`

GetGitUrlOk returns a tuple with the GitUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitUrl

`func (o *UnpaginatedCollectionVersionResponse) SetGitUrl(v string)`

SetGitUrl sets GitUrl field to given value.

### HasGitUrl

`func (o *UnpaginatedCollectionVersionResponse) HasGitUrl() bool`

HasGitUrl returns a boolean if a field has been set.

### GetGitCommitSha

`func (o *UnpaginatedCollectionVersionResponse) GetGitCommitSha() string`

GetGitCommitSha returns the GitCommitSha field if non-nil, zero value otherwise.

### GetGitCommitShaOk

`func (o *UnpaginatedCollectionVersionResponse) GetGitCommitShaOk() (*string, bool)`

GetGitCommitShaOk returns a tuple with the GitCommitSha field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitCommitSha

`func (o *UnpaginatedCollectionVersionResponse) SetGitCommitSha(v string)`

SetGitCommitSha sets GitCommitSha field to given value.

### HasGitCommitSha

`func (o *UnpaginatedCollectionVersionResponse) HasGitCommitSha() bool`

HasGitCommitSha returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


