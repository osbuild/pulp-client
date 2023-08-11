# PythonPythonPackageContent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Repository** | Pointer to **string** | A URI of a repository the new content unit should be associated with. | [optional] 
**Artifact** | Pointer to **string** | Artifact file representing the physical content | [optional] 
**RelativePath** | **string** | Path where the artifact is located relative to distributions base_path | 
**File** | Pointer to ***os.File** | An uploaded file that may be turned into the artifact of the content unit. | [optional] 
**Upload** | Pointer to **string** | An uncommitted upload that may be turned into the artifact of the content unit. | [optional] 
**Sha256** | Pointer to **string** | The SHA256 digest of this package. | [optional] [default to ""]
**Summary** | Pointer to **string** | A one-line summary of what the package does. | [optional] 
**Description** | Pointer to **string** | A longer description of the package that can run to several paragraphs. | [optional] 
**DescriptionContentType** | Pointer to **string** | A string stating the markup syntax (if any) used in the distribution’s description, so that tools can intelligently render the description. | [optional] 
**Keywords** | Pointer to **string** | Additional keywords to be used to assist searching for the package in a larger catalog. | [optional] 
**HomePage** | Pointer to **string** | The URL for the package&#39;s home page. | [optional] 
**DownloadUrl** | Pointer to **string** | Legacy field denoting the URL from which this package can be downloaded. | [optional] 
**Author** | Pointer to **string** | Text containing the author&#39;s name. Contact information can also be added, separated with newlines. | [optional] 
**AuthorEmail** | Pointer to **string** | The author&#39;s e-mail address.  | [optional] 
**Maintainer** | Pointer to **string** | The maintainer&#39;s name at a minimum; additional contact information may be provided. | [optional] 
**MaintainerEmail** | Pointer to **string** | The maintainer&#39;s e-mail address. | [optional] 
**License** | Pointer to **string** | Text indicating the license covering the distribution | [optional] 
**RequiresPython** | Pointer to **string** | The Python version(s) that the distribution is guaranteed to be compatible with. | [optional] 
**ProjectUrl** | Pointer to **string** | A browsable URL for the project and a label for it, separated by a comma. | [optional] 
**ProjectUrls** | Pointer to **map[string]interface{}** | A dictionary of labels and URLs for the project. | [optional] 
**Platform** | Pointer to **string** | A comma-separated list of platform specifications, summarizing the operating systems supported by the package. | [optional] 
**SupportedPlatform** | Pointer to **string** | Field to specify the OS and CPU for which the binary package was compiled.  | [optional] 
**RequiresDist** | Pointer to **map[string]interface{}** | A JSON list containing names of some other distutils project required by this distribution. | [optional] 
**ProvidesDist** | Pointer to **map[string]interface{}** | A JSON list containing names of a Distutils project which is contained within this distribution. | [optional] 
**ObsoletesDist** | Pointer to **map[string]interface{}** | A JSON list containing names of a distutils project&#39;s distribution which this distribution renders obsolete, meaning that the two projects should not be installed at the same time. | [optional] 
**RequiresExternal** | Pointer to **map[string]interface{}** | A JSON list containing some dependency in the system that the distribution is to be used. | [optional] 
**Classifiers** | Pointer to **map[string]interface{}** | A JSON list containing classification values for a Python package. | [optional] 

## Methods

### NewPythonPythonPackageContent

`func NewPythonPythonPackageContent(relativePath string, ) *PythonPythonPackageContent`

NewPythonPythonPackageContent instantiates a new PythonPythonPackageContent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPythonPythonPackageContentWithDefaults

`func NewPythonPythonPackageContentWithDefaults() *PythonPythonPackageContent`

NewPythonPythonPackageContentWithDefaults instantiates a new PythonPythonPackageContent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRepository

