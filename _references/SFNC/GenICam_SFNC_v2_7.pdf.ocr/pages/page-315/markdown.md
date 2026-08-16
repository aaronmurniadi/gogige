|  GEN<i>CAM |   | emva  |
| --- | --- | --- |
|  Version 2.7.1 | Standard Features Naming Convention  |   |

# 17 Sequencer Control

The Sequencer Control chapter describes the model and features related to the control of the sequencers that can be used to change some features of the device automatically based on different events and signals.

## 17.1 Sequencer control model

The Sequencer Control Model section describes the usage of the features related to the sequencing of automatic changes to features by the device. It describes the basic Sequencer model and the typical behavior of the device when a sequencer is used.

### Sequencer overview

The purpose of a sequencer is to allow the user of a camera to define a series of feature sets for image acquisition which can consecutively be activated during the acquisition by the camera. Accordingly, the proposed sequence is configured by a list of parameter sets.

Each of these sequencer sets contains the settings for a number of camera features. Similar to user sets, the actual settings of the camera are overwritten when one of these sequencer sets is loaded. The order in which the features are applied to the camera depends on the design of the vendor. It is recommended to apply all the image related settings to the camera, before the first frame of this sequence is captured.

The sequencer sets can be loaded and saved by selecting them using SequencerSetSelector. The Execution of the sequencer is completely controlled by the device.

### Configuration of a sequencer set

The index of the adjustable sequencer set is given by the SequencerSetSelector. The number of available sequencer sets is directly given by the range of this feature.

The features which are actually part of a sequencer set are defined by the camera manufacturer. These features can be read by SequencerFeatureSelector and activated by SequencerFeatureEnable[SequencerFeatureSelector]. This configuration is the same for all Sequencer Sets.

To configure a sequencer set the camera has to be switched into configuration mode by SequencerConfigurationMode. Then the user has to select the desired sequencer set he wants to modify with the SequencerSetSelector. After the user has changed all the needed camera settings it is possible to store all these settings within a selected sequencer set by SequencerSetSave[SequencerSetSelector]. The user can also read back this settings by SequencerSetLoad[SequencerSetSelector].

To permit a flexible usage, more than one possibility to go from one sequencer set to another can exist. Such a path is selected by SequencerPathSelector[SequencerSetSelector]. Each path and therefore the transition between different sequencer sets is based on a defined trigger and an aimed next sequencer set which is selectable by SequencerSetNext[SequencerSetSelector][SequencerPathSelector]. After the trigger occurs the settings of the next set are active.