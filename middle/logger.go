package middle

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/rs/zerolog"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 记录请求日志
		logger := zerolog.New(os.Stdout).With().Str("request_id", uuid.New().String()).Logger()
		logger.Info().Msg("Request received")
		c.Next()
	}
}
