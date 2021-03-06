# cirello.io/oversight

[![travis-ci](https://api.travis-ci.org/ucirello/oversight.svg?branch=master)](https://travis-ci.org/ucirello/oversight)
[![GoDoc](https://godoc.org/cirello.io/oversight?status.svg)](https://godoc.org/cirello.io/oversight)
[![gocover.io](https://gocover.io/_badge/cirello.io/oversight)](https://gocover.io/cirello.io/oversight)
[![Go Report Card](https://goreportcard.com/badge/github.com/ucirello/oversight)](https://goreportcard.com/report/github.com/ucirello/oversight)
[![SLA](https://img.shields.io/badge/SLA-95%25-brightgreen.svg)](https://github.com/ucirello/public/blob/master/SLA.md)

Package oversight makes a complete implementation of the Erlang supervision
trees.

Refer to: http://erlang.org/doc/design_principles/sup_princ.html

go get [-u -f] cirello.io/oversight

http://godoc.org/cirello.io/oversight


## Quickstart
```
supervise := oversight.New(
	oversight.Processes(func(ctx context.Context) error {
		select {
		case <-ctx.Done():
			return nil
		case <-time.After(time.Second):
			log.Println(1)
		}
		return nil
	}),
)

ctx, cancel := context.WithCancel(context.Background())
defer cancel()
if err := supervise.Start(ctx); err != nil {
	log.Fatal(err)
}
```