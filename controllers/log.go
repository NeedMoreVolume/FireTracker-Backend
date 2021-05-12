package controllers

import (
	"context"
	"errors"
	"github.com/NeedMoreVolume/FireTracker/gen/log"
	"github.com/NeedMoreVolume/FireTracker/services"
	"github.com/rs/zerolog"
	"time"

	logSvc "github.com/NeedMoreVolume/FireTracker/gen/log"
)

// log service example implementation.
// The example methods log the requests and return zero values.
type logController struct {
	logger     zerolog.Logger
	logService *services.LogService
}

// NewLog returns the log service implementation.
func NewLog(logger zerolog.Logger, ls *services.LogService) logSvc.Service {
	return &logController{
		logger.With().Str("controller", "log").Logger(),
		ls,
	}
}

// Add a log to existing fire
func (s *logController) Create(_ context.Context, p *logSvc.Log) (res *logSvc.Log, err error) {
	res = &logSvc.Log{}
	s.logger.Debug().Msg("log.create")

	newLog, err := logToModel(p)
	if err != nil {
		err = log.MakeBadRequest(err)
		return
	}

	newLog, err = s.logService.Create(newLog)
	if err != nil {
		return
	}

	res = logToTransport(newLog)
	return
}

// Get log and data friends
func (s *logController) Get(_ context.Context, p *logSvc.Log) (res *logSvc.Log, err error) {
	res = &logSvc.Log{}
	s.logger.Debug().Msg("log.get")
	out, err := s.logService.Get(uint(*p.ID))
	res = logToTransport(out)
	return
}

// Update something about a log specifically
func (s *logController) Update(_ context.Context, p *logSvc.Log) (res *logSvc.Log, err error) {
	res = &logSvc.Log{}
	s.logger.Debug().Msg("log.update")
	if p.ID == nil {
		err = log.MakeBadRequest(errors.New("log id is required"))
		return
	}
	tempLog, err := s.logService.Get(uint(*p.ID))
	if err != nil {
		return
	}
	if p.AddedAt != nil {
		tempLog.AddedAt, err = time.Parse(time.RFC3339, *p.AddedAt)
		if err != nil {
			err = log.MakeBadRequest(errors.New("time must be in RFC3339 format"))
			return
		}
	}
	if p.Name != nil {
		tempLog.Name = *p.Name
	}
	if p.FireID != nil {
		tempLog.FireID = uint(*p.FireID)
	}
	if p.Size != nil {
		tempLog.Size = *p.Size
	}
	if p.Weather != nil {
		tempLog.Weather, err = logWeatherToModel(p.Weather)
		if err != nil {
			err = log.MakeBadRequest(err)
			return
		}
		tempLog.WeatherID = &tempLog.Weather.ID
	}

	updatedLog, err := s.logService.Update(tempLog)
	if err != nil {
		return
	}
	res = logToTransport(updatedLog)
	return
}

// Delete some log specifically
func (s *logController) Delete(_ context.Context, p *logSvc.Log) (err error) {
	s.logger.Debug().Msg("log.delete")

	if p.ID == nil || *p.ID == 0 {
		err = log.MakeBadRequest(errors.New("must provide a valid id to delete by"))
		return
	}

	err = s.logService.Delete(uint(*p.ID))

	return
}

// List fires
func (s *logController) List(_ context.Context, p *logSvc.LogListPayload) (res *logSvc.LogList, err error) {
	s.logger.Debug().Msg("log.list")

	theseLogs, count, err := s.logService.List(p)
	if err != nil {
		return
	}

	total := int(count)
	res = &logSvc.LogList{
		Logs: []*logSvc.Log{},
		Pagination: &logSvc.Pagination{
			Total: &total,
		},
	}

	for _, thisLog := range theseLogs {
		res.Logs = append(res.Logs, logToTransport(thisLog))
	}

	return
}
