# greenpass

This repository contains a library and a command line utility in Go
for parsing QR codes containing EU Digital COVID certificates.

## Command line

Run the `greenpass` utility under [cmd/greenpass](cmd/greenpass/). It currently
supports PNG and JPEG.

Usage:
```
$ ./greenpass
Usage: ./greenpass <file name>
```

Example:

Download an example QR code, like
https://github.com/ehn-dcc-development/dcc-testdata/blob/main/IE/png/1_qr.png .

Then run
```
$ ./greenpass 1_qr.png
map[interface {}]interface {}{
    -260: map[interface {}]interface {}{
        0x1: map[interface {}]interface {}{
            "dob": "1988-06-07",
            "nam": map[interface {}]interface {}{
                "fn":  "Bloggs",
                "fnt": "BLOGGS",
                "gn":  "Jane",
                "gnt": "JANE",
            },
            "v": []interface {}{
                map[interface {}]interface {}{
                    "ci": "URN:UVCI:01:IE:52d0dc929c884cf8998a7987f0b9d863#2",
                    "co": "IE",
                    "dn": uint64(0x1),
                    "dt": "2021-05-06",
                    "is": "HSE",
                    "ma": "ORG-100030215",
                    "mp": "EU/1/20/1528",
                    "sd": uint64(0x2),
                    "tg": "840539006",
                    "vp": "1119349007",
                },
            },
            "ver": "1.0.4",
        },
    },
    0x1: "IE",
    0x4: uint64(0x60c71a90),
    0x6: uint64(0x60bdced4),
}
```

## Library

You can use `Read` or `Decode`.

* `Read(r io.Reader)` receives a reader, like a file handle, and returns the
  COVID certificate object. It calls `Decode` under the hood.
* `Decode` receives an `image.Image` and returns the COVID certificate object.

Examples:

`Read`:
```
fd, _ := os.Open("myqrcode.png")
gp, _ := greenpass.Read(fd)
fmt.Printf("%+v", gp)
```

`Decode`:
```
img, _, _ := image.Decode(reader)
gp, _ := greenpass.Decode(img)
fmt.Printf("%+v", gp)
```

## Useful links

* [Technical Specifications for EU Digital COVID Certificates](https://ec.europa.eu/health/sites/default/files/ehealth/docs/covid-certificate_json_specification_en.pdf)
* [Electronic Health Certificate specification](https://github.com/ehn-dcc-development/hcert-spec)
* [Digital COVID Certificate JSON schema](https://github.com/ehn-dcc-development/ehn-dcc-schema)
* [Digital COVID Certificate test data](https://github.com/ehn-dcc-development/dcc-testdata)
* [Online Digital COVID Certificate decoder](https://ehealth.vyncke.org/) with a
  technical explanation

