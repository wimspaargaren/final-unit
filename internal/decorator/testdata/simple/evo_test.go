package password

func X() string {
	return "expected string"
}

func defaultPass() string {
	return "SuperSecret123"
}

func defaultHash() string {
	return "$argon2id$v=19$m=65536,t=1,p=4$uXcC6G/Usn8zRo5hkjA15w$BCq8OA"
}

func defaultManager() *manager {
	return NewDefaultManager().(*manager)
}
