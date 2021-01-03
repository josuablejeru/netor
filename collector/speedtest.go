package collector

import (
	"github.com/showwin/speedtest-go/speedtest"
	log "github.com/sirupsen/logrus"
)

// PrepareSession fetches all targets to run the speedtest on
func PrepareSession() (targets speedtest.Servers) {
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
