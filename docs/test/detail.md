## Get Test

### URL

  GET /test/:id

### Header
```bash
Content-Type: application/json
Authorization: Bearer $ACCESSTOKEN
```

### Success Response:
#### Status Code: 200
#### Data (Test):
```json
{
    "id": "uint64",
    "testName": "string",
    "testDescription": "string",
    "createdAt": "time.Time",
    "updatedAt": "time.Time"
}
```

### Error Response:
#### Status Code: 400
#### Data:
```json
"error message"
```

--------------------

### Data (Test):

| AttributeName | Required | Type | Comment |
|---------------|---------:|-----:|--------:|
|id|Yes|uint64||
|testName|Yes|string||
|testDescription|No|string||
|createdAt|No|time.Time||
|updatedAt|No|time.Time||