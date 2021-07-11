run:
	clear && go build -race
	go test ./... -v
	./httptool -parallel 2 http://google.com  http://adjust.com  facebook.com  twitter.com  adjust.com

runAll:
	clear && go build -race
	go test ./... -v
	./httptool -parallel 3 http://adjust.com http://google.com http://facebook.com http://yahoo.com http://yandex.com http://twitter.com http://reddit.com/r/funny http://reddit.com/r/notfunny http://baroquemusiclibrary.com http://adjust.com http://google.com http://facebook.com http://yahoo.com http://yandex.com http://twitter.com http://reddit.com/r/funny http://reddit.com/r/notfunny http://baroquemusiclibrary.com http://adjust.com http://google.com http://facebook.com http://yahoo.com http://yandex.com http://twitter.com http://reddit.com/r/funny http://reddit.com/r/notfunny http://baroquemusiclibrary.com http://adjust.com http://google.com http://facebook.com http://yahoo.com http://yandex.com http://twitter.com http://reddit.com/r/funny http://reddit.com/r/notfunny http://baroquemusiclibrary.com http://adjust.com http://google.com http://facebook.com http://yahoo.com http://yandex.com http://twitter.com http://reddit.com/r/funny http://reddit.com/r/notfunny http://baroquemusiclibrary.com http://adjust.com http://google.com http://facebook.com http://yahoo.com http://yandex.com http://twitter.com http://reddit.com/r/funny http://reddit.com/r/notfunny http://baroquemusiclibrary.com http://adjust.com http://google.com http://facebook.com http://yahoo.com http://yandex.com http://twitter.com http://reddit.com/r/funny http://reddit.com/r/notfunny http://baroquemusiclibrary.com http://adjust.com http://google.com http://facebook.com http://yahoo.com http://yandex.com http://twitter.com http://reddit.com/r/funny http://reddit.com/r/notfunny http://baroquemusiclibrary.com http://adjust.com http://google.com http://facebook.com http://yahoo.com http://yandex.com http://twitter.com http://reddit.com/r/funny http://reddit.com/r/notfunny http://baroquemusiclibrary.com http://adjust.com http://google.com http://facebook.com http://yahoo.com http://yandex.com http://twitter.com http://reddit.com/r/funny http://reddit.com/r/notfunny http://baroquemusiclibrary.com http://adjust.com http://google.com http://facebook.com http://yahoo.com http://yandex.com http://twitter.com http://reddit.com/r/funny http://reddit.com/r/notfunny http://baroquemusiclibrary.com http://adjust.com http://google.com http://facebook.com http://yahoo.com http://yandex.com http://twitter.com http://reddit.com/r/funny http://reddit.com/r/notfunny http://baroquemusiclibrary.com

.Phony:
	run
