package commands

import "github.com/SherClockHolmes/webpush-go"

type SetClientSubscriptionCommandDetails struct {
	ClientSubscription *webpush.Subscription `mapstructure:",squash"`
}

type SetClientSubscriptionCommand struct {
	WorkerCommand
	SetClientSubscriptionCommandDetails
}
