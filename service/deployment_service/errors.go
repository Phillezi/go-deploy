package deployment_service

import (
	"fmt"
)

var (
	DeploymentNotFoundErr  = fmt.Errorf("deployment not found")
	NonUniqueFieldErr      = fmt.Errorf("non unique field")
	InvalidTransferCodeErr = fmt.Errorf("invalid transfer code")
)