`func (o *PythonPythonPackageContent) GetRepository() string`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *PythonPythonPackageContent) GetRepositoryOk() (*string, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *PythonPythonPackageContent) SetRepository(v string)`

SetRepository sets Repository field to given value.

### HasRepository

`func (o *PythonPythonPackageContent) HasRepository() bool`

HasRepository returns a boolean if a field has been set.

### GetArtifact

`func (o *PythonPythonPackageContent) GetArtifact() string`

GetArtifact returns the Artifact field if non-nil, zero value otherwise.

### GetArtifactOk

`func (o *PythonPythonPackageContent) GetArtifactOk() (*string, bool)`

GetArtifactOk returns a tuple with the Artifact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifact

`func (o *PythonPythonPackageContent) SetArtifact(v string)`

SetArtifact sets Artifact field to given value.

### HasArtifact

`func (o *PythonPythonPackageContent) HasArtifact() bool`

HasArtifact returns a boolean if a field has been set.

### GetRelativePath

`func (o *PythonPythonPackageContent) GetRelativePath() string`

GetRelativePath returns the RelativePath field if non-nil, zero value otherwise.

### GetRelativePathOk

`func (o *PythonPythonPackageContent) GetRelativePathOk() (*string, bool)`

GetRelativePathOk returns a tuple with the RelativePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelativePath

`func (o *PythonPythonPackageContent) SetRelativePath(v string)`

SetRelativePath sets RelativePath field to given value.


### GetFile

`func (o *PythonPythonPackageContent) GetFile() *os.File`

GetFile returns the File field if non-nil, zero value otherwise.

### GetFileOk

`func (o *PythonPythonPackageContent) GetFileOk() (**os.File, bool)`

GetFileOk returns a tuple with the File field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFile

`func (o *PythonPythonPackageContent) SetFile(v *os.File)`

SetFile sets File field to given value.

### HasFile

`func (o *PythonPythonPackageContent) HasFile() bool`

HasFile returns a boolean if a field has been set.

### GetUpload

`func (o *PythonPythonPackageContent) GetUpload() string`

GetUpload returns the Upload field if non-nil, zero value otherwise.

### GetUploadOk

`func (o *PythonPythonPackageContent) GetUploadOk() (*string, bool)`

GetUploadOk returns a tuple with the Upload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpload

`func (o *PythonPythonPackageContent) SetUpload(v string)`

SetUpload sets Upload field to given value.

### HasUpload

`func (o *PythonPythonPackageContent) HasUpload() bool`

HasUpload returns a boolean if a field has been set.

### GetSha256

`func (o *PythonPythonPackageContent) GetSha256() string`

GetSha256 returns the Sha256 field if non-nil, zero value otherwise.

### GetSha256Ok

`func (o *PythonPythonPackageContent) GetSha256Ok() (*string, bool)`

GetSha256Ok returns a tuple with the Sha256 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha256

`func (o *PythonPythonPackageContent) SetSha256(v string)`

SetSha256 sets Sha256 field to given value.

### HasSha256

`func (o *PythonPythonPackageContent) HasSha256() bool`

HasSha256 returns a boolean if a field has been set.

### GetSummary

`func (o *PythonPythonPackageContent) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *PythonPythonPackageContent) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *PythonPythonPackageContent) SetSummary(v string)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *PythonPythonPackageContent) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### GetDescription

