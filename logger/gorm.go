// author : 颜洪毅
// e-mail : yhyzgn@gmail.com
// time   : 2020-04-09 18:27
// version: 1.0.0
// desc   :

package logger

import (
	"context"
	"time"

	"github.com/pkg/errors"

	"gorm.io/gorm/logger"
)

// Gorm ...
type Gorm struct {
	level logger.LogLevel
}

// NewGorm ...
func NewGorm() *Gorm {
	return &Gorm{}
}

// LogMode ...
func (g *Gorm) LogMode(level logger.LogLevel) logger.Interface {
	g.level = level
	return g
}

// Info ...
func (g *Gorm) Info(ctx context.Context, s string, i ...interface{}) {
	if g.level >= logger.Info {
		InfoF(s, i...)
	}
}

// Warn ...
func (g *Gorm) Warn(ctx context.Context, s string, i ...interface{}) {
	if g.level >= logger.Warn {
		WarnF(s, i...)
	}
}

// Error ...
func (g *Gorm) Error(ctx context.Context, s string, i ...interface{}) {
	if g.level >= logger.Error {
		ErrorF(s, i...)
	}
}

// Trace ...
func (g *Gorm) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	if g.level <= logger.Silent {
		return
	}

	switch {
	case err != nil && g.level >= logger.Error && (!errors.Is(err, logger.ErrRecordNotFound)):
		sql, rows := fc()
		if rows == -1 {
			Trace(sql, "-", err.Error())
		} else {
			Trace(sql, rows, err.Error())
		}
	default:
		sql, rows := fc()
		if rows == -1 {
			Trace(sql)
		} else {
			Trace(sql, rows)
		}
	}
}
