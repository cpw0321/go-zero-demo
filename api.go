package main

import (
	"cpw/go-zero-demo/internal/middleware"
	"flag"
	"fmt"

	"cpw/go-zero-demo/internal/config"
	"cpw/go-zero-demo/internal/handler"
	"cpw/go-zero-demo/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf, rest.WithNotAllowedHandler(middleware.NewCorsMiddleware().Handler()))
	defer server.Stop()
	server.Use(middleware.NewCorsMiddleware().Handle)

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
