package main

import (
    "context"
    "fmt"
    "log"
    "net/http"
    "os"
    "os/signal"
    "syscall"
    "time"

    "github.com/rs/cors"

    "git02.smartosc.com/production/smartedc-connector/smartedc"
)

const (
    defaultGrpcAddress = "127.0.0.1:9090"
)

func main() {
    httpServer()
}

func httpServer() {
    errChan := make(chan error)
    c := make(chan os.Signal, 1)
    signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
    srv := smartedc.NewSmartEdcServer()

    // Routes
    mux := http.NewServeMux()
    mux.HandleFunc("/", srv.HandleTerminalTest)
    mux.HandleFunc("/payment/e-wallet/pay", srv.HandleEWalletPayment)
    mux.HandleFunc("/payment/credit-card/pay", srv.HandleCreditCardPayment)

    s := &http.Server{
        Handler: cors.AllowAll().Handler(mux),
        Addr:    defaultGrpcAddress,
    }

    go func() {
        fmt.Println("server started at " + defaultGrpcAddress)
        err := s.ListenAndServe()
        if nil != err {
            log.Fatal(err)
        }
    }()

    <-c
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    if err := s.Shutdown(ctx); err != nil {
        log.Fatal(err)
    }
    defer cancel()
    log.Printf("exit %e", <-errChan)
}
