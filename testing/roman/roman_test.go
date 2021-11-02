package roman

import "testing"

func TestToRoman(t *testing.T) {
	testCases := []struct {
		num    int
		expect string
		err    error
	}{
		{0, "", ErrOutOfRange},
		{1, "I", nil},
		{2, "II", nil},
		{3, "III", nil},
		{4, "IV", nil},
		{5, "V", nil},
		{6, "VI", nil},
		{7, "VII", nil},
		{8, "VIII", nil},
		{9, "IX", nil},
		{10, "X", nil},
		{50, "L", nil},
		{100, "C", nil},
		{500, "D", nil},
		{1000, "M", nil},
		{31, "XXXI", nil},
		{148, "CXLVIII", nil},
		{294, "CCXCIV", nil},
		{312, "CCCXII", nil},
		{421, "CDXXI", nil},
		{528, "DXXVIII", nil},
		{621, "DCXXI", nil},
		{782, "DCCLXXXII", nil},
		{870, "DCCCLXX", nil},
		{941, "CMXLI", nil},
		{1043, "MXLIII", nil},
		{1110, "MCX", nil},
		{1226, "MCCXXVI", nil},
		{1301, "MCCCI", nil},
		{1485, "MCDLXXXV", nil},
		{1509, "MDIX", nil},
		{1607, "MDCVII", nil},
		{1754, "MDCCLIV", nil},
		{1832, "MDCCCXXXII", nil},
		{1993, "MCMXCIII", nil},
		{2074, "MMLXXIV", nil},
		{2152, "MMCLII", nil},
		{2212, "MMCCXII", nil},
		{2343, "MMCCCXLIII", nil},
		{2499, "MMCDXCIX", nil},
		{2574, "MMDLXXIV", nil},
		{2646, "MMDCXLVI", nil},
		{2723, "MMDCCXXIII", nil},
		{2892, "MMDCCCXCII", nil},
		{2975, "MMCMLXXV", nil},
		{3051, "MMMLI", nil},
		{3185, "MMMCLXXXV", nil},
		{3250, "MMMCCL", nil},
		{3313, "MMMCCCXIII", nil},
		{3408, "MMMCDVIII", nil},
		{3501, "MMMDI", nil},
		{3610, "MMMDCX", nil},
		{3743, "MMMDCCXLIII", nil},
		{3844, "MMMDCCCXLIV", nil},
		{3888, "MMMDCCCLXXXVIII", nil},
		{3940, "MMMCMXL", nil},
		{3999, "MMMCMXCIX", nil},
		{4000, "", ErrOutOfRange},
	}

	for _, testCase := range testCases {
		got, err := ToRoman(testCase.num)
		if got != testCase.expect {
			t.Errorf("ToRoman(%d) expect:%s got:%s", testCase.num, testCase.expect, got)
		}

		if err != testCase.err {
			t.Errorf("ToRoman(%d) expect error:%v got:%v", testCase.num, testCase.err, err)
		}
	}
}
