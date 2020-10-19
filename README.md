# clickhz
example sqbuf with clickhouse

```

CREATE TABLE IF NOT EXISTS applogs (
    errType     String,
    caller      String,
    message     String,
    dt          DateTime
)
ENGINE = MergeTree()
ORDER BY (errType, dt)
SETTINGS index_granularity = 8192

```
