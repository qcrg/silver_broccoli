package auth_drivers

import jwt_auth "github.com/qcrg/silver_broccoli/auth/drivers/jwt"

func RegisterAll() {
	jwt_auth.RegisterJWT()
}
