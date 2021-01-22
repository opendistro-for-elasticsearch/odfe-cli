# Anomaly Detection (ad)

## Description

Use the Anomaly Detection commands to create, configure, and manage detectors.


## Available commands

* [create](./ad.md#1-create)
* [get](./ad.md#2-get)
* [update](./ad.md#3-update)
* [start](./ad.md#4-start)
* [stop](./ad.md#5-stop)
* [delete](./ad.md#6-delete)



## 1. create

### Description

Create detectors based on a local JSON files

### Synopsis

```
odfe-cli ad create json-file-path ... [flags]
```

### Flags

`--generate-template`  Output sample detector configuration

### Example

The following example shows creation of e-commerce sample detector.

#### Step 1: Generate a sample configuration

The following example shows how to generate a template formatted in JSON for the `--generate-template` parameter.

```
$ odfe-cli ad create --generate-template
{
  "name": "Detector Name",
  "description": "A brief description",
  "time_field": "",
  "index": [],
  "features": [
    {
      "aggregation_type": [
        "count"
      ],
      "enabled": false,
      "field": []
    }
  ],
  "filter": {},
  "interval": "10m",
  "window_delay": "1m",
  "start": false,
  "partition_field": ""
} 

```

| Field Name  | Description |
| ----------------  | ------------- | --------------- |
| 1.0               | [Anomaly Detection](https://opendistro.github.io/for-elasticsearch-docs/docs/ad/)  | 1.12.0 |

#### Step 2: Save the template as new file

```
$ odfe-cli ad create --generate-tempalte > ecommerce-sample-detector.json
```

#### Step 3: Fill in template file

Open the parameter template file in your text editor and remove any of the parameters 
that you don't need. For example, you might update the template down to the following. 

```
{
  "name": "e-commerce-sample-detector",
  "description": "E-commerce data",
  "time_field": "utc_time",
  "index": ["kibana_sample_data_ecommerce*"],
  "features": [{
    "aggregation_type": ["sum"],
    "enabled": true,
    "field":["total_quantity"]
  }],
  "filter": {
    "bool": {
      "filter": {
        "term": {
          "currency": "EUR"
        }
    }}
  },
  "interval": "5m",
  "window_delay": "1m",
  "start": true
}
```

#### Step 4: Create the detector from command line on prod cluster

```
$ odfe-cli ad create ecommerce-sample-detector.json --profile prod
100.00% [=================================================] 1 / 1
Successfully created 1 detector(s)
```

## 2. get

### Description

Get detectors based on a list of IDs, names, or name regex patterns.

### Synopsis

```
$ odfe-cli ad get detector_name ... [flags] 
```

### Flags

`--id Detector ID`  Input is detector ID


### Example

The following example display `ecommerce-sample-detector` 

```
$ odfe-cli ad get ecommerce-sample-detector
```
The following example display detector with id `mqLmNertifkEsffEff` 

```
$ odfe-cli ad get --id mqLmNertifkEsffEff
```

## 3. update

### Description

To begin, use `odfe-cli ad get detector-name > detector_to_be_updated.json` to download the detector.
Modify the file, and then use update command to update the detector.

### Synopsis

```
$ odfe-cli ad update file-path ... [flags]
```

### Flags

`--force`    Stop detector and update forcefully

`--start`   Start detector if update is successful

### Example

The following example will update ecommerce-sample-detector.
#### Step 1: Download ecommerce-sample-detector.json
```
$ odfe-cli ad get ecommerce-sample-detector > ecommerce-sample-detector.json
$ cat ecommerce-sample-detector.json
{
  "ID": "0r1YA3QBCzk8KIr9KdNn",
  "name": "ecommerce-sample-detector",
  "description": "Ecommerce data",
  "time_field": "utc_time",
  "indices": [
    "kibana_sample_data_ecommerce*"
  ],
  "features": [
    {
      "feature_name": "sum_total_quantity",
      "feature_enabled": true,
      "aggregation_query": {
        "sum_total_quantity": {
          "sum": {
            "field": "total_quantity"
          }
        }
      }
    }
  ],
  "filter_query": {
    "bool": {
      "must": [
        {
          "bool": {
            "filter": [
              {
                "term": {
                  "currency": {
                    "value": "EUR",
                    "boost": 1.0
                  }
                }
              }
            ],
            "adjust_pure_negative": true,
            "boost": 1.0
          }
        }
      ],
      "adjust_pure_negative": true,
      "boost": 1.0
    }
  },
  "detection_interval": "5m",
  "window_delay": "1m",
  "last_update_time": 1597783943520,
  "schema_version": 0
}

```
#### Step 2: Manually make changes to ecommerce-sample-detector.json
Update detection_interval to 2m

```
$ cat ecommerce-sample-detector.json
{
  "ID": "0r1YA3QBCzk8KIr9KdNn",
  "name": "ecommerce-sample-detector",
  "description": "Ecommerce data",
  "time_field": "utc_time",
  "indices": [
    "kibana_sample_data_ecommerce*"
  ],
  "features": [
    {
      "feature_name": "sum_total_quantity",
      "feature_enabled": true,
      "aggregation_query": {
        "sum_total_quantity": {
          "sum": {
            "field": "total_quantity"
          }
        }
      }
    }
  ],
  "filter_query": {
    "bool": {
      "must": [
        {
          "bool": {
            "filter": [
              {
                "term": {
                  "currency": {
                    "value": "EUR",
                    "boost": 1.0
                  }
                }
              }
            ],
            "adjust_pure_negative": true,
            "boost": 1.0
          }
        }
      ],
      "adjust_pure_negative": true,
      "boost": 1.0
    }
  },
  "detection_interval": "2m",
  "window_delay": "1m",
  "last_update_time": 1597783943520,
  "schema_version": 0
}

```

#### Step 3: Update detector from command line

```
$ odfe-cli ad update ecommerce-sample-detector.json
```

#### Step 3: Update detector by  force  and restart.

```
$ odfe-cli ad update-detectors ecommerce-sample-detector.json --force --start
100.00% [=================================================] 1 / 1
Successfully updated and restarted 1 detector(s)
```

## 4. start

### Description

Start detectors based on list of IDs, names, or name regex patterns.

### Synopsis

```
$ odfe-cli ad start detector_name ... [flags] 
```

### Flags

`--id Detector ID`  Input is detector ID


### Example

The following example starts `ecommerce-sample-detector`  and `ecommerce-sample-detector1`

```
$ odfe-cli ad start ecommerce-sample-*
2 detector(s) matched by name ecommerce-sample-*
ecommerce-sample-detector
ecommerce-sample-detector1
odfe-cli will start above matched detector(s). Do you want to proceed? please type (y)es or (n)o: yes
100.00% [===============================================] 2 / 2
```
The following example starts detector with id `mqLmNertifkEsffEff` 

```
$ odfe-cli ad start --id mqLmNertifkEsffEff
100.00% [===============================================] 1 / 1
```

## 5. stop

### Description

Stop detectors based on list of IDs, names, or name regex patterns.

### Synopsis

```
$ odfe-cli ad stop detector_name ... [flags] 
```

### Flags

`--id Detector ID`  Input is detector ID


### Example

The following example stop `ecommerce-sample-detector`  and `ecommerce-sample-detector1`

```
$ odfe-cli ad stop ecommerce-sample-*
2 detector(s) matched by name ecommerce-sample-*
ecommerce-sample-detector
ecommerce-sample-detector1
odfe-cli will stop above matched detector(s). Do you want to proceed? please type (y)es or (n)o: yes
100.00% [===============================================] 2 / 2
```
The following example stops detector with id `mqLmNertifkEsffEff` 

```
$ odfe-cli ad stop --id mqLmNertifkEsffEff
100.00% [===============================================] 1 / 1
```

## 6. delete

### Description

Delete detectors based on list of IDs, names, or name regex patterns.

### Synopsis

```
$ odfe-cli ad delete detector_name ... [flags]
```

### Options

`--force`   Delete the detector even if it is running
`--id`          Input is detector ID


### Example

The following example deletes `ecommerce-sample-detector1` but failed to delete  `ecommerce-sample-detector`  since it is still running.

```
$ odfe-cli ad delete ecommerce-sample-*
2 detector(s) matched by name ecommerce-sample-*
ecommerce-sample-detector
ecommerce-sample-detector1
odfe-cli will delete above matched detector(s). Do you want to proceed? please type (y)es or (n)o: yes
50.00% [=======>__________________________________________] 1 / 2
delete command failed.
failed to delete 1 following detector(s)
ecommerce-sample-detector     Reason: Detector job is running: 3r1YA3QBCzk8KIr9KtOp
```
The following example deletes `ecommerce-sample-detector` forcefully since detector job is running.

```
$odfe-cli ad delete ecommerce-sample-* --force
 1 detector(s) matched by name ecommerce-sample-*
ecommerce-sample-detector
odfe-cli will delete above matched detector(s). Do you want to proceed? please type (y)es or (n)o: yes
100.00% [===============================================] 1 / 1
```
 The following example delete detector with id `mqLmNertifkEsffEff` 

```
$ odfe-cli ad delete --id mqLmNertifkEsffEff
100.00% [===============================================] 1 / 1
```

