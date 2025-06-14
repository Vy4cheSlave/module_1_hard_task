package generate_password

import (
	"math/rand"
	"bytes"
	"time"
)

// letters - список допустимых символов в пароле
const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func GeneratePassword(n int) (string, error) {
	// TODO: ваш код
	rand.Seed(time.Now().UnixNano())
	indexLen := len(letters)
	var buf bytes.Buffer
	var randVal int
	for i := 0; i < n; i++ {
		randVal = rand.Intn(indexLen)
		buf.WriteByte(letters[randVal])
	}

	return buf.String(), nil


	panic("implement me")
}
