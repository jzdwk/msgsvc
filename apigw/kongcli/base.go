/*
@Time : 20-7-28
@Author : jzd
@Project: apigw
*/
package kongcli

import "github.com/kong/go-kong/kong"

//kong support http method
type HttpType int

const (
	HttpGet     HttpType = 0
	HttpPost    HttpType = 1
	HttpDelete  HttpType = 2
	HttpPut     HttpType = 3
	HttpPatch   HttpType = 4
	HttpHead    HttpType = 5
	HttpOptions HttpType = 6
	HttpAny     HttpType = 7

	HTTP  = 0
	HTTPS = 1
	GRPC  = 2

	EmptyKoUUID = "-"
)

const (
	plugin = "plugin"
)

var HttpMethodMap = map[HttpType]string{
	HttpGet:     "GET",
	HttpPost:    "POST",
	HttpDelete:  "DELETE",
	HttpPut:     "PUT",
	HttpPatch:   "PATCH",
	HttpHead:    "HEAD",
	HttpOptions: "OPTIONS",
	HttpAny:     "-",
}

var HttpProtocolMap = map[int]string{
	HTTP:  "http",
	HTTPS: "https",
	GRPC:  "grpc",
}

func (t *HttpType) ToStrings() []*string {
	if *t == HttpAny {
		get, post, deletes, put, patch, head, options := "GET", "POST", "DELETE", "PUT", "PATCH", "HEAD", "OPTIONS"
		return []*string{&get, &post, &deletes, &put, &patch, &head, &options}
	}
	method := HttpMethodMap[*t]
	return []*string{&method}
}

type KongClientWrap struct {
	cli *kong.Client
}

func NewKongClient() (*KongClientWrap, error) {
	if cli, err := newKongClient(); err == nil {
		return &KongClientWrap{cli: cli}, nil
	} else {
		return nil, err
	}
}
