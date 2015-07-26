package fpoimages

import (
//    "log"
    "image"
    "image/draw"
    "image/png"
    "net/http"
    "regexp"
    "strconv"
    "strings"
    "image/color"
    "encoding/hex"
)

func init() {
    http.HandleFunc("/favicon.ico", handleIcon)
    http.HandleFunc("/", handler)
}

func handleIcon(w http.ResponseWriter, r *http.Request) {
}

func handler(w http.ResponseWriter, r *http.Request) {
    var width, height int
    var fullsize, transparent bool
    var userColor, setColor string
    path := strings.ToLower(strings.TrimPrefix(r.URL.Path, "/"))

    fullsize = boolify(r.FormValue("f"))
    transparent = boolify(r.FormValue("t"))
    userColor = r.FormValue("c")

    if isValidDimensions(path) {
        width, height = extractDimensions(path, fullsize)
    } else {
        width, height = 1, 1
    }
    if isValidColor(userColor) {
        setColor = userColor
    } else {
        setColor = "C0C0C0"
    }
    if userColor == "" {
        transparent = true
    }
    w.Header().Set("Content-Type", "image/png")
    png.Encode(w, drawImage(width, height, setColor, transparent))
}

func drawImage(width, height int, hexColor string, transparent bool) (img draw.Image) {
    img = image.NewRGBA(image.Rect(0, 0, width, height))
    if transparent {
        draw.Draw(img, img.Bounds(), image.Transparent, image.ZP, draw.Src)
        return
    } else {
        r, g, b := convertHex(hexColor)
        rgbColor := color.RGBA{r, g, b, 255}
        srcImg := image.NewUniform(rgbColor)
        draw.Draw(img, img.Bounds(), srcImg, image.ZP, draw.Src)
        return
    }
    return
}

func boolify(boolString string) (boolBool bool) {
    if boolString == "true" {
        return true
    } else {
        return false
    }
}

func isValidDimensions(path string) (isValid bool) {
    reg := regexp.MustCompile("^\\d{1,4}[xX]\\d{1,4}$")
    isValid = reg.MatchString(path)
//    log.Println("isValidDimension: ", isValid)
    return
}

func isValidColor(hex string) (isValid bool) {
    reg := regexp.MustCompile("^[0-9a-fA-F]{6}$")
    isValid = reg.MatchString(hex)
//    log.Println("isValidColor: ", isValid)
    return
}

func extractDimensions(path string, fullSize bool) (width, height int) {
    dimensions := strings.Split(path, "x")
    givenWidth, _ := strconv.Atoi(dimensions[0])
    givenHeight, _ := strconv.Atoi(dimensions[1])
    if len(dimensions) == 2 {
        if fullSize {
            width, height = givenWidth, givenHeight
        } else {
            width, height = applyGcd(givenWidth, givenHeight)
        }
    } else {
        width, height = 1, 1
    }
    return
}

func convertHex(hexColor string) (r, g, b uint8) {
//    log.Println(hex.DecodeString(hexColor[0:2]))
    rbyte, _ := hex.DecodeString(hexColor[0:2])
    gbyte, _ := hex.DecodeString(hexColor[2:4])
    bbyte, _ := hex.DecodeString(hexColor[4:6])
    r = rbyte[0]
    g = gbyte[0]
    b = bbyte[0]
    return
}

func applyGcd(x, y int) (int, int) {
    a, b := x, y
    if a == 0 || b == 0 {
        a = 1
    } else {
        for b != 0 {
            a, b = b, a%b
        }
    }
    return x / a, y / a
}
