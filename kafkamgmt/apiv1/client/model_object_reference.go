/*
 * Kafka Management API
 *
 * Kafka Management API is a REST API to manage Kafka instances
 *
 * API version: 1.12.0
 * Contact: rhosak-support@redhat.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package kafkamgmtclient

import (
	"encoding/json"
)

// ObjectReference struct for ObjectReference
type ObjectReference struct {
	Id string `json:"id"`
	Kind string `json:"kind"`
	Href string `json:"href"`
}

// NewObjectReference instantiates a new ObjectReference object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewObjectReference(id string, kind string, href string) *ObjectReference {
	this := ObjectReference{}
	this.Id = id
	this.Kind = kind
	this.Href = href
	return &this
}

// NewObjectReferenceWithDefaults instantiates a new ObjectReference object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewObjectReferenceWithDefaults() *ObjectReference {
	this := ObjectReference{}
	return &this
}

// GetId returns the Id field value
func (o *ObjectReference) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ObjectReference) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ObjectReference) SetId(v string) {
	o.Id = v
}

// GetKind returns the Kind field value
func (o *ObjectReference) GetKind() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Kind
}

// GetKindOk returns a tuple with the Kind field value
// and a boolean to check if the value has been set.
func (o *ObjectReference) GetKindOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Kind, true
}

// SetKind sets field value
func (o *ObjectReference) SetKind(v string) {
	o.Kind = v
}

// GetHref returns the Href field value
func (o *ObjectReference) GetHref() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Href
}

// GetHrefOk returns a tuple with the Href field value
// and a boolean to check if the value has been set.
func (o *ObjectReference) GetHrefOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Href, true
}

// SetHref sets field value
func (o *ObjectReference) SetHref(v string) {
	o.Href = v
}

func (o ObjectReference) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["kind"] = o.Kind
	}
	if true {
		toSerialize["href"] = o.Href
	}
	return json.Marshal(toSerialize)
}

type NullableObjectReference struct {
	value *ObjectReference
	isSet bool
}

func (v NullableObjectReference) Get() *ObjectReference {
	return v.value
}

func (v *NullableObjectReference) Set(val *ObjectReference) {
	v.value = val
	v.isSet = true
}

func (v NullableObjectReference) IsSet() bool {
	return v.isSet
}

func (v *NullableObjectReference) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableObjectReference(val *ObjectReference) *NullableObjectReference {
	return &NullableObjectReference{value: val, isSet: true}
}

func (v NullableObjectReference) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableObjectReference) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


