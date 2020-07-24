## Updata Test

### URL

  PUT /test/:id

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
    "updatedAt": "time.Time"
}
```

### Success Response:
#### Status Code: 200
#### Data (rowsAffected int64):
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
|testName|Yes|string||
|testDescription|No|string||
|updatedAt|No|time.Time||

### Where:

| AttributeName | Required | Type |
|---------------|---------:|-----:|
|id|Yes|uint64|