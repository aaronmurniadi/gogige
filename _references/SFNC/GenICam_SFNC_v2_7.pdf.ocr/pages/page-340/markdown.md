|  GEN<i>CAM |   | emva  |
| --- | --- | --- |
|  Version 2.7.1 | Standard Features Naming Convention  |   |

## 19.2 Source Control features

This section describes the features related to the devices that support multiple video sources that are transmitted over a single link. The virtual Stream channels of a Link are used to transport the different video sources over a common physical connection.

The main feature in this section is the source selector feature (SourceSelector). This feature enables the features associated to a given video source to be controlled on a per video source basis even if they pertain to different feature categories. For instance, it can enable a user to independently set the Width feature (Image Format Control category) and the Gain feature (Analog Control category) for the two sources supported by a given device.

An example of independent features setting for a dual source device would be:

SourceSelector = Source1
Width[SourceSelector] = 320
Gain[SourceSelector] = 60
AcquisitionStart[SourceSelector]

...
AcquisitionStop[SourceSelector]

SourceSelector = Source2
Width[SourceSelector] = 240
Gain[SourceSelector] = 90
AcquisitionStart[SourceSelector]

...
AcquisitionStop[SourceSelector]

### Features selected by the Source Selector Feature

The source selector feature can be an optional selector to a number of features defined in this document based on the specificity of product it represents.

In order to simplify the standard text and feature descriptions, the optional source selector is not propagated to all the features of the SFNC that it can potentially select. Table 19-126 summarizes which features could potentially be selected by the source selector.

Table 19-126: Source Selector Potential Selectees

|  Categories | Potential Selectees  |
| --- | --- |
|  Device Control | The DeviceScanType feature can potentially be selected by the source selector feature. **Note:** The DeviceTemperatureSelector, DeviceClockSelector and DeviceSerialPortSelector features may have more enumeration  |