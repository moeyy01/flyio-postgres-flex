package flycheck

import (
	"context"

	"github.com/fly-apps/postgres-flex/pkg/flypg"
	"github.com/pkg/errors"
	"github.com/superfly/fly-checks/check"
)

// PostgreSQLRole outputs current role
func PostgreSQLRole(ctx context.Context, checks *check.CheckSuite) (*check.CheckSuite, error) {
	node, err := flypg.NewNode()
	if err != nil {
		return checks, errors.Wrap(err, "failed to initialize node")
	}

	conn, err := node.RepMgr.NewLocalConnection(ctx)
	if err != nil {
		return checks, errors.Wrap(err, "failed to connect to local node")
	}

	// Cleanup connections
	checks.OnCompletion = func() {
		conn.Close(ctx)
	}

	checks.AddCheck("role", func() (string, error) {
		role, err := node.RepMgr.CurrentRole(ctx, conn)
		if err != nil {
			return "failed", errors.Wrap(err, "failed to check role")
		}

		if role == flypg.PrimaryRoleName {
			return "primary", nil
		} else if role == flypg.StandbyRoleName {
			return "replica", nil
		} else {
			return "unknown", nil
		}
	})
	return checks, nil
}