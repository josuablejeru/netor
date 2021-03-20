package collector

import (
	"github.com/showwin/speedtest-go/speedtest"
	log "github.com/sirupsen/logrus"
)

// GetTargets fetches all targets to run the speedtest on
func GetTargets() (targets speedtest.Servers) {
	user, err := speedtest.FetchUserInfo()
	if err != nil {
		log.Fatal(err)
	}

	serverList, err := speedtest.FetchServerList(user)
	if err != nil {
		log.Fatal(err)
	}

	targets, _ = serverList.FindServer([]int{})

	return targets
}

// StartSpeedtesting retuns
func StartSpeedtesting(target *speedtest.Server) {
	go func() {
		target.PingTest()
		target.DownloadTest(false)
		target.UploadTest(false)
	}()
}
