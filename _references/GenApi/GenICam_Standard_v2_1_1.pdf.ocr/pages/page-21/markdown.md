|  GENICAM |   | ![img-33.jpeg](img-33.jpeg) emva  |
| --- | --- | --- |
|  Version 2.1.1 | Standard  |   |

- Changes in the SubMinor version number are bug fixes only; always use the file with the highest available SubMinor version number

Example: Version 1.3.0 is compatible with version 1.1.*, 1.2.* and 1.3.* (where * means don't care). It is not compatible with version 2.*.*. If version 1.3.2 is available, it should be used instead of 1.3.0.

#### 2.7.1 Versioning the Schema

The attributes SchemaMajorVersion, SchemaMinorVersion, and SchemaSubMinorVersion describe the version of the GenApi schema used for the XML file. These attributes are mandatory. They are for information purposes. In addition, the Major and Minor schema version numbers are encoded in the namespace (see xmlns entry) and the schema's file name (see xsi:schemaLocation entry).

In the example, the namespace reads “http://www.genicam.org/GenApi/Version_1_1”. A program seeking the schema file might either retrieve it over the internet using the URL or look at the file path given optionally in the second part of the schemaLocation. In the example, the path reads “.././GenApi/GenApiSchema_Version_1_1.xsd” and assumes that the XML file is stored within the folder structure of the GenICam reference implementation.

The xmlns:xsi entry “http://www.w3.org/2001/XMLSchema-instance” describes the namespace of the schema language itself.

Note that an implementation supporting, e.g., schemas up to version 1.3.* must have three schema files present: for versions 1.0.*, 1.2.*, and 1.3.*. This is required for backward compatibility – since older XML files come with an older namespace, they need older schema files. On the other hand, an XML file using a later schema version not yet supported by the implementation, say 1.4.*, needs to be rejected, hence the necessity to have the version number coded in the schema's namespace.

#### 2.7.2 Versioning the Camera Description File

The MajorVersion, MinorVersion, and SubMinorVersion attributes describe the version of XML file itself. The camera vendor is responsible for following the compatibility rules.

What does backward compatibility mean with respect to camera description files? Assume a camera that in version 1.0 has only a single feature implemented. Now assume the camera's firmware is extended to have another feature. There are two ways to deal with this situation in the camera description file. If the feature is just added to the XML file, this implicitly states that the feature is always there. Because this is not true with older cameras, the new file will not be backward compatible, and consequently it must get the version number 2.0.

A second, smarter way to deal with the situation is to introduce an inquiry register in the camera(!) where the user can check to see if the new feature is present or not. The new feature can now be added in a way that lets the user learn from the access mode of the feature whether the feature is present or not. This makes the new file backward compatible and its version number would be 1.1. Of course, this is possible only if an inquiry mechanism has been implemented in the camera from the beginning. The benefit of using the second method is that only one camera description file must be maintained for a whole family of cameras.