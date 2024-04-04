package main

import (
    "fmt"
    "net"
    "time"
)

func main() {
    // Mulai mendengarkan pada port 8080
    ln, err := net.Listen("tcp", ":8080")
    if err != nil {
        fmt.Println("Error listening:", err)
        return
    }
    defer ln.Close()

    fmt.Println("Server is listening on port 8080...")

    for {
        // Terima koneksi
        conn, err := ln.Accept()
        if err != nil {
            fmt.Println("Error accepting connection:", err)
            continue
        }

        // Set timeout untuk membaca dan menulis
        conn.SetReadDeadline(time.Now().Add(10 * time.Second))
        conn.SetWriteDeadline(time.Now().Add(10 * time.Second))

        // Proses koneksi
        go handleConnection(conn)
    }
}

func handleConnection(conn net.Conn) {
    defer conn.Close()

    // Baca data dari koneksi
    buffer := make([]byte, 1024)
    _, err := conn.Read(buffer)
    if err != nil {
        fmt.Println("Error reading:", err)
        return
    }

    // Cetak data yang diterima
    fmt.Println("Received:", string(buffer))

    // Kirim balik pesan ke klien
    response := []byte("Hello from server")
    _, err = conn.Write(response)
    if err != nil {
        fmt.Println("Error writing:", err)
        return
    }
}
