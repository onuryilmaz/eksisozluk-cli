[![Build Status](https://travis-ci.org/onuryilmaz/eksisozluk-cli.svg?branch=master)](https://travis-ci.org/onuryilmaz/eksisozluk-cli)


### Build
```go
go build
```

### Run

Başlık ile arama:
```go
eksisozluk-cli baslik "golang" [--sukela] [--page=SAYFA_SAYISI] [--limit=ENTRY_LIMITI] [--output=json|console]
```

Gündem Başlıkları:
```go
eksisozluk-cli gundem [--page=SAYFA_SAYISI] [--limit=BASLIK_LIMITI]  [--output=json|console]
```

DEBE:
```go
eksisozluk-cli debe [--limit=DEBE_LIMITI] [--output=json|console]
```

### TODO
- Gündem: live; coloring for windows
- Documentation (Readme, Animated GIF, Godoc)

### Bilgi
* Ekşi Sözlük, [Ekşi Teknoloji](https://eksisozluk.com/eksi-teknoloji--1631416) şirketinin tescilli markasıdır.
* [Ekşi Sözlük Kullanım Koşulları](https://eksisozluk.com/eksi-sozluk-kullanim-kosullari--2602576)
