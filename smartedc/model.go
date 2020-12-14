package smartedc

import "encoding/xml"

type CommonRequest struct {
    XMLName         xml.Name `json:"-" xml:"xml"`
    TcpHost         string   `json:"tcp_host,omitempty" xml:"-"`
    TcpPort         string   `json:"tcp_port,omitempty" xml:"-"`
    SerialPort      string   `json:"serial_port,omitempty" xml:"-"`
    ConnectionType  string   `json:"connection_type,omitempty" xml:"-"`
    TradeType       string   `json:"trade_type,omitempty" xml:"trade_type,omitempty"`
    Amount          float64  `json:"amount,omitempty" xml:"amount,omitempty"`
    PosRefNo        string   `json:"pos_ref_no,omitempty" xml:"pos_ref_no,omitempty"`
    TransactionType string   `json:"transaction_type,omitempty" xml:"transaction_type,omitempty"`
}

type CommonResponse struct {
    XMLName      xml.Name `json:"-" xml:"xml"`
    PosRefNo     string   `json:"pos_ref_no" xml:"pos_ref_no,omitempty"`
    ResponseCode string   `json:"response_code" xml:"response_code,omitempty"`
    ResponseMsg  string   `json:"response_msg" xml:"response_msg,omitempty"`
    InvoiceNo    string   `json:"invoice_no" xml:"invoice_no,omitempty"`
    Amount       float64  `json:"amount" xml:"amount,omitempty"`
    ErrorMessage string   `json:"error_message" xml:"error_message,omitempty"`
}

type EWalletPaymentRequest struct {
    CommonRequest
    ServiceType string `json:"service_type,omitempty" xml:"service_type,omitempty"`
}

type EWalletPaymentResponse struct {
    CommonResponse
    TransactionId string `json:"transaction_id" xml:"transaction_id,omitempty"`
}

type CreditCardPaymentRequest struct {
    CommonRequest
}

type CreditCardPaymentResponse struct {
    CommonResponse
    CardNo           string `json:"card_no" xml:"card_no,omitempty"`
    CardApprovalCode string `json:"card_approval_code" xml:"card_approval_code,omitempty"`
}

type QueryEWalletPaymentStatusRequest struct {
    CommonRequest
    InvoiceNo string `json:"invoice_no,omitempty" xml:"invoice_no,omitempty"`
}

type QueryEWalletPaymentStatusResponse struct {
    CommonResponse
    TransactionId string `json:"transaction_id" xml:"transaction_id,omitempty"`
}

type VoidEWalletPaymentRequest struct {
    CommonRequest
    InvoiceNo string `json:"invoice_no,omitempty" xml:"invoice_no,omitempty"`
}

type VoidEWalletPaymentResponse struct {
    CommonResponse
    TransactionId string `json:"transaction_id" xml:"transaction_id,omitempty"`
}

type VoidCreditCardPaymentRequest struct {
    CommonRequest
    InvoiceNo        string `json:"invoice_no,omitempty" xml:"invoice_no,omitempty"`
    CardApprovalCode string `json:"card_approval_code,omitempty" xml:"card_approval_code,omitempty"`
}

type VoidCreditCardPaymentResponse struct {
    CommonResponse
}
