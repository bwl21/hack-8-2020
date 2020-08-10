package cmd

import (
	"fmt"

	"github.com/bwl21/zupfmanager/pkg/api"
	"github.com/spf13/cobra"
)

// generateSchemaCmd represents the generateSchema command
var generateSchemaCmd = &cobra.Command{
	Use:   "generate-schema",
	Short: "Produces the GraphQL schema as JSON",
	RunE: func(cmd *cobra.Command, args []string) error {
		s, err := api.NewServer()
		if err != nil {
			return err
		}
		ss, err := s.SchemaJSON()
		if err != nil {
			return err
		}

		fmt.Println(string(ss))
		return nil
	},
}

func init() {
	rootCmd.AddCommand(generateSchemaCmd)
}
