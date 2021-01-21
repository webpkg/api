## Destroy Test

### URL

  DELETE /test/:id

### Header
```bash
Content-Type: application/json
Authorization: Bearer $ACCESSTOKEN
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

### Where:

| AttributeName | Required | Type |
|---------------|---------:|-----:|
|id|Yes|uint64|