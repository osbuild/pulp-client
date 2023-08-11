# DebPackageResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PulpHref** | Pointer to **string** |  | [optional] [readonly] 
**PulpCreated** | Pointer to **time.Time** | Timestamp of creation. | [optional] [readonly] 
**Artifact** | Pointer to **string** | Artifact file representing the physical content | [optional] 
**RelativePath** | Pointer to **string** | Path where the artifact is located relative to distributions base_path | [optional] 
**Distribution** | Pointer to **string** | Name of the distribution. | [optional] 
**Component** | Pointer to **string** | Name of the component. | [optional] 
**Md5** | Pointer to **string** | The MD5 checksum if available. | [optional] [readonly] 
**Sha1** | Pointer to **string** | The SHA-1 checksum if available. | [optional] [readonly] 
**Sha224** | Pointer to **string** | The SHA-224 checksum if available. | [optional] [readonly] 
**Sha256** | Pointer to **string** | The SHA-256 checksum if available. | [optional] [readonly] 
**Sha384** | Pointer to **string** | The SHA-384 checksum if available. | [optional] [readonly] 
**Sha512** | Pointer to **string** | The SHA-512 checksum if available. | [optional] [readonly] 
**Package** | Pointer to **string** |  | [optional] [readonly] 
**Source** | Pointer to **string** |  | [optional] [readonly] 
**Version** | Pointer to **string** |  | [optional] [readonly] 
**Architecture** | Pointer to **string** |  | [optional] [readonly] 
**Section** | Pointer to **string** |  | [optional] [readonly] 
**Priority** | Pointer to **string** |  | [optional] [readonly] 
**Origin** | Pointer to **string** |  | [optional] [readonly] 
**Tag** | Pointer to **string** |  | [optional] [readonly] 
**Bugs** | Pointer to **string** |  | [optional] [readonly] 
**Essential** | Pointer to **string** |  | [optional] [readonly] 
**BuildEssential** | Pointer to **string** |  | [optional] [readonly] 
**InstalledSize** | Pointer to **string** |  | [optional] [readonly] 
**Maintainer** | Pointer to **string** |  | [optional] [readonly] 
**OriginalMaintainer** | Pointer to **string** |  | [optional] [readonly] 
**Description** | Pointer to **string** |  | [optional] [readonly] 
**DescriptionMd5** | Pointer to **string** |  | [optional] [readonly] 
**Homepage** | Pointer to **string** |  | [optional] [readonly] 
**BuiltUsing** | Pointer to **string** |  | [optional] [readonly] 
**AutoBuiltPackage** | Pointer to **string** |  | [optional] [readonly] 
**MultiArch** | Pointer to **string** |  | [optional] [readonly] 
**Breaks** | Pointer to **string** |  | [optional] [readonly] 
**Conflicts** | Pointer to **string** |  | [optional] [readonly] 
**Depends** | Pointer to **string** |  | [optional] [readonly] 
**Recommends** | Pointer to **string** |  | [optional] [readonly] 
**Suggests** | Pointer to **string** |  | [optional] [readonly] 
**Enhances** | Pointer to **string** |  | [optional] [readonly] 
**PreDepends** | Pointer to **string** |  | [optional] [readonly] 
**Provides** | Pointer to **string** |  | [optional] [readonly] 
**Replaces** | Pointer to **string** |  | [optional] [readonly] 

## Methods

### NewDebPackageResponse

`func NewDebPackageResponse() *DebPackageResponse`

NewDebPackageResponse instantiates a new DebPackageResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDebPackageResponseWithDefaults

`func NewDebPackageResponseWithDefaults() *DebPackageResponse`

NewDebPackageResponseWithDefaults instantiates a new DebPackageResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPulpHref

`func (o *DebPackageResponse) GetPulpHref() string`

GetPulpHref returns the PulpHref field if non-nil, zero value otherwise.

### GetPulpHrefOk

`func (o *DebPackageResponse) GetPulpHrefOk() (*string, bool)`

GetPulpHrefOk returns a tuple with the PulpHref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpHref

`func (o *DebPackageResponse) SetPulpHref(v string)`

SetPulpHref sets PulpHref field to given value.

### HasPulpHref

`func (o *DebPackageResponse) HasPulpHref() bool`

HasPulpHref returns a boolean if a field has been set.

