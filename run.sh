#!/bin/bash

go build -o bookings cmd/web/*.go
./bookings -production=false -cache=false -dbhost=localhost -dbport=5432 -dbname=bookings -dbuser=booking -dbpass=booking