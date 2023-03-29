+++
description = ""
draft = true
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

Begin with a merged indoor point cloud (LAZ format) - typical process might involve the use of a Geoslam scanner and Geoslam Connect software to produce an aligned and merged indoor point cloud.

**Steps**

1. Create a GeoSync Z-Tools project and copy the indoor point cloud into the Results folder.  Create a text file named georeference.txt - use notepad to make notes - [example](https://ztools.blob.core.windows.net/$root/georeference.txt)
2. 