package main

import (
            "fmt"
            "math/rand"
            "time"
            "sync"
)

const letters = "abcdefghijklmnopqrstuvwxyz"
var chans [30] chan string
var result string
var wg sync.WaitGroup
var isFinished bool = false

func allocateChans() {
    for i := range chans {
        chans[i] = make(chan string)
    }
}

func generateRandomString(n int) {
    s := make([]byte, n)

    seed := rand.NewSource(time.Now().UnixNano())
    generator := rand.New(seed)

    for i := range s{
        s[i] = letters[generator.Intn(len(letters))]
    }

    fmt.Println("String inicial: " + string(s))
    chans[0] <- string(s)
}

func readBuffer(in chan string) []rune {
    s := <- in
    return []rune(s)
}

func updateString(str []rune, out chan string) {
    for i:= range str {
        if str[i] >96 && str[i]<123 {
            s := toUpper(str,i)
            time.Sleep(1 * time.Second)
            out <- s
            result = s
            break
            } else if i == len(str) -1  {
                isFinished = true
        }
    }
}

func readBufferAndUpdateString(rank int, in chan string, out chan string) {
    fmt.Println("Thread ", rank)

    for {
        r := readBuffer(in)
        updateString(r, out)
        if isFinished {
            finish(out)
            break
        }
    }
}


func finish(ch chan string) {
    close(ch)
    wg.Done()
}

func toUpper(str []rune, index int) string {
    str[index] = str[index] - 32
    return string(str)
}

func initiateRoutines() {
    var j int

    wg.Add(30)
    for i:=0;i<30;i++{
        j = (i+1)%30
        go readBufferAndUpdateString(i, chans[i], chans[j])
    }
    wg.Wait()
}

func main() {
    allocateChans()
    go generateRandomString(80)
    initiateRoutines()

    fmt.Println("Resultado final: " + result)
}