package main

import (
  "flag"
  "fmt"
  "os"
  "redis"
  "strings"
  "syscall"
)

func printVersion() {
  fmt.Println(fmt.Sprintf("%s 0.0.1", os.Args[0]))
  os.Exit(0)
}

func listConfig(client *redis.Client, key string) {
  value, err := client.Hgetall(key)

  if err != nil {
    fmt.Fprintln(os.Stderr, err)
    os.Exit(1)
  }

  for key, value := range value.StringMap() {
    fmt.Println(fmt.Sprintf("%s=%s", key, value))
  }
}

func addConfig(client *redis.Client, key string, nameAndValue string) {
  parts := strings.Split(nameAndValue, "=")
  name := parts[0]
  value := strings.Join(parts[1:len(parts)], "=")
  _, err := client.Hset(key, name, value)

  if err != nil {
    fmt.Fprintln(os.Stderr, err)
    os.Exit(1)
  }
}

func removeConfig(client *redis.Client, key string, name string) {
  _, err := client.Hdel(key, name)

  if err != nil {
    fmt.Fprintln(os.Stderr, err)
    os.Exit(1)
  }
}

func run(client *redis.Client, key string, command string) {
  value, err := client.Hgetall(key)

  if err != nil {
    fmt.Fprintln(os.Stderr, err)
    os.Exit(111)
  }

  env := make([]string, 0)

  for key, value := range value.StringMap() {
    env = append(env, fmt.Sprintf("%s=%s", key, value))
  }

  syscall.Exec("/bin/sh", []string{"/bin/sh", "-c", command}, env)
}

func main() {
  version := flag.Bool("version", false, "Print version and exit")
  list := flag.Bool("list", false, "List config vars")
  netaddr := flag.String("netaddr", "tcp:127.0.0.1:6379", "Redis netaddr (e.g. tcp:120.0.0.1:6379")
  dbIndex := flag.Int("db", 0, "Redis database index")
  name := flag.String("name", "default", "Config name")
  runCommand := flag.String("run", "", "Command to run")
  remove := flag.String("remove", "", "Config var to remove")
  add := flag.String("add", "", "Config var to add")
  flag.Parse()

  if *version {
    printVersion()
  } else {
    client := redis.New(*netaddr, *dbIndex, "")
    key := fmt.Sprintf("redis-env:%s", *name)
    fmt.Println(key)

    if len(*runCommand) > 0 {
      run(client, key, *runCommand)
    } else if len(*remove) > 0 {
      removeConfig(client, key, *remove)
    } else if len(*add) > 0 {
      addConfig(client, key, *add)
    } else if *list {
      listConfig(client, key)
    }
  }
}
