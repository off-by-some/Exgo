package main

import (
    . "gopkg.in/godo.v1"
)

func tasks(p *Project) {
    // HACK:
    Env = "GOPATH=$GOPATH"

    p.Task("default", D{"build"})

    p.Task("bootstrap", D{}, func() error {
        return Run("gopm get", In{"."})
    })

    p.Task("build", D{"run"}, func() error {
        return Run("gopm build", In{"."})
    }).Watch("**/*.go")

    p.Task("run", D{}, func() error {
        println("Starting server")
        return Start("./output", In{"."})
    }).Watch("**/*.go")

}

func main() {
    Godo(tasks)
}
