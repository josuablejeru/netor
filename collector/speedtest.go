package collector

import (
	"github.com/showwin/speedtest-go/speedtest"
	log "github.com/sirupsen/logrus"
)

// PrepareSession fetches all targets to run the speedtest on
func PrepareSession() (targets speedtest.Servers) {
	user, err := speedtest.FetchUserInfo()
	if err != nil {
		log.Fatal("Could not fetch required user information")
	}

	serverList, err := speedtest.FetchServerList(user)
	if err != nil {
		log.Fatal("could not fetch server list")
	}

	targets, _ = serverList.FindServer([]int{})

	return targets
}
