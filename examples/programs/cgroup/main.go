package main

import (
	"github.com/ehids/ebpfmanager"
	"github.com/sirupsen/logrus"
)

var m = &manager.Manager{
	Probes: []*manager.Probe{
		{
			Section:      "cgroup_skb/egress",
			CGroupPath:   "/sys/fs/cgroup/unified",
			EbpfFuncName: "cgroup_egress_func",
		},
	},
}

func main() {
	// Initialize the manager
	if err := m.Init(recoverAssets()); err != nil {
		logrus.Fatal(err)
	}

	// Start the manager
	if err := m.Start(); err != nil {
		logrus.Fatal(err)
	}

	logrus.Println("successfully started, head over to /sys/kernel/debug/tracing/trace_pipe")

	// Generate some network traffic to trigger the probe
	trigger()

	// Close the manager
	if err := m.Stop(manager.CleanAll); err != nil {
		logrus.Fatal(err)
	}
}
