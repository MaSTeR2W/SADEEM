# sadeem


## PREREQUISITES

1- Golang version 1.21.5

2- Postgres version 16.1

## HOW TO RUN

1- Go to .env file.

2- Change connection options with yours.

3- Then enter the following in console.

```bash
go run main.go
```

### HOW TO LOGIN

1- Open postman.
2- Add request with (localhost:1323/api/v1/registration/login) method:POST.
3- Send the request with: (email:sadeem@sadeem.ly, password:12345678)
4- You will get JWT token as a response.

note:
   send the JWT Token in Authorization header:(Bearer <token>) with all following requests.

### Attachments

1-postman_collection.json

use it to explore all methods provided by the api.

### Remarks

1- You can change the language of The error message from English (default) to Arabic by add a query param "?lang=ar".