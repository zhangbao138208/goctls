{{.head}}

package server

import (
	{{if .notStream}}"context"{{end}}
	{{if .hasLock}}
           "fmt"
           	"github.com/go-locks/distlock/mutex"
           	"time"{{end}}


	{{.imports}}

)

type {{.server}}Server struct {
	svcCtx *svc.ServiceContext
	{{.unimplementedServer}}
}

func New{{.server}}Server(svcCtx *svc.ServiceContext) *{{.server}}Server {
	return &{{.server}}Server{
		svcCtx: svcCtx,
	}
}

{{.funcs}}
