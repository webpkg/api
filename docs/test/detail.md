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

### Test
```go
// Test model
type Test struct {
	// @column $columnName=id,$dataType=bigint(20) unsigned,PrimaryKey
	ID uint64 `json:"id"`
	// @column $columnName=test_name,$dataType=varchar(127)
	TestName string `json:"testName"`
	// @column $columnName=test_description,$dataType=varchar(255)
	TestDescription *string `json:"testDescription"`
	// @column $columnName=created_at,$dataType=timestamp
	CreatedAt *time.Time `json:"createdAt"`
	// @column $columnName=updated_at,$dataType=timestamp
	UpdatedAt *time.Time `json:"updatedAt"`
}
```