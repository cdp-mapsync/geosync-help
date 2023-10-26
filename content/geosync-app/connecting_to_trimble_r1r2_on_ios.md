---
title: Connecting to Trimble R1/R2 on iOS
---

In order to use a Trimble R1/R2 with GeoSync Cloud, the Trimble Mobile Manager app will need to be installed. Below is a link to the app on the App Store.

[https://apps.apple.com/us/app/trimble-mobile-manager/id1477956695](https://apps.apple.com/us/app/trimble-mobile-manager/id1477956695)

Connecting Receiver to Trimble Mobile Manager

1. Press the power button on the receiver until the LED starts blinking.
2. To pair with the receiver from an:
   * iOS device, go to the iOS Bluetooth settings screen.
   * Android device, in Trimble Mobile Manager tap  and tap Position Source. In the GNSS receiver type field, select Bluetooth Receiver and then tap List Connected Devices.
3. The device scans for nearby devices. Receivers are listed by model name and serial number. If you are not sure which receiver is yours, check the serial number label on the underside of the receiver.
4. If the receiver is not listed, tap the Refresh icon to rescan. If there are a number of Bluetooth devices in the vicinity you may need to refresh the list a couple of times.
5. If you are prompted to enter a PIN when connecting to a Trimble R series receiver, the default PIN is 0000. Note however that this may have been changed for your receiver in the web interface used to configure the receiver settings.
6. If you have connected to the receiver using the iOS Bluetooth settings screen but it is not listed in Trimble Mobile Manager, return to the iOS Bluetooth settings screen and turn off Bluetooth and then turn on Bluetooth. Return to Trimble Mobile Manager and tap the List Connected Devices button again.
7. Once connected, the Trimble Mobile Manager Home screen shows the name and serial number of the connected receiver.

Using Receiver Position in GeoSync Cloud

1. Once positions are being received in the Trimble Mobile Manager app, a connection can be made from the GeoSync Cloud app.
2. With a GeoSync Cloud map open, go to the settings in the bottom left corner.
3. On the settings window, change the location source to Trimble Websocket and click save.
4. Turn on geolocation and positions should start appearing in the app. To verify location data, use the location tab in the left menu panel.
