package simpleutils

import (
	"github.com/sirupsen/logrus"
)

type ContextHook struct {
}

func (hook ContextHook) Levels() []logrus.Level {
	return logrus.AllLevels
}
func (hook ContextHook) Fire(entry *logrus.Entry) error {
	//if entry.Context != nil {
	//	if trace, ok := entry.Context.Value("simple_trace_id").(string); ok {
	//		entry.Data["simple_trace_id"] = trace
	//	}
	//}
	//if pc, file, line, ok := runtime.Caller(8); ok {
	//	funcName := runtime.FuncForPC(pc).Name()
	//	entry.Data["file"] = path.Base(file)
	//	entry.Data["func"] = path.Base(funcName)
	//	entry.Data["line"] = line
	//}
	return nil
}
