# Backend - Sistema de Controle Financeiro

## Estrutura do Projeto

```
backend/
├── cmd/
│   ├── server/          # Servidor principal
│   └── migrate/          # Script de migrations
├── internal/
│   ├── config/          # Configurações (DB, etc)
│   ├── entity/          # Entidades do domínio
│   ├── handlers/        # Controllers HTTP
│   ├── repositories/    # Acesso a dados
│   ├── routes/          # Definição de rotas
│   └── usecases/        # Lógica de negócio
└── migrations/           # SQL migrations
```

## Configuração

1. Instalar dependências:
```bash
go mod download
```

2. Executar migrations:
```bash
cd cmd/migrate
go run main.go
```

3. Executar servidor:
```bash
cd cmd/server
go run main.go
```

O servidor estará rodando em `http://localhost:8080`

## API Endpoints

### Transações
- `GET /api/transactions` - Listar todas
- `GET /api/transactions/:id` - Buscar por ID
- `POST /api/transactions` - Criar
- `PUT /api/transactions/:id` - Atualizar
- `DELETE /api/transactions/:id` - Excluir
- `POST /api/transactions/:id/installments/:installment/pay` - Pagar parcela

### Categorias
- `GET /api/categories` - Listar todas
- `GET /api/categories/:id` - Buscar por ID
- `POST /api/categories` - Criar
- `PUT /api/categories/:id` - Atualizar
- `DELETE /api/categories/:id` - Excluir

### Dashboard
- `GET /api/dashboard/summary` - Resumo financeiro

