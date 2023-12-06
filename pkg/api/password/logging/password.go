package password

import (
	"time"

	"github.com/labstack/echo/v4"

	stems "github.com/Soapstone-Services/go-template-2024"
	"github.com/Soapstone-Services/go-template-2024/pkg/api/password"
)

// New creates new password logging service
func New(svc password.Service, logger stems.Logger) *LogService {
	return &LogService{
		Service: svc,
		logger:  logger,
	}
}

// LogService represents password logging service
type LogService struct {
	password.Service
	logger stems.Logger
}

const name = "password"

// Change logging
func (ls *LogService) Change(c echo.Context, id int, oldPass, newPass string) (err error) {
	defer func(begin time.Time) {
		ls.logger.Log(
			c,
			name, "Change password request", err,
			map[string]interface{}{
				"req":  id,
				"took": time.Since(begin),
			},
		)
	}(time.Now())
	return ls.Service.Change(c, id, oldPass, newPass)
}
