[![Build Status](https://travis-ci.org/onuryilmaz/eksisozluk-cli.svg?branch=master)](https://travis-ci.org/onuryilmaz/eksisozluk-cli)
[![Go Report Card](http://goreportcard.com/badge/onuryilmaz/eksisozluk-cli)](http://goreportcard.com/report/onuryilmaz/eksisozluk-cli)

### Ekşisözlük CLI Nedir?
* 

### Kullanım

#### İndir:
* İşletim sisteminize uygun paketi bağlantıdan indirebilirsiniz.
* İndirdiğiniz dosyanın adını **eksisozluk-cli** olarak değiştirerek örnek komutları çalıştırabilirsiniz.

#### Başlık ile Arama:
```shell
eksisozluk-cli baslik "golang" [--sukela] [--page=SAYFA_SAYISI] [--limit=ENTRY_LIMITI] [--output=json|console]
```

#### Gündem Başlıkları:
```shell
eksisozluk-cli gundem [--page=SAYFA_SAYISI] [--limit=BASLIK_LIMITI] [--output=json|console]
```

#### DEBE:
```shell
eksisozluk-cli debe [--limit=DEBE_LIMITI] [--output=json|console]
```


### TODO
- Gündem: live; coloring for windows
- Documentation 
 - Readme with detailed usage & Animated GIFs
 - Godoc
 - Github Pages

### Bilgi
* Ekşi Sözlük, [Ekşi Teknoloji](https://eksisozluk.com/eksi-teknoloji--1631416) şirketinin tescilli markasıdır.
* [Ekşi Sözlük Kullanım Koşulları](https://eksisozluk.com/eksi-sozluk-kullanim-kosullari--2602576)
