package test

import (
	uuid "github.com/satori/go.uuid"
	"testing"
)

func TestGenerateUUID(t *testing.T) {
	println(uuid.NewV4().String())
}

// 生成唯一码
func GetUUID() string {
	return uuid.NewV4().String()
}
