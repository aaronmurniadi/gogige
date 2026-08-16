There must always be 2ⁿ entries in the LUT, where n is the number of input in the LUT. The index must be 0 based and row 0 refers to the entry with all inputs at logic level 0 (False). Index value of 2ⁿ-1 refers to the entry with all inputs at logic level 1 (True).

### 12.2.9 LogicBlockLUTValue

|  Name | LogicBlockLUTValue[LogicBlockSelector][LogicBlockLUTIndex]  |
| --- | --- |
|  Category | LogicBlockControl  |
|  Level | Optional  |
|  Interface | IBoolean  |
|  Access | Read/Write  |
|  Unit | -  |
|  Visibility | Guru  |
|  Values | True False  |

Read or Write the Value associated with the entry at index LogicBlockLUTIndex of the selected LUT.

### 12.2.10 LogicBlockLUTValueAll

|  Name | LogicBlockLUTValueAll[LogicBlockSelector]  |
| --- | --- |
|  Category | LogicBlockControl  |
|  Level | Optional  |
|  Interface | IInteger  |
|  Access | Read/Write  |
|  Unit | -  |
|  Visibility | Guru  |
|  Values | ≥0  |

Sets the values of all the output bits of the selected LUT in one access ignoring LogicBlockLUTIndex. LogicBlockLUTValueAll value can be any binary number and each bit correspond to the output value for the corresponding index (i.e. Bit 0 = LUT Index 0 output binary value).

### 12.2.11 LogicBlockLUTSelector

|  Name | LogicBlockLUTSelector[LogicBlockSelector]  |
| --- | --- |
|  Category | LogicBlockControl  |