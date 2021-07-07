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

### Locations

A location is formed by an id, a feature, name and coordinates

## How does it work

Nowadays we can just import directly in database by choosing between features and countries.

```bash
# Import the features
~$ bin/processor features

# Import a country
~$ bin/processor country ES
```