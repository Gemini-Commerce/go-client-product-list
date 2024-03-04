/*
Collection Service

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1
Contact: info@gemini-commerce.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package product-list

import (
	"encoding/json"
)

// checks if the ProductlistDeleteProductListResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProductlistDeleteProductListResponse{}

// ProductlistDeleteProductListResponse struct for ProductlistDeleteProductListResponse
type ProductlistDeleteProductListResponse struct {
	List *ProductlistProductListEntity `json:"list,omitempty"`
	Errors []ProductlistProductListError `json:"errors,omitempty"`
}

// NewProductlistDeleteProductListResponse instantiates a new ProductlistDeleteProductListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductlistDeleteProductListResponse() *ProductlistDeleteProductListResponse {
	this := ProductlistDeleteProductListResponse{}
	return &this
}

// NewProductlistDeleteProductListResponseWithDefaults instantiates a new ProductlistDeleteProductListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductlistDeleteProductListResponseWithDefaults() *ProductlistDeleteProductListResponse {
	this := ProductlistDeleteProductListResponse{}
	return &this
}

// GetList returns the List field value if set, zero value otherwise.
func (o *ProductlistDeleteProductListResponse) GetList() ProductlistProductListEntity {
	if o == nil || IsNil(o.List) {
		var ret ProductlistProductListEntity
		return ret
	}
	return *o.List
}

// GetListOk returns a tuple with the List field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductlistDeleteProductListResponse) GetListOk() (*ProductlistProductListEntity, bool) {
	if o == nil || IsNil(o.List) {
		return nil, false
	}
	return o.List, true
}

// HasList returns a boolean if a field has been set.
func (o *ProductlistDeleteProductListResponse) HasList() bool {
	if o != nil && !IsNil(o.List) {
		return true
	}

	return false
}

// SetList gets a reference to the given ProductlistProductListEntity and assigns it to the List field.
func (o *ProductlistDeleteProductListResponse) SetList(v ProductlistProductListEntity) {
	o.List = &v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *ProductlistDeleteProductListResponse) GetErrors() []ProductlistProductListError {
	if o == nil || IsNil(o.Errors) {
		var ret []ProductlistProductListError
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductlistDeleteProductListResponse) GetErrorsOk() ([]ProductlistProductListError, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *ProductlistDeleteProductListResponse) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []ProductlistProductListError and assigns it to the Errors field.
func (o *ProductlistDeleteProductListResponse) SetErrors(v []ProductlistProductListError) {
	o.Errors = v
}

func (o ProductlistDeleteProductListResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProductlistDeleteProductListResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.List) {
		toSerialize["list"] = o.List
	}
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	return toSerialize, nil
}

type NullableProductlistDeleteProductListResponse struct {
	value *ProductlistDeleteProductListResponse
	isSet bool
}

func (v NullableProductlistDeleteProductListResponse) Get() *ProductlistDeleteProductListResponse {
	return v.value
}

func (v *NullableProductlistDeleteProductListResponse) Set(val *ProductlistDeleteProductListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableProductlistDeleteProductListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableProductlistDeleteProductListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductlistDeleteProductListResponse(val *ProductlistDeleteProductListResponse) *NullableProductlistDeleteProductListResponse {
	return &NullableProductlistDeleteProductListResponse{value: val, isSet: true}
}

func (v NullableProductlistDeleteProductListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProductlistDeleteProductListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


