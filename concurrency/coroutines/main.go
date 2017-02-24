// http://www.golangpatterns.info/concurrency/coroutines
package main

import "fmt"

func integers () chan int {
    yield := make (chan int);
    count := 0;
    go func () {
        for {
            yield <- count;
            count++;
        }
    } ();
    return yield
}

func main(){
  resume := integers();
  generateInteger := func() int{
    return <-resume
  }

  for i := 0; i < 3; i++{
    fmt.Println(generateInteger())
  }
}
