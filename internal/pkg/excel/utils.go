package excel

const _lettersAmount = 26

// Returns an Excel column name corresponding to the given number.
func numberToExcelColumn(n int) string {
	if n <= 0 {
		return ""
	}
	n--
	return numberToExcelColumn(n/_lettersAmount) + string(rune('A'+n%_lettersAmount))
}
