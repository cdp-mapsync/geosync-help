---
title: Publish a GeoJSON Layer
_template: new_doc
---


### Publish Scenario

You have a local geojson layer (might be converted from some existing GIS data) and you want to create a cloud layer that can be added to a GeoSync Map.

**Keys:**

* the geojson layer must be in a geographic projection (lat,lng)
* if the geojson does not have a fully populated **muid_g** field then you must select "Generate ID" when you import the data
* if the cloud layer already has data then importing a geojson layer will append data to the cloud layer **unless you select overwrite when importing**

{{< vimeo 737554763 >}}
