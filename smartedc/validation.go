package smartedc

import (
    "google.golang.org/grpc/codes"
    "google.golang.org/grpc/status"
)

func validateCreditCardRequest(r *CreditCardPaymentRequest) error {
    if "" == r.TradeType {
        return status.Error(codes.InvalidArgument, "Trade type is empty")
    }

    if 0 == r.Amount {
        return status.Error(codes.InvalidArgument, "Payment amount must not be zero")
    }

    if "" == r.PosRefNo {
        return status.Error(codes.InvalidArgument, "POS reference number is missing")
    }

    if "" == r.TransactionType {
        return status.Error(codes.InvalidArgument, "Transaction type is empty")
    }

    return nil
}

func validateEWalletRequest(r *EWalletPaymentRequest) error {
    if "" == r.TradeType {
        return status.Error(codes.InvalidArgument, "Trade type is empty")
    }

    if 0 == r.Amount {
        return status.Error(codes.InvalidArgument, "Payment amount must not be zero")
    }

    if "" == r.PosRefNo {
        return status.Error(codes.InvalidArgument, "POS reference number is missing")
    }

    if "" == r.TransactionType {
        return status.Error(codes.InvalidArgument, "Transaction type is empty")
    }

    if "" == r.ServiceType {
        return status.Error(codes.InvalidArgument, "Service type is empty")
    }

    return nil
}