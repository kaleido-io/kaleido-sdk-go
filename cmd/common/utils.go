package common

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"

	eth "github.com/ethereum/go-ethereum/common"
)

// PrintJSONObject print a single json object
func PrintJSONObject(v interface{}) error {
	b, err := json.Marshal(v)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(string(b))
	}
	return err
}

// PrintJSON print a json object or an array
func PrintJSON(v interface{}) error {
	val := reflect.ValueOf(v)
	if val.Kind() == reflect.Array {
		for i := 0; i < val.Len(); i++ {
			if err := PrintJSONObject(val.Index(i)); err != nil {
				return err
			}
		}
	} else {
		return PrintJSONObject(v)
	}
	return nil
}

// EthereumAddress implements the pflag Value interface for flags that require ethereum address
type EthereumAddress struct {
	address string
}

func (value *EthereumAddress) String() string {
	return value.address
}

// Set sets the value, returns an error if it is not a valid ethereum address
func (value *EthereumAddress) Set(address string) error {
	if !eth.IsHexAddress(address) {
		return errors.New("must be a valid ethereum address")
	}
	value.address = address
	return nil
}

// Type returns the type string for an ethereum address
func (value *EthereumAddress) Type() string {
	return "ethereum-address"
}
