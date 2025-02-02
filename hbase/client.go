package hbaseclient

import (
	"context"
	"log"
	"time"

	"github.com/apache/thrift/lib/go/thrift"
	"github.com/balampbv/hbase-migrate/hbase/gen-go/hbase"
)

const (
	accessKeyID    = "ACCESSKEYID"
	accesSignature = "ACCESSSIGNATURE"
)

type HBaseClient struct {
	client *hbase.THBaseServiceClient
}

type HBaseConfig struct {
	URI      string `mapstructure:"uri"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
}

type ColumnFamily struct {
	Name        string        `json:"name"`
	TTL         time.Duration `json:"ttl,omitempty"`         // Time to Live for the column family
	MaxVersions int32         `json:"maxVersions,omitempty"` // Max versions of cells to store
}

type CreateTableRequest struct {
	NamespaceName  string         `json:"namespaceName"`
	TableName      string         `json:"tableName"`
	ColumnFamilies []ColumnFamily `json:"columnFamilies"`
}

// NewHBaseClient creates a new HBase client using the given configuration
func NewHBaseClient(conf HBaseConfig) (*HBaseClient, error) {
	protocolFactory := thrift.NewTBinaryProtocolFactoryConf(nil)
	transport, err := thrift.NewTHttpClient(conf.URI)
	if err != nil {
		return nil, err
	}

	httClient, ok := transport.(*thrift.THttpClient)
	if !ok {
		return nil, err
	}

	// Set the access key ID and access signature
	httClient.SetHeader(accessKeyID, conf.User)
	httClient.SetHeader(accesSignature, conf.Password)

	client := hbase.NewTHBaseServiceClientFactory(transport, protocolFactory)
	if err := transport.Open(); err != nil {
		return nil, err
	}
	return &HBaseClient{client: client}, nil
}

// CreateTable creates a new table in HBase with the given column families
func (h *HBaseClient) CreateTable(request CreateTableRequest) error {
	// Prepare the column families
	var cfDescriptors []*hbase.TColumnFamilyDescriptor
	for _, cf := range request.ColumnFamilies {
		cfDescriptors = append(cfDescriptors, &hbase.TColumnFamilyDescriptor{
			Name:        []byte(cf.Name),
			TimeToLive:  thrift.Int32Ptr(int32(cf.TTL.Seconds())),
			MaxVersions: &cf.MaxVersions,
		})
	}

	// Prepare the table descriptor
	tableDesc := &hbase.TTableDescriptor{
		TableName: &hbase.TTableName{
			Ns:        []byte(request.NamespaceName),
			Qualifier: []byte(request.TableName),
		},
		Columns: cfDescriptors,
	}

	// Create the table using the Thrift client
	err := h.client.CreateTable(context.Background(), tableDesc, nil)
	if err != nil {
		return err
	}

	log.Printf("Table '%s' created successfully", request.TableName)
	return nil
}
