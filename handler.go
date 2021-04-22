package jobs

import (
	"log"
	"runtime"
	"strings"

	cron "github.com/robfig/cron/v3"
)

//CONST_CONFIG_FILE is cheese
const (
	Scheduled = "Scheduled"
	Adhoc     = "AdHoc"
)

// Start Initialises the Required Jobs
func Start() {
	//funcName = "Start"
	logit("JobHandler", "JOB SCHEDULER")
	c := cron.New()
	// Insert Jobs Here
	c.AddFunc("@every 10s", func() { RunJobHeartBeat(Scheduled) })
	logit("RunJobHeartBeat", "@every 10s")
	c.AddFunc("@every 1m", func() { RunJobFXSPOT(Scheduled) })
	logit("RunJobFXSPOT", "@every 1m")
	//log.Println(runtime.Caller(0))
	//log.Println(c.Entries())
	c.Start()
}

func logit(actionType string, data string) {
	_, caller, _, _ := runtime.Caller(1)
	outcall := strings.Split(caller, "/")
	depth := len(outcall) - 1
	depth2 := depth - 1
	log.Println(len(outcall), depth, depth2)
	callerName := outcall[depth2] + "/" + outcall[depth]
	log.Println(callerName, actionType, data)
}
