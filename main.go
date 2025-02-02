package main

import (
	hbaseclient "github.com/balampbv/hbase-migrate/hbase"
	"time"
)

func main() {
	client, err := hbaseclient.NewHBaseClient(hbaseclient.HBaseConfig{
		URI: "http://localhost:9090",
	})
	if err != nil {
		panic(err)
	}
	err = client.CreateTable(hbaseclient.CreateTableRequest{
		NamespaceName: "default",
		TableName:     "test_table",
		ColumnFamilies: []hbaseclient.ColumnFamily{
			{
				Name:        "cf1",
				TTL:         48 * time.Hour,
				MaxVersions: 5,
			},
		},
	})
	if err != nil {
		panic(err)
	}
}
