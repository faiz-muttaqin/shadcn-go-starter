package util

import (
	"context"
	"os"
)

var (
	JwtSecretKey = []byte(os.Getenv("JWT_SECRET_KEY"))
	Context      = context.Background()
)
