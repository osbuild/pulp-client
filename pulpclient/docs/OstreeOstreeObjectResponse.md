# OstreeOstreeObjectResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PulpHref** | Pointer to **string** |  | [optional] [readonly] 
**PulpCreated** | Pointer to **time.Time** | Timestamp of creation. | [optional] [readonly] 
**Artifact** | **string** | Artifact file representing the physical content | 
**RelativePath** | **string** | Path where the artifact is located relative to distributions base_path | 
**Checksum** | **string** |  | 
**Typ** | **int64** |              The type of an object. All values are described by the mapping declared at             https://lazka.github.io/pgi-docs/OSTree-1.0/enums.html#OSTree.ObjectType              | 

## Methods

### NewOstreeOstreeObjectResponse

`func NewOstreeOstreeObjectResponse(artifact string, relativePath string, checksum string, typ int64, ) *OstreeOstreeObjectResponse`

NewOstreeOstreeObjectResponse instantiates a new OstreeOstreeObjectResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOstreeOstreeObjectResponseWithDefaults

`func NewOstreeOstreeObjectResponseWithDefaults() *OstreeOstreeObjectResponse`

NewOstreeOstreeObjectResponseWithDefaults instantiates a new OstreeOstreeObjectResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPulpHref

`func (o *OstreeOstreeObjectResponse) GetPulpHref() string`

GetPulpHref returns the PulpHref field if non-nil, zero value otherwise.

### GetPulpHrefOk

`func (o *OstreeOstreeObjectResponse) GetPulpHrefOk() (*string, bool)`

GetPulpHrefOk returns a tuple with the PulpHref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpHref

`func (o *OstreeOstreeObjectResponse) SetPulpHref(v string)`

SetPulpHref sets PulpHref field to given value.

### HasPulpHref

`func (o *OstreeOstreeObjectResponse) HasPulpHref() bool`

HasPulpHref returns a boolean if a field has been set.

### GetPulpCreated

`func (o *OstreeOstreeObjectResponse) GetPulpCreated() time.Time`

GetPulpCreated returns the PulpCreated field if non-nil, zero value otherwise.

### GetPulpCreatedOk

`func (o *OstreeOstreeObjectResponse) GetPulpCreatedOk() (*time.Time, bool)`

GetPulpCreatedOk returns a tuple with the PulpCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpCreated

`func (o *OstreeOstreeObjectResponse) SetPulpCreated(v time.Time)`

SetPulpCreated sets PulpCreated field to given value.

### HasPulpCreated

`func (o *OstreeOstreeObjectResponse) HasPulpCreated() bool`

HasPulpCreated returns a boolean if a field has been set.

### GetArtifact

`func (o *OstreeOstreeObjectResponse) GetArtifact() string`

GetArtifact returns the Artifact field if non-nil, zero value otherwise.

### GetArtifactOk

`func (o *OstreeOstreeObjectResponse) GetArtifactOk() (*string, bool)`

GetArtifactOk returns a tuple with the Artifact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifact

`func (o *OstreeOstreeObjectResponse) SetArtifact(v string)`

SetArtifact sets Artifact field to given value.


### GetRelativePath

`func (o *OstreeOstreeObjectResponse) GetRelativePath() string`

GetRelativePath returns the RelativePath field if non-nil, zero value otherwise.

### GetRelativePathOk

`func (o *OstreeOstreeObjectResponse) GetRelativePathOk() (*string, bool)`

GetRelativePathOk returns a tuple with the RelativePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelativePath

`func (o *OstreeOstreeObjectResponse) SetRelativePath(v string)`

SetRelativePath sets RelativePath field to given value.


### GetChecksum

`func (o *OstreeOstreeObjectResponse) GetChecksum() string`

GetChecksum returns the Checksum field if non-nil, zero value otherwise.

### GetChecksumOk

`func (o *OstreeOstreeObjectResponse) GetChecksumOk() (*string, bool)`

GetChecksumOk returns a tuple with the Checksum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChecksum

`func (o *OstreeOstreeObjectResponse) SetChecksum(v string)`

SetChecksum sets Checksum field to given value.


### GetTyp

`func (o *OstreeOstreeObjectResponse) GetTyp() int64`

GetTyp returns the Typ field if non-nil, zero value otherwise.

### GetTypOk

`func (o *OstreeOstreeObjectResponse) GetTypOk() (*int64, bool)`

GetTypOk returns a tuple with the Typ field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTyp

`func (o *OstreeOstreeObjectResponse) SetTyp(v int64)`

SetTyp sets Typ field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


