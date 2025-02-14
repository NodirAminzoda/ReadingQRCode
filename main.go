package main

import (
	"fmt"
	"image"
	"os"

	"github.com/makiuchi-d/gozxing"
	"github.com/makiuchi-d/gozxing/qrcode"

	_ "image/png"
)

func ReadingQR(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Не смогли открыть файл с QR-кодом")
		return "", err
	}
	defer file.Close()

	// Декодирование изображения
	img, _, err := image.Decode(file)
	if err != nil {
		fmt.Println("Не смогли декодировать изображение")
		return "", err
	}

	// Преобразование изображения в битовую карту
	bitmap, err := gozxing.NewBinaryBitmap(gozxing.NewHybridBinarizer(gozxing.NewLuminanceSourceFromImage(img)))
	if err != nil {
		fmt.Println("Не смогли преобразовать изображения в битовую карту")
	}
	// Читаем QR-код
	reader := qrcode.NewQRCodeReader()
	result, err := reader.Decode(bitmap, nil)
	if err != nil {
		fmt.Println("Не смогли расшифровать QR-код")
		return "", err
	}

	return result.GetText(), nil
}

func main() {
	decodedText, err := ReadingQR("./qrcode_example.png")
	if err != nil {
		fmt.Println("Ошибка:", err)
	} else {
		fmt.Println("Декодированный текст:", decodedText)
	}
}
