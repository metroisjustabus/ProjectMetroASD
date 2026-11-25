# Tech Stack

Stable technology choices for MetroIsJustABus. Changing any of these needs human approval
(`AGENTS.md` §2). Versions below are the ones currently in the repository.

| Area | Choice | Where it is defined |
|---|---|---|
| Backend | Go 1.23 + Gin 1.10 | `backend/go.mod`, `Dockerfile` |
| Frontend | React 18 + Vite 5 (Node 20) | `frontend/package.json`, `Dockerfile` |
| Authentication | JWT issued by the Core Backend; passwords hashed with BCrypt (cost ≥ 10) | `AGENTS.md` §2 |
| Database | SQLite, one database file per service | `AGENTS.md` §2–3 |
| Async messaging | RabbitMQ message broker | `docs/requirements.md` §6 |
| Containers | Docker, multi-stage build on Alpine | `Dockerfile` |
| Image registry | GitHub Container Registry (GHCR) | `.github/workflows/cd.yml` |
| CI/CD | GitHub Actions + self-hosted runner for the development environment | `.github/workflows/` |
| Code quality | `gofmt`, `go vet`, ESLint | `.github/workflows/ci-docker.yml` |
| Tests | `go test` | `.github/workflows/ci-docker.yml` |

## Not yet in the repository

These are decided but not implemented yet: SQLite databases, JWT/BCrypt auth, RabbitMQ,
and the separate Ticketing and Notification services.
