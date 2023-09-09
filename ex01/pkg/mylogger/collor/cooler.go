package collor

import "fmt"

// ColorizeString ...
// Преобразвание цвета строки
// Принимает текст и цвет
// Возвращает цыетной текст
func ColorizeString(text string, color string) string {
	colorCodes := map[string]string{
		"black":  "30",
		"red":    "31",
		"green":  "32",
		"yellow": "33",
		"blue":   "34",
		"white":  "37",
	}

	colorCode, ok := colorCodes[color]
	if !ok {
		return text
	}

	return fmt.Sprintf("\033[%sm%s\033[0m", colorCode, text)
}
