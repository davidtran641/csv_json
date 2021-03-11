package convert

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

func TestConvertKey(t *testing.T) {
	cases := []struct {
		input string
		want  string
	}{
		{
			"English",
			"en",
		},
	}

	for _, c := range cases {
		tc := c
		t.Run(tc.input, func(t *testing.T) {
			got := convert(tc.input)
			want := tc.want
			assertEqual(t, want, got)
		})
	}
}

func TestReadCsv(t *testing.T) {
	cases := []struct {
		input string
		want  [][]string
	}{
		{
			"a,b,c",
			[][]string{{"a", "b", "c"}},
		},
		{
			"a,b,c\n1,2,34",
			[][]string{{"a", "b", "c"}, {"1", "2", "34"}},
		},
	}

	for _, c := range cases {
		tc := c
		t.Run(tc.input, func(t *testing.T) {
			reader := strings.NewReader(tc.input)
			got := readCSV(reader)
			want := tc.want
			assertEqual(t, want, got)
		})
	}
}

func TestConvertHeader(t *testing.T) {
	cases := []struct {
		input [][]string
		want  []map[string]string
	}{
		{
			[][]string{{"a", "b", "c"}},
			[]map[string]string{},
		},
		{
			[][]string{{"a", "b", "c"}, {"1", "2", "34"}},
			[]map[string]string{map[string]string{"a": "1", "b": "2", "c": "34"}},
		},
	}

	for i, c := range cases {
		tc := c
		t.Run(fmt.Sprintf("test:%d", i), func(t *testing.T) {
			got := convertHeader(tc.input)
			want := tc.want
			assertEqual(t, want, got)
		})
	}
}

func TestTransformHeader(t *testing.T) {
	convert := func(key string) string {
		switch key {
		case "a":
			return "a1"
		case "b":
			return "b1"
		default:
			return ""
		}

	}
	cases := []struct {
		input map[string]string
		want  map[string]string
	}{
		{
			map[string]string{"a": "1", "b": "2", "c": "34"},
			map[string]string{"a1": "1", "b1": "2"},
		},
	}

	for i, c := range cases {
		tc := c
		t.Run(fmt.Sprintf("test:%d", i), func(t *testing.T) {
			got := transformHeader(tc.input, convert)
			want := tc.want
			assertEqual(t, want, got)
		})
	}
}

func testParse(t *testing.T) {
	cases := []struct {
		input string
		want  []map[string]string
	}{
		{
			"English,Malay,NonKey\n1,2,3\na,b,c",
			[]map[string]string{
				map[string]string{"English": "1", "Malay": "2"},
				map[string]string{"English": "a", "Malay": "b"},
			},
		},
	}

	for i, c := range cases {
		tc := c
		t.Run(fmt.Sprintf("test:%d", i), func(t *testing.T) {
			reader := strings.NewReader(tc.input)
			got := Parse(reader)
			want := tc.want
			assertEqual(t, want, got)
		})
	}
}

func TestCheckEmptyRow(t *testing.T) {
	cases := []struct {
		input []string
		want  bool
	}{
		{
			[]string{"", ""},
			true,
		},
		{
			[]string{"", "a"},
			false,
		},
	}

	for i, c := range cases {
		tc := c
		t.Run(fmt.Sprintf("test:%d", i), func(t *testing.T) {
			got := isEmptyRow(tc.input)
			want := tc.want
			assertEqual(t, want, got)
		})
	}
}

func TestCheckEmptyDict(t *testing.T) {
	cases := []struct {
		input map[string]string
		want  bool
	}{
		{
			map[string]string{"en": "", "vi": ""},
			true,
		},
		{
			map[string]string{"en": "a", "vi": ""},
			false,
		},
	}

	for i, c := range cases {
		tc := c
		t.Run(fmt.Sprintf("test:%d", i), func(t *testing.T) {
			got := isEmptyDict(tc.input)
			want := tc.want
			assertEqual(t, want, got)
		})
	}
}

func assertEqual(t *testing.T, want interface{}, got interface{}) {
	t.Helper()
	if reflect.DeepEqual(want, got) {
		return
	}

	t.Errorf("want %v, but got %v", want, got)
}
