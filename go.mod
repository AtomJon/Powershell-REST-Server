module github.com/AtomJon/subscriptrestserver

go 1.17

replace github.com/AtomJon/subscriptrestserver/handler => ./handler

replace github.com/AtomJon/subscriptrestserver/resource => ./resource

replace github.com/AtomJon/subscriptrestserver/executor => ./executor

require github.com/mattn/go-zglob v0.0.3
