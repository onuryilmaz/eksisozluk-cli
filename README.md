[![Build Status](https://travis-ci.org/onuryilmaz/eksisozluk-cli.svg?branch=master)](https://travis-ci.org/onuryilmaz/eksisozluk-cli)


### Build
```go
go build
```

### Run

Başlık ile arama:
```go
eksisozluk-cli baslik "golang" [--page=PAGE_NUMBER] [--limit=ENTRY_LIMIT]
```

Gündem Başlıkları:
```go
eksisozluk-cli gundem [--page=PAGE_NUMBER] [--limit=ENTRY_LIMIT]
```

DEBE:
```go
eksisozluk-cli debe
```

### TODO
- [ ] Gündem: live; coloring (Coloring: error when limit < 15)
- [ ] Başlık: "şükela mod"
- [ ] Debe: Limit
- [ ] JSON, CSV export
- [ ] Logger


### Bilgi
* Ekşi Sözlük, [Ekşi Teknoloji](https://eksisozluk.com/eksi-teknoloji--1631416) şirketinin tescilli markasıdır.
* [Ekşi Sözlük Kullanım Koşulları](https://eksisozluk.com/eksi-sozluk-kullanim-kosullari--2602576)
