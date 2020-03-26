# InvitesSystem

This system is responsible for sending invites to list of customers
who are within 100km of given office location for some food and drinks on us. 

This program that will read the full list of customers and
output the names and user ids of matching customers (within 100km),
sorted by User ID (ascending).

# # How to run this code

1. Have the mongodb instance running in background (docker exec -it mongodb bash)
2. go run *.go
3. curl -d '{"url": "https://s3.amazonaws.com/intercom-take-home-test/customers.txt"}' -H "Content-Type:application/json" --request POST 'http://localhost:8001/customers'
