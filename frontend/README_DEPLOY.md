# Deploy do Frontend

## Opção 1: Railway (Recomendado se já usa Railway)

1. No Railway, crie um novo serviço
2. Conecte o repositório
3. Configure:
   - **Root Directory**: `frontend`
   - O Railway detectará automaticamente o `nixpacks.toml`
4. Adicione variável de ambiente:
   - `PUBLIC_API_URL`: URL do seu backend (ex: `https://seu-backend.railway.app`)

## Opção 2: Vercel (Mais fácil para SvelteKit)

1. Instale a CLI do Vercel: `npm i -g vercel`
2. No diretório `frontend`, execute: `vercel`
3. Ou conecte o repositório no dashboard da Vercel
4. Configure:
   - **Framework Preset**: SvelteKit
   - **Root Directory**: `frontend`
   - **Build Command**: `npm run build`
   - **Output Directory**: `.svelte-kit`

### Variáveis de Ambiente na Vercel:
- `PUBLIC_API_URL`: URL do seu backend

## Opção 3: Netlify

1. Conecte o repositório no Netlify
2. Configure:
   - **Base directory**: `frontend`
   - **Build command**: `npm run build`
   - **Publish directory**: `.svelte-kit`

## Configuração da API

No arquivo `frontend/src/lib/api/api.js`, a URL da API está configurada como `/api` (proxy).

Para produção, você precisa:

1. **Opção A**: Configurar proxy no servidor (Vercel/Netlify fazem isso automaticamente)
2. **Opção B**: Usar variável de ambiente `PUBLIC_API_URL`

### Atualizar api.js para usar variável de ambiente:

```javascript
const API_BASE = import.meta.env.PUBLIC_API_URL || '/api';
```

