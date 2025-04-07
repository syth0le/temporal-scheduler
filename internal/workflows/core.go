package workflows

import (
	"fmt"
	"time"

	"temporal-docs/internal/activities"
	"temporal-docs/internal/utils"

	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/workflow"
)

const (
	IPAddressQueueName            = "ip-address"
	IPAddressAndLocationQueueName = "ip-address-location"
)

// GetAddressFromIP is the Temporal Workflow that retrieves the IP address and location info.
func GetAddressFromIP(ctx workflow.Context) error {
	// Define the activity options, including the retry policy
	ao := workflow.ActivityOptions{
		ActivityID:          utils.GenerateAUID(),
		StartToCloseTimeout: 10 * time.Second,
		RetryPolicy: &temporal.RetryPolicy{
			InitialInterval:    time.Second, //amount of time that must elapse before the first retry occurs
			MaximumInterval:    time.Minute, //maximum interval between retries
			BackoffCoefficient: 2,           //how much the retry interval increases
			MaximumAttempts:    2,           // Uncomment this if you want to limit attempts
		},
	}
	ctx = workflow.WithActivityOptions(ctx, ao)

	var ipActivities *activities.IPActivityManager

	var ip string
	err := workflow.ExecuteActivity(ctx, ipActivities.GetIP).Get(ctx, &ip)
	if err != nil {
		return fmt.Errorf("failed to get IP: %s", err)
	}

	var location string
	err = workflow.ExecuteActivity(ctx, ipActivities.GetLocationInfo, ip).Get(ctx, &location)
	if err != nil {
		fmt.Println(fmt.Errorf("failed to get location: %s", err))
		return nil
	}

	fmt.Println(fmt.Sprintf("Your IP is %s and your location is %v", ip, location))

	return nil
}

// GetOnlyIP is the Temporal Workflow that retrieves only the IP address
func GetOnlyIP(ctx workflow.Context) error {
	// Define the activity options, including the retry policy
	ao := workflow.ActivityOptions{
		ActivityID:          utils.GenerateAUID(),
		StartToCloseTimeout: 10 * time.Second,
		RetryPolicy: &temporal.RetryPolicy{
			InitialInterval:    time.Second, //amount of time that must elapse before the first retry occurs
			MaximumInterval:    time.Minute, //maximum interval between retries
			BackoffCoefficient: 2,           //how much the retry interval increases
			MaximumAttempts:    2,           // Uncomment this if you want to limit attempts
		},
	}
	ctx = workflow.WithActivityOptions(ctx, ao)

	var ipActivities *activities.IPActivityManager

	var ip string
	err := workflow.ExecuteActivity(ctx, ipActivities.GetIP).Get(ctx, &ip)
	if err != nil {
		return fmt.Errorf("failed to get IP: %s", err)
	}

	fmt.Println(fmt.Sprintf("Your IP is %s", ip))

	return nil
}
