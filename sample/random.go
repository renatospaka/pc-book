package sample

import (
	"math/rand"
	"time"

	"github.com/google/uuid"
	"github.com/renatospaka/pc-book/pb"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func randomKeyboardLayout() pb.Keyboard_Layout {
	switch rand.Intn(3) {
	case 1:
		return pb.Keyboard_QWERTY
	case 2:
		return pb.Keyboard_QWERTZ
	default:
		return pb.Keyboard_AZERTY
	}
}

func randomBool() bool {
	return rand.Intn(2) == 1
}

func randomCPUBrand() string {
	return randomStringFromSel("Intel", "AMD")
}

func randomGPUBrand() string {
	return randomStringFromSel("NVIDIA", "AMD")
}

func randomCPUName(brand string) string {
	if brand == "Intel" {
		return randomStringFromSel(
			"Xeon E-2286M",
			"Core i9-9980HK",
			"Core i7-9750H",
			"Core i5-940F",
			"Core i3-100SG1",
		)
	}
	return randomStringFromSel(
		"Ryzen 7 PRO 2700U",
		"Ryzen 5 PRO 3500",
		"Ryzen 3 PRO 3200GE",
	)
}

func randomGPUName(brand string) string {
	if brand == "NVIDIA" {
		return randomStringFromSel(
			"RTX 2060",
			"RTX 2070",
			"RTX 1660-Ti",
			"RTX 1070",
		)
	}
	return randomStringFromSel(
		"RX 590",
		"RX 580",
		"RX 5700-XT",
		"RX VEGA-56",
	)
}

func randomScreenPanel() pb.Screen_Panel {
	if rand.Intn(2) == 1 {
		return pb.Screen_IPS
	}
	return pb.Screen_OLED
}

func randomScreenResolution() *pb.Screen_Resolution {
	height := randomInt(13, 17)
	width := height * 16 / 9
	screen := &pb.Screen_Resolution{
		Width:  uint32(width),
		Height: uint32(height),
	}
	return screen
}

func randomID() string {
	return uuid.NewString()
}

func randomLaptopBrand() string {
	return randomStringFromSel("Dell", "Apple", "Lenovo", "Acer")
}

func randomLaptopName(brand string) string {
	switch brand {
	case "Dell":
		return randomStringFromSel("Latitude", "Vostro", "XPS", "Alienware")

	case "Apple":
		return randomStringFromSel("Macbook Air", "Macbook Pro")

	case "Lenovo":
		return randomStringFromSel("Thinkpad X1", "Thinkpad P1", "Thinkpad PS3")

	default:
		return randomStringFromSel("Predator", "Nitro", "Aspire")
	}
}

func randomInt(min, max int) int {
	return min + rand.Intn(max-min+1)
}

func randomFloat32(min, max float32) float32 {
	return min + rand.Float32()*(max-min)
}

func randomFloat64(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}

func randomStringFromSel(a ...string) string {
	n := len(a)
	if n == 0 {
		return ""
	}
	return a[rand.Intn(n)]
}
