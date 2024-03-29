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

// checks if the ProductlistSearchProductListsByIdsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProductlistSearchProductListsByIdsResponse{}

// ProductlistSearchProductListsByIdsResponse struct for ProductlistSearchProductListsByIdsResponse
type ProductlistSearchProductListsByIdsResponse struct {
	Results []ProductlistProductListEntity `json:"results,omitempty"`
}

// NewProductlistSearchProductListsByIdsResponse instantiates a new ProductlistSearchProductListsByIdsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductlistSearchProductListsByIdsResponse() *ProductlistSearchProductListsByIdsResponse {
	this := ProductlistSearchProductListsByIdsResponse{}
	return &this
}

// NewProductlistSearchProductListsByIdsResponseWithDefaults instantiates a new ProductlistSearchProductListsByIdsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductlistSearchProductListsByIdsResponseWithDefaults() *ProductlistSearchProductListsByIdsResponse {
	this := ProductlistSearchProductListsByIdsResponse{}
	return &this
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *ProductlistSearchProductListsByIdsResponse) GetResults() []ProductlistProductListEntity {
	if o == nil || IsNil(o.Results) {
		var ret []ProductlistProductListEntity
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductlistSearchProductListsByIdsResponse) GetResultsOk() ([]ProductlistProductListEntity, bool) {
	if o == nil || IsNil(o.Results) {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *ProductlistSearchProductListsByIdsResponse) HasResults() bool {
	if o != nil && !IsNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given []ProductlistProductListEntity and assigns it to the Results field.
func (o *ProductlistSearchProductListsByIdsResponse) SetResults(v []ProductlistProductListEntity) {
	o.Results = v
}

func (o ProductlistSearchProductListsByIdsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProductlistSearchProductListsByIdsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Results) {
		toSerialize["results"] = o.Results
	}
	return toSerialize, nil
}

type NullableProductlistSearchProductListsByIdsResponse struct {
	value *ProductlistSearchProductListsByIdsResponse
	isSet bool
}

func (v NullableProductlistSearchProductListsByIdsResponse) Get() *ProductlistSearchProductListsByIdsResponse {
	return v.value
}

func (v *NullableProductlistSearchProductListsByIdsResponse) Set(val *ProductlistSearchProductListsByIdsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableProductlistSearchProductListsByIdsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableProductlistSearchProductListsByIdsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductlistSearchProductListsByIdsResponse(val *ProductlistSearchProductListsByIdsResponse) *NullableProductlistSearchProductListsByIdsResponse {
	return &NullableProductlistSearchProductListsByIdsResponse{value: val, isSet: true}
}

func (v NullableProductlistSearchProductListsByIdsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProductlistSearchProductListsByIdsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


