// +build iterator

package json

import (
	"github.com/json-iterator/go"
)

var iter = jsoniter.ConfigCompatibleWithStandardLibrary

var Marshal = iter.Marshal
var Unmarshal = iter.Unmarshal
var Valid = iter.Valid
