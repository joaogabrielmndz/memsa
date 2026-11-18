# memsa (Memory + Sub-Agents)

Orquestrador CLI minimalista escrito em Go para executar fluxos multi-agentes usando o **OpenCode** com **memória persistente em disco** e **compressão extrema de prompts (estilo Caveman)**, reduzindo o consumo de tokens em mais de 65%.

---

## O Problema

Sessões longas de agentes de IA sofrem com degradação de contexto e custo exponencial: a cada nova interação, todo o histórico de mensagens e saídas de ferramentas anteriores é reenviado ao modelo. Além disso, prompts verbosos e explicações intermediárias consomem tokens desnecessários entre etapas técnicas.

## Como o `memsa` Resolve

O `memsa` substitui o histórico acumulado em memória RAM/chat por **passagem de estado via arquivos Markdown** e **isolamento de processos**:

1. **Zero-Context Loop (Isolamento por Tarefa):** Cada item do plano de execução roda em um processo separado (`opencode run`). Quando uma tarefa termina, o processo encerra e 100% da janela de contexto é descartada antes da próxima etapa.
2. **AI Memory em Disco (`.ai-memory/`):** O estado do projeto vive em arquivos enxutos no repositório (`PLAN.md` e `PROGRESS.md`), permitindo que cada novo agente leia apenas o estado atual sem carregar conversas passadas.
3. **Protocolo Caveman (`AGENTS.md`):** Sub-agentes operam sob regras estritas de comunicação telegráfica (sem artigos, saudações ou explicações redundantes) nas etapas internas de planejamento, execução e log.
4. **Agnóstico de Modelo:** Como delega a execução ao OpenCode CLI, herda automaticamente todos os provedores, chaves de API e modelos locais já configurados na máquina do usuário.

---

## Arquitetura do Projeto

```text
memsa/
├── main.go                      # Entrypoint da aplicação
├── cmd/                         # Definição dos comandos CLI (Cobra)
│   ├── root.go                  # Comando raiz e configuração global
│   ├── init.go                  # Inicializa .ai-memory/ e AGENTS.md
│   ├── plan.go                  # Gera tarefas atômicas no PLAN.md
│   └── run.go                   # Loop isolado de execução de tarefas
└── internal/
    ├── memory/
    │   └── manager.go           # Leitura, parse e atualização do estado em disco
    ├── runner/
    │   └── opencode.go          # Wrapper de execução headless do OpenCode CLI
    └── templates/
        ├── embed.go             # Empacotamento estático dos templates no binário
        ├── AGENTS.md            # Regras de compressão Caveman e protocolo de memória
        └── PLAN.md              # Estrutura inicial de checklist
```

## Fluxo de Execução
```text
[ Usuário ]
     │
     ▼
  memsa plan "objetivo" ──► OpenCode (Contexto Limpo) ──► Grava .ai-memory/PLAN.md
     │
     ▼
  memsa run ──► Loop para cada item '- [ ]' em PLAN.md:
                  1. Inicia processo isolado: opencode run
                  2. Injeta regras de AGENTS.md + histórico curto de PROGRESS.md
                  3. Executa tarefa única + valida build/testes
                  4. Atualiza PROGRESS.md (máx. 3 linhas) e encerra processo
                  5. Marca '- [x]' no PLAN.md e repete com contexto zerado
```

## Pre-réqusitos
- Go 1.23+ (apenas para compilar a partir do código-fonte)
- OpenCode CLI instalado e configurado no PATH (opencode)

## Instalação
```bash
git clone [https://github.com/seu-usuario/memsa.git](https://github.com/seu-usuario/memsa.git)
cd memsa
go mod tidy
go build -ldflags="-s -w" -o ~/.local/bin/memsa .
```

## Cross-compilation para distribuição
```bash
# Linux (64-bit)
GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o memsa-linux-amd64 .

# Windows (64-bit)
GOOS=windows GOARCH=amd64 go build -ldflags="-s -w" -o memsa-windows-amd64.exe .
```

## Uso

Dentro da raiz de qualquer projeto que você deseja automatizar:

1. **Inicializar o workspace de memória:**

```bash
memsa init
```
>Cria a pasta .ai-memory/ (PLAN.md, PROGRESS.md) e o arquivo AGENTS.md.

2. **Gerar o plano de tarefas atômicas:**

```Bash
memsa plan "implementar autenticação JWT e testes unitários"
```

3. **Executar as tarefas em ciclos isolados:**
```Bash
# Executa todas as tarefas pendentes
memsa run

# Ou executa apenas N tarefas por vez
memsa run --limit 1
```
---

## Comandos para Validar e Enviar ao Git

```bash
Rode na raiz de `~/P/memsa`:

go mod tidy
go build -o memsa .
./memsa --help
```
**Se estiver tudo certo:**

```Bash
git init
git add .
git commit -m "feat: initial skeleton with cobra cli, ai-memory and opencode runner"
```