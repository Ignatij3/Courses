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
		t.Run(tc.name, func(t *testing.T) {
			// Here you test tc.value against a test function.
			t.Log(tc.value)
		})
	}
}
/*
go test -v
=== RUN   TestTLog
=== PAUSE TestTLog
=== CONT  TestTLog
=== RUN   TestTLog/test_1
    parallel_test.go:20: 1
=== RUN   TestTLog/test_2
    parallel_test.go:20: 2
=== RUN   TestTLog/test_3
    parallel_test.go:20: 3
=== RUN   TestTLog/test_4
    parallel_test.go:20: 4
--- PASS: TestTLog (0.00s)
    --- PASS: TestTLog/test_1 (0.00s)
    --- PASS: TestTLog/test_2 (0.00s)
    --- PASS: TestTLog/test_3 (0.00s)
    --- PASS: TestTLog/test_4 (0.00s)
PASS
ok      test0   0.248s
*/
