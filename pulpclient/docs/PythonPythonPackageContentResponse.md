# PythonPythonPackageContentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PulpHref** | Pointer to **string** |  | [optional] [readonly] 
**PulpCreated** | Pointer to **time.Time** | Timestamp of creation. | [optional] [readonly] 
**Artifact** | Pointer to **string** | Artifact file representing the physical content | [optional] 
**Filename** | Pointer to **string** | The name of the distribution package, usually of the format: {distribution}-{version}(-{build tag})?-{python tag}-{abi tag}-{platform tag}.{packagetype} | [optional] [readonly] 
**Packagetype** | Pointer to **string** | The type of the distribution package (e.g. sdist, bdist_wheel, bdist_egg, etc) | [optional] [readonly] 
**Name** | Pointer to **string** | The name of the python project. | [optional] [readonly] 
**Version** | Pointer to **string** | The packages version number. | [optional] [readonly] 
**Sha256** | Pointer to **string** | The SHA256 digest of this package. | [optional] [default to ""]
**MetadataVersion** | Pointer to **string** | Version of the file format | [optional] [readonly] 
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

### NewPythonPythonPackageContentResponse

`func NewPythonPythonPackageContentResponse() *PythonPythonPackageContentResponse`

NewPythonPythonPackageContentResponse instantiates a new PythonPythonPackageContentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPythonPythonPackageContentResponseWithDefaults

`func NewPythonPythonPackageContentResponseWithDefaults() *PythonPythonPackageContentResponse`

NewPythonPythonPackageContentResponseWithDefaults instantiates a new PythonPythonPackageContentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPulpHref

`func (o *PythonPythonPackageContentResponse) GetPulpHref() string`

GetPulpHref returns the PulpHref field if non-nil, zero value otherwise.

### GetPulpHrefOk

`func (o *PythonPythonPackageContentResponse) GetPulpHrefOk() (*string, bool)`

GetPulpHrefOk returns a tuple with the PulpHref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpHref

`func (o *PythonPythonPackageContentResponse) SetPulpHref(v string)`

SetPulpHref sets PulpHref field to given value.

### HasPulpHref

`func (o *PythonPythonPackageContentResponse) HasPulpHref() bool`

HasPulpHref returns a boolean if a field has been set.

### GetPulpCreated

`func (o *PythonPythonPackageContentResponse) GetPulpCreated() time.Time`

GetPulpCreated returns the PulpCreated field if non-nil, zero value otherwise.

### GetPulpCreatedOk

`func (o *PythonPythonPackageContentResponse) GetPulpCreatedOk() (*time.Time, bool)`

GetPulpCreatedOk returns a tuple with the PulpCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpCreated

`func (o *PythonPythonPackageContentResponse) SetPulpCreated(v time.Time)`

SetPulpCreated sets PulpCreated field to given value.

### HasPulpCreated

`func (o *PythonPythonPackageContentResponse) HasPulpCreated() bool`

HasPulpCreated returns a boolean if a field has been set.

### GetArtifact

`func (o *PythonPythonPackageContentResponse) GetArtifact() string`

GetArtifact returns the Artifact field if non-nil, zero value otherwise.

### GetArtifactOk

`func (o *PythonPythonPackageContentResponse) GetArtifactOk() (*string, bool)`

GetArtifactOk returns a tuple with the Artifact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifact

`func (o *PythonPythonPackageContentResponse) SetArtifact(v string)`

SetArtifact sets Artifact field to given value.

### HasArtifact

`func (o *PythonPythonPackageContentResponse) HasArtifact() bool`

HasArtifact returns a boolean if a field has been set.

### GetFilename

`func (o *PythonPythonPackageContentResponse) GetFilename() string`

GetFilename returns the Filename field if non-nil, zero value otherwise.

### GetFilenameOk

`func (o *PythonPythonPackageContentResponse) GetFilenameOk() (*string, bool)`

GetFilenameOk returns a tuple with the Filename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilename

