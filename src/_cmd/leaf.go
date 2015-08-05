package main

import(
    "os"
    "log"
    
    "gopkg.in/alecthomas/kingpin.v2"
)

var(
    app = kingpin.New("leaf", "Lua Enhanced Application Framework")
    
    //Run subcommand
    runcmd         = app.Command("run", "Runs a Leaf Application")
    runcmd_apppath = runcmd.Arg("path", "The path to your application").Required().String()
)

func main() {
    switch kingpin.MustParse(app.Parse(os.Args[1:])) {
        // Register user
        case runcmd.FullCommand():
            run(*runcmd_apppath)
    }    
    log.Println("Exiting")
}

func run(appname string) {
    
}

func install(appurl string) {

}

func search(appname string) {

}

func uninstall(appname string) {

}

func update(appname string) {

}