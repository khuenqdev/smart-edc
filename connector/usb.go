package connector

import (
    "bytes"
    "fmt"
    "github.com/tarm/serial"
    "google.golang.org/grpc/codes"
    "google.golang.org/grpc/status"
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
    s := checkPorts()

    if nil == s {
        return nil, status.Error(codes.Internal, "Cannot open any serial port")
    }

    fmt.Println("Writing data")
    _, err := s.Write([]byte(data))

    if nil != err {
        fmt.Println("Failed!")
        return nil, status.Error(codes.Unavailable, "Cannot send data to terminal")
    }

    var buf []byte
    fmt.Println("Reading response")

    for !validResponseData(buf) {
        x := make([]byte, 2048)
        readLen, _ := s.Read(x)
        buf = append(buf, x...)
        fmt.Println("Read length", readLen)
        buf = bytes.Trim(buf, "\x00")
    }

    err = s.Close()

    if nil != err {
        return nil, status.Error(codes.Unavailable, "Cannot close terminal connection")
    }

    fmt.Println("Terminal connection closed")
    return buf, nil
}

func checkPorts() *serial.Port {
    portList := getPortList()

    for _, p := range portList {
        connection := &serial.Config{
            Name:        p,
            Baud:        9600,
        }
        s, err := serial.OpenPort(connection)

        if nil != err {
            continue
        }

        fmt.Println("Opened port", p)
        return s
    }

    return nil
}

func getPortList() []string {
    return []string{
        "COM1",
        "COM2",
        "COM3",
        "COM4",
        "COM5",
        "/dev/ttyACM0",
        "/dev/ttyACM1",
        "/dev/ttyACM2",
        "/dev/ttyACM3",
        "/dev/ttyACM4",
    }
}
