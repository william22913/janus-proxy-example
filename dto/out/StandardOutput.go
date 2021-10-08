package out

type StandardOutput struct{
	Success bool `json:"success"`
	Transaction string `json:"transaction"`
	Request string `json:"request"`
	Payload interface{} `json:"payload"`
}