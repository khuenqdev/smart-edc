package connector

type Connector interface {
    SendData(string) ([]byte, error)
}
