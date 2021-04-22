package jobs

import (
	"log"
	"sync"

	"github.com/davecgh/go-spew/spew"
	tools "github.com/mt1976/mwtgostringtools"
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
	c.AddFunc("@every 10s", func() { RunJobHeartBeat("") })
	c.AddFunc("@every 1m", func() { RunJobFXSPOT("SCHEDULED") })
	spew.Dump(c.Entries())
	c.Start()
}

func logit(funcName string, actionType string, data string) {
	log.Println(funcName, tools.Dquote(actionType), data)
}
