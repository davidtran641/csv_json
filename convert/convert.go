package convert

import (
	"encoding/csv"
	"io"
	"log"
)

func Parse(reader io.Reader) []map[string]string {
	rows := readCSV(reader)
	rowsDict := convertHeader(rows)

	results := []map[string]string{}

	for _, row := range rowsDict {
		res := transformHeader(row, convert)
		if isEmptyDict(res) {
			continue
		}
		results = append(results, res)
	}

	return results
}

type ConvertFunc func(key string) string

func convert(key string) string {
	dict := map[string]string{
		"English":            "en",
		"Malay":              "ms",
		"Simplified Chinese": "zh",
		"Vietnamese":         "vi",
		"Thai":               "th",
		"Burmese (Unicode)":  "my",
		"Khmer":              "kh",
	}
	v, ok := dict[key]
	if ok {
		return v
	}
	log.Printf("Key not found: %s", key)
	return ""
}

func readCSV(reader io.Reader) [][]string {
	r := csv.NewReader(reader)
	result := [][]string{}
	for {
		row, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Printf("Error while reading %v", err)
			continue
		}
		result = append(result, row)
	}
	return result
}

func convertHeader(rows [][]string) []map[string]string {
	result := []map[string]string{}

	if len(rows) < 2 {
		return result
	}

	header := rows[0]
	for _, row := range rows[1:] {
		if isEmptyRow(row) {
			continue
		}
		dict := make(map[string]string)
		for c, val := range row {
			if c < len(header) {
				dict[header[c]] = val
			}
		}

		result = append(result, dict)
	}

	return result
}

func isEmptyRow(row []string) bool {
	for _, v := range row {
		if v != "" {
			return false
		}
	}
	return true
}

func transformHeader(dict map[string]string, convert ConvertFunc) map[string]string {
	result := make(map[string]string)
	for k, v := range dict {
		newKey := convert(k)
		if newKey != "" {
			result[newKey] = v
		}
	}
	return result
}

func isEmptyDict(dict map[string]string) bool {
	for _, v := range dict {
		if v != "" {
			return false
		}
	}
	return true
}
