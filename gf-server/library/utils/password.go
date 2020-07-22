package utils

import "golang.org/x/crypto/bcrypt"

// PasswordCheck 密码检查(工具类)
// p1 数据库密码
// p2 前端传递密码
// false 校验失败
func CompareHashAndPassword(p1 string, p2 string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(p1), []byte(p2)); err != nil {
		return false
	}
	return true
}

// EncryptedPassword: 加密密码(工具类)
func EncryptedPassword(p1 string) (p2 string, err error) {
	if byTes, err := bcrypt.GenerateFromPassword([]byte(p1), bcrypt.DefaultCost); err == nil { // 加密密码
		return string(byTes), nil
	}
	return p2, err
}
