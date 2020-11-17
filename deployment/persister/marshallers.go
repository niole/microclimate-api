package persister

import (
	api "api/protobuf"
	sql "database/sql"
	"log"
)

// TODO use generated enum map things
var statusMap = map[string]api.Deployment_Status{
	"unreachable":   api.Deployment_UNREACHABLE,
	"starting_up":   api.Deployment_STARTING_UP,
	"failed":        api.Deployment_FAILED,
	"running":       api.Deployment_RUNNING,
	"shutting_down": api.Deployment_SHUTTING_DOWN,
}

func scanOneDeployment(scanner func(dest ...interface{}) error, saveToDeployment *api.Deployment) error {
	var id string
	var ownerUserId string
	var status string
	var name string
	err := scanner(&id, &ownerUserId, &status, &name)

	*saveToDeployment = api.Deployment{
		Id:          id,
		OwnerUserId: ownerUserId,
		Status:      statusMap[status],
		Name:        name,
	}

	return err
}

func ScanOneDeployment(result *sql.Row, saveToDeployment *api.Deployment) error {
	return scanOneDeployment(result.Scan, saveToDeployment)
}

func ScanManyDeployments(result *sql.Rows, err error) ([]api.Deployment, error) {
	defer result.Close()

	deployments := make([]api.Deployment, 0)
	if err == nil {
		for result.Next() {
			var deployment api.Deployment
			err = scanOneDeployment(result.Scan, &deployment)
			if err == nil {
				deployments = append(deployments, deployment)
			}
		}
	} else {
		log.Printf("Something went wrong while unmarshalling deployments, error: %v", err)
	}

	return deployments, err
}
