- Stream0TransferStopped, Stream1TransferStopped, ...: Transfer on the stream is stopped.
- Stream0TransferOverflow, Stream1TransferOverflow, ...: Transfer on the stream is in overflow.

### 12.2.7 LogicBlockInputInverter

|  Name | LogicBlockInputInverter[LogicBlockSelector][LogicBlockInputSelector]  |
| --- | --- |
|  Category | LogicBlockControl  |
|  Level | Optional  |
|  Interface | IBoolean  |
|  Access | Read/Write  |
|  Unit | -  |
|  Visibility | Guru  |
|  Values | True False  |

Selects if the selected Logic Block Input source signal is inverted. This feature is not available when the LogicBlockInputSource is set to True or False.

Note: When applied to a clock input, if LogicBlockInputInverter is set to False, this corresponds to a rising edge clock activation and if set to True to a falling edge activation.

Possible values are:

- True: The Logic Block Input is inverted.
- False: The Logic Block Input is not inverted.

### 12.2.8 LogicBlockLUTIndex

|  Name | LogicBlockLUTIndex[LogicBlockSelector]  |
| --- | --- |
|  Category | LogicBlockControl  |
|  Level | Optional  |
|  Interface | IInteger  |
|  Access | Read/Write  |
|  Unit | -  |
|  Visibility | Guru  |
|  Values | ≥0  |

Controls the index of the truth table to access in the selected LUT.