// Package coreclient calls the Core service REST API (trip instance existence, capacity) with a
// timeout, bounded retries with backoff and a cached fallback (NFR2, AGENTS.md section 3).
package coreclient
