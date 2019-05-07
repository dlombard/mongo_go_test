FROM iron/go:dev
WORKDIR /app
ENV SRC_DIR=/go/src/github.com/dlombard/test_go

ADD . $SRC_DIR

CMD ls -alrt $SRC_DIR
RUN cd $SRC_DIR; go build -o testgoapp; cp testgoapp /app/
ENTRYPOINT [ "./testgoapp" ]