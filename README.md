# Simple go server playground project

## starting route
```
http://localhost:8888
```


## implemented routes

```
GET:         /
GET:         /bikes
POST:        /bikes
GET:         /bikes/{id}
PUT:         /bikes/{id}:soldout  //custom method
DELETE:      /bikes/{id}

```

## response structure

```
{
    "data":
    {
        "id":1,
        "desc":"Haro Midway",
        "frame":"3 tubes Crmo, Internal HT \u0026 mid BB - 20.5 \u0026 21 TT",
        "gearing":"25/9",
        "customerPrice":{"currency_code":"CHF","units":459,"nanos":50},
        "soldOut":false
    }
}

```

```
{
    "data":[
        {"data":
            {"id":1,"desc":"Haro Midway","frame":"3 tubes Crmo, Internal HT \u0026 mid BB - 20.5 \u0026 21 TT","gearing":"25/9","customerPrice":{"currency_code":"CHF","units":459,"nanos":50},"soldOut":false}
        },
        {"data":
            {"id":2,"desc":"Haro CK AM 2019","frame":"3 tubes Crmo, Internal HT \u0026 mid BB - 20.5 \u0026 21 TT","gearing":"25/9","customerPrice":{"currency_code":"CHF","units":759},"soldOut":true}
        },
        {"data":
            {"id":3,"desc":"Haro Extreme 2019","frame":"All tubes Crmo, Internal HT \u0026 mid BB - 20.5 \u0026 21 TT","gearing":"25/9","customerPrice":{"currency_code":"CHF","units":1250},"soldOut":false}
        }
    ]
}

```