`func (o *PythonPythonPackageContent) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PythonPythonPackageContent) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PythonPythonPackageContent) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PythonPythonPackageContent) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDescriptionContentType

`func (o *PythonPythonPackageContent) GetDescriptionContentType() string`

GetDescriptionContentType returns the DescriptionContentType field if non-nil, zero value otherwise.

### GetDescriptionContentTypeOk

`func (o *PythonPythonPackageContent) GetDescriptionContentTypeOk() (*string, bool)`

GetDescriptionContentTypeOk returns a tuple with the DescriptionContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptionContentType

`func (o *PythonPythonPackageContent) SetDescriptionContentType(v string)`

SetDescriptionContentType sets DescriptionContentType field to given value.

### HasDescriptionContentType

`func (o *PythonPythonPackageContent) HasDescriptionContentType() bool`

HasDescriptionContentType returns a boolean if a field has been set.

### GetKeywords

`func (o *PythonPythonPackageContent) GetKeywords() string`

GetKeywords returns the Keywords field if non-nil, zero value otherwise.

### GetKeywordsOk

`func (o *PythonPythonPackageContent) GetKeywordsOk() (*string, bool)`

GetKeywordsOk returns a tuple with the Keywords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeywords

`func (o *PythonPythonPackageContent) SetKeywords(v string)`

SetKeywords sets Keywords field to given value.

### HasKeywords

`func (o *PythonPythonPackageContent) HasKeywords() bool`

HasKeywords returns a boolean if a field has been set.

### GetHomePage

`func (o *PythonPythonPackageContent) GetHomePage() string`

GetHomePage returns the HomePage field if non-nil, zero value otherwise.

### GetHomePageOk

`func (o *PythonPythonPackageContent) GetHomePageOk() (*string, bool)`

GetHomePageOk returns a tuple with the HomePage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHomePage

`func (o *PythonPythonPackageContent) SetHomePage(v string)`

SetHomePage sets HomePage field to given value.

### HasHomePage

`func (o *PythonPythonPackageContent) HasHomePage() bool`

HasHomePage returns a boolean if a field has been set.

### GetDownloadUrl

`func (o *PythonPythonPackageContent) GetDownloadUrl() string`

GetDownloadUrl returns the DownloadUrl field if non-nil, zero value otherwise.

### GetDownloadUrlOk

`func (o *PythonPythonPackageContent) GetDownloadUrlOk() (*string, bool)`

GetDownloadUrlOk returns a tuple with the DownloadUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadUrl

`func (o *PythonPythonPackageContent) SetDownloadUrl(v string)`

SetDownloadUrl sets DownloadUrl field to given value.

### HasDownloadUrl

`func (o *PythonPythonPackageContent) HasDownloadUrl() bool`

HasDownloadUrl returns a boolean if a field has been set.

### GetAuthor

`func (o *PythonPythonPackageContent) GetAuthor() string`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *PythonPythonPackageContent) GetAuthorOk() (*string, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *PythonPythonPackageContent) SetAuthor(v string)`

SetAuthor sets Author field to given value.

### HasAuthor

`func (o *PythonPythonPackageContent) HasAuthor() bool`

HasAuthor returns a boolean if a field has been set.

### GetAuthorEmail

`func (o *PythonPythonPackageContent) GetAuthorEmail() string`

GetAuthorEmail returns the AuthorEmail field if non-nil, zero value otherwise.

### GetAuthorEmailOk

`func (o *PythonPythonPackageContent) GetAuthorEmailOk() (*string, bool)`

GetAuthorEmailOk returns a tuple with the AuthorEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorEmail

`func (o *PythonPythonPackageContent) SetAuthorEmail(v string)`

SetAuthorEmail sets AuthorEmail field to given value.

### HasAuthorEmail

`func (o *PythonPythonPackageContent) HasAuthorEmail() bool`

HasAuthorEmail returns a boolean if a field has been set.

### GetMaintainer

`func (o *PythonPythonPackageContent) GetMaintainer() string`

GetMaintainer returns the Maintainer field if non-nil, zero value otherwise.

### GetMaintainerOk

`func (o *PythonPythonPackageContent) GetMaintainerOk() (*string, bool)`

GetMaintainerOk returns a tuple with the Maintainer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintainer

`func (o *PythonPythonPackageContent) SetMaintainer(v string)`

SetMaintainer sets Maintainer field to given value.

### HasMaintainer

`func (o *PythonPythonPackageContent) HasMaintainer() bool`

HasMaintainer returns a boolean if a field has been set.

### GetMaintainerEmail

`func (o *PythonPythonPackageContent) GetMaintainerEmail() string`

