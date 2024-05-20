FROM node:lts-bullseye as build-stage

WORKDIR /app

# Копируем package.json и package-lock.json для установки зависимостей
COPY package*.json ./

# Устанавливаем необходимые пакеты для сборки node-sass и других зависимостей
RUN apt-get update && apt-get install -y python3 make g++ && rm -rf /var/lib/apt/lists/*

# Устанавливаем зависимости проекта
RUN npm install

# Копируем остальные файлы проекта и собираем приложение
COPY . .
RUN npm run build

# Используем nginx для сервировки статических файлов
FROM nginx:stable-alpine as production-stage
COPY --from=build-stage /app/dist /usr/share/nginx/html
COPY nginx.conf /etc/nginx/nginx.conf

EXPOSE 3000
CMD ["nginx", "-g", "daemon off;"]