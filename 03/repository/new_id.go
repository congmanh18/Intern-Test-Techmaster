package repository

import (
	"time"

	"math/rand"

	"github.com/oklog/ulid/v2"
)

// Tạo ULID mới
func newULID() string {
	t := time.Now()
	entropy := ulid.Monotonic(rand.New(rand.NewSource(t.UnixNano())), 0)
	return ulid.MustNew(ulid.Timestamp(t), entropy).String()
}
