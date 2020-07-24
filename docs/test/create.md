## Create Test

### URL

  POST /test/

### Header
```bash
Content-Type: application/json
Authorization: Bearer $ACCESSTOKEN
```

### Payload (Test)
```json
{
    "id": "uint64",
    "testName": "string",
    "testDescription": "string",
    "createdAt": "time.Time"
}
```

### Success Response:
#### Status Code: 200
#### Data (id uint64):
```json

```

### Error Response:
#### Status Code: 400
#### Data:
```json
"error message"
```

### Error Response:
#### Status Code: 401
#### Data:
```json
"invalid token"
```

### Error Response:
#### Status Code: 403
#### Data:
```json
"permission denied"
```

--------------------
### Data:

| AttributeName | Required | Type | Validator |
|---------------|---------:|-----:|----------:|
|id|Yes|uint64||
|testName|Yes|string||
|testDescription|No|string||
|createdAt|No|time.Time||