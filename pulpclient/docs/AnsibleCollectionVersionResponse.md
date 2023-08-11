# AnsibleCollectionVersionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PulpHref** | Pointer to **string** |  | [optional] [readonly] 
**Artifact** | Pointer to **string** | Artifact file representing the physical content | [optional] 
**PulpCreated** | Pointer to **time.Time** | Timestamp of creation. | [optional] [readonly] 
**Sha256** | Pointer to **string** | The SHA-256 checksum if available. | [optional] [readonly] 
**Md5** | Pointer to **string** | The MD5 checksum if available. | [optional] [readonly] 
**Sha1** | Pointer to **string** | The SHA-1 checksum if available. | [optional] [readonly] 
**Sha224** | Pointer to **string** | The SHA-224 checksum if available. | [optional] [readonly] 
**Sha384** | Pointer to **string** | The SHA-384 checksum if available. | [optional] [readonly] 
**Sha512** | Pointer to **string** | The SHA-512 checksum if available. | [optional] [readonly] 
**Id** | Pointer to **string** | A collection identifier. | [optional] [readonly] 
**Authors** | Pointer to **[]string** | A list of the CollectionVersion content&#39;s authors. | [optional] [readonly] 
**Contents** | Pointer to **map[string]interface{}** | A JSON field with data about the contents. | [optional] [readonly] 
**Dependencies** | Pointer to **map[string]interface{}** | A dict declaring Collections that this collection requires to be installed for it to be usable. | [optional] [readonly] 
**Description** | Pointer to **string** | A short summary description of the collection. | [optional] [readonly] 
**DocsBlob** | Pointer to **map[string]interface{}** | A JSON field holding the various documentation blobs in the collection. | [optional] [readonly] 
**Manifest** | Pointer to **map[string]interface{}** | A JSON field holding MANIFEST.json data. | [optional] [readonly] 
**Files** | Pointer to **map[string]interface{}** | A JSON field holding FILES.json data. | [optional] [readonly] 
**Documentation** | Pointer to **string** | The URL to any online docs. | [optional] [readonly] 
**Homepage** | Pointer to **string** | The URL to the homepage of the collection/project. | [optional] [readonly] 
**Issues** | Pointer to **string** | The URL to the collection issue tracker. | [optional] [readonly] 
**License** | Pointer to **[]string** | A list of licenses for content inside of a collection. | [optional] [readonly] 
**Name** | Pointer to **string** | The name of the collection. | [optional] [readonly] 
**Namespace** | Pointer to **string** | The namespace of the collection. | [optional] [readonly] 
**OriginRepository** | Pointer to **string** | The URL of the originating SCM repository. | [optional] [readonly] 
**Tags** | Pointer to [**[]AnsibleTagResponse**](AnsibleTagResponse.md) |  | [optional] [readonly] 
**Version** | Pointer to **string** | The version of the collection. | [optional] [readonly] 
**RequiresAnsible** | Pointer to **NullableString** | The version of Ansible required to use the collection. Multiple versions can be separated with a comma. | [optional] [readonly] 

## Methods

### NewAnsibleCollectionVersionResponse

`func NewAnsibleCollectionVersionResponse() *AnsibleCollectionVersionResponse`

NewAnsibleCollectionVersionResponse instantiates a new AnsibleCollectionVersionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnsibleCollectionVersionResponseWithDefaults

`func NewAnsibleCollectionVersionResponseWithDefaults() *AnsibleCollectionVersionResponse`

NewAnsibleCollectionVersionResponseWithDefaults instantiates a new AnsibleCollectionVersionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPulpHref

`func (o *AnsibleCollectionVersionResponse) GetPulpHref() string`

GetPulpHref returns the PulpHref field if non-nil, zero value otherwise.

### GetPulpHrefOk

`func (o *AnsibleCollectionVersionResponse) GetPulpHrefOk() (*string, bool)`

GetPulpHrefOk returns a tuple with the PulpHref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpHref

`func (o *AnsibleCollectionVersionResponse) SetPulpHref(v string)`

SetPulpHref sets PulpHref field to given value.

### HasPulpHref

`func (o *AnsibleCollectionVersionResponse) HasPulpHref() bool`

HasPulpHref returns a boolean if a field has been set.

### GetArtifact

`func (o *AnsibleCollectionVersionResponse) GetArtifact() string`

GetArtifact returns the Artifact field if non-nil, zero value otherwise.

### GetArtifactOk

`func (o *AnsibleCollectionVersionResponse) GetArtifactOk() (*string, bool)`

GetArtifactOk returns a tuple with the Artifact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifact

`func (o *AnsibleCollectionVersionResponse) SetArtifact(v string)`

SetArtifact sets Artifact field to given value.

### HasArtifact

