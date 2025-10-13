package main

import (
	"fmt"
	"strings"
)

// Утечка памяти: Основная проблема заключается в том, что при создании среза v[:100],
// новая строка justString продолжает ссылаться на исходный массив байт огромной строки v.
//  Это означает, что весь массив размером 1KB (1024 байта) не может быть освобожден сборщиком мусора,
//  хотя используется только 100 байт.

// Глобальная переменная: Использование глобальной переменной justString
// может привести к неожиданным побочным эффектам и усложняет тестирование.

// Отсутствие обработки ошибок: Нет проверки на то,
// что строка v действительно имеет длину не менее 100 символов.

// var justString string

// func someFunc() {
//   v := createHugeString(1 << 10)
//   justString = v[:100]
// }

// func main() {
//   someFunc()
// }

func createHugeString(size int) string { 
	return strings.Repeat("wb", size)
}

func someFunc() (string, error) {
    v := createHugeString(1 << 10)
    
    if len(v) < 100 {
        return "", fmt.Errorf("string too short: expected at least 100 bytes, got %d", len(v))
    }
    
    result := make([]byte, 100)
    copy(result, v[:100])
    
    return string(result), nil
}

func main() {
    justString, err := someFunc()
	
    if err != nil {
        fmt.Printf("Error: %v\n", err)
        return
    }
    
    fmt.Printf("Result: %s\n", justString)
    fmt.Printf("Length: %d\n", len(justString))
}