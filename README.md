# CRC16 ITU algo

[![Build Status](https://travis-ci.org/qsun/crc16itu.svg?branch=master)](https://travis-ci.org/qsun/crc16itu)

This crc16/ITU algo is used in some GPS devices.

# Usage

```go
package main

import "github.com/qsun/crc16itu"

.....
.....


checksum := crc16itu.CRC16ITU([]byte{data[2], data[3], data[4], data[5]})
log.Println("Checksum: ", checksum)
```
