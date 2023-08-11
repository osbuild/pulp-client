# AnsibleRepositorySignature

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContentUnits** | **[]interface{}** | List of collection version hrefs to sign, use * to sign all content in repository | 
**SigningService** | **string** | A signing service to use to sign the collections | 

## Methods

### NewAnsibleRepositorySignature

`func NewAnsibleRepositorySignature(contentUnits []interface{}, signingService string, ) *AnsibleRepositorySignature`

NewAnsibleRepositorySignature instantiates a new AnsibleRepositorySignature object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnsibleRepositorySignatureWithDefaults

`func NewAnsibleRepositorySignatureWithDefaults() *AnsibleRepositorySignature`

NewAnsibleRepositorySignatureWithDefaults instantiates a new AnsibleRepositorySignature object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContentUnits

`func (o *AnsibleRepositorySignature) GetContentUnits() []interface{}`

GetContentUnits returns the ContentUnits field if non-nil, zero value otherwise.

### GetContentUnitsOk

`func (o *AnsibleRepositorySignature) GetContentUnitsOk() (*[]interface{}, bool)`

GetContentUnitsOk returns a tuple with the ContentUnits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentUnits

`func (o *AnsibleRepositorySignature) SetContentUnits(v []interface{})`

SetContentUnits sets ContentUnits field to given value.


### GetSigningService

`func (o *AnsibleRepositorySignature) GetSigningService() string`

GetSigningService returns the SigningService field if non-nil, zero value otherwise.

### GetSigningServiceOk

`func (o *AnsibleRepositorySignature) GetSigningServiceOk() (*string, bool)`

GetSigningServiceOk returns a tuple with the SigningService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigningService

`func (o *AnsibleRepositorySignature) SetSigningService(v string)`

SetSigningService sets SigningService field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


