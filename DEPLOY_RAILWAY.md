# 🚀 Deploy Completo no Railway

## Estrutura do Projeto no Railway

Você terá **3 serviços** no mesmo projeto Railway:

1. **PostgreSQL** (Database)
2. **Backend** (API Go)
3. **Frontend** (SvelteKit)

---

## 📋 Passo a Passo

### 1️⃣ Criar Projeto no Railway

1. Acesse [Railway.app](https://railway.app)
2. Clique em **"New Project"**
3. Selecione **"Deploy from GitHub repo"**
4. Escolha seu repositório `manager`

---

### 2️⃣ Adicionar Banco de Dados PostgreSQL

1. No projeto criado, clique em **"New"**
2. Selecione **"Database"** → **"Add PostgreSQL"**
3. O Railway criará automaticamente um banco PostgreSQL
4. **Anote a URL de conexão** (você verá nas variáveis de ambiente)

---

### 3️⃣ Configurar Backend

1. No projeto, clique em **"New"** → **"GitHub Repo"**
2. Selecione o mesmo repositório `manager`

3. **Configurações do Backend:**
   - **Service Name**: `backend` (ou `api`)
   - **Root Directory**: Deixe **VAZIO** (raiz do projeto)
   - O Railway detectará automaticamente o `nixpacks.toml` ou `Dockerfile`

4. **Variáveis de Ambiente:**
   - Vá em **"Variables"** do serviço backend
   - Adicione:
     ```
     DATABASE_URL=<URL_DO_POSTGRESQL_DO_RAILWAY>
     PORT=8080
     GIN_MODE=release
     ```
   - **Importante**: Use a URL do PostgreSQL que o Railway criou (não a hardcoded)

5. **Atualizar código do backend** para usar variável de ambiente:
   - O arquivo `backend/internal/config/database.go` precisa ler `DATABASE_URL` do ambiente

---

### 4️⃣ Configurar Frontend

1. No mesmo projeto, clique em **"New"** → **"GitHub Repo"**
2. Selecione o mesmo repositório `manager`

3. **Configurações do Frontend:**
   - **Service Name**: `frontend` (ou `web`)
   - **Root Directory**: `frontend`
   - O Railway detectará automaticamente o `nixpacks.toml` do frontend

4. **Variáveis de Ambiente:**
   - Vá em **"Variables"** do serviço frontend
   - Adicione:
     ```
     PUBLIC_API_URL=https://seu-backend.railway.app/api
     PORT=5173
     ```
   - **Importante**: Substitua `seu-backend.railway.app` pela URL real do seu backend
   - Você encontra a URL do backend em: **Backend Service** → **Settings** → **Domains**

---

### 5️⃣ Executar Migrations

Após o backend estar rodando:

**Opção A: Via Railway CLI**
```bash
# Instalar Railway CLI
npm i -g @railway/cli

# Login
railway login

# Conectar ao projeto
railway link

# Executar migration
railway run --service backend sh -c "cd backend/cmd/migrate && go run main.go"
```

**Opção B: Via Terminal Web do Railway**
1. No dashboard do backend, vá em **"Deployments"**
2. Clique no deployment mais recente
3. Use o terminal web para executar:
   ```bash
   cd backend/cmd/migrate
   go run main.go
   ```

---

## 🔧 Ajustes Necessários no Código

### Backend - Usar variável de ambiente para DB

Atualize `backend/internal/config/database.go`:

```go
package config

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

// InitDB inicializa e retorna o pool de conexões do PostgreSQL
func InitDB() (*pgxpool.Pool, error) {
	// Usar variável de ambiente ou fallback para hardcoded
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		dbURL = "postgresql://postgres:fuTjLCHygdFJTpRARveKdgGtkwFOzpgc@mainline.proxy.rlwy.net:18337/railway"
	}

	config, err := pgxpool.ParseConfig(dbURL)
	if err != nil {
		return nil, fmt.Errorf("failed to parse database URL: %w", err)
	}

	config.MaxConns = 25
	config.MinConns = 5
	config.MaxConnLifetime = 5 * time.Minute
	config.MaxConnIdleTime = 1 * time.Minute

	pool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		return nil, fmt.Errorf("failed to create connection pool: %w", err)
	}

	// Testar conexão
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := pool.Ping(ctx); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	return pool, nil
}
```

---

## ✅ Verificação Final

### Testar Backend:
1. Acesse: `https://seu-backend.railway.app/api/dashboard/summary`
2. Deve retornar JSON (mesmo que vazio se não houver dados)

### Testar Frontend:
1. Acesse: `https://seu-frontend.railway.app`
2. A interface deve carregar
3. Verifique o console do navegador (F12) para erros

---

## 🌐 Configurar Domínios Customizados (Opcional)

1. **Backend**: Settings → Domains → Add Custom Domain
2. **Frontend**: Settings → Domains → Add Custom Domain
3. Atualize `PUBLIC_API_URL` no frontend com o novo domínio

---

## 📝 Checklist de Deploy

- [ ] Projeto criado no Railway
- [ ] PostgreSQL adicionado
- [ ] Backend configurado (Root Directory vazio)
- [ ] Variável `DATABASE_URL` configurada no backend
- [ ] Backend deployado com sucesso
- [ ] Migrations executadas
- [ ] Frontend configurado (Root Directory: `frontend`)
- [ ] Variável `PUBLIC_API_URL` configurada no frontend
- [ ] Frontend deployado com sucesso
- [ ] Testado acesso ao backend
- [ ] Testado acesso ao frontend
- [ ] Frontend consegue se comunicar com backend

---

## 🐛 Troubleshooting

**Erro: "Cannot connect to database"**
- Verifique se `DATABASE_URL` está correta
- Certifique-se de que o PostgreSQL está rodando

**Erro: "CORS" no frontend**
- O backend já tem CORS configurado
- Verifique se `PUBLIC_API_URL` está correto

**Erro: "404" nas requisições da API**
- Certifique-se que `PUBLIC_API_URL` termina com `/api`
- Exemplo correto: `https://backend.railway.app/api`

**Frontend não encontra a API**
- Verifique a variável `PUBLIC_API_URL` no Railway
- Verifique os logs do frontend no Railway
