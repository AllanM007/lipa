package main

// import (
// 	"context"
// 	"fmt"
// 	"time"

// 	"go.temporal.io/sdk/temporal"
// 	"go.temporal.io/sdk/workflow"
// )

// type BankingService interface {
// 	Withdraw(accountNumber string, amount int, referenceID string) (string, error)
// 	Deposit(accountNumber string, amount int, referenceID string) (string, error)
// }

// type Activities struct {
// 	bank BankingService
// }

// type PaymentDetails struct {
// 	ReferenceID   string
// 	SourceAccount string
// 	TargetAccount string
// 	Amount        int
// }

// func MoneyTransfer(ctx workflow.Context, input PaymentDetails) (string, error) {
// 	// RetryPolicy specifies how to automatically handle retries if an Activity fails.
// 	retrypolicy := &temporal.RetryPolicy{
// 		InitialInterval:        time.Second,
// 		BackoffCoefficient:     2.0,
// 		MaximumInterval:        100 * time.Second,
// 		MaximumAttempts:        0, // unlimited retries
// 		NonRetryableErrorTypes: []string{"ErrInvalidAccount", "ErrInsufficientFunds"},
// 	}

// 	options := workflow.ActivityOptions{
// 		// Timeout options specify when to automatically timeout Activity functions.
// 		StartToCloseTimeout: time.Minute,
// 		// Optionally provide a customized RetryPolicy.
// 		// Temporal retries failed Activities by default.
// 		RetryPolicy: retrypolicy,
// 	}

// 	// Apply the options.
// 	ctx = workflow.WithActivityOptions(ctx, options)

// 	// Withdraw money.
// 	var withdrawOutput string

// 	withdrawErr := workflow.ExecuteActivity(ctx, "Withdraw", input).Get(ctx, &withdrawOutput)

// 	if withdrawErr != nil {
// 		return "", withdrawErr
// 	}

// 	// Deposit money.
// 	var depositOutput string

// 	depositErr := workflow.ExecuteActivity(ctx, "Deposit", input).Get(ctx, &depositOutput)

// 	if depositErr != nil {
// 		// The deposit failed; put money back in original account.
// 		var result string
// 		refundErr := workflow.ExecuteActivity(ctx, "Refund", input).Get(ctx, &result)

// 		if refundErr != nil {
// 			return "",
// 				fmt.Errorf("Deposit: failed to deposit money into %v: %v. Money could not be returned to %v: %w",
// 					input.TargetAccount, depositErr, input.SourceAccount, refundErr)
// 		}

// 		return "", fmt.Errorf("Deposit: failed to deposit money into %v: Money returned to %v: %w",
// 			input.TargetAccount, input.SourceAccount, depositErr)
// 	}

// 	result := fmt.Sprintf("Transfer complete (transaction IDs: %s, %s)", withdrawOutput, depositOutput)
// 	return result, nil
// }

// func (a *Activities) Withdraw(ctx context.Context, data PaymentDetails) (string, error) {
// 	referenceID := fmt.Sprintf("%s-withdrawal", data.ReferenceID)
// 	confirmation, err := a.bank.Withdraw(data.SourceAccount, data.Amount, referenceID)
// 	return confirmation, err
// }

// func (a *Activities) Deposit(ctx context.Context, data PaymentDetails) (string, error) {
// 	referenceID := fmt.Sprintf("%s-deposit", data.ReferenceID)
// 	confirmation, err := a.bank.Deposit(data.TargetAccount, data.Amount, referenceID)
// 	return confirmation, err
// }

// func (a *Activities) Refund(ctx context.Context, data PaymentDetails) (string, error) {
// 	referenceID := fmt.Sprintf("%s-refund", data.ReferenceID)
// 	confirmation, err := a.bank.Deposit(data.SourceAccount, data.Amount, referenceID)
// 	return confirmation, err
// }
