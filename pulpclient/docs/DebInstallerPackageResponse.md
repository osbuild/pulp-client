# DebInstallerPackageResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PulpHref** | Pointer to **string** |  | [optional] [readonly] 
**PulpCreated** | Pointer to **time.Time** | Timestamp of creation. | [optional] [readonly] 
**Artifact** | Pointer to **string** | Artifact file representing the physical content | [optional] 
**RelativePath** | Pointer to **string** | Path where the artifact is located relative to distributions base_path | [optional] 
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

### NewDebInstallerPackageResponse

`func NewDebInstallerPackageResponse() *DebInstallerPackageResponse`

NewDebInstallerPackageResponse instantiates a new DebInstallerPackageResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDebInstallerPackageResponseWithDefaults

`func NewDebInstallerPackageResponseWithDefaults() *DebInstallerPackageResponse`

NewDebInstallerPackageResponseWithDefaults instantiates a new DebInstallerPackageResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPulpHref

`func (o *DebInstallerPackageResponse) GetPulpHref() string`

GetPulpHref returns the PulpHref field if non-nil, zero value otherwise.

### GetPulpHrefOk

`func (o *DebInstallerPackageResponse) GetPulpHrefOk() (*string, bool)`

GetPulpHrefOk returns a tuple with the PulpHref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpHref

`func (o *DebInstallerPackageResponse) SetPulpHref(v string)`

SetPulpHref sets PulpHref field to given value.

### HasPulpHref

`func (o *DebInstallerPackageResponse) HasPulpHref() bool`

HasPulpHref returns a boolean if a field has been set.

### GetPulpCreated

`func (o *DebInstallerPackageResponse) GetPulpCreated() time.Time`

GetPulpCreated returns the PulpCreated field if non-nil, zero value otherwise.

### GetPulpCreatedOk

`func (o *DebInstallerPackageResponse) GetPulpCreatedOk() (*time.Time, bool)`

GetPulpCreatedOk returns a tuple with the PulpCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpCreated

`func (o *DebInstallerPackageResponse) SetPulpCreated(v time.Time)`

SetPulpCreated sets PulpCreated field to given value.

### HasPulpCreated

`func (o *DebInstallerPackageResponse) HasPulpCreated() bool`

HasPulpCreated returns a boolean if a field has been set.

### GetArtifact

`func (o *DebInstallerPackageResponse) GetArtifact() string`

GetArtifact returns the Artifact field if non-nil, zero value otherwise.

### GetArtifactOk

`func (o *DebInstallerPackageResponse) GetArtifactOk() (*string, bool)`

GetArtifactOk returns a tuple with the Artifact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifact

`func (o *DebInstallerPackageResponse) SetArtifact(v string)`

SetArtifact sets Artifact field to given value.

### HasArtifact

`func (o *DebInstallerPackageResponse) HasArtifact() bool`

HasArtifact returns a boolean if a field has been set.

### GetRelativePath

`func (o *DebInstallerPackageResponse) GetRelativePath() string`

GetRelativePath returns the RelativePath field if non-nil, zero value otherwise.

### GetRelativePathOk

`func (o *DebInstallerPackageResponse) GetRelativePathOk() (*string, bool)`

GetRelativePathOk returns a tuple with the RelativePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelativePath

`func (o *DebInstallerPackageResponse) SetRelativePath(v string)`

SetRelativePath sets RelativePath field to given value.

### HasRelativePath

`func (o *DebInstallerPackageResponse) HasRelativePath() bool`

HasRelativePath returns a boolean if a field has been set.

### GetMd5

`func (o *DebInstallerPackageResponse) GetMd5() string`

GetMd5 returns the Md5 field if non-nil, zero value otherwise.

### GetMd5Ok

`func (o *DebInstallerPackageResponse) GetMd5Ok() (*string, bool)`

GetMd5Ok returns a tuple with the Md5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMd5

