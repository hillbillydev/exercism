package romannumerals

import "testing"

func TestConvert(t *testing.T) {
	tc := append(romanNumeralTests, []romanNumeralTest{
		{0, "", true},
		{-1, "", true},
		{3001, "", true},
	}...)

	for _, test := range tc {
		actual, err := Convert(test.arabic)
		if err == nil && test.hasError {
			t.Errorf("Convert(%d) should return an error.", test.arabic)
			continue
		}
		if err != nil && !test.hasError {
			var _ error = err
			t.Errorf("Convert(%d) should not return an error.", test.arabic)
			continue
		}
		if !test.hasError && actual != test.roman {
			t.Errorf("Convert(%d): %s, expected %s", test.arabic, actual, test.roman)
		}
	}
}

func BenchmarkConvert(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range romanNumeralTests {
			Convert(test.arabic)
		}
	}
}
