package structs

import (
	"testing"

	"github.com/go-viper/mapstructure/v2"
	"github.com/stretchr/testify/require"
)

func TestMapTo(t *testing.T) {
	in := map[string]any{
		"hello": "world",
	}
	out := &output{}

	err := MapTo(in, out, func(config *mapstructure.DecoderConfig) {
		config.TagName = "json2"
	})
	require.NoError(t, err)
	require.Equal(t, in["hello"], out.Hello)
}

type output struct {
	Hello string `json2:"hello"`
}
