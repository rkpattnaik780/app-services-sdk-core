package models
import (
    "errors"
)
// 
type LogLevel int

const (
    DEBUG_LOGLEVEL LogLevel = iota
    TRACE_LOGLEVEL
    WARN_LOGLEVEL
    ERROR_LOGLEVEL
    SEVERE_LOGLEVEL
    WARNING_LOGLEVEL
    INFO_LOGLEVEL
    CONFIG_LOGLEVEL
    FINE_LOGLEVEL
    FINER_LOGLEVEL
    FINEST_LOGLEVEL
)

func (i LogLevel) String() string {
    return []string{"DEBUG", "TRACE", "WARN", "ERROR", "SEVERE", "WARNING", "INFO", "CONFIG", "FINE", "FINER", "FINEST"}[i]
}
func ParseLogLevel(v string) (any, error) {
    result := DEBUG_LOGLEVEL
    switch v {
        case "DEBUG":
            result = DEBUG_LOGLEVEL
        case "TRACE":
            result = TRACE_LOGLEVEL
        case "WARN":
            result = WARN_LOGLEVEL
        case "ERROR":
            result = ERROR_LOGLEVEL
        case "SEVERE":
            result = SEVERE_LOGLEVEL
        case "WARNING":
            result = WARNING_LOGLEVEL
        case "INFO":
            result = INFO_LOGLEVEL
        case "CONFIG":
            result = CONFIG_LOGLEVEL
        case "FINE":
            result = FINE_LOGLEVEL
        case "FINER":
            result = FINER_LOGLEVEL
        case "FINEST":
            result = FINEST_LOGLEVEL
        default:
            return 0, errors.New("Unknown LogLevel value: " + v)
    }
    return &result, nil
}
func SerializeLogLevel(values []LogLevel) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
