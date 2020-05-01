package main

import (
    "time"
    "math/rand"
)

<<<<<<< HEAD
func genRandDie() int8 { // this defines the return type
=======
func genRandDie()  {
>>>>>>> bc86b1679638b5b488fb8e1f6f10b4253e76655d
    // Set the seed so we get a unique random value every time
    rand.Seed(time.Now().UnixNano())
    return rand.Intn(5) + 1 // (0 - 5) + 1
}
