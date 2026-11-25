# Conventions

How work is organised in this repository. Binding rules are in `AGENTS.md`; Git details are
in `docs/GITHUB_GUIDELINES.md`.

## Repository layout

| Path | Content |
|---|---|
| `services/core/`, `services/ticketing/`, `services/notification/` | One Go/Gin service each, with its own `go.mod` (see `docs/architecture.md`) |
| `web/` | React/Vite web app |
| `contracts/` | OpenAPI and event schemas between components |
| `deploy/` | Docker Compose files per environment |
| `.github/workflows/` | CI and CD pipelines |
| `docs/requirements.md` | Requirements, assumptions, traceability, decomposition |
| `docs/architecture.md` | Repository structure and component layout |
| `docs/planning/` | Task plan, project summary, data model, AI-log template, Git directives |
| `docs/seed/` | Seed data definitions |
| `docs/AI_LOGS.md` | AI usage log (one entry per task) |
| `convo_log_files/AAAA-MM-DD/LOG-NNN-<name>.md` | Full session record for a log entry |
| `.ai/steering/` | Stable context for agents (this folder) |
| `.ai/sdd/specs/NNN-<name>/` | One spec per feature |

Do not create new top-level directories or services without asking (`AGENTS.md` §3).

## Spec-Driven Development

1. Copy `.ai/sdd/specs/000-feature-example/` to `.ai/sdd/specs/NNN-<name>/` (next free number).
2. Write `idea.md` → `plan.md` → `requirements.md` → `design.md` → `tasks.md`, in that order.
3. `.status` holds exactly one line and is the only source of truth for the spec state.
   The agent may set `*:draft`; only a human sets `*:approved` (`AGENTS.md` §7).
4. No production code before `tasks:approved`.

## Git

| Item | Format |
|---|---|
| Branch | `feature/M3-<task>`, `hotfix/<desc>`; never commit to `main` |
| Commit | `<type>(<scope>): <description>` + `Issue #<n>: [<FR/NFR>] <title>` + `Fixes #<n>` |
| Types | `feat`, `fix`, `docs`, `style`, `refactor`, `test`, `chore` |
| PR | To `develop`, using `.github/pull_request_template.md`; ≥ 1 human review and green CI |

## Code

- Cite requirement IDs (`FR13`, `NFR3`, `AM-05`) in comments on non-obvious logic.
- Go code must pass `gofmt` and `go vet`; frontend code must pass `npm run lint`.
- Tests live next to the code (`*_test.go` in Go).

## Data

- Table and column names follow `docs/planning/Dados_Seed.md` exactly (snake_case, plural tables).
- IDs are UUID strings; timestamps are UTC ISO-8601 text, shown in `Europe/Lisbon`.
- Money is integer cents (`price_cents`); "unlimited" is `NULL`, never `-1`.
- Enums are `TEXT` with a `CHECK (... IN (...))` constraint.
- Schema changes go through versioned migrations; never edit an applied migration.

## AI usage log

Every task adds an entry to `docs/AI_LOGS.md` (next `LOG-NNN`, in the summary table and the
detailed section), following `docs/planning/Guia_Reg_IA.md`. Only reference commits, PRs and CI
runs that exist. Long session records go in `convo_log_files/AAAA-MM-DD/LOG-NNN-<name>.md`.
