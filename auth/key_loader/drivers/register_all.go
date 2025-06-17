package key_loader_drivers

import local_pem_loader "github.com/qcrg/silver_broccoli/auth/key_loader/drivers/local_pem"

func RegisterAll() {
	local_pem_loader.RegisterLocalPEM()
}
