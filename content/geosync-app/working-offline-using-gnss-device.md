---
_template: new_doc
description: ""
lastmod: ""
subtitle: ""
title: "Working Offline using GNSS device"
---
Configure the GeoSync App to use a Trimble R1 or R2 GNSS device on a tablet or phone.  You need to the GeoSync Cloud App installed and the Trimble GNSS Status App installed.  The connection utilizes Bluetooth so the GNSS device must be paired to your phone or tablet.  See the help documentation and video for setting up a GNSS device.  

Use SBAS as a GNSS device source since the device will be used in areas without connection in the GNSS Status app.

Assuming the GNSS is connected and configured in the GNSS Status app and in the GeoSync app, perform the data collection as normal as if the device were connected, by using the following steps:

* turn on the GNSS device location 

 ![](/images/locationtool.png)

* when connected and using GNSS, location tool will have a green dot showing enabled and there will be a GPS location on the map screen (blue dot) with an accuracy label displayed

![](/images/locationtoolenabled.png)

![](/images/gpslocation.png)

(the accuracy will be better when not in an office on a computer)

* when in position to start the GPS data collection, press the compass tool (new data collection tool) 

![](/images/newdatacollection.png)

* pick the data collection report form to open

![](/images/collectionforms.png)

* add data to the report form 

![](/images/hydrantform.png)

* press save when finished entering information on the report

![](/images/savereport.png)

It is IMPORTANT to know that the symbols will NOT appear in the typical format since the device is offline.  The symbols will NOT appear on the map until the device is back within WI-FI range and a connection.  Once connected to WI-FI, the project map will update.

This can be alarming for users as they are concerned that the data is NOT being collected.  TO see the data collection points that are being collected, turn on the HISTORY POINTS.  

HISTORY POINTS will show where the data is being collected and once the device is within WI-FI range, the regular symbols will appear on the map.

Here is how to turn on the history points:

![](/images/historytool.png)

![](/images/historypoitns.png)
