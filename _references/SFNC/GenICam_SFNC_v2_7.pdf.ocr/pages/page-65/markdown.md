|  **GEN<i>CAM** |   | emva  |
| --- | --- | --- |
|  Version 2.7.1 | Standard Features Naming Convention  |   |

Table 2-15: Sequencer Control Summary

|  Name | Level | Interface | Access | Unit | Visibility | Description  |
| --- | --- | --- | --- | --- | --- | --- |
|  SequencerControl | O | ICategory | R | - | E | Category for the Sequencer Control features.  |
|  SequencerMode | R | IEnumeration | R/W | - | E | Controls if the sequencer mechanism is active.  |
|  SequencerConfigurationMode | R | IEnumeration | R/W | - | E | Controls if the sequencer configuration mode is active.  |
|  SequencerFeatureSelector | R | IEnumeration | R/W | - | E | Selects which sequencer features to control.  |
|  SequencerFeatureEnable[SequencerFeatureSelector] | R | IBoolean | R/(W) | - | E | Enables the selected feature and make it active in all the sequencer sets.  |
|  SequencerSetSelector | R | IInteger | R/W | - | E | Selects the sequencer set to which further feature settings applies.  |
|  SequencerSetSave[SequencerSetSelector] | R | ICommand | (R)/W | - | E | Saves the current device state to the sequencer set selected by the SequencerSetSelector.  |
|  SequencerSetLoad[SequencerSetSelector] | R | ICommand | (R)/W | - | E | Loads the sequencer set selected by SequencerSetSelector in the device.  |
|  SequencerSetActive | O | IInteger | R | - | E | Contains the currently active sequencer set.  |
|  SequencerSetStart | R | IInteger | R/W | - | E | Sets the initial/start sequencer set, which is the first set used within a sequencer.  |
|  SequencerPathSelector[SequencerSetSelector] | R | IInteger | R/W | - | E | Selects to which branching path further path settings applies.  |
|  SequencerSetNext[SequencerSetSelector][SequencerPathSelector] | R | IInteger | R/W | - | E | Specifies the next sequencer set.  |
|  SequencerTriggerSource[SequencerSetSelector][SequencerPathSelector] | R | IEnumeration | R/W | - | E | Specifies the internal signal or physical input line to use as the sequencer trigger source.  |
|  SequencerTriggerActivation[SequencerSetSelector][SequencerPathSelector] | R | IEnumeration | R/W | - | E | Specifies the activation mode of the sequencer trigger.  |

## 2.16 File Access Control

Contains the features related to the generic file access of a device (See the File Access Control chapter for details).