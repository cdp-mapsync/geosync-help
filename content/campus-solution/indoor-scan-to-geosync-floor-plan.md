+++
description = ""
lastmod = ""
subtitle = ""
title = "Indoor Scan to GeoSync Floor Plan"

+++
Process takes a merged indoor point cloud and generates 2d floor plans for use in a GeoSync Cloud Campus project.  Optionally, point clouds can be published and hosted as part of the Campus project but are not included in this topic. The process examples assume Kentucky locations.

Videos of most of the steps can be found here:

This process requires the following software packages:

1. GeoSync Z-Tools
2. QGIS Version 3.22 or later
3. GeoSync Cloud Publisher (Publisher User)

Begin with a merged indoor point cloud (LAZ format - typ: meters) - typical process might involve the use of a Geoslam scanner and Geoslam Connect software to produce an aligned and merged indoor point cloud.

**Steps**

1. **Create a GeoSync Z-Tools project** and copy the indoor point cloud into the Results folder.  Create a text file named georeference.txt - use notepad to make notes - [example](https://ztools.blob.core.windows.net/$root/georeference.txt)
2. **Create a QGIS project** - set projection to KY Single Zone - USFT and add KYAPED ArcGIS Imagery layer.  Zoom to the location of the indoor point cloud.  Recommend installing the Coordinate Capture plugin for pulling reference coordinates.
3. **Determine the first floor elevation** - Open the Z-Tools map and find the location of the indoor point cloud.  Use the right-click to estimate the first floor elevation. (NOTE: if you have other better sources for elevation then use the best value)  Record this information in the georeference.txt document.
4. **Decimate the indoor point cloud** - use the Filter LAS to decimate the point cloud to 20 percent.  Rename to merged_20pct.laz and copy to Results folder.
5. 