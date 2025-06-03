package ctxkeys

type ctxKey string

const (
	CtxKeyApp            ctxKey = "app"
	CtxKeyRuntime        ctxKey = "runtime"
	CtxKeyEnv            ctxKey = "env"
	CtxKeyAppVersion     ctxKey = "app_version"
	CtxKeyCorrelationID  ctxKey = "correlation_id"
	CtxKeyPath           ctxKey = "path"
	CtxKeyMethod         ctxKey = "method"
	CtxKeyIP             ctxKey = "ip"
	CtxKeyPort           ctxKey = "port"
	CtxKeySrcIP          ctxKey = "src_ip"
	CtxKeyRT             ctxKey = "rt"
	CtxKeyRC             ctxKey = "rc"
	CtxKeyHeader         ctxKey = "header"
	CtxKeyRequest        ctxKey = "req"
	CtxKeyResponse       ctxKey = "resp"
	CtxKeyError          ctxKey = "error"
	CtxKeyTraceID        ctxKey = "trace_id"
	CtxKeyAdditionalData ctxKey = "additional_data"
)
