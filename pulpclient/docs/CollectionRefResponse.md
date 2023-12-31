# CollectionRefResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** |  | 
**Href** | Pointer to **string** | Returns link to a collection. | [optional] [readonly] 

## Methods

### NewCollectionRefResponse

`func NewCollectionRefResponse(id string, name string, ) *CollectionRefResponse`

NewCollectionRefResponse instantiates a new CollectionRefResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCollectionRefResponseWithDefaults

`func NewCollectionRefResponseWithDefaults() *CollectionRefResponse`

NewCollectionRefResponseWithDefaults instantiates a new CollectionRefResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CollectionRefResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CollectionRefResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CollectionRefResponse) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *CollectionRefResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CollectionRefResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CollectionRefResponse) SetName(v string)`

SetName sets Name field to given value.


### GetHref

`func (o *CollectionRefResponse) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *CollectionRefResponse) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *CollectionRefResponse) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *CollectionRefResponse) HasHref() bool`

HasHref returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


