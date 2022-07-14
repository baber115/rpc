package client

/**
curl --location --request POST 'http://localhost:1234/jsonrpc' \
--header 'Content-Type: application/json' \
--data-raw '{
    "method": "HelloServer.Hello",
    "params": [
        "alice01"
    ]
}'
*/

/**
curl --location --request POST 'http://localhost:1234/jsonrpc' \
--header 'Content-Type: application/json' \
--data-raw '{
    "method": "HelloServer.Calc",
    "params": [
        {
            "a": 1,
            "b": 2
        }
    ]
}'
*/
