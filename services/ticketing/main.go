// Entry point of the Ticketing & Occupancy service: reads configuration from environment variables,
// opens its own SQLite database, runs migrations, registers routes, starts the ticket expiry job
// (FR12) and serves GET /health (NFR7). Not implemented yet.
package main
