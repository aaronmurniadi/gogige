|  GENICAM |   | ![img-46.jpeg](img-46.jpeg) emva  |
| --- | --- | --- |
|  Version 2.1.1 | Standard  |   |

<IntReg Name="Gain">
    <Address>0x1234</Address>
    <Length>2</Length>
    <AccessMode>RW</AccessMode>
    <pPort>Device</pPort>
    <Sign>Unsigned</Sign>
    <Endianess>BigEndian</Endianess>
</IntReg>

For example, the default [Minimum, Maximum] range of an integer is created from the native Minimum / Maximum of its number representation. For an <Integer> node (int64_t) this would be [INT64_MIN, INT64_MAX]. For the <IntReg> node shown above (uint16_t) this would be [0, UINT16_MAX].</IntReg>

The <Sign> element can have the value Singed or Unsigned. Since all GenApi integer values are internally represented as signed 64-bit, register with less than 64-bit are automatically convert to a signed int64. The Sign element is used to manage sign bit extension while converting a register to a value. Unsigned int64 registers are not available since they cannot be converted to unsigned int64. The largest unsigned integer accessible with a MaskedIntReg (see below) equals 2^63-1.

The <Endianess> element can have the values LittleEndian or BigEndian and refers to the endianess of the device as seen through the transport layer. The transport layer must attempt to not change the endianess. Note that the implementation must be aware of whether it is running itself on a little-endian or big-endian machine.</Endianess>

Sometimes integers are not byte aligned, but are packed into a register. In this case, a MaskedIntReg is used. It inherits the elements and attributes from the Register node. The following XML code is an example for a 12 bit integer packed into a 2 byte register. The <LSB> and <MSB> elements denote the least significant bit and the most significant bit respectively.</MSB>

<MaskedIntReg Name="Offset">
    <Address>0x2345</Address>
    <Length>2</Length>
    <AccessMode>RW</AccessMode>
    <pPort>Device</pPort>
    <LSB>11</LSB>
    <MSB>0</MSB>
    <Sign>Unsigned</Sign>
    <Endianess>BigEndian</Endianess>
</MaskedIntReg>

In the case where only a single bit must be mapped – which is quite common for presence inquiry bits – instead of using an <LSB> and an <MSB> element with the same value, you can also use a <Bit> entry.</Bit>

<MaskedIntReg Name="OffsetInq">
    <Address>0x2345</Address>
    <Length>2</Length>
    <AccessMode>RW</AccessMode>
    <pPort>Device</pPort>
    <Bit>15</Bit>
    <Sign>Unsigned</Sign>
    <Endianess>BigEndian</Endianess>