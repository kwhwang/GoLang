import (
    "github.com/alouca/gosnmp"
    "log"
)

s, err := gosnmp.NewGoSNMP("61.147.69.87", "public", gosnmp.Version2c, 5)
if err != nil {
    log.Fatal(err)
}
resp, err := s.Get(".1.3.6.1.2.1.1.1.0")
if err == nil {
    for _, v := range resp.Variables {
        switch v.Type {
        case gosnmp.OctetString:
            log.Printf("Response: %s : %s : %s \n", v.Name, v.Value.(string), v.Type.String())
        }
    }
}