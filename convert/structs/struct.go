package structs

import (
	"github.com/go-viper/mapstructure/v2"
)

func MapTo(in map[string]any, out interface{}, opts ...Option) error {
	config := mapstructure.DecoderConfig{
		WeaklyTypedInput: true,
		TagName:          "json",
		Result:           out,
	}
	for _, o := range opts {
		o(&config)
	}

	decoder, err := mapstructure.NewDecoder(&config)
	if err != nil {
		return err
	}
	return decoder.Decode(in)
}

type Option func(config *mapstructure.DecoderConfig)