`func (o *PythonPythonPackageContentResponse) SetFilename(v string)`

SetFilename sets Filename field to given value.

### HasFilename

`func (o *PythonPythonPackageContentResponse) HasFilename() bool`

HasFilename returns a boolean if a field has been set.

### GetPackagetype

`func (o *PythonPythonPackageContentResponse) GetPackagetype() string`

GetPackagetype returns the Packagetype field if non-nil, zero value otherwise.

### GetPackagetypeOk

`func (o *PythonPythonPackageContentResponse) GetPackagetypeOk() (*string, bool)`

GetPackagetypeOk returns a tuple with the Packagetype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackagetype

`func (o *PythonPythonPackageContentResponse) SetPackagetype(v string)`

SetPackagetype sets Packagetype field to given value.

### HasPackagetype

`func (o *PythonPythonPackageContentResponse) HasPackagetype() bool`

HasPackagetype returns a boolean if a field has been set.

### GetName

`func (o *PythonPythonPackageContentResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PythonPythonPackageContentResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PythonPythonPackageContentResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PythonPythonPackageContentResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetVersion

`func (o *PythonPythonPackageContentResponse) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *PythonPythonPackageContentResponse) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *PythonPythonPackageContentResponse) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *PythonPythonPackageContentResponse) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetSha256

`func (o *PythonPythonPackageContentResponse) GetSha256() string`

GetSha256 returns the Sha256 field if non-nil, zero value otherwise.

### GetSha256Ok

`func (o *PythonPythonPackageContentResponse) GetSha256Ok() (*string, bool)`

GetSha256Ok returns a tuple with the Sha256 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha256

`func (o *PythonPythonPackageContentResponse) SetSha256(v string)`

SetSha256 sets Sha256 field to given value.

### HasSha256

`func (o *PythonPythonPackageContentResponse) HasSha256() bool`

HasSha256 returns a boolean if a field has been set.

### GetMetadataVersion

`func (o *PythonPythonPackageContentResponse) GetMetadataVersion() string`

GetMetadataVersion returns the MetadataVersion field if non-nil, zero value otherwise.

### GetMetadataVersionOk

`func (o *PythonPythonPackageContentResponse) GetMetadataVersionOk() (*string, bool)`

GetMetadataVersionOk returns a tuple with the MetadataVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataVersion

`func (o *PythonPythonPackageContentResponse) SetMetadataVersion(v string)`

SetMetadataVersion sets MetadataVersion field to given value.

### HasMetadataVersion

`func (o *PythonPythonPackageContentResponse) HasMetadataVersion() bool`

HasMetadataVersion returns a boolean if a field has been set.

### GetSummary

