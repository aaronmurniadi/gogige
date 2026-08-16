|  LogicBlockInputNumber[LogicBlockSelector] | O | IInteger | R/(W) | - | G | Specifies the number of active signal inputs of the Logic Block.  |
| --- | --- | --- | --- | --- | --- | --- |
|  LogicBlockInputSelector[LogicBlockSelector] | O | IInteger | R/W | - | G | Selects the Logic Block's input to configure.  |
|  LogicBlockInputSource[LogicBlockSelector][LogicBlockInputSelector] | O | IEnumeration | R/W | - | G | Selects the source signal for the input into the Logic Block.  |
|  LogicBlockInputInverter[LogicBlockSelector][LogicBlockInputSelector] | O | IBoolean | R/W | - | G | Selects if the selected Logic Block Input source signal is inverted.  |
|  LogicBlockLUTIndex[LogicBlockSelector] | O | IInteger | R/W | - | G | Controls the index of the truth table to access in the selected LUT.  |
|  LogicBlockLUTValue[LogicBlockSelector][LogicBlockLUTIndex] | O | IBoolean | R/W | - | G | Read or Write the Value associated with the entry at index LogicBlockLUTIndex of the selected LUT.  |
|  LogicBlockLUTValueAll[LogicBlockSelector] | O | IInteger | R/W | - | G | Sets the values of all the output bits of the selected LUT in one access ignoring LogicBlockLUTIndex.  |
|  LogicBlockLUTSelector[LogicBlockSelector] | O | IEnumeration | R/W | - | G | Selects which of the two LUTs to configure when the selected Logic Block is a Latched dual LUTs (i.  |

## 2.11 Software Signal Control

Contains the features related to the control of the Software Signal (See the Software Signal Control chapter for details).

Table 2-11: Software Signal Control Summary

|  Name | Level | Interface | Access | Unit | Visibility | Description  |
| --- | --- | --- | --- | --- | --- | --- |
|  SoftwareSignalControl | O | ICategory | R | - | B | Category that contains the Software Signal Control features.  |
|  SoftwareSignalSelector | O | IEnumeration | R/W | - | B | Selects which Software Signal features to control.  |
|  SoftwareSignalPulse[SoftwareSignalSelector] | O | ICommand | (R)/W | - | B | Generates a pulse signal that can be used as a software trigger.  |