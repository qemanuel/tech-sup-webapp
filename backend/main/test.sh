#!/bin/bash

test_workers(){
    echo "GET Workers:"
    curl   "http://localhost:8010/api/v1/workers/"
    echo "GET Worker 1:"
    curl  "http://localhost:8010/api/v1/workers/1"
    echo "GET Worker 2:"
    curl  "http://localhost:8010/api/v1/workers/2"
    echo "DELETE Worker 1"
    curl   --request DELETE "http://localhost:8010/api/v1/workers/1"
    echo "GET Worker 1:"
    curl  "http://localhost:8010/api/v1/workers/1"
    echo "UPDATE Worker 2:"
    curl   --header "Content-Type: application/json" \
           --request POST \
           --data '{"name":"anibal","email":"anibal12345@gmail.com", "phone": "15123123123"}'\
           "http://localhost:8010/api/v1/workers/2"
    echo "GET Worker 2:"
    curl  "http://localhost:8010/api/v1/workers/2"
}

test_customers(){
    echo "GET Customers:"
    curl   "http://localhost:8010/api/v1/customers/"
    echo "GET Customer 1:"
    curl  "http://localhost:8010/api/v1/customers/1"
    echo "GET Customer 2:"
    curl  "http://localhost:8010/api/v1/customers/2"
    echo "DELETE Customer 1"
    curl   --request DELETE "http://localhost:8010/api/v1/customers/1"
    echo "GET Customer 1:"
    curl  "http://localhost:8010/api/v1/customers/1"
    echo "UPDATE Customer 2:"
    curl   --header "Content-Type: application/json" \
           --request POST \
           --data '{"name":"amazon","email":"amazon12345@gmail.com", "phone": "15123123123"}'\
           "http://localhost:8010/api/v1/customers/2"
    echo "GET Customer 2:"
    curl  "http://localhost:8010/api/v1/customers/2"
}

test_devices(){
    echo "GET Devices:"
    curl   "http://localhost:8010/api/v1/devices/"
    echo "GET Device 1:"
    curl  "http://localhost:8010/api/v1/devices/1"
    echo "GET Device 2:"
    curl  "http://localhost:8010/api/v1/devices/2"
    echo "DELETE Device 1"
    curl   --request DELETE "http://localhost:8010/api/v1/devices/1"
    echo "GET Device 1:"
    curl  "http://localhost:8010/api/v1/devices/1"
    echo "UPDATE Device 2:"
    curl   --header "Content-Type: application/json" \
           --request POST \
           --data '{"owner_id":"2","kind":"xbox","brand":"microsoft","model": "360", "serial":"123123"}'\
           "http://localhost:8010/api/v1/devices/2"
    echo "GET Device 2:"
    curl  "http://localhost:8010/api/v1/devices/2"
}
test_workers
test_customers
test_devices
