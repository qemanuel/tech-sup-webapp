#!/bin/bash

test_workers(){
    echo -e "\nGET Workers:"
    curl   "http://localhost:8010/api/v1/workers/"
    echo -e "\nGET Worker 1:"
    curl  "http://localhost:8010/api/v1/workers/1"
    echo -e "\nGET Worker 2:"
    curl  "http://localhost:8010/api/v1/workers/2"
    echo -e "\nDELETE Worker 1"
    curl   --request DELETE "http://localhost:8010/api/v1/workers/1"
    echo -e "\nGET Worker 1:\n"
    curl  "http://localhost:8010/api/v1/workers/1"
    echo -e "\nUPDATE Worker 2:"
    curl   --header "Content-Type: application/json" \
           --request POST \
           --data '{"name":"anibal","email":"anibal12345@gmail.com", "phone": "15123123123"}'\
           "http://localhost:8010/api/v1/workers/2"
    echo -e "\nGET Worker 2:"
    curl  "http://localhost:8010/api/v1/workers/2"
}

test_customers(){
    echo -e "\nGET Customers:"
    curl   "http://localhost:8010/api/v1/customers/"
    echo -e "\nGET Customer 1:"
    curl  "http://localhost:8010/api/v1/customers/1"
    echo -e "\nGET Customer 2:"
    curl  "http://localhost:8010/api/v1/customers/2"
    echo -e "\nDELETE Customer 1"
    curl   --request DELETE "http://localhost:8010/api/v1/customers/1"
    echo -e "\nGET Customer 1:"
    curl  "http://localhost:8010/api/v1/customers/1"
    echo -e "\nUPDATE Customer 2:"
    curl   --header "Content-Type: application/json" \
           --request POST \
           --data '{"name":"amazon","email":"amazon12345@gmail.com", "phone": "15123123123"}'\
           "http://localhost:8010/api/v1/customers/2"
    echo -e "\nGET Customer 2:"
    curl  "http://localhost:8010/api/v1/customers/2"
}

test_devices(){
    echo -e "\nGET Devices:"
    curl   "http://localhost:8010/api/v1/devices/"
    echo -e "\nGET Device 1:"
    curl  "http://localhost:8010/api/v1/devices/1"
    echo -e "\nGET Device 2:"
    curl  "http://localhost:8010/api/v1/devices/2"
    echo -e "\nDELETE Device 1"
    curl   --request DELETE "http://localhost:8010/api/v1/devices/1"
    echo -e "\nGET Device 1:"
    curl  "http://localhost:8010/api/v1/devices/1"
    echo -e "\nUPDATE Device 2:"
    curl   --header "Content-Type: application/json" \
           --request POST \
           --data '{"owner_id":"2","kind":"xbox","brand":"microsoft","model": "360", "serial":"123123"}'\
           "http://localhost:8010/api/v1/devices/2"
    echo -e "\nGET Device 2:"
    curl  "http://localhost:8010/api/v1/devices/2"
}
test_workers
test_customers
test_devices
