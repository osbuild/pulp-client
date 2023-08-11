# OstreeOstreeDistributionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PulpHref** | Pointer to **string** |  | [optional] [readonly] 
**PulpCreated** | Pointer to **time.Time** | Timestamp of creation. | [optional] [readonly] 
**BasePath** | **string** | The base (relative) path component of the published url. Avoid paths that                     overlap with other distribution base paths (e.g. \&quot;foo\&quot; and \&quot;foo/bar\&quot;) | 
**BaseUrl** | Pointer to **string** | The URL for accessing the publication as defined by this distribution. | [optional] [readonly] 
**ContentGuard** | Pointer to **NullableString** | An optional content-guard. | [optional] 
**Hidden** | Pointer to **bool** | Whether this distribution should be shown in the content app. | [optional] [default to false]
**PulpLabels** | Pointer to **map[string]string** |  | [optional] 
**Name** | **string** | A unique name. Ex, &#x60;rawhide&#x60; and &#x60;stable&#x60;. | 
**Repository** | Pointer to **NullableString** | The latest RepositoryVersion for this Repository will be served. | [optional] 
**RepositoryVersion** | Pointer to **NullableString** | RepositoryVersion to be served | [optional] 

## Methods

### NewOstreeOstreeDistributionResponse

`func NewOstreeOstreeDistributionResponse(basePath string, name string, ) *OstreeOstreeDistributionResponse`

NewOstreeOstreeDistributionResponse instantiates a new OstreeOstreeDistributionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOstreeOstreeDistributionResponseWithDefaults

`func NewOstreeOstreeDistributionResponseWithDefaults() *OstreeOstreeDistributionResponse`

NewOstreeOstreeDistributionResponseWithDefaults instantiates a new OstreeOstreeDistributionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPulpHref

`func (o *OstreeOstreeDistributionResponse) GetPulpHref() string`

GetPulpHref returns the PulpHref field if non-nil, zero value otherwise.

### GetPulpHrefOk

`func (o *OstreeOstreeDistributionResponse) GetPulpHrefOk() (*string, bool)`

GetPulpHrefOk returns a tuple with the PulpHref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpHref

`func (o *OstreeOstreeDistributionResponse) SetPulpHref(v string)`

SetPulpHref sets PulpHref field to given value.

### HasPulpHref

`func (o *OstreeOstreeDistributionResponse) HasPulpHref() bool`

HasPulpHref returns a boolean if a field has been set.

### GetPulpCreated

`func (o *OstreeOstreeDistributionResponse) GetPulpCreated() time.Time`

GetPulpCreated returns the PulpCreated field if non-nil, zero value otherwise.

### GetPulpCreatedOk

`func (o *OstreeOstreeDistributionResponse) GetPulpCreatedOk() (*time.Time, bool)`

GetPulpCreatedOk returns a tuple with the PulpCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpCreated

`func (o *OstreeOstreeDistributionResponse) SetPulpCreated(v time.Time)`

SetPulpCreated sets PulpCreated field to given value.

### HasPulpCreated

`func (o *OstreeOstreeDistributionResponse) HasPulpCreated() bool`

HasPulpCreated returns a boolean if a field has been set.

### GetBasePath

`func (o *OstreeOstreeDistributionResponse) GetBasePath() string`

GetBasePath returns the BasePath field if non-nil, zero value otherwise.

### GetBasePathOk

`func (o *OstreeOstreeDistributionResponse) GetBasePathOk() (*string, bool)`

GetBasePathOk returns a tuple with the BasePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasePath

`func (o *OstreeOstreeDistributionResponse) SetBasePath(v string)`

SetBasePath sets BasePath field to given value.


### GetBaseUrl

`func (o *OstreeOstreeDistributionResponse) GetBaseUrl() string`

GetBaseUrl returns the BaseUrl field if non-nil, zero value otherwise.

### GetBaseUrlOk

`func (o *OstreeOstreeDistributionResponse) GetBaseUrlOk() (*string, bool)`

GetBaseUrlOk returns a tuple with the BaseUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseUrl

`func (o *OstreeOstreeDistributionResponse) SetBaseUrl(v string)`

SetBaseUrl sets BaseUrl field to given value.

### HasBaseUrl

`func (o *OstreeOstreeDistributionResponse) HasBaseUrl() bool`

HasBaseUrl returns a boolean if a field has been set.

### GetContentGuard

`func (o *OstreeOstreeDistributionResponse) GetContentGuard() string`

GetContentGuard returns the ContentGuard field if non-nil, zero value otherwise.

