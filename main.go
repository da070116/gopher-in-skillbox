package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

func manage(attrs []string) (result string) {
	// manage -- возвращает результат в зависимости от числа аргументов на входе
	switch {
	case len(attrs) == 1 || len(attrs) == 2:
		result = readFileData(attrs...)
	case len(attrs) == 3:
		data := readFileData(attrs[:2]...)
		writeToFile(attrs[2], data)

	default:
		panic("Error: use 1, 2 or 3 args")
	}
	return
}

func writeToFile(fileName string, data string) {
	file, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)
	fmt.Printf("write %+v into %v\n", data, fileName)
	_, err = file.WriteString(data)
	if err != nil {
		panic(err)
	}

}

func readFromFile(fileName string) []byte {
	// проверить наличие:
	f, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("Файл %v не найден\n", fileName)
		return nil
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			panic(err)
		}
	}(f)

	// проверить размер:
	fileInfo, err := f.Stat()
	if err != nil || fileInfo.Size() == 0 {
		fmt.Printf("Файл %v пуст\n", fileInfo.Name())
		return nil
	}
	// считать в буфер по размеру файла:
	buf := make([]byte, fileInfo.Size())
	if _, err := io.ReadFull(f, buf); err != nil {
		panic(err)
	}
	return buf
}

func readFileData(fileNames ...string) (result string) {
	//readFileData -- прочитать данные из одного или нескольких файлов, результат объединить в строку
	data := make([]string, len(fileNames))
	for _, name := range fileNames {
		fileContent := readFromFile(name)
		if fileContent != nil {
			data = append(data, string(fileContent))
		}
	}
	result = strings.Join(data, "\n")
	result = strings.Trim(result, "\n")
	return
}

func main() {
	// получить параметры без имени из аргументов запуска
	flag.Parse()
	var rawParameters = flag.Args()

	fmt.Printf("%v\n", manage(rawParameters))
}
