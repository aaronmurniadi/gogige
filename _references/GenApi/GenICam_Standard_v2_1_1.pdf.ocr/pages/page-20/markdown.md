|  GENICAM |   | ![img-32.jpeg](img-32.jpeg) emva  |
| --- | --- | --- |
|  Version 2.1.1 | Standard  |   |

cache for ImagerHeightReg will be invalidated when the content of the BinningReg node changes, a  link must be introduced between the two nodes. The sole purpose of this link is to document the hidden dependency between the two features and to make sure that the cache is always coherent.

### 2.7 Identifying and Versioning a Camera Description File

It must be possible to identify a camera description file, and thus the described camera, in a unique manner. In addition, a camera description file will typically evolve over time, e.g., when features are added to the corresponding camera product. This creates the necessity for a versioning mechanism. The GenApi syntax itself will also evolve over time, e.g., when new node types are added, thus a versioning mechanism for the schema is also required.

The necessary means are found in the attribute list of the <RegisterDescription> element, which is the outermost bracket of the XML file. Here is an example:</RegisterDescription>

<RegisterDescription
    ModelName="Example01"
    VendorName="Test"
    ToolTip="Example 01 from the GenApi standard"
    StandardNameSpace="None"
    SchemaMajorVersion="1"
    SchemaMinorVersion="1"
    SchemaSubMinorVersion="0"
    MajorVersion="1"
    MinorVersion="0"
    SubMinorVersion="0"
    ProductGuid="1F3C6A72-7842-4edd-9130-E2E90A2058BA"
    VersionGuid="7645D2A1-A41E-4ac6-B486-1531FB7BECE6"
    xmlns="http://www.genicam.org/GenApi/Version_1_1"
    xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
    xsi:schemaLocation="http://www.genicam.org/GenApi/Version_1_1../GenApiSchema_Version_1_1.xsd">

The camera described is identified by the VendorName: / ModelName pair. Assuming that vendor names are mutually exclusive due to trade marks, this scheme creates unique names. The ToolTip attribute is used to provide additional information about the device that can be displayed to the user, e.g., in a selection list of devices found on a bus.

The attribute StandardNameSpace is deprecated (for details see section 2.8.1).

The versioning of the different items in a camera description file follows common rules, and a three part version number is used:

<Major>.<Minor>.<SubMinor>

An example would be '1.4.2'.

The following compatibility rules apply:

- Files with a higher Major version number are not backward compatible
- Files with a higher Minor version number are backward compatible