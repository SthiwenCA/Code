package main

import (
    "fmt"
    "net"
    "time"
)

func main() {
    // Dial dengan timeout
    conn, err := net.DialTimeout("tcp", "localhost:8080", 5*time.Second)
    if err != nil {
        fmt.Println("Error connecting:", err)
        return
    }
    defer conn.Close()

    // Set timeout untuk menulis dan membaca
    conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
    conn.SetReadDeadline(time.Now().Add(10 * time.Second))

    // Kirim pesan ke server
    message := []byte("Hello from client")
    _, err = conn.Write(message)
    if err != nil {
        fmt.Println("Error writing:", err)
        return
    }

    // Baca balasan dari server
    buffer := make([]byte, 1024)
    _, err = conn.Read(buffer)
    if err != nil {
        fmt.Println("Error reading:", err)
        return
    }

    // Cetak balasan dari server
    fmt.Println("Response:", string(buffer))
}
