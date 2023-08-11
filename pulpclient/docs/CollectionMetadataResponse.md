# CollectionMetadataResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Authors** | Pointer to **[]string** |  | [optional] [readonly] 
**Contents** | Pointer to **map[string]interface{}** |  | [optional] [readonly] 
**Dependencies** | Pointer to **map[string]interface{}** |  | [optional] [readonly] 
**Description** | Pointer to **string** |  | [optional] [readonly] 
**Documentation** | Pointer to **string** |  | [optional] [readonly] 
**Homepage** | Pointer to **string** |  | [optional] [readonly] 
**Issues** | Pointer to **string** |  | [optional] [readonly] 
**License** | Pointer to **[]string** |  | [optional] [readonly] 
**Repository** | Pointer to **string** |  | [optional] [readonly] 
**Tags** | **[]string** |  | 

## Methods

### NewCollectionMetadataResponse

`func NewCollectionMetadataResponse(tags []string, ) *CollectionMetadataResponse`

NewCollectionMetadataResponse instantiates a new CollectionMetadataResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCollectionMetadataResponseWithDefaults

`func NewCollectionMetadataResponseWithDefaults() *CollectionMetadataResponse`

NewCollectionMetadataResponseWithDefaults instantiates a new CollectionMetadataResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthors

`func (o *CollectionMetadataResponse) GetAuthors() []string`

GetAuthors returns the Authors field if non-nil, zero value otherwise.

### GetAuthorsOk

`func (o *CollectionMetadataResponse) GetAuthorsOk() (*[]string, bool)`

GetAuthorsOk returns a tuple with the Authors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthors

`func (o *CollectionMetadataResponse) SetAuthors(v []string)`

SetAuthors sets Authors field to given value.

### HasAuthors

`func (o *CollectionMetadataResponse) HasAuthors() bool`

HasAuthors returns a boolean if a field has been set.

### GetContents

`func (o *CollectionMetadataResponse) GetContents() map[string]interface{}`

GetContents returns the Contents field if non-nil, zero value otherwise.

### GetContentsOk

`func (o *CollectionMetadataResponse) GetContentsOk() (*map[string]interface{}, bool)`

GetContentsOk returns a tuple with the Contents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContents

`func (o *CollectionMetadataResponse) SetContents(v map[string]interface{})`

SetContents sets Contents field to given value.

### HasContents

`func (o *CollectionMetadataResponse) HasContents() bool`

HasContents returns a boolean if a field has been set.

### GetDependencies

`func (o *CollectionMetadataResponse) GetDependencies() map[string]interface{}`

GetDependencies returns the Dependencies field if non-nil, zero value otherwise.

### GetDependenciesOk

`func (o *CollectionMetadataResponse) GetDependenciesOk() (*map[string]interface{}, bool)`

GetDependenciesOk returns a tuple with the Dependencies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependencies

`func (o *CollectionMetadataResponse) SetDependencies(v map[string]interface{})`

SetDependencies sets Dependencies field to given value.

### HasDependencies

`func (o *CollectionMetadataResponse) HasDependencies() bool`

HasDependencies returns a boolean if a field has been set.

### GetDescription

`func (o *CollectionMetadataResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CollectionMetadataResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CollectionMetadataResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CollectionMetadataResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDocumentation

`func (o *CollectionMetadataResponse) GetDocumentation() string`

GetDocumentation returns the Documentation field if non-nil, zero value otherwise.

### GetDocumentationOk

`func (o *CollectionMetadataResponse) GetDocumentationOk() (*string, bool)`

GetDocumentationOk returns a tuple with the Documentation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentation

`func (o *CollectionMetadataResponse) SetDocumentation(v string)`

SetDocumentation sets Documentation field to given value.

### HasDocumentation

`func (o *CollectionMetadataResponse) HasDocumentation() bool`

HasDocumentation returns a boolean if a field has been set.

### GetHomepage

`func (o *CollectionMetadataResponse) GetHomepage() string`

GetHomepage returns the Homepage field if non-nil, zero value otherwise.

### GetHomepageOk

`func (o *CollectionMetadataResponse) GetHomepageOk() (*string, bool)`

GetHomepageOk returns a tuple with the Homepage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHomepage

`func (o *CollectionMetadataResponse) SetHomepage(v string)`

SetHomepage sets Homepage field to given value.

### HasHomepage

`func (o *CollectionMetadataResponse) HasHomepage() bool`

HasHomepage returns a boolean if a field has been set.

### GetIssues

`func (o *CollectionMetadataResponse) GetIssues() string`

GetIssues returns the Issues field if non-nil, zero value otherwise.

### GetIssuesOk

`func (o *CollectionMetadataResponse) GetIssuesOk() (*string, bool)`

GetIssuesOk returns a tuple with the Issues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssues

`func (o *CollectionMetadataResponse) SetIssues(v string)`

SetIssues sets Issues field to given value.

### HasIssues

`func (o *CollectionMetadataResponse) HasIssues() bool`

HasIssues returns a boolean if a field has been set.

### GetLicense

`func (o *CollectionMetadataResponse) GetLicense() []string`

GetLicense returns the License field if non-nil, zero value otherwise.

### GetLicenseOk

`func (o *CollectionMetadataResponse) GetLicenseOk() (*[]string, bool)`

GetLicenseOk returns a tuple with the License field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicense

`func (o *CollectionMetadataResponse) SetLicense(v []string)`

SetLicense sets License field to given value.

### HasLicense

`func (o *CollectionMetadataResponse) HasLicense() bool`

HasLicense returns a boolean if a field has been set.

### GetRepository

`func (o *CollectionMetadataResponse) GetRepository() string`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *CollectionMetadataResponse) GetRepositoryOk() (*string, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *CollectionMetadataResponse) SetRepository(v string)`

SetRepository sets Repository field to given value.

### HasRepository

`func (o *CollectionMetadataResponse) HasRepository() bool`

HasRepository returns a boolean if a field has been set.

### GetTags

`func (o *CollectionMetadataResponse) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CollectionMetadataResponse) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CollectionMetadataResponse) SetTags(v []string)`

SetTags sets Tags field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


