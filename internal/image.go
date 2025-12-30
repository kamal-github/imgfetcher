package internal

import (
	"encoding/json"
	"fmt"
	"image"
	"image/jpeg"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"

	"github.com/fogleman/gg"
)

const (
	// fontPath       = "./Roboto-Bold.ttf"
	labelBarHeight = 80
	labelFontSize  = 48
)

func getFontPath() (string, error) {
	exe, err := os.Executable()
	if err != nil {
		return "", err
	}
	return filepath.Join(filepath.Dir(exe), "Roboto-Bold.ttf"), nil
}

type pixabayResponse struct {
	Hits []struct {
		WebformatURL string `json:"webformatURL"`
	} `json:"hits"`
}

func processJob(job Job, outDir string) error {
	imgURL, err := fetchImageURL(job.Item, job.Index)
	if err != nil {
		return err
	}

	resp, err := http.Get(imgURL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	img, _, err := image.Decode(resp.Body)
	if err != nil {
		return err
	}

	labeled := addLabel(img, strings.ToUpper(job.Item))

	categoryDir := filepath.Join(outDir, job.Category)
	if err := os.MkdirAll(categoryDir, 0755); err != nil {
		return err
	}

	filename := fmt.Sprintf("%s_%d.jpg", sanitize(job.Item), job.Index)
	filePath := filepath.Join(categoryDir, filename)

	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	return jpeg.Encode(file, labeled, &jpeg.Options{Quality: 90})
}

func fetchImageURL(item string, index int) (string, error) {
	apiKey, err := GetPixabayAPIKey()
	if err != nil {
		return "", err
	}

	searchURL := fmt.Sprintf(
		"https://pixabay.com/api/?key=%s&q=%s&image_type=photo&safesearch=true&per_page=5",
		apiKey,
		url.QueryEscape(item),
	)

	resp, err := http.Get(searchURL)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var result pixabayResponse
	if err = json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", err
	}

	if len(result.Hits) < index {
		return "", fmt.Errorf("not enough images for %s", item)
	}

	return result.Hits[index-1].WebformatURL, nil
}

func addLabel(img image.Image, text string) image.Image {
	w := img.Bounds().Dx()
	h := img.Bounds().Dy() + labelBarHeight

	dc := gg.NewContext(w, h)

	// background
	dc.SetRGB(1, 1, 1)
	dc.Clear()

	// image
	dc.DrawImage(img, 0, 0)

	// black label bar
	dc.SetRGB(0, 0, 0)
	dc.DrawRectangle(0, float64(h-labelBarHeight), float64(w), labelBarHeight)
	dc.Fill()

	// text
	fontPath, err := getFontPath()
	if err == nil {
		_ = dc.LoadFontFace(fontPath, labelFontSize)
	}

	_ = dc.LoadFontFace(fontPath, labelFontSize)
	dc.SetRGB(1, 1, 1)
	dc.DrawStringAnchored(
		text,
		float64(w/2),
		float64(h-labelBarHeight/2),
		0.5,
		0.5,
	)

	return dc.Image()
}

func sanitize(s string) string {
	return strings.ReplaceAll(strings.ToLower(s), " ", "_")
}
