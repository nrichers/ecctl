// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package cmdelasticsearchplan

import (
	"path/filepath"

	"github.com/spf13/cobra"

	"github.com/elastic/ecctl/pkg/deployment/elasticsearch/plan"
	"github.com/elastic/ecctl/pkg/ecctl"
	"github.com/elastic/ecctl/pkg/util"
)

var listElasticsearchPlansCmd = &cobra.Command{
	Use:     "history <cluster id>",
	Short:   "Lists the plan history",
	Aliases: []string{"attempts"},
	PreRunE: cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		p, err := plan.GetHistory(plan.GetHistoryParams{ClusterParams: util.ClusterParams{
			API:       ecctl.Get().API,
			ClusterID: args[0],
		}})
		if err != nil {
			return err
		}

		return ecctl.Get().Formatter.Format(
			filepath.Join("elasticsearch", "planhistory"),
			p,
		)
	},
}
