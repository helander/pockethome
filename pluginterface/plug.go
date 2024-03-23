package pluginterface

type Host interface {
    InfoLog(msg string) error
    ErrorLog(msg string) error
    PropertyLog(msg string) error
}

type Controller interface {
    Start(host Host) error
}

type Plug interface {
    GetPlugName() string
    Create(name string, url string, extraJson string) *Controller
}
