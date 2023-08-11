# DebReleaseComponentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PulpHref** | Pointer to **string** |  | [optional] [readonly] 
**PulpCreated** | Pointer to **time.Time** | Timestamp of creation. | [optional] [readonly] 
**Component** | **string** | Name of the component. | 
**Distribution** | **string** | Name of the distribution. | 
**Codename** | **string** |  | 
**Suite** | **string** |  | 

## Methods

### NewDebReleaseComponentResponse

`func NewDebReleaseComponentResponse(component string, distribution string, codename string, suite string, ) *DebReleaseComponentResponse`

NewDebReleaseComponentResponse instantiates a new DebReleaseComponentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDebReleaseComponentResponseWithDefaults

`func NewDebReleaseComponentResponseWithDefaults() *DebReleaseComponentResponse`

NewDebReleaseComponentResponseWithDefaults instantiates a new DebReleaseComponentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPulpHref

`func (o *DebReleaseComponentResponse) GetPulpHref() string`

GetPulpHref returns the PulpHref field if non-nil, zero value otherwise.

### GetPulpHrefOk

`func (o *DebReleaseComponentResponse) GetPulpHrefOk() (*string, bool)`

GetPulpHrefOk returns a tuple with the PulpHref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpHref

`func (o *DebReleaseComponentResponse) SetPulpHref(v string)`

SetPulpHref sets PulpHref field to given value.

### HasPulpHref

`func (o *DebReleaseComponentResponse) HasPulpHref() bool`

HasPulpHref returns a boolean if a field has been set.

### GetPulpCreated

`func (o *DebReleaseComponentResponse) GetPulpCreated() time.Time`

GetPulpCreated returns the PulpCreated field if non-nil, zero value otherwise.

### GetPulpCreatedOk

`func (o *DebReleaseComponentResponse) GetPulpCreatedOk() (*time.Time, bool)`

GetPulpCreatedOk returns a tuple with the PulpCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpCreated

`func (o *DebReleaseComponentResponse) SetPulpCreated(v time.Time)`

SetPulpCreated sets PulpCreated field to given value.

### HasPulpCreated

`func (o *DebReleaseComponentResponse) HasPulpCreated() bool`

HasPulpCreated returns a boolean if a field has been set.

### GetComponent

`func (o *DebReleaseComponentResponse) GetComponent() string`

GetComponent returns the Component field if non-nil, zero value otherwise.

### GetComponentOk

`func (o *DebReleaseComponentResponse) GetComponentOk() (*string, bool)`

GetComponentOk returns a tuple with the Component field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponent

`func (o *DebReleaseComponentResponse) SetComponent(v string)`

SetComponent sets Component field to given value.


### GetDistribution

`func (o *DebReleaseComponentResponse) GetDistribution() string`

GetDistribution returns the Distribution field if non-nil, zero value otherwise.

### GetDistributionOk

`func (o *DebReleaseComponentResponse) GetDistributionOk() (*string, bool)`

GetDistributionOk returns a tuple with the Distribution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistribution

`func (o *DebReleaseComponentResponse) SetDistribution(v string)`

SetDistribution sets Distribution field to given value.


### GetCodename

`func (o *DebReleaseComponentResponse) GetCodename() string`

GetCodename returns the Codename field if non-nil, zero value otherwise.

### GetCodenameOk

`func (o *DebReleaseComponentResponse) GetCodenameOk() (*string, bool)`

GetCodenameOk returns a tuple with the Codename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodename

`func (o *DebReleaseComponentResponse) SetCodename(v string)`

SetCodename sets Codename field to given value.


### GetSuite

`func (o *DebReleaseComponentResponse) GetSuite() string`

GetSuite returns the Suite field if non-nil, zero value otherwise.

### GetSuiteOk

`func (o *DebReleaseComponentResponse) GetSuiteOk() (*string, bool)`

GetSuiteOk returns a tuple with the Suite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuite

`func (o *DebReleaseComponentResponse) SetSuite(v string)`

SetSuite sets Suite field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


