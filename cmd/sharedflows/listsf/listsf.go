package listsf

import (
	"net/url"
	"../../shared"
	"path"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all shared flows in the organization.",
	Long:  "Lists all shared flows in the organization.",
	Run: func(cmd *cobra.Command, args []string) {
		u, _ := url.Parse(shared.BaseURL)
		u.Path = path.Join(u.Path, shared.RootArgs.Org, "sharedflows")
		shared.GetHttpClient(u.String(), shared.RootArgs.Token)
	},
}

func init() {

}