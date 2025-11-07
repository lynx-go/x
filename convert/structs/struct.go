package structs

import "github.com/mitchellh/mapstructure"

func MapTo(in map[string]any, out interface{}) error {
	config := mapstructure.DecoderConfig{
		WeaklyTypedInput: true,
		TagName:          "json",
		Result:           out,
	}
	decoder, err := mapstructure.NewDecoder(&config)
	if err != nil {
		return err
	}
	return decoder.Decode(in)
}
