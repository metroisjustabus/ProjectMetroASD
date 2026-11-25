<!-- Purpose: describes how the repository is organised into components (services, web app, contracts, deploy), what each folder contains, the current migration state and the open design questions. Complements docs/requirements.md section 6. -->

# Architecture and Repository Structure

The component split (which service does what, and who owns which data) is defined in
`docs/requirements.md` §6 and `AGENTS.md` §3. This document explains how that split is
reflected in the folders of the repository.

## Overview

```
services/
  core/            Core Backend: users, network, schedules, trips, disruptions
  ticketing/       Ticketing & Occupancy: ticket types, tickets, validations, occupancy
  notification/    Notification: subscriptions, preferences, notifications
web/               Web App (React + Vite), passenger UI and staff back-office
contracts/         The agreed interfaces between components (REST + events)
deploy/            Docker Compose files per environment
```

```
Web App ──REST──► Core ────────────── publishes ─┐
   │                ▲                             │
   ├──REST──► Ticketing ── REST (timeout/retry) ──┘ ── publishes ─┐
   │                                                              ▼
   └──REST──► Notification ◄──────────── consumes ──────────── RabbitMQ
```

## Why one folder and one Go module per service

- **Boundaries are enforced by the compiler.** Each service has its own `go.mod`, so it cannot
  import another service's code. Services talk only through REST or events (NFR4,
  `AGENTS.md` §3).
- **Independently deployable.** Each service gets its own image, as required by the project
  statement ("independently deployable supporting components").
- **Explicit interfaces.** The contracts between services live in `contracts/`, in a neutral
  format, and are written before the code.
- **Clear ownership.** One folder per service maps to one owner in the task plan and in
  `.github/CODEOWNERS`.

Trade-off: more setup (one Dockerfile and one CI job per service), and small shared helpers
(health, logging, retry) are copied into each service rather than shared, to avoid coupling.

## Layout of a service

Every service follows the same layout:

| Path | Content |
|---|---|
| `go.mod` | The service's own Go module |
| `main.go` | Startup only: configuration from environment variables, database, migrations, routes, `GET /health` |
| `Dockerfile` | Image for this service only |
| `internal/api/` | HTTP layer: routes, request/response types, JWT and role checks. No business rules |
| `internal/domain/` | Business rules as plain Go, with no database or HTTP. Easiest to unit-test |
| `internal/store/` | SQLite access, transactions, idempotency constraints |
| `internal/events/` | Publishes events to RabbitMQ (Core, Ticketing) |
| `internal/consumer/` | Consumes events from RabbitMQ (Notification) |
| `internal/coreclient/` | REST client for Core with timeout, retries and fallback (Ticketing) |
| `migrations/` | Versioned SQL; an applied migration is never edited |
| `*_test.go` | Tests next to the code they test |

## Components

### `services/core/`

| | |
|---|---|
| Owns | `users`, `stops`, `lines`, `line_stops`, `day_types`, `day_type_exceptions`, `trip_templates`, `trip_instances`, `stop_schedules`, `disruptions` |
| Requirements | FR1–FR8; staff maintenance of network, schedules and disruptions |
| Publishes | `TripDelayed`, `TripCancelled`, `TripDiverted`, `DisruptionResolved` |

### `services/ticketing/`

| | |
|---|---|
| Owns | `ticket_types`, `passenger_tickets`, `ticket_validations`, `trip_occupancy` |
| Requirements | FR10–FR17 |
| Calls | Core (trip existence and capacity), with timeout, retries and cached fallback (NFR2) |
| Publishes | `TicketPurchased`, `TicketExpired` |

### `services/notification/`

| | |
|---|---|
| Owns | `subscriptions`, `notification_preferences`, `notifications` |
| Requirements | FR9, FR18–FR20 |
| Consumes | All six events, deduplicated by `event_id` |

If Notification is down, events wait in RabbitMQ and the rest of the system keeps working (NFR1).

### `web/`

Owns no data. `src/pages/passenger/` and `src/pages/staff/` hold the UI; `src/api/` has one
client per service. If Ticketing is down, ticket pages show the feature as unavailable while
network pages keep working.

### `contracts/`

| Path | Content |
|---|---|
| `openapi/<service>.yaml` | REST contract of each service |
| `events/envelope.json` | Fields common to every event (`event_id`, type, time, payload) |
| `events/<Event>.json` | JSON Schema of each event payload |

### `deploy/`

| File | Use |
|---|---|
| `docker-compose.local.yml` | Local stack built from source |
| `docker-compose.dev.yml` | Development environment, deployed by CD on the self-hosted runner |
| `docker-compose.prod.yml` | Production, with the chosen deployment strategy |
| `.env.example` | Configuration names with dummy values; real secrets are never committed |

## Current state (2026-11-25)

| Item | State |
|---|---|
| `services/core/` | Contains the existing skeleton (`/health`, `/api/hello`) and its tests, moved from `backend/` |
| `web/` | Contains the existing React app, moved from `frontend/` |
| `services/ticketing/`, `services/notification/` | Placeholders only: each file holds a comment describing its purpose. Not built by CI yet |
| `contracts/`, `deploy/docker-compose.prod.yml`, `deploy/.env.example` | Placeholders only |
| Images | Still **one image**: the root `Dockerfile` builds Core and the web app together, and CI/CD use it. The per-service `Dockerfile`s are placeholders until the images are split |
| `deploy/docker-compose.local.yml` | Not changed; it still points to the old `backend/` and `frontend/` folders |

## Next steps

1. Split the images: one Dockerfile per service plus the web app, and a CI/CD matrix.
2. Write the contracts (OpenAPI and event schemas) before implementing endpoints.
3. Add RabbitMQ and one database volume per service to the compose files.

## Open questions (team decisions)

1. **FR9 needs ticket holders inside Notification.** Notification cannot read Ticketing's
   database, so it would keep its own copy built from `TicketPurchased` / `TicketExpired`.
   That table is not in the 17-table model. Linked to IN-01, which is still open.
2. **`AGENTS.md` §3.2 lists REST only for Web App → Core / Ticketing,** but passengers manage
   subscriptions in Notification through the web app.
3. **Reliable event publishing.** If Core saves a change and crashes before publishing, the
   event is lost. The usual fix is an outbox table written in the same transaction, which is
   also not in the model.
4. **JWT verification** in Ticketing and Notification: shared secret or public key?
5. **Location of system / end-to-end tests,** which span several components.
6. **`trip_occupancy` ownership:** `Dados_Seed.md` says "Ticketing / Core", `AGENTS.md` says
   Ticketing only.
7. **Frontend serving and gateway:** the web app in its own nginx container that forwards
   `/api/...` to the services, or keep serving it from Core.
