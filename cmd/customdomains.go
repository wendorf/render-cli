package cmd

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/renderinc/cli/pkg/http"
	"github.com/renderinc/cli/pkg/table"
	"github.com/spf13/cobra"
)

type ShortServer struct {
	Id string `json:"id"`
	Name string `json:"name"`
}

type CustomDomain struct {
	Id string `json:"id"`
	Name string `json:"name"`
	IsApex bool `json:"isApex"`
	PublicSuffix string `json:"publicSuffix"`
	RedirectForName string `json:"redirectForName"`
	Verified bool `json:"verified"`
	CreatedAt string `json:"createdAt"`
	Server ShortServer `json:"server""`
}

type CursorCustomDomain struct {
	Cursor string `json:"cursor"`
	CustomDomain CustomDomain `json:"customDomain"`
}

// customDomainsCmd represents the customDomains command
var customDomainsCmd = &cobra.Command{
	Use:   "custom-domains",
	Short: "Get custom domains by service id",
	Long: `It's hack day!'`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return errors.New("provide only one argument")
		}

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		serviceId := args[0]
		if err := table.Print([]string{"Id", "Name", "CreatedAt"}, customDomains(serviceId)); err != nil {
			panic(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(customDomainsCmd)
}

func customDomains(serviceId string) []CustomDomain {
	jsonString, err := http.Request(fmt.Sprintf("services/%s/custom-domains", serviceId))
	if err != nil {
		panic(err)
	}

	var cursorCustomDomains []CursorCustomDomain
	if err := json.Unmarshal(jsonString, &cursorCustomDomains); err != nil {
		panic(err)
	}

	var customDomains []CustomDomain
	for _, cursorCustomDomain := range cursorCustomDomains {
		customDomains = append(customDomains, cursorCustomDomain.CustomDomain)
	}
	return customDomains
}