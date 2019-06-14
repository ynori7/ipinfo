# IpInfo
This is an API interface for checking your IP address, looking up an IP address, and for DNS lookup. See the [IpInfo Website](https://www.ip-info.xyz) for more details.

# *Running Locally*

Note: Depends on the [GeoLite2](https://dev.maxmind.com/geoip/geoip2/geolite2/) database.

After downloading the source and the GeoLite2 database, simply start the service, optionally specifying the path to the database:

```
go run main.go --geo-db-path "./GeoLite2-City.mmdb" --server-host ":8081"
```

# *Endpoints*

## *WhatsMyIp*
This endpoint will return data about the IP of the requester

### Request
```
GET  /ip
```

### Response
```
{
  "ip": "111.11.11.111",
  "geolocation": {
    "country_code": "DE",
    "country": "Germany",
    "city": "Garching bei Munchen",
    "lat": 48.249,
    "long": 11.651,
    "timezone": "Europe/Berlin"
  }
}
```

### *Errors*
The following errors may be returned:

HTTP code|error id
---|---
400|invalid_ip
500|internal_server_error

### CURL
```
curl http://localhost:8081/ip
```

## *IpLookup*
This endpoint will return data about the specified IP

### Request
```
GET  /ip/{ipAddress}
```

### Response
```
{
  "ip": "198.252.206.211",
  "hostnames": [
    "stackoverflow.com."
  ],
  "geolocation": {
    "country_code": "US",
    "country": "United States",
    "lat": 37.751,
    "long": -97.822,
    "timezone": "America/Chicago"
  }
}
```

### *Errors*
The following errors may be returned:

HTTP code|error id
---|---
400|missing_ip
400|invalid_ip
500|internal_server_error

### CURL
```
curl http://localhost:8081/ip/198.252.206.211
```

## *HostLookup*
This endpoint will return data about the specified IP

### Request
```
GET  /host/{hostname}
```

### Response
```
{
  "ips": [
    "151.101.1.69",
    "151.101.129.69",
    "151.101.65.69",
    "151.101.193.69"
  ],
  "geolocation": {
    "country_code": "US",
    "country": "United States",
    "lat": 37.751,
    "long": -97.822,
    "timezone": "America/Chicago"
  }
}

```

### *Errors*
The following errors may be returned:

HTTP code|error id
---|---
400|missing_hostname
400|invalid_hostname
404|not_found
500|internal_server_error

### CURL
```
curl http://localhost:8081/host/stackoverflow.com
```