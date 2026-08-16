|  GENICAM |   | ![img-41.jpeg](img-41.jpeg) emva  |
| --- | --- | --- |
|  Version 2.1.1 | Standard  |   |

added to the address. Alternatively the offset can be taken from a node <pIndex pOffset="OffsetValue">. If neither Offset nor pOffset attribute is given the register's length is taken as offset.

The <pInvalidator> element contains the name of a node that when changed, will invalidate the content of this node as described in section 2.6.

The following example shows how to use this mechanism for indirect addressing (see also Figure 13):

<Integer Name="BaseAddress">
    <Value>0xff00</Value>
</Integer>

<IntReg Name="Gain">
    <Address>0x04</Address>
    <pAddress>BaseAddress</pAddress>
    <Length>4</Length>
    <AccessMode>RW</AccessMode>
    <pPort>Device</pPort>
    <Sign>Unsigned</Sign>
    <Endianess>LittleEndian</Endianess>
</IntReg>

<IntReg Name="Offset">
    <Address>0x08</Address>
    <pAddress>BaseAddress</pAddress>
    <Length>4</Length>
    <AccessMode>RW</AccessMode>
    <pPort>Device</pPort>
    <Sign>Unsigned</Sign>
    <Endianess>LittleEndian</Endianess>
</IntReg>

This example mimics a C/C++ struct of the form:

struct { // BaseAddress 0xff00
    uint32_t Reserved;
    uint32_t Gain; // Offset 0x04
    uint32_t Offset; // Offset 0x08
};

The value for the struct's base address comes from a BaseAddress constant integer node and is fed into the node using a <pAddress> element. Each element of the (Gain and Offset) struct has an offset that is added to the base address using an <Address> element.