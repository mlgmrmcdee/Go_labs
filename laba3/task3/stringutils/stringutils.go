package stringutils

// Reverse переворачивает строку
func Reverse(s string) string {
    runes := []rune(s) // используем руны, чтобы корректно работать с Unicode
    n := len(runes)
    for i := 0; i < n/2; i++ {
        runes[i], runes[n-1-i] = runes[n-1-i], runes[i]
    }
    return string(runes)
}
