// Package store is the Core service SQLite access layer. It owns users, stops, lines, line_stops,
// day_types, day_type_exceptions, trip_templates, trip_instances, stop_schedules and disruptions.
// No other service may read or write these tables (AGENTS.md section 3).
package store
