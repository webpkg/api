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
    "testName": "string",
    "testDescription": "string"
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

--------------------

### Data (Test):

| AttributeName | Required | Type | Validator | Comment |
|---------------|---------:|-----:|----------:|--------:|
|testName|Yes|string|||
|testDescription|No|string|||

### Where:

| AttributeName | Required | Type |
|---------------|---------:|-----:|
|id|Yes|uint64|