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

// checks if the ProductlistUpdateProductListResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProductlistUpdateProductListResponse{}

// ProductlistUpdateProductListResponse struct for ProductlistUpdateProductListResponse
type ProductlistUpdateProductListResponse struct {
	List                 *ProductlistProductListEntity `json:"list,omitempty"`
	Errors               []ProductlistProductListError `json:"errors,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ProductlistUpdateProductListResponse ProductlistUpdateProductListResponse

// NewProductlistUpdateProductListResponse instantiates a new ProductlistUpdateProductListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductlistUpdateProductListResponse() *ProductlistUpdateProductListResponse {
	this := ProductlistUpdateProductListResponse{}
	return &this
}

// NewProductlistUpdateProductListResponseWithDefaults instantiates a new ProductlistUpdateProductListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductlistUpdateProductListResponseWithDefaults() *ProductlistUpdateProductListResponse {
	this := ProductlistUpdateProductListResponse{}
	return &this
}

// GetList returns the List field value if set, zero value otherwise.
func (o *ProductlistUpdateProductListResponse) GetList() ProductlistProductListEntity {
	if o == nil || IsNil(o.List) {
		var ret ProductlistProductListEntity
		return ret
	}
	return *o.List
}

// GetListOk returns a tuple with the List field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductlistUpdateProductListResponse) GetListOk() (*ProductlistProductListEntity, bool) {
	if o == nil || IsNil(o.List) {
		return nil, false
	}
	return o.List, true
}

// HasList returns a boolean if a field has been set.
func (o *ProductlistUpdateProductListResponse) HasList() bool {
	if o != nil && !IsNil(o.List) {
		return true
	}

	return false
}

// SetList gets a reference to the given ProductlistProductListEntity and assigns it to the List field.
func (o *ProductlistUpdateProductListResponse) SetList(v ProductlistProductListEntity) {
	o.List = &v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *ProductlistUpdateProductListResponse) GetErrors() []ProductlistProductListError {
	if o == nil || IsNil(o.Errors) {
		var ret []ProductlistProductListError
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductlistUpdateProductListResponse) GetErrorsOk() ([]ProductlistProductListError, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *ProductlistUpdateProductListResponse) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []ProductlistProductListError and assigns it to the Errors field.
func (o *ProductlistUpdateProductListResponse) SetErrors(v []ProductlistProductListError) {
	o.Errors = v
}

func (o ProductlistUpdateProductListResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProductlistUpdateProductListResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.List) {
		toSerialize["list"] = o.List
	}
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ProductlistUpdateProductListResponse) UnmarshalJSON(data []byte) (err error) {
	varProductlistUpdateProductListResponse := _ProductlistUpdateProductListResponse{}

	err = json.Unmarshal(data, &varProductlistUpdateProductListResponse)

	if err != nil {
		return err
	}

	*o = ProductlistUpdateProductListResponse(varProductlistUpdateProductListResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "list")
		delete(additionalProperties, "errors")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *ProductlistUpdateProductListResponse) GetValue() interface{} {
	if o == nil || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}

// SetValue populates the value of well-known types
func (o *ProductlistUpdateProductListResponse) SetValue(value interface{}) {
	if o == nil || IsNil(value) {
		return
	}
	if IsNil(o.AdditionalProperties) {
		o.AdditionalProperties = map[string]interface{}{}
	}
	o.AdditionalProperties["value"] = value
	return
}

type NullableProductlistUpdateProductListResponse struct {
	value *ProductlistUpdateProductListResponse
	isSet bool
}

func (v NullableProductlistUpdateProductListResponse) Get() *ProductlistUpdateProductListResponse {
	return v.value
}

func (v *NullableProductlistUpdateProductListResponse) Set(val *ProductlistUpdateProductListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableProductlistUpdateProductListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableProductlistUpdateProductListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductlistUpdateProductListResponse(val *ProductlistUpdateProductListResponse) *NullableProductlistUpdateProductListResponse {
	return &NullableProductlistUpdateProductListResponse{value: val, isSet: true}
}

func (v NullableProductlistUpdateProductListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProductlistUpdateProductListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
