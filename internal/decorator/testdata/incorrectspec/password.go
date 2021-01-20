package password

type manager struct{}

func (m *manager) ComparePassword(password, hash string) (bool, error) {
	return false, nil
}
