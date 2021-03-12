package config

import (
	"errors"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadCsv(t *testing.T) {
	cases := []struct {
		input string
		want  *Config
		err   error
	}{
		{
			"",
			nil,
			errors.New(""),
		},
		{
			"{}",
			&Config{},
			nil,
		},
		{
			`{
				"header_mapping" : {
					"en" :"en-new"
				}
			}`,
			&Config{HeaderMapping: map[string]string{"en": "en-new"}},
			nil,
		},
	}

	for _, c := range cases {
		tc := c
		t.Run(tc.input, func(t *testing.T) {
			reader := strings.NewReader(tc.input)
			got, err := parseConfig(reader)
			want := tc.want
			assert.Equal(t, want, got)

			if tc.err != nil {
				assert.Error(t, err)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}
