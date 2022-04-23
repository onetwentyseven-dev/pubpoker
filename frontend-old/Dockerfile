FROM node:14 as builder

WORKDIR /app
ADD . /app/

RUN npm ci && npm run build

FROM alpine:latest AS release
WORKDIR /app

RUN apk --no-cache add tzdata ca-certificates

COPY --from=builder /app/.nuxt /app/.nuxt
WORKDIR /app

LABEL maintainer="David Douglas <david@onetwentyseven.dev>"

CMD ["npm", "run", "start"]