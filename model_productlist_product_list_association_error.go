/*
Collection Service

API for managing collection

API version: v1
Contact: info@gemini-commerce.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package product-list

import (
	"encoding/json"
)

// checks if the ProductlistProductListAssociationError type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProductlistProductListAssociationError{}

// ProductlistProductListAssociationError struct for ProductlistProductListAssociationError
type ProductlistProductListAssociationError struct {
	Code *string `json:"code,omitempty"`
	Message *string `json:"message,omitempty"`
}

// NewProductlistProductListAssociationError instantiates a new ProductlistProductListAssociationError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductlistProductListAssociationError() *ProductlistProductListAssociationError {
	this := ProductlistProductListAssociationError{}
	return &this
}

// NewProductlistProductListAssociationErrorWithDefaults instantiates a new ProductlistProductListAssociationError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductlistProductListAssociationErrorWithDefaults() *ProductlistProductListAssociationError {
	this := ProductlistProductListAssociationError{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *ProductlistProductListAssociationError) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductlistProductListAssociationError) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *ProductlistProductListAssociationError) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *ProductlistProductListAssociationError) SetCode(v string) {
	o.Code = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *ProductlistProductListAssociationError) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductlistProductListAssociationError) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *ProductlistProductListAssociationError) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *ProductlistProductListAssociationError) SetMessage(v string) {
	o.Message = &v
}

func (o ProductlistProductListAssociationError) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProductlistProductListAssociationError) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	return toSerialize, nil
}

type NullableProductlistProductListAssociationError struct {
	value *ProductlistProductListAssociationError
	isSet bool
}

func (v NullableProductlistProductListAssociationError) Get() *ProductlistProductListAssociationError {
	return v.value
}

func (v *NullableProductlistProductListAssociationError) Set(val *ProductlistProductListAssociationError) {
	v.value = val
	v.isSet = true
}

func (v NullableProductlistProductListAssociationError) IsSet() bool {
	return v.isSet
}

func (v *NullableProductlistProductListAssociationError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductlistProductListAssociationError(val *ProductlistProductListAssociationError) *NullableProductlistProductListAssociationError {
	return &NullableProductlistProductListAssociationError{value: val, isSet: true}
}

func (v NullableProductlistProductListAssociationError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProductlistProductListAssociationError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

