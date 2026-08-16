|  GENICAM |   | ![img-49.jpeg](img-49.jpeg) emva  |
| --- | --- | --- |
|  Version 2.1.1 | Standard  |   |

<ToolTip>Inquiry register for video format 7 color codes</ToolTip>
<Address>0x14</Address>
<pAddress>VFormat7ModeCsrBase</pAddress>
<Length>4</Length>
<AccessMode>RO</AccessMode>
<pPort>Device</pPort>
<Endianess>BigEndian</Endianess>
<StructEntry Name="VFormat7Mono8InqReg">
    <ToolTip>Inquiry for ColorCode Mono8</ToolTip>
    <Bit>31</Bit>
</StructEntry>
<StructEntry Name="VFormat7YUV422InqReg">
    <ToolTip>Inquiry for ColorCode YUV8 422</ToolTip>
    <Bit>29</Bit>
</StructEntry>
<StructEntry Name="VFormat7Raw8InqReg">
    <Bit>24</Bit>
</StructEntry>
</StructReg>

The StructReg node contains the same elements as the MaskedInt element. In addition it contains one or more <StructEntry> elements which in turn can contain again the same elements as the MaskedInt element. A pre-processor replaces the StructReg node with a set of MaskedInt nodes: From each <StructEntry> element one MaskedInt node is created which gets the Name attribute from the StructEntry element, all its sub-elements, plus all elements from the StructReg node which are not present already in the <StructEntry> element. Thus the first MaskedInt node created from the example above would look like this.

<MaskedInt Name="VFormat7Mono8InqReg">
    <Address>0x14</Address>
    <pAddress>VFormat7ModeCsrBase</pAddress>
    <Length>4</Length>
    <AccessMode>RO</AccessMode>
    <pPort>Device</pPort>
    <Endianess>BigEndian</Endianess>
    <ToolTip>Inquiry for ColorCode Mono8</ToolTip>
    <Bit>31</Bit>
</MaskedInt>

Note that the <ToolTip> element was selected from the <StructEntry> element, not from the <StructReg> node. In contrast the entry with the Name VFormat7Raw8InqReg would inherit the <ToolTip> element from the <StructReg> node because it has no own. The <StructReg> element has an Comment attribute which describes it.</StructReg></ToolTip></StructEntry></991>

#### 2.8.7 Boolean

The Boolean node maps the integer value in the <OnValue> element to true and the integer value in the <OffValue> element to false. The Boolean node implements the IBoolean interface and inherits the elements and attributes from the Node node. The following example shows how to use this capacity for a Trigger node that can be displayed in a GUI as a check box:</OffValue></OnValue>