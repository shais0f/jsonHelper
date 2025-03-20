package validateJSON

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

// Тесты для функции Execute
func TestExecute(t *testing.T) {
	tests := []struct {
		name    string
		args    []string
		wantOut string
		wantErr bool
		setup   func() string
		cleanup func(string)
	}{
		{
			name:    "No arguments",
			args:    []string{},
			wantOut: "Укажите путь к JSON файлу в формате path={путь}\nПример: program path=file.json\n",
			wantErr: false,
		},
		{
			name:    "Missing path prefix",
			args:    []string{"file.json"},
			wantOut: "Аргумент должен начинаться с 'path='\nПример: program path=file.json\n",
			wantErr: false,
		},
		{
			name:    "Empty path",
			args:    []string{"path="},
			wantOut: "Укажите путь к файлу после 'path='\n",
			wantErr: false,
		},
		{
			name:    "Valid JSON",
			args:    []string{"path=test.json"},
			wantOut: "JSON валиден\n",
			setup: func() string {
				path := "test.json"
				err := os.WriteFile(path, []byte(`{"name": "John"}`), 0644)
				if err != nil {
					t.Fatal(err)
				}
				return "path=" + path
			},
			cleanup: func(path string) {
				os.Remove(strings.TrimPrefix(path, "path="))
			},
		},
		{
			name:    "Invalid JSON",
			args:    []string{"path=test.json"},
			wantOut: "ошибка синтаксиса JSON на строке 1: invalid character 'i' looking for beginning of value\n",
			setup: func() string {
				path := "test.json"
				err := os.WriteFile(path, []byte(`{"name": invalid}`), 0644)
				if err != nil {
					t.Fatal(err)
				}
				return "path=" + path
			},
			cleanup: func(path string) {
				os.Remove(strings.TrimPrefix(path, "path="))
			},
		},
		{
			name:    "Non-existent file",
			args:    []string{"path=nonexistent.json"},
			wantOut: "ошибка открытия файла: open nonexistent.json: no such file or directory\n",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var filePath string
			if tt.setup != nil {
				filePath = tt.setup()
				tt.args[0] = filePath
			}

			old := os.Stdout
			r, w, _ := os.Pipe()
			os.Stdout = w
			outC := make(chan string)

			go func() {
				var buf bytes.Buffer
				io.Copy(&buf, r)
				outC <- buf.String()
			}()

			Execute(tt.args)

			w.Close()
			os.Stdout = old
			out := <-outC

			if out != tt.wantOut {
				t.Errorf("Execute() output = %q, want %q", out, tt.wantOut)
			}

			if tt.cleanup != nil {
				tt.cleanup(filePath)
			}
		})
	}
}

// Тесты для функции validateJSON
func TestValidateJSON(t *testing.T) {
	tests := []struct {
		name    string
		content string
		wantErr bool
		errMsg  string
	}{
		{
			name:    "Valid JSON",
			content: `{"name": "John", "age": 30}`,
			wantErr: false,
		},
		{
			name:    "Invalid JSON - syntax error",
			content: `{"name": "John", "age": }`,
			wantErr: true,
			errMsg:  "ошибка синтаксиса JSON на строке 1: invalid character '}' looking for beginning of value",
		},
		{
			name:    "Empty JSON",
			content: `{}`,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tmpFile, err := os.CreateTemp("", "test*.json")
			if err != nil {
				t.Fatal(err)
			}
			defer os.Remove(tmpFile.Name())

			if _, err := tmpFile.Write([]byte(tt.content)); err != nil {
				t.Fatal(err)
			}
			tmpFile.Close()

			err = validateJSON(tmpFile.Name())

			if (err != nil) != tt.wantErr {
				t.Errorf("validateJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr && err.Error() != tt.errMsg {
				t.Errorf("validateJSON() error = %q, want %q", err.Error(), tt.errMsg)
			}
		})
	}
}

// Тест для функции Help
func TestHelp(t *testing.T) {
	expected := "validateJSON path={file} - проверяет валидность JSON файла"
	result := Help()
	if result != expected {
		t.Errorf("Help() = %q, want %q", result, expected)
	}
}
