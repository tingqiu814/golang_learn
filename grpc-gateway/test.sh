#!/bin/bash
curl -X POST -k "http://localhost:8080/v1/example/echo" -d '{"value": "CoreOS is hiring!"}'
curl -vX POST -k "http://localhost:8080/v1/example/addItem" -d '{"userId":2,"itemType":3,"itemQuantity":3,"orderId":"123","source":12323,"prptys":[{"itemId":"qwer","id":12,"value":123,"info":"info"}],"itemName":"12321"}'
