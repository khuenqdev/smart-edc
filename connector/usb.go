package connector

import (
    "fmt"
    "github.com/tarm/serial"
    "time"
)

const defaultSerialPort = "COM1"

type UsbConnector struct {
    SerialPort string
}

func NewUsbConnector(SerialPort string) *UsbConnector {
    if "" == SerialPort {
        SerialPort = defaultSerialPort
    }

    return &UsbConnector{SerialPort: SerialPort}
}

func (c *UsbConnector) SendData(data string) ([]byte, error) {
    // Configure connection
    connection := &serial.Config{
        Name: c.SerialPort,
        Baud: 9600,
        StopBits: serial.Stop1,
        ReadTimeout: time.Millisecond * 500,
        Parity: serial.ParityOdd,
        Size: 8,
    }
    fmt.Println("Opening serial port", c.SerialPort)
    s, err := serial.OpenPort(connection)

    if nil != err {
        fmt.Println("Failed!")
        return nil, err
    }

    defer s.Close()

    fmt.Println("Writing data")
    _, err = s.Write([]byte(data))

    if nil != err {
        fmt.Println("Failed!")
        return nil, err
    }

    time.Sleep(10000)
    buf := make([]byte, 2048)
    fmt.Println("Reading response")
    _, err = s.Read(buf)

    if err != nil {
        fmt.Println("Failed!")
        return nil, err
    }

    return buf, nil
}
