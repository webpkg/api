## Get Tests

### URL

  GET /test/?key=&page=1&pagesize=20

### Header
```bash
Content-Type: application/json
Authorization: Bearer $ACCESSTOKEN
```

### Success Response:
#### Status Code: 200
#### Data (TestCollection):
```json
[
    {
        "id": "uint64",
        "testName": "string",
        "testDescription": "string",
        "createdAt": "time.Time",
        "updatedAt": "time.Time"
    }
]
```

### Error Response:
#### Status Code: 400
#### Data:
```json
"error message"
```

--------------------

### Query:

| QueryName | Required | Type |
|-----------|---------:|-----:|
|  key      |    No    |string|
|  page     |    No    |  int |
|  pagesize |    No    |  int |

### Data (Test):

| AttributeName | Required | Type | Comment |
|---------------|---------:|-----:|--------:|
|id|Yes|uint64||
|testName|Yes|string||
|testDescription|No|string||
|createdAt|No|time.Time||
|updatedAt|No|time.Time||