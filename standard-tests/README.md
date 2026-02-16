To make go recognize tests, the following 3 must be satisfied:
1. **Filename**: your test file must end in `_test.go`
2. **Function Name**: Test functions must start with `Test` and followed by a capital letter.
E.g. `Testadd` and `testAdd` are ignored, but `TestAdd` works.
3. **Signature**: Every test function takes exactly one argument `(t *testing.T)`

