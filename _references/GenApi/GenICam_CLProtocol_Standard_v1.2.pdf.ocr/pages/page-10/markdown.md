|  GEN<i>CAM |   | ![img-12.jpeg](img-12.jpeg) emva  |
| --- | --- | --- |
|  V1.2 | CLProtocol Standard Module  |   |

- The CLProtocol driver library checks which XML descriptions the camera can provide itself. In order to support this, the camera might implement a Manifest register as described in the GenICam GenCP standard.
- The CLProtocol driver library itself might contain suitable XML description, e.g. compiled in as Windows resource.
- The directory containing the CLProtocol driver library may contains additional XML files. The name of these files must be <XML ID>.xml, e.g.:

"SchemaVersion.1.0@MyVendor#MyFamily@XMLVersion.1.2.3.xml"

Note that the retrieval of the XML files stored on the file system is performed by the CLProtocol module of the reference implementation so the CLProtocol driver library does not have to implement that part.

If a XML ID is retrieved two immediate checks are made:

■ If the SchemaVersion cannot be handle by the GenApi version used the XML ID is discarded.
- If the DeviceID template contained in the XML ID does not match the current DeviceID the XML ID is discarded as well.

Example 1: A XML ID

"SchemaVersion.1.2@CameraManufacturer@XMLVersion.1.2.3.xml"

would be rejected by GenICam v2.0 because that version can handle only schema versions v1.0 and v1.1.

Example 2 : If the DeviceID is "MyVendor#Family1" a XML ID

"SchemaVersion.1.2@MyVendor#Family2@XMLVersion.1.2.3.xml"

would not match (wrong family) and be discarded.

Finally the list of not rejected XML IDs is sorted according to the following rules:

■ A higher SchemaVersion number goes first.
■ Within the same SchemaVersion a longer DeviceID template goes first
■ Within the same SchemaVersion and DeviceID template a higher DeviceVersion number goes first

Example:

"SchemaVersion.1.1@MyVendor#Family2@XMLVersion.1.2.0.xml"
"SchemaVersion.1.1@MyVendor#Family2@XMLVersion.1.0.0.xml"
"SchemaVersion.1.1@MyVendor@XMLVersion.3.0.0.xml"
"SchemaVersion.1.0@MyVendor@XMLVersion.3.0.0.xml"

The user can select a XML ID (possibly the fist one which is the best matching) and use this to retrieve the XML description itself. Using this description GenApi can then give access to the camera features.

### Summary

The following list summarizes the steps a client program has to take in order to retrieve an XML description for an already connected camera.

1. Retrieve a sorted list of XML IDs
2. Optionally present the list to the user to select one. The default selection is the first and – due to the sorting – best matching XML ID
3. Retrieve the XML description associated with the selected XML ID

## 5 Handling the Baud Rate

A special feature of the camera is the BaudRate. It is special because when changed it must be changed in the camera and the frame grabber at the same time; otherwise connection to the camera is lost. The frame grabber's baud rate can be changed by the CLProtocol driver library via the CLAllSerial interface which also provides means to query a list of possible baud rates supported by the grabber (see section 6.2).