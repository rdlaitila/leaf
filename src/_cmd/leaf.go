package main

import(
    "os"
    
    "gopkg.in/alecthomas/kingpin.v2"
)

var(
    app = kingpin.New("leaf", "Lua Enhanced Application Framework")
    
    //Run subcommand
    runcmd         = app.Command("run", "Runs a Leaf Application")
    runcmd_apppath = runcmd.Arg("name/path", "The path to your application").Required().String()
)

func main() {
    switch kingpin.MustParse(app.Parse(os.Args[1:])) {
        case runcmd.FullCommand():
            run(*runcmd_apppath)
    }
}

func run(apppath string) {
    
}

func start(apppath string) {

}

func install(appurl string) {

}

func search(appname string) {

}

func uninstall(appname string) {

}

func update(appname string) {

}