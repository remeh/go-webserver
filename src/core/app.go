package core;

import (
    "fmt"
    "encoding/json"
    "io/ioutil"
    "container/list"
    "net/http"
    "time"
)

// ---------------------- 
// Declarations

/**
 * The application struct.
 * @author Rémy MATHIEU
 */
type App struct {
    pages   list.List;
    Router  Router;
}

// ---------------------- 
// Methods

/**
 * App initialization.
 */
func (a *App) Init() {
    fmt.Println("------------------------");
    fmt.Println("⡀⣀ ⢀⡀ ⣀⣀  ⢀⡀ ⣇⡀ ⠃ ⢀⣀   ⢀⣀ ⠄ ⣀⣀  ⣀⡀ ⡇ ⢀⡀   ⢀⡀ ⢀⡀   ⡀ ⢀ ⢀⡀ ⣇⡀ ⢀⣀ ⢀⡀ ⡀⣀ ⡀⢀ ⢀⡀ ⡀⣀");
    fmt.Println("⠏  ⠣⠭ ⠇⠇⠇ ⠣⠭ ⠇⠸   ⠭⠕   ⠭⠕ ⠇ ⠇⠇⠇ ⡧⠜ ⠣ ⠣⠭   ⣑⡺ ⠣⠜   ⠱⠱⠃ ⠣⠭ ⠧⠜ ⠭⠕ ⠣⠭ ⠏  ⠱⠃ ⠣⠭ ⠏ ");
    fmt.Println("------------------------");

    // init the router
    a.InitRouter();

    /*
    // Read the JSON configuration
    a.readConfiguration();
    */
}

func (a *App) InitRouter() {
    a.Router.Init();
    a.Router.Start();
}

func (a *App) Start() {
    // TODO
    http.ListenAndServe(":8080", nil);
}

func (a *App) readConfiguration() {
    /*
     * Read the file
     */

    data, err := ioutil.ReadFile("app/config.json"); // XXX hardcoded filenae
    if (err != nil) {
        fmt.Printf(" x Unable to read the config : error while reading the file : \n%s\n",err);
        return;
    }

    /*
     * Unmarshal the json.
     */

    var config ConfigurationFormat;
    err = json.Unmarshal(data, &config);
    if (err != nil) {
        fmt.Printf(" x Unable to read the router configuration : error while unmarshaling the data : \n%s\n",err);
    }

    /*
     * Evaluate the configuration.
    // TODO remove
     */
    //a.router.evalutateConfiguration(config);
}

func logAccess(request *http.Request, fail bool, end string) {
    start := " -";
    if (fail) {
        start = " x";
    }
    referer := "";
    if (len(request.Referer()) != 0) {
        referer = fmt.Sprintf("with referer[%s]", request.Referer());
    }
    userAgent := "";
    if (len(request.UserAgent()) != 0) {
        userAgent = fmt.Sprintf(" UserAgent[%s] ", request.UserAgent());
    }
    var ip string = request.Header.Get("X-Forwarded-For");
    if (len(ip) == 0) {
        ip = request.RemoteAddr;
    }
    fmt.Printf("%s [%s] %s -> %s %s for %s %s%s%s\n", start, time.Now().Format("2006-01-02 15:04:05"), ip, request.Method, request.Proto, request.URL.Path, referer, userAgent, end);
}
