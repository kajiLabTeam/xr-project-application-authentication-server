package utils

import (
	"math/rand"
	"time"

	"github.com/oklog/ulid/v2"
)

func GenerateUlid() (string, error) {
	// 疑似乱数生成器の初期化
	entropy := ulid.Monotonic(rand.New(rand.NewSource(time.Now().UnixNano())), 0)

	// ULID生成
	ulid, err := ulid.New(ulid.Timestamp(time.Now()), entropy)
	if err != nil {
		return "", err
	}

	return ulid.String(), nil
}
