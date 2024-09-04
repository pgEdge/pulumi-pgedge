package main

import (
    "context"
    "fmt"
    "log"
    "os"
    "github.com/jackc/pgx/v4"
)

func main() {
    os.Setenv("PGHOST", "")
    os.Setenv("PGUSER", "")
    os.Setenv("PGDATABASE", "")
    os.Setenv("PGSSLMODE", "")
    os.Setenv("PGPASSWORD", "")

    ctx := context.Background()

    config, err := pgx.ParseConfig("")
    if err != nil {
        log.Fatalf("Failed to parse configuration: %v", err)
    }

    conn, err := pgx.ConnectConfig(ctx, config)
    if err != nil {
        log.Fatalf("Unable to connect to database: %v", err)
    }
    defer conn.Close(ctx)

    rows, err := conn.Query(ctx, "SELECT node_name FROM spock.node")
    if err != nil {
        log.Fatalf("Failed to execute query: %v", err)
    }
    defer rows.Close()

    var nodeNames []string
    for rows.Next() {
        var nodeName string
        if err := rows.Scan(&nodeName); err != nil {
          log.Fatalf("Failed to scan row: %v", err)
        }
        nodeNames = append(nodeNames, nodeName)
    }
    fmt.Println(nodeNames)
}
