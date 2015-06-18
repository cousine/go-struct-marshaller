Overview
========

This is a small helper library to marshal a struct into a map and vice-versa.

Status
======

At the moment only marshaling into maps is implemented.

Installation
============

`go get github.com/cousine/go-struct-marshaller`

Usage
=====

    import "github.com/cousine/go-struct-marshaller"

    type MarshalStruct struct {
      Field      string `stmarsh:"field_name"`
      OtherField int    `strmarsh:"otherFieldName"`
    }

    func main() {
      ms := &MarshalStruct {
        Field: "Value",
        OtherField: "OtherValye",
      }

      st := structmarshaller.New(ms)
      var stmap map[string]interface{} = st.Encode()
    }
