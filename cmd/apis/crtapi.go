// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package apis

import (
	"github.com/spf13/cobra"
	"github.com/srinandan/apigeecli/apiclient"
	"github.com/srinandan/apigeecli/client/apis"
	proxybundle "github.com/srinandan/apigeecli/cmd/apis/proxybundle"
)

//CreateCmd to create api
var CreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Creates an API proxy in an Apigee Org",
	Long:  "Creates an API proxy in an Apigee Org",
	Args: func(cmd *cobra.Command, args []string) (err error) {
		apiclient.SetApigeeOrg(org)
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		if proxy != "" {
			_, err = apis.CreateProxy(name, proxy)
		} else if oasDoc != "" {
			err = GenerateAPIProxyDefFromOAS(name, oasDoc)
			if err != nil {
				return err
			}

			err = proxybundle.GenerateAPIProxyBundle(name)
			if err != nil {
				return err
			}

			if importProxy {
				_, err = apis.CreateProxy(name, name+".zip")
			}

		} else {
			_, err = apis.CreateProxy(name, "")
		}

		return
	},
}

var proxy, oasDoc string
var importProxy bool

func init() {

	CreateCmd.Flags().StringVarP(&name, "name", "n",
		"", "API Proxy name")
	CreateCmd.Flags().StringVarP(&proxy, "proxy", "p",
		"", "API Proxy Bundle path")
	CreateCmd.Flags().StringVarP(&oasDoc, "oas", "f",
		"", "Open API 3.0 Specification file")
	CreateCmd.Flags().BoolVarP(&importProxy, "import", "",
		true, "Import API Proxy after generation from spec")

	_ = CreateCmd.MarkFlagRequired("name")
}
