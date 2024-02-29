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

// checks if the ProductlistDeleteProductListAssociationResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProductlistDeleteProductListAssociationResponse{}

// ProductlistDeleteProductListAssociationResponse struct for ProductlistDeleteProductListAssociationResponse
type ProductlistDeleteProductListAssociationResponse struct {
	Association *ProductlistProductListAssociation `json:"association,omitempty"`
	Errors []ProductlistProductListAssociationError `json:"errors,omitempty"`
}

// NewProductlistDeleteProductListAssociationResponse instantiates a new ProductlistDeleteProductListAssociationResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductlistDeleteProductListAssociationResponse() *ProductlistDeleteProductListAssociationResponse {
	this := ProductlistDeleteProductListAssociationResponse{}
	return &this
}

// NewProductlistDeleteProductListAssociationResponseWithDefaults instantiates a new ProductlistDeleteProductListAssociationResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductlistDeleteProductListAssociationResponseWithDefaults() *ProductlistDeleteProductListAssociationResponse {
	this := ProductlistDeleteProductListAssociationResponse{}
	return &this
}

// GetAssociation returns the Association field value if set, zero value otherwise.
func (o *ProductlistDeleteProductListAssociationResponse) GetAssociation() ProductlistProductListAssociation {
	if o == nil || IsNil(o.Association) {
		var ret ProductlistProductListAssociation
		return ret
	}
	return *o.Association
}

// GetAssociationOk returns a tuple with the Association field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductlistDeleteProductListAssociationResponse) GetAssociationOk() (*ProductlistProductListAssociation, bool) {
	if o == nil || IsNil(o.Association) {
		return nil, false
	}
	return o.Association, true
}

// HasAssociation returns a boolean if a field has been set.
func (o *ProductlistDeleteProductListAssociationResponse) HasAssociation() bool {
	if o != nil && !IsNil(o.Association) {
		return true
	}

	return false
}

// SetAssociation gets a reference to the given ProductlistProductListAssociation and assigns it to the Association field.
func (o *ProductlistDeleteProductListAssociationResponse) SetAssociation(v ProductlistProductListAssociation) {
	o.Association = &v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *ProductlistDeleteProductListAssociationResponse) GetErrors() []ProductlistProductListAssociationError {
	if o == nil || IsNil(o.Errors) {
		var ret []ProductlistProductListAssociationError
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductlistDeleteProductListAssociationResponse) GetErrorsOk() ([]ProductlistProductListAssociationError, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *ProductlistDeleteProductListAssociationResponse) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []ProductlistProductListAssociationError and assigns it to the Errors field.
func (o *ProductlistDeleteProductListAssociationResponse) SetErrors(v []ProductlistProductListAssociationError) {
	o.Errors = v
}

func (o ProductlistDeleteProductListAssociationResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProductlistDeleteProductListAssociationResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Association) {
		toSerialize["association"] = o.Association
	}
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	return toSerialize, nil
}

type NullableProductlistDeleteProductListAssociationResponse struct {
	value *ProductlistDeleteProductListAssociationResponse
	isSet bool
}

func (v NullableProductlistDeleteProductListAssociationResponse) Get() *ProductlistDeleteProductListAssociationResponse {
	return v.value
}

func (v *NullableProductlistDeleteProductListAssociationResponse) Set(val *ProductlistDeleteProductListAssociationResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableProductlistDeleteProductListAssociationResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableProductlistDeleteProductListAssociationResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductlistDeleteProductListAssociationResponse(val *ProductlistDeleteProductListAssociationResponse) *NullableProductlistDeleteProductListAssociationResponse {
	return &NullableProductlistDeleteProductListAssociationResponse{value: val, isSet: true}
}

func (v NullableProductlistDeleteProductListAssociationResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProductlistDeleteProductListAssociationResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


