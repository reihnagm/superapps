# Api Spec

## LOGIN 

Request : 
- Method : POST
- Endpoint : `/api/v1/login`
- Body :

```json
{
    "val": "089670558381",
    "password": "12345678",
}
```

Response : 

```json
{
    "status": 200,
    "error": false,
    "message": "Successfully",
    "data": {
        "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJpZCI6IjUyNzU3OWI5LTBmMDItNDY2NS05MWIyLTY3MzVlNzc1YTg1ZSJ9.3ELrZBcwmgUy_mtNL5xiWoq36v7_rnkrfLcz9bnRmwc"
    }
}
```

## REGISTER

Request : 
- Method : POST
- Endpoint : `/api/v1/register`
- Body :

```json
{
    "email": "reihanagam7@gmail.com",
    "phone": "089670558381",
    "password": "12345678",
    "app_name": "myapp"
}
```

Response : 

```json
{
    "status": 200,
    "error": false,
    "message": "Successfully",
    "data": {
        "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJpZCI6IjAzMWZjMWVmLWRmYzEtNDA3Ny05NTM3LTFmNmFjYWI3ZDJjNSJ9.BsjXvoHoEkoLVTY2-XLXJgYu8jfU5AG-aKxN5r-iN34"
    }
}
```

## VERIFY OTP 

Request : 
- Method : POST
- Endpoint : `/api/v1/verify-otp`
- Body :

```json
{
    "email": "reihanagam8@gmail.com",
    "otp": "SEEA"
}
```

Response : 

```json
{
    "status": 200,
    "error": false,
    "message": "Successfully",
    "data": {
        "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJpZCI6IjAzMWZjMWVmLWRmYzEtNDA3Ny05NTM3LTFmNmFjYWI3ZDJjNSJ9.BsjXvoHoEkoLVTY2-XLXJgYu8jfU5AG-aKxN5r-iN34"
    }
}
```

## RESEND OTP

Request : 
- Method : POST
- Endpoint : `/api/v1/resend-otp`
- Body :

```json
{
    "email": "reihanagam8@gmail.com",
}
```

Response : 

```json
{
    "status": 200,
    "error": false,
    "message": "Successfully",
    "data": {
        "otp": "A8F1"
    }
}
```

## NEWS ALL

Request : 
- Method : GET
- Endpoint : `/api/v1/news?page=1&limit=10&search=&app_name=`

Response : 

```json
{
    "status": 200,
    "error": false,
    "message": "Successfully",
    "total": 3,
    "per_page": 1,
    "prev_page": 1,
    "next_page": 2,
    "current_page": 1,
    "next_url": "http://localhost:5004?page=2",
    "prev_url": "http://localhost:5004?page=1",
    "data": [
        {
            "id": "f25f0679-1920-48c9-a6e2-63bcf315498c",
            "title": "Title 1",
            "desc": "lorem metus viverra tortor, eu tempus orci leo imperdiet purus. Nam nec consectetur mi, nec vestibulum nulla. Donec vel dui lectus. Etiam non scelerisque odio. Pellentesque eleifend nisi et odio commodo gravida. Vivamus pellentesque elementum eros, vitae ultrices magna sollicitudin ac.",
            "images": [],
            "app": {
                "id": "ada7582e-04dc-486d-819c-467986c1cb91",
                "name": "My App"
            },
            "user": {
                "fullname": "",
                "email": "reihanagam7@gmail.com",
                "phone": "089670558381"
            },
            "created_at": "2023-10-06 16:56"
        },
    ]
}
```

## CREATE NEWS

Request : 
- Method : POST
- Endpoint : `/api/v1/news`
- Body :

```json
{
    "id": "409630ea-8daa-40c8-9196-a2a3aa473fdf",
    "title": "Test",
    "desc": "lorem metus viverra tortor, eu tempus orci leo imperdiet purus. Nam nec consectetur mi, nec vestibulum nulla. Donec vel dui lectus.",
}
```

## ASSIGN NEWS IMG

Request : 
- Method : POST
- Endpoint : `/api/v1/news-upload-image`
- Body :

```json
{
    "id": "409630ea-8daa-40c8-9196-a2a3aa473fdf",
    "title": "Test",
    "desc": "lorem metus viverra tortor, eu tempus orci leo imperdiet purus. Nam nec consectetur mi, nec vestibulum nulla. Donec vel dui lectus.",
}
```
