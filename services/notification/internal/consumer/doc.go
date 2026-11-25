// Package consumer reads events from RabbitMQ and ignores repeats by event_id, because delivery is
// at-least-once (AM-15, AGENTS.md section 4.4).
package consumer
