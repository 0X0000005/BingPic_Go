package scheduler

import (
	"log"

	"github.com/robfig/cron/v3"
)

type Scheduler struct {
	cron *cron.Cron
}

func NewScheduler() *Scheduler {
	return &Scheduler{
		cron: cron.New(),
	}
}

func (s *Scheduler) AddJob(schedule string, job func()) error {
	_, err := s.cron.AddFunc(schedule, job)
	return err
}

func (s *Scheduler) Start() {
	s.cron.Start()
	log.Println("调度器已启动")
}

func (s *Scheduler) Stop() {
	s.cron.Stop()
	log.Println("调度器已停止")
}
