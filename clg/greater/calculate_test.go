package greater

//import (
//	"testing"
//)
//
//func Test_CLG_Greater(t *testing.T) {
//	testCases := []struct {
//		A        float64
//		B        float64
//		Expected float64
//	}{
//		{
//			A:        3.5,
//			B:        3.5,
//			Expected: 3.5,
//		},
//		{
//			A:        3.5,
//			B:        12.5,
//			Expected: 12.5,
//		},
//		{
//			A:        35.5,
//			B:        14.5,
//			Expected: 35.5,
//		},
//		{
//			A:        -3.5,
//			B:        7.5,
//			Expected: 7.5,
//		},
//		{
//			A:        12.5,
//			B:        4.5,
//			Expected: 12.5,
//		},
//		{
//			A:        17,
//			B:        65,
//			Expected: 65,
//		},
//		{
//			A:        65,
//			B:        17,
//			Expected: 65,
//		},
//	}
//
//	for i, testCase := range testCases {
//		output := testMaybeNewCollection(t).Greater(testCase.A, testCase.B)
//
//		if output != testCase.Expected {
//			t.Fatal("case", i+1, "expected", testCase.Expected, "got", output)
//		}
//	}
//}
