package listapp

import (
	"../../shared"
	"github.com/spf13/cobra"
	"net/url"
	"path"
)

var Cmd = &cobra.Command{
	Use:   "list",
	Short: "Returns a list of Developer Applications",
	Long:  "Returns a list of app IDs within an organization based on app status",
	RunE: func(cmd *cobra.Command, args []string) error {
		u, _ := url.Parse(shared.BaseURL)
		u.Path = path.Join(u.Path, shared.RootArgs.Org, "apps")
		q := u.Query()
		if expand {
			q.Set("expand", "true")
		} else {
			q.Set("expand", "false")
		}
		if count != "" {
			q.Set("row", count)
		}
		u.RawQuery = q.Encode()
		return shared.HttpClient(u.String())
	},
}

var expand = false
var count string

func init() {

	Cmd.Flags().StringVarP(&count, "count", "c",
		"", "Number of apps; limit is 1000")

	Cmd.Flags().BoolVarP(&expand, "expand", "x",
		false, "Expand Details")
}
