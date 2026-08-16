|  GEN<i>CAM |   | ![img-7.jpeg](img-7.jpeg) emva  |
| --- | --- | --- |
|  V1.2 | CLProtocol Standard Module  |   |

For example in order to address all cameras of a certain family the corresponding DeviceID template would looks like this:

"c:\program files\MyVendorDir#MyDriver.dll#MyVendor#MyFamily1"

A DeviceID template is said to match a DeviceID if the left part of the DeviceID string is identical to the DeviceID template.

For example the template given above would match the following DeviceIDs

"c:\program files\MyVendorDir#MyDriver.dll#MyVendor#MyFamily1#MyModelA#Version_2a#SerNo234"
"c:\program files\MyVendorDir#MyDriver.dll#MyVendor#MyFamily1#MyModelB#Version_2b#SerNo432"

but not this one

"c:\program files\MyVendorDir#MyDriver.dll#MyVendor#MyFamily2#MyModelC#Version_2a#SerNo345"

because the family is different.

#### Short DeviceID (Templates)

A short DeviceID or short DeviceID template is just an original string with the first two items – the DLL directory and file name including the trailing hash sign – missing. For example if a DeviceID template reads

"c:\program files\MyVendorDir\Win32_i86#MyDriver.dll#MyVendor#MyFamily1"

the corresponding short DeviceID is

"MyVendor#MyFamily1"

#### Probing a Device

Ideally a customer being about to setup a frame grabber port is just presented a list of all CLProtocol driver libraries installed in the system, each being represented by the corresponding manufacturer name. However it may not be possible for each driver library to fully automatically identify the camera attached to the selected port. For those cases the CLProtocol driver library provides a list of DeviceID templates for the user to select one.

For example the CLProtocol driver library of a VendorA might be able to deal with two camera families Family1 and Family2 but for example might no be able to automatically distinguish between cameras of the two families, because they implement very different protocols. In this case VendorA's CLProtocol driver library would supply the following two DeviceID templates:

"c:\program files\MyVendorDir#MyDriver.dll#VendorA#Family1"
"c:\program files\MyVendorDir#MyDriver.dll#VendorA#Family2"

A VendorB whose CLProtocol driver library can do a fully automated detection of all cameras would only supply a single DeviceID template like this:

"c:\program files\MyVendorDir#MyDriver.dll#VendorB"

A VendorC however might not bother with automatic identification altogether and just enumerates all camera models the driver library can deal with:

"c:\program files\MyVendorDir#MyDriver.dll#VendorC#Family1#ModelX"
"c:\program files\MyVendorDir#MyDriver.dll#VendorC#Family1#ModelY"
"c:\program files\MyVendorDir#MyDriver.dll#VendorC#Family2#ModelZ"

In a system were CLProtocol driver libraries from vendors A, B, and C are installed at the same time the user setting up a frame grabber port would get presented the following list of short DeviceID templates to select one:

"VendorA#Family1"
"VendorA#Family2"
"VendorB"
"VendorC#Family1#ModelX"
"VendorC#Family1#ModelY"
"VendorC#Family2#ModelZ"

After the user has selected a DeviceID template the CLProtocol driver library should be able to probe and identify the attached camera using the DeviceID template as a hint. If the identification is successful the