### GetContentGuardOk

`func (o *OstreeOstreeDistributionResponse) GetContentGuardOk() (*string, bool)`

GetContentGuardOk returns a tuple with the ContentGuard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentGuard

`func (o *OstreeOstreeDistributionResponse) SetContentGuard(v string)`

SetContentGuard sets ContentGuard field to given value.

### HasContentGuard

`func (o *OstreeOstreeDistributionResponse) HasContentGuard() bool`

HasContentGuard returns a boolean if a field has been set.

### SetContentGuardNil

`func (o *OstreeOstreeDistributionResponse) SetContentGuardNil(b bool)`

 SetContentGuardNil sets the value for ContentGuard to be an explicit nil

### UnsetContentGuard
`func (o *OstreeOstreeDistributionResponse) UnsetContentGuard()`

UnsetContentGuard ensures that no value is present for ContentGuard, not even an explicit nil
### GetHidden

`func (o *OstreeOstreeDistributionResponse) GetHidden() bool`

GetHidden returns the Hidden field if non-nil, zero value otherwise.

### GetHiddenOk

`func (o *OstreeOstreeDistributionResponse) GetHiddenOk() (*bool, bool)`

GetHiddenOk returns a tuple with the Hidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHidden

`func (o *OstreeOstreeDistributionResponse) SetHidden(v bool)`

SetHidden sets Hidden field to given value.

### HasHidden

`func (o *OstreeOstreeDistributionResponse) HasHidden() bool`

HasHidden returns a boolean if a field has been set.

### GetPulpLabels

`func (o *OstreeOstreeDistributionResponse) GetPulpLabels() map[string]string`

GetPulpLabels returns the PulpLabels field if non-nil, zero value otherwise.

### GetPulpLabelsOk

`func (o *OstreeOstreeDistributionResponse) GetPulpLabelsOk() (*map[string]string, bool)`

GetPulpLabelsOk returns a tuple with the PulpLabels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpLabels

`func (o *OstreeOstreeDistributionResponse) SetPulpLabels(v map[string]string)`

SetPulpLabels sets PulpLabels field to given value.

### HasPulpLabels

`func (o *OstreeOstreeDistributionResponse) HasPulpLabels() bool`

HasPulpLabels returns a boolean if a field has been set.

### GetName

`func (o *OstreeOstreeDistributionResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OstreeOstreeDistributionResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OstreeOstreeDistributionResponse) SetName(v string)`

SetName sets Name field to given value.


### GetRepository

`func (o *OstreeOstreeDistributionResponse) GetRepository() string`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *OstreeOstreeDistributionResponse) GetRepositoryOk() (*string, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *OstreeOstreeDistributionResponse) SetRepository(v string)`

SetRepository sets Repository field to given value.

### HasRepository

`func (o *OstreeOstreeDistributionResponse) HasRepository() bool`

HasRepository returns a boolean if a field has been set.

### SetRepositoryNil

`func (o *OstreeOstreeDistributionResponse) SetRepositoryNil(b bool)`

 SetRepositoryNil sets the value for Repository to be an explicit nil

### UnsetRepository
`func (o *OstreeOstreeDistributionResponse) UnsetRepository()`

UnsetRepository ensures that no value is present for Repository, not even an explicit nil
### GetRepositoryVersion

`func (o *OstreeOstreeDistributionResponse) GetRepositoryVersion() string`

GetRepositoryVersion returns the RepositoryVersion field if non-nil, zero value otherwise.

### GetRepositoryVersionOk

`func (o *OstreeOstreeDistributionResponse) GetRepositoryVersionOk() (*string, bool)`

GetRepositoryVersionOk returns a tuple with the RepositoryVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepositoryVersion

`func (o *OstreeOstreeDistributionResponse) SetRepositoryVersion(v string)`

SetRepositoryVersion sets RepositoryVersion field to given value.

### HasRepositoryVersion

`func (o *OstreeOstreeDistributionResponse) HasRepositoryVersion() bool`

HasRepositoryVersion returns a boolean if a field has been set.

### SetRepositoryVersionNil

`func (o *OstreeOstreeDistributionResponse) SetRepositoryVersionNil(b bool)`

 SetRepositoryVersionNil sets the value for RepositoryVersion to be an explicit nil

### UnsetRepositoryVersion
`func (o *OstreeOstreeDistributionResponse) UnsetRepositoryVersion()`

UnsetRepositoryVersion ensures that no value is present for RepositoryVersion, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


