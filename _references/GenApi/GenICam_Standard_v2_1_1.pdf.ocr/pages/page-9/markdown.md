|  GENICAM |   | ![img-13.jpeg](img-13.jpeg) emva  |
| --- | --- | --- |
|  Version 2.1.1 | Standard  |   |

The Gain node in Figure 3 is of the IntReg type, which allows the extraction an integer from a register. Looked at from the Root node, it is a feature of the camera. The Root node, therefore, contains a link named pFeature referencing the Gain node. To read and write the Gain registers, the Gain node needs access to the camera port, and thus it contains a link to the Device node. The link is named pPort and references the Device node.

The Gain node contains all of the information required to extract a two byte unsigned integer in BigEndian mode. The complete camera description file looks like this:

<?xml version="1.0" encoding="utf-8"?>
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
    xsi:schemaLocation="http://www.genicam.org/GenApi/Version_1_1
    http://www.genicam.org/GenApi/GenApiSchema_Version_1_1.xsd">

    <Category Name="Root">
    <ToolTip>Entry for traversing the node graph</ToolTip>
    <pFeature>Gain</pFeature>
    </Category>

    <IntReg Name="Gain">
    <ToolTip>Access node for the camera's Gain feature</ToolTip>
    <Address>0x0815</Address>
    <Length>2</Length>
    <AccessMode>RW</AccessMode>
    <pPort>Device</pPort>
    <Sign>Unsigned</Sign>
    <Endianess>BigEndian</Endianess>
    </IntReg>

    <Port Name="Device">
    <ToolTip> Port node giving access to the camera</ToolTip>
    </Port>

</RegisterDescription>

The <?xml> node is a processing element giving hints about the encoding of the file and is always the same.

The <RegisterDescription> element is the outermost bracket encapsulating all nodes of the camera. The camera is identified by the ModelName and VendorName attributes (model