# ConnectorNamespacePatchRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Annotations** | Pointer to [**[]ConnectorNamespaceRequestMetaAnnotations**](ConnectorNamespaceRequestMetaAnnotations.md) |  | [optional] 


## Methods

### NewConnectorNamespacePatchRequest

`func NewConnectorNamespacePatchRequest() *ConnectorNamespacePatchRequest`

NewConnectorNamespacePatchRequest instantiates a new ConnectorNamespacePatchRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectorNamespacePatchRequestWithDefaults

`func NewConnectorNamespacePatchRequestWithDefaults() *ConnectorNamespacePatchRequest`

NewConnectorNamespacePatchRequestWithDefaults instantiates a new ConnectorNamespacePatchRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


### GetName

`func (o *ConnectorNamespacePatchRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ConnectorNamespacePatchRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ConnectorNamespacePatchRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ConnectorNamespacePatchRequest) HasName() bool`

HasName returns a boolean if a field has been set.


### GetAnnotations

`func (o *ConnectorNamespacePatchRequest) GetAnnotations() []ConnectorNamespaceRequestMetaAnnotations`

GetAnnotations returns the Annotations field if non-nil, zero value otherwise.

### GetAnnotationsOk

`func (o *ConnectorNamespacePatchRequest) GetAnnotationsOk() (*[]ConnectorNamespaceRequestMetaAnnotations, bool)`

GetAnnotationsOk returns a tuple with the Annotations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotations

`func (o *ConnectorNamespacePatchRequest) SetAnnotations(v []ConnectorNamespaceRequestMetaAnnotations)`

SetAnnotations sets Annotations field to given value.

### HasAnnotations

`func (o *ConnectorNamespacePatchRequest) HasAnnotations() bool`

HasAnnotations returns a boolean if a field has been set.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