`func (o *DebInstallerPackageResponse) SetMd5(v string)`

SetMd5 sets Md5 field to given value.

### HasMd5

`func (o *DebInstallerPackageResponse) HasMd5() bool`

HasMd5 returns a boolean if a field has been set.

### GetSha1

`func (o *DebInstallerPackageResponse) GetSha1() string`

GetSha1 returns the Sha1 field if non-nil, zero value otherwise.

### GetSha1Ok

`func (o *DebInstallerPackageResponse) GetSha1Ok() (*string, bool)`

GetSha1Ok returns a tuple with the Sha1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha1

`func (o *DebInstallerPackageResponse) SetSha1(v string)`

SetSha1 sets Sha1 field to given value.

### HasSha1

`func (o *DebInstallerPackageResponse) HasSha1() bool`

HasSha1 returns a boolean if a field has been set.

### GetSha224

`func (o *DebInstallerPackageResponse) GetSha224() string`

GetSha224 returns the Sha224 field if non-nil, zero value otherwise.

### GetSha224Ok

`func (o *DebInstallerPackageResponse) GetSha224Ok() (*string, bool)`

GetSha224Ok returns a tuple with the Sha224 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha224

`func (o *DebInstallerPackageResponse) SetSha224(v string)`

SetSha224 sets Sha224 field to given value.

### HasSha224

`func (o *DebInstallerPackageResponse) HasSha224() bool`

HasSha224 returns a boolean if a field has been set.

### GetSha256

`func (o *DebInstallerPackageResponse) GetSha256() string`

GetSha256 returns the Sha256 field if non-nil, zero value otherwise.

### GetSha256Ok

`func (o *DebInstallerPackageResponse) GetSha256Ok() (*string, bool)`

GetSha256Ok returns a tuple with the Sha256 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha256

`func (o *DebInstallerPackageResponse) SetSha256(v string)`

SetSha256 sets Sha256 field to given value.

### HasSha256

`func (o *DebInstallerPackageResponse) HasSha256() bool`

HasSha256 returns a boolean if a field has been set.

### GetSha384

`func (o *DebInstallerPackageResponse) GetSha384() string`

GetSha384 returns the Sha384 field if non-nil, zero value otherwise.

### GetSha384Ok

`func (o *DebInstallerPackageResponse) GetSha384Ok() (*string, bool)`

GetSha384Ok returns a tuple with the Sha384 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha384

`func (o *DebInstallerPackageResponse) SetSha384(v string)`

SetSha384 sets Sha384 field to given value.

### HasSha384

`func (o *DebInstallerPackageResponse) HasSha384() bool`

HasSha384 returns a boolean if a field has been set.

### GetSha512

`func (o *DebInstallerPackageResponse) GetSha512() string`

GetSha512 returns the Sha512 field if non-nil, zero value otherwise.

### GetSha512Ok

`func (o *DebInstallerPackageResponse) GetSha512Ok() (*string, bool)`

GetSha512Ok returns a tuple with the Sha512 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha512

`func (o *DebInstallerPackageResponse) SetSha512(v string)`

SetSha512 sets Sha512 field to given value.

### HasSha512

`func (o *DebInstallerPackageResponse) HasSha512() bool`

HasSha512 returns a boolean if a field has been set.

### GetPackage

`func (o *DebInstallerPackageResponse) GetPackage() string`

GetPackage returns the Package field if non-nil, zero value otherwise.

### GetPackageOk

`func (o *DebInstallerPackageResponse) GetPackageOk() (*string, bool)`

GetPackageOk returns a tuple with the Package field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackage

`func (o *DebInstallerPackageResponse) SetPackage(v string)`

SetPackage sets Package field to given value.

### HasPackage

`func (o *DebInstallerPackageResponse) HasPackage() bool`

HasPackage returns a boolean if a field has been set.

### GetSource

