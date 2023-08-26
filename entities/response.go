package entities

type Response struct {
    Status       int `json:"status"`
	Error 		 bool `json:"error"`
    Message   	 string `json:"message"`
    Data         map[string]interface{} `json:"data"`
}