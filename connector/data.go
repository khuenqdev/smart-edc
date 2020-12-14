package connector

import (
    "encoding/xml"
    "fmt"
    "strings"
)

type CommonResponse struct {
    XMLName      xml.Name `json:"-" xml:"xml"`
    PosRefNo     string   `json:"pos_ref_no,omitempty" xml:"pos_ref_no,omitempty"`
    ResponseCode string   `json:"response_code,omitempty" xml:"response_code,omitempty"`
    ResponseMsg  string   `json:"response_msg,omitempty" xml:"response_msg,omitempty"`
    InvoiceNo    string   `json:"invoice_no,omitempty" xml:"invoice_no,omitempty"`
    Amount       float64  `json:"amount,omitempty" xml:"amount,omitempty"`
    ErrorMessage string   `json:"error_message,omitempty" xml:"error_message,omitempty"`
}

func validResponseData(r []byte) bool {
    data := &CommonResponse{}

    fmt.Println(string(r))
    err := xml.Unmarshal(r, data)

    if nil != err {
        if strings.Contains(err.Error(), "XML syntax error on line 1: expected element name after <") {
            return true
        }

        return false
    }

    return true
}
