# Stage 1: Build
FROM node:20-alpine AS builder

WORKDIR /app
RUN apk add --no-cache bash

# Копируем package.json и package-lock.json
COPY package.json package-lock.json tsconfig.json ./

ENV NODE_ENV=development
RUN npm install --legacy-peer-deps

# Копируем исходники и public
COPY ./src ./src
COPY ./public ./public

# Сборка production
RUN npx react-scripts build

# Stage 2: Production
FROM nginx:alpine
COPY --from=builder /app/build /usr/share/nginx/html
EXPOSE 3000
CMD ["nginx", "-g", "daemon off;"]