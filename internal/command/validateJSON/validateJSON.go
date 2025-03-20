package validateJSON

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
)

// Execute выполняет команду валидация JSON файлов
func Execute(args []string) {
	if len(args) < 1 {
		fmt.Println("Укажите путь к JSON файлу в формате path={путь}")
		fmt.Println("Пример: program path=file.json")
		return
	}

	filePathArg := args[0]
	if !strings.HasPrefix(filePathArg, "path=") {
		fmt.Println("Аргумент должен начинаться с 'path='")
		fmt.Println("Пример: program path=file.json")
		return
	}

	filePath := strings.TrimPrefix(filePathArg, "path=")
	if filePath == "" {
		fmt.Println("Укажите путь к файлу после 'path='")
		return
	}

	err := validateJSON(filePath)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("JSON валиден")
}

// Help возвращает справку по команде
func Help() string {
	return "validateJSON path={file} - проверяет валидность JSON файла"
}

func validateJSON(filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("ошибка открытия файла: %v", err)
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	data, err := io.ReadAll(file)
	if err != nil {
		return fmt.Errorf("ошибка чтения файла: %v", err)
	}

	var v interface{}
	err = json.Unmarshal(data, &v)
	if err != nil {
		if syntaxErr, ok := err.(*json.SyntaxError); ok {
			content := string(data[:syntaxErr.Offset])
			line := strings.Count(content, "\n") + 1
			return fmt.Errorf("ошибка синтаксиса JSON на строке %d: %v", line, err)
		}
		return fmt.Errorf("ошибка проверки JSON: %v", err)
	}

	return nil
}
