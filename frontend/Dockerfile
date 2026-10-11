# Build stage
FROM node:24-alpine AS builder

WORKDIR /app

# Install dependencies first for better layer caching
COPY package.json package-lock.json ./
RUN npm ci

# Copy source and build static assets
COPY . .
RUN npm run build

# Final stage
FROM nginx:1.29-alpine

# Serve the built SPA, falling back to index.html for client-side routes
COPY nginx.conf /etc/nginx/conf.d/default.conf
COPY --from=builder /app/dist /usr/share/nginx/html

EXPOSE 80

CMD ["nginx", "-g", "daemon off;"]
