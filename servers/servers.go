package servers

import (
	"net/http"
)

type Server interface {
	Run(http.Handler)
	Stop()
}
