# Hudu Teams Bot

Bot para Microsoft Teams que integra com o [Hudu](https://hudu.com).

## Pré-requisitos

- [mise](https://mise.jdx.dev) instalado
- [ngrok](https://ngrok.com) instalado (para desenvolvimento local)
- Conta no [Azure Portal](https://portal.azure.com)

## 1. Setup local

```bash
git clone git@github.com:Gabrielbdd/hudu-teams-bot.git
cd hudu-teams-bot
mise trust && mise install
cp .env.example .env
```

## 2. Registrar o bot no Azure

### 2.1 App Registration (Entra ID)

1. Acesse [Azure Portal > App registrations](https://portal.azure.com/#view/Microsoft_AAD_RegisteredApps/ApplicationsListBlade)
2. **New registration**
3. Nome: `Hudu Bot`
4. Supported account types: **Accounts in any organizational directory (Multitenant)**
5. Clique **Register**
6. Copie o **Application (client) ID** → coloque em `TEAMS_APP_ID` no `.env`
7. Vá em **Certificates & secrets** > **New client secret**
8. Copie o **Value** do secret → coloque em `TEAMS_APP_PASSWORD` no `.env`

### 2.2 Azure Bot Resource

1. Acesse [Azure Portal > Create a resource](https://portal.azure.com/#create/hub) > busque **Azure Bot**
2. **Bot handle**: `hudu-bot`
3. **Type of App**: Multi Tenant
4. **App ID**: use o Application (client) ID do passo anterior (selecione "Use existing app registration")
5. Clique **Create**
6. Após criar, vá em **Configuration** > **Messaging endpoint**: deixe em branco por enquanto (vamos preencher com a URL do ngrok)
7. Vá em **Channels** > adicione **Microsoft Teams**

## 3. Rodar o bot

```bash
# Terminal 1: inicia o bot
mise run dev

# Terminal 2: expõe via ngrok
mise run ngrok
```

O ngrok vai mostrar uma URL como `https://abcd1234.ngrok-free.app`.

Copie essa URL e atualize o **Messaging endpoint** no Azure Bot:

```
https://abcd1234.ngrok-free.app/api/messages
```

## 4. Instalar no Teams

### 4.1 Empacotar o manifest

```bash
mise run package
```

Isso gera `dist/hudu-bot.zip`.

### 4.2 Sideload no Teams

1. Abra o **Microsoft Teams**
2. Vá em **Apps** > **Manage your apps** > **Upload an app**
3. Selecione **Upload a custom app**
4. Escolha o arquivo `dist/hudu-bot.zip`
5. Clique **Add**

O bot vai aparecer no chat. Mande uma mensagem e ele responde com "Hello! You said: ...".

## Tasks disponíveis

| Task | Comando | Descrição |
|---|---|---|
| dev | `mise run dev` | Roda o bot |
| build | `mise run build` | Compila o binário |
| test | `mise run test` | Roda os testes |
| lint | `mise run lint` | Roda o linter |
| fmt | `mise run fmt` | Formata o código |
| ngrok | `mise run ngrok` | Expõe o bot via ngrok |
| package | `mise run package` | Empacota o manifest do Teams |

## Variáveis de ambiente

| Variável | Descrição |
|---|---|
| `TEAMS_APP_ID` | Application (client) ID do App Registration |
| `TEAMS_APP_PASSWORD` | Client secret do App Registration |
| `TEAMS_TENANT_ID` | Tenant ID (opcional para multi-tenant) |
| `HUDU_BASE_URL` | URL da instância do Hudu |
| `HUDU_API_KEY` | API key do Hudu |
| `APP_PORT` | Porta do servidor (default: 3978) |
| `APP_ENV` | Ambiente (development/production) |
