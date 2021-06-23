/*
 * Connector Service Fleet Manager
 *
 * Connector Service Fleet Manager is a Rest API to manage connectors.
 *
 * API version: 0.0.3
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package connectormgmtclient

import (
	"encoding/json"
)

// ConnectorTypeListAllOf struct for ConnectorTypeListAllOf
type ConnectorTypeListAllOf struct {

	Items *[]ConnectorType `json:"items,omitempty"`

}

// NewConnectorTypeListAllOf instantiates a new ConnectorTypeListAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConnectorTypeListAllOf() *ConnectorTypeListAllOf {
	this := ConnectorTypeListAllOf{}
	return &this
}

// NewConnectorTypeListAllOfWithDefaults instantiates a new ConnectorTypeListAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConnectorTypeListAllOfWithDefaults() *ConnectorTypeListAllOf {
	this := ConnectorTypeListAllOf{}


	return &this
}


// GetItems returns the Items field value if set, zero value otherwise.
func (o *ConnectorTypeListAllOf) GetItems() []ConnectorType {
	if o == nil || o.Items == nil {
		var ret []ConnectorType
		return ret
	}
	return *o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorTypeListAllOf) GetItemsOk() (*[]ConnectorType, bool) {
	if o == nil || o.Items == nil {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *ConnectorTypeListAllOf) HasItems() bool {
	if o != nil && o.Items != nil {
		return true
	}

	return false
}

// SetItems gets a reference to the given []ConnectorType and assigns it to the Items field.
func (o *ConnectorTypeListAllOf) SetItems(v []ConnectorType) {
	o.Items = &v
}


func (o ConnectorTypeListAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	
	if o.Items != nil {
		toSerialize["items"] = o.Items
	}
    
	return json.Marshal(toSerialize)
}

type NullableConnectorTypeListAllOf struct {
	value *ConnectorTypeListAllOf
	isSet bool
}

func (v NullableConnectorTypeListAllOf) Get() *ConnectorTypeListAllOf {
	return v.value
}

func (v *NullableConnectorTypeListAllOf) Set(val *ConnectorTypeListAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableConnectorTypeListAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableConnectorTypeListAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConnectorTypeListAllOf(val *ConnectorTypeListAllOf) *NullableConnectorTypeListAllOf {
	return &NullableConnectorTypeListAllOf{value: val, isSet: true}
}

func (v NullableConnectorTypeListAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConnectorTypeListAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

