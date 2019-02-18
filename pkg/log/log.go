package log

import (
	"fmt"

	"github.com/go-logr/logr"
	"github.com/go-logr/zapr"
	"go.uber.org/zap"
	"sigs.k8s.io/controller-runtime/pkg/runtime/log"
)

// Logger is a simple logging interface for Go.
var Logger logr.Logger

func init() {
	// Build a zap development logger.
	zapLog, err := zap.NewDevelopment()
	if err != nil {
		panic(fmt.Sprintf("error building logger: %v", err))
	}
	// zapr defines an implementation of the Logger
	// interface built on top of Zap (go.uber.org/zap).
	Logger = zapr.NewLogger(zapLog)
	Logger.Info("started zapr logger")
}

// SetRuntimeLogger sets a concrete logging implementation for all
// controller-runtime deferred Loggers.
func SetRuntimeLogger(logger logr.Logger) {
	// Set a concrete logging implementation for all
	// controller-runtime deferred Loggers.
	log.SetLogger(logger)
}
