# AI Logs — Registo Diário de Uso de IA

Evidências e métricas do uso de LLMs e agentes de IA ao longo do projeto
(template: `docs/planning/Guia_Reg_IA.md`, Parte 2). Os logs completos das sessões estão em
`convo_log_files/`.

## Campos

| Campo | Descrição |
|---|---|
| **ID Registo** | Identificador sequencial (`LOG-NNN`) |
| **Data** | Data da sessão (AAAA-MM-DD) |
| **Membro Equipa** | Quem conduziu/supervisionou o agente |
| **Fase SDLC** | Fase do ciclo de desenvolvimento |
| **Tarefa / Módulo** | Tarefa do plano ou módulo afetado |
| **Ferramenta IA** | Ferramenta e modelo usados |
| **Prompt / Contexto** | Prompt resumido e ficheiros de contexto fornecidos |
| **Autonomia (%)** | Percentagem do trabalho feito pelo agente |
| **Tempo Humano (min)** | Minutos de esforço humano (prompting, revisão, correção) |
| **Resultado / Erros** | Resultado obtido e erros detetados |
| **Evidência** | Commit SHA, PR, CI run ou ficheiro de log |

## Resumo

> **POR COMPLETAR** — registo criado retroativamente para commits que já existem mas não
> tinham entrada neste log. Os commits e ficheiros indicados são reais; os campos marcados
> `POR COMPLETAR` têm de ser preenchidos pelo membro que fez o trabalho. Se a tarefa não usou IA,
> preencher **Ferramenta IA** com `Nenhuma` e **Autonomia** com `0%`.

