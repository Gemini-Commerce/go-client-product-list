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

// checks if the ProductlistGetProductListsCountResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProductlistGetProductListsCountResponse{}

// ProductlistGetProductListsCountResponse struct for ProductlistGetProductListsCountResponse
type ProductlistGetProductListsCountResponse struct {
	TotalSize *int64 `json:"totalSize,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ProductlistGetProductListsCountResponse ProductlistGetProductListsCountResponse

// NewProductlistGetProductListsCountResponse instantiates a new ProductlistGetProductListsCountResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductlistGetProductListsCountResponse() *ProductlistGetProductListsCountResponse {
	this := ProductlistGetProductListsCountResponse{}
	return &this
}

// NewProductlistGetProductListsCountResponseWithDefaults instantiates a new ProductlistGetProductListsCountResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductlistGetProductListsCountResponseWithDefaults() *ProductlistGetProductListsCountResponse {
	this := ProductlistGetProductListsCountResponse{}
	return &this
}

// GetTotalSize returns the TotalSize field value if set, zero value otherwise.
func (o *ProductlistGetProductListsCountResponse) GetTotalSize() int64 {
	if o == nil || IsNil(o.TotalSize) {
		var ret int64
		return ret
	}
	return *o.TotalSize
}

// GetTotalSizeOk returns a tuple with the TotalSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductlistGetProductListsCountResponse) GetTotalSizeOk() (*int64, bool) {
	if o == nil || IsNil(o.TotalSize) {
		return nil, false
	}
	return o.TotalSize, true
}

// HasTotalSize returns a boolean if a field has been set.
func (o *ProductlistGetProductListsCountResponse) HasTotalSize() bool {
	if o != nil && !IsNil(o.TotalSize) {
		return true
	}

	return false
}

// SetTotalSize gets a reference to the given int64 and assigns it to the TotalSize field.
func (o *ProductlistGetProductListsCountResponse) SetTotalSize(v int64) {
	o.TotalSize = &v
}

func (o ProductlistGetProductListsCountResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProductlistGetProductListsCountResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TotalSize) {
		toSerialize["totalSize"] = o.TotalSize
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ProductlistGetProductListsCountResponse) UnmarshalJSON(data []byte) (err error) {
	varProductlistGetProductListsCountResponse := _ProductlistGetProductListsCountResponse{}

	err = json.Unmarshal(data, &varProductlistGetProductListsCountResponse)

	if err != nil {
		return err
	}

	*o = ProductlistGetProductListsCountResponse(varProductlistGetProductListsCountResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "totalSize")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *ProductlistGetProductListsCountResponse) GetValue() interface{} {
	if o == nil || IsNil(o.Type) || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}
// SetValue populate the value of well-known types
func (o *ProductlistGetProductListsCountResponse) SetValue(value interface{}) {
	if o == nil || IsNil(o.Type) || IsNil(value) {
		return
	}
    if IsNil(o.AdditionalProperties) {
        o.AdditionalProperties = map[string]interface{}{}
    }
	o.AdditionalProperties["value"] = value
	return
}
type NullableProductlistGetProductListsCountResponse struct {
	value *ProductlistGetProductListsCountResponse
	isSet bool
}

func (v NullableProductlistGetProductListsCountResponse) Get() *ProductlistGetProductListsCountResponse {
	return v.value
}

func (v *NullableProductlistGetProductListsCountResponse) Set(val *ProductlistGetProductListsCountResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableProductlistGetProductListsCountResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableProductlistGetProductListsCountResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductlistGetProductListsCountResponse(val *ProductlistGetProductListsCountResponse) *NullableProductlistGetProductListsCountResponse {
	return &NullableProductlistGetProductListsCountResponse{value: val, isSet: true}
}

func (v NullableProductlistGetProductListsCountResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProductlistGetProductListsCountResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


