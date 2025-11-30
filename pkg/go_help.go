package pgk

import (
	"context"
	"fmt"
	"io"
	"log/slog"
)

/*
对slog的日志格式进行自定义
使用: 调一下InitSimpleLogger
*/

// SimpleHandler 自定义的极简日志处理器
type SimpleHandler struct {
	w io.Writer
}

// NewSimpleHandler 创建处理器
func NewSimpleHandler(w io.Writer) *SimpleHandler {
	return &SimpleHandler{w: w}
}

// Enabled 启用所有级别（你可以根据需要修改，比如只启用 Info 以上）
func (h *SimpleHandler) Enabled(ctx context.Context, level slog.Level) bool {
	return true
}

// Handle 核心打印逻辑
func (h *SimpleHandler) Handle(ctx context.Context, r slog.Record) error {
	// 1. 格式化时间
	timeStr := r.Time.Format("15:04:05.000")

	// 2. 处理 Level 部分
	var levelStr string
	if r.Level != slog.LevelInfo {
		// 只有非 INFO 级别才显示 level=XXX
		levelStr = fmt.Sprintf("level=%s ", r.Level.String())
	}

	// 3. 打印基础部分：时间 [Level] 消息
	// 格式：14:30:05.123 [level=DEBUG ]消息内容
	_, err := fmt.Fprintf(h.w, "%s %s%s", timeStr, levelStr, r.Message)
	if err != nil {
		return err
	}

	// 4. 处理后续的 Key-Value 参数 (如果有)
	// 例如：slog.Info("msg", "k", "v") -> 后面追加 k=v
	r.Attrs(func(a slog.Attr) bool {
		fmt.Fprintf(h.w, " %s=%v", a.Key, a.Value)
		return true
	})

	// 换行
	fmt.Fprintln(h.w)
	return nil
}

// WithAttrs 必须实现的方法（这里简单返回h本身，暂不支持子logger携带属性）
// 如果你需要 log.With(...) 功能，这里需要完整实现，对于简单Bot通常不需要
func (h *SimpleHandler) WithAttrs(attrs []slog.Attr) slog.Handler { return h }
func (h *SimpleHandler) WithGroup(name string) slog.Handler       { return h }
