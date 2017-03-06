## Coroutines
Users of C#, Lua, Python, and others might notice the resemblance between Go's generators/iterators and coroutines.  Obviously, the name "goroutine" stems from this similarity.

The differences between coroutines and goroutines are:

goroutines imply parallelism; coroutines in general do not
goroutines communicate via channels; coroutines communicate via yield and resume operations
In general, goroutines are much more powerful than coroutines.  In particular, it is easy to port coroutine logic to goroutines and gain parallelism in the process.

For example, the following coroutine (in Lua) implements a simple generator:

```Lua
function integers ()
    local count = 0
    return function ()
        while true do
            yield (count)
            count = count + 1
        end
    end
end
...
function generateInteger ()
    return resume(integers)
end
...
generateInteger() => 0
generateInteger() => 1
generateInteger() => 2
```

The same can be done in Go, though a channel is used instead of a yield:

```
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
...
resume := integers();
func generateInteger () int {
    return <-resume
}

generateInteger() => 0
generateInteger() => 1
generateInteger() => 2
```
