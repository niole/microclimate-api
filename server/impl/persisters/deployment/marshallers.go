package deployment

import (
	api "api/protobuf"
	sql "database/sql"
)

// TODO use generated enum map things
var statusMap = map[string]api.Deployment_Status{
	"unreachable":   api.Deployment_UNREACHABLE,
	"starting_up":   api.Deployment_STARTING_UP,
	"failed":        api.Deployment_FAILED,
	"running":       api.Deployment_RUNNING,
	"shutting_down": api.Deployment_SHUTTING_DOWN,
}

func ScanOneDeployment(result *sql.Row, saveToDeployment *api.Deployment) error {
	var id string
	var ownerUserId string
	var domain string
	var status string
	err := result.Scan(&id, &ownerUserId, &domain, &status)

	*saveToDeployment = api.Deployment{
		Id:          id,
		OwnerUserId: ownerUserId,
		Domain:      domain,
		Status:      statusMap[status],
	}

	return err
}
