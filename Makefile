run:
	clear && go build -race
	./httptool http://google.com  http://adjust.com  facebook.com  twitter.com

.Phony:
	run