GetMaintainerEmail returns the MaintainerEmail field if non-nil, zero value otherwise.

### GetMaintainerEmailOk

`func (o *PythonPythonPackageContent) GetMaintainerEmailOk() (*string, bool)`

GetMaintainerEmailOk returns a tuple with the MaintainerEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintainerEmail

`func (o *PythonPythonPackageContent) SetMaintainerEmail(v string)`

SetMaintainerEmail sets MaintainerEmail field to given value.

### HasMaintainerEmail

`func (o *PythonPythonPackageContent) HasMaintainerEmail() bool`

HasMaintainerEmail returns a boolean if a field has been set.

### GetLicense

`func (o *PythonPythonPackageContent) GetLicense() string`

GetLicense returns the License field if non-nil, zero value otherwise.

### GetLicenseOk

`func (o *PythonPythonPackageContent) GetLicenseOk() (*string, bool)`

GetLicenseOk returns a tuple with the License field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicense

`func (o *PythonPythonPackageContent) SetLicense(v string)`

SetLicense sets License field to given value.

### HasLicense

`func (o *PythonPythonPackageContent) HasLicense() bool`

HasLicense returns a boolean if a field has been set.

### GetRequiresPython

`func (o *PythonPythonPackageContent) GetRequiresPython() string`

GetRequiresPython returns the RequiresPython field if non-nil, zero value otherwise.

### GetRequiresPythonOk

`func (o *PythonPythonPackageContent) GetRequiresPythonOk() (*string, bool)`

GetRequiresPythonOk returns a tuple with the RequiresPython field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiresPython

`func (o *PythonPythonPackageContent) SetRequiresPython(v string)`

SetRequiresPython sets RequiresPython field to given value.

### HasRequiresPython

`func (o *PythonPythonPackageContent) HasRequiresPython() bool`

HasRequiresPython returns a boolean if a field has been set.

### GetProjectUrl

`func (o *PythonPythonPackageContent) GetProjectUrl() string`

GetProjectUrl returns the ProjectUrl field if non-nil, zero value otherwise.

### GetProjectUrlOk

`func (o *PythonPythonPackageContent) GetProjectUrlOk() (*string, bool)`

GetProjectUrlOk returns a tuple with the ProjectUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectUrl

`func (o *PythonPythonPackageContent) SetProjectUrl(v string)`

SetProjectUrl sets ProjectUrl field to given value.

### HasProjectUrl

`func (o *PythonPythonPackageContent) HasProjectUrl() bool`

HasProjectUrl returns a boolean if a field has been set.

### GetProjectUrls

`func (o *PythonPythonPackageContent) GetProjectUrls() map[string]interface{}`

GetProjectUrls returns the ProjectUrls field if non-nil, zero value otherwise.

### GetProjectUrlsOk

`func (o *PythonPythonPackageContent) GetProjectUrlsOk() (*map[string]interface{}, bool)`

GetProjectUrlsOk returns a tuple with the ProjectUrls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectUrls

`func (o *PythonPythonPackageContent) SetProjectUrls(v map[string]interface{})`

SetProjectUrls sets ProjectUrls field to given value.

### HasProjectUrls

`func (o *PythonPythonPackageContent) HasProjectUrls() bool`

HasProjectUrls returns a boolean if a field has been set.

### GetPlatform

