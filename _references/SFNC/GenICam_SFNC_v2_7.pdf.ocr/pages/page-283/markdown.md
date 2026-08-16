|  Values | AND OR LUT LatchedLUT ...  |
| --- | --- |

Selects the combinational logic Function of the Logic Block to configure.

Possible values are:

- AND: Selects a Logic Block that does the logical AND of all the inputs.
- OR: Selects a Logic Block that does the logical OR of all the inputs.
- LUT: Selects a Logic Block that does a Look Up Table Transformation on all the inputs.
- LatchedLUT: Selects a Logic Block with 2 LUTs as inputs to a Flip-Flop.
- ...

### 12.2.4 LogicBlockInputNumber

|  Name | LogicBlockInputNumber[LogicBlockSelector]  |
| --- | --- |
|  Category | LogicBlockControl  |
|  Level | Optional  |
|  Interface | IInteger  |
|  Access | Read/(Write)  |
|  Unit | -  |
|  Visibility | Guru  |
|  Values | ≥2  |

Specifies the number of active signal inputs of the Logic Block.

Note that this feature might be read-only for Logical blocks that have a fixed number of inputs.

### 12.2.5 LogicBlockInputSelector

|  Name | LogicBlockInputSelector[LogicBlockSelector]  |
| --- | --- |
|  Category | LogicBlockControl  |
|  Level | Optional  |
|  Interface | IInteger  |
|  Access | Read/Write  |
|  Unit | -  |