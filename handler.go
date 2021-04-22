package jobs

import (
	"log"
	"sync"

	tools "github.com/mt1976/mwtgostringtools"
	cron "github.com/robfig/cron/v3"
)

//CONST_CONFIG_FILE is cheese
const (
	Scheduled = "Scheduled"
	Adhoc     = "Ad-Hoc"
)

func main() {

	Start()

	for {
		wg := sync.WaitGroup{}
		wg.Add(1)
		wg.Wait()
	}

}

// Start Initialises the Required Jobs
func Start() {
	logit("JobHandler", "INIT", "JOB SCHEDULER")
	c := cron.New()
	// Insert Jobs Here
	c.AddFunc("@every 10s", func() { RunJobHeartBeat(Scheduled) })
	c.AddFunc("@every 1m", func() { RunJobFXSPOT(Scheduled) })
	//
	log.Println(c.Entries())
	c.Start()
}

func logit(funcName string, actionType string, data string) {
	log.Println(funcName, tools.Dquote(actionType), data)
}
