|  GENICAM |   | ![img-51.jpeg](img-51.jpeg) emva  |
| --- | --- | --- |
|  Version 2.1.1 | Standard  |   |

<DisplayPrecision> elements which is a non-negative number. The last two elements map to the corresponding stdio items. Here an example:</DisplayPrecision>

<float name="Exposure">
<pvalue>ExposureReg</pvalue>
<min>0.02</min>
<max>10.0</max>
<unit>ms</unit>
<representation>PureNumber</representation>
<displaynotation>Fixed</displaynotation>
<displayprecision>3</displayprecision>
</float>

As you can see in the example, <DisplayPrecision> is the amount of all digits left and right of the decimal point. This applies to all <DisplayNotation> in the same manner.</DisplayNotation>

The Float node can also work like a multiplexer or a value table like shown in the following examples:

<float name="Multiplexer">
<pindex>Selector</pindex>
<pvalueindexed index="10">SomeFloat1</pvalue>
<pvalueindexed index="20">SomeFloat2</pvalue>
<valuedefault>0</valuedefault>
</float>

<float name="Table">
<pindex>Selector</pindex>
<valueindexed index="10">100</value>
<valueindexed index="20">200</value>
<pvaluedefault>SomeNode</pvaluedefault>
</float>

The  entry refers to an Integer node. Depending on its value one of the ValueIndexed> or  entries is selected which behave like Value or pValue entries respectively. Note that selecting an entry also means forwarding its properties Unit, Representation, DisplayNotation, and DisplayPrecision.

The two entry types can be mixed. If the index does not match the value given in <ValueDefault> or <pValueDefault> respectively is returned.</pValueDefault>

A FloatReg node can be used to extract a floating point value from a byte aligned register. The FloatReg node inherits the elements and nodes of the Register node. It also has an <Endianess> element. The Length can be either 4 bytes (single precision float) or 8 bytes (double precision float). The number format has to be according to IEEE standard 754-1985.</Endianess>

#### 2.8.10 Enumeration, EnumEntry

The Enumeration node maps a name to an index value and implements the IEnumeration interface. The Enumeration node holds a list of EnumEntries with each representing a possible {name, index} pair. The Enumeration node inherits the elements and attributes of the Node node. In addition, it has either a <Value> element that represents the current index value or a <pValue> element that connects to a node with an Integer interface.</pValue></Value>