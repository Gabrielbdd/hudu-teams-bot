# Hudu Teams Bot

Bot para Microsoft Teams que integra com o Hudu (plataforma de documentação de TI).

## Stack

- **Linguagem**: Go (gerenciado via mise)
- **Gerenciador de ferramentas**: [mise](https://mise.jdx.dev) — versões, tasks, variáveis de ambiente
- **API**: Hudu REST API v1 (JSON, autenticação via API key no header)

## Setup

```bash
# Instalar dependências (Go via mise)
mise install

# Copiar variáveis de ambiente
cp .env.example .env
# Editar .env com suas credenciais

# Ativar ambiente (automático se mise está configurado no shell)
mise trust
```

## Mise

Toda a configuração do projeto está em `mise.toml`. O mise gerencia:

- **Versão do Go**: definida em `[tools]`, pinada via `mise.lock` (commitado no repo)
- **Tasks**: build, test, lint, run, etc. — executar com `mise run <task>`
- **Variáveis de ambiente**: carregadas do `.env` via diretiva `_.file` no `[env]`

### Tasks disponíveis

```bash
mise run build      # Compila o binário
mise run test       # Roda os testes
mise run lint       # Roda o linter
mise run dev        # Roda em modo desenvolvimento
mise run fmt        # Formata o código
```

### Variáveis de ambiente

As variáveis ficam no arquivo `.env` (não commitado). O mise carrega automaticamente via `mise.toml`.
Veja `.env.example` para referência.

## Estrutura do projeto

```
hudu-teams-bot/
├── cmd/
│   └── bot/              # Entrypoint da aplicação
│       └── main.go
├── internal/
│   ├── config/           # Carregamento de configuração
│   ├── hudu/             # Cliente da API do Hudu
│   ├── teams/            # Integração com Microsoft Teams
│   └── bot/              # Lógica principal do bot
├── mise.toml             # Configuração do mise (tools, tasks, env)
├── mise.lock             # Lock file do mise (commitado)
├── .env                  # Variáveis de ambiente (NÃO commitado)
├── .env.example          # Exemplo de variáveis de ambiente
├── go.mod
└── go.sum
```

## Hudu API

- **Base URL**: `https://<instancia>.hudu.com/api/v1/`
- **Autenticação**: header `x-api-key` com a API key
- **Paginação**: query param `?page=N`, 25 resultados por página
- **Documentação**: disponível em cada instância em Admin > API Keys > View API Documentation (Swagger)

### Principais recursos da API

| Recurso | Descrição |
|---|---|
| Companies | Empresas/clientes — unidade organizacional principal |
| Assets | Ativos documentados (servidores, workstations, etc.) |
| Asset Layouts | Templates que definem a estrutura dos ativos |
| Asset Passwords | Credenciais armazenadas por ativo |
| Articles | Artigos da knowledge base |
| Procedures | SOPs/runbooks com checklists |
| Websites | Monitoramento de sites/domínios |
| Activity Logs | Logs de auditoria |
| Expirations | Itens expirando (domínios, certificados, garantias) |
| Magic Dash | Cards customizados no dashboard |

## Convenções de código

- Seguir as convenções idiomáticas de Go
- Usar `internal/` para pacotes privados
- Nomes de pacotes curtos e em minúsculas
- Tratar todos os erros explicitamente (não usar `_` para ignorar erros)
- Usar `context.Context` como primeiro parâmetro em funções que fazem I/O
- Testes no mesmo pacote com sufixo `_test.go`
- Usar `testify` para assertions nos testes quando necessário
