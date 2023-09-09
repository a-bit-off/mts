/*
Задание 1. System design.
Спроектируйте систему сбора ошибок.
*/

package main

import (
	"os"

	"ex01/pkg/mylogger"
)

func main() {
	log := mylogger.NewLogger(os.Stdout)
	log.Debug("sum debug")
	log.Info("sum information")
	log.Warn("sum warn")
	log.Error("sum error")
}
