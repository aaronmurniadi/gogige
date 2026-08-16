|  GEN<I>CAM |   | ![img-6.jpeg](img-6.jpeg) emva  |
| --- | --- | --- |
|  V1.2 | CLProtocol Standard Module  |   |

If the environment variable does not exist it must be created.

If the driver libraries are uninstalled the location entry must be removed from the list of locations leaving any others in place. If no other entry is left the environment variable must be deleted.

While a CLProtocol driver library is under development it can be compiled in Debug or Release mode. In order to simplify the life of the developer the Debug version of a driver library named XXX.dll/XXX.so should be named XXX.debug.dll/XXX.debug.so. The following rules apply when the driver libraries are enumerated by the CLProtocol module (CCLPort wrapper class):

- If the CLProtocol module is compiled in Debug mode a driver library named XXX.dll/XXX.so is loaded only if there is no corresponding driver library named XXX.debug.dll/XXX.debug.so in the same directory.
- If the CLProtocol module is compiled in Release mode a driver library named XXX.debug.dll/XXX.debug.so is loaded only if there is no corresponding driver library named XXX.dll/XXX.so in the same directory.

## 3 Selecting a CLProtocol Driver Library and Identifying a Camera

The key problem when setting up the CLProtocol driver library is to identify the manufacturer and model name of the camera connected to a frame grabber port. This information is required in order to select the right CLProtocol driver library but it is also required by the CLProtocol driver library itself for adapting its behavior to different camera models of the same vendor.

It would be nice if the manufacturer name as well as the model name of an arbitrary Camera Link camera could be determined automatically just by probing the frame grabber port. However, this kind of plug&play mechanism will stay a dream for Camera Link because for historical reasons there is no standard protocol for the serial port of Camera Link cameras and cameras of different vendors can behave very differently. Probing a camera with different protocol variants would take too long and could even drive some camera models in an undefined state from which they might not recover.

So there is no way around the user selecting at least the camera manufacturer name and thus the CLProtocol driver library for each frame grabber port manually. After that has been done the CLProtocol driver library can identify the camera in a more or less automatic way because the vendors should know their cameras well enough in order to automate that task. Nevertheless, if for some reason that automation is not possible the standard provides means to deal with that situation, too.

The whole identification process is based on string identifiers (IDs) which are enumerated by the system and (partially) selected by the customers.

### DeviceID

The identifier resulting from the camera identification process is called the DeviceID. It contains all data required to uniquely identify a device and its corresponding CLProtocol driver library. This data is assembled in a string which is composed of tokens separated by the hash ('#') sign:

"DriverDirectory#DriverFileName#Manufacturer#Family#Model#Version#SerialNumber"

The first two tokens describe the directory where the CLProtocol driver library is found (without trailing back slash) and the file name of the driver library. The other tokens are from left to right the camera's manufacturer, family, model, version, and serial number. Each of these latter tokens must follow the naming convention for C variables, i.e. they must match the following regular expression:

[a-zA-Z_] [a-zA-Z0-9_]*

Either the serial number or the serial number and the version token can be omitted. Here two examples for valid DeviceIDs:

"c:\program files\MyVendorDir\Win32_i86#MyDriver.dll#MyVendor#MyFamily1#MyModelA#Ver_2a#SerNo123"
"c:\program files\MyVendorDir\Win32_i86#MyDriver.dll#MyVendor#MyFamily1#MyModelA"

### DeviceID Templates

In order to address a subset of possible DeviceIDs a DeviceID template can be formed by the DeviceID from the right up to but not including the manufacturer name.