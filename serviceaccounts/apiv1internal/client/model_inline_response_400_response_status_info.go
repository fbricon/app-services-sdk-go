/*
 * Service Accounts API Documentation
 *
 * This is the API documentation for Service Accounts
 *
 * API version: 1.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package serviceaccountsclient

import (
	"encoding/json"
)

// InlineResponse400ResponseStatusInfo struct for InlineResponse400ResponseStatusInfo
type InlineResponse400ResponseStatusInfo struct {

	ReasonPhrase *string `json:"reasonPhrase,omitempty"`

	StatusCode *int32 `json:"statusCode,omitempty"`

	Family *string `json:"family,omitempty"`

}

// NewInlineResponse400ResponseStatusInfo instantiates a new InlineResponse400ResponseStatusInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse400ResponseStatusInfo() *InlineResponse400ResponseStatusInfo {
	this := InlineResponse400ResponseStatusInfo{}
	return &this
}

// NewInlineResponse400ResponseStatusInfoWithDefaults instantiates a new InlineResponse400ResponseStatusInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse400ResponseStatusInfoWithDefaults() *InlineResponse400ResponseStatusInfo {
	this := InlineResponse400ResponseStatusInfo{}




	return &this
}


// GetReasonPhrase returns the ReasonPhrase field value if set, zero value otherwise.
func (o *InlineResponse400ResponseStatusInfo) GetReasonPhrase() string {
	if o == nil || o.ReasonPhrase == nil {
		var ret string
		return ret
	}
	return *o.ReasonPhrase
}

// GetReasonPhraseOk returns a tuple with the ReasonPhrase field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse400ResponseStatusInfo) GetReasonPhraseOk() (*string, bool) {
	if o == nil || o.ReasonPhrase == nil {
		return nil, false
	}
	return o.ReasonPhrase, true
}

// HasReasonPhrase returns a boolean if a field has been set.
func (o *InlineResponse400ResponseStatusInfo) HasReasonPhrase() bool {
	if o != nil && o.ReasonPhrase != nil {
		return true
	}

	return false
}

// SetReasonPhrase gets a reference to the given string and assigns it to the ReasonPhrase field.
func (o *InlineResponse400ResponseStatusInfo) SetReasonPhrase(v string) {
	o.ReasonPhrase = &v
}


// GetStatusCode returns the StatusCode field value if set, zero value otherwise.
func (o *InlineResponse400ResponseStatusInfo) GetStatusCode() int32 {
	if o == nil || o.StatusCode == nil {
		var ret int32
		return ret
	}
	return *o.StatusCode
}

// GetStatusCodeOk returns a tuple with the StatusCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse400ResponseStatusInfo) GetStatusCodeOk() (*int32, bool) {
	if o == nil || o.StatusCode == nil {
		return nil, false
	}
	return o.StatusCode, true
}

// HasStatusCode returns a boolean if a field has been set.
func (o *InlineResponse400ResponseStatusInfo) HasStatusCode() bool {
	if o != nil && o.StatusCode != nil {
		return true
	}

	return false
}

// SetStatusCode gets a reference to the given int32 and assigns it to the StatusCode field.
func (o *InlineResponse400ResponseStatusInfo) SetStatusCode(v int32) {
	o.StatusCode = &v
}


// GetFamily returns the Family field value if set, zero value otherwise.
func (o *InlineResponse400ResponseStatusInfo) GetFamily() string {
	if o == nil || o.Family == nil {
		var ret string
		return ret
	}
	return *o.Family
}

// GetFamilyOk returns a tuple with the Family field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse400ResponseStatusInfo) GetFamilyOk() (*string, bool) {
	if o == nil || o.Family == nil {
		return nil, false
	}
	return o.Family, true
}

// HasFamily returns a boolean if a field has been set.
func (o *InlineResponse400ResponseStatusInfo) HasFamily() bool {
	if o != nil && o.Family != nil {
		return true
	}

	return false
}

// SetFamily gets a reference to the given string and assigns it to the Family field.
func (o *InlineResponse400ResponseStatusInfo) SetFamily(v string) {
	o.Family = &v
}


func (o InlineResponse400ResponseStatusInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	
	if o.ReasonPhrase != nil {
		toSerialize["reasonPhrase"] = o.ReasonPhrase
	}
    
	if o.StatusCode != nil {
		toSerialize["statusCode"] = o.StatusCode
	}
    
	if o.Family != nil {
		toSerialize["family"] = o.Family
	}
    
	return json.Marshal(toSerialize)
}

type NullableInlineResponse400ResponseStatusInfo struct {
	value *InlineResponse400ResponseStatusInfo
	isSet bool
}

func (v NullableInlineResponse400ResponseStatusInfo) Get() *InlineResponse400ResponseStatusInfo {
	return v.value
}

func (v *NullableInlineResponse400ResponseStatusInfo) Set(val *InlineResponse400ResponseStatusInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse400ResponseStatusInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse400ResponseStatusInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse400ResponseStatusInfo(val *InlineResponse400ResponseStatusInfo) *NullableInlineResponse400ResponseStatusInfo {
	return &NullableInlineResponse400ResponseStatusInfo{value: val, isSet: true}
}

func (v NullableInlineResponse400ResponseStatusInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse400ResponseStatusInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

