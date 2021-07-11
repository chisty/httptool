run:
	clear && go build -race
	./httptool http://google.com  http://adjust.com  http//facebook.com  http//twitter.com

.Phony:
	run
