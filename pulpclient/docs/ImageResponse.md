# ImageResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | File name. | 
**Path** | **string** | File path. | 
**Platforms** | **string** | Compatible platforms. | 
**Artifact** | [**NullableArtifactResponse**](ArtifactResponse.md) |  | 

## Methods

### NewImageResponse

`func NewImageResponse(name string, path string, platforms string, artifact NullableArtifactResponse, ) *ImageResponse`

NewImageResponse instantiates a new ImageResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImageResponseWithDefaults

`func NewImageResponseWithDefaults() *ImageResponse`

NewImageResponseWithDefaults instantiates a new ImageResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ImageResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ImageResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ImageResponse) SetName(v string)`

SetName sets Name field to given value.


### GetPath

`func (o *ImageResponse) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *ImageResponse) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *ImageResponse) SetPath(v string)`

SetPath sets Path field to given value.


### GetPlatforms

`func (o *ImageResponse) GetPlatforms() string`

GetPlatforms returns the Platforms field if non-nil, zero value otherwise.

### GetPlatformsOk

`func (o *ImageResponse) GetPlatformsOk() (*string, bool)`

GetPlatformsOk returns a tuple with the Platforms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatforms

`func (o *ImageResponse) SetPlatforms(v string)`

SetPlatforms sets Platforms field to given value.


### GetArtifact

`func (o *ImageResponse) GetArtifact() ArtifactResponse`

GetArtifact returns the Artifact field if non-nil, zero value otherwise.

### GetArtifactOk

`func (o *ImageResponse) GetArtifactOk() (*ArtifactResponse, bool)`

GetArtifactOk returns a tuple with the Artifact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifact

`func (o *ImageResponse) SetArtifact(v ArtifactResponse)`

SetArtifact sets Artifact field to given value.


### SetArtifactNil

`func (o *ImageResponse) SetArtifactNil(b bool)`

 SetArtifactNil sets the value for Artifact to be an explicit nil

### UnsetArtifact
`func (o *ImageResponse) UnsetArtifact()`

UnsetArtifact ensures that no value is present for Artifact, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


