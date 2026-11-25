// Package events publishes Core events to RabbitMQ: TripDelayed, TripCancelled, TripDiverted and
// DisruptionResolved, following the schemas in contracts/events/. Publishing must never block or
// fail a business request when the broker is down (NFR1).
package events
