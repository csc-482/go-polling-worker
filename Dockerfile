FROM golang:1.11.5
LABEL Andy Vadnais

ENV APP /~/Programs/go/src/go-polling-worker/
WORKDIR /~/Programs/go/src/go-polling-worker/
 
ADD . $APP

RUN cd ${APP} && go get -d -v github.com/jamespearly/loggly
RUN go build pollingWorker.go


ENV LOGGLY_TOKEN 1e81e107-fe9a-48e4-8ac7-75c9f6c1fa43
ENV TRN_API_KEY d8b929fb-27a1-48dd-a6ac-ef2092db1291

CMD ./pollingWorker


## build in /Users/andyvadnais/Programs/go/src/go-polling-worker
## use command (to build): docker build -t go-polling-worker:1.0.0 .
## and command (to run): docker run go-polling-worker:1.0.0
