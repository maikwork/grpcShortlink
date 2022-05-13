package domain

import (
	"fmt"
	"testing"
)

type testchk struct {
	t    string
	link string
	resp bool
}

func TestCheckLink(t *testing.T) {
	tests := [...]testchk{
		{"Латиница нижний регистр", "qwerrqw", true},
		{"Только цифры", "12345", true},
		{"Латиница и цифры", "asdf1234", true},
		{"Латиница верхниго регистра", "ASDFG", true},
		{"Латиница и нижней подчеркивание", "asdf_", true},
		{"Пустая строка", "", false},
		{"Строка с некоректными символами", "йцукеукйе", false},
		{"Строка превышающая 10 элементов", "qwertyasdfghq", false},
	}

	for _, v := range tests {
		chk := checkLink(v.link)
		if chk != v.resp {
			fmt.Errorf("%v: %v", v.t, chk)
		}
	}
}
