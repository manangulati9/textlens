package lib

import (
	"os"
	"strings"
)

func SetWaylandEnvs() {
	envs := []string{"XCURSOR_SIZE=24", "QT_QPA_PLATFORM=wayland"}
	for _, env := range envs {
		parts := strings.Split(env, "=")
		if _, didFind := os.LookupEnv(parts[0]); !didFind {
			os.Setenv(parts[0], parts[1])
		}
	}
}
