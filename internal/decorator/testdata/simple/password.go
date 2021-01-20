package password

type manager struct{}

func (m *manager) GenerateHash(password string) (string, error) {
	return "", nil
}

func (m *manager) ComparePassword(password, hash string) (bool, error) {
	return true, nil
}
