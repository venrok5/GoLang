package main

import (
	"bufio"
	"bytes"
	"strings"
	"testing"
)

var testOk = "1\n2\n3\n" // \n потому что у нас когда вводим надо жать ентер
var testOkResult = "1\n2\n3\n"

func TestOk(t *testing.T) {
	in := bufio.NewReader(strings.NewReader(testOk))
	out := new(bytes.Buffer)
	err := uniq(in, out)
	if err != nil {
		t.Errorf("test for OK Failed 0 - error")
	}
	result := out.String()
	if result != testOkResult {
		t.Errorf("test for OK Failed 1 - results not match\n %v %v",
			result, testOkResult)
	}
}

// В тесте мы запускаем нашу функцию на тестовых данных. Если получаем ошибку, говорим, что тест сфейлился.
// Если ошибки не было, проверим, что результат работы программы корректный, сверившись с образцом.

// Чтобы запустить тестирование, надо ввести в командной строке go test − v. v — это значит verbals, то есть
// мы будем видеть результаты: какие тесты заработали, а какие закончились неудачей. Убедитесь, что тест
// проходит корректно.

// Но это не все, что можно сделать. Еще остался случай, когда мы должны вернуть ошибку, если поток ввода
// не отсортирован. Давайте напишем еще один тест.

var testFail = "1\n2\n1\n"

func TestForError(t *testing.T) {
	in := bufio.NewReader(strings.NewReader(testFail))
	out := new(bytes.Buffer)
	err := uniq(in, out)
	if err == nil {
		t.Errorf("test for OK Failed 2 - error: %v", err)
	}
}

// В него мы будем подавать заранее неправильные, то есть не отсортированные данные, и говорить, что тест
// сфейлился, если программа отрабатывает без ошибки, то есть не замечает некорректность ввода.

// Запустите тесты еще раз. Теперь, мы можем быть не просто уверены, что программа работает корректно, но и
// что мы сможем быстро проверить её корректность, даже если будем вносить в неё изменения, потому что у нас
// есть тесты.
