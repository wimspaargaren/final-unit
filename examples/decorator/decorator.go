package decorator

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"fmt"
	"strings"

	"golang.org/x/crypto/argon2"
)

// PasswordConfig config for passwords
type PasswordConfig struct {
	time    uint32
	memory  uint32
	threads uint8
	keyLen  uint32
}

// PWManager the password manager containing the config
type PWManager struct {
	Config *PasswordConfig
}

// DefaultManager default manager
func DefaultManager() *PWManager {
	return &PWManager{
		Config: &PasswordConfig{
			time:    1,
			memory:  64 * 1024,
			threads: 4,
			keyLen:  32,
		},
	}
}

// GeneratePassword generate password hash from string
func (m *PWManager) GeneratePassword(password string) (string, error) {
	salt := make([]byte, 16)
	if _, err := rand.Read(salt); err != nil {
		return "", err
	}

	hash := argon2.IDKey([]byte(password), salt, m.Config.time, m.Config.memory, m.Config.threads, m.Config.keyLen)

	b64Salt := base64.RawStdEncoding.EncodeToString(salt)
	b64Hash := base64.RawStdEncoding.EncodeToString(hash)

	format := "$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s"
	full := fmt.Sprintf(format, argon2.Version, m.Config.memory, m.Config.time, m.Config.threads, b64Salt, b64Hash)
	return full, nil
}

// ComparePassword compares password with given hash
func (m *PWManager) ComparePassword(password, hash string) (bool, error) {
	parts := strings.Split(hash, "$")

	c := &PasswordConfig{}
	_, err := fmt.Sscanf(parts[3], "m=%d,t=%d,p=%d", &c.memory, &c.time, &c.threads)
	if err != nil {
		return false, err
	}

	salt, err := base64.RawStdEncoding.DecodeString(parts[4])
	if err != nil {
		return false, err
	}

	decodedHash, err := base64.RawStdEncoding.DecodeString(parts[5])
	if err != nil {
		return false, err
	}
	m.Config.keyLen = uint32(len(decodedHash))

	comparisonHash := argon2.IDKey([]byte(password), salt, m.Config.time, m.Config.memory, m.Config.threads, m.Config.keyLen)

	return (subtle.ConstantTimeCompare(decodedHash, comparisonHash) == 1), nil
}
