package kongcli

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"github.com/hbagdi/go-kong/kong"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"net/http"
	"net/url"
	"os"
	"path"
	"strings"
)

type KongClientConfig struct {
	Address   string
	Workspace string

	TLSServerName string

	TLSCACert string

	TLSSkipVerify bool
	Debug         bool

	Headers []string
}

// HeaderRoundTripper injects Headers into requests
// made via RT.
type HeaderRoundTripper struct {
	headers []string
	rt      http.RoundTripper
}

// RoundTrip satisfies the RoundTripper interface.
func (t *HeaderRoundTripper) RoundTrip(req *http.Request) (*http.Response,
	error) {
	newRequest := new(http.Request)
	*newRequest = *req
	newRequest.Header = make(http.Header, len(req.Header))
	for k, s := range req.Header {
		newRequest.Header[k] = append([]string(nil), s...)
	}
	for _, s := range t.headers {
		split := strings.SplitN(s, ":", 2)
		if len(split) >= 2 {
			newRequest.Header[split[0]] = append([]string(nil), split[1])
		}
	}
	return t.rt.RoundTrip(newRequest)
}

// GetKongClient returns a Kong client
func getKongClient(opt KongClientConfig) (*kong.Client, error) {
	c := &http.Client{}
	url, err := url.Parse(opt.Address)
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse kong address")
	}
	if opt.Workspace != "" {
		url.Path = path.Join(url.Path, opt.Workspace)
	}
	if url.Scheme == "http" {
		kongClient, err := kong.NewClient(kong.String(url.String()), c)
		if err != nil {
			return nil, errors.Wrap(err, "creating client for Kong's Admin API")
		}
		if opt.Debug {
			kongClient.SetDebugMode(true)
			kongClient.SetLogger(os.Stderr)
		}
		return kongClient, nil
	}

	var tlsConfig tls.Config
	if opt.TLSSkipVerify {
		tlsConfig.InsecureSkipVerify = true
	}
	if opt.TLSServerName != "" {
		tlsConfig.ServerName = opt.TLSServerName
	}

	if opt.TLSCACert != "" {
		certPool := x509.NewCertPool()
		ok := certPool.AppendCertsFromPEM([]byte(opt.TLSCACert))
		if !ok {
			return nil, errors.New("failed to load TLSCACert")
		}
		tlsConfig.RootCAs = certPool
	}

	defaultTransport := http.DefaultTransport.(*http.Transport)
	defaultTransport.TLSClientConfig = &tlsConfig
	c.Transport = defaultTransport
	if len(opt.Headers) > 0 {
		c.Transport = &HeaderRoundTripper{
			headers: opt.Headers,
			rt:      defaultTransport,
		}
	}
	kongClient, err := kong.NewClient(kong.String(url.String()), c)
	if err != nil {
		return nil, errors.Wrap(err, "creating client for Kong's Admin API")
	}
	return kongClient, nil
}

func newKongClient() (*kong.Client, error) {
	protocol := "http"
	ip := "myecs.jzd"
	port := 65101
	address := fmt.Sprintf("%v://%v:%d", protocol, ip, port)
	logrus.Infof("starting to connect kong on %v.", address)
	opt := KongClientConfig{
		Address: address,
	}
	return getKongClient(opt)
}
