# Copyright 2016 The Kubernetes Authors.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

FROM golang:1.19.6-alpine

WORKDIR /app
COPY go.mod go.sum main.go ./
RUN go mod download && go mod verify

ADD ./main.go .
#RUN go build -v -o main .
#RUN CGO_ENABLED=0 GOOS=linux go build -v -o main .
RUN CGO_ENABLED=0 go build -ldflags="-s -w" -o main .
FROM scratch
#FROM golang:1.19.6-alpine

WORKDIR /app
COPY --chown=nonroot:nonroot --from=0 /app/main .
COPY --chown=nonroot:nonroot ./public/index.html public/index.html
COPY --chown=nonroot:nonroot ./public/script.js public/script.js
COPY --chown=nonroot:nonroot ./public/style.css public/style.css

USER 5000:5000

CMD ["/app/main"]
EXPOSE 3000
