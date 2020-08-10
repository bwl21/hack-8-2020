package cmd

import (
	"encoding/json"
	"net/http"
	"os"

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

		mux := http.NewServeMux()
		apisrv.Serve(mux)

		log.WithField("address", cfg.Addr).Info("serving zupfmanager")
		err = http.ListenAndServe(cfg.Addr, mux)
		if err != nil {
			return err
		}

		return nil
	},
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
}