`func (o *AnsibleCollectionVersionResponse) HasArtifact() bool`

HasArtifact returns a boolean if a field has been set.

### GetPulpCreated

`func (o *AnsibleCollectionVersionResponse) GetPulpCreated() time.Time`

GetPulpCreated returns the PulpCreated field if non-nil, zero value otherwise.

### GetPulpCreatedOk

`func (o *AnsibleCollectionVersionResponse) GetPulpCreatedOk() (*time.Time, bool)`

GetPulpCreatedOk returns a tuple with the PulpCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpCreated

`func (o *AnsibleCollectionVersionResponse) SetPulpCreated(v time.Time)`

SetPulpCreated sets PulpCreated field to given value.

### HasPulpCreated

`func (o *AnsibleCollectionVersionResponse) HasPulpCreated() bool`

HasPulpCreated returns a boolean if a field has been set.

### GetSha256

`func (o *AnsibleCollectionVersionResponse) GetSha256() string`

GetSha256 returns the Sha256 field if non-nil, zero value otherwise.

### GetSha256Ok

`func (o *AnsibleCollectionVersionResponse) GetSha256Ok() (*string, bool)`

GetSha256Ok returns a tuple with the Sha256 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha256

`func (o *AnsibleCollectionVersionResponse) SetSha256(v string)`

SetSha256 sets Sha256 field to given value.

### HasSha256

`func (o *AnsibleCollectionVersionResponse) HasSha256() bool`

HasSha256 returns a boolean if a field has been set.

### GetMd5

`func (o *AnsibleCollectionVersionResponse) GetMd5() string`

GetMd5 returns the Md5 field if non-nil, zero value otherwise.

### GetMd5Ok

`func (o *AnsibleCollectionVersionResponse) GetMd5Ok() (*string, bool)`

GetMd5Ok returns a tuple with the Md5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMd5

`func (o *AnsibleCollectionVersionResponse) SetMd5(v string)`

SetMd5 sets Md5 field to given value.

### HasMd5

`func (o *AnsibleCollectionVersionResponse) HasMd5() bool`

HasMd5 returns a boolean if a field has been set.

### GetSha1

`func (o *AnsibleCollectionVersionResponse) GetSha1() string`

GetSha1 returns the Sha1 field if non-nil, zero value otherwise.

### GetSha1Ok

`func (o *AnsibleCollectionVersionResponse) GetSha1Ok() (*string, bool)`

GetSha1Ok returns a tuple with the Sha1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha1

`func (o *AnsibleCollectionVersionResponse) SetSha1(v string)`

SetSha1 sets Sha1 field to given value.

### HasSha1

`func (o *AnsibleCollectionVersionResponse) HasSha1() bool`

HasSha1 returns a boolean if a field has been set.

### GetSha224

`func (o *AnsibleCollectionVersionResponse) GetSha224() string`

GetSha224 returns the Sha224 field if non-nil, zero value otherwise.

### GetSha224Ok

`func (o *AnsibleCollectionVersionResponse) GetSha224Ok() (*string, bool)`

GetSha224Ok returns a tuple with the Sha224 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha224

`func (o *AnsibleCollectionVersionResponse) SetSha224(v string)`

SetSha224 sets Sha224 field to given value.

### HasSha224

`func (o *AnsibleCollectionVersionResponse) HasSha224() bool`

HasSha224 returns a boolean if a field has been set.

### GetSha384

`func (o *AnsibleCollectionVersionResponse) GetSha384() string`

GetSha384 returns the Sha384 field if non-nil, zero value otherwise.

### GetSha384Ok

`func (o *AnsibleCollectionVersionResponse) GetSha384Ok() (*string, bool)`

GetSha384Ok returns a tuple with the Sha384 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha384

`func (o *AnsibleCollectionVersionResponse) SetSha384(v string)`

SetSha384 sets Sha384 field to given value.

### HasSha384

`func (o *AnsibleCollectionVersionResponse) HasSha384() bool`

HasSha384 returns a boolean if a field has been set.

### GetSha512

`func (o *AnsibleCollectionVersionResponse) GetSha512() string`

GetSha512 returns the Sha512 field if non-nil, zero value otherwise.

### GetSha512Ok

`func (o *AnsibleCollectionVersionResponse) GetSha512Ok() (*string, bool)`

GetSha512Ok returns a tuple with the Sha512 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha512

`func (o *AnsibleCollectionVersionResponse) SetSha512(v string)`

SetSha512 sets Sha512 field to given value.

### HasSha512

`func (o *AnsibleCollectionVersionResponse) HasSha512() bool`

HasSha512 returns a boolean if a field has been set.

### GetId

`func (o *AnsibleCollectionVersionResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AnsibleCollectionVersionResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AnsibleCollectionVersionResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AnsibleCollectionVersionResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAuthors

`func (o *AnsibleCollectionVersionResponse) GetAuthors() []string`

GetAuthors returns the Authors field if non-nil, zero value otherwise.

### GetAuthorsOk

`func (o *AnsibleCollectionVersionResponse) GetAuthorsOk() (*[]string, bool)`

GetAuthorsOk returns a tuple with the Authors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthors

`func (o *AnsibleCollectionVersionResponse) SetAuthors(v []string)`

SetAuthors sets Authors field to given value.

### HasAuthors

`func (o *AnsibleCollectionVersionResponse) HasAuthors() bool`

HasAuthors returns a boolean if a field has been set.

### GetContents

`func (o *AnsibleCollectionVersionResponse) GetContents() map[string]interface{}`

GetContents returns the Contents field if non-nil, zero value otherwise.

### GetContentsOk

`func (o *AnsibleCollectionVersionResponse) GetContentsOk() (*map[string]interface{}, bool)`

GetContentsOk returns a tuple with the Contents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContents

`func (o *AnsibleCollectionVersionResponse) SetContents(v map[string]interface{})`

SetContents sets Contents field to given value.

### HasContents

`func (o *AnsibleCollectionVersionResponse) HasContents() bool`

HasContents returns a boolean if a field has been set.

### GetDependencies

`func (o *AnsibleCollectionVersionResponse) GetDependencies() map[string]interface{}`

GetDependencies returns the Dependencies field if non-nil, zero value otherwise.

### GetDependenciesOk

`func (o *AnsibleCollectionVersionResponse) GetDependenciesOk() (*map[string]interface{}, bool)`

GetDependenciesOk returns a tuple with the Dependencies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependencies

`func (o *AnsibleCollectionVersionResponse) SetDependencies(v map[string]interface{})`

SetDependencies sets Dependencies field to given value.

### HasDependencies

`func (o *AnsibleCollectionVersionResponse) HasDependencies() bool`

HasDependencies returns a boolean if a field has been set.

### GetDescription

`func (o *AnsibleCollectionVersionResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AnsibleCollectionVersionResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AnsibleCollectionVersionResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AnsibleCollectionVersionResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDocsBlob

`func (o *AnsibleCollectionVersionResponse) GetDocsBlob() map[string]interface{}`

GetDocsBlob returns the DocsBlob field if non-nil, zero value otherwise.

### GetDocsBlobOk

`func (o *AnsibleCollectionVersionResponse) GetDocsBlobOk() (*map[string]interface{}, bool)`

GetDocsBlobOk returns a tuple with the DocsBlob field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocsBlob

`func (o *AnsibleCollectionVersionResponse) SetDocsBlob(v map[string]interface{})`

SetDocsBlob sets DocsBlob field to given value.

### HasDocsBlob

`func (o *AnsibleCollectionVersionResponse) HasDocsBlob() bool`

HasDocsBlob returns a boolean if a field has been set.

### GetManifest

`func (o *AnsibleCollectionVersionResponse) GetManifest() map[string]interface{}`

GetManifest returns the Manifest field if non-nil, zero value otherwise.

### GetManifestOk

`func (o *AnsibleCollectionVersionResponse) GetManifestOk() (*map[string]interface{}, bool)`

GetManifestOk returns a tuple with the Manifest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManifest

`func (o *AnsibleCollectionVersionResponse) SetManifest(v map[string]interface{})`

SetManifest sets Manifest field to given value.

### HasManifest

`func (o *AnsibleCollectionVersionResponse) HasManifest() bool`

HasManifest returns a boolean if a field has been set.

### GetFiles

`func (o *AnsibleCollectionVersionResponse) GetFiles() map[string]interface{}`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *AnsibleCollectionVersionResponse) GetFilesOk() (*map[string]interface{}, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *AnsibleCollectionVersionResponse) SetFiles(v map[string]interface{})`

SetFiles sets Files field to given value.

### HasFiles

`func (o *AnsibleCollectionVersionResponse) HasFiles() bool`

HasFiles returns a boolean if a field has been set.

### GetDocumentation

`func (o *AnsibleCollectionVersionResponse) GetDocumentation() string`

GetDocumentation returns the Documentation field if non-nil, zero value otherwise.

### GetDocumentationOk

`func (o *AnsibleCollectionVersionResponse) GetDocumentationOk() (*string, bool)`

GetDocumentationOk returns a tuple with the Documentation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentation

`func (o *AnsibleCollectionVersionResponse) SetDocumentation(v string)`

SetDocumentation sets Documentation field to given value.

### HasDocumentation

`func (o *AnsibleCollectionVersionResponse) HasDocumentation() bool`

HasDocumentation returns a boolean if a field has been set.

### GetHomepage

`func (o *AnsibleCollectionVersionResponse) GetHomepage() string`

GetHomepage returns the Homepage field if non-nil, zero value otherwise.

### GetHomepageOk

`func (o *AnsibleCollectionVersionResponse) GetHomepageOk() (*string, bool)`

GetHomepageOk returns a tuple with the Homepage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHomepage

`func (o *AnsibleCollectionVersionResponse) SetHomepage(v string)`

SetHomepage sets Homepage field to given value.

### HasHomepage

`func (o *AnsibleCollectionVersionResponse) HasHomepage() bool`

HasHomepage returns a boolean if a field has been set.

### GetIssues

`func (o *AnsibleCollectionVersionResponse) GetIssues() string`

GetIssues returns the Issues field if non-nil, zero value otherwise.

### GetIssuesOk

`func (o *AnsibleCollectionVersionResponse) GetIssuesOk() (*string, bool)`

GetIssuesOk returns a tuple with the Issues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssues

`func (o *AnsibleCollectionVersionResponse) SetIssues(v string)`

SetIssues sets Issues field to given value.

### HasIssues

`func (o *AnsibleCollectionVersionResponse) HasIssues() bool`

HasIssues returns a boolean if a field has been set.

### GetLicense

`func (o *AnsibleCollectionVersionResponse) GetLicense() []string`

GetLicense returns the License field if non-nil, zero value otherwise.

### GetLicenseOk

`func (o *AnsibleCollectionVersionResponse) GetLicenseOk() (*[]string, bool)`

GetLicenseOk returns a tuple with the License field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicense

`func (o *AnsibleCollectionVersionResponse) SetLicense(v []string)`

SetLicense sets License field to given value.

### HasLicense

`func (o *AnsibleCollectionVersionResponse) HasLicense() bool`

HasLicense returns a boolean if a field has been set.

### GetName

`func (o *AnsibleCollectionVersionResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AnsibleCollectionVersionResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AnsibleCollectionVersionResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AnsibleCollectionVersionResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNamespace

`func (o *AnsibleCollectionVersionResponse) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *AnsibleCollectionVersionResponse) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *AnsibleCollectionVersionResponse) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *AnsibleCollectionVersionResponse) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetOriginRepository

`func (o *AnsibleCollectionVersionResponse) GetOriginRepository() string`

GetOriginRepository returns the OriginRepository field if non-nil, zero value otherwise.

### GetOriginRepositoryOk

`func (o *AnsibleCollectionVersionResponse) GetOriginRepositoryOk() (*string, bool)`

GetOriginRepositoryOk returns a tuple with the OriginRepository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginRepository

`func (o *AnsibleCollectionVersionResponse) SetOriginRepository(v string)`

SetOriginRepository sets OriginRepository field to given value.

### HasOriginRepository

`func (o *AnsibleCollectionVersionResponse) HasOriginRepository() bool`

HasOriginRepository returns a boolean if a field has been set.

### GetTags

`func (o *AnsibleCollectionVersionResponse) GetTags() []AnsibleTagResponse`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *AnsibleCollectionVersionResponse) GetTagsOk() (*[]AnsibleTagResponse, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *AnsibleCollectionVersionResponse) SetTags(v []AnsibleTagResponse)`

SetTags sets Tags field to given value.

### HasTags

`func (o *AnsibleCollectionVersionResponse) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetVersion

`func (o *AnsibleCollectionVersionResponse) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *AnsibleCollectionVersionResponse) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *AnsibleCollectionVersionResponse) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *AnsibleCollectionVersionResponse) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetRequiresAnsible

`func (o *AnsibleCollectionVersionResponse) GetRequiresAnsible() string`

GetRequiresAnsible returns the RequiresAnsible field if non-nil, zero value otherwise.

### GetRequiresAnsibleOk

`func (o *AnsibleCollectionVersionResponse) GetRequiresAnsibleOk() (*string, bool)`

GetRequiresAnsibleOk returns a tuple with the RequiresAnsible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiresAnsible

`func (o *AnsibleCollectionVersionResponse) SetRequiresAnsible(v string)`

SetRequiresAnsible sets RequiresAnsible field to given value.

### HasRequiresAnsible

`func (o *AnsibleCollectionVersionResponse) HasRequiresAnsible() bool`

HasRequiresAnsible returns a boolean if a field has been set.

### SetRequiresAnsibleNil

`func (o *AnsibleCollectionVersionResponse) SetRequiresAnsibleNil(b bool)`

 SetRequiresAnsibleNil sets the value for RequiresAnsible to be an explicit nil

### UnsetRequiresAnsible
`func (o *AnsibleCollectionVersionResponse) UnsetRequiresAnsible()`

UnsetRequiresAnsible ensures that no value is present for RequiresAnsible, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


