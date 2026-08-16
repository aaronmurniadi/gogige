|  ![img-39.jpeg](img-39.jpeg) CAM |   | ![img-40.jpeg](img-40.jpeg) emva  |
| --- | --- | --- |
|  Version 1.6 | GenTL Standard  |   |

This file has the information file name “tlguru_system_rev1.xml” and is located in the virtual register map starting at address 0xF0F00000 (C style notation) with the length of 0x3BF bytes.

The memory alignment is not further restricted (byte aligned) in a GenTL module. If the platform or the transport layer technology requests a certain memory alignment it has to be taken into account in the GenTL Producer implementation.

##### 4.1.2.2 Local Directory

URLs in the form “file:///filepath.extension[?SchemaVersion=1.0.0]” or “file:filename.extension[?SchemaVersion=1.0.0]” indicate that a file is present somewhere on the machine running the GenTL Consumer. This notation follows the URL definition as in the RFC 3986 for local files. Entries in italics must be replaced with the actual values, for example:

file:///C|program%20files/genicam/xml/genapi/tlguru/tlguru_system_rev1.xml?SchemaVersion=1.0.0

This would apply to an uncompressed XML file on an English Microsoft Windows operating system's C drive.

Optionally the “///” behind the “file:” can be omitted to be compatible with the GigE Vision notation. This notation does not specify the exact location. A graphical user interface could be used to determine the location using a file dialog for example.

In order to comply with some Windows notations it is also allowed to replace the ‘|’ after the drive letter with a ‘:’.

It is recommended to put the vendor, model or device and version information in the file name separated by an underscore. For example: tlguru_system_rev1 for the first version of the System module file of the GenTL Producer company TLGuru.

Supported extensions are:

- ‘xml’ for uncompressed XML description files
- ‘zip’ for zip-compressed XML description files

##### 4.1.2.3 Vendor Web Site (optional)

If a URL in the form “http://host/path/filename.extension[?SchemaVersion=1.0.0]” is present, it indicates that the XML description document can be downloaded from the vendor’s web site. This notation follows the URL definition as in the RFC 3986 for the http protocol. Entries in italics must be replaced with the actual values, e.g.,

http://www.tlguru.org/xml/tlguru_system_rev1.xml

This would apply to an uncompressed XML file found on the web site of the TLGuru company in the xml sub directory.

It is recommended to put the vendor, model or device and version information in the file name separated by an underscore. For example: tlguru_system_rev1 for the first version of the System module file of the GenTL Producer company TLGuru.

Supported extensions are: