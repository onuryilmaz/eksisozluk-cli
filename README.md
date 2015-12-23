[![Build Status](https://travis-ci.org/onuryilmaz/eksisozluk-cli.svg?branch=master)](https://travis-ci.org/onuryilmaz/eksisozluk-cli)
[![Go Report Card](http://goreportcard.com/badge/onuryilmaz/eksisozluk-cli)](http://goreportcard.com/report/onuryilmaz/eksisozluk-cli) ![License Mozilla v2](http://b.repl.ca/v1/License-Mozilla_v2-red.png) [![GoDoc](https://godoc.org/github.com/onuryilmaz/eksisozluk-cli?status.svg)](https://godoc.org/github.com/onuryilmaz/eksisozluk-cli)

### Ekşisözlük CLI Nedir?
* Ekşi Sözlük komut satırı arayüzü; metin analizi, veri madenciliği gibi işlemleriniz için veri toplanmasını otomatik olarak yapmanıza yardımcı olur. Ayrıca gündeme komut satırından bakarak "nerd" olma şansınız da var!

### Kullanım

##### İndir:
* İşletim sisteminize uygun paketi [bağlantıdan](https://github.com/onuryilmaz/eksisozluk-cli/releases) indirebilirsiniz.
* İndirdiğiniz dosyanın adını **eksisozluk-cli** olarak değiştirerek örnek komutları çalıştırabilirsiniz.

##### Başlık ile Arama:
```shell
eksisozluk-cli baslik "golang" [--sukela] [--page=SAYFA_SAYISI] [--limit=ENTRY_LIMITI] [--output=json|console]
```
![Ekşisözlük CLI - Başlık ile Arama](http://i.imgur.com/kpJxrhQ.gif)

##### Gündem Başlıkları:
```shell
eksisozluk-cli gundem [--page=SAYFA_SAYISI] [--limit=BASLIK_LIMITI] [--output=json|console]
```
![Ekşisözlük Gündem](http://im.ezgif.com/tmp/ezgif-3294122878.gif)

##### DEBE:
```shell
eksisozluk-cli debe [--limit=DEBE_LIMITI] [--output=json|console]
```
![Ekşisözlük CLI - DEBE](http://i.imgur.com/0pqzU23.gif)

### Katkıda Bulunun
* Ekşi Sözlük CLI [Go](https://golang.org/) dili ile açık kaynak kodlu olarak yazıldı. 
* Eklenmesini istediğiniz bir özellik ya da bulduğunuz bir hata olursa [issue açarak](https://github.com/onuryilmaz/eksisozluk-cli/issues) bildirebilirsiniz.

#### Derleme
* Go 1.5 ve üzeri bir versiyon ile aşağıdaki komutları kullanarak uygulamayı derleyebilirsiniz:
```
go get github.com/tools/godep
go get github.com/onuryilmaz/eksisozluk-cli
cd $GOPATH/src/github.com/onuryilmaz/eksisozluk-cli
godep restore
go build
```

### Bilgi
* Ekşi Sözlük, [Ekşi Teknoloji](https://eksisozluk.com/eksi-teknoloji--1631416) şirketinin tescilli markasıdır.
* Geliştirilen uygulama bilgisayarınızdan Ekşi Sözlük web sitesine bağlanarak bilgi alır. [Ekşi Sözlük Kullanım Koşulları](https://eksisozluk.com/eksi-sozluk-kullanim-kosullari--2602576)'na aykırı kullanımlarda sorumluluk bu uygulamayı kullanana aittir.

### TODO
- Documentation
 - Godoc
 - Releases
 - Unit Tests
 - Github Pages
