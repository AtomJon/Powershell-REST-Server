module github.com/AtomJon/SubscriptRESTServer

go 1.17

replace github.com/AtomJon/SubscriptRESTServer/handler => ./handler

replace github.com/AtomJon/SubscriptRESTServer/resource => ./resource

replace github.com/AtomJon/SubscriptRESTServer/executor => ./executor

require github.com/mattn/go-zglob v0.0.3
