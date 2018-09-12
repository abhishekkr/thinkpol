
**[TBD]**
## local adapter

#### request body constructs:

_global_
* type: `joycamp` (if people want more kinds)

_global_
* ProcDef: joycamp Proc def

_global_
* webhook: to call once job is done 
> * webhook-type: `http(s)`
> * webhook-method (if webhook-type is `http(s)`)
> * webhook-url-params (if webhook-type is `http(s)`)
> * webhook-body (if webhook-type is `http(s)`)
> * webhook-headers (if webhook-type is `http(s)`)
> * webhook-delay

_global_
* cron: nil if provided is not a scheduler job, else timer
> * cron-times: `<int>`; negative means upper count on failures before discard; zero means infinite

> sample

```
{
  "type": "joycamp",
  "proc-def": {
    <raw-json like joycamp proc def>
  },
  "webhook": {
    "type": "http",
    "url": "http://127.0.0.1:9988/webhook",
    "method": "post",
    "params": {
      "status": "$thinkpol.proc.status"
    },
    "body": <raw-json like "$thinkpol.proc.stdout">,
    "headers": {
      "X-THINKPOL": "v0.0.1"
    },
    "delay": "1" //in seconds
  },
  "cron": {
    "timer": "*/5 * * * *",
    "times": 2
  }
}
```

---
