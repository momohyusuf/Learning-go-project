package utils

import (
	"crypto/rand"
	"math/big"
	"os"
	"strings"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func GenerateRandomString(length int) string {
	var sb strings.Builder
	for i := 0; i < length; i++ {
		index, _ := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		sb.WriteByte(charset[index.Int64()])
	}
	return sb.String()
}

func SaveTaskToFile(jsonData []byte, fileName string) string {
	os.WriteFile(fileName, []byte(string(jsonData)), 0666)
	return "Task has been created successfully"
}
