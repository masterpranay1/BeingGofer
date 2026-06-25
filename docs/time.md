**Handling Time in Golang**

- we will use the core **time** package
- `presentTime := time.Now()`
- We can add format to the time.
- `presentTime.Format("01-02-2006")`
- Yes we have remember this exact date
- `fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))`
- This is very different that other languages. Monday above is like compulsory thing.

**Creating a new date**

```go
createdAt := time.Date(2020, time.August, 10, 23, 23, 0, 0, time.UTC)
// first is year
// then comes the month. Month is a type of time.month 
// then date, hour, minute, second, nanosec
// last one is location, we have bunch option for location like UTC, Local
```

There are many methods in the time package. Definitely worth giving a doc read.