### GetPulpCreated

`func (o *DebPackageResponse) GetPulpCreated() time.Time`

GetPulpCreated returns the PulpCreated field if non-nil, zero value otherwise.

### GetPulpCreatedOk

`func (o *DebPackageResponse) GetPulpCreatedOk() (*time.Time, bool)`

GetPulpCreatedOk returns a tuple with the PulpCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpCreated

`func (o *DebPackageResponse) SetPulpCreated(v time.Time)`

SetPulpCreated sets PulpCreated field to given value.

### HasPulpCreated

`func (o *DebPackageResponse) HasPulpCreated() bool`

HasPulpCreated returns a boolean if a field has been set.

### GetArtifact

`func (o *DebPackageResponse) GetArtifact() string`

GetArtifact returns the Artifact field if non-nil, zero value otherwise.

### GetArtifactOk

`func (o *DebPackageResponse) GetArtifactOk() (*string, bool)`

GetArtifactOk returns a tuple with the Artifact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifact

`func (o *DebPackageResponse) SetArtifact(v string)`

SetArtifact sets Artifact field to given value.

### HasArtifact

`func (o *DebPackageResponse) HasArtifact() bool`

HasArtifact returns a boolean if a field has been set.

### GetRelativePath

`func (o *DebPackageResponse) GetRelativePath() string`

GetRelativePath returns the RelativePath field if non-nil, zero value otherwise.

### GetRelativePathOk

`func (o *DebPackageResponse) GetRelativePathOk() (*string, bool)`

GetRelativePathOk returns a tuple with the RelativePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelativePath

`func (o *DebPackageResponse) SetRelativePath(v string)`

SetRelativePath sets RelativePath field to given value.

### HasRelativePath

`func (o *DebPackageResponse) HasRelativePath() bool`

HasRelativePath returns a boolean if a field has been set.

### GetDistribution

`func (o *DebPackageResponse) GetDistribution() string`

GetDistribution returns the Distribution field if non-nil, zero value otherwise.

### GetDistributionOk

`func (o *DebPackageResponse) GetDistributionOk() (*string, bool)`

GetDistributionOk returns a tuple with the Distribution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistribution

`func (o *DebPackageResponse) SetDistribution(v string)`

SetDistribution sets Distribution field to given value.

### HasDistribution

`func (o *DebPackageResponse) HasDistribution() bool`

HasDistribution returns a boolean if a field has been set.

### GetComponent

`func (o *DebPackageResponse) GetComponent() string`

GetComponent returns the Component field if non-nil, zero value otherwise.

### GetComponentOk

`func (o *DebPackageResponse) GetComponentOk() (*string, bool)`

GetComponentOk returns a tuple with the Component field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponent

`func (o *DebPackageResponse) SetComponent(v string)`

SetComponent sets Component field to given value.

### HasComponent

`func (o *DebPackageResponse) HasComponent() bool`

HasComponent returns a boolean if a field has been set.

### GetMd5

`func (o *DebPackageResponse) GetMd5() string`

GetMd5 returns the Md5 field if non-nil, zero value otherwise.

### GetMd5Ok

`func (o *DebPackageResponse) GetMd5Ok() (*string, bool)`

GetMd5Ok returns a tuple with the Md5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMd5

`func (o *DebPackageResponse) SetMd5(v string)`

SetMd5 sets Md5 field to given value.

### HasMd5

`func (o *DebPackageResponse) HasMd5() bool`

HasMd5 returns a boolean if a field has been set.

### GetSha1

`func (o *DebPackageResponse) GetSha1() string`

GetSha1 returns the Sha1 field if non-nil, zero value otherwise.

### GetSha1Ok

`func (o *DebPackageResponse) GetSha1Ok() (*string, bool)`

GetSha1Ok returns a tuple with the Sha1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha1

`func (o *DebPackageResponse) SetSha1(v string)`

SetSha1 sets Sha1 field to given value.

### HasSha1

`func (o *DebPackageResponse) HasSha1() bool`

HasSha1 returns a boolean if a field has been set.

### GetSha224

`func (o *DebPackageResponse) GetSha224() string`

GetSha224 returns the Sha224 field if non-nil, zero value otherwise.

### GetSha224Ok

`func (o *DebPackageResponse) GetSha224Ok() (*string, bool)`

