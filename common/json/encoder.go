// +build !iterator

package json

import "encoding/json"

var Marshal = json.Marshal
var Unmarshal = json.Unmarshal
var Valid = json.Valid
