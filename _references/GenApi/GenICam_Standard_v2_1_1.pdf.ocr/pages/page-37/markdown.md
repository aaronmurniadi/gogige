|  GENICAM |   | ![img-52.jpeg](img-52.jpeg) emva  |
| --- | --- | --- |
|  Version 2.1.1 | Standard  |   |

The following example sIf a complete MaskedInt entry hows an Enumeration describing the camera's ColorCode. If the ColorCodeReg is set to 1, for example, the camera is configured to Mono16.

<Enumeration Name="ColorCode">
    <EnumEntry Name="Mono8">
    <Value>0</Value>
    </EnumEntry>
    <EnumEntry Name="Mono16">
    <Value>1</Value>
    </EnumEntry>
    <EnumEntry Name="YUV422">
    <Value>3</Value>
    </EnumEntry>
    ColorCodeReg


</Enumeration>

<IntReg Name="ColorCodeReg">
    <Address>0x1234</Address>
    <Length>1</Length>
    <AccessMode>RW</AccessMode>
    Device
    <Sign>Unsigned</Sign>
    <Endianess>BigEndian</Endianess>
</IntReg>

Quite often, some of the EnumEntries in the list are temporarily unavailable and thus should not be presented to the user. To describe this with GenICam, you can have  and  elements in the EnumEntry sub-nodes, just as you can have with any other node.

Typically, the implementation will pre-process the camera description file and will create a separate node with the Name “EnumerationName_EnumEntryName” for each EnumEntry. Instead of the EnumEntry itself, a  element is placed in the Enumeration node. The original name of the EnumEntry is copied to the  element inside the newly created EnumEntry node. The index value represented by the EnumEntry is copied to the EnumEntry’s  element. Note that  entries must not be set manually.

Sometimes it makes sense to map a list of EnumEntries to a list of numbers. For example an Enumeration GainList could have the values {Low, Mean, High} and have a Float alias GainAbs with the values {1.0, 10.0, 100.0}. In order to express this,  elements supports a  entry which holds the float number alias of the respective entry. If the entry is not present the default NumericValue is the integer value of the EnumEntry. An Enumeration can be referenced inside the XML file by a float pointer, e.g. the  pointer of a Float node. On reading the NumericValue of the current EnumEntry is retrieved. On writing the absolute difference between the value written and the NumericValues of all writable EnumEntries are computed and the EnumEntry with the smallest absolute difference is chosen.

Enumeration nodes can also have an  element. For a description see section 2.8.4.

If an Enumeration nodes has a <PollingTime> entry the polling takes only place while the enumeration's value is set to the value of an EnumEntry which has a <IsSelfClearing> entry</IsSelfClearing></PollingTime>