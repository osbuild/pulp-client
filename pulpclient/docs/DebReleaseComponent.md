# DebReleaseComponent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Repository** | Pointer to **string** | A URI of a repository the new content unit should be associated with. | [optional] 
**Component** | **string** | Name of the component. | 
**Distribution** | **string** | Name of the distribution. | 
**Codename** | **string** |  | 
**Suite** | **string** |  | 

## Methods

### NewDebReleaseComponent

`func NewDebReleaseComponent(component string, distribution string, codename string, suite string, ) *DebReleaseComponent`

NewDebReleaseComponent instantiates a new DebReleaseComponent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDebReleaseComponentWithDefaults

`func NewDebReleaseComponentWithDefaults() *DebReleaseComponent`

NewDebReleaseComponentWithDefaults instantiates a new DebReleaseComponent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRepository

`func (o *DebReleaseComponent) GetRepository() string`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *DebReleaseComponent) GetRepositoryOk() (*string, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *DebReleaseComponent) SetRepository(v string)`

SetRepository sets Repository field to given value.

### HasRepository

`func (o *DebReleaseComponent) HasRepository() bool`

HasRepository returns a boolean if a field has been set.

### GetComponent

`func (o *DebReleaseComponent) GetComponent() string`

GetComponent returns the Component field if non-nil, zero value otherwise.

### GetComponentOk

`func (o *DebReleaseComponent) GetComponentOk() (*string, bool)`

GetComponentOk returns a tuple with the Component field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponent

`func (o *DebReleaseComponent) SetComponent(v string)`

SetComponent sets Component field to given value.


### GetDistribution

`func (o *DebReleaseComponent) GetDistribution() string`

GetDistribution returns the Distribution field if non-nil, zero value otherwise.

### GetDistributionOk

`func (o *DebReleaseComponent) GetDistributionOk() (*string, bool)`

GetDistributionOk returns a tuple with the Distribution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistribution

`func (o *DebReleaseComponent) SetDistribution(v string)`

SetDistribution sets Distribution field to given value.


### GetCodename

`func (o *DebReleaseComponent) GetCodename() string`

GetCodename returns the Codename field if non-nil, zero value otherwise.

### GetCodenameOk

`func (o *DebReleaseComponent) GetCodenameOk() (*string, bool)`

GetCodenameOk returns a tuple with the Codename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodename

`func (o *DebReleaseComponent) SetCodename(v string)`

SetCodename sets Codename field to given value.


### GetSuite

`func (o *DebReleaseComponent) GetSuite() string`

GetSuite returns the Suite field if non-nil, zero value otherwise.

### GetSuiteOk

`func (o *DebReleaseComponent) GetSuiteOk() (*string, bool)`

GetSuiteOk returns a tuple with the Suite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuite

`func (o *DebReleaseComponent) SetSuite(v string)`

SetSuite sets Suite field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


