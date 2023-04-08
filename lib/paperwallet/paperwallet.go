package paperwallet

import (
	"fmt"
	"image"

	"github.com/ejfhp/gobop/lib/graphic"
)

type PaperWallet struct {
	Address string
	WIF     string
	Graphic image.Image
	Config  *Config
}

func (p *PaperWallet) Image() (image.Image, error) {
	qrAddress, err := graphic.QRCode(p.Address)
	if err != nil {
		return nil, fmt.Errorf("cannot make QRCodde of address: %w", err)
	}
	qrKey, err := graphic.QRCode(p.Address)
	if err != nil {
		return nil, fmt.Errorf("cannot make QRCodde of key: %w", err)
	}
	//ADDRESS
	walletImage, err := graphic.AddText(p.Graphic, p.Address, p.Config.FontName, p.Config.Address.Color, p.Config.Address.Size, p.Config.Address.Rotation, p.Config.Address.X, p.Config.Address.Y)
	if err != nil {
		return nil, fmt.Errorf("cannot print address: %w", err)
	}
	//ADDRESS QRCODE
	walletImage = graphic.AddImage(walletImage, qrAddress, p.Config.AddressQR.Size, p.Config.AddressQR.Rotation, p.Config.AddressQR.X, p.Config.AddressQR.Y)

	//KEY
	walletImage, err = graphic.AddText(walletImage, p.WIF, p.Config.FontName, p.Config.Key.Color, p.Config.Key.Size, p.Config.Key.Rotation, p.Config.Key.X, p.Config.Key.Y)
	if err != nil {
		return nil, fmt.Errorf("cannot print address: %w", err)
	}

	//KEY QRCODE
	walletImage = graphic.AddImage(walletImage, qrKey, p.Config.KeyQR.Size, p.Config.KeyQR.Rotation, p.Config.KeyQR.X, p.Config.KeyQR.Y)
	return walletImage, nil
}
