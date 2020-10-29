FROM golang:1.15-alpine AS build
WORKDIR /app/out
COPY . .

# Spiegazione CGO_ENABLED=0 : https://stackoverflow.com/questions/55106186/no-such-file-or-directory-with-docker-scratch-image
RUN CGO_ENABLED=0 go build -o ./auth main.go

FROM scratch AS bin
EXPOSE 8200
COPY --from=build /app/out/auth /app

CMD [ "/app" ]