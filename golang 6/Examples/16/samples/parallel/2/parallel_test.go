package parallel

import "testing"

func TestTLog(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name  string
		value int
	}{
		{name: "test 1", value: 1},
		{name: "test 2", value: 2},
		{name: "test 3", value: 3},
		{name: "test 4", value: 4},
	}
	for _, tc := range tests {
		value := tc.value // shadowing
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()  //  +++
			// Here you test tc.value against a test function.
			t.Log(value)
		})
	}
}
/*
=== RUN   TestTLog
=== PAUSE TestTLog
=== CONT  TestTLog
=== RUN   TestTLog/test_1
=== PAUSE TestTLog/test_1
=== RUN   TestTLog/test_2
=== PAUSE TestTLog/test_2
=== RUN   TestTLog/test_3
=== PAUSE TestTLog/test_3
=== RUN   TestTLog/test_4
=== PAUSE TestTLog/test_4
=== CONT  TestTLog/test_1
    parallel_test.go:21: 1
=== CONT  TestTLog/test_3
=== CONT  TestTLog/test_2
=== CONT  TestTLog/test_4
=== CONT  TestTLog/test_3
    parallel_test.go:21: 3
=== CONT  TestTLog/test_2
    parallel_test.go:21: 2
=== CONT  TestTLog/test_4
    parallel_test.go:21: 4
--- PASS: TestTLog (0.00s)
    --- PASS: TestTLog/test_1 (0.00s)
    --- PASS: TestTLog/test_3 (0.00s)
    --- PASS: TestTLog/test_2 (0.00s)
    --- PASS: TestTLog/test_4 (0.00s)
PASS
ok      test2   0.257s
*/
