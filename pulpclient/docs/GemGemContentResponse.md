# GemGemContentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PulpHref** | Pointer to **string** |  | [optional] [readonly] 
**PulpCreated** | Pointer to **time.Time** | Timestamp of creation. | [optional] [readonly] 
**Artifacts** | **map[string]interface{}** | A dict mapping relative paths inside the Content to the correspondingArtifact URLs. E.g.: {&#39;relative/path&#39;: &#39;/artifacts/1/&#39; | [readonly] 
**Checksum** | Pointer to **string** | SHA256 checksum of the gem | [optional] [readonly] 
**Name** | Pointer to **string** | Name of the gem | [optional] [readonly] 
**Version** | Pointer to **string** | Version of the gem | [optional] [readonly] 
**Prerelease** | Pointer to **bool** | Whether the gem is a prerelease | [optional] [readonly] 
**Dependencies** | Pointer to **map[string]string** |  | [optional] [readonly] 
**RequiredRubyVersion** | Pointer to **string** | Required ruby version of the gem | [optional] [readonly] 
**RequiredRubygemsVersion** | Pointer to **string** | Required rubygems version of the gem | [optional] [readonly] 

## Methods

### NewGemGemContentResponse

`func NewGemGemContentResponse(artifacts map[string]interface{}, ) *GemGemContentResponse`

NewGemGemContentResponse instantiates a new GemGemContentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGemGemContentResponseWithDefaults

`func NewGemGemContentResponseWithDefaults() *GemGemContentResponse`

NewGemGemContentResponseWithDefaults instantiates a new GemGemContentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPulpHref

`func (o *GemGemContentResponse) GetPulpHref() string`

GetPulpHref returns the PulpHref field if non-nil, zero value otherwise.

### GetPulpHrefOk

`func (o *GemGemContentResponse) GetPulpHrefOk() (*string, bool)`

GetPulpHrefOk returns a tuple with the PulpHref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpHref

`func (o *GemGemContentResponse) SetPulpHref(v string)`

SetPulpHref sets PulpHref field to given value.

### HasPulpHref

`func (o *GemGemContentResponse) HasPulpHref() bool`

HasPulpHref returns a boolean if a field has been set.

### GetPulpCreated

`func (o *GemGemContentResponse) GetPulpCreated() time.Time`

GetPulpCreated returns the PulpCreated field if non-nil, zero value otherwise.

### GetPulpCreatedOk

`func (o *GemGemContentResponse) GetPulpCreatedOk() (*time.Time, bool)`

GetPulpCreatedOk returns a tuple with the PulpCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpCreated

`func (o *GemGemContentResponse) SetPulpCreated(v time.Time)`

SetPulpCreated sets PulpCreated field to given value.

### HasPulpCreated

`func (o *GemGemContentResponse) HasPulpCreated() bool`

HasPulpCreated returns a boolean if a field has been set.

### GetArtifacts

`func (o *GemGemContentResponse) GetArtifacts() map[string]interface{}`

GetArtifacts returns the Artifacts field if non-nil, zero value otherwise.

### GetArtifactsOk

`func (o *GemGemContentResponse) GetArtifactsOk() (*map[string]interface{}, bool)`

GetArtifactsOk returns a tuple with the Artifacts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifacts

`func (o *GemGemContentResponse) SetArtifacts(v map[string]interface{})`

SetArtifacts sets Artifacts field to given value.


### GetChecksum

`func (o *GemGemContentResponse) GetChecksum() string`

GetChecksum returns the Checksum field if non-nil, zero value otherwise.

### GetChecksumOk

`func (o *GemGemContentResponse) GetChecksumOk() (*string, bool)`

GetChecksumOk returns a tuple with the Checksum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChecksum

`func (o *GemGemContentResponse) SetChecksum(v string)`

SetChecksum sets Checksum field to given value.

### HasChecksum

`func (o *GemGemContentResponse) HasChecksum() bool`

HasChecksum returns a boolean if a field has been set.

### GetName

`func (o *GemGemContentResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GemGemContentResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GemGemContentResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GemGemContentResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetVersion

`func (o *GemGemContentResponse) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *GemGemContentResponse) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *GemGemContentResponse) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *GemGemContentResponse) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetPrerelease

`func (o *GemGemContentResponse) GetPrerelease() bool`

GetPrerelease returns the Prerelease field if non-nil, zero value otherwise.

### GetPrereleaseOk

`func (o *GemGemContentResponse) GetPrereleaseOk() (*bool, bool)`

GetPrereleaseOk returns a tuple with the Prerelease field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrerelease

`func (o *GemGemContentResponse) SetPrerelease(v bool)`

SetPrerelease sets Prerelease field to given value.

### HasPrerelease

`func (o *GemGemContentResponse) HasPrerelease() bool`

HasPrerelease returns a boolean if a field has been set.

### GetDependencies

`func (o *GemGemContentResponse) GetDependencies() map[string]string`

GetDependencies returns the Dependencies field if non-nil, zero value otherwise.

### GetDependenciesOk

`func (o *GemGemContentResponse) GetDependenciesOk() (*map[string]string, bool)`

GetDependenciesOk returns a tuple with the Dependencies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependencies

`func (o *GemGemContentResponse) SetDependencies(v map[string]string)`

SetDependencies sets Dependencies field to given value.

### HasDependencies

`func (o *GemGemContentResponse) HasDependencies() bool`

HasDependencies returns a boolean if a field has been set.

### GetRequiredRubyVersion

`func (o *GemGemContentResponse) GetRequiredRubyVersion() string`

GetRequiredRubyVersion returns the RequiredRubyVersion field if non-nil, zero value otherwise.

### GetRequiredRubyVersionOk

`func (o *GemGemContentResponse) GetRequiredRubyVersionOk() (*string, bool)`

GetRequiredRubyVersionOk returns a tuple with the RequiredRubyVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredRubyVersion

`func (o *GemGemContentResponse) SetRequiredRubyVersion(v string)`

SetRequiredRubyVersion sets RequiredRubyVersion field to given value.

### HasRequiredRubyVersion

`func (o *GemGemContentResponse) HasRequiredRubyVersion() bool`

HasRequiredRubyVersion returns a boolean if a field has been set.

### GetRequiredRubygemsVersion

`func (o *GemGemContentResponse) GetRequiredRubygemsVersion() string`

GetRequiredRubygemsVersion returns the RequiredRubygemsVersion field if non-nil, zero value otherwise.

### GetRequiredRubygemsVersionOk

`func (o *GemGemContentResponse) GetRequiredRubygemsVersionOk() (*string, bool)`

GetRequiredRubygemsVersionOk returns a tuple with the RequiredRubygemsVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredRubygemsVersion

`func (o *GemGemContentResponse) SetRequiredRubygemsVersion(v string)`

SetRequiredRubygemsVersion sets RequiredRubygemsVersion field to given value.

### HasRequiredRubygemsVersion

`func (o *GemGemContentResponse) HasRequiredRubygemsVersion() bool`

HasRequiredRubygemsVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


