package main

import (
    "fmt"
    "sync"
    "time"
    "crypto/sha256"
)

var ( appVersion = "2.4" )

func 8wUCilXjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 153
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8BQ7KbZSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 75
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RlAtcIceWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 97
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dAO12EBHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 22
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WUF4sxRGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 290
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mXEGSFsBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 95
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aduGamBHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 293
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func m0MFJbcUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 11
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sMN6f7YOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 145
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SWh8JpAtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 45
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MHFeKfP5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 300
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bC7LOapxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 128
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4Mq8Q2FmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 40
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func foHtX9YRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 207
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wlVPKkREWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 79
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func M5me5GYOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 141
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ROjzaFLjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 24
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QoTFfNwMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 10
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bF5DICFCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 169
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FOaCUCsEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 77
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kp5azQCYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 205
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UGVmLM95Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 111
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zy41C5qKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 28
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GZ8pN9QIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 53
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qBKA3kSyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 193
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zqTVX95DWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 296
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func T4YckCESWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 164
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func E7s1mtGWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 293
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BCFBVpbYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 235
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NSljfPuOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 131
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aMsH764fWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8XG2mpR1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 36
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OUteViO7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 263
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YZp5UOOMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 286
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aXPJwY9MWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 135
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func POLPISWAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 123
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func umtxt0wRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 234
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Y7aUDZpDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 182
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TEVkrw1DWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 202
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KS74RAX6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 255
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1N0Eqxs8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 100
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ISlsxf7FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 11
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IRNNG8iWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 271
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jhaETBSlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 58
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7xKPR9qAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 192
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pjhzDNqzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 271
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xFAZwWMhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 275
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HM9hmuwOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 284
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dZEZSUM1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 132
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qFSnhUqzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 130
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LTYOVlNfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 143
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mFi48kInWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 13
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ALM1WWRTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 111
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ttYZ3uwEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 217
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func O4f9m7tRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 142
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZHnbtcHHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HNzzBFmTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 102
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YPlGGWqzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 37
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func po9u5onIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 201
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func R8m7woCTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 289
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VWfndw09Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 40
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rtzO6PCkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 57
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JTDL6rM4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 55
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pERiPzoxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 298
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func h3ZtRbYyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 45
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jh2QUiDxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 63
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8TNfTRO0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 129
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nqoiC0z7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 173
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qLA7YKilWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 291
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HlZmQkEaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 168
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AO6X3sHFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 199
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WdTJ2PtWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 193
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gSEJwzTqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 36
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gLqA5sLRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 57
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vPUDH1myWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 262
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pYjsqTVzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 74
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BACBDOE5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 34
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pZcFWxyxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 14
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZEM59bm6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 252
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nuuBASfXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 83
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func q9PztadqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 296
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YKeVK4gCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 290
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LJHjWfVbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 259
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hUWLhhN5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 135
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UaOCriqjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 82
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SIEYgmTKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 290
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7MmjgcbKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 243
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sRyKW9I0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 134
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Hm8PPP0cWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 248
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZikTkOt7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 117
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dEGozOlgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 296
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6RYru2oMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 277
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func R9fJ4CemWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 247
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wUpujEh5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 153
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rDbBYb2qWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 206
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xAZoURtpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 14
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7Mwh0spGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 113
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hoB1t5s8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 81
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2OjgCn6SWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yobKuEKJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 35
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func x3zKsisaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 187
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KlEacL62Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 253
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hj2GWXhPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 17
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TzO7XHl7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 77
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pD9Tmd1HWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 22
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0vkrw2A0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 48
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CSYSIp3FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 278
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4Ui9I9vKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 93
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dhi7QQGPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 46
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func l6zWv9LXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 97
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LoZgq6TzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 156
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func w9GJfZfEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 66
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TesRDmMdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 62
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9rtGyBWWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 203
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rLiaFahCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 237
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SrYAWSqnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 264
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func upGmKAJHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 181
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9ChyMnRjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 290
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4nrSDaSkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 231
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BAq82OR2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 285
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4fkpzorhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 169
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8tdoK1zPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 171
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3ERZg9PRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 181
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dXpACtyoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 234
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UjVExKJmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 11
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QYBHtvBwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 97
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EgLP7DwJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 201
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6APoAKDyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 68
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hvCeZcc6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 128
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jtCZsoCgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 293
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eWTm4exWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 276
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2qd4sUX2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 190
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MB1SmL2EWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 61
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hC6Kz0zoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 220
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RjyKIdMtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func d2o4zRvPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 117
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fJl3meFhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 165
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7xE8epWMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 278
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Nkf8DeW1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 274
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ExffZZqfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 146
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func houwjrNNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 138
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YFug6lGCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 206
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kdreb0UmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 42
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NQ3BYU3NWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 17
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DJvswKm4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 111
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oUr7UUujWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 99
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sPnEwVnyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 51
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gl5rwASsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 122
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OZiNrd2LWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 78
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JRgSXhU8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 53
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Txqa0yDVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 87
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1EDBOd2kWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 292
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VmwV1WmCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 298
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1GujSl5QWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 267
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VBtWq0R1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 135
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZYkoVcPiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 48
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 88R5N8H4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 160
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sIm4dOHoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 133
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func e3OyxJpaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Z6C7RE1jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 259
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xL5XalblWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 295
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DzyNstB9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 264
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1Giu0KP7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 26
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func i6gYeTP3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 276
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VFlGnnbhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 176
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SEDMDNGQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 100
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aTNriek2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 296
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xAanwGcSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 42
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 066H8F0gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 34
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wkqmVUJOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 120
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WRtpIIDNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 15
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func C0Zw6AvpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 250
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5WwOEaHBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 107
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IRCee3UeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 119
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zTFPKHE0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 218
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sB5Jl0NyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 64
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func d1czFXyWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 165
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func V6BhkQEVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 52
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6qXdTmy0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 119
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GJaITxYUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 224
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WTO3dSzbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 17
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RlWaMd2eWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 105
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func l5hDTUW6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 111
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pqESETSoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 289
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NqTpHEqkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 281
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rGXmtAMmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 176
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Bgo6IesYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 97
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KR5gJaPgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 13
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2DwO7DuEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 97
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Qz7gTWYZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 141
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Vf2A1iAHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 251
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uSG9ThNVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 250
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0qCAMxLOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 129
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UzpuJohTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 47
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YzKy6PkWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 212
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func N8uSasDEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 294
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gVc9RBqwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 288
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AxgdTONNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 82
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func w6jOQu83Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 215
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func J2alckHxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 187
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func V8Ml2haBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 289
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Z37LqN3sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 63
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kXa2BsdhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 53
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3VtWCI09Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 285
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZQc8b59mWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 244
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2TYyJtz4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 244
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lm58AHOnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 202
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QuculS9qWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 216
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oVarcvfbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 42
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nNoBUsOOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 31
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XAD49tN5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 217
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TE2VRB10Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 264
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yJSsfut3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 142
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nbPcNPTLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 206
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4G6dqzgSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 219
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xxNf9jTzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 297
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1kLgzFR3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 145
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lo79CFs4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 53
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zDv1m0OMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 216
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DXWfLGzzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 223
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gy03qg5JWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 225
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Fh1tOWgxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 113
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func grgIZjTyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 11
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1iCaa7KyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 43
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dL1ugKtDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 171
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func f9yoVllBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 122
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MmUWg4rTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 72
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gt2Od0AiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 100
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rjXGkhXCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 99
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6rKvtF7mWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 33
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sgNGLOCeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 28
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ruHYfUdGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 135
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WfEjPQcxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 259
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Uv9ftwwOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 215
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OndFKhDwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 288
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bvFreagRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 263
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Xe3WBx8EWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 63
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0Lgs08mfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 196
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gCijSRH9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 55
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XHS3YTOAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 206
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tUncIbzNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 246
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qnGzDatrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 81
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pGK4zwegWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 35
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qMb93sRmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 218
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func p229BYfaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 34
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QoMdOUXTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 140
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xIQXkF1VWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 127
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func O1CVZlMpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 10
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8sfLuwD0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 72
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mgvctxmUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 31
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4vBsKi6FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 258
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1YFzkuJPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 278
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XuPqFqENWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 17
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9MJRFwS5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 188
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func teTV3sLiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 101
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