GetSha224Ok returns a tuple with the Sha224 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha224

`func (o *DebPackageResponse) SetSha224(v string)`

SetSha224 sets Sha224 field to given value.

### HasSha224

`func (o *DebPackageResponse) HasSha224() bool`

HasSha224 returns a boolean if a field has been set.

### GetSha256

`func (o *DebPackageResponse) GetSha256() string`

GetSha256 returns the Sha256 field if non-nil, zero value otherwise.

### GetSha256Ok

`func (o *DebPackageResponse) GetSha256Ok() (*string, bool)`

GetSha256Ok returns a tuple with the Sha256 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha256

`func (o *DebPackageResponse) SetSha256(v string)`

SetSha256 sets Sha256 field to given value.

### HasSha256

`func (o *DebPackageResponse) HasSha256() bool`

HasSha256 returns a boolean if a field has been set.

### GetSha384

`func (o *DebPackageResponse) GetSha384() string`

GetSha384 returns the Sha384 field if non-nil, zero value otherwise.

### GetSha384Ok

`func (o *DebPackageResponse) GetSha384Ok() (*string, bool)`

GetSha384Ok returns a tuple with the Sha384 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha384

`func (o *DebPackageResponse) SetSha384(v string)`

SetSha384 sets Sha384 field to given value.

### HasSha384

`func (o *DebPackageResponse) HasSha384() bool`

HasSha384 returns a boolean if a field has been set.

### GetSha512

`func (o *DebPackageResponse) GetSha512() string`

GetSha512 returns the Sha512 field if non-nil, zero value otherwise.

### GetSha512Ok

`func (o *DebPackageResponse) GetSha512Ok() (*string, bool)`

GetSha512Ok returns a tuple with the Sha512 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha512

`func (o *DebPackageResponse) SetSha512(v string)`

SetSha512 sets Sha512 field to given value.

### HasSha512

`func (o *DebPackageResponse) HasSha512() bool`

HasSha512 returns a boolean if a field has been set.

### GetPackage

`func (o *DebPackageResponse) GetPackage() string`

GetPackage returns the Package field if non-nil, zero value otherwise.

### GetPackageOk

`func (o *DebPackageResponse) GetPackageOk() (*string, bool)`

GetPackageOk returns a tuple with the Package field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackage

`func (o *DebPackageResponse) SetPackage(v string)`

SetPackage sets Package field to given value.

### HasPackage

`func (o *DebPackageResponse) HasPackage() bool`

HasPackage returns a boolean if a field has been set.

### GetSource

`func (o *DebPackageResponse) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *DebPackageResponse) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *DebPackageResponse) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *DebPackageResponse) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetVersion

`func (o *DebPackageResponse) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *DebPackageResponse) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *DebPackageResponse) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *DebPackageResponse) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetArchitecture

`func (o *DebPackageResponse) GetArchitecture() string`

GetArchitecture returns the Architecture field if non-nil, zero value otherwise.

### GetArchitectureOk

`func (o *DebPackageResponse) GetArchitectureOk() (*string, bool)`

GetArchitectureOk returns a tuple with the Architecture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchitecture

`func (o *DebPackageResponse) SetArchitecture(v string)`

SetArchitecture sets Architecture field to given value.

### HasArchitecture

`func (o *DebPackageResponse) HasArchitecture() bool`

HasArchitecture returns a boolean if a field has been set.

### GetSection

`func (o *DebPackageResponse) GetSection() string`

GetSection returns the Section field if non-nil, zero value otherwise.

### GetSectionOk

`func (o *DebPackageResponse) GetSectionOk() (*string, bool)`

GetSectionOk returns a tuple with the Section field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSection

`func (o *DebPackageResponse) SetSection(v string)`

SetSection sets Section field to given value.

### HasSection

`func (o *DebPackageResponse) HasSection() bool`

HasSection returns a boolean if a field has been set.

### GetPriority

`func (o *DebPackageResponse) GetPriority() string`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *DebPackageResponse) GetPriorityOk() (*string, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *DebPackageResponse) SetPriority(v string)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *DebPackageResponse) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetOrigin

`func (o *DebPackageResponse) GetOrigin() string`

GetOrigin returns the Origin field if non-nil, zero value otherwise.

### GetOriginOk

