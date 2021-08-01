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
$ go run . 1_qr.png
Name             : Jane Bloggs
Date of Birth    : 1988-06-07
Status           : partially vaccinated (1 of 2 doses)
Vaccination date : 2021-05-06
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

## Generated code

Some data structures are generated via
[`tools/generate_types`](tools/generate_types). If you need to regenerate them
(e.g. because they were updated), just run `go generate` in the top level
directory.

## Useful links

* [Technical Specifications for EU Digital COVID Certificates](https://ec.europa.eu/health/sites/default/files/ehealth/docs/covid-certificate_json_specification_en.pdf)
* [Electronic Health Certificate specification](https://github.com/ehn-dcc-development/hcert-spec)
* [Digital COVID Certificate JSON schema](https://github.com/ehn-dcc-development/ehn-dcc-schema)
* [Digital COVID Certificate test data](https://github.com/ehn-dcc-development/dcc-testdata)
* [Online Digital COVID Certificate decoder](https://ehealth.vyncke.org/) with a
  technical explanation
* [Reverse engineering digital COVID-19 certificates](https://harrisonsand.com/posts/covid-certificates/)
