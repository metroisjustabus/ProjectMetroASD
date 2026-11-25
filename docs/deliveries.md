# Deliveries

What each milestone of the course project must deliver (from the project statement, v1.0,
section 5), and where the evidence is in this repository. Status as of 2026-11-25.

## Milestone 1 – Requirements (5%) – 27 Oct 2026 — Delivered

| Required | Status | Evidence |
|---|---|---|
| Main functional and non-functional requirements | Done | `docs/requirements.md` §2–3 (FR1–FR20, NFR1–NFR8) |
| Ambiguities, inconsistencies, missing information, assumptions | Done | `docs/requirements.md` §4 (AM, IN, MI) |
| Traceability to the project description | Done | `docs/requirements.md` §5 |
| Initial decomposition into components | Done | `docs/requirements.md` §6 |
| Evidence of LLM / agent use | Done | `docs/requirements.md` §7; `docs/AI_LOGS.md` LOG-001, LOG-002 |

## Milestone 2 – CI/CDE/CD pipeline (5%) – 17 Nov 2026 — Delivered

| Required | Status | Evidence |
|---|---|---|
| Automatic build from the repository | Done | `.github/workflows/ci-docker.yml` |
| Automated tests in the pipeline | Done | `tests` job; `services/core/main_test.go` (PR #8) |
| Basic code quality checks | Done | `quality` job: gofmt, go vet, ESLint (PR #4) |
| Containerized application | Done | `Dockerfile` (single image: Go backend + React build) |
| Pipeline produces a deployable version | Done | Image pushed to GHCR tagged with the commit SHA (`cd.yml`) |
| Pipeline deploys to the development environment | Done | `deploy-and-verify-dev` job on the self-hosted runner (PR #9) |

## Final delivery (90%) – 11 Dec 2026 — In progress

### Source code and engineering artifacts

| Required | Status |
|---|---|
| Source code, scripts, configuration, models, documentation | In progress: only the skeleton backend/frontend exists |
| Complete CI/CDE/CD and deployment configuration | In progress: production deploy is commented out in `cd.yml` |
| Final version deployed and running in production | Not started |

The system itself must also cover the technical description: web app, core backend,
independently deployable supporting service(s), notification across its own boundary,
sync + async communication, failure handling, idempotency, and separate dev/prod
environments with a deployment strategy and monitoring.

### Development evidence and metrics

| Required | Status |
|---|---|
| Evidence and metrics on LLM / agent use | In progress: `docs/AI_LOGS.md`, `convo_log_files/` |
| Records of agent activity, human intervention, validation, defects, rework, effort | In progress: several entries still marked POR COMPLETAR |

### Final report

Must cover: problem analysis and requirements; system design (domain/data models,
architecture, interfaces, communication); traceability; AI-assisted and agentic process;
testing and verification strategy; CI/CDE/CD process; deployment and operation; evidence,
metrics and critical analysis; team contributions. Status: not started.

### Video (10 min)

Must show: passenger-side functionality, interactions between components, normal and failure
scenarios, and proof that it is the deployed version. Status: not started.

### Self- and peer-evaluation

Form provided by the teachers. Status: not started.

### Defense

Sign up for a slot in Inforestudante before the final deadline. Every member must be able to
justify any part of the project.
