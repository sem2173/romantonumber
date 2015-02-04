package romanToNumber

import "testing"

func TestRomanToNumber(t *testing.T) {
  tests := []struct {
    input string
    expected int
  } {
    {input: "I", expected: 1},
    {input: "II", expected: 2},
    {input: "III", expected: 3},
    {input: "V", expected: 5},
    {input: "IV", expected: 4},

    
  }

  for _, testCase := range tests {
    actual := RomanToNumber(testCase.input)
    expected := testCase.expected
    if actual != expected {
      t.Errorf("%d != %d", actual, expected)
    }
  }
}
