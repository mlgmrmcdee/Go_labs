package main

import (
	"bufio"
	"crypto/md5"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"hash"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("=== Утилита хэширования данных ===")

	stdin := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\nВыберите действие:")
		fmt.Println("1 - Вычислить хэш строки")
		fmt.Println("2 - Проверить целостность данных")
		fmt.Println("0 - Выход")

		// читаем строку и парсим выбор
		fmt.Print("Введите номер действия: ")
		line, err := stdin.ReadString('\n')
		if err != nil {
			fmt.Println("Ошибка чтения ввода:", err)
			return
		}
		line = strings.TrimSpace(line)
		choice, err := strconv.Atoi(line)
		if err != nil {
			fmt.Println("Неверный ввод, ожидается число. Попробуйте снова.")
			continue
		}

		switch choice {
		case 1:
			calcHash(stdin)
		case 2:
			verifyHash(stdin)
		case 0:
			fmt.Println("Выход из программы.")
			return
		default:
			fmt.Println("Неизвестная команда.")
		}
	}
}

// --- Функция для выбора алгоритма хэширования ---
func selectHashFunc(stdin *bufio.Reader) (hash.Hash, string) {
	fmt.Println("\nВыберите алгоритм хэширования:")
	fmt.Println("1 - MD5")
	fmt.Println("2 - SHA-256")
	fmt.Println("3 - SHA-512")

	fmt.Print("Введите номер: ")
	line, err := stdin.ReadString('\n')
	if err != nil {
		fmt.Println("Ошибка чтения ввода:", err)
		os.Exit(1)
	}
	line = strings.TrimSpace(line)
	choice, err := strconv.Atoi(line)
	if err != nil {
		fmt.Println("Неверный выбор. Используется SHA-256 по умолчанию.")
		return sha256.New(), "SHA-256"
	}

	switch choice {
	case 1:
		return md5.New(), "MD5"
	case 2:
		return sha256.New(), "SHA-256"
	case 3:
		return sha512.New(), "SHA-512"
	default:
		fmt.Println("Неверный выбор. Используется SHA-256 по умолчанию.")
		return sha256.New(), "SHA-256"
	}
}

// --- Вычисление хэша строки ---
func calcHash(stdin *bufio.Reader) {
	fmt.Print("\nВведите строку для хэширования: ")
	input, _ := stdin.ReadString('\n')
	input = strings.TrimSpace(input)

	h, name := selectHashFunc(stdin)
	h.Write([]byte(input))
	hashValue := hex.EncodeToString(h.Sum(nil))

	fmt.Printf("\nАлгоритм: %s\n", name)
	fmt.Printf("Хэш: %s\n", hashValue)
}

// --- Проверка целостности данных ---
func verifyHash(stdin *bufio.Reader) {
	fmt.Print("\nВведите исходную строку: ")
	input, _ := stdin.ReadString('\n')
	input = strings.TrimSpace(input)
	fmt.Print("Введите ожидаемый хэш: ")
	storedHash, _ := stdin.ReadString('\n')
	storedHash = strings.TrimSpace(storedHash)

	h, name := selectHashFunc(stdin)
	h.Write([]byte(input))
	computedHash := hex.EncodeToString(h.Sum(nil))

	fmt.Printf("\nАлгоритм: %s\n", name)
	fmt.Printf("Вычисленный хэш: %s\n", computedHash)

	if strings.EqualFold(computedHash, storedHash) {
		fmt.Println("✅ Целостность данных подтверждена: хэши совпадают.")
	} else {
		fmt.Println("❌ Нарушена целостность данных: хэши не совпадают.")
	}
}
