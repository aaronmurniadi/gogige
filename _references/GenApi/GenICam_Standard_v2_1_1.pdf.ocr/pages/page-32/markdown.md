|  GENICAM |   | ![img-47.jpeg](img-47.jpeg) emva  |
| --- | --- | --- |
|  Version 2.1.1 | Standard  |   |

</MaskedIntReg>

The numbering of the bits differs between big-endian and little-endian as is shown for a 32 bit integer below:

|  Little-Endian: | MSB | ... | LSB  |
| --- | --- | --- | --- |
|   |  31 | ... | 0  |
|  Big-Endian: | MSB | ... | LSB  |
|   |  0 | ... | 31  |

The LSB is the bit which maps to the  \( 2^{n} \)  digit. Note that with big-endian the equation MSB  \( \leq \)  LSB holds true while with little-endian the opposite holds true: LSB  \( \leq \)  MSB.

The Integer node type is used to merge the Value and the Minimum, Maximum, Increment and ValidValueSet parameters from different sources. The Integer node inherits the elements and attributes from the Node node. The value of the dynamic parameters are read form the  \( \langle pValue\rangle \)  element. The restriction parameters can be overwritten either with constants using the  \( \langle Min\rangle \) ,  \( \langle Max\rangle \) ,  \( \langle Inc\rangle \) , and  \( \langle ValidValueSet\rangle \)  elements or with pointers to other Integer nodes using the  \( \langle pMin\rangle \) ,  \( \langle pMax\rangle \) , and  \( \langle pInc\rangle \)  elements.

The <ValidValueSet> element is used to inform the uses that the integer has a limited number of valid values. This element is useful for cases like binning where the set of valid values cannot be described with an increment. The “;” character is used to mark the separation between valid values. The user only receives the value of the set that are between min and max.</ValidValueSet>

<Integer Name="Binning">
    <pValue>BinningReg</pValue>
    <ValidValuesSet>1;2;4;8;16</ValidValuesSet>
</Integer>

The Value normally comes from another node using the  \( \langle pValue\rangle \)  element. Alternatively, a constant can be given inside a  \( \langle Value\rangle \)  element. In this case, the node is a “floating” variable that can be set by the user to any value allowed by the restriction parameters. The given constant is the start value. A typical example is the following Index node that can be set to the values 0, 2, 4, ..., 254:

<Integer Name="Index">
    <Value>0</Value>
    <Min>0</Min>
    <Max>255</Max>
    <Inc>2</Inc>
</Integer>

If the  \( \langle pValue\rangle \)  element is used with an Integer node it can be optionally surrounded by any numbers of  \( \langle pValueCopy\rangle \)  elements like in the following example.

<Integer Name="Replicator">
    <pValueCopy>SomeInt1</pValueCopy>
    <pValue>SomeInt2</pValue>
    <pValueCopy>SomeInt3</pValueCopy>
    <pValueCopy>SomeInt4</pValueCopy>
</Integer>