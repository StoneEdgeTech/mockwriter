package mockwriter

import ()

type MockWriter struct {
	Written []byte
}

func (m *MockWriter) Write(p []byte) (n int, err error) {
	m.Written = p
	return len(m.Written), nil
}