`func (o *DebInstallerPackageResponse) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *DebInstallerPackageResponse) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *DebInstallerPackageResponse) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *DebInstallerPackageResponse) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetVersion

`func (o *DebInstallerPackageResponse) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *DebInstallerPackageResponse) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *DebInstallerPackageResponse) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *DebInstallerPackageResponse) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetArchitecture

`func (o *DebInstallerPackageResponse) GetArchitecture() string`

GetArchitecture returns the Architecture field if non-nil, zero value otherwise.

### GetArchitectureOk

`func (o *DebInstallerPackageResponse) GetArchitectureOk() (*string, bool)`

GetArchitectureOk returns a tuple with the Architecture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchitecture

`func (o *DebInstallerPackageResponse) SetArchitecture(v string)`

SetArchitecture sets Architecture field to given value.

### HasArchitecture

`func (o *DebInstallerPackageResponse) HasArchitecture() bool`

HasArchitecture returns a boolean if a field has been set.

### GetSection

`func (o *DebInstallerPackageResponse) GetSection() string`

GetSection returns the Section field if non-nil, zero value otherwise.

### GetSectionOk

`func (o *DebInstallerPackageResponse) GetSectionOk() (*string, bool)`

GetSectionOk returns a tuple with the Section field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSection

`func (o *DebInstallerPackageResponse) SetSection(v string)`

SetSection sets Section field to given value.

### HasSection

`func (o *DebInstallerPackageResponse) HasSection() bool`

HasSection returns a boolean if a field has been set.

### GetPriority

`func (o *DebInstallerPackageResponse) GetPriority() string`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *DebInstallerPackageResponse) GetPriorityOk() (*string, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *DebInstallerPackageResponse) SetPriority(v string)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *DebInstallerPackageResponse) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetOrigin

`func (o *DebInstallerPackageResponse) GetOrigin() string`

GetOrigin returns the Origin field if non-nil, zero value otherwise.

### GetOriginOk

`func (o *DebInstallerPackageResponse) GetOriginOk() (*string, bool)`

GetOriginOk returns a tuple with the Origin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigin

`func (o *DebInstallerPackageResponse) SetOrigin(v string)`

SetOrigin sets Origin field to given value.

### HasOrigin

`func (o *DebInstallerPackageResponse) HasOrigin() bool`

HasOrigin returns a boolean if a field has been set.

### GetTag

`func (o *DebInstallerPackageResponse) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *DebInstallerPackageResponse) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *DebInstallerPackageResponse) SetTag(v string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *DebInstallerPackageResponse) HasTag() bool`

HasTag returns a boolean if a field has been set.

### GetBugs

`func (o *DebInstallerPackageResponse) GetBugs() string`

GetBugs returns the Bugs field if non-nil, zero value otherwise.

### GetBugsOk

`func (o *DebInstallerPackageResponse) GetBugsOk() (*string, bool)`

GetBugsOk returns a tuple with the Bugs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBugs

`func (o *DebInstallerPackageResponse) SetBugs(v string)`

SetBugs sets Bugs field to given value.

### HasBugs

`func (o *DebInstallerPackageResponse) HasBugs() bool`

HasBugs returns a boolean if a field has been set.

### GetEssential

`func (o *DebInstallerPackageResponse) GetEssential() string`

GetEssential returns the Essential field if non-nil, zero value otherwise.

### GetEssentialOk

`func (o *DebInstallerPackageResponse) GetEssentialOk() (*string, bool)`

GetEssentialOk returns a tuple with the Essential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEssential

`func (o *DebInstallerPackageResponse) SetEssential(v string)`

SetEssential sets Essential field to given value.

### HasEssential

`func (o *DebInstallerPackageResponse) HasEssential() bool`

HasEssential returns a boolean if a field has been set.

### GetBuildEssential

`func (o *DebInstallerPackageResponse) GetBuildEssential() string`

GetBuildEssential returns the BuildEssential field if non-nil, zero value otherwise.

### GetBuildEssentialOk

