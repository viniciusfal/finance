# Guia de Deploy - Railway

## Deploy do Backend e Frontend no Railway

### 1. Backend (Primeiro serviço)

1. **Criar novo serviço no Railway**
   - Clique em "New Project" → "GitHub Repo"
   - Selecione seu repositório

2. **Configurar Backend**
   - **Service Name**: `backend` ou `financy-api`
   - **Root Directory**: Deixe vazio (raiz do projeto)
   - O Railway detectará automaticamente o `nixpacks.toml` ou `Dockerfile`

3. **Variáveis de Ambiente** (se necessário):
   - `PORT`: `8080` (Railway define automaticamente)
   - `GIN_MODE`: `release`

4. **Banco de Dados PostgreSQL**
   - Adicione um serviço PostgreSQL no Railway
   - O backend já está configurado com a URL do banco

5. **Executar Migrations**
   - Após o deploy, execute as migrations manualmente ou crie um script
   - Você pode usar o Railway CLI ou executar via terminal

### 2. Frontend (Segundo serviço)

1. **Criar novo serviço no Railway**
   - No mesmo projeto, clique em "New" → "GitHub Repo"
   - Selecione o mesmo repositório

2. **Configurar Frontend**
   - **Service Name**: `frontend` ou `financy-web`
   - **Root Directory**: `frontend`
   - O Railway detectará automaticamente o `nixpacks.toml` do frontend

3. **Variáveis de Ambiente IMPORTANTES**:
   - `PUBLIC_API_URL`: URL do seu backend
     - Exemplo: `https://seu-backend.railway.app/api`
     - Você encontra a URL no dashboard do serviço backend no Railway

4. **Porta**
   - O Railway define automaticamente a porta
   - O SvelteKit usa a variável `PORT` automaticamente

### 3. Estrutura no Railway

```
Projeto: Financy
├── PostgreSQL (Database)
├── Backend Service
│   └── Root: / (raiz)
│   └── URL: https://backend-xxxx.railway.app
└── Frontend Service
    └── Root: /frontend
    └── URL: https://frontend-xxxx.railway.app
    └── Env: PUBLIC_API_URL=https://backend-xxxx.railway.app/api
```

### 4. Configurar Domínios (Opcional)

1. **Backend**: Configure um domínio customizado se quiser
2. **Frontend**: Configure um domínio customizado
3. Atualize `PUBLIC_API_URL` no frontend com o novo domínio do backend

### 5. Executar Migrations

Após o deploy do backend, você precisa executar as migrations:

**Opção A: Via Railway CLI**
```bash
railway run --service backend cd backend/cmd/migrate && go run main.go
```

**Opção B: Via Terminal do Railway**
- No dashboard do backend, vá em "Deployments" → "View Logs"
- Ou use o terminal web do Railway

**Opção C: Criar endpoint de migration** (não recomendado para produção)
- Adicione um endpoint temporário no backend para executar migrations

### 6. Verificar Deploy

1. **Backend**: Acesse `https://seu-backend.railway.app/api/dashboard/summary`
   - Deve retornar JSON com os dados do dashboard

2. **Frontend**: Acesse `https://seu-frontend.railway.app`
   - Deve carregar a interface
   - Verifique no console do navegador se há erros de CORS ou API

### 7. Troubleshooting

**Erro de CORS no Frontend:**
- O backend já tem CORS configurado para aceitar todas as origens
- Se ainda houver problema, verifique se `PUBLIC_API_URL` está correto

**Erro 404 na API:**
- Verifique se `PUBLIC_API_URL` termina com `/api`
- Exemplo correto: `https://backend.railway.app/api`
- Exemplo errado: `https://backend.railway.app` (falta `/api`)

**Migrations não executadas:**
- Execute manualmente via Railway CLI ou terminal
- Verifique se o banco de dados está conectado corretamente

### 8. Atualizar API para Produção

O arquivo `frontend/src/lib/api/api.js` já está configurado para usar `PUBLIC_API_URL` se disponível.

Certifique-se de definir a variável de ambiente `PUBLIC_API_URL` no serviço do frontend no Railway.
