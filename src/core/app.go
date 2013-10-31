package core;

import (
    "fmt"
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
    router  Router;
}

// ---------------------- 
// Methods

/**
 * App initialization.
 */
func (a *App) Init() {
    fmt.Println("------------------------\n");
    fmt.Println(" ⡀⣀ ⢀⡀ ⣀⣀  ⢀⡀ ⣇⡀   ⣰⡁ ⡀⣀   ");
    fmt.Println(" ⠏  ⠣⠭ ⠇⠇⠇ ⠣⠭ ⠇⠸ ⠶ ⢸  ⠏  \n");
    fmt.Println("------------------------");

    // init the router
    a.InitRouter();

    // start
    a.Start(8080);
}

func (a *App) InitRouter() {
    a.router.Init();
    a.router.Start();
}

func (a *App) Start(port int) {
    http.ListenAndServe(fmt.Sprintf(":%d", port), nil);
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
