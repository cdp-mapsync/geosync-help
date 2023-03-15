+++
description = ""
draft = true
lastmod = ""
subtitle = ""
title = "GIS Tools - Flip Direction using Reverse Flag"

+++
SQL Layer line directions can be flipped by setting the _reverse flag = yes

Requirements:

* SQL Line Layer and FORM must be setup to support _reverse flag
* GIS Manager in Z-Tools must be configured to connect to the hosted SQL database.  
* Clear Cache needs to be setup in GIS Manager in order to easily refresh the map display after line direction is flipped

Steps:

1. From GeoSync App set the _reverse flag on line segments you want to flip
2. From Z-Tools GIS Manager - Refresh the SQL Database layers list
3. Check the layer with _reverse flags
4. Run the Flip Direction tool
5. Run the Clear Cache tool