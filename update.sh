cd byte-server
pm2 stop server.ts
git pull
bun install
pm2 start server.ts