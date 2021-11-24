rollercoasters
==============

```sh
curl localhost:8080/coasters -i

curl localhost:8080/coasters -i -X POST -H 'Content-Type: application/json' --data-raw '{
  "name": "Name", "manufacturer": "Manu", "inPark": "Inpa", "height":55
}'

curl localhost:8080/coasters -i
```
