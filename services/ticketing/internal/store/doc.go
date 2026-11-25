// Package store is the Ticketing SQLite access layer. It owns ticket_types, passenger_tickets,
// ticket_validations and trip_occupancy. Validation runs in one transaction, with the idempotency
// key enforced by a UNIQUE constraint (NFR3, AGENTS.md section 4).
package store
