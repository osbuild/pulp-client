# OstreeImportCommitsToRef

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Artifact** | **string** | An artifact representing OSTree content compressed as a tarball. | 
**RepositoryName** | **string** | The name of a repository that contains the compressed OSTree content. | 
**Ref** | **string** | The name of a ref branch that holds the reference to the last commit. | 

## Methods

### NewOstreeImportCommitsToRef

`func NewOstreeImportCommitsToRef(artifact string, repositoryName string, ref string, ) *OstreeImportCommitsToRef`

NewOstreeImportCommitsToRef instantiates a new OstreeImportCommitsToRef object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOstreeImportCommitsToRefWithDefaults

`func NewOstreeImportCommitsToRefWithDefaults() *OstreeImportCommitsToRef`

NewOstreeImportCommitsToRefWithDefaults instantiates a new OstreeImportCommitsToRef object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArtifact

`func (o *OstreeImportCommitsToRef) GetArtifact() string`

GetArtifact returns the Artifact field if non-nil, zero value otherwise.

### GetArtifactOk

`func (o *OstreeImportCommitsToRef) GetArtifactOk() (*string, bool)`

GetArtifactOk returns a tuple with the Artifact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifact

`func (o *OstreeImportCommitsToRef) SetArtifact(v string)`

SetArtifact sets Artifact field to given value.


### GetRepositoryName

`func (o *OstreeImportCommitsToRef) GetRepositoryName() string`

GetRepositoryName returns the RepositoryName field if non-nil, zero value otherwise.

### GetRepositoryNameOk

`func (o *OstreeImportCommitsToRef) GetRepositoryNameOk() (*string, bool)`

GetRepositoryNameOk returns a tuple with the RepositoryName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepositoryName

`func (o *OstreeImportCommitsToRef) SetRepositoryName(v string)`

SetRepositoryName sets RepositoryName field to given value.


### GetRef

`func (o *OstreeImportCommitsToRef) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *OstreeImportCommitsToRef) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *OstreeImportCommitsToRef) SetRef(v string)`

SetRef sets Ref field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


