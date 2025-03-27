##############
## FRONTEND ##
##############
FROM node:20-slim AS build-frontend-web

RUN corepack enable
WORKDIR /app

# deps
COPY frontend/web/package.json frontend/web/pnpm-lock.yaml ./
RUN --mount=type=cache,id=pnpm,target=/root/.local/share/pnpm/store pnpm fetch --frozen-lockfile
RUN --mount=type=cache,id=pnpm,target=/root/.local/share/pnpm/store pnpm install --frozen-lockfile

# gen pwa assets
COPY frontend/web/pwa-assets.config.ts .
COPY frontend/web/src/consts.ts src/consts.ts
COPY frontend/web/public/logo.svg public/logo.svg
RUN pnpm gen:pwa-assets

COPY frontend/web .
ARG VERSION=UNKNOWN
ENV VERSION=$VERSION
RUN pnpm build
# copy apple-touch-icon.png becuase reading the correct filename seems to be hard
RUN cp dist/apple-touch-icon-180x180.png dist/apple-touch-icon.png


#############
## BACKEND ##
#############
FROM golang:1.24-alpine AS build-backend

WORKDIR /app

COPY backend/go.mod backend/go.sum ./
RUN go mod download && go mod verify

COPY backend .
RUN go build -v -o ./adnexos cmd/adnexos/main.go


#########
## RUN ##
#########
FROM alpine

WORKDIR /app

COPY --from=build-frontend-web /app/dist pb_public/
COPY --from=build-backend /app/adnexos adnexos

EXPOSE 8090
CMD ["./adnexos", "serve", "--http=0.0.0.0:8090"]
