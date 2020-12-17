package connector

import (
    "bytes"
    "fmt"
    "github.com/tarm/serial"
    "google.golang.org/grpc/codes"
    "google.golang.org/grpc/status"
    "io/ioutil"
    "strings"
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
        readLen, err := s.Read(x)

        if nil != err {
            return nil, status.Error(codes.Unavailable, "Cannot read data from terminal")
        }

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
    var portNames []string
    windowsPorts := []string{
        "COM1",
        "COM2",
        "COM3",
        "COM4",
        "COM5",
        "COM6",
        "COM7",
        "COM8",
        "COM9",
        "COM10",
    }

    portNames = append(portNames, windowsPorts...)

    files, err := ioutil.ReadDir("/dev")

    if err != nil {
        return portNames
    }

    for _, file := range files {
        if strings.Contains(file.Name(), "ttyACM") || strings.Contains(file.Name(), "tty.usbmodem") {
            fmt.Println(file.Name())
            portNames = append(portNames, "/dev/" + file.Name())
        }
    }

    return portNames
}
