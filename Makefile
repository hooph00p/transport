# Never made a makefile before
default:
	go get github.com/satori/go.uuid
	go get github.com/gin-gonic/gin
	go get github.com/itsjamie/gin-cors
	go build

install:
	go install

clean:
	go clean
