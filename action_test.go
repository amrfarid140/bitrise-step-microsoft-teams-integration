package main

import (
	"reflect"
	"testing"
)

func TestParseActions(t *testing.T) {
	var tests = []struct {
		input    string
		expected []Action
	}{
		{
			input:    `{"text": "some invalid text"}`,
			expected: nil,
		},
		{
			input: `[
				{
					"text": "Some text",
					"targets": [
						{
							"uri": "www.google.com"
						}
					]
				}
			]`,
			expected: []Action{
				{
					Text: "Some text",
					Targets: []ActionTarget{
						{
							URI: "www.google.com",
						},
					},
				},
			},
		},
		{
			input: `[
				{
					"text": "Some text",
					"targets": [
						{
							"uri": "www.google.com",
							"os": "default"
						}
					]
				}
			]`,
			expected: []Action{
				{
					Text: "Some text",
					Targets: []ActionTarget{
						{
							URI: "www.google.com",
							OS:  "default",
						},
					},
				},
			},
		},
		{
			input: `[
				{
					"text": "Some text",
					"targets": [
						{
							"uri": "www.google.com",
							"os": "android"
						},
						{
							"uri": "www.google.com",
							"os": "iOS"
						},
						{
							"uri": "www.google.com",
							"os": "windows"
						}
					]
				}
			]`,
			expected: []Action{
				{
					Text: "Some text",
					Targets: []ActionTarget{
						{
							URI: "www.google.com",
							OS:  "android",
						},
						{
							URI: "www.google.com",
							OS:  "iOS",
						},
						{
							URI: "www.google.com",
							OS:  "windows",
						},
					},
				},
			},
		},
	}

	for _, test := range tests {
		if output := parseActions(test.input); !reflect.DeepEqual(output, test.expected) {
			t.Errorf("Test failed: config input was %v, expected %v", test.input, test.expected)
		}
	}
}
