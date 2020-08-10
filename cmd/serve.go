//go:generate sh -c "cd ../pkg/webui && yarn build && cd -; rice embed-go"

package cmd

import (
	"encoding/json"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"

	rice "github.com/GeertJohan/go.rice"
	"github.com/bwl21/zupfmanager/pkg/api"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve <config.json>",
	Short: "Starts the zupfmanager server",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg, err := getConfig(args[0])
		if err != nil {
			return err
		}

		apisrv, err := api.NewServer()
		if err != nil {
			return err
		}

		debugProxyAddr, _ := cmd.Flags().GetString("debug-proxy-addr")

		mux := http.NewServeMux()
		apisrv.Serve(mux)
		if debugProxyAddr == "" {
			// WebUI is a single-page app, hence any path that does not resolve to a static file must result in /index.html.
			// As a (rather crude) fix we intercept the response writer to find out if the FileServer returned an error. If so
			// we return /index.html instead.
			dws := http.FileServer(rice.MustFindBox("../pkg/webui/build").HTTPBox())
			mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
				dws.ServeHTTP(&interceptResponseWriter{
					ResponseWriter: w,
					errH: func(rw http.ResponseWriter, code int) {
						r.URL.Path = "/"
						rw.Header().Set("Content-Type", "text/html; charset=utf-8")
						dws.ServeHTTP(rw, r)
					},
				}, r)
			})
		} else {
			proxyURL, err := url.Parse(debugProxyAddr)
			if err != nil {
				return err
			}
			mux.Handle("/", httputil.NewSingleHostReverseProxy(proxyURL))
		}

		log.WithField("address", cfg.Addr).Info("serving zupfmanager")
		err = http.ListenAndServe(cfg.Addr, mux)
		if err != nil {
			return err
		}

		return nil
	},
}

type interceptResponseWriter struct {
	http.ResponseWriter
	errH func(http.ResponseWriter, int)
}

func (w *interceptResponseWriter) WriteHeader(status int) {
	if status >= http.StatusBadRequest {
		w.errH(w.ResponseWriter, status)
		w.errH = nil
	} else {
		w.ResponseWriter.WriteHeader(status)
	}
}

func (w *interceptResponseWriter) Write(p []byte) (n int, err error) {
	if w.errH == nil {
		return len(p), nil
	}
	return w.ResponseWriter.Write(p)
}

type serverConfig struct {
	Addr string `json:"address"`
}

func getConfig(fn string) (*serverConfig, error) {
	f, err := os.Open(fn)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var cfg serverConfig
	err = json.NewDecoder(f).Decode(&cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}

func init() {
	rootCmd.AddCommand(serveCmd)

	serveCmd.Flags().String("debug-proxy-addr", "", "bypass the webui server with this proxy")
}
