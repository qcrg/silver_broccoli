package initiator

import (
	"sync"

	"github.com/joho/godotenv"
)

func init_all() {
	godotenv.Load()
	InitLogger()
}

var DefaultInitAll = sync.OnceFunc(init_all)
