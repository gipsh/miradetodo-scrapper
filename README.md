# miradetodo-scrapper

Extract the content links from http://miradetodo.co  so you can watch or download 


## usage 

1. Go to http://miradetodo.co and select the movie or tv show you want to watch  and copy te link 

2. Then run the scrapper 

```bash
./miradetodo-scrapper https://miradetodo.co/aves-de-presa-y-la-fantabulosa-emancipacion-de-una-harley-quinn-2020-720p-hd/
```

You will get the list of links for each provieder:


```bash
[*] link: [https://player.miradetodo.net/generator/embed/04426] 
	 0. https://fastplay.to/d8w9w0vv7k2i.html
	 1. https://vidoza.net/iu63w23ee633.html
	 2. http://ok.ru/video/1634522827337
	 3. https://1fichier.com/?m7g8sf7fuxndu08e79bv
	 4. https://jawcloud.co/8am7ysdvdbqa.html
	 5. https://vev.io/n5o40wx7pnr7
	 6. https://vup.to/c223azpstivn.html
	 7. https://vidlox.me/v4tmzhqom0tx
	 8. https://clipwatching.com/d5yx32vj052p/Birds.Of.Prey.2020-MDT.mp4.html
	 9. https://www.yourupload.com/embed/87R8nH21W3T2

```

3. Watch directly on your browser or use [youtube-dl](https://github.com/ytdl-org/youtube-dl) to download 


## build 

as usual, get deps first

```bash
go get
```

then build 


```bash
go build
```

