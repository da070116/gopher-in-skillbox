package internal

func caesarEncode(rawString string, shift int) string {
	// caesarEncode - Зашифровать строку, используя шифр Цезаря
	encoded := make([]rune, len(rawString))
	for i, runicLetter := range rawString {
		if runicLetter-int32(shift) < 32 {
			encoded[i] = runicLetter
		} else {
			encoded[i] = runicLetter - int32(shift)
		}
	}
	return string(encoded)
}

func caesarDecode(rawString string, shift int) string {
	// caesarDecode - Расшифровать строку, используя шифр Цезаря
	decoded := make([]rune, len(rawString))
	for _, runicLetter := range rawString {
		if runicLetter-int32(shift) < 32 {
			decoded = append(decoded, runicLetter)
		} else {
			decoded = append(decoded, runicLetter+int32(shift))
		}
	}
	return string(decoded)
}
