# Build stage
FROM node:20-alpine AS builder

WORKDIR /app

# Copy package files
COPY package*.json ./
RUN npm install --legacy-peer-deps

# Copy source code
COPY . .

# Run svelte-kit sync to generate .svelte-kit directory
RUN npx svelte-kit sync || true

# Build the application
RUN npm run build

# Production stage
FROM node:20-alpine

WORKDIR /app

# Copy package files
COPY package*.json ./

# Install only production dependencies
RUN npm install --omit=dev --legacy-peer-deps

# Copy built application from builder
COPY --from=builder /app/build ./build
COPY --from=builder /app/package.json ./

# Expose port (Railway define automaticamente)
EXPOSE 3000

# Start the application
ENV PORT=3000
ENV HOST=0.0.0.0
CMD ["node", "build"]
