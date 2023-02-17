package main

import (
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
	"gov-cms/app/internal/config"
	"gov-cms/app/internal/handler"
	"gov-cms/app/internal/svc"
	"gov-cms/common/result"
	"net/http"
)

var configFile = flag.String("f", "etc/govcms.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf, rest.WithUnauthorizedCallback(result.JwtUnauthorizedResult), rest.WithCustomCors(nil, notAllowedFn))
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)
	
	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}

func notAllowedFn(w http.ResponseWriter) {
	//w.Header().Add("Access-Control-Allow-Headers", "Authorization")
}
