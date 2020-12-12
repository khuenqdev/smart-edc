package connector

import (
    "fmt"
    "net"
    "time"
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

    conn, err := net.DialTimeout(c.Protocol, address, 10000000)

    if nil != err {
        fmt.Println("Failed!")
        return nil, err
    }

    if err = conn.SetReadDeadline(time.Now().Add(10 * time.Second)); nil != err {
        fmt.Println("Failed!")
        return nil, err
    }

    defer conn.Close()

    for {
        // Send data to payment terminal
        fmt.Println("Writing data")
        _, err = conn.Write([]byte(data))

        if nil != err {
            fmt.Println("Failed!")
            return nil, err
        }

        time.Sleep(10000)
        response := make([]byte, 2048)
        fmt.Println("Reading response")
        _, err = conn.Read(response)

        if nil != err {
            fmt.Println("Failed!")
            return nil, err
        }

        return response, nil
    }

}
