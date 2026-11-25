# Product

**MetroIsJustABus** is a public-transport platform for a Coimbra-inspired urban bus network.
It is the course project of Automated Software Engineering 2026/2027. Full requirements:
`docs/requirements.md`.

## Users

| User | What they do |
|---|---|
| **Visitor** (not signed in) | Consults the network, lines, stops, schedules and trips. Consultation is public (IN-02). |
| **Passenger** (`PASSENGER`) | Signs up and signs in (FR1), buys and views tickets (FR10–FR11), validates tickets (FR13), subscribes to lines, stops or trips (FR18), sets notification preferences (FR19) and receives notifications (FR9, FR20). |
| **Operational staff** (`STAFF`) | Maintains network and schedules, manages ticket types, registers and resolves disruptions, inspects ticket usage (FR14) and monitors occupancy (FR15). Staff accounts are created by seed/admin only (MI-05). |

## Main capabilities

| Area | Requirements |
|---|---|
| Network and schedule consultation | FR2–FR5 |
| Disruptions: delays, cancellations, diversions | FR5–FR8, AM-09–AM-11, MI-04 |
| Tickets: purchase, wallet, expiry, validation, inspection | FR10–FR14, AM-03–AM-05 |
| Occupancy and capacity | FR15–FR17, AM-06, AM-07 |
| Subscriptions and notifications | FR9, FR18–FR20, AM-12, AM-13, MI-06 |

## Components

From `docs/requirements.md` §6:

| Component | Responsibility |
|---|---|
| Web App | Passenger UI and staff back-office. Owns no data. |
| Core Backend | Auth (issues JWT), network, schedules, trip instances, disruptions. Publishes disruption events. |
| Ticketing & Occupancy | Ticket types, purchase, validation, occupancy. Publishes ticket events. |
| Notification | Subscriptions, preferences, matching affected passengers, deduplicated notifications. |
| Message broker | RabbitMQ transport for events. |

## Out of scope

- Refunds, transfers and ticket cancellation (MI-03).
- Real payments and real validator hardware: validators are simulated (AM-04).
- High load: demo scale; correctness under concurrency matters more than throughput (MI-08).
- Multiple languages: English only (MI-09).

## Open team decisions

- **IN-01**: what "active tickets matching the affected trip" means for FR9. Do not decide this in code.
