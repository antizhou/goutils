package es

import (
	"context"

	"github.com/olivere/elastic"
	"github.com/pkg/errors"
)

type EsHelper struct {
	username string
	password string
	hosts    []string
	sniff    bool
}

func Init(hosts []string) *EsHelper {
	return &EsHelper{
		hosts: hosts,
	}
}

func (o *EsHelper) SetUsername(username string) {
	if username == "" {
		return
	}
	o.username = username
}

func (o *EsHelper) SetPassword(password string) {
	if password == "" {
		return
	}
	o.password = password
}

func (o *EsHelper) SetNiff() {
	o.sniff = true
}

func (o *EsHelper) GetClient() (*elastic.Client, error) {
	if len(o.hosts) == 0 {
		return nil, errors.New("es hosts is empty")
	}

	options := []elastic.ClientOptionFunc{
		elastic.SetURL(o.hosts...),
	}

	if o.sniff {
		options = append(options,
			elastic.SetSniff(true))
	}

	if o.username != "" && o.password != "" {
		options = append(options,
			elastic.SetBasicAuth(o.username, o.password))
	}

	client, err := elastic.NewClient(options...)
	if err != nil {
		return nil, err
	}
	return client, nil
}

func (o *EsHelper) IndexRequest(esindex string, estype string, data interface{}) *elastic.BulkIndexRequest {
	return elastic.NewBulkIndexRequest().Index(esindex).Type(estype).Doc(data)
}

func (o *EsHelper) DeleteRequest(esindex string, estype string, id string) *elastic.BulkDeleteRequest {
	return elastic.NewBulkDeleteRequest().Index(esindex).Type(estype).Id(id)
}

func (o *EsHelper) Bulk(client *elastic.Client, requests ...elastic.BulkableRequest) (*elastic.BulkResponse, error) {
	service := client.Bulk().Add(requests...)
	return service.Do(context.TODO())
}
