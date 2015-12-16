[![Build Status](https://travis-ci.org/onuryilmaz/eksisozluk-cli.svg?branch=master)](https://travis-ci.org/onuryilmaz/eksisozluk-cli)


### Build
```go
go build
```

### Run

Başlık ile arama:
```go
eksisozluk-cli baslik "golang" [--page=SAYFA_SAYISI] [--limit=ENTRY_LIMITI]
```

Gündem Başlıkları:
```go
eksisozluk-cli gundem [--page=SAYFA_SAYISI] [--limit=BASLIK_LIMITI]
```

DEBE:
```go
eksisozluk-cli debe [--limit=DEBE_LIMITI]
```

### TODO
- [ ] Gündem: live; coloring (Coloring: error when limit < 15)
- [ ] Başlık: "şükela mod"
- [ ] JSON, CSV export
- [ ] Logger


### Bilgi
* Ekşi Sözlük, [Ekşi Teknoloji](https://eksisozluk.com/eksi-teknoloji--1631416) şirketinin tescilli markasıdır.
* [Ekşi Sözlük Kullanım Koşulları](https://eksisozluk.com/eksi-sozluk-kullanim-kosullari--2602576)