`func (o *PythonPythonPackageContentResponse) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *PythonPythonPackageContentResponse) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *PythonPythonPackageContentResponse) SetSummary(v string)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *PythonPythonPackageContentResponse) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### GetDescription

`func (o *PythonPythonPackageContentResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PythonPythonPackageContentResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PythonPythonPackageContentResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PythonPythonPackageContentResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDescriptionContentType

`func (o *PythonPythonPackageContentResponse) GetDescriptionContentType() string`

GetDescriptionContentType returns the DescriptionContentType field if non-nil, zero value otherwise.

### GetDescriptionContentTypeOk

`func (o *PythonPythonPackageContentResponse) GetDescriptionContentTypeOk() (*string, bool)`

GetDescriptionContentTypeOk returns a tuple with the DescriptionContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptionContentType

`func (o *PythonPythonPackageContentResponse) SetDescriptionContentType(v string)`

SetDescriptionContentType sets DescriptionContentType field to given value.

### HasDescriptionContentType

`func (o *PythonPythonPackageContentResponse) HasDescriptionContentType() bool`

HasDescriptionContentType returns a boolean if a field has been set.

### GetKeywords

`func (o *PythonPythonPackageContentResponse) GetKeywords() string`

GetKeywords returns the Keywords field if non-nil, zero value otherwise.

### GetKeywordsOk

`func (o *PythonPythonPackageContentResponse) GetKeywordsOk() (*string, bool)`

GetKeywordsOk returns a tuple with the Keywords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeywords

`func (o *PythonPythonPackageContentResponse) SetKeywords(v string)`

SetKeywords sets Keywords field to given value.

### HasKeywords

`func (o *PythonPythonPackageContentResponse) HasKeywords() bool`

HasKeywords returns a boolean if a field has been set.

### GetHomePage

`func (o *PythonPythonPackageContentResponse) GetHomePage() string`

GetHomePage returns the HomePage field if non-nil, zero value otherwise.

### GetHomePageOk

`func (o *PythonPythonPackageContentResponse) GetHomePageOk() (*string, bool)`

GetHomePageOk returns a tuple with the HomePage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHomePage

`func (o *PythonPythonPackageContentResponse) SetHomePage(v string)`

SetHomePage sets HomePage field to given value.

### HasHomePage

`func (o *PythonPythonPackageContentResponse) HasHomePage() bool`

HasHomePage returns a boolean if a field has been set.

### GetDownloadUrl

`func (o *PythonPythonPackageContentResponse) GetDownloadUrl() string`

GetDownloadUrl returns the DownloadUrl field if non-nil, zero value otherwise.

### GetDownloadUrlOk

`func (o *PythonPythonPackageContentResponse) GetDownloadUrlOk() (*string, bool)`

GetDownloadUrlOk returns a tuple with the DownloadUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadUrl

`func (o *PythonPythonPackageContentResponse) SetDownloadUrl(v string)`

SetDownloadUrl sets DownloadUrl field to given value.

### HasDownloadUrl

`func (o *PythonPythonPackageContentResponse) HasDownloadUrl() bool`

HasDownloadUrl returns a boolean if a field has been set.

### GetAuthor

`func (o *PythonPythonPackageContentResponse) GetAuthor() string`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *PythonPythonPackageContentResponse) GetAuthorOk() (*string, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *PythonPythonPackageContentResponse) SetAuthor(v string)`

SetAuthor sets Author field to given value.

### HasAuthor

`func (o *PythonPythonPackageContentResponse) HasAuthor() bool`

HasAuthor returns a boolean if a field has been set.

### GetAuthorEmail

`func (o *PythonPythonPackageContentResponse) GetAuthorEmail() string`

GetAuthorEmail returns the AuthorEmail field if non-nil, zero value otherwise.

### GetAuthorEmailOk

`func (o *PythonPythonPackageContentResponse) GetAuthorEmailOk() (*string, bool)`

GetAuthorEmailOk returns a tuple with the AuthorEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorEmail

`func (o *PythonPythonPackageContentResponse) SetAuthorEmail(v string)`

SetAuthorEmail sets AuthorEmail field to given value.

### HasAuthorEmail

`func (o *PythonPythonPackageContentResponse) HasAuthorEmail() bool`

HasAuthorEmail returns a boolean if a field has been set.

### GetMaintainer

`func (o *PythonPythonPackageContentResponse) GetMaintainer() string`

GetMaintainer returns the Maintainer field if non-nil, zero value otherwise.

### GetMaintainerOk

`func (o *PythonPythonPackageContentResponse) GetMaintainerOk() (*string, bool)`

GetMaintainerOk returns a tuple with the Maintainer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintainer

`func (o *PythonPythonPackageContentResponse) SetMaintainer(v string)`

SetMaintainer sets Maintainer field to given value.

### HasMaintainer

`func (o *PythonPythonPackageContentResponse) HasMaintainer() bool`

HasMaintainer returns a boolean if a field has been set.

### GetMaintainerEmail

`func (o *PythonPythonPackageContentResponse) GetMaintainerEmail() string`

GetMaintainerEmail returns the MaintainerEmail field if non-nil, zero value otherwise.

### GetMaintainerEmailOk

`func (o *PythonPythonPackageContentResponse) GetMaintainerEmailOk() (*string, bool)`

GetMaintainerEmailOk returns a tuple with the MaintainerEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintainerEmail

`func (o *PythonPythonPackageContentResponse) SetMaintainerEmail(v string)`

SetMaintainerEmail sets MaintainerEmail field to given value.

### HasMaintainerEmail

`func (o *PythonPythonPackageContentResponse) HasMaintainerEmail() bool`

HasMaintainerEmail returns a boolean if a field has been set.

### GetLicense

`func (o *PythonPythonPackageContentResponse) GetLicense() string`

GetLicense returns the License field if non-nil, zero value otherwise.

### GetLicenseOk

`func (o *PythonPythonPackageContentResponse) GetLicenseOk() (*string, bool)`

GetLicenseOk returns a tuple with the License field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicense

`func (o *PythonPythonPackageContentResponse) SetLicense(v string)`

SetLicense sets License field to given value.

### HasLicense

`func (o *PythonPythonPackageContentResponse) HasLicense() bool`

HasLicense returns a boolean if a field has been set.

### GetRequiresPython

`func (o *PythonPythonPackageContentResponse) GetRequiresPython() string`

GetRequiresPython returns the RequiresPython field if non-nil, zero value otherwise.

### GetRequiresPythonOk

`func (o *PythonPythonPackageContentResponse) GetRequiresPythonOk() (*string, bool)`

GetRequiresPythonOk returns a tuple with the RequiresPython field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiresPython

`func (o *PythonPythonPackageContentResponse) SetRequiresPython(v string)`

SetRequiresPython sets RequiresPython field to given value.

### HasRequiresPython

`func (o *PythonPythonPackageContentResponse) HasRequiresPython() bool`

HasRequiresPython returns a boolean if a field has been set.

### GetProjectUrl

`func (o *PythonPythonPackageContentResponse) GetProjectUrl() string`

GetProjectUrl returns the ProjectUrl field if non-nil, zero value otherwise.

### GetProjectUrlOk

`func (o *PythonPythonPackageContentResponse) GetProjectUrlOk() (*string, bool)`

GetProjectUrlOk returns a tuple with the ProjectUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectUrl

`func (o *PythonPythonPackageContentResponse) SetProjectUrl(v string)`

SetProjectUrl sets ProjectUrl field to given value.

### HasProjectUrl

`func (o *PythonPythonPackageContentResponse) HasProjectUrl() bool`

HasProjectUrl returns a boolean if a field has been set.

### GetProjectUrls

`func (o *PythonPythonPackageContentResponse) GetProjectUrls() map[string]interface{}`

GetProjectUrls returns the ProjectUrls field if non-nil, zero value otherwise.

### GetProjectUrlsOk

`func (o *PythonPythonPackageContentResponse) GetProjectUrlsOk() (*map[string]interface{}, bool)`

GetProjectUrlsOk returns a tuple with the ProjectUrls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectUrls

`func (o *PythonPythonPackageContentResponse) SetProjectUrls(v map[string]interface{})`

SetProjectUrls sets ProjectUrls field to given value.

### HasProjectUrls

`func (o *PythonPythonPackageContentResponse) HasProjectUrls() bool`

HasProjectUrls returns a boolean if a field has been set.

### GetPlatform

`func (o *PythonPythonPackageContentResponse) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *PythonPythonPackageContentResponse) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *PythonPythonPackageContentResponse) SetPlatform(v string)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *PythonPythonPackageContentResponse) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### GetSupportedPlatform

