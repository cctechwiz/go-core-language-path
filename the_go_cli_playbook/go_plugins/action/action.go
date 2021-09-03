package main

import (
	"flag"
	"log"
	"plugin"
)

func main() {
	path := flag.String("plugin", "", "Plugin to execute")

	flag.Parse()

	if *path == "" {
		log.Fatal("Path to plugin must be provided")
	}

	p, err := plugin.Open(*path)
	if err != nil {
		log.Fatal(err)
	}

	f, err := p.Lookup("ThingToDo")
	if err != nil {
		log.Fatal(err)
	}
	thingToDo, ok := f.(func())
	if !ok {
		log.Fatal("COuld not find function 'ThingToDo' in plugin")
	}
	thingToDo()
	log.Println("Did the thing")
}
