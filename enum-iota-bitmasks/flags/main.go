package main

import "fmt"

type MyConfig uint8

const (
	VERBOSE MyConfig = 1 << iota
	CONFIG_FROM_DISK
	DATABASE_REQUIRED
	LOGGER_ACTIVATED
	DEBUG
	FLOAT_SUPPORT
	RECOVERY_MODE
	REBOOT_ON_FAILURE
)

func MyComplexFunction(conf MyConfig, databaseDsn string) {
	fmt.Printf("conf: %08b\n", conf)
	test := conf & REBOOT_ON_FAILURE
	fmt.Printf("test: %08b\n", test)

	// Показывает, что флаг DATABASE_REQUIRED включен или нет
	test2 := conf & DATABASE_REQUIRED
	fmt.Printf("test2: %08b\n", test2)

	// Переключить определенный битовый флаг
	// toggle FLOAT_SUPPORT => Activate
	conf = conf ^ FLOAT_SUPPORT
	test = conf & FLOAT_SUPPORT
	fmt.Printf("test : %08b\n", test)

	// Очистить определенный битовый флаг
	conf = conf &^ FLOAT_SUPPORT
	test = conf & FLOAT_SUPPORT
	fmt.Printf("test : %08b\n", test)
}

func main() {
	MyComplexFunction(VERBOSE|REBOOT_ON_FAILURE, "test")
}

// | - оператор побитового ИЛИ, установить флаг
// &^ - операция побитового И НЕ, очистить флаг
// ^ - операция побитового XOR, переключить флаг
// & - операция побитового И, проверить флаг
