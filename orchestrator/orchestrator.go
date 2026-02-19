package orchestrator

import (
	"os"
	"strings"
	"time"

	"github.com/Shivamingale3/ProcPipe/config"
	"github.com/Shivamingale3/ProcPipe/logger"
	"github.com/Shivamingale3/ProcPipe/monitor"
	"github.com/Shivamingale3/ProcPipe/notify"
	"github.com/Shivamingale3/ProcPipe/process"
)

// Run is the main entry point: spawns process, monitors, notifies.
func Run(cfg *config.Config) (int, error) {
	n := createNotifier(cfg)
	cmdStr := strings.Join(cfg.Command, " ")
	logger.Info("Starting: %s", cmdStr)

	proc, err := process.Start(cfg.Command)
	if err != nil {
		return 1, err
	}
	defer proc.Close()

	startTime := time.Now()
	host, _ := os.Hostname()

	if err := n.SendStarted(notify.StartInfo{
		Command: cmdStr, Host: host,
		Directory: mustGetwd(), StartTime: startTime,
	}); err != nil {
		logger.Warn("Telegram send failed: %s", err)
	}

	logger.Success("Telegram connected — watching output")

	mon := monitor.New(proc.Output(), cfg.Monitor.LogTailLines, cfg.Monitor.InputPatterns)
	mon.Start()

	return handleEvents(mon, proc, n, cmdStr, host, startTime, cfg.DryRun), nil
}
