# message format define

a common format:
```json
{
  "id": "uuid",
  "type": "service type",
  "operation":"create/delete/update",
  "content": {
    "self_define": "self_define"
  }
}
```

## apigw
service create example
```json
{
  "id": "uuid",
  "type": "apigw",
  "operation":"create",
  "callback":"http://localhost:8000/callback",
  "content": {
    "callback":"status update callback url",
    "service": {
      "name": "test",
      "upstream_name": "test_upstream",
      "schema": "http",
      "endpoints": [
        "localhost:8000",
        "localhost:8001"
      ],
      "auth_type":"apikey"
    }
  }
}
```