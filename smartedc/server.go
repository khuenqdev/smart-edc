package smartedc

import (
    "encoding/json"
    "fmt"
    "git02.smartosc.com/production/smartedc-connector/connector"
    "io/ioutil"
    "net/http"
)

type SmartEdcService interface {
    HandleCreditCardPayment(w http.ResponseWriter, r *http.Request)
    HandleEWalletPayment(w http.ResponseWriter, r *http.Request)
}

type SmartEdcServer struct {
}

func NewSmartEdcServer() *SmartEdcServer {
    return &SmartEdcServer{}
}

func (s *SmartEdcServer) HandleTerminalTest(w http.ResponseWriter, r *http.Request) {
    _, _ = fmt.Fprintf(w, "SmartEDC Gateway Accessed!")
    fmt.Println("Gateway accessed!")
}

func (s *SmartEdcServer) HandleCreditCardPayment(w http.ResponseWriter, r *http.Request) {
    w.Header().Add("Content-Type", "application/json")
    reqBody, _ := ioutil.ReadAll(r.Body)
    request := &CreditCardPaymentRequest{}
    terminalResponse := make([]byte, 2048)
    jsonEncoder := json.NewEncoder(w)

    err := json.Unmarshal(reqBody, request)

    if nil != err {
        _ = jsonEncoder.Encode(createCreditCardErrorResponse(request, err, "Cannot decode credit card request"))
        return
    }

    if err = validateCreditCardRequest(request); nil != err {
        _ = jsonEncoder.Encode(createCreditCardErrorResponse(request, err, "Invalid credit card request"))
        return
    }

    requestData, err := BuildRequestData(request)

    if nil != err {
        _ = jsonEncoder.Encode(createCreditCardErrorResponse(request, err, "Cannot convert credit card request to XML"))
        return
    }

    switch request.ConnectionType {
    case connectionTypeWifi:
        fmt.Println("Creating WIFI connection to the payment terminal")
        wifiConnector := connector.NewWifiConnector(request.TcpHost, request.TcpPort, "tcp")
        terminalResponse, err = wifiConnector.SendData(requestData)

        if nil != err {
            _ = jsonEncoder.Encode(createCreditCardErrorResponse(request, err, "Error when processing credit card payment via WIFI connection"))
            return
        }
    case connectionTypeUsb:
        usbConnector := connector.NewUsbConnector(request.SerialPort)
        terminalResponse, err = usbConnector.SendData(requestData)

        if nil != err {
            _ = jsonEncoder.Encode(createCreditCardErrorResponse(request, err, "Error when processing credit card payment via USB connection"))
            return
        }
    }

    fmt.Println("Data successfully sent to the payment terminal")
    resp, err := PrepareResponseData(terminalResponse, &CreditCardPaymentResponse{})

    if nil != err {
        _ = jsonEncoder.Encode(createCreditCardErrorResponse(request, err, "Error when handling credit card payment response"))
        return
    }

    err = jsonEncoder.Encode(resp)

    if nil != err {
        _ = jsonEncoder.Encode(createCreditCardErrorResponse(request, err, "Cannot encode credit card response"))
        return
    }
}

func (s *SmartEdcServer) HandleEWalletPayment(w http.ResponseWriter, r *http.Request) {
    w.Header().Add("Content-Type", "application/json")
    reqBody, _ := ioutil.ReadAll(r.Body)
    request := &EWalletPaymentRequest{}
    terminalResponse := make([]byte, 2048)
    jsonEncoder := json.NewEncoder(w)

    err := json.Unmarshal(reqBody, request)

    if nil != err {
        _ = jsonEncoder.Encode(createEWalletErrorResponse(request, err, "Cannot decode e-wallet request"))
        return
    }

    if err = validateEWalletRequest(request); nil != err {
        _ = jsonEncoder.Encode(createEWalletErrorResponse(request, err, "Invalid e-wallet request"))
        return
    }

    requestData, err := BuildRequestData(request)

    if nil != err {
        _ = jsonEncoder.Encode(createEWalletErrorResponse(request, err, "Cannot convert credit card request to XML"))
        return
    }

    switch request.ConnectionType {
    case connectionTypeWifi:
        fmt.Println("Creating WIFI connection to the payment terminal")
        wifiConnector := connector.NewWifiConnector(request.TcpHost, request.TcpPort, "tcp")
        terminalResponse, err = wifiConnector.SendData(requestData)

        if nil != err {
            _ = jsonEncoder.Encode(createEWalletErrorResponse(request, err, "Error when processing e-wallet payment via WIFI connection"))
            return
        }
    case connectionTypeUsb:
        usbConnector := connector.NewUsbConnector(request.SerialPort)
        terminalResponse, err = usbConnector.SendData(requestData)

        if nil != err {
            _ = jsonEncoder.Encode(createEWalletErrorResponse(request, err, "Error when processing e-wallet payment via USB connection"))
            return
        }
    }

    fmt.Println("Data successfully sent to the payment terminal")
    resp, err := PrepareResponseData(terminalResponse, &EWalletPaymentResponse{})

    if nil != err {
        _ = jsonEncoder.Encode(createEWalletErrorResponse(request, err, "Error when handling e-wallet payment response"))
        return
    }

    err = jsonEncoder.Encode(resp)

    if nil != err {
        _ = jsonEncoder.Encode(createEWalletErrorResponse(request, err, "Cannot encode e-wallet response"))
        return
    }
}