package json

import (
	jsoniter "github.com/json-iterator/go"
	"github.com/json-iterator/go/extra"
	"github.com/samber/lo"
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

func MustMarshalToString(v interface{}) string {
	return lo.Must1(json.MarshalToString(v))
}

func MustMarshal(v interface{}) []byte {
	return lo.Must1(json.Marshal(v))
}