`func (o *DebInstallerPackageResponse) GetBuildEssentialOk() (*string, bool)`

GetBuildEssentialOk returns a tuple with the BuildEssential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildEssential

`func (o *DebInstallerPackageResponse) SetBuildEssential(v string)`

SetBuildEssential sets BuildEssential field to given value.

### HasBuildEssential

`func (o *DebInstallerPackageResponse) HasBuildEssential() bool`

HasBuildEssential returns a boolean if a field has been set.

### GetInstalledSize

`func (o *DebInstallerPackageResponse) GetInstalledSize() string`

GetInstalledSize returns the InstalledSize field if non-nil, zero value otherwise.

### GetInstalledSizeOk

`func (o *DebInstallerPackageResponse) GetInstalledSizeOk() (*string, bool)`

GetInstalledSizeOk returns a tuple with the InstalledSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstalledSize

`func (o *DebInstallerPackageResponse) SetInstalledSize(v string)`

SetInstalledSize sets InstalledSize field to given value.

### HasInstalledSize

`func (o *DebInstallerPackageResponse) HasInstalledSize() bool`

HasInstalledSize returns a boolean if a field has been set.

### GetMaintainer

`func (o *DebInstallerPackageResponse) GetMaintainer() string`

GetMaintainer returns the Maintainer field if non-nil, zero value otherwise.

### GetMaintainerOk

`func (o *DebInstallerPackageResponse) GetMaintainerOk() (*string, bool)`

GetMaintainerOk returns a tuple with the Maintainer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintainer

`func (o *DebInstallerPackageResponse) SetMaintainer(v string)`

SetMaintainer sets Maintainer field to given value.

### HasMaintainer

`func (o *DebInstallerPackageResponse) HasMaintainer() bool`

HasMaintainer returns a boolean if a field has been set.

### GetOriginalMaintainer

`func (o *DebInstallerPackageResponse) GetOriginalMaintainer() string`

GetOriginalMaintainer returns the OriginalMaintainer field if non-nil, zero value otherwise.

### GetOriginalMaintainerOk

`func (o *DebInstallerPackageResponse) GetOriginalMaintainerOk() (*string, bool)`

GetOriginalMaintainerOk returns a tuple with the OriginalMaintainer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalMaintainer

`func (o *DebInstallerPackageResponse) SetOriginalMaintainer(v string)`

SetOriginalMaintainer sets OriginalMaintainer field to given value.

### HasOriginalMaintainer

`func (o *DebInstallerPackageResponse) HasOriginalMaintainer() bool`

HasOriginalMaintainer returns a boolean if a field has been set.

### GetDescription

