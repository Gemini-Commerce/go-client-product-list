/*
Collection Service

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1
Contact: info@gemini-commerce.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package productlist

import (
	"encoding/json"
)

// checks if the ProductlistProductListError type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProductlistProductListError{}

// ProductlistProductListError struct for ProductlistProductListError
type ProductlistProductListError struct {
	Code *string `json:"code,omitempty"`
	Message *string `json:"message,omitempty"`
}

// NewProductlistProductListError instantiates a new ProductlistProductListError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductlistProductListError() *ProductlistProductListError {
	this := ProductlistProductListError{}
	return &this
}

// NewProductlistProductListErrorWithDefaults instantiates a new ProductlistProductListError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductlistProductListErrorWithDefaults() *ProductlistProductListError {
	this := ProductlistProductListError{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *ProductlistProductListError) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductlistProductListError) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *ProductlistProductListError) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *ProductlistProductListError) SetCode(v string) {
	o.Code = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *ProductlistProductListError) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductlistProductListError) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *ProductlistProductListError) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *ProductlistProductListError) SetMessage(v string) {
	o.Message = &v
}

func (o ProductlistProductListError) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProductlistProductListError) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	return toSerialize, nil
}

type NullableProductlistProductListError struct {
	value *ProductlistProductListError
	isSet bool
}

func (v NullableProductlistProductListError) Get() *ProductlistProductListError {
	return v.value
}

func (v *NullableProductlistProductListError) Set(val *ProductlistProductListError) {
	v.value = val
	v.isSet = true
}

func (v NullableProductlistProductListError) IsSet() bool {
	return v.isSet
}

func (v *NullableProductlistProductListError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductlistProductListError(val *ProductlistProductListError) *NullableProductlistProductListError {
	return &NullableProductlistProductListError{value: val, isSet: true}
}

func (v NullableProductlistProductListError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProductlistProductListError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


