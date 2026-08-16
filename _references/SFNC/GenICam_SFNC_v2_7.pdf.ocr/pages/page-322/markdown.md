|  Interface | IBoolean  |
| --- | --- |
|  Access | Read/(Write)  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | True False  |

Enables the selected feature and make it active in all the sequencer sets.

### 17.9 SequencerSetSelector

|  Name | SequencerSetSelector  |
| --- | --- |
|  Category | SequencerControl  |
|  Level | Recommended  |
|  Interface | IInteger  |
|  Access | Read/Write  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | ≥0  |

Selects the sequencer set to which further feature settings applies.

### 17.10 SequencerSetSave

|  Name | SequencerSetSave[SequencerSetSelector]  |
| --- | --- |
|  Category | SequencerControl  |
|  Level | Recommended  |
|  Interface | ICommand  |
|  Access | (Read)/Write  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | -  |

Saves the current device state to the sequencer set selected by the SequencerSetSelector.