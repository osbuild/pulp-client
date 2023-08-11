# CollectionVersionCopyMove

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CollectionVersions** | **[]string** | A list of collection versions to move or copy. | 
**DestinationRepositories** | **[]string** | List of repository HREFs to put content in. | 
**SigningService** | Pointer to **string** | HREF for a signing service. This will be used to sign the collection before moving putting it in any new repositories. | [optional] 

## Methods

### NewCollectionVersionCopyMove

`func NewCollectionVersionCopyMove(collectionVersions []string, destinationRepositories []string, ) *CollectionVersionCopyMove`

NewCollectionVersionCopyMove instantiates a new CollectionVersionCopyMove object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCollectionVersionCopyMoveWithDefaults

`func NewCollectionVersionCopyMoveWithDefaults() *CollectionVersionCopyMove`

NewCollectionVersionCopyMoveWithDefaults instantiates a new CollectionVersionCopyMove object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCollectionVersions

`func (o *CollectionVersionCopyMove) GetCollectionVersions() []string`

GetCollectionVersions returns the CollectionVersions field if non-nil, zero value otherwise.

### GetCollectionVersionsOk

`func (o *CollectionVersionCopyMove) GetCollectionVersionsOk() (*[]string, bool)`

GetCollectionVersionsOk returns a tuple with the CollectionVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionVersions

`func (o *CollectionVersionCopyMove) SetCollectionVersions(v []string)`

SetCollectionVersions sets CollectionVersions field to given value.


### GetDestinationRepositories

`func (o *CollectionVersionCopyMove) GetDestinationRepositories() []string`

GetDestinationRepositories returns the DestinationRepositories field if non-nil, zero value otherwise.

### GetDestinationRepositoriesOk

`func (o *CollectionVersionCopyMove) GetDestinationRepositoriesOk() (*[]string, bool)`

GetDestinationRepositoriesOk returns a tuple with the DestinationRepositories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationRepositories

`func (o *CollectionVersionCopyMove) SetDestinationRepositories(v []string)`

SetDestinationRepositories sets DestinationRepositories field to given value.


### GetSigningService

`func (o *CollectionVersionCopyMove) GetSigningService() string`

GetSigningService returns the SigningService field if non-nil, zero value otherwise.

### GetSigningServiceOk

`func (o *CollectionVersionCopyMove) GetSigningServiceOk() (*string, bool)`

GetSigningServiceOk returns a tuple with the SigningService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigningService

`func (o *CollectionVersionCopyMove) SetSigningService(v string)`

SetSigningService sets SigningService field to given value.

### HasSigningService

`func (o *CollectionVersionCopyMove) HasSigningService() bool`

HasSigningService returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


