package test

type MockRepository struct {
	IncrValue int64
}

func NewMockRepository() *MockRepository {
	return &MockRepository{}
}

func (m *MockRepository) Save(str string, value any, exp int) error {
	return nil
}

func (m *MockRepository) Get(key string) any {
	return m.IncrValue
}

func (m *MockRepository) Incr(key string) int64 {
	m.IncrValue++
	return m.IncrValue
}

func (m *MockRepository) Expr(key string, timeout int) bool {
	return true
}
