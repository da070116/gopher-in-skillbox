package main

import (
	"flag"
)

func m25Task1() bool {
	// m25Task1 -- программа для нахождения подстроки в кириллической подстроке.
	var haystack, needle string
	var symbolMatched bool

	flag.StringVar(&haystack, "str", "", "Строка для поиска")
	flag.StringVar(&needle, "substr", "", "Что искать")

	flag.Parse()

	runeHaystack := []rune(haystack)
	runeNeedle := []rune(needle)

	s := runeNeedle // временный срез для проверки вхождений, не трогая исходное значение

	for _, r := range runeHaystack {
		if len(s) == 0 {
			break
		}
		if r == s[0] {
			// если символы совпали, ставим флаг в true и уменьшаем размер среза
			symbolMatched = true
			s = s[1:]
		} else {
			// символы не совпали, сбрасываем флаг и восстанавливаем срез до исходного
			symbolMatched = false
			s = runeNeedle
		}
	}

	//// а в нормальном случае строки 18-36 заменяются на
	//symbolMatched = strings.Contains(haystack, needle)

	return symbolMatched
}
