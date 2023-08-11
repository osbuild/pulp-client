# OstreeOstreeConfigResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PulpHref** | Pointer to **string** |  | [optional] [readonly] 
**PulpCreated** | Pointer to **time.Time** | Timestamp of creation. | [optional] [readonly] 
**Artifact** | **string** | Artifact file representing the physical content | 
**RelativePath** | **string** | Path where the artifact is located relative to distributions base_path | 

## Methods

### NewOstreeOstreeConfigResponse

`func NewOstreeOstreeConfigResponse(artifact string, relativePath string, ) *OstreeOstreeConfigResponse`

NewOstreeOstreeConfigResponse instantiates a new OstreeOstreeConfigResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOstreeOstreeConfigResponseWithDefaults

`func NewOstreeOstreeConfigResponseWithDefaults() *OstreeOstreeConfigResponse`

NewOstreeOstreeConfigResponseWithDefaults instantiates a new OstreeOstreeConfigResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPulpHref

`func (o *OstreeOstreeConfigResponse) GetPulpHref() string`

GetPulpHref returns the PulpHref field if non-nil, zero value otherwise.

### GetPulpHrefOk

`func (o *OstreeOstreeConfigResponse) GetPulpHrefOk() (*string, bool)`

GetPulpHrefOk returns a tuple with the PulpHref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpHref

`func (o *OstreeOstreeConfigResponse) SetPulpHref(v string)`

SetPulpHref sets PulpHref field to given value.

### HasPulpHref

`func (o *OstreeOstreeConfigResponse) HasPulpHref() bool`

HasPulpHref returns a boolean if a field has been set.

### GetPulpCreated

`func (o *OstreeOstreeConfigResponse) GetPulpCreated() time.Time`

GetPulpCreated returns the PulpCreated field if non-nil, zero value otherwise.

### GetPulpCreatedOk

`func (o *OstreeOstreeConfigResponse) GetPulpCreatedOk() (*time.Time, bool)`

GetPulpCreatedOk returns a tuple with the PulpCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpCreated

`func (o *OstreeOstreeConfigResponse) SetPulpCreated(v time.Time)`

SetPulpCreated sets PulpCreated field to given value.

### HasPulpCreated

`func (o *OstreeOstreeConfigResponse) HasPulpCreated() bool`

HasPulpCreated returns a boolean if a field has been set.

### GetArtifact

`func (o *OstreeOstreeConfigResponse) GetArtifact() string`

GetArtifact returns the Artifact field if non-nil, zero value otherwise.

### GetArtifactOk

`func (o *OstreeOstreeConfigResponse) GetArtifactOk() (*string, bool)`

GetArtifactOk returns a tuple with the Artifact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifact

`func (o *OstreeOstreeConfigResponse) SetArtifact(v string)`

SetArtifact sets Artifact field to given value.


### GetRelativePath

`func (o *OstreeOstreeConfigResponse) GetRelativePath() string`

GetRelativePath returns the RelativePath field if non-nil, zero value otherwise.

### GetRelativePathOk

`func (o *OstreeOstreeConfigResponse) GetRelativePathOk() (*string, bool)`

GetRelativePathOk returns a tuple with the RelativePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelativePath

`func (o *OstreeOstreeConfigResponse) SetRelativePath(v string)`

SetRelativePath sets RelativePath field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


