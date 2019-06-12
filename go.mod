module opslabGo

replace (
	golang.org/x/crypto => github.com/golang/crypto v0.0.0-20190611184440-5c40567a22f8
	golang.org/x/net => github.com/golang/net v0.0.0-20190611141213-3f473d35a33a
	golang.org/x/sys => github.com/golang/sys v0.0.0-20190610200419-93c9922d18ae
	golang.org/x/text => github.com/golang/text v0.3.0
	golang.org/x/tools => github.com/golang/tools v0.0.0-20190611222205-d73e1c7e250b
)

require (
	github.com/PuerkitoBio/goquery v1.5.0
	github.com/go-redis/redis v6.15.2+incompatible //indirect
	github.com/google/logger v1.0.1
	github.com/sirupsen/logrus v1.4.2 // indirect
)

go 1.12
