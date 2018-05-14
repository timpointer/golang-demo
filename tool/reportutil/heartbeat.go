package reportutil

import "time"

type Heartbeat struct {
	pulse     <-chan time.Time
	heartbeat chan interface{}
	count     *int64
}

func NewHeartbeat(d time.Duration) Heartbeat {
	return Heartbeat{pulse: time.Tick(d), heartbeat: make(chan interface{}), count: new(int64)}
}

// 发送心跳信息
func (h *Heartbeat) SendPluse() {
	select {
	case <-h.pulse:
		select {
		case h.heartbeat <- *h.count:
		default:
		}
	default:
	}
}

// 监听心跳
func (h *Heartbeat) Output() <-chan interface{} {
	return h.heartbeat
}

// 心跳记录数器
func (h *Heartbeat) Add() {
	*h.count++
}

// 关闭心跳
func (h *Heartbeat) Close() {
	close(h.heartbeat)
}
