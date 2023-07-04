package proxyserver

import (
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/sirupsen/logrus"
	xproxy "golang.org/x/net/proxy"
)

func New(serverPort int, proxyUrl *url.URL) *http.Server {
	transfer := func(destination io.WriteCloser, source io.ReadCloser) (err error) {
		defer func() { err = destination.Close() }()
		defer func() { err = source.Close() }()
		_, err = io.Copy(destination, source)
		return
	}
	return &http.Server{
		Addr: fmt.Sprintf(":%d", serverPort),
		// ReadTimeout: 10 * time.Second, WriteTimeout: 10 * time.Second, // - Don't use this settings
		MaxHeaderBytes: 1 << 20,
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			log := logrus.WithField("Source", "StartProxyServer > Handler")
			if r.Method == http.MethodConnect {
				var proxyDialer xproxy.Dialer = xproxy.Direct
				if proxyUrl != nil {
					var err error
					proxyDialer, err = xproxy.FromURL(proxyUrl, xproxy.Direct)
					if err != nil {
						log.WithError(err).Error("xproxy.FromURL(proxyUrl, xproxy.Direct)")
						http.Error(w, err.Error(), http.StatusServiceUnavailable)
						return
					}
				}
				dest_conn, err := proxyDialer.Dial("tcp", r.Host)
				if err != nil {
					log.WithError(err).Error("proxyNow.Dialer.Dial")
					http.Error(w, err.Error(), http.StatusServiceUnavailable)
					return
				}
				w.WriteHeader(http.StatusOK)
				hijacker, ok := w.(http.Hijacker)
				if !ok {
					log.Error("wrong type assertion w.(http.Hijacker)")
					http.Error(w, "Hijacking not supported", http.StatusInternalServerError)
					return
				}
				client_conn, _, err := hijacker.Hijack()
				if err != nil {
					log.WithError(err).Error("hijacker.Hijack()")
					http.Error(w, err.Error(), http.StatusServiceUnavailable)
				}
				go transfer(dest_conn, client_conn)
				go transfer(client_conn, dest_conn)
			} else { // r.Method == http.MethodConnect
				t, ok := http.DefaultTransport.(*http.Transport)
				if !ok {
					log.Error("wrong type assertion http.DefaultTransport.(*http.Transport)")
					http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
					return
				}
				if proxyUrl != nil {
					t.Proxy = http.ProxyURL(proxyUrl)
				}
				resp, err := t.RoundTrip(r)
				if err != nil {
					log.WithError(err).Error("t.RoundTrip")
					http.Error(w, err.Error(), http.StatusServiceUnavailable)
					return
				}
				defer resp.Body.Close()
				for key, val := range resp.Header {
					for _, v := range val {
						w.Header().Add(key, v)
					}
				}
				w.WriteHeader(resp.StatusCode)
				if _, err := io.Copy(w, resp.Body); err != nil {
					log.WithError(err).Error("io.Copy")
					http.Error(w, err.Error(), http.StatusServiceUnavailable)
					return
				}

			} // if r.Method == http.MethodConnect
		}), // http.HandlerFunc
	}
} // StartProxy
