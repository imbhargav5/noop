# noop :zap:

A no-op in golang. **It does nothing**.

<img src="http://www.richardcollison.net/wp-content/uploads/kfp-mr-ping-secret.jpg" alt="Nothing" />


## installation

```bash
go get github.com/bhargav175/noop
```

## Usage

```golang
import "github.com/bhargav175/noop"

// ...

var a func()

if want_to_do_something == true {

  a = func() {
    //do something
  }

}else{

  a = noop.Noop

}

a()

// ...

```
