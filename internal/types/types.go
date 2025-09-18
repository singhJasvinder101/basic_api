package types

// type Student struct {
// 	Id int `json:"id"`
// 	Name string `json:"name"`
// 	Email string `json:"email"`
// 	Age int `json:"age"`
// }



type Student struct {
    ID    int    `json:"id"`
    Name  string `json:"name" validate:"required"`
    Email string `json:"email" validate:"required,email"`
    Age   int    `json:"age" validate:"gte=1,lte=100"`
}

type APIResponse struct {
    Status string      `json:"status"`
    Data   interface{} `json:"data,omitempty"`
    Error  string      `json:"error,omitempty"`
}
