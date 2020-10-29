package env

import (
	"os"
)

// JwtSecret segreto per il jwt
var JwtSecret = []byte(os.Getenv("JWT_SECRET"))