package main

import (
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
	"hello-world-temporal/app"
	"log"
)

func main() {
	// default Options :localhost:7233, default namespace
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln("Unable create temporal client", err)
	}
	defer c.Close()

	w := worker.New(c, app.GreetingTaskQueue, worker.Options{})
	w.RegisterWorkflow(app.GreetingWorkflow)
	w.RegisterActivity(app.ComposeGreeting)

	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("unable to start worker", err)
	}
}
