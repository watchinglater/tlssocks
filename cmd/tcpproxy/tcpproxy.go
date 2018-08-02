package main

import (
	"flag"

	"github.com/google/tcpproxy"
	"go.uber.org/zap"
)

var logger *zap.Logger

func init() {
	l, _ := zap.NewProduction()
	logger = l
}

func main() {
	defer logger.Sync()
	flagDestination := flag.String("destination", "", "address of destination server like 127.0.0.1:8000")
	flagAddr := flag.String("addr", "", "where to listen like 127.0.0.1:8001")

	flag.Parse()
	if *flagDestination == "" {
		flag.Usage()
		logger.Fatal("empty socks server")
	}
	if *flagAddr == "" {
		flag.Usage()
		logger.Fatal("empty addr - I do not know where to listen")
	}

	logger.Info(
		"starting tcp proxy",
		zap.String("addr", *flagAddr),
		zap.String("destination", *flagDestination),
	)
	var p tcpproxy.Proxy
	p.AddRoute(*flagAddr, tcpproxy.To(*flagDestination))
	logger.Fatal("sgutting down", zap.Error(p.Run()))
}