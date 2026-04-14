package qrcode

import (
	"TicketRservation/env"
	"fmt"

	qrcodehandler "github.com/skip2/go-qrcode"
)

type QrcodeService struct {
	// the path to store the qrcode files
	path string
}

func NewQrcodeService() *QrcodeService {

	// get the path from the env var

	path := env.QrcodePath.GetValue()

	return &QrcodeService{path: path}
}

func (service *QrcodeService) NewQrcode(content, name string) {

	// to do ---> make the qrcode encoding based on a secret key

	name = service.path + name

	fmt.Println("new qr code is created with this path ", name)

	qrcodehandler.WriteFile(content, qrcodehandler.Medium, 256, name)

}

// checks if a qrcode exist and return the complete path of it
func (service *QrcodeService) GetQrcode(name string) string {

	return service.path + name
}