`func (o *DebInstallerPackageResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DebInstallerPackageResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DebInstallerPackageResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DebInstallerPackageResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDescriptionMd5

`func (o *DebInstallerPackageResponse) GetDescriptionMd5() string`

GetDescriptionMd5 returns the DescriptionMd5 field if non-nil, zero value otherwise.

### GetDescriptionMd5Ok

`func (o *DebInstallerPackageResponse) GetDescriptionMd5Ok() (*string, bool)`

GetDescriptionMd5Ok returns a tuple with the DescriptionMd5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptionMd5

`func (o *DebInstallerPackageResponse) SetDescriptionMd5(v string)`

SetDescriptionMd5 sets DescriptionMd5 field to given value.

### HasDescriptionMd5

`func (o *DebInstallerPackageResponse) HasDescriptionMd5() bool`

HasDescriptionMd5 returns a boolean if a field has been set.

### GetHomepage

`func (o *DebInstallerPackageResponse) GetHomepage() string`

GetHomepage returns the Homepage field if non-nil, zero value otherwise.

### GetHomepageOk

`func (o *DebInstallerPackageResponse) GetHomepageOk() (*string, bool)`

GetHomepageOk returns a tuple with the Homepage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHomepage

`func (o *DebInstallerPackageResponse) SetHomepage(v string)`

SetHomepage sets Homepage field to given value.

### HasHomepage

`func (o *DebInstallerPackageResponse) HasHomepage() bool`

HasHomepage returns a boolean if a field has been set.

### GetBuiltUsing

`func (o *DebInstallerPackageResponse) GetBuiltUsing() string`

GetBuiltUsing returns the BuiltUsing field if non-nil, zero value otherwise.

### GetBuiltUsingOk

`func (o *DebInstallerPackageResponse) GetBuiltUsingOk() (*string, bool)`

GetBuiltUsingOk returns a tuple with the BuiltUsing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuiltUsing

`func (o *DebInstallerPackageResponse) SetBuiltUsing(v string)`

SetBuiltUsing sets BuiltUsing field to given value.

### HasBuiltUsing

`func (o *DebInstallerPackageResponse) HasBuiltUsing() bool`

HasBuiltUsing returns a boolean if a field has been set.

### GetAutoBuiltPackage

`func (o *DebInstallerPackageResponse) GetAutoBuiltPackage() string`

GetAutoBuiltPackage returns the AutoBuiltPackage field if non-nil, zero value otherwise.

### GetAutoBuiltPackageOk

`func (o *DebInstallerPackageResponse) GetAutoBuiltPackageOk() (*string, bool)`

GetAutoBuiltPackageOk returns a tuple with the AutoBuiltPackage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoBuiltPackage

`func (o *DebInstallerPackageResponse) SetAutoBuiltPackage(v string)`

SetAutoBuiltPackage sets AutoBuiltPackage field to given value.

### HasAutoBuiltPackage

`func (o *DebInstallerPackageResponse) HasAutoBuiltPackage() bool`

HasAutoBuiltPackage returns a boolean if a field has been set.

### GetMultiArch

`func (o *DebInstallerPackageResponse) GetMultiArch() string`

GetMultiArch returns the MultiArch field if non-nil, zero value otherwise.

### GetMultiArchOk

`func (o *DebInstallerPackageResponse) GetMultiArchOk() (*string, bool)`

GetMultiArchOk returns a tuple with the MultiArch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiArch

`func (o *DebInstallerPackageResponse) SetMultiArch(v string)`

SetMultiArch sets MultiArch field to given value.

### HasMultiArch

`func (o *DebInstallerPackageResponse) HasMultiArch() bool`

HasMultiArch returns a boolean if a field has been set.

### GetBreaks

`func (o *DebInstallerPackageResponse) GetBreaks() string`

GetBreaks returns the Breaks field if non-nil, zero value otherwise.

### GetBreaksOk

`func (o *DebInstallerPackageResponse) GetBreaksOk() (*string, bool)`

GetBreaksOk returns a tuple with the Breaks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBreaks

`func (o *DebInstallerPackageResponse) SetBreaks(v string)`

SetBreaks sets Breaks field to given value.

### HasBreaks

`func (o *DebInstallerPackageResponse) HasBreaks() bool`

HasBreaks returns a boolean if a field has been set.

### GetConflicts

`func (o *DebInstallerPackageResponse) GetConflicts() string`

GetConflicts returns the Conflicts field if non-nil, zero value otherwise.

### GetConflictsOk

`func (o *DebInstallerPackageResponse) GetConflictsOk() (*string, bool)`

GetConflictsOk returns a tuple with the Conflicts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConflicts

`func (o *DebInstallerPackageResponse) SetConflicts(v string)`

SetConflicts sets Conflicts field to given value.

### HasConflicts

`func (o *DebInstallerPackageResponse) HasConflicts() bool`

HasConflicts returns a boolean if a field has been set.

### GetDepends

`func (o *DebInstallerPackageResponse) GetDepends() string`

GetDepends returns the Depends field if non-nil, zero value otherwise.

### GetDependsOk

`func (o *DebInstallerPackageResponse) GetDependsOk() (*string, bool)`

GetDependsOk returns a tuple with the Depends field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepends

`func (o *DebInstallerPackageResponse) SetDepends(v string)`

SetDepends sets Depends field to given value.

### HasDepends

`func (o *DebInstallerPackageResponse) HasDepends() bool`

HasDepends returns a boolean if a field has been set.

### GetRecommends

`func (o *DebInstallerPackageResponse) GetRecommends() string`

GetRecommends returns the Recommends field if non-nil, zero value otherwise.

### GetRecommendsOk

`func (o *DebInstallerPackageResponse) GetRecommendsOk() (*string, bool)`

GetRecommendsOk returns a tuple with the Recommends field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecommends

`func (o *DebInstallerPackageResponse) SetRecommends(v string)`

SetRecommends sets Recommends field to given value.

### HasRecommends

`func (o *DebInstallerPackageResponse) HasRecommends() bool`

HasRecommends returns a boolean if a field has been set.

### GetSuggests

`func (o *DebInstallerPackageResponse) GetSuggests() string`

GetSuggests returns the Suggests field if non-nil, zero value otherwise.

### GetSuggestsOk

`func (o *DebInstallerPackageResponse) GetSuggestsOk() (*string, bool)`

GetSuggestsOk returns a tuple with the Suggests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuggests

`func (o *DebInstallerPackageResponse) SetSuggests(v string)`

SetSuggests sets Suggests field to given value.

### HasSuggests

`func (o *DebInstallerPackageResponse) HasSuggests() bool`

HasSuggests returns a boolean if a field has been set.

### GetEnhances

`func (o *DebInstallerPackageResponse) GetEnhances() string`

GetEnhances returns the Enhances field if non-nil, zero value otherwise.

### GetEnhancesOk

`func (o *DebInstallerPackageResponse) GetEnhancesOk() (*string, bool)`

GetEnhancesOk returns a tuple with the Enhances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnhances

`func (o *DebInstallerPackageResponse) SetEnhances(v string)`

SetEnhances sets Enhances field to given value.

### HasEnhances

`func (o *DebInstallerPackageResponse) HasEnhances() bool`

HasEnhances returns a boolean if a field has been set.

### GetPreDepends

`func (o *DebInstallerPackageResponse) GetPreDepends() string`

GetPreDepends returns the PreDepends field if non-nil, zero value otherwise.

### GetPreDependsOk

`func (o *DebInstallerPackageResponse) GetPreDependsOk() (*string, bool)`

GetPreDependsOk returns a tuple with the PreDepends field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreDepends

`func (o *DebInstallerPackageResponse) SetPreDepends(v string)`

SetPreDepends sets PreDepends field to given value.

### HasPreDepends

`func (o *DebInstallerPackageResponse) HasPreDepends() bool`

HasPreDepends returns a boolean if a field has been set.

### GetProvides

`func (o *DebInstallerPackageResponse) GetProvides() string`

GetProvides returns the Provides field if non-nil, zero value otherwise.

### GetProvidesOk

`func (o *DebInstallerPackageResponse) GetProvidesOk() (*string, bool)`

GetProvidesOk returns a tuple with the Provides field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvides

`func (o *DebInstallerPackageResponse) SetProvides(v string)`

SetProvides sets Provides field to given value.

### HasProvides

`func (o *DebInstallerPackageResponse) HasProvides() bool`

HasProvides returns a boolean if a field has been set.

### GetReplaces

`func (o *DebInstallerPackageResponse) GetReplaces() string`

GetReplaces returns the Replaces field if non-nil, zero value otherwise.

### GetReplacesOk

`func (o *DebInstallerPackageResponse) GetReplacesOk() (*string, bool)`

GetReplacesOk returns a tuple with the Replaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplaces

`func (o *DebInstallerPackageResponse) SetReplaces(v string)`

SetReplaces sets Replaces field to given value.

### HasReplaces

`func (o *DebInstallerPackageResponse) HasReplaces() bool`

HasReplaces returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


