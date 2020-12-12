package smartedc

import (
    "encoding/xml"
    "fmt"
)

func buildRequestData(r interface{}) (string, error) {
    requestData, err := xml.Marshal(r)

    if nil != err {
        return "", err
    }

    return string(requestData), nil
}

func prepareResponseData(r []byte, output interface{}) (interface{}, error) {
    err := xml.Unmarshal(r, output)
    return output, err
}

func createCreditCardErrorResponse(r *CreditCardPaymentRequest, err error, prefix string) *CreditCardPaymentResponse {
    return &CreditCardPaymentResponse{
        CommonResponse: CommonResponse{
            PosRefNo:     r.PosRefNo,
            ResponseCode: responseCodeGeneral,
            ResponseMsg:  responseMsgFail,
            InvoiceNo:    "",
            Amount:       r.Amount,
            ErrorMessage: fmt.Sprintf("%s: %s", prefix, err.Error()),
        },
        CardNo:           "",
        CardApprovalCode: "",
    }
}

func createEWalletErrorResponse(r *EWalletPaymentRequest, err error, prefix string) *EWalletPaymentResponse {
    return &EWalletPaymentResponse{
        CommonResponse: CommonResponse{
            PosRefNo:     r.PosRefNo,
            ResponseCode: responseCodeGeneral,
            ResponseMsg:  responseMsgFail,
            InvoiceNo:    "",
            Amount:       r.Amount,
            ErrorMessage: fmt.Sprintf("%s: %s", prefix, err.Error()),
        },
        TransactionId: "",
    }
}
