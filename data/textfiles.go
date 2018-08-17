package data 

const (
Sqs = `sqs:
  classes:
    - standard
    - fifo
  free_tier:
    requests:
      value: 1000000
      period: monthly
  limits:
    request_payload:
      value: 64
      units: KB
  price:
    requests:
      per: 1000000
      standard: 0.4
      fifo: 0.5
    data:
      period: monthly
      in:
        type: flat
        price: 0
      out:
        type: ladder
        units: GB
        bands:
          - from: 0
            to: 1
            price: 0.0
          - from: 1
            to: 10000
            price: 0.09
          - from: 10000
            to: 40000
            price: 0.085
          - from: 40000
            to: 150000
            price: 0.07
          - from: 150000
            to: 500000
            price: 0.05
          - from: 500000
            poa: true


`
)
