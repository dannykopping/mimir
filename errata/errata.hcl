version = "0.1"

options {
  prefix   = "err-mimir-"
  base_url = "https://grafana.com/docs/mimir/latest/errata/"
}

error "label-name-too-long" {
  message    = "Received a series whose label name length exceeds the limit, label: '%.200s' series: '%.200s'"
  categories = ["validation", "label"]
  guide      = file("guides/label-name-too-long.md")
  args       = [
    arg("label", "string"),
    arg("series", "string"),
  ]
  labels     = {
    http_response_code : 400
    level : "warning"
    flag : "validation.max-length-label-name"
  }
}

error "label-value-too-long" {
  message    = "received a series whose label value length exceeds the limit, value: '%.200s' (truncated) series: '%.200s'"
  categories = ["validation", "label"]
  guide      = file("guides/label-value-too-long.md")
  args       = [
    arg("label", "string"),
    arg("series", "string"),
  ]
  labels     = {
    http_response_code : 400
    level : "warning"
    flag : "validation.max-length-label-value"
  }
}