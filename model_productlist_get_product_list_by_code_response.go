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

// checks if the ProductlistGetProductListByCodeResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProductlistGetProductListByCodeResponse{}

// ProductlistGetProductListByCodeResponse struct for ProductlistGetProductListByCodeResponse
type ProductlistGetProductListByCodeResponse struct {
	List *ProductlistProductListEntity `json:"list,omitempty"`
}

// NewProductlistGetProductListByCodeResponse instantiates a new ProductlistGetProductListByCodeResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductlistGetProductListByCodeResponse() *ProductlistGetProductListByCodeResponse {
	this := ProductlistGetProductListByCodeResponse{}
	return &this
}

// NewProductlistGetProductListByCodeResponseWithDefaults instantiates a new ProductlistGetProductListByCodeResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductlistGetProductListByCodeResponseWithDefaults() *ProductlistGetProductListByCodeResponse {
	this := ProductlistGetProductListByCodeResponse{}
	return &this
}

// GetList returns the List field value if set, zero value otherwise.
func (o *ProductlistGetProductListByCodeResponse) GetList() ProductlistProductListEntity {
	if o == nil || IsNil(o.List) {
		var ret ProductlistProductListEntity
		return ret
	}
	return *o.List
}

// GetListOk returns a tuple with the List field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductlistGetProductListByCodeResponse) GetListOk() (*ProductlistProductListEntity, bool) {
	if o == nil || IsNil(o.List) {
		return nil, false
	}
	return o.List, true
}

// HasList returns a boolean if a field has been set.
func (o *ProductlistGetProductListByCodeResponse) HasList() bool {
	if o != nil && !IsNil(o.List) {
		return true
	}

	return false
}

// SetList gets a reference to the given ProductlistProductListEntity and assigns it to the List field.
func (o *ProductlistGetProductListByCodeResponse) SetList(v ProductlistProductListEntity) {
	o.List = &v
}

func (o ProductlistGetProductListByCodeResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProductlistGetProductListByCodeResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.List) {
		toSerialize["list"] = o.List
	}
	return toSerialize, nil
}

type NullableProductlistGetProductListByCodeResponse struct {
	value *ProductlistGetProductListByCodeResponse
	isSet bool
}

func (v NullableProductlistGetProductListByCodeResponse) Get() *ProductlistGetProductListByCodeResponse {
	return v.value
}

func (v *NullableProductlistGetProductListByCodeResponse) Set(val *ProductlistGetProductListByCodeResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableProductlistGetProductListByCodeResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableProductlistGetProductListByCodeResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductlistGetProductListByCodeResponse(val *ProductlistGetProductListByCodeResponse) *NullableProductlistGetProductListByCodeResponse {
	return &NullableProductlistGetProductListByCodeResponse{value: val, isSet: true}
}

func (v NullableProductlistGetProductListByCodeResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProductlistGetProductListByCodeResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


