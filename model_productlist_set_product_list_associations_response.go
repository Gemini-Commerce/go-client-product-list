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

// checks if the ProductlistSetProductListAssociationsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProductlistSetProductListAssociationsResponse{}

// ProductlistSetProductListAssociationsResponse struct for ProductlistSetProductListAssociationsResponse
type ProductlistSetProductListAssociationsResponse struct {
	Associations []ProductlistProductListAssociation `json:"associations,omitempty"`
	Errors []ProductlistProductListAssociationError `json:"errors,omitempty"`
}

// NewProductlistSetProductListAssociationsResponse instantiates a new ProductlistSetProductListAssociationsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductlistSetProductListAssociationsResponse() *ProductlistSetProductListAssociationsResponse {
	this := ProductlistSetProductListAssociationsResponse{}
	return &this
}

// NewProductlistSetProductListAssociationsResponseWithDefaults instantiates a new ProductlistSetProductListAssociationsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductlistSetProductListAssociationsResponseWithDefaults() *ProductlistSetProductListAssociationsResponse {
	this := ProductlistSetProductListAssociationsResponse{}
	return &this
}

// GetAssociations returns the Associations field value if set, zero value otherwise.
func (o *ProductlistSetProductListAssociationsResponse) GetAssociations() []ProductlistProductListAssociation {
	if o == nil || IsNil(o.Associations) {
		var ret []ProductlistProductListAssociation
		return ret
	}
	return o.Associations
}

// GetAssociationsOk returns a tuple with the Associations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductlistSetProductListAssociationsResponse) GetAssociationsOk() ([]ProductlistProductListAssociation, bool) {
	if o == nil || IsNil(o.Associations) {
		return nil, false
	}
	return o.Associations, true
}

// HasAssociations returns a boolean if a field has been set.
func (o *ProductlistSetProductListAssociationsResponse) HasAssociations() bool {
	if o != nil && !IsNil(o.Associations) {
		return true
	}

	return false
}

// SetAssociations gets a reference to the given []ProductlistProductListAssociation and assigns it to the Associations field.
func (o *ProductlistSetProductListAssociationsResponse) SetAssociations(v []ProductlistProductListAssociation) {
	o.Associations = v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *ProductlistSetProductListAssociationsResponse) GetErrors() []ProductlistProductListAssociationError {
	if o == nil || IsNil(o.Errors) {
		var ret []ProductlistProductListAssociationError
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductlistSetProductListAssociationsResponse) GetErrorsOk() ([]ProductlistProductListAssociationError, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *ProductlistSetProductListAssociationsResponse) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []ProductlistProductListAssociationError and assigns it to the Errors field.
func (o *ProductlistSetProductListAssociationsResponse) SetErrors(v []ProductlistProductListAssociationError) {
	o.Errors = v
}

func (o ProductlistSetProductListAssociationsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProductlistSetProductListAssociationsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Associations) {
		toSerialize["associations"] = o.Associations
	}
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	return toSerialize, nil
}

type NullableProductlistSetProductListAssociationsResponse struct {
	value *ProductlistSetProductListAssociationsResponse
	isSet bool
}

func (v NullableProductlistSetProductListAssociationsResponse) Get() *ProductlistSetProductListAssociationsResponse {
	return v.value
}

func (v *NullableProductlistSetProductListAssociationsResponse) Set(val *ProductlistSetProductListAssociationsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableProductlistSetProductListAssociationsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableProductlistSetProductListAssociationsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductlistSetProductListAssociationsResponse(val *ProductlistSetProductListAssociationsResponse) *NullableProductlistSetProductListAssociationsResponse {
	return &NullableProductlistSetProductListAssociationsResponse{value: val, isSet: true}
}

func (v NullableProductlistSetProductListAssociationsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProductlistSetProductListAssociationsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


