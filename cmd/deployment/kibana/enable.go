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

package cmdkibana

import (
	"path/filepath"

	"github.com/spf13/cobra"

	cmdutil "github.com/elastic/ecctl/cmd/util"
	"github.com/elastic/ecctl/pkg/deployment/kibana"
	"github.com/elastic/ecctl/pkg/ecctl"
	"github.com/elastic/ecctl/pkg/util"
)

var enableKibanaCmd = &cobra.Command{
	Use:     "enable <deployment id>",
	Short:   "Enables a kibana instance in the selected deployment",
	PreRunE: cmdutil.MinimumNArgsAndUUID(1),

	RunE: func(cmd *cobra.Command, args []string) error {
		var track, _ = cmd.Flags().GetBool("track")
		res, err := kibana.Enable(kibana.DeploymentParams{
			API: ecctl.Get().API,
			ID:  args[0],
			TrackParams: util.TrackParams{
				Track:  track,
				Output: ecctl.Get().Config.OutputDevice,
			},
		})
		if err != nil {
			return err
		}
		return ecctl.Get().Formatter.Format(filepath.Join("kibana", "enable"), res)
	},
}

func init() {
	enableKibanaCmd.Flags().Bool("track", false, cmdutil.TrackFlagMessage)
}
