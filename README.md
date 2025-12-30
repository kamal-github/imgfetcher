# imgfetcher

`imgfetcher` is a Go-based CLI tool that generates **printable learning cards for kids** by downloading free images, adding **large readable labels**, and organizing them into **category-wise folders**.

Designed for:

* Early learning
* Kids-friendly flash cards
* DIY printable cards

---

## ✨ Features

* 📄 YAML-driven input (no code changes needed)
* 🧠 Category-wise organization (Kitchen / Bedroom / Bathroom, etc.)
* 🖼️ Downloads **multiple images per item**
* 🏷️ Adds **bold labels directly on images**
* ⚡ Concurrent image fetching
* 🔐 Secure API key handling via environment variables
* 🧱 Production-grade CLI using Cobra

---

## 📦 Installation

### Option 1: Run locally

```bash
git clone https://github.com/kamal-github/imgfetcher.git
cd imgfetcher
go mod tidy
```

### Option 2: Install globally (recommended later)

```bash
go install github.com/kamal-github/imgfetcher@latest
```

---

## 🔑 Pixabay API Key (Required)

`imgfetcher` uses **Pixabay** for free, watermark-free images.

### Step 1: Get a free API key

👉 [https://pixabay.com/api/docs/](https://pixabay.com/api/docs/)

### Step 2: Set environment variable

```bash
export PIXABAY_API_KEY="your_pixabay_api_key"
```

> 💡 **Important**
>
> * The API key is **not** passed as a CLI flag
> * The program will **fail fast** if this variable is not set
> * This follows best practices for security and CI/CD usage

### ❌ If missing, you’ll see:

```
PIXABAY_API_KEY environment variable not set
```

---

## 🧾 YAML Input Format

Create a YAML file describing categories and items.

### `home_items.yaml`

```yaml
categories:
  LivingRoom:
    - television
    - sofa
    - remote control
    - fan

  Kitchen:
    - refrigerator
    - microwave oven
    - kitchen stove
    - mixer grinder

  Bedroom:
    - bed
    - lamp
    - cupboard
    - clock

  Bathroom:
    - toilet
    - bathroom sink
    - shower
    - mirror
```

You can modify this file anytime without touching Go code.

---

## 🚀 Usage

```bash
imgfetcher generate \
  --input home_items.yaml \
  --out images \
  --images-per-item 2 \
  --workers 5
```

### Flags

| Flag                | Description        | Default      |
| ------------------- | ------------------ | ------------ |
| `--input`, `-i`     | YAML input file    | **required** |
| `--out`, `-o`       | Output directory   | `images`     |
| `--images-per-item` | Images per item    | `2`          |
| `--workers`         | Concurrent workers | `5`          |

---

## 📁 Output Structure

```
images/
├── Kitchen/
│   ├── refrigerator_1.jpg
│   ├── refrigerator_2.jpg
│
├── Bedroom/
│   ├── bed_1.jpg
│   ├── lamp_1.jpg
│
├── Bathroom/
│   ├── toilet_1.jpg
│
└── LivingRoom/
    ├── television_1.jpg
```

Each image contains:

* A **large clear photo**
* A **black label bar**
* **High-contrast white text** (ideal for kids)

---

## 🖨️ Printing Tips (Recommended)

* Print one image per page (A4)
* Use thick paper (200–250 GSM)
* Laminate if possible
* One concept per card

Perfect for:

* Speech therapy
* Visual learning
* Autism-friendly teaching

---

## 🛠️ Development Notes

* CLI built with **Cobra**
* Image rendering uses **fogleman/gg**
* YAML parsing via `yaml.v3`
* Secrets handled via environment variables
* Internal logic isolated under `internal/`

---

## 🧩 Roadmap (Optional Enhancements)

* 📄 PDF export (`imgfetcher pdf`)
* 🌍 Multi-language labels (EN / HI)
* 🧪 YAML validation command
* 🔌 Multiple image providers (Pexels, Unsplash)
* 🧠 Action cards (washing hands, sleeping)

---

## ❤️ Why this exists

This tool was built to help parents create **custom, meaningful learning material** for their kids — especially for children who learn best with **visual structure and repetition**.

---