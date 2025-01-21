package env

import (
	"os"
	"testing"
)

func TestLoad(t *testing.T) {

	filename := ".env"

	tests := []struct {
		name      string
		envLine   string
		key       string
		wantValue string
		wantErr   bool
	}{
		{
			name:      "valid variable: doubleQuote value",
			envLine:   "VAR_TEST1=\"double quote value\"",
			key:       "VAR_TEST1",
			wantValue: "double quote value",
			wantErr:   false,
		},
		{
			name:      "valid variable: singleQuote value",
			envLine:   "VAR_TEST2='single quote value'",
			key:       "VAR_TEST2",
			wantValue: "single quote value",
			wantErr:   false,
		},
		{
			name:      "valid: value single in double",
			envLine:   "VAR_TEST9=\"value single' in double\"",
			key:       "VAR_TEST9",
			wantValue: "value single' in double",
			wantErr:   false,
		},
		{
			name:      "valid: value double in single",
			envLine:   "VAR_TEST10='value double\" in single'",
			key:       "VAR_TEST10",
			wantValue: "value double\" in single",
			wantErr:   false,
		},
		{
			name:      "valid: spaces before var name",
			envLine:   "   VAR_TEST5='space before var name'",
			key:       "VAR_TEST5",
			wantValue: "space before var name",
			wantErr:   false,
		},
		{
			name:      "valid: spaces after var name",
			envLine:   "VAR_TEST6    ='space after var name'",
			key:       "VAR_TEST6",
			wantValue: "space after var name",
			wantErr:   false,
		},
		{
			name:      "valid: spaces before value",
			envLine:   "VAR_TEST7=    'spaces before value'",
			key:       "VAR_TEST7",
			wantValue: "spaces before value",
			wantErr:   false,
		},
		{
			name:      "valid: spaces after value",
			envLine:   "VAR_TEST8='spaces after value'      ",
			key:       "VAR_TEST8",
			wantValue: "spaces after value",
			wantErr:   false,
		},
		{
			name:      "invalid: value start single end double",
			envLine:   "VAR_TEST3='start single end double\"",
			key:       "VAR_TEST3",
			wantValue: "",
			wantErr:   true,
		},
		{
			name:      "invalid: value start double end single",
			envLine:   "VAR_TEST4=\"start double end single'",
			key:       "VAR_TEST4",
			wantValue: "",
			wantErr:   true,
		},
		{
			name:      "invalid: key start _",
			envLine:   "_VAR_TEST='key start _'",
			key:       "_VAR_TEST",
			wantValue: "",
			wantErr:   true,
		},
		{
			name:      "invalid: key start number",
			envLine:   "1VAR_TEST='key start number",
			key:       "1VAR_TEST",
			wantValue: "",
			wantErr:   true,
		},
		{
			name:      "invalid: key contains a non alphanumeric or underscore char",
			envLine:   "VAR_TE.ST11='key contains a non alphanumeric or underscore char'",
			key:       "VAR_TE.ST11",
			wantValue: "",
			wantErr:   true,
		},
		{
			name:      "invalid: empty key",
			envLine:   "='empty key'",
			key:       "",
			wantValue: "",
			wantErr:   true,
		},
		{
			name:      "valid: empty value",
			envLine:   "VAR_TEST12=''",
			key:       "VAR_TEST12",
			wantValue: "",
			wantErr:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			file, err := os.Create(filename)
			if err != nil {
				t.Errorf("failed to create file: %v", err)
				return
			}
			defer file.Close()

			_, err = file.WriteString(tt.envLine)
			if err != nil {
				t.Errorf("failed to write to file: %v", err)
				return
			}

			gotErr := Load()
			gotValue := os.Getenv(tt.key)

			if (gotErr != nil) != tt.wantErr {
				t.Errorf("Load(), wantErr %v, got %v", tt.wantErr, gotErr)
				return
			}

			if gotValue != tt.wantValue {
				t.Errorf("Load(), want %v, got %v", tt.wantValue, gotValue)
				return
			}
		})

	}
	os.Remove(filename)
}

