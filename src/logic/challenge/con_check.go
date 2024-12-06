package challenge_

import (
	"os/exec"
	con "src/const"
	"src/database"
	"time"
)

func Con_check() {
	db := con.Db.Db
	var cons []database.Container
	err := db.Find(&cons).Error
	nowtime := time.Now()
	if err == nil {
		for _, con := range cons {
			timet := con.UpdatedAt
			if nowtime.Sub(timet).Minutes() >= 60 {
				cmd := exec.Command("docker", "stop", con.ContainerID)
				err := cmd.Run()
				if err != nil {
					continue
				}
				cmd = exec.Command("docker", "rm", con.ContainerID)
				err = cmd.Run()
				if err != nil {
					continue
				}

				db.Delete(&con)
			}
		}
	}
	time.AfterFunc(5*time.Minute, Con_check)
}