`func (o *PythonPythonPackageContent) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *PythonPythonPackageContent) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *PythonPythonPackageContent) SetPlatform(v string)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *PythonPythonPackageContent) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### GetSupportedPlatform

`func (o *PythonPythonPackageContent) GetSupportedPlatform() string`

GetSupportedPlatform returns the SupportedPlatform field if non-nil, zero value otherwise.

### GetSupportedPlatformOk

`func (o *PythonPythonPackageContent) GetSupportedPlatformOk() (*string, bool)`

GetSupportedPlatformOk returns a tuple with the SupportedPlatform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedPlatform

`func (o *PythonPythonPackageContent) SetSupportedPlatform(v string)`

SetSupportedPlatform sets SupportedPlatform field to given value.

### HasSupportedPlatform

`func (o *PythonPythonPackageContent) HasSupportedPlatform() bool`

HasSupportedPlatform returns a boolean if a field has been set.

### GetRequiresDist

`func (o *PythonPythonPackageContent) GetRequiresDist() map[string]interface{}`

GetRequiresDist returns the RequiresDist field if non-nil, zero value otherwise.

### GetRequiresDistOk

`func (o *PythonPythonPackageContent) GetRequiresDistOk() (*map[string]interface{}, bool)`

GetRequiresDistOk returns a tuple with the RequiresDist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiresDist

`func (o *PythonPythonPackageContent) SetRequiresDist(v map[string]interface{})`

SetRequiresDist sets RequiresDist field to given value.

### HasRequiresDist

`func (o *PythonPythonPackageContent) HasRequiresDist() bool`

HasRequiresDist returns a boolean if a field has been set.

### GetProvidesDist

`func (o *PythonPythonPackageContent) GetProvidesDist() map[string]interface{}`

GetProvidesDist returns the ProvidesDist field if non-nil, zero value otherwise.

### GetProvidesDistOk

`func (o *PythonPythonPackageContent) GetProvidesDistOk() (*map[string]interface{}, bool)`

GetProvidesDistOk returns a tuple with the ProvidesDist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvidesDist

`func (o *PythonPythonPackageContent) SetProvidesDist(v map[string]interface{})`

SetProvidesDist sets ProvidesDist field to given value.

### HasProvidesDist

`func (o *PythonPythonPackageContent) HasProvidesDist() bool`

HasProvidesDist returns a boolean if a field has been set.

### GetObsoletesDist

`func (o *PythonPythonPackageContent) GetObsoletesDist() map[string]interface{}`

GetObsoletesDist returns the ObsoletesDist field if non-nil, zero value otherwise.

### GetObsoletesDistOk

`func (o *PythonPythonPackageContent) GetObsoletesDistOk() (*map[string]interface{}, bool)`

GetObsoletesDistOk returns a tuple with the ObsoletesDist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObsoletesDist

`func (o *PythonPythonPackageContent) SetObsoletesDist(v map[string]interface{})`

SetObsoletesDist sets ObsoletesDist field to given value.

### HasObsoletesDist

`func (o *PythonPythonPackageContent) HasObsoletesDist() bool`

HasObsoletesDist returns a boolean if a field has been set.

### GetRequiresExternal

`func (o *PythonPythonPackageContent) GetRequiresExternal() map[string]interface{}`

GetRequiresExternal returns the RequiresExternal field if non-nil, zero value otherwise.

### GetRequiresExternalOk

`func (o *PythonPythonPackageContent) GetRequiresExternalOk() (*map[string]interface{}, bool)`

GetRequiresExternalOk returns a tuple with the RequiresExternal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiresExternal

`func (o *PythonPythonPackageContent) SetRequiresExternal(v map[string]interface{})`

SetRequiresExternal sets RequiresExternal field to given value.

### HasRequiresExternal

`func (o *PythonPythonPackageContent) HasRequiresExternal() bool`

HasRequiresExternal returns a boolean if a field has been set.

### GetClassifiers

`func (o *PythonPythonPackageContent) GetClassifiers() map[string]interface{}`

GetClassifiers returns the Classifiers field if non-nil, zero value otherwise.

### GetClassifiersOk

`func (o *PythonPythonPackageContent) GetClassifiersOk() (*map[string]interface{}, bool)`

GetClassifiersOk returns a tuple with the Classifiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassifiers

`func (o *PythonPythonPackageContent) SetClassifiers(v map[string]interface{})`

SetClassifiers sets Classifiers field to given value.

### HasClassifiers

`func (o *PythonPythonPackageContent) HasClassifiers() bool`

HasClassifiers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