func TestRead(t *testing.T) {

	filename := ".env"

	tests := []struct {
		name      string
		envLine   string
		key       string
		wantValue string
		wantOk    bool
		wantErr   bool
	}{
		{
			name:      "valid variable: doubleQuote value",
			envLine:   "VAR_TEST1=\"double quote value\"",
			key:       "VAR_TEST1",
			wantValue: "double quote value",
			wantOk:    true,
			wantErr:   false,
		},
		{
			name:      "valid variable: singleQuote value",
			envLine:   "VAR_TEST2='single quote value'",
			key:       "VAR_TEST2",
			wantValue: "single quote value",
			wantOk:    true,
			wantErr:   false,
		},
		{
			name:      "valid: value single in double",
			envLine:   "VAR_TEST9=\"value single' in double\"",
			key:       "VAR_TEST9",
			wantValue: "value single' in double",
			wantOk:    true,
			wantErr:   false,
		},
		{
			name:      "valid: value double in single",
			envLine:   "VAR_TEST10='value double\" in single'",
			key:       "VAR_TEST10",
			wantValue: "value double\" in single",
			wantOk:    true,
			wantErr:   false,
		},
		{
			name:      "valid: spaces before var name",
			envLine:   "   VAR_TEST5='space before var name'",
			key:       "VAR_TEST5",
			wantValue: "space before var name",
			wantOk:    true,
			wantErr:   false,
		},
		{
			name:      "valid: spaces after var name",
			envLine:   "VAR_TEST6    ='space after var name'",
			key:       "VAR_TEST6",
			wantValue: "space after var name",
			wantOk:    true,
			wantErr:   false,
		},
		{
			name:      "valid: spaces before value",
			envLine:   "VAR_TEST7=    'spaces before value'",
			key:       "VAR_TEST7",
			wantValue: "spaces before value",
			wantOk:    true,
			wantErr:   false,
		},
		{
			name:      "valid: spaces after value",
			envLine:   "VAR_TEST8='spaces after value'      ",
			key:       "VAR_TEST8",
			wantValue: "spaces after value",
			wantOk:    true,
			wantErr:   false,
		},
		{
			name:      "invalid: value start single end double",
			envLine:   "VAR_TEST3='start single end double\"",
			key:       "VAR_TEST3",
			wantValue: "",
			wantOk:    false,
			wantErr:   true,
		},
		{
			name:      "invalid: value start double end single",
			envLine:   "VAR_TEST4=\"start double end single'",
			key:       "VAR_TEST4",
			wantValue: "",
			wantOk:    false,
			wantErr:   true,
		},
		{
			name:      "invalid: key start _",
			envLine:   "_VAR_TEST='key start _'",
			key:       "_VAR_TEST",
			wantValue: "",
			wantOk:    false,
			wantErr:   true,
		},
		{
			name:      "invalid: key start number",
			envLine:   "1VAR_TEST='key start number",
			key:       "1VAR_TEST",
			wantValue: "",
			wantOk:    false,
			wantErr:   true,
		},
		{
			name:      "invalid: key contains a non alphanumeric or underscore char",
			envLine:   "VAR_TE.ST11='key contains a non alphanumeric or underscore char'",
			key:       "VAR_TE.ST11",
			wantValue: "",
			wantOk:    false,
			wantErr:   true,
		},
		{
			name:      "invalid: empty key",
			envLine:   "='empty key'",
			key:       "",
			wantValue: "",
			wantOk:    false,
			wantErr:   true,
		},
		{
			name:      "valid: empty value",
			envLine:   "VAR_TEST12=''",
			key:       "VAR_TEST12",
			wantValue: "",
			wantOk:    true,
			wantErr:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			file, err := os.Create(filename)
			if err != nil {
				t.Errorf("failed to create file: %v", err)
				return
			}
			defer file.Close()

			_, err = file.WriteString(tt.envLine)
			if err != nil {
				t.Errorf("failed to write to file: %v", err)
				return
			}

			gotMap, gotErr := Read()
			gotValue, gotOk := gotMap[tt.key]
			if gotErr != nil {
				if (gotErr != nil) != tt.wantErr {
					t.Errorf("Read(), wantErr %v, got %v", tt.wantErr, gotErr)
					return
				}
			}

			if gotOk != tt.wantOk {
				t.Errorf("Read(), wantOk %v, got %v", tt.wantOk, gotOk)
				return
			}

			if gotValue != tt.wantValue {
				t.Errorf("Read(), want %v, got %v", tt.wantValue, gotValue)
				return
			}
		})

	}
	os.Remove(filename)
}
