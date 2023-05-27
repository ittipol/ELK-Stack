GET orders/_search
{
  "size": 0,
  "aggs": {
    "purchased_at_date_histogram": {
      "date_histogram": {
        "field": "purchased_at",
        "format": "dd/MM/yyyy HH:ss:mm", 
        "calendar_interval": "quarter",
        //"fixed_interval": "30d",
        "keyed": true,
        "missing": "01/01/2000 00:00:00"
      },
      "aggs": {
        "total_amount_stats": {
          "stats": {
            "field": "total_amount"
          }
        }
      }
    }
  }
}