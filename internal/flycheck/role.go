package flycheck

import (
	"context"
	"encoding/json"
	"os/exec"

	"github.com/superfly/fly-checks/check"
)

type IsPrimaryCommandResult struct {
	IsPrimary *bool `json:"ismaster"`
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

		if data.IsPrimary == nil {
			return "error", nil
		}
		if *data.IsPrimary {
			return "primary", nil
		} else {
			return "replica", nil
		}
	})

	return checks, nil
}
