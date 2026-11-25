// Package domain decides who is affected by an event (subscriptions, FR20; ticket holders, FR9,
// with IN-01 still open), applies preferences (FR19) and builds the deduplication key
// user + disruption + type (AM-12).
package domain
