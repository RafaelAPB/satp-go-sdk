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

// checks if the BridgeInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BridgeInfo{}

// BridgeInfo Information about the bridge used for the token transfer.
type BridgeInfo struct {
	// The address of the token being transferred.
	TokenAddress *string `json:"tokenAddress,omitempty"`
}

// NewBridgeInfo instantiates a new BridgeInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBridgeInfo() *BridgeInfo {
	this := BridgeInfo{}
	return &this
}

// NewBridgeInfoWithDefaults instantiates a new BridgeInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBridgeInfoWithDefaults() *BridgeInfo {
	this := BridgeInfo{}
	return &this
}

// GetTokenAddress returns the TokenAddress field value if set, zero value otherwise.
func (o *BridgeInfo) GetTokenAddress() string {
	if o == nil || IsNil(o.TokenAddress) {
		var ret string
		return ret
	}
	return *o.TokenAddress
}

// GetTokenAddressOk returns a tuple with the TokenAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BridgeInfo) GetTokenAddressOk() (*string, bool) {
	if o == nil || IsNil(o.TokenAddress) {
		return nil, false
	}
	return o.TokenAddress, true
}

// HasTokenAddress returns a boolean if a field has been set.
func (o *BridgeInfo) HasTokenAddress() bool {
	if o != nil && !IsNil(o.TokenAddress) {
		return true
	}

	return false
}

// SetTokenAddress gets a reference to the given string and assigns it to the TokenAddress field.
func (o *BridgeInfo) SetTokenAddress(v string) {
	o.TokenAddress = &v
}

func (o BridgeInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BridgeInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TokenAddress) {
		toSerialize["tokenAddress"] = o.TokenAddress
	}
	return toSerialize, nil
}

type NullableBridgeInfo struct {
	value *BridgeInfo
	isSet bool
}

func (v NullableBridgeInfo) Get() *BridgeInfo {
	return v.value
}

func (v *NullableBridgeInfo) Set(val *BridgeInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableBridgeInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableBridgeInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBridgeInfo(val *BridgeInfo) *NullableBridgeInfo {
	return &NullableBridgeInfo{value: val, isSet: true}
}

func (v NullableBridgeInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBridgeInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


