Fanshim
=======

Turn a Pimoroni fanshim on and off according the to the temperature

```
version: "2"
services:
  fanshim:
    image: testingtony/fanshim
    privileged: true
    environment: 
      - ON_TEMP=65
      - OFF_TEMP=55
```