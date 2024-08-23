package database

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/log"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
	"gorm.io/gorm/utils"
)

type DatabaseLogger struct {
	IgnoreRecordNotFoundError bool
	Logger                    *log.Logger
	LogLevel                  gormLogger.LogLevel
	SlowThreshold             time.Duration
}

func NewDatabaseLogger() DatabaseLogger {
	prefixStyle := lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("43"))

	logger := log.With()
	logger.SetPrefix(prefixStyle.Render("database"))

	return DatabaseLogger{
		IgnoreRecordNotFoundError: false,
		Logger:                    logger,
		LogLevel:                  gormLogger.Warn,
		SlowThreshold:             100 * time.Millisecond,
	}
}

func (dbl DatabaseLogger) SetAsDefault() {
	gormLogger.Default = dbl
}

func (dbl DatabaseLogger) LogMode(level gormLogger.LogLevel) gormLogger.Interface {
	return DatabaseLogger{
		IgnoreRecordNotFoundError: dbl.IgnoreRecordNotFoundError,
		Logger:                    dbl.Logger,
		LogLevel:                  level,
		SlowThreshold:             dbl.SlowThreshold,
	}
}

func (dbl DatabaseLogger) Info(_ context.Context, str string, args ...interface{}) {
	if dbl.LogLevel < gormLogger.Info {
		return
	}

	dbl.Logger.Infof(str, args...)
}

func (dbl DatabaseLogger) Warn(_ context.Context, str string, args ...interface{}) {
	if dbl.LogLevel < gormLogger.Warn {
		return
	}

	dbl.Logger.Warnf(str, args...)
}

func (dbl DatabaseLogger) Error(_ context.Context, str string, args ...interface{}) {
	if dbl.LogLevel < gormLogger.Error {
		return
	}

	dbl.Logger.Errorf(str, args...)
}

func (dbl DatabaseLogger) Trace(_ context.Context, begin time.Time, fc func() (string, int64), err error) {
	// Silent
	if dbl.LogLevel <= 0 {
		return
	}
	slowLogStyle := lipgloss.NewStyle().Faint(true).Foreground(lipgloss.Color("120"))

	currentFile := utils.FileWithLineNum()
	newLog := dbl.Logger.With()
	newLog.SetCallerFormatter(func(_ string, _ int, _ string) string {
		return trimCallerPath(currentFile, 3)
	})

	dbl.Logger = newLog

	elapsed := time.Since(begin)
	sql, rows := fc()

	switch {
	case err != nil && dbl.LogLevel >= gormLogger.Error && (!errors.Is(err, gorm.ErrRecordNotFound) || !dbl.IgnoreRecordNotFoundError):
		dbl.Logger.Error(
			err.Error(),
			"duration",
			fmt.Sprintf("%.3fms", float64(elapsed.Nanoseconds())/1e6),
			"sql",
			sql,
			"rows",
			rows,
		)
	case elapsed > dbl.SlowThreshold && dbl.SlowThreshold != 0 && dbl.LogLevel >= gormLogger.Warn:
		slowLog := slowLogStyle.Render(fmt.Sprintf("SLOW SQL >= %v", dbl.SlowThreshold))

		dbl.Logger.Warn(
			fmt.Sprintf("%s %s", err.Error(), slowLog),
			"duration",
			fmt.Sprintf("%.3fms", float64(elapsed.Nanoseconds())/1e6),
			"sql",
			sql,
			"rows",
			rows,
		)
	case dbl.LogLevel == gormLogger.Info:
		dbl.Logger.Debug(
			"Chamado query",
			"duration",
			fmt.Sprintf("%.3fms", float64(elapsed.Nanoseconds())/1e6),
			"sql",
			sql,
			"rows",
			rows,
		)
	}
}

func trimCallerPath(path string, n int) string {
	if n <= 0 {
		return path
	}

	idx := strings.LastIndexByte(path, '/')
	if idx == -1 {
		return path
	}

	for i := 0; i < n-1; i++ {
		idx = strings.LastIndexByte(path[:idx], '/')
		if idx == -1 {
			return path
		}
	}

	return path[idx+1:]
}