| ID | Data | Membro | Fase SDLC | Tarefa / Módulo | Ferramenta IA | Autonomia | Tempo Humano | Evidência |
|---|---|---|---|---|---|---:|---:|---|
| [LOG-001](#log-001) | 2026-10-10 | João Canais | Requisitos | Extração de ambiguidades do PDF | Gemini 3.5 Pro | 70% | 30 min | `04e7263` |
| [LOG-002](#log-002) | 2026-10-27 | Rodrigo R. | Requisitos | Revisão e validação com agente | Claude Code (Opus) | 85% | 45 min | `e3a3304` |
| [LOG-003](#log-003) | 2026-11-10 | João Dias | DevOps / CI | Pipeline de Build no GitHub Actions | Claude Code | 90% | 15 min | `7f8a12b` |
| [LOG-004](#log-004) | 2026-11-18 | Rafael Mota | Backend / Seed | Modelos Prisma de Paragens e Linhas | Cursor / Claude 3.5 | 80% | 20 min | `3a91c4d` |
| [LOG-005](#log-005) | 2026-11-24 | João Canais | Ticketing | Lógica de Validação com Idempotência | GitHub Copilot | 60% | 50 min | `9b20e11` |
| [LOG-006](#log-006) | 2026-11-13 | João Canais | Governance / IA | AGENTS.md com diretivas para agentes (Tarefa 9) | Claude Code (Opus 5.5) | 80% | 30 min | `ce85178` |
| [LOG-007](#log-007) | 2026-11-13 | João Canais | Backend / Seed | Seed Data: rede, horários, tarifário e ocupação (Tarefas 11-14) | Claude Code (Opus 5.5) | 60% | 60 min | `1d8e916` |
| [LOG-008](#log-008) | 2026-11-13 | João Canais | Processo / Git | Normas GitHub, GitFlow e template de PR | Claude Code (Opus 5.5) | 85% | 15 min | `d64c3d2` |
| [LOG-009](#log-009) | 2026-11-15 | Rafael Gomes | Governance / Planeamento | Preparação SDD das tarefas M2 de quality gates e testes | ChatGPT Work / Codex | 60% | 60 min | LOG-009 file |
| [LOG-010](#log-010) | 2026-11-15 | Rafael Gomes | DevOps / CI | Implementação dos quality gates da Task 3 / spec 001 | ChatGPT Work / Codex | 75% | 15 min | `b2c63b5`, PR #4 |
| [LOG-011](#log-011) | 2026-11-15 | Rafael Gomes | Testing / CI | Implementação dos testes iniciais da Task 4 / spec 002 | ChatGPT Work / Codex | 75% | 25 min | `5a20084`, PR #8 |
| [LOG-012](#log-012) | 2026-10-06 a 2026-10-13 | JLDuni / Rodrigo R. (por confirmar) | Setup / Estrutura | Estrutura inicial do repositório e sugestão de tech stack | POR COMPLETAR | POR COMPLETAR | POR COMPLETAR | `9aaeb4b`, `388ffaa`, `8223204`, `7656a97` |
| [LOG-013](#log-013) | 2026-10-27 | João Canais | Requisitos | Secção de participação da IA e decomposição do sistema em requirements.md | POR COMPLETAR | POR COMPLETAR | POR COMPLETAR | `242cf62` |
| [LOG-014](#log-014) | 2026-11-03 a 2026-11-10 | Rodrigo R. | DevOps / CI | Port de Python para Go/Gin + React/Vite, Dockerfile e CI Docker | POR COMPLETAR | POR COMPLETAR | POR COMPLETAR | `6fd5db1`, `a5bb38d`, `ce56adf`, `abc6b3f`; LOG-014 files |
| [LOG-015](#log-015) | 2026-11-10 | Rodrigo R. | Processo / Git | CODEOWNERS e conversão do plano Excel para Markdown | POR COMPLETAR | POR COMPLETAR | POR COMPLETAR | `f2e303f`, `f20691f` |
| [LOG-016](#log-016) | 2026-11-14 a 2026-11-16 | JLD43 (por confirmar) | DevOps / CD | Pipeline CD: build/push GHCR e deploy para dev | POR COMPLETAR | POR COMPLETAR | POR COMPLETAR | `8156f3d` … `710f902`, PR #9 |
| [LOG-017](#log-017) | 2026-11-17 | João Canais | Governance / IA | Regra AGENTS.md: seguir GITHUB_GUIDELINES e registar em AI_LOGS | POR COMPLETAR | POR COMPLETAR | POR COMPLETAR | `b738c52`, PR #15 |
| [LOG-018](#log-018) | 2026-11-24 | Sofia Saraiva | Documentação / Processo | Análise do estado do repo, registo retroativo de commits sem log e reorganização da documentação | Claude Code (Opus 5.5) | POR COMPLETAR | POR COMPLETAR | commit pendente |

---

## Registos detalhados

### LOG-001

- **Data:** 2026-10-10
- **Membro:** João Canais
- **Fase SDLC:** Requisitos
- **Tarefa / Módulo:** Extração de ambiguidades do PDF
- **Ferramenta IA:** Gemini 3.5 Pro
- **Prompt / Contexto:** Prompt: 'Analisa o enunciado e lista ambiguidades AM-XX'
- **Autonomia:** 70%
- **Tempo humano:** 30 min
- **Resultado / Erros:** Identificou 26 itens. 2 foram rejeitados por decisão da equipa.
- **Evidência:** commit `04e7263`

### LOG-002

- **Data:** 2026-10-27
- **Membro:** Rodrigo R.
- **Fase SDLC:** Requisitos
- **Tarefa / Módulo:** Revisão e validação com agente
- **Ferramenta IA:** Claude Code (Opus)
- **Prompt / Contexto:** Leitura do repositório para procurar lacunas de requisitos
- **Autonomia:** 85%
- **Tempo humano:** 45 min
- **Resultado / Erros:** Detetou falta de requisitos de CI/CD. Tentou alterar AM-08 (revertido).
- **Evidência:** commit `e3a3304`

### LOG-003

- **Data:** 2026-11-10
- **Membro:** João Dias
- **Fase SDLC:** DevOps / CI
- **Tarefa / Módulo:** Pipeline de Build no GitHub Actions
- **Ferramenta IA:** Claude Code
- **Prompt / Contexto:** Prompt: 'Cria workflow .github/workflows/ci.yml com build e linter'
- **Autonomia:** 90%
- **Tempo humano:** 15 min
- **Resultado / Erros:** Pipeline criada e funcional à 1.ª tentativa. Passou em 100% dos linters.
- **Evidência:** commit `7f8a12b`

### LOG-004

- **Data:** 2026-11-18
- **Membro:** Rafael Mota
- **Fase SDLC:** Backend / Seed
- **Tarefa / Módulo:** Modelos Prisma de Paragens e Linhas
- **Ferramenta IA:** Cursor / Claude 3.5
- **Prompt / Contexto:** Contexto: Schema das tabelas 'stops' e 'lines' do Excel v4
- **Autonomia:** 80%
- **Tempo humano:** 20 min
- **Resultado / Erros:** Gerou os modelos e migrations SQL. Corrigido tipo de dados de lat/long.
- **Evidência:** commit `3a91c4d`

### LOG-005

- **Data:** 2026-11-24
- **Membro:** João Canais
- **Fase SDLC:** Ticketing
- **Tarefa / Módulo:** Lógica de Validação com Idempotência
- **Ferramenta IA:** GitHub Copilot
- **Prompt / Contexto:** Implementar chave de idempotência no endpoint /tickets/validate
- **Autonomia:** 60%
- **Tempo humano:** 50 min
- **Resultado / Erros:** IA gerou validação simples. Esforço humano para garantir NFR3 (cache/redis).
- **Evidência:** commit `9b20e11`

### LOG-006

- **Data:** 2026-11-13
- **Membro:** João Canais
- **Fase SDLC:** Governance / IA
- **Tarefa / Módulo:** AGENTS.md com diretivas para agentes (Tarefa 9)
- **Ferramenta IA:** Claude Code (Opus 5.5)
- **Prompt / Contexto:** Contexto: requirements.md, Dados_Seed, Diretivas GitHub. Prompt: 'Cria AGENTS.md com regras de arquitetura, idempotência (NFR3) e schemas BD'
- **Autonomia:** 80%
- **Tempo humano:** 30 min
- **Resultado / Erros:** AGENTS.md antigo descrevia outro projeto (TaskFlow Pro); reescrito. Regra de IDs ajustada após revisão humana.
- **Evidência:** commit `ce85178`

### LOG-007

- **Data:** 2026-11-13
- **Membro:** João Canais
- **Fase SDLC:** Backend / Seed
- **Tarefa / Módulo:** Seed Data: rede, horários, tarifário e ocupação (Tarefas 11-14)
- **Ferramenta IA:** Claude Code (Opus 5.5)
- **Prompt / Contexto:** Contexto: Dados_Seed, mapa Rede Metrobus, requirements.md (AM-01..07). Prompt: 'Define os dados seed só com tabelas'
- **Autonomia:** 60%
- **Tempo humano:** 60 min
- **Resultado / Erros:** 1.ª versão tinha paragens inventadas e UUIDs desnecessários; refeita com o mapa Metrobus. Coordenadas GPS não obtidas (OSM indisponível). Utilizadores seed removidos. Detetou 4 inconsistências no Dados_Seed.
- **Evidência:** commit `1d8e916`

### LOG-008

- **Data:** 2026-11-13
- **Membro:** João Canais
- **Fase SDLC:** Processo / Git
- **Tarefa / Módulo:** Normas GitHub, GitFlow e template de PR
- **Ferramenta IA:** Claude Code (Opus 5.5)
- **Prompt / Contexto:** Contexto: Diretivas_GIT_HUB_FLOW, Guia_Reg_IA. Prompt: 'Cria GITHUB_GUIDELINES.md com as normas do Excel'
- **Autonomia:** 85%
- **Tempo humano:** 15 min
- **Resultado / Erros:** Criados GITHUB_GUIDELINES.md e pull_request_template.md. 1.º commit ficou com mensagem mal formatada; corrigido antes do push.
- **Evidência:** commit `d64c3d2`

### LOG-009

- **Data:** 2026-11-15
- **Membro:** Rafael Gomes
- **Fase SDLC:** Governance / Planeamento
- **Tarefa / Módulo:** Preparação SDD das tarefas M2 de quality gates e testes
- **Ferramenta IA:** ChatGPT Work / Codex
- **Prompt / Contexto:** Contexto: enunciado, Google Sheet, AGENTS.md, requirements.md, workflow CI e estado do repositório. Resumo: identificar responsabilidades e preparar o fluxo SDD, Git e registo de IA para as tarefas 3 e 4.
- **Autonomia:** 60%
- **Tempo humano:** 60 min
- **Resultado / Erros:** Tarefas 3 e 4 identificadas; plano e duas specs preparados. Triggers, requisitos, designs e tasks aprovados com revisão humana. Sonar foi adiado. Implementação pausada até indicação da branch e autorização explícita.
- **Evidência:** `convo_log_files/2026-11-15/LOG-009-governance-planning.md`; commit pendente

### LOG-010

- **Data:** 2026-11-15
- **Membro:** Rafael Gomes
- **Fase SDLC:** DevOps / CI
- **Tarefa / Módulo:** Implementação dos quality gates da Task 3 / spec 001
- **Ferramenta IA:** ChatGPT Work / Codex
- **Prompt / Contexto:** Contexto: spec 001 aprovada, workflow CI, backend Go e frontend React/Vite. Resumo: implementar e validar gates obrigatórios de formatação, análise estática e lint sem commit nem push.
- **Autonomia:** 75%
- **Tempo humano:** 15 min
- **Resultado / Erros:** ESLint, lockfile e job quality configurados; triggers atualizados; lint, build, gofmt e go vet passaram. Falha deliberada de lint foi detetada e revertida. CI do PR #4 passou nos jobs quality e build-and-verify.
- **Evidência:** Issue #2; commit `b2c63b5`; PR #4; CI run `43753488922`; `convo_log_files/2026-11-15/LOG-010-code-quality-implementation.md`

### LOG-011

- **Data:** 2026-11-15
- **Membro:** Rafael Gomes
- **Fase SDLC:** Testing / CI
- **Tarefa / Módulo:** Implementação dos testes iniciais da Task 4 / spec 002
- **Ferramenta IA:** ChatGPT Work / Codex
- **Prompt / Contexto:** Contexto: spec 002 aprovada, endpoints Go/Gin e workflow CI. Resumo: criar testes unitários para os endpoints atuais, adicionar um job de testes bloqueante e preservar os smoke tests sem guardar prompts textuais.
- **Autonomia:** 75%
- **Tempo humano:** 25 min
- **Resultado / Erros:** Testes de /health e /api/hello adicionados; falha deliberada detetada e revertida; gofmt, go vet, go test, build Docker e smoke tests passaram. CI do PR #8 passou nos jobs quality, tests e build-and-verify.
- **Evidência:** Issue #5; commit `5a20084`; PR #8; CI run `43903698598`; `convo_log_files/2026-11-15/LOG-011-initial-tests-implementation.md`

### LOG-012

> **POR COMPLETAR** — registo retroativo; preencher os campos marcados.

- **Data:** 2026-10-06 a 2026-10-13
- **Membro:** POR COMPLETAR (autores git: `JLDuni`, `JLD43`, `Rodrigo R`)
- **Fase SDLC:** Setup / Estrutura
- **Tarefa / Módulo:** Estrutura inicial do repositório (`src/`, `.ai/sdd`, `.ai/steering`, `.gitignore`) e sugestão de tech stack (`docs/teckStack.md`)
- **Ferramenta IA:** POR COMPLETAR
- **Prompt / Contexto:** POR COMPLETAR
- **Autonomia:** POR COMPLETAR
- **Tempo humano:** POR COMPLETAR
- **Resultado / Erros:** POR COMPLETAR
- **Evidência:** commits `9aaeb4b`, `388ffaa`, `8223204`, `7656a97`

### LOG-013

> **POR COMPLETAR** — registo retroativo; preencher os campos marcados.

- **Data:** 2026-10-27
- **Membro:** João Canais
- **Fase SDLC:** Requisitos
- **Tarefa / Módulo:** `docs/requirements.md` — secção 6 (decomposição do sistema) e secção 7 (evidência de uso de IA)
- **Ferramenta IA:** POR COMPLETAR
- **Prompt / Contexto:** POR COMPLETAR
- **Autonomia:** POR COMPLETAR
- **Tempo humano:** POR COMPLETAR
- **Resultado / Erros:** POR COMPLETAR
- **Evidência:** commit `242cf62`

### LOG-014

> **POR COMPLETAR** — registo retroativo; preencher os campos marcados.

- **Data:** 2026-11-03 a 2026-11-10
- **Membro:** Rodrigo R.
- **Fase SDLC:** DevOps / CI
- **Tarefa / Módulo:** Port do esqueleto da aplicação de Python/Flask para Go/Gin (backend) e React/Vite (frontend); Dockerfile multi-stage; workflow `ci-docker.yml`
- **Ferramenta IA:** POR COMPLETAR (as conversas guardadas indicam um assistente de chat; confirmar qual)
- **Prompt / Contexto:** POR COMPLETAR (conversa completa em `convo_log_files/2026-11-03/`)
- **Autonomia:** POR COMPLETAR
- **Tempo humano:** POR COMPLETAR
- **Resultado / Erros:** POR COMPLETAR
- **Evidência:** commits `6fd5db1`, `30d1d64`, `a5bb38d`, `ce56adf`, `9f1a0c2`, `abc6b3f`; `convo_log_files/2026-11-03/LOG-014-port-to-go.md`; `convo_log_files/2026-11-03/LOG-014-port-to-go-p2.md`

### LOG-015

> **POR COMPLETAR** — registo retroativo; preencher os campos marcados.

- **Data:** 2026-11-10
- **Membro:** Rodrigo R.
- **Fase SDLC:** Processo / Git
- **Tarefa / Módulo:** Ficheiro CODEOWNERS e conversão do plano Excel para Markdown (`docs/excels/*`)
- **Ferramenta IA:** POR COMPLETAR
- **Prompt / Contexto:** POR COMPLETAR
- **Autonomia:** POR COMPLETAR
- **Tempo humano:** POR COMPLETAR
- **Resultado / Erros:** POR COMPLETAR
- **Evidência:** commits `f2e303f`, `f20691f`, `d58a82e`, `3ab4e50`

### LOG-016

> **POR COMPLETAR** — registo retroativo; preencher os campos marcados.

- **Data:** 2026-11-14 a 2026-11-16
- **Membro:** POR COMPLETAR (autor git: `JLD43`)
- **Fase SDLC:** DevOps / CD
- **Tarefa / Módulo:** Pipeline CD (`cd.yml`): build e push da imagem para GHCR e deploy para o ambiente de desenvolvimento no self-hosted runner (`deploy/docker-compose.dev.yml`) (Tarefas 5–7)
- **Ferramenta IA:** POR COMPLETAR
- **Prompt / Contexto:** POR COMPLETAR
- **Autonomia:** POR COMPLETAR
- **Tempo humano:** POR COMPLETAR
- **Resultado / Erros:** POR COMPLETAR (o histórico mostra ~20 commits de correção — `fix`, `testing CD`, healthcheck removido e reposto, nome da imagem GHCR alterado — descrever as falhas e o retrabalho)
- **Evidência:** commits `8156f3d` a `710f902`; PR #9

### LOG-017

> **POR COMPLETAR** — registo retroativo; preencher os campos marcados.

- **Data:** 2026-11-17
- **Membro:** João Canais
- **Fase SDLC:** Governance / IA
- **Tarefa / Módulo:** AGENTS.md — tornar obrigatório seguir `GITHUB_GUIDELINES.md` e registar cada tarefa em `AI_LOGS.md`; criação de `docs/AI_LOGS.md`
- **Ferramenta IA:** POR COMPLETAR
- **Prompt / Contexto:** POR COMPLETAR
- **Autonomia:** POR COMPLETAR
- **Tempo humano:** POR COMPLETAR
- **Resultado / Erros:** POR COMPLETAR
- **Evidência:** commit `b738c52`; PR #15

### LOG-018

- **Data:** 2026-11-24 a 2026-11-25
- **Membro:** Sofia Saraiva
- **Fase SDLC:** Documentação / Processo
- **Tarefa / Módulo:** Análise do estado do repositório face ao enunciado; registo  dos commits sem entrada no AI_LOGS; reorganização de `convo_log_files/`; limpeza e estruturação da documentação (branch `feature/M3-structuring-docs`)
- **Ferramenta IA:** Claude
- **Prompt / Contexto:** Contexto: enunciado do projeto (PDF), AGENTS.md, histórico git, `docs/`, `convo_log_files/`. Resumo: avaliar onde está o projeto; marcar como "por completar" os commits feitos sem log; mover as conversas de `convo_log_files/3/11/` para a convenção `AAAA-MM-DD/LOG-NNN-*.md`; não alterar LOG-003 e LOG-005 (exemplos). Depois: renomear `docs/excels/` para `docs/planning/` com extensão `.md`, mover o CODEOWNERS, apagar duplicados, preencher os ficheiros vazios e atualizar os estados do plano, confirmando cada passo com a humana antes de mover ficheiros.
- **Autonomia:** POR COMPLETAR
- **Tempo humano:** POR COMPLETAR
- **Resultado / Erros:** Criadas entradas LOG-012 a LOG-017 (por completar) a partir do histórico git; conversas do port para Go movidas para `convo_log_files/2026-11-03/` e associadas a LOG-014. `docs/excels/` → `docs/planning/*.md` (referências atualizadas em AGENTS.md, GITHUB_GUIDELINES.md e AI_LOGS.md; registos históricos não alterados). CODEOWNERS movido para `.github/` (antes estava em `.github/workflows/`, onde o GitHub o ignora) e caminhos `/src/`, `/tests/` corrigidos para `/backend/`, `/frontend/`. Apagados `docs/teckStack.md` (duplicado), `.github/workflows/requirements.txt` e o bloco Python comentado do Dockerfile. Preenchidos `.ai/steering/*`, `docs/deliveries.md` e o template `000-feature-example` só com conteúdo já existente nos documentos do projeto. Estados de `Plano_Tarefas.md` e `Resumo_Projeto.md` atualizados com evidência.
- **Evidência:** commit pendente
