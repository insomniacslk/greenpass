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
&greenpass.CovidCertificate{
    Ver: "1.0.4",
    Nam: greenpass.PersonName{Fn:"Bloggs", Fnt:"BLOGGS", Gn:"Jane", Gnt:"JANE"},
    DoB: "1988-06-07",
    V:   {
        {Tg:"840539006", Vp:"1119349007", Mp:"EU/1/20/1528", Ma:"ORG-100030215", Dn:1, Sd:2, Dt:"2021-05-06", Co:"IE", Is:"HSE", Ci:"URN:UVCI:01:IE:52d0dc929c884cf8998a7987f0b9d863#2"},
    },
    T:  nil,
    R:  nil,
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
* [Reverse engineering digital COVID-19 certificates](https://harrisonsand.com/posts/covid-certificates/)
