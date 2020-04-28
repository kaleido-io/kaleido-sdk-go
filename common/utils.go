package common

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	eth "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"golang.org/x/crypto/sha3"
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

// RootNodeHash is the 0x0 hash for the root node in the idregistry
const RootNodeHash = "0x0000000000000000000000000000000000000000000000000000000000000000"

func encodePacked(tokens ...string) []byte {
<<<<<<< HEAD
	stringTy, _ := abi.NewType("string",[]abi.ArgumentMarshaling{})
=======

	components := []abi.ArgumentMarshaling{}
	stringTy, _ := abi.NewType("string", "type", components)
>>>>>>> da25172d9121b8c8c21fb5a9df8a2f00238d8346

	arguments := abi.Arguments{}
	argument := abi.Argument{
		Type: stringTy,
	}

	for range tokens {
		arguments = append(arguments, argument)
	}

	bytes, _ := arguments.Pack(tokens)
	return bytes
}

<<<<<<< HEAD
func keccak256(bytes []byte) string {
=======
func Keccak256(bytes []byte) string {
>>>>>>> da25172d9121b8c8c21fb5a9df8a2f00238d8346
	hash := sha3.NewLegacyKeccak256()
	hash.Write(bytes)

	var buf []byte
	buf = hash.Sum(buf)

	return hexutil.Encode(buf)
}

// ChildHash calculates the hash of a child node given a parent node
func ChildHash(parentHex string, child string) string {
	intermediateHash := Keccak256([]byte(child))
	toHash := parentHex + intermediateHash[2:]
	hexBytes, _ := hexutil.Decode(toHash)
	hash := Keccak256(hexBytes)

	return hash
}

// PathHash calculates the hash for a path in the registry
func PathHash(path string) (string, error) {
	if path == "" {
		return "", errors.New("Can't calculate hash for empty path")
	}
	if path == "/" {
		return RootNodeHash, nil
	}

	hash := RootNodeHash
	tokens := strings.Split(path, "/")
	for _, token := range tokens {
		if token != "" {
			hash = ChildHash(hash, token)
		}
	}

	return hash, nil
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
