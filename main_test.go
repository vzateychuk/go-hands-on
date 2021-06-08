package main

import (
	"bufio"
	"bytes"
	"strings"
	"testing"
)

var dataSorted = `1
2
3
3
4
4
5
`
var expectedResult = `1
2
3
4
5
`
var dataNotSorted = `1
2
1
`

func TestOk(t *testing.T) {
	// arrange
	in := bufio.NewReader(strings.NewReader(dataSorted))
	out := new(bytes.Buffer)

	// act
	err := uniq(in, out)
	result := out.String()

	// assert не было ошибок
	if err != nil {
		t.Errorf("test OK failed: Error happend")
	}
	// assert значения совпадают с ожидаемыми
	if result != expectedResult {
		t.Errorf("test OK failed: Result not match expected:\nExpected:%v\nReceived:%v", expectedResult, result)
	}
}

func TestDataNotSorted(t *testing.T) {
	// arrange not sorted
	in := bufio.NewReader(strings.NewReader(dataNotSorted))
	out := new(bytes.Buffer)

	// act
	err := uniq(in, out)

	// assert что была ошибка
	if err == nil {
		t.Errorf("test TestDataNotSorted failed: %v\n", err.Error())
	}

}
