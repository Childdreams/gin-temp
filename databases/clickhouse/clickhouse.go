package clickhouse

import (
	"app/settings"
	"context"
	"github.com/ClickHouse/clickhouse-go/v2"
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
	"log"
	"os"
)

var ClickHouse driver.Conn

var (
	ClickHouseUri string
)

func init() {
	settings.RequireEnvs([]string{
		"CLICKHOUSE_URI",
	})

	ClickHouseUri = os.Getenv("CLICKHOUSE_URI")
	options, err := clickhouse.ParseDSN(ClickHouseUri)
	options.MaxIdleConns = 10
	options.MaxOpenConns = 20
	ClickHouse, err = clickhouse.Open(options)
	if err != nil {
		log.Fatal(err)
	}

	if err := ClickHouse.Ping(context.Background()); err != nil {
		if exception, ok := err.(*clickhouse.Exception); ok {
			log.Fatalf("[%d] %s \n%s\n", exception.Code, exception.Message, exception.StackTrace)
		} else {
			log.Fatal(err)
		}
		return
	}
}