`func (o *PythonPythonPackageContentResponse) GetSupportedPlatform() string`

GetSupportedPlatform returns the SupportedPlatform field if non-nil, zero value otherwise.

### GetSupportedPlatformOk

`func (o *PythonPythonPackageContentResponse) GetSupportedPlatformOk() (*string, bool)`

GetSupportedPlatformOk returns a tuple with the SupportedPlatform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedPlatform

`func (o *PythonPythonPackageContentResponse) SetSupportedPlatform(v string)`

SetSupportedPlatform sets SupportedPlatform field to given value.

### HasSupportedPlatform

`func (o *PythonPythonPackageContentResponse) HasSupportedPlatform() bool`

HasSupportedPlatform returns a boolean if a field has been set.

### GetRequiresDist

`func (o *PythonPythonPackageContentResponse) GetRequiresDist() map[string]interface{}`

GetRequiresDist returns the RequiresDist field if non-nil, zero value otherwise.

### GetRequiresDistOk

`func (o *PythonPythonPackageContentResponse) GetRequiresDistOk() (*map[string]interface{}, bool)`

GetRequiresDistOk returns a tuple with the RequiresDist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiresDist

`func (o *PythonPythonPackageContentResponse) SetRequiresDist(v map[string]interface{})`

SetRequiresDist sets RequiresDist field to given value.

