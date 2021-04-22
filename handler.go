package jobs

import (
	"fmt"
	"log"
	"sync"

	"github.com/davecgh/go-spew/spew"
	cron "github.com/robfig/cron/v3"
)

func main() {

	Start()

	for {
		wg := sync.WaitGroup{}
		wg.Add(1)
		wg.Wait()
	}

}

func Start() {
	log.Println("INSTANCE")
	c := cron.New()
	c.AddFunc("@every 10s", func() { fmt.Println("tick every 10 seconds") })
	c.AddFunc("@every 1m", func() { RunJobFXSPOT("SCHEDULED") })
	spew.Dump(c.Entries())
	c.Start()
}
