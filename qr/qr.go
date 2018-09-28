package qr

import (
	terminal "github.com/lizebang/qrcode-terminal"
	qrcode "github.com/skip2/go-qrcode"
)

// QR generates a QR code and prints it on the command line.
func QR(content string) {
	terminal.QRCode(content, terminal.BrightBlack, terminal.BrightWhite, qrcode.Highest)
}
