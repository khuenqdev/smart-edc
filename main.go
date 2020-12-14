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

    "github.com/gorilla/handlers"
    "github.com/gorilla/mux"

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
    r := mux.NewRouter()
    r.Use(commonMiddleware)
    r.HandleFunc("/", srv.HandleTerminalTest)
    r.HandleFunc("/payment/e-wallet/pay", srv.HandleEWalletPayment).Methods("POST")
    r.HandleFunc("/payment/credit-card/pay", srv.HandleCreditCardPayment).Methods("POST")
    //r.HandleFunc("/payment/e-wallet/void", srv.HandleEWalletVoid)
    //r.HandleFunc("/payment/credit-card/void", srv.HandleCreditCardVoid)
    //r.HandleFunc("/payment/e-wallet/status", srv.HandleEWalletStatus)

    s := &http.Server{
        Handler: handlers.CORS()(r),
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

func commonMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.Header().Add("Content-Type", "application/json")
        next.ServeHTTP(w, r)
    })
}
