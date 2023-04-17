---
_template: new_doc
---

+++
description = ""
lastmod = ""
subtitle = ""
title = "GNSS to CAD (PNEZD) Export"

+++
**Export captured points using the GeoSync app to a text file (PNEZD) ready for import into a CAD workflow.**

The export tool allows you to:

* select point number, Z, rod height, and description fields for export
* combine multiple fields for the description
* use rod heights to adjust the Z value to the ground
* adjust the Z to elevation by using the included geoid model if Z values represent height above ellipsoid (HAE).  Geoid adjustments are intended for GIS mapping applications - not for surveying or geodetic applications
* export can be either a tab or comma delimited file
* project to Northing and Easting coordinates based on the projection you set

**Steps:**

1. Open Z-Tools and pick the project you want to use for the export.  You should set the default projection for this projection to the projection you want the exported northing and easting coordinates to be in.
2. Open GIS Tools - Cloud / GIS Export Tab
   * From Cloud Layers tab refresh the cloud project list and find the cloud layer you want to export
   * Check the cloud layer (geojson) in the right panel and click the Layers download button to download the GIS layer from the cloud to your local project
   * Change the tab to GIS Data and check the GIS Layer.  Using the Geojson Import tool, import the data into your Master Layers folder
3. Switch to the Master Layers Tab
   * Click refresh layers on the left to show the list of available GIS layers
   * Select the layer you want to export and the data will be populated in the data grid on the right
   * Use the data grid and map to review the data and select the rows you want to export
   * Using the Tools - Save menu item, you can open the PNEZD tool for generating your export
4. Switch to the File Browser Tab to find the PNEZD export.  Exports are located under the \\GIS\\_pnezd folder.  Double click on the file to preview or browse to open Windows Explorer

Watch the quick video:

{{<vimeo 787192117>}}
