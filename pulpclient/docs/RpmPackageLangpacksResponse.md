# RpmPackageLangpacksResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PulpHref** | Pointer to **string** |  | [optional] [readonly] 
**PulpCreated** | Pointer to **time.Time** | Timestamp of creation. | [optional] [readonly] 
**Matches** | **map[string]interface{}** | Langpacks matches. | 
**Digest** | **NullableString** | Langpacks digest. | 

## Methods

### NewRpmPackageLangpacksResponse

`func NewRpmPackageLangpacksResponse(matches map[string]interface{}, digest NullableString, ) *RpmPackageLangpacksResponse`

NewRpmPackageLangpacksResponse instantiates a new RpmPackageLangpacksResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRpmPackageLangpacksResponseWithDefaults

`func NewRpmPackageLangpacksResponseWithDefaults() *RpmPackageLangpacksResponse`

NewRpmPackageLangpacksResponseWithDefaults instantiates a new RpmPackageLangpacksResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPulpHref

`func (o *RpmPackageLangpacksResponse) GetPulpHref() string`

GetPulpHref returns the PulpHref field if non-nil, zero value otherwise.

### GetPulpHrefOk

`func (o *RpmPackageLangpacksResponse) GetPulpHrefOk() (*string, bool)`

GetPulpHrefOk returns a tuple with the PulpHref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpHref

`func (o *RpmPackageLangpacksResponse) SetPulpHref(v string)`

SetPulpHref sets PulpHref field to given value.

### HasPulpHref

`func (o *RpmPackageLangpacksResponse) HasPulpHref() bool`

HasPulpHref returns a boolean if a field has been set.

### GetPulpCreated

`func (o *RpmPackageLangpacksResponse) GetPulpCreated() time.Time`

GetPulpCreated returns the PulpCreated field if non-nil, zero value otherwise.

### GetPulpCreatedOk

`func (o *RpmPackageLangpacksResponse) GetPulpCreatedOk() (*time.Time, bool)`

GetPulpCreatedOk returns a tuple with the PulpCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpCreated

`func (o *RpmPackageLangpacksResponse) SetPulpCreated(v time.Time)`

SetPulpCreated sets PulpCreated field to given value.

### HasPulpCreated

`func (o *RpmPackageLangpacksResponse) HasPulpCreated() bool`

HasPulpCreated returns a boolean if a field has been set.

### GetMatches

`func (o *RpmPackageLangpacksResponse) GetMatches() map[string]interface{}`

GetMatches returns the Matches field if non-nil, zero value otherwise.

### GetMatchesOk

`func (o *RpmPackageLangpacksResponse) GetMatchesOk() (*map[string]interface{}, bool)`

GetMatchesOk returns a tuple with the Matches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatches

`func (o *RpmPackageLangpacksResponse) SetMatches(v map[string]interface{})`

SetMatches sets Matches field to given value.


### SetMatchesNil

`func (o *RpmPackageLangpacksResponse) SetMatchesNil(b bool)`

 SetMatchesNil sets the value for Matches to be an explicit nil

### UnsetMatches
`func (o *RpmPackageLangpacksResponse) UnsetMatches()`

UnsetMatches ensures that no value is present for Matches, not even an explicit nil
### GetDigest

`func (o *RpmPackageLangpacksResponse) GetDigest() string`

GetDigest returns the Digest field if non-nil, zero value otherwise.

### GetDigestOk

`func (o *RpmPackageLangpacksResponse) GetDigestOk() (*string, bool)`

GetDigestOk returns a tuple with the Digest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigest

`func (o *RpmPackageLangpacksResponse) SetDigest(v string)`

SetDigest sets Digest field to given value.


### SetDigestNil

`func (o *RpmPackageLangpacksResponse) SetDigestNil(b bool)`

 SetDigestNil sets the value for Digest to be an explicit nil

### UnsetDigest
`func (o *RpmPackageLangpacksResponse) UnsetDigest()`

UnsetDigest ensures that no value is present for Digest, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


