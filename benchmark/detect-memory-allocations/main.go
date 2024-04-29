package main

import "practice_go/benchmark/basic"

func main() {
	basic.ConcatenateBuffer("first", "second")
	basic.ConcatenateJoin("first", "second")
}

// GODEBUG— это переменная среды, которая принимает список пар ключ-значение.
// Здесь мы сообщаем среде выполнения go генерировать трассировку стека для
// каждого выделения и освобождения. Затем мы добавляем "&>> trace.log"
// перенаправление стандартного вывода и стандартной ошибки в файл трассировки.log.
// Он создаст этот файл, если он не существует, а если он существует, к нему будут добавлены журналы.
// GODEBUG=allocfreetrace=1 ./allocDetect &>> trace.log
// cat trace.log | grep -n ./benchmark/basic/bench.go