`func (o *DebPackageResponse) GetOriginOk() (*string, bool)`

GetOriginOk returns a tuple with the Origin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigin

`func (o *DebPackageResponse) SetOrigin(v string)`

SetOrigin sets Origin field to given value.

### HasOrigin

`func (o *DebPackageResponse) HasOrigin() bool`

HasOrigin returns a boolean if a field has been set.

### GetTag

`func (o *DebPackageResponse) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *DebPackageResponse) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *DebPackageResponse) SetTag(v string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *DebPackageResponse) HasTag() bool`

HasTag returns a boolean if a field has been set.

### GetBugs

`func (o *DebPackageResponse) GetBugs() string`

GetBugs returns the Bugs field if non-nil, zero value otherwise.

### GetBugsOk

`func (o *DebPackageResponse) GetBugsOk() (*string, bool)`

GetBugsOk returns a tuple with the Bugs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBugs

`func (o *DebPackageResponse) SetBugs(v string)`

SetBugs sets Bugs field to given value.

### HasBugs

`func (o *DebPackageResponse) HasBugs() bool`

HasBugs returns a boolean if a field has been set.

### GetEssential

`func (o *DebPackageResponse) GetEssential() string`

GetEssential returns the Essential field if non-nil, zero value otherwise.

### GetEssentialOk

`func (o *DebPackageResponse) GetEssentialOk() (*string, bool)`

GetEssentialOk returns a tuple with the Essential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEssential

`func (o *DebPackageResponse) SetEssential(v string)`

SetEssential sets Essential field to given value.

### HasEssential

`func (o *DebPackageResponse) HasEssential() bool`

HasEssential returns a boolean if a field has been set.

### GetBuildEssential

`func (o *DebPackageResponse) GetBuildEssential() string`

GetBuildEssential returns the BuildEssential field if non-nil, zero value otherwise.

### GetBuildEssentialOk

`func (o *DebPackageResponse) GetBuildEssentialOk() (*string, bool)`

GetBuildEssentialOk returns a tuple with the BuildEssential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildEssential

`func (o *DebPackageResponse) SetBuildEssential(v string)`

SetBuildEssential sets BuildEssential field to given value.

### HasBuildEssential

`func (o *DebPackageResponse) HasBuildEssential() bool`

HasBuildEssential returns a boolean if a field has been set.

### GetInstalledSize

`func (o *DebPackageResponse) GetInstalledSize() string`

GetInstalledSize returns the InstalledSize field if non-nil, zero value otherwise.

### GetInstalledSizeOk

`func (o *DebPackageResponse) GetInstalledSizeOk() (*string, bool)`

GetInstalledSizeOk returns a tuple with the InstalledSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstalledSize

`func (o *DebPackageResponse) SetInstalledSize(v string)`

SetInstalledSize sets InstalledSize field to given value.

### HasInstalledSize

`func (o *DebPackageResponse) HasInstalledSize() bool`

HasInstalledSize returns a boolean if a field has been set.

### GetMaintainer

`func (o *DebPackageResponse) GetMaintainer() string`

GetMaintainer returns the Maintainer field if non-nil, zero value otherwise.

### GetMaintainerOk

`func (o *DebPackageResponse) GetMaintainerOk() (*string, bool)`

GetMaintainerOk returns a tuple with the Maintainer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintainer

`func (o *DebPackageResponse) SetMaintainer(v string)`

SetMaintainer sets Maintainer field to given value.

### HasMaintainer

`func (o *DebPackageResponse) HasMaintainer() bool`

HasMaintainer returns a boolean if a field has been set.

### GetOriginalMaintainer

`func (o *DebPackageResponse) GetOriginalMaintainer() string`

GetOriginalMaintainer returns the OriginalMaintainer field if non-nil, zero value otherwise.

### GetOriginalMaintainerOk

`func (o *DebPackageResponse) GetOriginalMaintainerOk() (*string, bool)`

GetOriginalMaintainerOk returns a tuple with the OriginalMaintainer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalMaintainer

`func (o *DebPackageResponse) SetOriginalMaintainer(v string)`

SetOriginalMaintainer sets OriginalMaintainer field to given value.

### HasOriginalMaintainer

`func (o *DebPackageResponse) HasOriginalMaintainer() bool`

HasOriginalMaintainer returns a boolean if a field has been set.

