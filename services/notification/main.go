// Entry point of the Notification service: reads configuration from environment variables, opens
// its own SQLite database, runs migrations, starts the RabbitMQ consumer, registers routes and
// serves GET /health (NFR7). It must start even when the broker is down (NFR1). Not implemented yet.
package main