### HasRequiresDist

`func (o *PythonPythonPackageContentResponse) HasRequiresDist() bool`

HasRequiresDist returns a boolean if a field has been set.

### GetProvidesDist

`func (o *PythonPythonPackageContentResponse) GetProvidesDist() map[string]interface{}`

GetProvidesDist returns the ProvidesDist field if non-nil, zero value otherwise.

### GetProvidesDistOk

`func (o *PythonPythonPackageContentResponse) GetProvidesDistOk() (*map[string]interface{}, bool)`

GetProvidesDistOk returns a tuple with the ProvidesDist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvidesDist

`func (o *PythonPythonPackageContentResponse) SetProvidesDist(v map[string]interface{})`

SetProvidesDist sets ProvidesDist field to given value.

### HasProvidesDist

`func (o *PythonPythonPackageContentResponse) HasProvidesDist() bool`

HasProvidesDist returns a boolean if a field has been set.

### GetObsoletesDist

`func (o *PythonPythonPackageContentResponse) GetObsoletesDist() map[string]interface{}`

GetObsoletesDist returns the ObsoletesDist field if non-nil, zero value otherwise.

### GetObsoletesDistOk

`func (o *PythonPythonPackageContentResponse) GetObsoletesDistOk() (*map[string]interface{}, bool)`

GetObsoletesDistOk returns a tuple with the ObsoletesDist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObsoletesDist

`func (o *PythonPythonPackageContentResponse) SetObsoletesDist(v map[string]interface{})`

SetObsoletesDist sets ObsoletesDist field to given value.

### HasObsoletesDist

`func (o *PythonPythonPackageContentResponse) HasObsoletesDist() bool`

HasObsoletesDist returns a boolean if a field has been set.

### GetRequiresExternal

`func (o *PythonPythonPackageContentResponse) GetRequiresExternal() map[string]interface{}`

GetRequiresExternal returns the RequiresExternal field if non-nil, zero value otherwise.

### GetRequiresExternalOk

`func (o *PythonPythonPackageContentResponse) GetRequiresExternalOk() (*map[string]interface{}, bool)`

GetRequiresExternalOk returns a tuple with the RequiresExternal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiresExternal

`func (o *PythonPythonPackageContentResponse) SetRequiresExternal(v map[string]interface{})`

SetRequiresExternal sets RequiresExternal field to given value.

### HasRequiresExternal

`func (o *PythonPythonPackageContentResponse) HasRequiresExternal() bool`

HasRequiresExternal returns a boolean if a field has been set.

### GetClassifiers

`func (o *PythonPythonPackageContentResponse) GetClassifiers() map[string]interface{}`

GetClassifiers returns the Classifiers field if non-nil, zero value otherwise.

### GetClassifiersOk

`func (o *PythonPythonPackageContentResponse) GetClassifiersOk() (*map[string]interface{}, bool)`

GetClassifiersOk returns a tuple with the Classifiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassifiers

`func (o *PythonPythonPackageContentResponse) SetClassifiers(v map[string]interface{})`

SetClassifiers sets Classifiers field to given value.

### HasClassifiers

`func (o *PythonPythonPackageContentResponse) HasClassifiers() bool`

HasClassifiers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