### GetDescription

`func (o *DebPackageResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DebPackageResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DebPackageResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DebPackageResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDescriptionMd5

`func (o *DebPackageResponse) GetDescriptionMd5() string`

GetDescriptionMd5 returns the DescriptionMd5 field if non-nil, zero value otherwise.

### GetDescriptionMd5Ok

`func (o *DebPackageResponse) GetDescriptionMd5Ok() (*string, bool)`

GetDescriptionMd5Ok returns a tuple with the DescriptionMd5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptionMd5

`func (o *DebPackageResponse) SetDescriptionMd5(v string)`

SetDescriptionMd5 sets DescriptionMd5 field to given value.

### HasDescriptionMd5

`func (o *DebPackageResponse) HasDescriptionMd5() bool`

HasDescriptionMd5 returns a boolean if a field has been set.

### GetHomepage

`func (o *DebPackageResponse) GetHomepage() string`

GetHomepage returns the Homepage field if non-nil, zero value otherwise.

### GetHomepageOk

`func (o *DebPackageResponse) GetHomepageOk() (*string, bool)`

GetHomepageOk returns a tuple with the Homepage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHomepage

`func (o *DebPackageResponse) SetHomepage(v string)`

SetHomepage sets Homepage field to given value.

### HasHomepage

`func (o *DebPackageResponse) HasHomepage() bool`

HasHomepage returns a boolean if a field has been set.

### GetBuiltUsing

`func (o *DebPackageResponse) GetBuiltUsing() string`

GetBuiltUsing returns the BuiltUsing field if non-nil, zero value otherwise.

### GetBuiltUsingOk

`func (o *DebPackageResponse) GetBuiltUsingOk() (*string, bool)`

GetBuiltUsingOk returns a tuple with the BuiltUsing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuiltUsing

`func (o *DebPackageResponse) SetBuiltUsing(v string)`

SetBuiltUsing sets BuiltUsing field to given value.

### HasBuiltUsing

`func (o *DebPackageResponse) HasBuiltUsing() bool`

HasBuiltUsing returns a boolean if a field has been set.

### GetAutoBuiltPackage

`func (o *DebPackageResponse) GetAutoBuiltPackage() string`

GetAutoBuiltPackage returns the AutoBuiltPackage field if non-nil, zero value otherwise.

### GetAutoBuiltPackageOk

`func (o *DebPackageResponse) GetAutoBuiltPackageOk() (*string, bool)`

GetAutoBuiltPackageOk returns a tuple with the AutoBuiltPackage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoBuiltPackage

`func (o *DebPackageResponse) SetAutoBuiltPackage(v string)`

SetAutoBuiltPackage sets AutoBuiltPackage field to given value.

### HasAutoBuiltPackage

`func (o *DebPackageResponse) HasAutoBuiltPackage() bool`

HasAutoBuiltPackage returns a boolean if a field has been set.

### GetMultiArch

`func (o *DebPackageResponse) GetMultiArch() string`

GetMultiArch returns the MultiArch field if non-nil, zero value otherwise.

### GetMultiArchOk

`func (o *DebPackageResponse) GetMultiArchOk() (*string, bool)`

GetMultiArchOk returns a tuple with the MultiArch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiArch

`func (o *DebPackageResponse) SetMultiArch(v string)`

SetMultiArch sets MultiArch field to given value.

### HasMultiArch

`func (o *DebPackageResponse) HasMultiArch() bool`

HasMultiArch returns a boolean if a field has been set.

### GetBreaks

`func (o *DebPackageResponse) GetBreaks() string`

GetBreaks returns the Breaks field if non-nil, zero value otherwise.

### GetBreaksOk

`func (o *DebPackageResponse) GetBreaksOk() (*string, bool)`

GetBreaksOk returns a tuple with the Breaks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBreaks

`func (o *DebPackageResponse) SetBreaks(v string)`

SetBreaks sets Breaks field to given value.

### HasBreaks

`func (o *DebPackageResponse) HasBreaks() bool`

HasBreaks returns a boolean if a field has been set.

### GetConflicts

`func (o *DebPackageResponse) GetConflicts() string`

GetConflicts returns the Conflicts field if non-nil, zero value otherwise.

### GetConflictsOk

`func (o *DebPackageResponse) GetConflictsOk() (*string, bool)`

GetConflictsOk returns a tuple with the Conflicts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConflicts

`func (o *DebPackageResponse) SetConflicts(v string)`

SetConflicts sets Conflicts field to given value.

### HasConflicts

`func (o *DebPackageResponse) HasConflicts() bool`

HasConflicts returns a boolean if a field has been set.

### GetDepends

`func (o *DebPackageResponse) GetDepends() string`

GetDepends returns the Depends field if non-nil, zero value otherwise.

### GetDependsOk

`func (o *DebPackageResponse) GetDependsOk() (*string, bool)`

GetDependsOk returns a tuple with the Depends field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepends

`func (o *DebPackageResponse) SetDepends(v string)`

SetDepends sets Depends field to given value.

### HasDepends

`func (o *DebPackageResponse) HasDepends() bool`

HasDepends returns a boolean if a field has been set.

### GetRecommends

`func (o *DebPackageResponse) GetRecommends() string`

GetRecommends returns the Recommends field if non-nil, zero value otherwise.

### GetRecommendsOk

`func (o *DebPackageResponse) GetRecommendsOk() (*string, bool)`

GetRecommendsOk returns a tuple with the Recommends field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecommends

`func (o *DebPackageResponse) SetRecommends(v string)`

SetRecommends sets Recommends field to given value.

### HasRecommends

`func (o *DebPackageResponse) HasRecommends() bool`

HasRecommends returns a boolean if a field has been set.

### GetSuggests

`func (o *DebPackageResponse) GetSuggests() string`

GetSuggests returns the Suggests field if non-nil, zero value otherwise.

### GetSuggestsOk

`func (o *DebPackageResponse) GetSuggestsOk() (*string, bool)`

GetSuggestsOk returns a tuple with the Suggests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuggests

`func (o *DebPackageResponse) SetSuggests(v string)`

SetSuggests sets Suggests field to given value.

### HasSuggests

`func (o *DebPackageResponse) HasSuggests() bool`

HasSuggests returns a boolean if a field has been set.

### GetEnhances

`func (o *DebPackageResponse) GetEnhances() string`

GetEnhances returns the Enhances field if non-nil, zero value otherwise.

### GetEnhancesOk

`func (o *DebPackageResponse) GetEnhancesOk() (*string, bool)`

GetEnhancesOk returns a tuple with the Enhances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnhances

`func (o *DebPackageResponse) SetEnhances(v string)`

SetEnhances sets Enhances field to given value.

### HasEnhances

`func (o *DebPackageResponse) HasEnhances() bool`

HasEnhances returns a boolean if a field has been set.

### GetPreDepends

`func (o *DebPackageResponse) GetPreDepends() string`

GetPreDepends returns the PreDepends field if non-nil, zero value otherwise.

### GetPreDependsOk

`func (o *DebPackageResponse) GetPreDependsOk() (*string, bool)`

GetPreDependsOk returns a tuple with the PreDepends field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreDepends

`func (o *DebPackageResponse) SetPreDepends(v string)`

SetPreDepends sets PreDepends field to given value.

### HasPreDepends

`func (o *DebPackageResponse) HasPreDepends() bool`

HasPreDepends returns a boolean if a field has been set.

### GetProvides

`func (o *DebPackageResponse) GetProvides() string`

GetProvides returns the Provides field if non-nil, zero value otherwise.

### GetProvidesOk

`func (o *DebPackageResponse) GetProvidesOk() (*string, bool)`

GetProvidesOk returns a tuple with the Provides field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvides

`func (o *DebPackageResponse) SetProvides(v string)`

SetProvides sets Provides field to given value.

### HasProvides

`func (o *DebPackageResponse) HasProvides() bool`

HasProvides returns a boolean if a field has been set.

### GetReplaces

`func (o *DebPackageResponse) GetReplaces() string`

GetReplaces returns the Replaces field if non-nil, zero value otherwise.

### GetReplacesOk

`func (o *DebPackageResponse) GetReplacesOk() (*string, bool)`

GetReplacesOk returns a tuple with the Replaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplaces

`func (o *DebPackageResponse) SetReplaces(v string)`

SetReplaces sets Replaces field to given value.

### HasReplaces

`func (o *DebPackageResponse) HasReplaces() bool`

HasReplaces returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


