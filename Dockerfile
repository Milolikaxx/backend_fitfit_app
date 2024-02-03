FROM golang

WORKDIR /app

ARG CACHEBUST=0

RUN git clone -b main https://Milolikaxx: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@github.com/Milolikaxx/backend_fitfit_app.git .

RUN go mod download

ENV GIN_MODE=release

COPY ./ ./

RUN go install github.com/cosmtrek/air@latest

# Don't forget to add .air.toml .gitignore
RUN air init

EXPOSE 8080

CMD ["air"]