/*
Задание 1. System design.
Спроектируйте систему сбора ошибок.
*/

package main

import (
	"ex01/pkg/mylogger"
	"os"
)

func main() {
	log := mylogger.NewLogger(os.Stdout)
	log.Debug("sum debug")
	log.Info("sum information")
	log.Warn("sum warn")
	log.Error("sum error")
