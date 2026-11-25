# Principles

Architecture and engineering principles that every change must respect. The binding rules
are in `AGENTS.md`; this file explains the reasoning behind them.

## 1. Each service owns its data (NFR4)

Every service has its own SQLite database and is the only one that writes to it. Data from
another service is reached through its API or its events, never through its database.
Cross-service references are logical IDs without a foreign key.

**Why:** the project must be a real distributed system with independently deployable
components, not one application split into folders.

## 2. Failures stay local (NFR1, NFR2)

A failure in Notification or Ticketing must not stop passengers from consulting the network
and schedules. No service refuses to start because another one is down. Every outbound HTTP
call has a timeout, bounded retries with backoff, and a fallback.

**Why:** the project description requires that the failure of one part does not make the
whole system unavailable.

## 3. Sync for questions, events for facts (NFR6)

REST is used when the caller needs an answer now (Web App → Core / Ticketing, Ticketing →
Core). Events through RabbitMQ are used to announce that something happened (`TripDelayed`,
`TripCancelled`, `TripDiverted`, `DisruptionResolved`, `TicketPurchased`, `TicketExpired`).
Business flows never call Notification synchronously.

## 4. Effects happen once (NFR3, AM-04, AM-05, AM-12, AM-15)

Retries and duplicate messages are expected. Ticket purchase and validation take an
`Idempotency-Key`; the key is enforced by a `UNIQUE` constraint in the same transaction as
the effect. Event consumers deduplicate by `event_id`. A passenger gets at most one
notification per user + disruption + type.

**Why:** the project description explicitly forbids issuing the same ticket, recording the
same validation, or sending the same notification twice.

## 5. History is kept (MI-07)

Business records are never hard-deleted. Use status fields, `is_active` or `RESOLVED`.

## 6. The server decides who can do what

Authentication and role are checked on the server for every protected endpoint. A passenger
only sees their own tickets, subscriptions and notifications; the user id comes from the
JWT, never from the request.

## 7. Everything observable (NFR7)

Every service exposes `GET /health` and writes structured logs.

## 8. Every change is verified (NFR8)

Every push is built, tested and checked automatically. Every behaviour change ships with
tests, and AI-generated output is reviewed by a human before it is merged.
