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

// InlineResponse400CauseSuppressed struct for InlineResponse400CauseSuppressed
type InlineResponse400CauseSuppressed struct {

	StackTrace *[]InlineResponse400CauseStackTrace `json:"stackTrace,omitempty"`

	Message *string `json:"message,omitempty"`

	LocalizedMessage *string `json:"localizedMessage,omitempty"`

}

// NewInlineResponse400CauseSuppressed instantiates a new InlineResponse400CauseSuppressed object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse400CauseSuppressed() *InlineResponse400CauseSuppressed {
	this := InlineResponse400CauseSuppressed{}
	return &this
}

// NewInlineResponse400CauseSuppressedWithDefaults instantiates a new InlineResponse400CauseSuppressed object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse400CauseSuppressedWithDefaults() *InlineResponse400CauseSuppressed {
	this := InlineResponse400CauseSuppressed{}




	return &this
}


// GetStackTrace returns the StackTrace field value if set, zero value otherwise.
func (o *InlineResponse400CauseSuppressed) GetStackTrace() []InlineResponse400CauseStackTrace {
	if o == nil || o.StackTrace == nil {
		var ret []InlineResponse400CauseStackTrace
		return ret
	}
	return *o.StackTrace
}

// GetStackTraceOk returns a tuple with the StackTrace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse400CauseSuppressed) GetStackTraceOk() (*[]InlineResponse400CauseStackTrace, bool) {
	if o == nil || o.StackTrace == nil {
		return nil, false
	}
	return o.StackTrace, true
}

// HasStackTrace returns a boolean if a field has been set.
func (o *InlineResponse400CauseSuppressed) HasStackTrace() bool {
	if o != nil && o.StackTrace != nil {
		return true
	}

	return false
}

// SetStackTrace gets a reference to the given []InlineResponse400CauseStackTrace and assigns it to the StackTrace field.
func (o *InlineResponse400CauseSuppressed) SetStackTrace(v []InlineResponse400CauseStackTrace) {
	o.StackTrace = &v
}


// GetMessage returns the Message field value if set, zero value otherwise.
func (o *InlineResponse400CauseSuppressed) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse400CauseSuppressed) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *InlineResponse400CauseSuppressed) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *InlineResponse400CauseSuppressed) SetMessage(v string) {
	o.Message = &v
}


// GetLocalizedMessage returns the LocalizedMessage field value if set, zero value otherwise.
func (o *InlineResponse400CauseSuppressed) GetLocalizedMessage() string {
	if o == nil || o.LocalizedMessage == nil {
		var ret string
		return ret
	}
	return *o.LocalizedMessage
}

// GetLocalizedMessageOk returns a tuple with the LocalizedMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse400CauseSuppressed) GetLocalizedMessageOk() (*string, bool) {
	if o == nil || o.LocalizedMessage == nil {
		return nil, false
	}
	return o.LocalizedMessage, true
}

// HasLocalizedMessage returns a boolean if a field has been set.
func (o *InlineResponse400CauseSuppressed) HasLocalizedMessage() bool {
	if o != nil && o.LocalizedMessage != nil {
		return true
	}

	return false
}

// SetLocalizedMessage gets a reference to the given string and assigns it to the LocalizedMessage field.
func (o *InlineResponse400CauseSuppressed) SetLocalizedMessage(v string) {
	o.LocalizedMessage = &v
}


func (o InlineResponse400CauseSuppressed) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	
	if o.StackTrace != nil {
		toSerialize["stackTrace"] = o.StackTrace
	}
    
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
    
	if o.LocalizedMessage != nil {
		toSerialize["localizedMessage"] = o.LocalizedMessage
	}
    
	return json.Marshal(toSerialize)
}

type NullableInlineResponse400CauseSuppressed struct {
	value *InlineResponse400CauseSuppressed
	isSet bool
}

func (v NullableInlineResponse400CauseSuppressed) Get() *InlineResponse400CauseSuppressed {
	return v.value
}

func (v *NullableInlineResponse400CauseSuppressed) Set(val *InlineResponse400CauseSuppressed) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse400CauseSuppressed) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse400CauseSuppressed) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse400CauseSuppressed(val *InlineResponse400CauseSuppressed) *NullableInlineResponse400CauseSuppressed {
	return &NullableInlineResponse400CauseSuppressed{value: val, isSet: true}
}

func (v NullableInlineResponse400CauseSuppressed) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse400CauseSuppressed) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

