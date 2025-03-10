package json

import (
	jsoniter "github.com/json-iterator/go"
	"github.com/json-iterator/go/extra"
)

func init() {
	extra.RegisterFuzzyDecoders()
}

var json = jsoniter.ConfigCompatibleWithStandardLibrary

var Marshal = json.Marshal
var Unmarshal = json.Unmarshal
var MarshalIndent = json.MarshalIndent
var MarshalToString = json.MarshalToString
var UnmarshalFromString = json.UnmarshalFromString
var NewEncoder = json.NewEncoder
var Get = json.Get
var NewDecoder = json.NewDecoder
var Valid = json.Valid
