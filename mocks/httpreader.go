package mocks

import "fmt"

type MockErrorReader struct{}

func (m *MockErrorReader) Read(p []byte) (n int, err error) {
	return 0, fmt.Errorf("failed to read response body") // Gera um erro sempre que tentar ler
}
