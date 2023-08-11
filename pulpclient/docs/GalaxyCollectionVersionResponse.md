# GalaxyCollectionVersionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Version** | **string** |  | 
**Href** | Pointer to **string** | Get href. | [optional] [readonly] 
**Namespace** | Pointer to **map[string]interface{}** | Create a namespace dict. | [optional] [readonly] 
**Collection** | Pointer to **map[string]interface{}** | Create a collection dict. | [optional] [readonly] 
**Artifact** | Pointer to **map[string]interface{}** | Create an artifact dict. | [optional] [readonly] 
**Metadata** | [**CollectionMetadataResponse**](CollectionMetadataResponse.md) |  | 

## Methods

### NewGalaxyCollectionVersionResponse

`func NewGalaxyCollectionVersionResponse(version string, metadata CollectionMetadataResponse, ) *GalaxyCollectionVersionResponse`

NewGalaxyCollectionVersionResponse instantiates a new GalaxyCollectionVersionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGalaxyCollectionVersionResponseWithDefaults

`func NewGalaxyCollectionVersionResponseWithDefaults() *GalaxyCollectionVersionResponse`

NewGalaxyCollectionVersionResponseWithDefaults instantiates a new GalaxyCollectionVersionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersion

`func (o *GalaxyCollectionVersionResponse) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *GalaxyCollectionVersionResponse) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *GalaxyCollectionVersionResponse) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetHref

`func (o *GalaxyCollectionVersionResponse) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *GalaxyCollectionVersionResponse) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *GalaxyCollectionVersionResponse) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *GalaxyCollectionVersionResponse) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetNamespace

`func (o *GalaxyCollectionVersionResponse) GetNamespace() map[string]interface{}`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *GalaxyCollectionVersionResponse) GetNamespaceOk() (*map[string]interface{}, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *GalaxyCollectionVersionResponse) SetNamespace(v map[string]interface{})`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *GalaxyCollectionVersionResponse) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetCollection

`func (o *GalaxyCollectionVersionResponse) GetCollection() map[string]interface{}`

GetCollection returns the Collection field if non-nil, zero value otherwise.

### GetCollectionOk

`func (o *GalaxyCollectionVersionResponse) GetCollectionOk() (*map[string]interface{}, bool)`

GetCollectionOk returns a tuple with the Collection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollection

`func (o *GalaxyCollectionVersionResponse) SetCollection(v map[string]interface{})`

SetCollection sets Collection field to given value.

### HasCollection

`func (o *GalaxyCollectionVersionResponse) HasCollection() bool`

HasCollection returns a boolean if a field has been set.

### GetArtifact

`func (o *GalaxyCollectionVersionResponse) GetArtifact() map[string]interface{}`

GetArtifact returns the Artifact field if non-nil, zero value otherwise.

### GetArtifactOk

`func (o *GalaxyCollectionVersionResponse) GetArtifactOk() (*map[string]interface{}, bool)`

GetArtifactOk returns a tuple with the Artifact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifact

`func (o *GalaxyCollectionVersionResponse) SetArtifact(v map[string]interface{})`

SetArtifact sets Artifact field to given value.

### HasArtifact

`func (o *GalaxyCollectionVersionResponse) HasArtifact() bool`

HasArtifact returns a boolean if a field has been set.

### GetMetadata

`func (o *GalaxyCollectionVersionResponse) GetMetadata() CollectionMetadataResponse`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GalaxyCollectionVersionResponse) GetMetadataOk() (*CollectionMetadataResponse, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GalaxyCollectionVersionResponse) SetMetadata(v CollectionMetadataResponse)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


