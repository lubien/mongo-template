package flycheck

import (
	"context"
	"encoding/json"
	"os/exec"

	"github.com/superfly/fly-checks/check"
)

type IsPrimaryCommandResult struct {
	IsPrimary   *bool  `json:"ismaster"`
	IsSecondary *bool  `json:"secondary"`
	Info        string `json:"info"`
}

// MachineRole outputs current role
func MachineRole(ctx context.Context, checks *check.CheckSuite) (*check.CheckSuite, error) {
	checks.AddCheck("role", func() (string, error) {
		cmd := exec.Command("mongosh", "--eval", "\"db.runCommand('ismaster');\"", "--quiet", "--json")

		output, cmdErr := cmd.Output()
		if cmdErr != nil {
			return "", cmdErr
		}

		var data IsPrimaryCommandResult
		err := json.Unmarshal([]byte(output), &data)
		if err != nil {
			return "", err
		}

		if *data.IsPrimary == true {
			return "primary", nil
		} else if *data.IsSecondary == true {
			return "replica", nil
		}

		if data.Info == "Does not have a valid replica set config" {
			return "non-initialized", nil
		}

		if data.Info != "" {
			return data.Info, nil
		}

		return "error", nil
	})

	return checks, nil
}
