// Пакет mathutils: утилиты для математических операций.
// В этом файле реализована функция вычисления факториала числа с поддержкой произвольно больших результатов (используется math/big).

package mathutils

import (
	"fmt"
	"math/big"
)

// Factorial вычисляет факториал n (n!).
// Принимает n типа int64. Для n < 0 возвращает ошибку.
// Результат возвращается как *big.Int, чтобы поддерживать большие значения.
func Factorial(n int64) (*big.Int, error) {
	if n < 0 {
		return nil, fmt.Errorf("negative input: %d", n)
	}

	result := big.NewInt(1)
	if n <= 1 {
		return result, nil
	}

	// tmp используется для умножения в цикле, чтобы не выделять новый big.Int на каждом шаге
	tmp := new(big.Int)
	for i := int64(2); i <= n; i++ {
		tmp.SetInt64(i)
		result.Mul(result, tmp)
	}

	return result, nil
}
