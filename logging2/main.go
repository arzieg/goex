package main

import (
	"errors"
	"log/slog"
	"os"
)

type PaymentError struct {
	Code    string
	Message string
	Cause   error
}

func (pe PaymentError) Error() string {
	return pe.Message
}

func (pe PaymentError) LogValue() slog.Value {
	return slog.GroupValue(
		slog.String("code", pe.Code),
		slog.String("message", pe.Message),
		slog.String("cause", pe.Cause.Error()),
	)
}

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	causeErr := errors.New("network timeout")
	err := PaymentError{
		Code:    "GATEWAY_UNREACHABLE",
		Message: "Failed to reach payment gateway",
		Cause:   causeErr,
	}

	logger.Error("Payment operation failed", slog.Any("error", err))
}
