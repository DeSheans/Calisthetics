FROM node:23.4 AS dependencies

RUN npm install -g pnpm
WORKDIR /calisthetic
COPY package.json pnpm-lock.yaml ./
RUN pnpm install

FROM node:23.4 AS builder
WORKDIR /calisthetic
COPY . .
COPY --from=dependencies /calisthetic/node_modules ./node_modules
# RUN pnpm run build

FROM node:23.4 AS runner

RUN npm install -g pnpm

WORKDIR /calisthetic

COPY --from=builder /calisthetic/src ./src
COPY --from=builder /calisthetic/public ./public
COPY --from=builder /calisthetic/package.json ./package.json
COPY --from=builder /calisthetic/tsconfig.json ./tsconfig.json
COPY --from=builder /calisthetic/.next ./.next
COPY --from=builder /calisthetic/node_modules ./node_modules

CMD ["pnpm", "dev"]