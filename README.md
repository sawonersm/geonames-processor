# Geonames Processor


## What is this?
This repository is a fast-way geolocations processor based on the world size database   [Geonames](https://www.geonames.org/)

## Structure

Database is structures in two tables:

### Features

A feature is a way of classificate locations. It is composed by:
- class: It is a high level classification. Country, road, city, mountain ...
- code: It is a more specific aggrupation.

For more information, please, visit the official [documentation](http://www.geonames.org/export/codes.html)

```sql
# Single result example
id: 100
class: H
code: OCN
name: ocean
description: one of the major divisions of the vast expanse of salt water covering part of the earth
created_at: 2021-07-07 08:18:24
updated_at: 2021-07-07 08:18:24
deleted_at: NULL
```

### Locations

A location is formed by an id, a feature, name and coordinates


```sql
# Single result example
geoname_id: 10105996
name: Las Palmeras
latitude: 37.6865501
longitude: -0.7360200
feature: 348
created_at: 2021-07-07 08:34:19
updated_at: 2021-07-07 08:34:19
deleted_at: NULL
```

## How does it work

Nowadays we can just import directly in database by choosing between features and countries.

```bash
# Import the features
~$ bin/processor features

# Import a country
~$ bin/processor country ES
```

## Sponsor

Thanks to LCApps for supporting this project

![LCApps logo](./doc/asset/logo-lcapps.png)

