package reportutil

import (
	"fmt"
	"time"
)

type Heartbeat struct {
	pulse     <-chan time.Time
	heartbeat chan HeartbeatPacket
	count     *int64
	started   time.Time
}

type HeartbeatPacket struct {
	Count    int64
	Duration time.Duration
}

func NewHeartbeat(d time.Duration) Heartbeat {
	return Heartbeat{pulse: time.Tick(d), heartbeat: make(chan HeartbeatPacket), count: new(int64), started: time.Now()}
}

// 发送心跳信息
func (h *Heartbeat) SendPluse() {
	select {
	case <-h.pulse:
		select {
		case h.heartbeat <- HeartbeatPacket{Count: *h.count, Duration: time.Since(h.started)}:
		default:
		}
	default:
	}
}

// 监听心跳
func (h *Heartbeat) Output() <-chan HeartbeatPacket {
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

func PrintHeartbeat(hp HeartbeatPacket) {
	fmt.Printf("count:%d ,duration:%.2f\n", hp.Count, hp.Duration.Minutes())
}
