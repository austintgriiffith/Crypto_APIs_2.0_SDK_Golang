/*
CryptoAPIs

Crypto APIs 2.0 is a complex and innovative infrastructure layer that radically simplifies the development of any Blockchain and Crypto related applications. Organized around REST, Crypto APIs 2.0 can assist both novice Bitcoin/Ethereum enthusiasts and crypto experts with the development of their blockchain applications. Crypto APIs 2.0 provides unified endpoints and data, raw data, automatic tokens and coins forwardings, callback functionalities, and much more.

API version: 2.0.0
Contact: developers@cryptoapis.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cryptoapis

import (
	"encoding/json"
	"fmt"
)

// AddressTokensTransactionUnconfirmedToken - struct for AddressTokensTransactionUnconfirmedToken
type AddressTokensTransactionUnconfirmedToken struct {
	AddressTokensTransactionUnconfirmedErc20 *AddressTokensTransactionUnconfirmedErc20
	AddressTokensTransactionUnconfirmedErc721 *AddressTokensTransactionUnconfirmedErc721
	AddressTokensTransactionUnconfirmedOmni *AddressTokensTransactionUnconfirmedOmni
}

// AddressTokensTransactionUnconfirmedErc20AsAddressTokensTransactionUnconfirmedToken is a convenience function that returns AddressTokensTransactionUnconfirmedErc20 wrapped in AddressTokensTransactionUnconfirmedToken
func AddressTokensTransactionUnconfirmedErc20AsAddressTokensTransactionUnconfirmedToken(v *AddressTokensTransactionUnconfirmedErc20) AddressTokensTransactionUnconfirmedToken {
	return AddressTokensTransactionUnconfirmedToken{ AddressTokensTransactionUnconfirmedErc20: v}
}

// AddressTokensTransactionUnconfirmedErc721AsAddressTokensTransactionUnconfirmedToken is a convenience function that returns AddressTokensTransactionUnconfirmedErc721 wrapped in AddressTokensTransactionUnconfirmedToken
func AddressTokensTransactionUnconfirmedErc721AsAddressTokensTransactionUnconfirmedToken(v *AddressTokensTransactionUnconfirmedErc721) AddressTokensTransactionUnconfirmedToken {
	return AddressTokensTransactionUnconfirmedToken{ AddressTokensTransactionUnconfirmedErc721: v}
}

// AddressTokensTransactionUnconfirmedOmniAsAddressTokensTransactionUnconfirmedToken is a convenience function that returns AddressTokensTransactionUnconfirmedOmni wrapped in AddressTokensTransactionUnconfirmedToken
func AddressTokensTransactionUnconfirmedOmniAsAddressTokensTransactionUnconfirmedToken(v *AddressTokensTransactionUnconfirmedOmni) AddressTokensTransactionUnconfirmedToken {
	return AddressTokensTransactionUnconfirmedToken{ AddressTokensTransactionUnconfirmedOmni: v}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *AddressTokensTransactionUnconfirmedToken) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into AddressTokensTransactionUnconfirmedErc20
	err = json.Unmarshal(data, &dst.AddressTokensTransactionUnconfirmedErc20)
	if err == nil {
		jsonAddressTokensTransactionUnconfirmedErc20, _ := json.Marshal(dst.AddressTokensTransactionUnconfirmedErc20)
		if string(jsonAddressTokensTransactionUnconfirmedErc20) == "{}" { // empty struct
			dst.AddressTokensTransactionUnconfirmedErc20 = nil
		} else {
			match++
		}
	} else {
		dst.AddressTokensTransactionUnconfirmedErc20 = nil
	}

	// try to unmarshal data into AddressTokensTransactionUnconfirmedErc721
	err = json.Unmarshal(data, &dst.AddressTokensTransactionUnconfirmedErc721)
	if err == nil {
		jsonAddressTokensTransactionUnconfirmedErc721, _ := json.Marshal(dst.AddressTokensTransactionUnconfirmedErc721)
		if string(jsonAddressTokensTransactionUnconfirmedErc721) == "{}" { // empty struct
			dst.AddressTokensTransactionUnconfirmedErc721 = nil
		} else {
			match++
		}
	} else {
		dst.AddressTokensTransactionUnconfirmedErc721 = nil
	}

	// try to unmarshal data into AddressTokensTransactionUnconfirmedOmni
	err = json.Unmarshal(data, &dst.AddressTokensTransactionUnconfirmedOmni)
	if err == nil {
		jsonAddressTokensTransactionUnconfirmedOmni, _ := json.Marshal(dst.AddressTokensTransactionUnconfirmedOmni)
		if string(jsonAddressTokensTransactionUnconfirmedOmni) == "{}" { // empty struct
			dst.AddressTokensTransactionUnconfirmedOmni = nil
		} else {
			match++
		}
	} else {
		dst.AddressTokensTransactionUnconfirmedOmni = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.AddressTokensTransactionUnconfirmedErc20 = nil
		dst.AddressTokensTransactionUnconfirmedErc721 = nil
		dst.AddressTokensTransactionUnconfirmedOmni = nil

		return fmt.Errorf("Data matches more than one schema in oneOf(AddressTokensTransactionUnconfirmedToken)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("Data failed to match schemas in oneOf(AddressTokensTransactionUnconfirmedToken)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src AddressTokensTransactionUnconfirmedToken) MarshalJSON() ([]byte, error) {
	if src.AddressTokensTransactionUnconfirmedErc20 != nil {
		return json.Marshal(&src.AddressTokensTransactionUnconfirmedErc20)
	}

	if src.AddressTokensTransactionUnconfirmedErc721 != nil {
		return json.Marshal(&src.AddressTokensTransactionUnconfirmedErc721)
	}

	if src.AddressTokensTransactionUnconfirmedOmni != nil {
		return json.Marshal(&src.AddressTokensTransactionUnconfirmedOmni)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *AddressTokensTransactionUnconfirmedToken) GetActualInstance() (interface{}) {
	if obj.AddressTokensTransactionUnconfirmedErc20 != nil {
		return obj.AddressTokensTransactionUnconfirmedErc20
	}

	if obj.AddressTokensTransactionUnconfirmedErc721 != nil {
		return obj.AddressTokensTransactionUnconfirmedErc721
	}

	if obj.AddressTokensTransactionUnconfirmedOmni != nil {
		return obj.AddressTokensTransactionUnconfirmedOmni
	}

	// all schemas are nil
	return nil
}

type NullableAddressTokensTransactionUnconfirmedToken struct {
	value *AddressTokensTransactionUnconfirmedToken
	isSet bool
}

func (v NullableAddressTokensTransactionUnconfirmedToken) Get() *AddressTokensTransactionUnconfirmedToken {
	return v.value
}

func (v *NullableAddressTokensTransactionUnconfirmedToken) Set(val *AddressTokensTransactionUnconfirmedToken) {
	v.value = val
	v.isSet = true
}

func (v NullableAddressTokensTransactionUnconfirmedToken) IsSet() bool {
	return v.isSet
}

func (v *NullableAddressTokensTransactionUnconfirmedToken) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddressTokensTransactionUnconfirmedToken(val *AddressTokensTransactionUnconfirmedToken) *NullableAddressTokensTransactionUnconfirmedToken {
	return &NullableAddressTokensTransactionUnconfirmedToken{value: val, isSet: true}
}

func (v NullableAddressTokensTransactionUnconfirmedToken) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddressTokensTransactionUnconfirmedToken) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


