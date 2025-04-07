package utils

import (
	"strings"

	"github.com/google/uuid"
)

const (
	serviceNamePrefix    = "str"
	workflowEntityPrefix = "w"
	scheduleEntityPrefix = "s"
	activityEntityPrefix = "a"
)

func GenerateWUID() string {
	return generateUID(workflowEntityPrefix)
}

func GenerateSUID() string {
	return generateUID(scheduleEntityPrefix)
}

func GenerateAUID() string {
	return generateUID(activityEntityPrefix)
}

func generateUID(entityPrefix string) string {
	return serviceNamePrefix + entityPrefix + strings.Replace(uuid.New().String(), "-", "", -1)
}
