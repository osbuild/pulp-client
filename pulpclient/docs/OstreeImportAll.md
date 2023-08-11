# OstreeImportAll

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Artifact** | **string** | An artifact representing OSTree content compressed as a tarball. | 
**RepositoryName** | **string** | The name of a repository that contains the compressed OSTree content. | 

## Methods

### NewOstreeImportAll

`func NewOstreeImportAll(artifact string, repositoryName string, ) *OstreeImportAll`

NewOstreeImportAll instantiates a new OstreeImportAll object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOstreeImportAllWithDefaults

`func NewOstreeImportAllWithDefaults() *OstreeImportAll`

NewOstreeImportAllWithDefaults instantiates a new OstreeImportAll object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArtifact

`func (o *OstreeImportAll) GetArtifact() string`

GetArtifact returns the Artifact field if non-nil, zero value otherwise.

### GetArtifactOk

`func (o *OstreeImportAll) GetArtifactOk() (*string, bool)`

GetArtifactOk returns a tuple with the Artifact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifact

`func (o *OstreeImportAll) SetArtifact(v string)`

SetArtifact sets Artifact field to given value.


### GetRepositoryName

`func (o *OstreeImportAll) GetRepositoryName() string`

GetRepositoryName returns the RepositoryName field if non-nil, zero value otherwise.

### GetRepositoryNameOk

`func (o *OstreeImportAll) GetRepositoryNameOk() (*string, bool)`

GetRepositoryNameOk returns a tuple with the RepositoryName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepositoryName

`func (o *OstreeImportAll) SetRepositoryName(v string)`

SetRepositoryName sets RepositoryName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


