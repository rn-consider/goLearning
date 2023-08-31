package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// 生成加密密码
func generatePassword(password string) (string, error) {
	const cost = 10 //哈希成本
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	return string(bytes), err
}

// 密码效验
func checkPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func main() {
	password := "123456"
	hash, _ := generatePassword(password) // 不处理错误

	fmt.Println("Password:", password)
	fmt.Println("Hash:    ", hash)

	match := checkPassword(password, hash)
	fmt.Println("Match:   ", match)
}
