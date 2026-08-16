|  GENICAM |   | ![img-48.jpeg](img-48.jpeg) emva  |
| --- | --- | --- |
|  Version 2.1.1 | Standard  |   |

The GetValue method returns the value read from the pValue sub-node. The SetValue method writes to the pValue and the pValueCopy sub-nodes in the order in which they are given in the XML file. The [min, max] boundaries of the sub-nodes are combined to find the largest boundary which will fit all sub-nodes. The increments of all sub-nodes must be the same; otherwise the node becomes not writable. The ValidValueSet only considered the pValue node and ignores all pValueCopy.

The Integer node can also work like a multiplexer or a value table like shown in the following examples:

<Integer Name="Multiplexer">
    <pIndex>Selector</pIndex>
    <pValueIndexed Index="10">SomeInt1</pValue>
    <pValueIndexed Index="20">SomeInt2</pValue>
    <ValueDefault>0</ ValueDefault>
</Integer>

<Integer Name="Table">
    <pIndex>Selector</pIndex>
    <ValueIndexed Index="10">100</Value>
    <ValueIndexed Index="20">200</Value>
    <pValueDefault>SomeNode</pValueDefault>
</Integer>

The \( \text{<pIndex>} \) entry refers to an Integer node. Depending on its value one of the \( \text{<ValueIndexed>} \) or \( \text{<pValueIndexed>} \) entries is selected which behave like Value or pValue entries respectively. The two entry types can be mixed. If the index does not match any element the value given in \( \text{<ValueDefault>} \) or \( \text{<pValueDefault>} \) respectively is returned. Note that selecting an entry also forwards its properties Unit and Representation.

The <Unit> element denotes the physical meaning of a number.

The <Representation> element gives a hint about how to display the integer. If the element is Linear or Logarithmic, a slider with the appropriate behavior should be implemented. If the element is Boolean, a checkbox should be used. PureNumber means to use an edit box only with decimal display; HexNumber means the same with hexadecimal display. IPV4Address and MACAddress mean to show the numbers like an IP address (IP version 4) or a MAC address respectively.

Integer, IntReg and MaskedInt nodes can also have an  \( \langle pSelected\rangle \)  element. For a description see section 2.8.4.

#### 2.8.6 StructReg

MaskedInt node are often used to pick a field of bits from a register. If a complete MaskedInt entry is used, for each bit there is a lot of unnecessarily copied data in the camera description file because the different MaskedInt entries share most of their elements like, e.g. the <pPort> element, the <endianess> etc.</endianess></pPort>

In order to overcome this the StructReg node has been introduced. Here an example:

<StructReg Comment="VFormat7InqReg">