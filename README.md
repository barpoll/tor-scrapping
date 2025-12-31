# Thor'un Scraperi - Tor Network Automation Tool

Bu proje, siber tehdit aktörlerinin izlerini kaybettirmek için kullandığı Tor ağındaki (sızıntı siteleri, forumlar, marketler) verileri düzenli ve anonim bir şekilde toplamak amacıyla geliştirilmiş bir otomasyon aracıdır. CTI süreçlerindeki "Collection" (Toplama) ve "Automation" (Otomasyon) yetkinliklerini kazandırmayı hedefler.

## 🚀 Proje Amacı ve Özellikleri

* **Anonimlik:** Trafiği yerel Tor servisi (SOCKS5 Proxy) üzerinden yönlendirerek tam anonimlik sağlar ve IP sızıntısını önler.
* **Otomasyon:** `targets.yaml` dosyasındaki toplu hedef listesini temizleyerek (whitespace trimming) otomatik olarak işler.
* **Veri Toplama:** Erişilen onion sitelerinin hem HTML içeriğini hem de tam boy ekran görüntülerini (screenshot) saklar.
* **Hata Yönetimi:** Kapanmış veya ulaşılamayan (dead) adresler programı durdurmaz; araç hatayı loglayıp bir sonraki URL'ye geçer.

## 🛠 Kullanılan Teknolojiler

**Dil:** Go (Golang) - Performans ve eşzamanlılık avantajları nedeniyle tercih edilmiştir.

* **Kritik Kütüphaneler:**
    * `net/http`: HTTP istekleri için.
    * `golang.org/x/net/proxy`: SOCKS5 proxy desteği için.
    * `chromedp`: Sayfa renderlama ve screenshot alımı için.
    * `os`, `bufio`: Dosya okuma ve yazma işlemleri için.
* **Ağ Altyapısı:** Tor Service (Port: 9150/9050).

## 📋 Kurulum ve Kullanım

1. **Tor Servisini Başlatın:** Arka planda Tor Browser veya Tor Expert Bundle'ın çalıştığından emin olun.
2. **Bağımlılıkları Yükleyin:**
   ```bash
   go mod tidy
3. Hedefleri Belirleyin: targets.yaml dosyasına taramak istediğiniz .onion adreslerini ekleyin.
4. **Uygulamayı Çalıştırın:**
   ```bash
   go run main.go
