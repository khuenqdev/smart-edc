package connector

import (
    "bytes"
    "fmt"
    "net"
)

const (
    defaultHost     = "127.0.0.1"
    defaultPort     = "8020"
    defaultProtocol = "tcp"
)

type WifiConnector struct {
    Host     string
    Port     string
    Protocol string
}

func NewWifiConnector(host, port, protocol string) *WifiConnector {
    if "" == host {
        host = defaultHost
    }

    if "" == port {
        port = defaultPort
    }

    if "" == protocol {
        protocol = defaultProtocol
    }

    return &WifiConnector{
        Host:     host,
        Port:     port,
        Protocol: protocol,
    }
}

func (c *WifiConnector) SendData(data string) ([]byte, error) {
    address := fmt.Sprintf("%s:%s", c.Host, c.Port)
    fmt.Println("Dialing", address)

    tcpAddr, err := net.ResolveTCPAddr("tcp", address)

    if err != nil {
        fmt.Println("Resolve TCP Address failed:", err.Error())
        return nil, err
    }

    conn, err := net.DialTCP(c.Protocol, nil, tcpAddr)

    if nil != err {
        fmt.Println("Failed!")
        return nil, err
    }

    // Send data to payment terminal
    fmt.Println("Writing data")
    _, err = conn.Write([]byte(data))

    if nil != err {
        fmt.Println("Failed!")
        return nil, err
    }

    _ = conn.CloseWrite()

    response := make([]byte, 2048)
    fmt.Println("Reading response")

    for !validResponseData(response) {
        response = make([]byte, 2048)
        readLen, _ := conn.Read(response)
        response = bytes.Trim(response, "\x00")
        fmt.Println("Read length", readLen)
    }

    err = conn.Close()

    if nil != err {
        return nil, err
    }

    return response, nil
}
