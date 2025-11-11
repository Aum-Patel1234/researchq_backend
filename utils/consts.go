package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const CookieName string = "AuthCookie"
const TokenExpiryTime time.Duration = time.Hour

var JwtAlgo *jwt.SigningMethodHMAC = jwt.SigningMethodHS256
