#!/bin/bash

fill_workers(){ 
    curl   --header "Content-Type: application/json" \
           --request POST \
           --data '{"name":"esteban","email":"esteban@gmail.com", "phone": "1511111111"}'\
           "http://localhost:8010/api/v1/workers/"
    curl   --header "Content-Type: application/json" \
           --request POST \
           --data '{"name":"anibal","email":"anibal@gmail.com", "phone": "1522222222"}'\
           "http://localhost:8010/api/v1/workers/"
    curl   --header "Content-Type: application/json" \
           --request POST \
           --data '{"name":"ramon","email":"ramon@gmail.com", "phone": "1533333333"}'\
           "http://localhost:8010/api/v1/workers/"
    curl   --header "Content-Type: application/json" \
           --request POST \
           --data '{"name":"hilda","email":"hilda@gmail.com", "phone": "1544444444"}'\
           "http://localhost:8010/api/v1/workers/"
    curl   --header "Content-Type: application/json" \
           --request POST \
           --data '{"name":"susana","email":"susana@gmail.com", "phone": "1555555555"}'\
           "http://localhost:8010/api/v1/workers/"
}

fill_customers(){ 
    curl   --header "Content-Type: application/json" \
           --request POST \
           --data '{"name":"google","email":"google@gmail.com", "phone": "1511111111"}'\
           "http://localhost:8010/api/v1/customers/"
    curl   --header "Content-Type: application/json" \
           --request POST \
           --data '{"name":"amazon","email":"amazon@gmail.com", "phone": "1522222222"}'\
           "http://localhost:8010/api/v1/customers/"
    curl   --header "Content-Type: application/json" \
           --request POST \
           --data '{"name":"microsoft","email":"microsoft@gmail.com", "phone": "1533333333"}'\
           "http://localhost:8010/api/v1/customers/"
    curl   --header "Content-Type: application/json" \
           --request POST \
           --data '{"name":"nike","email":"nike@gmail.com", "phone": "1544444444"}'\
           "http://localhost:8010/api/v1/customers/"
    curl   --header "Content-Type: application/json" \
           --request POST \
           --data '{"name":"adidas","email":"adidas@gmail.com", "phone": "1555555555"}'\
           "http://localhost:8010/api/v1/customers/"
}

fill_devices(){ 
    curl   --header "Content-Type: application/json" \
           --request POST \
           --data '{"brand": "sony","kind": "ps","model": "ps5","owner_id": "1","serial": "111"}'\
           "http://localhost:8010/api/v1/devices/"
    curl   --header "Content-Type: application/json" \
           --request POST \
           --data '{"brand": "microsoft","kind": "xbox","model": "one","owner_id": "2","serial": "222"}'\
           "http://localhost:8010/api/v1/devices/"
    curl   --header "Content-Type: application/json" \
           --request POST \
           --data '{"brand": "sony","kind": "tv","model": "hjkg1234","owner_id": "3","serial": "333"}'\
           "http://localhost:8010/api/v1/devices/"
    curl   --header "Content-Type: application/json" \
           --request POST \
           --data '{"brand": "apple","kind": "cellphone","model": "iphone15","owner_id": "4","serial": "444"}'\
           "http://localhost:8010/api/v1/devices/"
    curl   --header "Content-Type: application/json" \
           --request POST \
           --data '{"brand": "dell","kind": "notebook","model": "ultrabook","owner_id": "5","serial": "555"}'\
           "http://localhost:8010/api/v1/devices/"
}

fill_jobs(){
       curl   --header "Content-Type: application/json" \
              --request POST \
              --data '{"device_id":"1",
                     "reason":"falla",
                     "observations": "",
                     "status": "ingressed",
                     "assigned_id": "1",
                     "author_id": "1"}'\
              "http://localhost:8010/api/v1/jobs/"
       curl   --header "Content-Type: application/json" \
              --request POST \
              --data '{"device_id":"2",
                     "reason":"falla",
                     "observations": "a",
                     "status": "ingressed",
                     "assigned_id": "2",
                     "author_id": "2"}'\
              "http://localhost:8010/api/v1/jobs/"
       curl   --header "Content-Type: application/json" \
              --request POST \
              --data '{"device_id":"3",
                     "reason":"falla",
                     "observations": "b",
                     "status": "ingressed",
                     "assigned_id": "1",
                     "author_id": "1"}'\
              "http://localhost:8010/api/v1/jobs/"
       curl   --header "Content-Type: application/json" \
              --request POST \
              --data '{"device_id":"4",
                     "reason":"falla",
                     "observations": "c",
                     "status": "ingressed",
                     "assigned_id": "2",
                     "author_id": "2"}'\
              "http://localhost:8010/api/v1/jobs/"
}

fill_workers
fill_customers
fill_devices
fill_jobs

# curl   'http://localhost:8010/api/v1/jobs/?deviceId="1"&authorId="1"&assignedId="1"'
#
#
#