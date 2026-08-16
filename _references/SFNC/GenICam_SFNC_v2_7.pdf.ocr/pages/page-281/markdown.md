/* Select the Latch dual LUTs Logic Block. */
LogicBlockSelector = LogicBlock0;
LogicBlockFunction[LogicBlock0] = LatchedLUT;

/* Initialize the Value LUT of the Logic Block. */
LogicBlockLUTSelector = Value;
LogicBlockInputNumber[LogicBlock0][value] = 3;
LogicBlockInputSelector[LogicBlock0][Value] = 0;
LogicBlockInputSource[LogicBlock0][Value][0] = FrameTriggerWait;
LogicBlockInputSelector[LogicBlock0][Value] = 1;
LogicBlockInputSource[LogicBlock0][Value][1] = FrameTrigger;
LogicBlockInputSelector[LogicBlock0][Value] = 2;
LogicBlockInputSource[LogicBlock0][Value][2] = False;
LogicBlockLUTValueAll[LogicBlock0][Value][2] = 0x08;

/* Initialize the Enable LUT of the Logic Block (output always 1). */
LogicBlockLUTSelector[LogicBlock0] = Enable;
LogicBlockInputNumber[LogicBlock0][Enable] = 3;
LogicBlockInputSelector[LogicBlock0][Enable] = 0;
LogicBlockInputSource[LogicBlock0][Enable][2] = True;
LogicBlockInputSelector[LogicBlock0][Enable] = 1;
LogicBlockInputSource[LogicBlock0][Enable][2] = True;
LogicBlockInputSelector[LogicBlock0][Enable] = 2;
LogicBlockInputSource[LogicBlock0][Enable][2] = True;
LogicBlockLUTValueAll[LogicBlock0][Enable] = 0xFF;

/* Set LatchedLUTs Logic Block as source to a Counter. */
CounterSelector = Counter0;
CounterEventSource[Counter0] = LogicBlock0;
CounterEventActivation[Counter0] = RisingEdge;
CounterResetSource[Counter0] = AcquisitionStart;
CounterResetActivation[Counter0] = RisingEdge;

## 12.2 Logic Block Control features

This section describes the features that control the Logic Block.

### 12.2.1 LogicBlockControl

|  Name | LogicBlockControl  |
| --- | --- |
|  Category | Root  |
|  Level | Optional  |
|  Interface | ICategory  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Guru  |