### 17.11 SequencerSetLoad

|  Name | SequencerSetLoad[SequencerSetSelector]  |
| --- | --- |
|  Category | SequencerControl  |
|  Level | Recommended  |
|  Interface | ICommand  |
|  Access | (Read)/Write  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | -  |

Loads the sequencer set selected by SequencerSetSelector in the device. Even if SequencerMode is off, this will change the device state to the configuration of the selected set.

### 17.12 SequencerSetActive

|  Name | SequencerSetActive  |
| --- | --- |
|  Category | SequencerControl  |
|  Level | Optional  |
|  Interface | IInteger  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | ≥0  |

Contains the currently active sequencer set.

### 17.13 SequencerSetStart

|  Name | SequencerSetStart  |
| --- | --- |
|  Category | SequencerControl  |
|  Level | Recommended  |
|  Interface | IInteger  |
|  Access | Read/Write  |
|  Unit | -  |
|  Visibility | Expert  |