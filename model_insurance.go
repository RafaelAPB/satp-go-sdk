/*
SATP Gateway Client (Business Logic Orchestrator)

SATP is a protocol operating between two gateways that conducts the transfer of a digital asset from one gateway to another. The protocol establishes a secure channel between the endpoints and implements a 2-phase commit to ensure the properties of transfer atomicity, consistency, isolation and durability.  This API defines the gateway client facing API (business logic orchestrator, or BLO), which is named API-Type 1 in the SATP-Core specification.  **Additional Resources**: - [Proposed SATP Charter](https://datatracker.ietf.org/doc/charter-ietf-satp/) - [SATP Core draft](https://datatracker.ietf.org/doc/draft-ietf-satp-core) - [SATP Crash Recovery draft](https://datatracker.ietf.org/doc/draft-belchior-satp-gateway-recovery/) - [SATP Architecture draft](https://datatracker.ietf.org/doc/draft-ietf-satp-architecture/) - [SATP Use-Cases draft](https://datatracker.ietf.org/doc/draft-ramakrishna-sat-use-cases/) - [SATP Data sharing draft](https://datatracker.ietf.org/doc/draft-ramakrishna-satp-data-sharing) - [SATP View Addresses draft](https://datatracker.ietf.org/doc/draft-ramakrishna-satp-views-addresses)

API version: v0.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package generated

import (
	"encoding/json"
)

// checks if the Insurance type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Insurance{}

// Insurance struct for Insurance
type Insurance struct {
	// The state of insurance applicability for the transaction.
	State *string `json:"state,omitempty"`
	// The fee amount for insurance, represented in USD.
	FeeAmountUsd *string `json:"feeAmountUsd,omitempty"`
}

// NewInsurance instantiates a new Insurance object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInsurance() *Insurance {
	this := Insurance{}
	return &this
}

// NewInsuranceWithDefaults instantiates a new Insurance object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInsuranceWithDefaults() *Insurance {
	this := Insurance{}
	return &this
}

// GetState returns the State field value if set, zero value otherwise.
func (o *Insurance) GetState() string {
	if o == nil || IsNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Insurance) GetStateOk() (*string, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *Insurance) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *Insurance) SetState(v string) {
	o.State = &v
}

// GetFeeAmountUsd returns the FeeAmountUsd field value if set, zero value otherwise.
func (o *Insurance) GetFeeAmountUsd() string {
	if o == nil || IsNil(o.FeeAmountUsd) {
		var ret string
		return ret
	}
	return *o.FeeAmountUsd
}

// GetFeeAmountUsdOk returns a tuple with the FeeAmountUsd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Insurance) GetFeeAmountUsdOk() (*string, bool) {
	if o == nil || IsNil(o.FeeAmountUsd) {
		return nil, false
	}
	return o.FeeAmountUsd, true
}

// HasFeeAmountUsd returns a boolean if a field has been set.
func (o *Insurance) HasFeeAmountUsd() bool {
	if o != nil && !IsNil(o.FeeAmountUsd) {
		return true
	}

	return false
}

// SetFeeAmountUsd gets a reference to the given string and assigns it to the FeeAmountUsd field.
func (o *Insurance) SetFeeAmountUsd(v string) {
	o.FeeAmountUsd = &v
}

func (o Insurance) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Insurance) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	if !IsNil(o.FeeAmountUsd) {
		toSerialize["feeAmountUsd"] = o.FeeAmountUsd
	}
	return toSerialize, nil
}

type NullableInsurance struct {
	value *Insurance
	isSet bool
}

func (v NullableInsurance) Get() *Insurance {
	return v.value
}

func (v *NullableInsurance) Set(val *Insurance) {
	v.value = val
	v.isSet = true
}

func (v NullableInsurance) IsSet() bool {
	return v.isSet
}

func (v *NullableInsurance) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInsurance(val *Insurance) *NullableInsurance {
	return &NullableInsurance{value: val, isSet: true}
}

func (v NullableInsurance) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInsurance) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


