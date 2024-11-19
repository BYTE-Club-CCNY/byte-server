FROM node:18-alpine AS base

WORKDIR /app

COPY package.json package-lock.json ./

RUN npm install --no-cache

COPY . .

EXPOSE 3000

CMD ["npm", "run", "serve"]
